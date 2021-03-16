package repository

import "github.com/vishnunanduz/go-rest-api/entity"

type PostRepo interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	Delete(post *entity.Post) error
	FindByID(id string) (*entity.Post, error)
}
