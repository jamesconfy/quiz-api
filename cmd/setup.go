package cmd

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/quiz-api/docs"
	"github.com/quiz-api/middleware"
	"github.com/quiz-api/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())

	v1 := router.Group("/api/v1")

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.QuestionRoute(v1)
	routes.ErrorRoute(router)

	router.Run(":" + port)
}
