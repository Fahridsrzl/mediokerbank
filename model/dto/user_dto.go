package dto

// contoh createprofile by id
// photo, idcard, salaryslip diisi setelah generate pathfile
type ProfileCreateDto struct {
	FirstName         "binding required"
	LastName          "binding required"
	Citizenship       "binding required"
	NationalId        "binding required"
	BirthPlace        "binding required"
	BirthDate         string "binding required"
	Gender            "binding required"
	MaritalStatus     "binding required"
	Occupation        "binding required"
	MonthlyIncome     "binding required"
	PhoneNumber       string "binding required"
	UrgentPhoneNumber string "binding required"
	Photo             string "pathfile"
	IdCard            string "pathfile"
	SalarySlip        string "pathfile"
	UserId            string "binding required"
}

// contoh response saat select user by id
// {
// 	"id": "adsdawdasd",
// 	"username": "hahaha",
// 	"email": "hahaha@gmail.com",
// 	"password": "hash#dawdads", "optional"
// 	"role": "user",
// 	"status": "verfied",
// 	"profile": {
// 		"firstName": "hahaha",
// 		"dll..."
// 	},
// 	"creditScore": 80,
// 	"balance": 150000,
// 	"loanActive": 1,
// 	"loans": [
// 		{
// 			"productName": "classic",
// 			"amount": 800000,
// 			"installmentPeriod": 5,
// 			"installmentAmount": 168000,
// 			"periodLeft": 2,
// 			"periodUnit": "month"
// 		}
// 	]
// }
