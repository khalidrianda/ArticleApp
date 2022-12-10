package repository

import (
	"github.com/khalidrianda/ArticleApp/features/articles/domain"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Author string
	Title  string
	Body   string
}

func FromDomain(a domain.Core) Article {
	return Article{
		Model:  gorm.Model{ID: a.ID},
		Author: a.Author,
		Title:  a.Title,
		Body:   a.Body,
	}
}

func ToDomain(a Article) domain.Core {
	return domain.Core{
		ID:     a.ID,
		Author: a.Author,
		Title:  a.Title,
		Body:   a.Body,
	}
}

func ToDomainArray(a []Article) []domain.Core {
	var resData []domain.Core
	for _, val := range a {
		resData = append(resData, domain.Core{
			ID:     val.ID,
			Author: val.Author,
			Title:  val.Title,
			Body:   val.Body,
		})
	}
	return resData
}
