package dal

import (
	"encoding/csv"
	"errors"
	"io"
	"ip2loc/consts"
	"ip2loc/models"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type CSVGeoIPRepository struct {
	Reader io.Reader
}

func NewCSVGeoIPRepository(reader io.Reader) *CSVGeoIPRepository {
	return &CSVGeoIPRepository{Reader: reader}
}

func (r *CSVGeoIPRepository) GetGeoIP(ip string) (models.GeoIPData, error) {
	var csvReader *csv.Reader
	if r.Reader == nil {
		filePath := viper.GetString("datastore.filePath")
		if filePath == "" {
			return models.GeoIPData{}, errors.New(consts.MissingCSVPath)
		}

		absPath, err := filepath.Abs(filePath)
		if err != nil {
			return models.GeoIPData{}, err
		}

		f, err := os.Open(absPath)
		if err != nil {
			return models.GeoIPData{}, err
		}

		defer f.Close()

		csvReader = csv.NewReader(f)
	} else {
		csvReader = csv.NewReader(r.Reader)
	}

	rowCounter := 1
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return models.GeoIPData{}, err
		}
		if len(rec) < 3 || len(rec) > 3 {
			log.Warn().Int("Row number", rowCounter).Msg("Row contains wrong number of values")
		}
		if rec[0] == ip {
			return models.GeoIPData{
				IP:      rec[0],
				City:    rec[1],
				Country: rec[2],
			}, nil
		}
	}
	return models.GeoIPData{}, errors.New(consts.IPNotFound)
}
