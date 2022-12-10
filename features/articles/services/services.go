package services

import "github.com/khalidrianda/ArticleApp/features/articles/domain"

type articleServices struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Services {
	return &articleServices{qry: repo}
}

func (as *articleServices) Create(newData domain.Core) (domain.Core, error) {
	res, err := as.qry.Insert(newData)
	if err != nil {
		return newData, err
	}

	return res, nil
}

func (as *articleServices) Show(query string, author string) ([]domain.Core, error) {
	if query != "" {
		query = "%" + query + "%"
	}
	if author != "" {
		author = "%" + author + "%"
	}

	res, err := as.qry.GetData(query, author)
	if err != nil {
		return nil, err
	}

	return res, nil
}
