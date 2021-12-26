package controllers

import (
	"github.com/garyburd/redigo/redis"
	"log"
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
	h.DB, err = redis.Dial("tcp","192.168.1.105:6380",redis.DialPassword("root"))
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
}
