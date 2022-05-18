package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"apiserver/configs"
	"apiserver/models"
	"apiserver/router"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	debugMode = true
)

func init() {
	// set log level
	if debugMode {
		log.SetLevel(log.DebugLevel)
		gin.SetMode(gin.DebugMode)
	} else {
		log.SetLevel(log.InfoLevel)
		gin.SetMode(gin.ReleaseMode)
	}
	format := &log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	}
	log.SetFormatter(format)
}

func main() {
	config := configs.LoadConfig()

	db, err := models.ConnectDB(config.Database)
	if err != nil {
		log.Fatal("Database connection failed")
	} else {
		log.Info("connection success")
	}

	router := router.NewRoutes(db)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Server.Port),
		Handler: router,
	}

	// gracefully shutdown the server & database connection
	go handleShutdown(db, server)

	// run the http server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %v\n", err)
	}
}

func handleShutdown(db *sql.DB, srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	// handle ctrl+c event here
	// for example, close database
	log.Warn("Closing DB connection and http server before complete shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// wait 5 second for all service to process and then shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Error: %v\n", err)
	}
	if err := db.Close(); err != nil {
		log.Fatalf("Database Connection Close Failed: %v\n", err)
	}
	os.Exit(0)
}
