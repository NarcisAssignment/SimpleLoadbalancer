package main

import (
	"io"
	"os"
	"simpleLoadbalancer/controller"
	"simpleLoadbalancer/model"
	"simpleLoadbalancer/routers"

	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	gin.DisableConsoleColor()
	f, _ := os.Create(os.Getenv("LOG_FILE"))
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()
	r.Use(ginprom.PromMiddleware(nil))

	r.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
	r.GET("/health", routers.HealthGET)

	model.ConnectDatabase()
	r.GET("/books", controller.FindBooks)
	r.GET("/books/:id", controller.FindBook)

	r.Run()
}
