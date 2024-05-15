package dal_test

import (
	"errors"
	"ip2loc/consts"
	"ip2loc/dal"
	"ip2loc/models"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockCSVReader struct {
	Records [][]string
	Err     error
}

func (m *MockCSVReader) Read(p []byte) (n int, err error) {
	if len(m.Records) == 0 {
		return 0, m.Err
	}
	record := strings.Join(m.Records[0], ",") + "\n"
	copy(p, record)
	m.Records = m.Records[1:]
	return len(record), nil
}

func TestCSVGeoIPRepository_GetGeoIP(t *testing.T) {
	tests := []struct {
		name           string
		ip             string
		records        [][]string
		expectedResult models.GeoIPData
		expectedError  error
	}{
		{
			name: "Valid IP found",
			ip:   "192.168.1.1",
			records: [][]string{
				{"192.168.1.1", "New York", "United States"},
				{"10.0.0.1", "London", "United Kingdom"},
			},
			expectedResult: models.GeoIPData{
				IP:      "192.168.1.1",
				City:    "New York",
				Country: "United States",
			},
			expectedError: nil,
		},
		{
			name:           "IP not found",
			ip:             "192.168.1.2",
			records:        [][]string{},
			expectedResult: models.GeoIPData{},
			expectedError:  errors.New(consts.IPNotFound),
		},
		{
			name: "Invalid CSV format",
			ip:   "192.168.1.1",
			records: [][]string{
				{"192.168.1.1", "New York"},
				{"10.0.0.1", "London", "United Kingdom"},
			},
			expectedResult: models.GeoIPData{},
			expectedError:  errors.New("csv: wrong number of fields in line 1"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &MockCSVReader{Records: tt.records}
			repo := dal.NewCSVGeoIPRepository(reader)

			result, err := repo.GetGeoIP(tt.ip)

			assert.Equal(t, tt.expectedResult, result)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
