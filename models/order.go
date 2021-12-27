package models

import (
	"errors"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type Order struct {
	OrderId     string  `json:"orderId" redis:"orderId"`
	UserId      string  `json:"userId" redis:"userId"`
	ProductId   string  `json:"productId" redis:"productId"`
	Price       float64 `json:"price" redis:"price"`
	OrderStatus int     `json:"orderStatus" redis:"orderStatus"`
	PayTime     string  `json:"payTime" redis:"payTime"`
	OrderTime   string  `json:"orderTime" redis:"orderTime"`
	Remark      string  `json:"remark" redis:"remark"`
	ProductNum  int     `json:"productNum" redis:"productNum"`
}

func (order Order) CreateOrder(c redis.Conn) error {
	_, err := c.Do("MULTI") //事务开始
	if err != nil {
		return err
	}
	_, err = c.Do("SADD", "order", order.OrderId)
	if err != nil {
		return err
	}
	_, err = c.Do("HSET", order.OrderId,
		"orderId", order.OrderId,
		"userId", order.UserId,
		"productId", order.ProductId,
		"price", order.Price,
		"orderStatus", order.OrderStatus,
		"payTime", order.PayTime,
		"orderTime", order.OrderTime,
		"remark", order.Remark,
		"productNum", order.ProductNum,
	)
	if err != nil {
		return err
	}
	//先查product的stockNum
	product := Product{}
	err1, productData := product.GetProduct(c, order.ProductId)
	if err1 != nil {
		return err1
	}
	stockNum := productData.StockNum - order.ProductNum
	_, err = c.Do("HSET", order.ProductId,
		"stockNum", stockNum,
	)
	if err != nil {
		return err
	}
	_, err = c.Do("EXEC") //事务结束
	if err != nil {
		return err
	}

	return nil
}

func (order Order) GetOrder(c redis.Conn, orderId string) (error, Order) {
	values, err := redis.Values(c.Do("HGETALL", orderId))
	if len(values) < 1 {
		return errors.New("Order is not defined"), order
	}
	if err != nil {
		return err, order
	} else {
		if err = redis.ScanStruct(values, &order); err != nil {
			return err, order
		} else {
			return nil, order
		}
	}
}

func DeleteOrder(c redis.Conn, orderId string) error {
	del, err := redis.Bool(c.Do("DEL", orderId))
	fmt.Println(del)
	if err != nil {
		return err
	}
	return nil
}

func (order Order) GetAllOrder(c redis.Conn) (error, []Order) {
	var listOrders []Order
	values, err := redis.Values(c.Do("KEYS", "order-*"))
	if err != nil {
		return err, listOrders
	}
	if len(values) < 1 {
		return errors.New("No orders here."), listOrders
	}

	for _, orderid := range values {
		orders := Order{}
		Rvalues, err := redis.Values(c.Do("HGETALL", orderid))
		if err != nil {
			return err, listOrders
		}
		err = redis.ScanStruct(Rvalues, &orders)
		if err != nil {
			return err, listOrders
		}
		listOrders = append(listOrders, orders)
	}
	return nil, listOrders
}

func (order Order) PayOrder(c redis.Conn, orderid string) error {
	_, err := c.Do("HSET", order.OrderId,
		"orderId", order.OrderId,
		"userId", order.UserId,
		"productId", order.ProductId,
		"price", order.Price,
		"orderStatus", order.OrderStatus,
		"payTime", order.PayTime,
		"orderTime", order.OrderTime,
		"remark", order.Remark,
		"productNum", order.ProductNum,
	)
	if err != nil {
		return err
	}
	return nil
}
