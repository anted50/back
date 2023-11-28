package main

import (
	"test/controller"
	"test/initializers"
	"test/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func init() {
	initializers.LoadEnvVariables()
	err := initializers.InitGorm()
	if err != nil {
		panic(err)
	}
}

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
	router.GET("/validate", middleware.RequireAuth, controller.Validate)

	router.POST("/get", controller.Get)
	router.GET("/list/:id", controller.Paginate)

	router.Run()
}
