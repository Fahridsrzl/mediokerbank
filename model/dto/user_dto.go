package dto

import "time"

type ProfileCreateDto struct {
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	Citizenship       string `json:"citizenship"`
	NationalID        string `json:"nationalID"`
	BirthPlace        string `json:"birthPlace"`
	BirthDate         string `json:"birthDate"`
	Gender            string `json:"gender"`
	MaritalStatus     string `json:"maritalStatus"`
	Occupation        string `json:"occupation"`
	MonthlyIncome     int    `json:"monthlyIncome"`
	PhoneNumber       string `json:"phoneNumber"`
	UrgentPhoneNumber string `json:"urgentPhoneNumber"`
	Photo             string `json:"photo"`
	IDCard            string `json:"idCard"`
	SalarySlip        string `json:"salarySlip"`
	UserID            string `json:"userID"`
}

type AddressCreateDto struct {
	AddressLine string `json:"addressLine"`
	City        string `json:"city"`
	Province    string `json:"province"`
	PostalCode  string `json:"postalCode"`
	Country     string `json:"country"`
}

type ResponseStatus struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
	Status      string    `json:"status"`
	CreditScore int       `json:"creditScore"`
	Balance     int       `json:"balance"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UserDto struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
	Status      string    `json:"status"`
	CreditScore int       `json:"creditScore"`
	Balance     int       `json:"balance"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
