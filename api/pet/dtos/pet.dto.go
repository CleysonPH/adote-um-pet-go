package dtos

type PetResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"nome"`
	History string `json:"historia"`
	Picture string `json:"foto"`
}

type PetRequest struct {
	Name    string `json:"nome" binding:"required,max=100"`
	History string `json:"historia" binding:"required,max=255"`
	Picture string `json:"foto" binding:"required,uri,max=255"`
}
