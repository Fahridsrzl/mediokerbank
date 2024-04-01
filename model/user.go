package model

import "time"

type User struct {
	Id          string    `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
	CreditScore string    `json:"creditScore"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Profile struct {
	ID                string    `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Citizenship       string    `json:"citizenship"`
	NationalID        string    `json:"nationalID"`
	BirthPlace        string    `json:"birthPlace"`
	BirthDate         string    `json:"birthDate"`
	Gender            string    `json:"gender"`
	MaritalStatus     string    `json:"maritalStatus"`
	Occupation        string    `json:"occupation"`
	MonthlyIncome     int       `json:"monthlyIncome"`
	PhoneNumber       string    `json:"phoneNumber"`
	UrgentPhoneNumber string    `json:"urgentPhoneNumber"`
	Photo             string    `json:"photo"`
	UserID            string    `json:"userID"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type Address struct {
    ID          string    `json:"id"`
    AddressLine string    `json:"addressLine"`
    City        string    `json:"city"`
    Province    string    `json:"province"`
    PostalCode  string    `json:"postalCode"`
    Country     string    `json:"country"`
    ProfileID   string    `json:"profileID"`
    CreatedAt   time.Time `json:"createdAt"`
    UpdatedAt   time.Time `json:"updatedAt"`
}

func (u User) IsValidRole() bool {
	return u.Role == "admin"
}
