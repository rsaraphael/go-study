package model

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `json:"username"`
	Pass     string `json:"pass"`
	Email    string `json:"email"`
}

// Set User's table name to be `profiles`
func (User) TableName() string {
	return "cardap_user"
}
