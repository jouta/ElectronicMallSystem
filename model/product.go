package model
type Product struct{
	ProductId            string `json:"productId" gorm:"column:product_id"`
	ProductName          string `json:"productName" gorm:"column:product_name"`
	ProductIntro         string `json:"productIntro" gorm:"column:product_intro"`
	Price                int    `json:"price" gorm:"column:price"`
	StockNum             int    `json:"stockNum" gorm:"column:stock_num"`
}
