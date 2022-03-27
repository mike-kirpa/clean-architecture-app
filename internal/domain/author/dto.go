package author

type CreateAuthorDTO struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UpdateAuthorDTO struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
