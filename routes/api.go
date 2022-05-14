package routes

import (
	"GDColumn/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", suc.SignupUsingEmail)
			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-email", lgc.LoginByEmail)
		}
	}
}