package main

import (
	"github.com/gin-gonic/gin"
	"github.com/teste/scr/controller/routes"
)

func main() {
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	router.Run(":8080")
}
