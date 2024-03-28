package dto

type StockProductCreateDto struct {
	CompanyName string `json:"companyName" binding:"required"`
	Risk        string `json:"risk" binding:"required"`
}

// risk = low, medium, high, very high
// rating = 1 - 5

type StockProductSearchByQueryDto struct {
	Rating int    `json:"rating"`
	Risk   string `json:"risk"`
}

type StockProductUpdateDto struct {
	CompanyName string `json:"companyName" binding:"required"`
	Rating      int    `json:"rating" binding:"required"`
	Risk        string `json:"risk" binding:"required"`
}
