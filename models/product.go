package models

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
)


type Product struct {
	ProductId   string    `json:"productId" redis:"productId"`
	ProductName  string `json:"productName" redis:"productName"`
	ProductIntro string `json:"productIntro" redis:"productIntro"`
	Price    string `json:"price" redis:"price"`
	StockNum int    `json:"stockNum" redis:"stockNum"`
	ProductImg string `json:"productImg" redis:"productImg"`
}

func (product Product) GetProduct(c redis.Conn, productId string) (error, Product) {
	values, err := redis.Values(c.Do("HGETALL", productId))
	fmt.Println(values)
	if len(values) < 1 {
		return errors.New("Product is not defined"), product
	}
	if err != nil {
		return err, product
	} else {
		if err = redis.ScanStruct(values, &product); err != nil {
			return err, product
		} else {
			return nil, product
		}
	}
}


//func  DeleteProduct(c redis.Conn, productId string) (error,del) {
//	del, err := redis.Bool(c.Do("DEL",productId))
//	if err != nil{
//		return  err, del
//	}
//	return  nil, del
//}


func (product Product) CreateProduct(c redis.Conn) error {
	_, err := c.Do("SADD", "product", product.ProductId)
	if err != nil {
		return err
	}
	_, err = c.Do("HSET", product.ProductId,
		           "productName", product.ProductName,
				   "productIntro", product.ProductIntro,
		           "price", product.Price,
				   "stockNum", product.StockNum,
		           "productImg", product.ProductImg,
		           "Id", product.ProductId)
	if err != nil {
		return err
	}
	return nil
}
