package models

type Location struct {
	IP           string `json:"ip"`
	CountryCode2 string `json:"country_code2"`
	CountryCode3 string `json:"country_code3"`
	CountryName  string `json:"country_name"`
	StateProv    string `json:"state_prov"`
	District     string `json:"district"`
	City         string `json:"city"`
	Zipcode      string `json:"zipcode"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
}