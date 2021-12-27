package controllers

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

type Response struct {
	status  bool
	message string
	result  string
}

type ConnRedis struct {
	DB redis.Conn
}

func (h *ConnRedis) Connect() {
	var err error
	//通用
	h.DB, err = redis.Dial("tcp", "127.0.0.1:6380")
	//qq用，注释掉就行
	//h.DB, err = redis.Dial("tcp","192.168.1.105:6380",redis.DialPassword("root"))
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
}
