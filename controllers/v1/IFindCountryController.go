package v1

import (
	"github.com/gin-gonic/gin"
)

type IFindCountryController interface {
	FindCountry(c *gin.Context)
}
