package utils

import (
	"log"

	"github.com/sawilkhan/analytics/models"
)

func LogRequest(ip string, location models.Location){
	log.Printf("IP: %s, Country: %s, City: %s, Latitude: %s, Longitude: %s",
        ip, location.CountryName, location.City, location.Latitude, location.Longitude)
}