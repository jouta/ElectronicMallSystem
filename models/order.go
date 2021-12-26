package models

import (
	"errors"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type Order struct {
	OrderId     string `json:"orderId" redis:"orderId"`
	UserId      string `json:"userId" redis:"userId"`
	ProductId   string `json:"productId" redis:"productId"`
	Price       string `json:"price" redis:"price"`
	OrderStatus int    `json:"orderStatus" redis:"orderStatus"`
	PayTime     string `json:"payTime" redis:"payTime"`
	OrderTime   string `json:"orderTime" redis:"orderTime"`
	Remark      string `json:"remark" redis:"remark"`
}

func (order Order) CreateOrder(c redis.Conn) error {
	_, err := c.Do("SADD", "order", order.OrderId)
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
	)
	if err != nil {
		return err
	}
	return nil
}

func (order Order) GetOrder(c redis.Conn, orderId string) (error, Order) {
	values, err := redis.Values(c.Do("HGETALL", orderId))
	fmt.Println(values)
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
	)
	if err != nil {
		return err
	}
	return nil
}
