package model

type Order struct {
	OrderId     string `json:"orderId" gorm:"column:order_id"`
	UserId      string `json:"userId" gorm:"column:user_id"`
	ProductId   string `json:"productId" gorm:"column:product_id"`
	Price       int    `json:"price" gorm:"column:price"`
	OrderStatus int    `json:"orderStatus" gorm:"column:order_status"`
	PayTime     string `json:"payTime" gorm:"column:pay_time"`
	OrderTime   string  `json:"orderTime" gorm:"column:order_time"`
}