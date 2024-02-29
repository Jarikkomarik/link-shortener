package middleware

import (
	"log"
	"net/http"

	"com.jarikkomarik.linkshortener/myError"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err, ok := recover().(error); ok == true {

				log.Printf("Panic occured : %s", err.Error())

				switch error := err.(type) {
				case myError.InvalidURLError:
					c.AbortWithStatusJSON(http.StatusBadRequest, error.Error())
				case myError.InvalidRecordId:
					c.AbortWithStatusJSON(http.StatusBadRequest, error.Error())
				default:
					c.AbortWithStatusJSON(http.StatusInternalServerError, error.Error())
				}
			}

		}()
		c.Next()
	}
}
