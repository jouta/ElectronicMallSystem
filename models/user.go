package models

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
)
/*
type User struct {
	Id    string
	Score int
	Name  string
}

 */
type User struct {
	UserId   string    `json:"userId" redis:"userId"`
	UserName string `json:"userName" redis:"userName"`
	PassWord string `json:"passWord" redis:"passWord"`
	Address  string    `json:"address" redis:"address"`
	UserType   int    `json:"userType" redis:"userType"`
}

func (user User) GetAllUser(c redis.Conn) (error, []User) {
	var listUsers []User
	values, err := redis.Values(c.Do("KEYS", "user-*"))
	if err != nil {
		return err, listUsers
	}
	if len(values) < 1 {
		return errors.New("No users here."), listUsers
	}

	for _,userid := range values {
		users := User{}
		Rvalues, err := redis.Values(c.Do("HGETALL", userid))
		if err != nil {
			return err, listUsers
		}
		err = redis.ScanStruct(Rvalues, &users)
		if err != nil {
			return err, listUsers
		}
		listUsers = append(listUsers, users)
	}
	return nil, listUsers
}


func (user User) GetUser(c redis.Conn, userid string) (error, User) {
	values, err := redis.Values(c.Do("HGETALL", userid))
	fmt.Println(values)
	if len(values) < 1 {
		return errors.New("User is not defined"), user
	}

	if err != nil {
		return err, user
	} else {
		if err = redis.ScanStruct(values, &user); err != nil {
			return err, user
		} else {
			return nil, user
		}
	}
}

func (user User) Create(c redis.Conn) error {
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

func  DeleteUser(c redis.Conn, userId string) (error) {
	_, err := redis.Bool(c.Do("DEL", userId))
	if err != nil{
		return  err
	}
	return  nil
}

/*
func (user User) GetTop(c redis.Conn, length int) ([]string, error) {
	values, err := redis.Strings(c.Do("SORT", "user", "by", "*->Score", "desc", "limit", 0, length))
	if err != nil {
		return []string{}, err
	}
	return values, nil
}
*/