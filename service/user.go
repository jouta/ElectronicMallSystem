package service

import (
	"errors"
	"fmt"
	"mall/model"
	"mall/repository"
	uuid "github.com/satori/go.uuid"
	"mall/utils"
)

type UserSrv interface {
	Get(user model.User) (*model.User, error)
	Exist(user model.User) *model.User
	ExistByUserID(id string) *model.User
	Add(user model.User) (*model.User, error)
	Edit(user model.User) (bool, error)
	Delete(id string) (bool, error)
}

type UserService struct {
	Repo repository.UserRepoInterface
}

func (srv *UserService) Get(user model.User) (*model.User, error) {
	return srv.Repo.Get(user)
}
func (srv *UserService) Exist(user model.User) *model.User {
	return srv.Repo.Exist(user)
}

func (srv *UserService) ExistByUserID(id string) *model.User {
	return srv.Repo.ExistByUserID(id)
}

func (srv *UserService) Add(user model.User) (*model.User, error) {
	//根据用户Id判断是否存在用户
	result := srv.Repo.ExistByUserID(user.UserId)
	if result != nil {
		return nil, errors.New("用户已经存在")
	}
	user.UserId = uuid.NewV4().String()
	if user.Password == "" {
		user.Password = utils.Md5("123456")
	}
	return srv.Repo.Add(user)
}

func (srv *UserService) Edit(user model.User) (bool, error) {
	if user.UserId == "" {
		return false, fmt.Errorf("参数错误")
	}

	exist := srv.Repo.ExistByUserID(user.UserId)
	if exist == nil {
		return false, errors.New("参数错误")
	}
	exist.UserName = user.UserName
	exist.Address = user.Address
	return srv.Repo.Edit(*exist)
}

func (srv *UserService) Delete(id string) (bool, error) {
	if id == "" {
		return false, errors.New("参数错误")
	}

	user := srv.ExistByUserID(id)
	if user == nil {
		return false, errors.New("参数错误")
	}
	return srv.Repo.Delete(*user)
}



