package api

import (
	"github.com/gin-gonic/gin"
)

func ApiRoute(router *gin.Engine) {
	router.GET("/user/:id", GetUser)
	router.POST("/user", PostUser)
	router.DELETE("/user/:id", DeleteUser)
}
