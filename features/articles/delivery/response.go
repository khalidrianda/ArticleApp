package delivery

import "github.com/khalidrianda/ArticleApp/features/articles/domain"

type ArticleResponse struct {
	ID     uint   `json:"id" form:"id"`
	Author string `json:"author" form:"author"`
	Title  string `json:"title" form:"title"`
	Body   string `json:"body" form:"body"`
}

func SuccessResponse(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
		"data":    data,
	}
}

func SuccessResponseNoData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func FailResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"message": msg,
	}
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "add":
		cnv := core.(domain.Core)
		res = ArticleResponse{ID: cnv.ID, Author: cnv.Author, Title: cnv.Title, Body: cnv.Body}
	case "get":
		var data []ArticleResponse
		cnv := core.([]domain.Core)
		for _, val := range cnv {
			data = append(data, ArticleResponse{ID: val.ID, Author: val.Author, Title: val.Title, Body: val.Body})
		}
		res = data
	}
	return res
}
