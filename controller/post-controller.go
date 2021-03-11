package controller

import (
	"encoding/json"
	"net/http"
	"rest/entity"
	"rest/errors"
	"rest/service"
)

var (
	postService service.PostService
)

type controller struct{}

type PostController interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPosts(w http.ResponseWriter, r *http.Request)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
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
