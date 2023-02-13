package contorller

import (
	"github.com/gin-gonic/gin"
	"shop/user/internal/shop/biz"
)

type UserController struct {
	biz biz.IBiz
}

func NewUserController(biz biz.IBiz) *UserController {
	return &UserController{biz: biz}
}

func (uc *UserController) Create(c *gin.Context) {

}
