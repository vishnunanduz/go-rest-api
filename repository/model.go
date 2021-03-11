package repository

import "rest/entity"

type PostRepo interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
