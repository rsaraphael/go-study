package model

type User struct {
	ID       int    `gorm:"primary_key"`
	Username string `json:"username"`
	Pass     string `json:"pass"`
	Email    string `json:"email"`
}

func (User) TableName() string {
	return "cardap_user"
}
