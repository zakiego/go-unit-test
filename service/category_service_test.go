package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zakiego/go-unit-test/entity"
	"github.com/zakiego/go-unit-test/repository"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, error := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, error)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	mockCategory := entity.Category{
		Id:   "2",
		Name: "Laptop",
	}
	categoryRepository.Mock.On("FindById", "2").Return(mockCategory)

	category, error := categoryService.Get("2")
	assert.NotNil(t, category)
	assert.Nil(t, error)
	assert.Equal(t, category.Id, mockCategory.Id)
	assert.Equal(t, category.Name, mockCategory.Name)
}
