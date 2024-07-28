package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sawilkhan/analytics/models"
)

func GetLocation(ip string) (models.Location, error) {
	apiUrl := fmt.Sprintf("https://api.ipgeolocation.io/ipgeo?apiKey=f2ea611e79f14c9795576a52b3e6a048&ip=%s&fields=geo", ip)
	res, err := http.Get(apiUrl)
	if err != nil{
		return models.Location{}, err
	}
	defer res.Body.Close()

	var location models.Location
	if err := json.NewDecoder(res.Body).Decode(&location); err != nil{
		return models.Location{}, err
	}
	return location, nil
}