package book

import (
	"clean-architecture-app/internal/adapters/api/author"
	"clean-architecture-app/internal/adapters/api/book"
	"context"
)

type service struct {
	storage       Storage
	authorService author.Service
	genreService  genre.Service
}

func NewService(storage Storage) book.Service {
	return &service{storage: storage}
}

func (s *service) Create(ctx context.Context, dto *CreateBookDTO) *Book {
	author := s.authorService.GetByUUID(ctx, dto.Author_UUID)
	genre := s.genreService.Genre_UUID(ctx, dto.Genre_UUID)
	return nil
}

func (s *service) GetByUUID(ctx context.Context, uuid string) *Book {
	return s.storage.GetOne(uuid)
}

func (s *service) GetAll(ctx context.Context, limit, offset int) []*Book {
	return s.storage.GetAll(limit, offset)
}
