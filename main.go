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
	PostRepo       repository.PostRepo       = repository.NewSQLiteRepository()
	postService    service.PostService       = service.NewPostService(PostRepo)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
	httpChiRouter  router.Router             = router.NewChiRouter()
)

func main() {

	const port string = ":8080"
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "server started")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPosts)

	httpRouter.SERVE(port)
}
