package models

import (
	"errors"
	"github.com/garyburd/redigo/redis"
)

type Product struct {
	ProductId    string
	ProductName  string
	ProductIntro string
	Price        string
	StockNum     int
	ProductImg string
}

func (product Product) GetProduct(c redis.Conn, productId string) (error, Product) {
	values, err := redis.Values(c.Do("HGETALL", productId))
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
