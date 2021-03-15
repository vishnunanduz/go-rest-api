package controller

import (
	"encoding/json"
	"net/http"
	"rest/service"
)

var (
	carDetailsService service.CarDetailsService
)

type controller struct{}

type CarDetailsController interface {
	GetCarDetails(w http.ResponseWriter, r *http.Request)
}

func NewCarDetailsController(service service.CarDetailsService) CarDetailsController {
	carDetailsService = service
	return &controller{}
}

func (*controller) GetCarDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	result := carDetailsService.GetDetails()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
