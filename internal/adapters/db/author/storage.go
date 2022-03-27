package author

import (
	"clean-architecture-app/internal/domain/author"

	"go.mongodb.org/mongo-driver/mongo"
)

type authorStorage struct {
	db *mongo.Database
}

func NewStorage(db *mongo.Database) author.Storage {
	return &authorStorage{
		db: db,
	}
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
