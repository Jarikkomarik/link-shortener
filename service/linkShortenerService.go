package service

import (
	"fmt"
	"log"
	"net/url"

	"com.jarikkomarik.linkshortener/myError"
	"com.jarikkomarik.linkshortener/repository"
)

type LinkShortenerService struct {
	repo *repository.LinkRepository
}

func NewLinkShortenerService(repo *repository.LinkRepository) *LinkShortenerService {
	return &LinkShortenerService{repo: repo}
}

func (service *LinkShortenerService) RegisterNewUrl(url, host string) string {
	if !isValidURL(url) {
		log.Println("Failed to register new url, url is invalid")
		panic(myError.InvalidURLError{})
	}
	id, err := service.repo.GetUrlId(url)
	if err != nil {
		id = service.repo.InsertRecord(url)
	}
	return fmt.Sprint(host, "/", id)

}

func (service *LinkShortenerService) GetOriginalUrl(id string) string {
	return service.repo.GetRecord(id)
}

func isValidURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}
	return true
}
