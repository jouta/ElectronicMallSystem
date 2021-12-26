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
	h.DB, err = redis.Dial("tcp", "127.0.0.1:6380")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
}
