package clusterModels

import "github.com/go-redis/redis/v8"

type User struct {
	UserId   string    `json:"userId" redis:"userId"`
	UserName string `json:"userName" redis:"userName"`
	PassWord string `json:"passWord" redis:"passWord"`
	Address  string    `json:"address" redis:"address"`
	UserType   int    `json:"userType" redis:"userType"`
}

//添加用户
func (user User) Create(c *redis.Client) error {
	_, err := c.Do("SADD", "user", user.UserId)
	if err != nil {
		return err
	}
	_, err = c.Do("HSET", user.UserId, "userId", user.UserId, "userName", user.UserName, "passWord", user.PassWord, "address", user.Address, "userType", user.UserType)
	if err != nil {
		return err
	}
	return nil
}
