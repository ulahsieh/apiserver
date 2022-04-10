package controller

import (
	"database/sql"
	"net/http"

	"apiserver/models"

	"github.com/gin-gonic/gin"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) GetKPSN(c *gin.Context) {
	csn := c.Param("csn")
	md := models.NewRepo(r.db)
	res, err := md.Find(csn)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Action Failed")
		return
	}

	if res == nil {
		c.JSON(http.StatusNotFound, "Record not found")
		return
	}

	c.JSON(http.StatusOK, res)
}
