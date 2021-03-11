package main

import (
	"rest/controller"
	router "rest/http"
	"rest/service"
)

var (
	carDetailsService    service.CarDetailsService       = service.NewCarDetailsService()
	carDetailsController controller.CarDetailsController = controller.NewCarDetailsController(carDetailsService)
	httpRouter           router.Router                   = router.NewMuxRouter()
	httpChiRouter        router.Router                   = router.NewChiRouter()
)

func main() {

	const port string = ":8000"

	httpRouter.GET("/posts", carDetailsController.GetCarDetails)

	httpRouter.SERVE(port)
}
