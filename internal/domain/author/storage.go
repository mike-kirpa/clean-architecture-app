package author

type Storage interface {
	GetOne(uuid string) *Author
	GetAll(limit, offset int) []*Author
	Create(author *Author) *Author
	Delete(author *Author) error
}
