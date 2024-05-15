package dal

import "ip2loc/models"

type IGeoIPRepository interface {
	GetGeoIP(ip string) (models.GeoIPData, error)
}
