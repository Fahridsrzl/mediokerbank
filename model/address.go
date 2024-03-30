package model

import "time"

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
