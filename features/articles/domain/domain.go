package domain

import "github.com/labstack/echo/v4"

type Core struct {
	ID     uint
	Author string
	Title  string
	Body   string
}

type Repository interface {
	Insert(newData Core) (Core, error)
	GetData(query string, author string) ([]Core, error)
}

type Services interface {
	Create(newData Core) (Core, error)
	Show(query string, author string) ([]Core, error)
}

type Handler interface {
	Insert() echo.HandlerFunc
	GetData() echo.HandlerFunc
}
