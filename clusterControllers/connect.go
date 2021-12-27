package clusterControllers

import (
	"github.com/go-redis/redis/v8"
)

type Response struct {
	status  bool
	message string
	result  string
}

type ConnRedis struct {
	DB *redis.ClusterClient
}

func (h *ConnRedis) Connect() {
	//var err error

	h.DB = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
}
