package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:  []string{"Content-Type", "Accept", "Origin", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))
	InitializeRoutes(r)

	r.Run(":8080")
}
