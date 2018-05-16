package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wongoo/go-app-example/api"
)

func main() {
	router := gin.Default()

	api.ApiRoute(router)

	router.Run(":8082")
}
