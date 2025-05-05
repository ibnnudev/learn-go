package repository

import "learn-go-testing/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
