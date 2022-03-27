package main

import (
	author3 "clean-architecture-app/internal/adapters/api/author"
	book3 "clean-architecture-app/internal/adapters/api/book"
	author2 "clean-architecture-app/internal/adapters/db/author"
	book2 "clean-architecture-app/internal/adapters/db/book"
	"clean-architecture-app/internal/domain/author"
	"clean-architecture-app/internal/domain/book"
)

func main() {
	bookStorage := book2.NewStorage()
	bookService := book.NewService(bookStorage)
	bookHandler := book3.NewHandler(bookService)

	authorStorage := author2.NewStorage()
	authorService := author.NewService(authorStorage)
	authorHandler := author3.NewHandler(authorService)

}
