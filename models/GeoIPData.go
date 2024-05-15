package models

type GeoIPData struct {
	Id      interface{} `json:"-"`
	IP      string      `json:"ip"`
	City    string      `json:"city"`
	Country string      `json:"country"`
}
