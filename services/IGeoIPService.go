package services

import "ip2loc/models"

type IGeoIPService interface {
	FindIPGeolocation(ip string) (models.GeoIPData, error)
}
