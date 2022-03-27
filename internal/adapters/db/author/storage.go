package author

import (
	"clean-architecture-app/internal/domain/author"
)

type authorStorage struct {
}

func NewStorage() author.Storage {
	return &authorStorage{}
}

func (as *authorStorage) GetOne(uuid string) *author.Author {
	return nil
}
func (as *authorStorage) GetAll(limit, offset int) []*author.Author {
	return nil
}
func (as *authorStorage) Create(book *author.Author) *author.Author {
	return nil
}
func (as *authorStorage) Delete(book *author.Author) error {
	return nil
}
