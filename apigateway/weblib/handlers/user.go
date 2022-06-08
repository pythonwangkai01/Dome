package handlers

import (
	"apigateway/pkg/utils"
	userpb "apigateway/services"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

//用户注册
func UserRegister(ctx *gin.Context) {
	var userReq userpb.UserRequest
	PanicIfUserError(ctx.ShouldBind(&userReq))
	//gin.key中取出服务实例
	us := ctx.Keys["UserService"].(userpb.UserService)
	//生成uid
	userReq.Uid = uint64(utils.UniqueId())
	//context.Background() 返回一个空的Context,我们可以用这个 空的 Context 作为 goroutine 的root 节点
	udr, err := us.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	ctx.JSON(http.StatusOK, gin.H{
		"code": udr.Code,
		"msg":  "注册成功",
		"data": udr.UserDetail,
	})
}

//用户登录
func UserLogin(ctx *gin.Context) {
	var userReq userpb.UserRequest
	PanicIfUserError(ctx.ShouldBind(&userReq))
	//gin.key中取出服务实例
	us := ctx.Keys["UserService"].(userpb.UserService)
	//context.Background() 返回一个空的Context,我们可以用这个 空的 Context 作为 goroutine 的root 节点
	udr, err := us.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)

	token, err := utils.GenerateToken(uint(udr.UserDetail.Uid))
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

//获取用户列表
func GetUsersList(ctx *gin.Context) {
	var userReq userpb.UserRequest
	PanicIfUserError(ctx.BindJSON(&userReq))
	//验证token 获取uid
	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	userReq.Uid = uint64(claim.Uid)
	us := ctx.Keys["UserService"].(userpb.UserService)
	udr, err := us.GetUsersList(context.Background(), &userReq)
	PanicIfUserError(err)
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "成功",
		"data": gin.H{
			"user":  udr.UserList,
			"count": udr.Count,
		},
	})

}

func GetUser(ctx *gin.Context) {
	var userReq userpb.UserRequest
	PanicIfUserError(ctx.BindJSON(&userReq))
	//验证token 获取uid
	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	userReq.Uid = uint64(claim.Uid) //查询人的uid
	us := ctx.Keys["UserService"].(userpb.UserService)
	udr, err := us.GetUser(context.Background(), &userReq)
	PanicIfUserError(err)
	ctx.JSON(http.StatusOK, gin.H{
		"code": udr.Code,
		"msg":  "成功",
		"data": udr.UserDetail,
	})
}

// admin***************************************************

//用户删除
func AdminUserDelte(ctx *gin.Context) {
	var userReq userpb.UserRequest
	PanicIfUserError(ctx.ShouldBind(&userReq))
	//验证token 获取id和uid
	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	userReq.Uid = uint64(claim.Uid)

	us := ctx.Keys["UserService"].(userpb.UserService)
	udr, err := us.AdminUserDelte(context.Background(), &userReq)
	PanicIfUserError(err)
	ctx.JSON(http.StatusOK, gin.H{
		"code": udr.Code,
		"msg":  "删除成功",
		"data": udr.UserDetail,
	})
}

//admin init user
func AdminUserRegister(ctx *gin.Context) {
	var userReq userpb.UserRequest
	PanicIfUserError(ctx.ShouldBind(&userReq))
	//gin.key中取出服务实例
	us := ctx.Keys["UserService"].(userpb.UserService)
	//获取权限uid
	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))
	userReq.CreateUid = uint32(claim.Uid)
	//生成创建admin用户的uid
	userReq.Uid = uint64(utils.UniqueId())
	//context.Background() 返回一个空的Context,我们可以用这个 空的 Context 作为 goroutine 的root 节点
	udr, err := us.AdminUserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	ctx.JSON(http.StatusOK, gin.H{
		"code": udr.Code,
		"msg":  "注册成功",
		"data": udr.UserDetail,
	})
}

func AdminUserLogin(ctx *gin.Context) {
	var userReq userpb.UserRequest
	PanicIfUserError(ctx.ShouldBind(&userReq))
	//gin.key中取出服务实例
	us := ctx.Keys["UserService"].(userpb.UserService)
	//context.Background() 返回一个空的Context,我们可以用这个 空的 Context 作为 goroutine 的root 节点
	udr, err := us.AdminUserLogin(context.Background(), &userReq)
	PanicIfUserError(err)

	token, err := utils.GenerateToken(uint(udr.UserDetail.Uid))
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
