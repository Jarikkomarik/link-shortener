package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"com.jarikkomarik.linkshortener/config"
	"com.jarikkomarik.linkshortener/controller"
	"com.jarikkomarik.linkshortener/docs"
	"com.jarikkomarik.linkshortener/middleware"
	"com.jarikkomarik.linkshortener/repository"
	"com.jarikkomarik.linkshortener/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 	URL shortener API
// @version	1.0
// @description URL shortener API
// @host 	localhost:8080
func main() {
	setupEnv()
	docs.SwaggerInfo.BasePath = ""

	ginInstance := gin.Default()
	repo := repository.NewLinkRepository(config.DatabaseConnection())
	service := service.NewLinkShortenerService(repo)
	controller := controller.NewLinkShortenerController(service)

	ginInstance.Use(middleware.ErrorHandler())

	ginInstance.POST("/shorten", controller.HandlePost)

	ginInstance.GET("/:id", controller.HandleGet)

	ginInstance.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server := &http.Server{
		Addr:           ":" + os.Getenv("SERVER_PORT"),
		Handler:        ginInstance,
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   2 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Starting server on port %s...", os.Getenv("SERVER_PORT"))

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
		panic(err)
	}

}

func setupEnv() {
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

}
