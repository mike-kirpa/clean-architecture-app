package adapters

import (
	"clean-architecture-app/internal/adapters"
	"clean-architecture-app/internal/book"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	bookURL  = "/books/:book_id"
	booksURL = "/books"
)

type handler struct {
	bookService book.Service
}

// Register implements adapters.Handler
func (h *handler) Register(router *httprouter.Router) {
	router.GET(booksURL, h.GetAllBook)
}

func NewHandler(service book.Service) adapters.Handler {
	return &handler{bookService: service}
}

func (h *handler) GetAllBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//books := h.bookService.GetAllBooks(context.Background(), 0, 0)
	w.Write([]byte("books"))
	w.WriteHeader(http.StatusOK)

}
