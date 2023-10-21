package dto

type ItemDTO struct {
	Name       string `json:"name"`
	CategoryID string `json:"category_id"`
}

type ItemUpdateDTO struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CategoryID string `json:"category_id"`
}
