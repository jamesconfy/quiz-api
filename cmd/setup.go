package cmd

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/quiz-api/routes"
)

func Setup() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	v1 := router.Group("/api/v1")
	v1.Use(gin.Logger())
	v1.Use(gin.Recovery())

	routes.QuestionRoute(v1)
	routes.ErrorRoute(router)

	router.Run(":" + port)
}
