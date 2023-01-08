package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/quiz-api/controller"
)

func ErrorRoute(router *gin.Engine) {
	router.NoRoute(controller.Error404())
}
