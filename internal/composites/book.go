package composites

import (
	"clean-architecture-app/internal/adapters/api"
	"clean-architecture-app/internal/adapters/api/book"
	book3 "clean-architecture-app/internal/adapters/db/book"
	book2 "clean-architecture-app/internal/domain/book"
)

type BookComposite struct {
	Storage book2.Storage
	Service book.Service
	Handler api.Handler
}

func NewBookComposite(mongoComposite MongoDBComposite) (*BookComposite, error) {
	storage := book3.NewStorage(mongoComposite.db)
	service := book2.NewService(storage)
	handler := book.NewHandler(service)
	return &BookComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
