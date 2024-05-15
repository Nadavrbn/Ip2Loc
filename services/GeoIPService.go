package services

import (
	"errors"
	"ip2loc/consts"
	"ip2loc/dal"
	"ip2loc/models"

	"github.com/spf13/viper"
)

const (
	CSV = "CSV"
)

type GeoIPService struct {
}

func NewGeoIPService() *GeoIPService {
	return &GeoIPService{}
}

func (s *GeoIPService) FindIPGeolocation(ip string) (models.GeoIPData, error) {
	var repository dal.IGeoIPRepository

	storageType := viper.GetString("datastore.type")
	switch storageType {
	case CSV:
		{
			repository = dal.NewCSVGeoIPRepository(nil)
		}
	case "":
		return models.GeoIPData{}, errors.New(consts.MissingDataStoreType)
	default:
		return models.GeoIPData{}, errors.New(consts.UnsupportedDataStore)
	}

	return repository.GetGeoIP(ip)
}
