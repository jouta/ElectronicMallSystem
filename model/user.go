package model

type User struct {
	UserId    string    `json:"userId" gorm:"column:user_id"`
	UserName  string    `json:"userName" gorm:"column:user_name"`
	Password  string    `json:"password" gorm:"column:password"`
	Address   string    `json:"address" gorm:"column:address"`
}