package dto

<<<<<<< HEAD
// contoh createprofile by id
// photo, idcard, salaryslip diisi setelah generate pathfile
// type ProfileCreateDto struct {
// 	FirstName         "binding required"
// 	LastName          "binding required"
// 	Citizenship       "binding required"
// 	NationalId        "binding required"
// 	BirthPlace        "binding required"
// 	BirthDate         string "binding required"
// 	Gender            "binding required"
// 	MaritalStatus     "binding required"
// 	Occupation        "binding required"
// 	MonthlyIncome     "binding required"
// 	PhoneNumber       string "binding required"
// 	UrgentPhoneNumber string "binding required"
// 	Photo             string "pathfile"
// 	IdCard            string "pathfile"
// 	SalarySlip        string "pathfile"
// 	UserId            string "binding required"
// }
=======
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
>>>>>>> users

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
	LoanActive  int       `json:"loanActive"`
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
	LoanActive  int       `json:"loanActive"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}