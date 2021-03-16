package main

import (
	"github.com/vishnunanduz/go-rest-api/controller"
	router "github.com/vishnunanduz/go-rest-api/http"
	"github.com/vishnunanduz/go-rest-api/repository"
	"github.com/vishnunanduz/go-rest-api/service"
)

var (
	postRepo       repository.PostRepo       = repository.NewSQLiteRepository()
	postService    service.PostService       = service.NewPostService(postRepo)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
	httpChiRouter  router.Router             = router.NewChiRouter()
)

func main() {

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPosts)

	httpRouter.SERVE(":8080")
}
