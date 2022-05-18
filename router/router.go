package router

import (
	"apiserver/controller"

	"database/sql"

	"github.com/gin-gonic/gin"
)

func NewRoutes(db *sql.DB) *gin.Engine {
	router := gin.Default()

	repo := controller.NewRepo(db)

	router.GET("/api/convergence/getkpsn/:csn", repo.GetKPSN)
	router.GET("/api/convergence/getreverse/:csn", repo.GetReverse)

	return router

}
