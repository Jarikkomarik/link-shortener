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

// HandlePost godoc
// @Summary Register new URL
// @Description Returns id that reffers to the original URL.
// @Param url query string true "URL to be shortened"
// @Success 200 {string} string "Successfully registered new URL"
// @Router /shorten [post]
func (controller *LinkShortenerController) HandlePost(ctx *gin.Context) {
	url := ctx.Query("url")
	host := ctx.Request.Host

	value := controller.service.RegisterNewUrl(url, host)
	log.Println("Sucsessfully registered new url")

	ctx.JSON(http.StatusOK, value)
}

// HandleGet godoc
// @Summary Redirect to original URL
// @Description Redirects to the original URL based on the provided ID
// @Param id path string true "ID of the shortened URL"
// @Success 302 {string} string "Successfully redirected to the original resource"
// @Router /{id} [get]
func (controller *LinkShortenerController) HandleGet(ctx *gin.Context) {
	id := ctx.Param("id")
	url := "https://" + controller.service.GetOriginalUrl(id)
	log.Println("Sucsessfully forwared to the original resource")
	ctx.Redirect(http.StatusFound, url)
}
