package v1

import (
	"net/http"

	"ip2loc/services"

	"github.com/gin-gonic/gin"
)

type FindCountryController struct {
	geoIPService services.IGeoIPService
}

func NewFindCountryController(service services.IGeoIPService) *FindCountryController {
	return &FindCountryController{
		geoIPService: service,
	}
}

func (qc *FindCountryController) FindCountry(c *gin.Context) {
	ip := c.Query("ip")

	if ip == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	result, err := qc.geoIPService.FindIPGeolocation(ip)

	if err != nil {
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}
