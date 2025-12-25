package models

type (
	CategoryResponse struct {
		ID        int    	`json:"id"`
		Nama      string 	`json:"nama_category"`
	}

	CategoryRequest struct {
		Nama string `json:"nama_category"`
	}

	UpdateRequest struct {
		Nama string `json:"nama_category"`
	}
)