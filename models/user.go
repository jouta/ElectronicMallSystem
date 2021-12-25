package models

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type User struct {
	Id    string
	Score int
	Name  string
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
	_, err := c.Do("SADD", "user", user.Id)
	if err != nil {
		return err
	}
	_, err = c.Do("HSET", user.Id, "Score", user.Score, "Name", user.Name, "Id", user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (user User) GetTop(c redis.Conn, length int) ([]string, error) {
	values, err := redis.Strings(c.Do("SORT", "user", "by", "*->Score", "desc", "limit", 0, length))
	if err != nil {
		return []string{}, err
	}
	return values, nil
}