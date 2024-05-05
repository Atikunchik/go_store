package model

type User struct {
	Login    string `gorm:"type:varchar;primary_key;unique;NOT NULL" json:"login"`
	Password string `json:"password"`
}
