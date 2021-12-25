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


func  DeleteProduct(c redis.Conn, productId string) (error) {
	del, err := redis.Bool(c.Do("DEL", productId))
	fmt.Println(del)
	if err != nil{
		return  err
	}
	return  nil
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
		           "productId", product.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func (product Product) GetAllProduct(c redis.Conn) (error, []Product) {
	var listProducts []Product
	values, err := redis.Values(c.Do("KEYS", "product-*"))
	if err != nil {
		return err, listProducts
	}
	if len(values) < 1 {
		return errors.New("No product here."), listProducts
	}

	for _,productId := range values {
		products := Product{}
		Rvalues, err := redis.Values(c.Do("HGETALL", productId))
		if err != nil {
			return err, listProducts
		}
		err = redis.ScanStruct(Rvalues, &products)
		if err != nil {
			return err, listProducts
		}
		listProducts = append(listProducts, products)
	}
	return nil, listProducts
}

