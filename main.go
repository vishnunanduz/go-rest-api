package main

import (
	"github.com/vishnunanduz/go-rest-api/cache"
	"github.com/vishnunanduz/go-rest-api/controller"
	router "github.com/vishnunanduz/go-rest-api/http"
	"github.com/vishnunanduz/go-rest-api/repository"
	"github.com/vishnunanduz/go-rest-api/service"
)

var (
	postRepo         repository.PostRepo       = repository.NewDynamoDBRepository()
	postService      service.PostService       = service.NewPostService(postRepo)
	postcacheService cache.PostCache           = cache.NewRedisCache("localhost:6379", 1, 20)
	postController   controller.PostController = controller.NewPostController(postService, postcacheService)
	httpRouter       router.Router             = router.NewMuxRouter()
	httpChiRouter    router.Router             = router.NewChiRouter()
)

func main() {

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.GET("/posts/{id}", postController.GetPostsById)
	httpRouter.POST("/posts", postController.AddPosts)

	httpRouter.SERVE(":8080")
}
