package repository

import (
	"errors"

	"github.com/khalidrianda/ArticleApp/features/articles/domain"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.Repository {
	return &repoQuery{db: db}
}

func (rq *repoQuery) Insert(newData domain.Core) (domain.Core, error) {
	var data Article = FromDomain(newData)

	if err := rq.db.Create(&data).Error; err != nil {
		return newData, err
	}

	newData = ToDomain(data)
	return newData, nil
}

func (rq *repoQuery) GetData(query string, author string) ([]domain.Core, error) {
	var resQry []Article
	var result *gorm.DB

	if query != "" && author != "" {
		result = rq.db.Where("author = ?", author).Where("title = ? OR author = ?", query, query).Order("created_at desc").Find(&resQry)
	} else if query != "" {
		result = rq.db.Where("title = ? OR author = ?", query, query).Order("created_at desc").Find(&resQry)
	} else if author != "" {
		result = rq.db.Where("author = ?", author).Order("created_at desc").Find(&resQry)
	} else {
		result = rq.db.Order("created_at desc").Find(&resQry)
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no data")
	}

	resData := ToDomainArray(resQry)
	return resData, nil
}
