package repository

import (
	"fmt"
	"mall/model"
    "github.com/jinzhu/gorm"
)
type UserRepository struct {
	DB *gorm.DB
}

type UserRepoInterface interface {
	Get(user model.User) (*model.User, error)
	Exist(user model.User) *model.User
	ExistByUserID(id string) *model.User
	Add(user model.User) (*model.User, error)
	Edit(user model.User) (bool, error)
	Delete(u model.User) (bool, error)
}

//管理员查询用户
func (repo *UserRepository) Get(user model.User) (*model.User, error) {
	if err := repo.DB.Where(&user).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

//根据用户名查询用户是否存在
func (repo *UserRepository) Exist(user model.User) *model.User {
	var count int
	repo.DB.Find(&user).Where("user_name = ?", user.UserName)
	if count > 0 {
		return &user
	}
	return nil
}

//根据用户ID查询用户是否存在
func (repo *UserRepository) ExistByUserID(id string) *model.User {
	var user model.User
	repo.DB.Where("user_id = ?", id).First(&user)
	return &user
}

//添加用户
func (repo *UserRepository) Add(user model.User) (*model.User, error) {
	if exist := repo.Exist(user); exist != nil {
		return nil, fmt.Errorf("用户注册已存在")
	}
	err := repo.DB.Create(&user).Error
	if err != nil {
		return nil, fmt.Errorf("用户注册失败")
	}
	return &user, nil
}

//修改用户
func (repo *UserRepository) Edit(user model.User) (bool, error) {
	err := repo.DB.Model(&user).Where("user_id=?", user.UserId).Updates(map[string]interface{}{
		"user_name": user.UserName,
		"address": user.Address,
		"password": user.Password,
	}).Error
	//err := repo.DB.Save(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

//删除用户
func (repo *UserRepository) Delete(user model.User) (bool, error) {
	err := repo.DB.Model(&user).Where("user_id=?", user.UserId).Delete(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}