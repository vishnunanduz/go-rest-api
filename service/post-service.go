package service

import (
	"errors"
	"math/rand"
	"strconv"

	"github.com/vishnunanduz/go-rest-api/entity"
	"github.com/vishnunanduz/go-rest-api/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindByID(id string) (*entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepo
)

func NewPostService(repository repository.PostRepo) PostService {
	repo = repository
	return &service{}
}

func (*service) Validate(post *entity.Post) error {

	if post == nil {
		err := errors.New("Post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("Title is empty")
		return err
	}
	return nil

}
func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)

}

func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()

}

func (*service) FindByID(id string) (*entity.Post, error) {
	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	return repo.FindByID(id)
}
