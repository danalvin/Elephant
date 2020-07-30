package models

// Employee -
type Employee struct {
	BaseModel
	Name        string  `json:"name"`
	NationalID  string  `json:"national_id"`
	PhoneNumber string  `json:"account"`
	Wage        float64 `json:"wage"`
	Active      bool    `json:"active"`
}
