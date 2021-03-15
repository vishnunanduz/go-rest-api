package service

import "net/http"

const carServiceUrl string = "https://myfakeapi.com/api/cars/1"

type CarService interface {
	FetchData()
}

type fetchCarData struct {
}

func NewCarService() CarService {
	return &fetchCarData{}
}

func (*fetchCarData) FetchData() {
	client := http.Client{}
	resp, _ := client.Get(carServiceUrl)

	carDataChannel <- resp

}
