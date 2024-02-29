package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"com.jarikkomarik.linkshortener/config"
	"com.jarikkomarik.linkshortener/controller"
	"com.jarikkomarik.linkshortener/middleware"
	"com.jarikkomarik.linkshortener/repository"
	"com.jarikkomarik.linkshortener/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//setupDNS() //remove when deployed

	setupEnv()

	ginInstance := gin.Default()
	repo := repository.NewLinkRepository(config.DatabaseConnection())
	service := service.NewLinkShortenerService(repo)
	controller := controller.NewLinkShortenerController(service)

	ginInstance.Use(middleware.ErrorHandler())

	ginInstance.POST("/", controller.HandlePost)

	ginInstance.GET("/:id", controller.HandleGet)

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

func setupDNS() {
	// Change the default DNS resolver in case of issue connect to Mongo
	resolver := &net.Resolver{
		PreferGo: true, // Use Go's DNS resolver implementation
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			// Customize DNS resolution here if needed
			// For example, you can use a custom DNS server
			return (&net.Dialer{}).DialContext(ctx, network, "8.8.4.4:53")
		},
	}

	net.DefaultResolver = resolver
}

func setupEnv() {
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

}
