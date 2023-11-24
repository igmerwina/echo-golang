package main

import (
	"dbgo/config"
	"dbgo/controller"
	auth "dbgo/middlerware"
	"os"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "dbgo/docs"
)

// @title API Documentation Employee
// @version 1
// @description Contoh Deskcripsi
// @host localhost:1323
// @BasePath /
func main() {
	config.Connect()

	e := echo.New()

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:" + config.CSRFTokenHader,
		ContextKey:  config.CSRFKey,
	}))

	e.GET("/index", controller.Index)
	e.POST("/sayhello", controller.SayHello)

	emp := e.Group("/emp")
	emp.Use(auth.Authentication())
	emp.PUT("/", controller.UpdateEmployee)
	emp.DELETE("/:id", controller.DeleteEmployee)

	item := e.Group("/item")
	item.Use(auth.Authentication())
	item.POST("/", controller.CreateItem)

	// login
	e.POST("/login", controller.Login)
	e.POST("/emp", controller.CreateEmployee)

	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	var PORT = os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start(":" + PORT))
}
