package model

import "time"

type User struct {
	Id          string    `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
	CreditScore string    `json:"creditScore`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt`
}

func (u User) IsValidRole() bool {
	return u.Role == "admin"
}
