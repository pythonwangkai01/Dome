package core

import (
	"context"
	"errors"
	"net/http"
	"user/model"
	userpb "user/services"

	"github.com/jinzhu/gorm"
)

func BuildUser(item model.User) *userpb.UserModel {
	userModel := userpb.UserModel{
		Id:        uint32(item.ID),
		Uid:       uint64(item.Uid),
		UserName:  item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
		Phone:     item.Phone,
		Address:   item.Address,
		Status:    int32(item.Status),
	}
	return &userModel
}

func (*UserService) UserLogin(ctx context.Context, req *userpb.UserRequest, resp *userpb.UserDetailResponse) error {
	var user model.User
	resp.Code = http.StatusOK
	if err := model.DB.Where("user_name=?", req.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			resp.Code = http.StatusBadRequest
			err := errors.New("数据库查询失败")
			return err
		}
		resp.Code = 500
		err := errors.New("username查询失败")
		return err
	}
	//校验密码
	if user.CheckPassWord(req.Password) == false {
		resp.Code = http.StatusBadRequest
		err := errors.New("密码错误")
		return err
	}
	resp.UserDetail = BuildUser(user)
	resp.Code = http.StatusOK
	return nil
}

func (*UserService) UserRegister(ctx context.Context, req *userpb.UserRequest, resp *userpb.UserDetailResponse) error {
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码不一致")
		return err
	}
	count := 0
	if err := model.DB.Model(&model.User{}).Where("user_name=?", req.UserName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		err := errors.New("用户名已存在")
		return err
	}
	userdata := model.User{
		UserName: req.UserName,
		Phone:    req.Phone,
		Address:  req.Address,
		Uid:      uint(req.Uid),
		Status:   1, //正常：1 封号：2 删除：3
	}
	//加密密码
	err := userdata.SetPassWord(req.Password)
	if err != nil {
		return err
	}
	if err := model.DB.Create(&userdata).Error; err != nil {
		return err
	}
	resp.UserDetail = BuildUser(userdata)
	resp.Code = http.StatusOK
	return nil

}

//获取用户列表
func (*UserService) GetUsersList(ctx context.Context, req *userpb.UserRequest, resp *userpb.UserListResponse) error {
	if req.Limit == 0 {
		req.Limit = 10
	}
	var userData []model.User
	var count uint32

	//查找用户
	//指定获取记录的最大值 offset 指定在开始返回记录之前要跳过的记录数量
	err := model.DB.Offset(req.Start).Limit(req.Limit).Find(&userData).Error
	if err != nil {
		return errors.New("mysql find:" + err.Error())
	}
	//返回protoc数据
	var userRes []*userpb.UserModel
	for _, item := range userData {
		userRes = append(userRes, BuildUser(item))
		count++
	}
	resp.UserList = userRes
	resp.Count = count
	return nil
}

func (*UserService) GetUser(ctx context.Context, req *userpb.UserRequest, resp *userpb.UserDetailResponse) error {
	userData := model.User{}
	model.DB.First(&userData, req.Id)
	userRes := BuildUser(userData)
	resp.UserDetail = userRes
	resp.Code = http.StatusOK
	return nil
}
