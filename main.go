package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	connStr := "//TODO: parse asp.net config"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	todoapi := &todoapi{db}

	router := gin.Default()

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.String(http.StatusNotFound, "")
	})

	todoapi.MapTodoApi(router)

	router.Run(":8080")
}
