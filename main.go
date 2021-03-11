package main

import (
	"fmt"
	"net/http"
	"rest/controller"
	router "rest/http"
	"rest/repository"
	"rest/service"
)

var (
	PostRepo       repository.PostRepo       = repository.NewFireStoreRepository()
	postService    service.PostService       = service.NewPostService(PostRepo)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
	httpChiRouter  router.Router             = router.NewChiRouter()
)

func main() {

	const port string = ":8080"
	httpChiRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "server started")
	})

	httpChiRouter.GET("/posts", postController.GetPosts)
	httpChiRouter.POST("/posts", postController.AddPosts)

	httpChiRouter.SERVE(port)
}
