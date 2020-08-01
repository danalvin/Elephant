package models

// User - user management will happen at django level
type User struct {
	BaseModel
	Avatar    string `json:"avatar"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
