package service

import "net/http"

const ownerServiceUrl string = "https://myfakeapi.com/api/users/1"

type OwnerService interface {
	FetchOwnerData()
}

type fetchOwnerData struct {
}

func NewOwnerService() OwnerService {
	return &fetchOwnerData{}
}

func (*fetchOwnerData) FetchOwnerData() {
	client := http.Client{}
	resp, _ := client.Get(ownerServiceUrl)

	ownerDataChannel <- resp

}
