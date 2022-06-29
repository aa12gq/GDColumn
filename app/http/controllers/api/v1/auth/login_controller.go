package auth

import (
	v1 "GDColumn/app/http/controllers/api/v1"
	"GDColumn/app/requests"
	"GDColumn/pkg/auth"
	"GDColumn/pkg/jwt"
	"GDColumn/pkg/response"
	_ "GDColumn/app/models/param"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	v1.BaseAPIController
}

// LoginByEmail 关于用户的路由，登录，注册，获取当前用户等等
// @Summary 发送请求用户登录
// @Description 需要用户名密码
// @Tags user 关于用户的路由，登录，注册，获取当前用户等等
// @Accept application/json
// @Produce application/json
// @Param body body requests.LoginByEmailRequest true "用户登录，需要提供用户的邮箱和密码"
// @Success 200 {json} {"token":token}
//@Router /auth/login/using-email [POST]
func (lc *LoginController) LoginByEmail(c *gin.Context) {

	request := requests.LoginByEmailRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByEmail); !ok {
		return
	}
	user, err := auth.Attempt(request.Email,request.Password)
	if err != nil {
		response.Unauthorized(c, "登录失败，该用户不存在或者密码错误")
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.NickName)
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

func (lc *LoginController) RefreshToken(c *gin.Context) {

	token, err := jwt.NewJWT().RefreshToken(c)

	if err != nil {
		response.Error(c, err, "令牌刷新失败")
	} else {
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}