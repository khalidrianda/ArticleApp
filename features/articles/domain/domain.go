package domain

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
