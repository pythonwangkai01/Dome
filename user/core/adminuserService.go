package core

import (
	"context"
	"errors"
	"net/http"
	"user/model"
	userpb "user/services"

	"github.com/jinzhu/gorm"
)

func BuildAdminUser(item model.AdminUser) *userpb.UserModel {
	userModel := userpb.UserModel{
		Id:        uint32(item.ID),
		Uid:       uint64(item.Uid),
		CreateUid: uint32(item.CreateUid),
		UserName:  item.UserName,
		Phone:     item.Phone,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &userModel
}

//admin用户登录
func (*UserService) AdminUserLogin(ctx context.Context, req *userpb.UserRequest, resp *userpb.UserDetailResponse) error {
	var adminuser model.AdminUser
	resp.Code = http.StatusOK
	if err := model.DB.Where("user_name=?", req.UserName).First(&adminuser).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			resp.Code = http.StatusBadRequest
			return nil
		}
		resp.Code = 500
		return nil
	}
	//校验密码
	if adminuser.CheckPassWord(req.Password) == false {
		resp.Code = http.StatusBadRequest
		return nil
	}
	resp.UserDetail = BuildAdminUser(adminuser)
	resp.Code = http.StatusOK
	return nil
}

//创建admin用户
func (*UserService) AdminUserRegister(ctx context.Context, req *userpb.UserRequest, resp *userpb.UserDetailResponse) error {
	//只有admin权限能创建admin
	if err := model.DB.Model(&model.AdminUser{}).Where("create_uid=?", req.CreateUid).First(&model.AdminUser{}).Error; err != nil {
		err := errors.New("无权限操作")
		return err
	}

	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码不一致")
		return err
	}
	count := 0
	if err := model.DB.Model(&model.AdminUser{}).Where("user_name=?", req.UserName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		err := errors.New("用户名已存在")
		return err
	}
	userdata := model.AdminUser{
		UserName:  req.UserName,
		Phone:     req.Phone,
		Uid:       uint(req.Uid),
		CreateUid: uint(req.CreateUid),
	}
	//加密密码
	err := userdata.SetPassWord(req.Password)
	if err != nil {
		return err
	}
	if err := model.DB.Create(&userdata).Error; err != nil {
		return err
	}

	resp.UserDetail = BuildAdminUser(userdata)
	resp.Code = http.StatusOK
	return nil

}

//上层加个手机验证吧（TODO），这里只处理数据层的逻辑
func (*UserService) AdminUserDelte(ctx context.Context, req *userpb.UserRequest, resp *userpb.UserDetailResponse) error {

	if err := model.DB.Model(&model.AdminUser{}).Where("uid=?", req.Uid).Error; err != nil {
		err := errors.New("无权限操作")
		return err
	}
	//删除需要admin权限到上面一层做
	err := model.DB.Model(&model.AdminUser{}).Where("uid=?", req.Uid).Find(&model.AdminUser{}).Error
	if err != nil {
		return errors.New("删除失败" + err.Error())

	}
	model.DB.Where("user_name=?", req.UserName).Delete(&model.User{})
	resp.Code = http.StatusOK
	return nil
}
