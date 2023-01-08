package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/quiz-api/controller"
)

func QuestionRoute(router *gin.RouterGroup) {
	router.GET("/", controller.Question())
}
