package handlers

import (
	"apigateway/pkg/utils"
	userpb "apigateway/services"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//用户注册
func UserRegister(ctx *gin.Context) {
	var userReq userpb.UserRequest
	PanicIfUserError(ctx.Bind(&userReq))
	//gin.key中取出服务实例
	us := ctx.Keys["UserService"].(userpb.UserService)
	//s生成uid
	userReq.Uid = uint64(utils.UniqueId())
	//context.Background() 返回一个空的Context,我们可以用这个 空的 Context 作为 goroutine 的root 节点
	udr, err := us.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	ctx.JSON(http.StatusOK, gin.H{
		"data": udr,
	})
}

//用户登录
func UserLogin(ctx *gin.Context) {
	var userReq userpb.UserRequest
	PanicIfUserError(ctx.Bind(&userReq))
	//gin.key中取出服务实例
	us := ctx.Keys["UserService"].(userpb.UserService)
	//context.Background() 返回一个空的Context,我们可以用这个 空的 Context 作为 goroutine 的root 节点
	udr, err := us.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)

	token, err := utils.GenerateToken(uint(udr.UserDetail.Id))
	PanicIfUserError(err)
	ctx.JSON(http.StatusOK, gin.H{
		"code": udr.Code,
		"msg":  "成功",
		"data": gin.H{
			"user":  udr.UserDetail,
			"token": token,
		},
	})
}

//用户删除
func UserDelte(ctx *gin.Context) {
	var userReq userpb.UserRequest
	PanicIfUserError(ctx.Bind(&userReq))
	//验证token 获取id和uid
	// claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	uid, _ := strconv.Atoi(ctx.PostForm("uid"))
	userReq.Id = uint64(id)
	userReq.Uid = uint64(uid)
	us := ctx.Keys["UserService"].(userpb.UserService)
	udr, err := us.UserDelte(context.Background(), &userReq)
	PanicIfUserError(err)
	ctx.JSON(http.StatusOK, gin.H{
		"data": udr.UserDetail,
	})
}

//获取用户列表
func GetUsersList(ctx *gin.Context) {
	var userReq userpb.UserRequest
	PanicIfUserError(ctx.Bind(&userReq))
	//验证token 获取uid
	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	userReq.Uid = uint64(claim.Id)
	us := ctx.Keys["UserService"].(userpb.UserService)
	udr, err := us.GetUsersList(context.Background(), &userReq)
	PanicIfUserError(err)
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user":  udr.UserList,
			"count": udr.Count,
		},
	})
}

func GetUser(ctx *gin.Context) {
	var userReq userpb.UserRequest
	PanicIfUserError(ctx.Bind(&userReq))
	//验证token 获取uid
	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	id, _ := strconv.Atoi(ctx.PostForm("id"))
	userReq.Id = uint64(id)
	userReq.Uid = uint64(claim.Id) //查询人的uid
	us := ctx.Keys["UserService"].(userpb.UserService)
	udr, err := us.GetUser(context.Background(), &userReq)
	PanicIfUserError(err)
	ctx.JSON(http.StatusOK, gin.H{
		"data": udr.UserDetail,
	})
}
