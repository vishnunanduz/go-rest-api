package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/vishnunanduz/go-rest-api/cache"
	"github.com/vishnunanduz/go-rest-api/entity"
	"github.com/vishnunanduz/go-rest-api/errors"
	"github.com/vishnunanduz/go-rest-api/service"
)

var (
	postService       service.PostService
	redisCacheService cache.PostCache
)

type controller struct{}

type PostController interface {
	GetPostsById(w http.ResponseWriter, r *http.Request)
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPosts(w http.ResponseWriter, r *http.Request)
}

func NewPostController(service service.PostService, cache cache.PostCache) PostController {
	postService = service
	redisCacheService = cache
	return &controller{}
}

func (*controller) GetPostsById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	postID := strings.Split(r.URL.Path, "/")[2]
	var post *entity.Post = redisCacheService.Get(postID)
	if post == nil {
		post, err := postService.FindByID(postID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errors.ServiceError{Message: "No post found"})
			return
		}

		redisCacheService.Set(postID, post)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(post)

	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(post)
	}

}

func (*controller) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error while reading data from DB"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (*controller) AddPosts(w http.ResponseWriter, r *http.Request) {
	var post entity.Post
	w.Header().Set("Content-type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error while unmarshalling"})
		return
	}

	err1 := postService.Validate(&post)
	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := postService.Create(&post)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}
