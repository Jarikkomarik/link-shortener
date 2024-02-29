package controller

import (
	"log"
	"net/http"

	"com.jarikkomarik.linkshortener/service"
	"github.com/gin-gonic/gin"
)

type LinkShortenerController struct {
	service *service.LinkShortenerService
}

func NewLinkShortenerController(service *service.LinkShortenerService) *LinkShortenerController {
	return &LinkShortenerController{service: service}
}

func (controller *LinkShortenerController) HandlePost(ctx *gin.Context) {
	url := ctx.Query("url")
	host := ctx.Request.Host

	value := controller.service.RegisterNewUrl(url, host)
	log.Println("Sucsessfully registered new url")

	ctx.JSON(http.StatusOK, value)
}

func (controller *LinkShortenerController) HandleGet(ctx *gin.Context) {
	id := ctx.Param("id")
	url := "https://" + controller.service.GetOriginalUrl(id)
	log.Println("Sucsessfully forwared to the original resource")
	ctx.Redirect(http.StatusFound, url)
}
