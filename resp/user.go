package resp

type User struct {
	Id        string `json:"id"`
	Key       string `json:"key"`
	UserId    string `json:"userId" gorm:"column:user_id"`
	UserName  string `json:"userName" gorm:"column:user_name"`
	Address   string `json:"address" gorm:"column:address"`
}