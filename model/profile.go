package model

import "time"

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
