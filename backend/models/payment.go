package models

// Payment -
type Payment struct {
	EmployeeID int64   `json:"employee_id" gorm:"foreignkey:employee_id"`
	Amount     float64 `json:"amount"`
	Confirmed  bool    `json:"confirmed"`
	DatePaid   string  `json:"date_paid"`
}
