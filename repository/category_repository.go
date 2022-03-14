package repository

import "github.com/zakiego/go-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
