package delivery

import "github.com/khalidrianda/ArticleApp/features/articles/domain"

type PostFormat struct {
	Author string `json:"author" form:"author"`
	Title  string `json:"title" form:"title"`
	Body   string `json:"body" form:"body"`
}

func ToDomain(i interface{}) domain.Core {
	var res domain.Core
	switch i.(type) {
	case PostFormat:
		cnv := i.(PostFormat)
		res = domain.Core{Author: cnv.Author, Title: cnv.Title, Body: cnv.Body}
	}
	return res
}
