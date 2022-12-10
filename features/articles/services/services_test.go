package services

import (
	"errors"
	"testing"

	"github.com/khalidrianda/ArticleApp/features/articles/domain"
	"github.com/khalidrianda/ArticleApp/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPendaki(t *testing.T) {
	repo := mocks.NewRepository(t)
	query := ""
	author := ""
	t.Run("Sukses Get All Data", func(t *testing.T) {
		repo.On("GetData", mock.Anything, mock.Anything).Return([]domain.Core{{ID: uint(1), Author: "same", Title: "buat app", Body: "bebas"}},
			nil).Once()
		srv := New(repo)
		res, err := srv.Show(query, author)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Get All Booking", func(t *testing.T) {
		repo.On("GetData", mock.Anything, mock.Anything).Return([]domain.Core{}, errors.New("no data")).Once()
		srv := New(repo)
		res, err := srv.Show(query, author)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		repo.AssertExpectations(t)
	})
}

func TestAddClimber(t *testing.T) {
	repo := new(mocks.Repository)
	input := domain.Core{Author: "same", Title: "buat app", Body: "bebas"}
	returnRespon := domain.Core{ID: uint(1), Author: "same", Title: "buat app", Body: "bebas"}

	t.Run("add climber success", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(returnRespon, nil).Once()

		usecase := New(repo)
		res, err := usecase.Create(input)
		assert.NoError(t, err)
		assert.Equal(t, returnRespon, res)
		repo.AssertExpectations(t)
	})

	t.Run("failed add climber", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(domain.Core{}, errors.New("some problem on database")).Once()

		usecase := New(repo)
		res, err := usecase.Create(domain.Core{})
		assert.EqualError(t, err, "some problem on database")
		assert.Equal(t, domain.Core{}, res)
		repo.AssertExpectations(t)
	})
}
