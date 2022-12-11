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
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestGetData(t *testing.T) {
	t.Run("GetData", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/article")

		userController := mockUserRepository{}
		userController.GetData("", "")

		var response GetArticleResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, response.Data[0].Author, "same")
		//
	})
	t.Run("GetData", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/article")

		falseUserController := New(e, mockServices{})
		falseUserController.GetData("", "")
		// (context)

		var response GetArticleResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, response.Message, "there is a problem on server.")
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
		// fmt.Println(jwtToken)
		// tmp := ""
		// req.Header.Set("Auhtorization", fmt.Sprintf("Bearer %v", tmp))
		context := e.NewContext(req, res)
		context.SetPath("/article")

		inputData := domain.Core{Author: "same", Title: "buat app", Body: "bebas"}
		userController := mockUserRepository{}
		if _, err := userController.Insert(inputData); err != nil {
			log.Fatal(err)
			return
		}

		response := domain.Core{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, response.Author, response.Author)
		assert.Equal(t, response.Title, response.Title)
	})
}

// MOCK OBJECT //
type mockServices struct{}

type mockAuthRepository struct{}

type mockUserRepository struct{}

func (m mockUserRepository) GetData(query string, author string) ([]domain.Core, error) {
	return []domain.Core{
		{ID: 1, Author: "same", Title: "membuat articleapp", Body: "bebas"},
	}, nil
}

func (m mockUserRepository) Insert(domain.Core) (domain.Core, error) {
	return domain.Core{ID: 1, Author: "same", Title: "membuat articleapp", Body: "bebas"}, nil
}

type mockFalseUserRepository struct{}

func (m mockFalseUserRepository) GetData(query string, author string) ([]domain.Core, error) {
	return nil, errors.New("no data")
}

func (m mockFalseUserRepository) Insert(domain.Core) (domain.Core, error) {
	return domain.Core{ID: 1, Author: "same", Title: "membuat articleapp", Body: "bebas"}, errors.New("there is a problem on server.")
}
