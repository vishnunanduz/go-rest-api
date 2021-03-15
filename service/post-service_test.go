package service

import (
	"testing"

	"github.com/vishnunanduz/go-rest-api/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)

}
func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Post is empty", err.Error())
}

func TestValidateTitle(t *testing.T) {
	post := entity.Post{ID: 1, Title: "", Text: "Nothing"}
	testService := NewPostService(nil)

	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, "Title is empty", err.Error())
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)
	var identifier int64 = 1
	// Expctations
	post := entity.Post{ID: identifier, Title: "A", Text: "Nothing"}
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)
	result, _ := testService.FindAll()

	// Mock assertion
	mockRepo.AssertExpectations(t)

	// Data assertion
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "Nothing", result[0].Text)
}

func TestCreatePost(t *testing.T) {

	mockRepo := new(MockRepository)
	post := entity.Post{Title: "A", Text: "Nothing"}

	// Expctations
	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)
	result, err := testService.Create(&post)

	// Mock assertion
	mockRepo.AssertExpectations(t)

	// Data assertion
	assert.NotNil(t, result.ID)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "Nothing", result.Text)
	assert.Nil(t, err)

}
