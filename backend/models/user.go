package models

type User struct {
	UUIDModel
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `gorm:"unique" json:"email"`
	ApiToken string `json:"api_token" gorm:"index"`
	Timestamps
}
