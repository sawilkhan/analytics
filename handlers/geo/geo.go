package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sawilkhan/analytics/services"
	"github.com/sawilkhan/analytics/utils"
)

func GeoHandler(w http.ResponseWriter, r *http.Request){
	ip := r.RemoteAddr
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != ""{
		ip = strings.Split(forwarded, ",")[0]
	}

	location, err := services.GetLocation(ip)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.LogRequest(ip, location)
	response := map[string]string{
		"status": "success",
	}
	
	json.NewEncoder(w).Encode(response)	
}