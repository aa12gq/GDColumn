package auth

import (
	v1 "GDColumn/app/http/controllers/api/v1"
	"GDColumn/app/requests"
	"GDColumn/pkg/auth"
	"GDColumn/pkg/jwt"
	"GDColumn/pkg/response"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	v1.BaseAPIController
}

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