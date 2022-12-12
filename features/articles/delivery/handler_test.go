package delivery

import (
	"bytes"
	"errors"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/khalidrianda/ArticleApp/features/articles/domain"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// func TestGetData(t *testing.T) {
// 	handler := mocks.NewHandler(t)
// 	var srv domain.Services
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	res := httptest.NewRecorder()
// 	context := e.NewContext(req, res)
// 	context.SetPath("/article")

// 	t.Run("Sukses Get All Data", func(t *testing.T) {
// 		handler.On("GetData", mock.Anything).Return(res).Once()
// 		srv := New
// 		srv.GetData()
// 		// assert.Nil(t, err)
// 		// var response GetArticleResponseFormat

// 		// json.Unmarshal([]byte(res.Body.Bytes()), &response)
// 		// assert.Equal(t, response.Data[0].Author, "same")
// 		assert.NotEmpty(t, res)
// 		handler.AssertExpectations(t)
// 	})

// }

func TestGetData(t *testing.T) {
	t.Run("GetData", func(t *testing.T) {
		var srv domain.Services = &mockServices{}
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/article")

		handler := New(e, srv)
		handler.GetData()(context)

		var response GetArticleResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, response.Data[0].Author, "same")
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "success get data", response.Message)
	})
	t.Run("GetData", func(t *testing.T) {
		e := echo.New()
		var srv domain.Services = &mockFalseService{}
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/article")

		handler := New(e, srv)
		handler.GetData()(context)

		var response GetArticleResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 200, res.Code)
		assert.Equal(t, response.Message, "no data")

	})
}

func TestInsert(t *testing.T) {
	t.Run("Test Create", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"author": "same",
			"title":  "membuat web",
			"body":   "bebas",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/article")

		handler := New(e, &mockServices{})
		handler.Insert()(context)

		var response ArticleResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "same", response.Data.Author)
		assert.Equal(t, 201, res.Code)
		assert.Equal(t, "success post the article", response.Message)
	})
	t.Run("Test Create Failed", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"author": "same",
			"title":  "membuat web",
			"body":   "bebas",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/article")

		handler := New(e, &mockFalseService{})
		handler.Insert()(context)

		var response ArticleResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "", response.Data.Author)
		assert.Equal(t, 500, res.Code)
		assert.Equal(t, "there is a problem on server", response.Message)
	})
}

// MOCK OBJECT //
type mockServices struct{}

func (s *mockServices) Create(newData domain.Core) (domain.Core, error) {
	return domain.Core{ID: 1, Author: "same", Title: "membuat articleapp", Body: "bebas"}, nil
}

func (s *mockServices) Show(query string, author string) ([]domain.Core, error) {
	return []domain.Core{
		{ID: 1, Author: "same", Title: "membuat articleapp", Body: "bebas"},
	}, nil
}

type mockFalseService struct{}

func (s *mockFalseService) Create(domain.Core) (domain.Core, error) {
	return domain.Core{Author: "same", Title: "membuat articleapp", Body: "bebas"}, errors.New("there is a problem on server.")
}

func (s *mockFalseService) Show(query string, author string) ([]domain.Core, error) {
	return nil, errors.New("no data")
}
