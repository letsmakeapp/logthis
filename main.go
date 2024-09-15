package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())

	r.POST("/log", func(c *gin.Context) {
		type Request struct {
			Content string `json:"content"`
		}

		var req Request
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(req.Content)
		c.Status(http.StatusNoContent)
	})
	if err := r.Run(fmt.Sprintf("0.0.0.0:%s", port)); err != nil {
		panic(err)
	}
}
