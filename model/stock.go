package model

import "time"

type StockProduct struct {
	Id          string    `json:"id"`
	CompanyName string    `json:"companyName"`
	Rating      int       `json:"rating"`
	Risk        string    `json:"risk"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
