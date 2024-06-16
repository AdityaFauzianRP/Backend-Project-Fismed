package service

import "time"

type StockBarang struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Total     string    `json:"total"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	UpdatedBy string    `json:"updated_by,omitempty"`
}

type RequestID struct {
	ID string `json:"id"`
}
