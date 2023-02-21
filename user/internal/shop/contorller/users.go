package contorller

import (
	"context"
	"github.com/gin-gonic/gin"
	userpb "shop/user/api-gen/v1"
	"shop/user/internal/pkg/core"
	"shop/user/internal/pkg/errno"
	"shop/user/internal/shop/biz"
)

type UserController struct {
	*userpb.UnimplementedUserServiceServer
	biz biz.IBiz
}

func NewUserController(biz biz.IBiz) *UserController {
	return &UserController{biz: biz}
}

func (uc *UserController) Create(c *gin.Context) {
	var request userpb.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
	}
	var ctx context.Context
	if err := uc.biz.Users().Create(ctx, &request); err != nil {
		core.WriteResponse(c, err, nil)
	}
	core.WriteResponse(c, nil, nil)
}

func (uc *UserController) CreateUser(ctx context.Context, request *userpb.CreateUserRequest) (*userpb.UserInfo, error) {
	return nil, nil
}

// PasswordLogin 使用密码登陆
func (uc *UserController) PasswordLogin(c *gin.Context) {
	var r userpb.CheckPassWordRequest
	// 1. 表单验证 使用gin 可以用vailador 自定义 来验证请求参数
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
	}
	// 2. 对比图形验证码
	// 3. 查询用户是否存在
	// 4. 对比加密的密码
}
