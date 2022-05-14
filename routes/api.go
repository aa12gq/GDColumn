package routes

import (
	"GDColumn/app/http/controllers/api/v1/auth"
	"GDColumn/app/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {

	v1 := r.Group("/v1")
	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup := v1.Group("/auth")
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist",
				middlewares.LimitPerRoute("60-H"), middlewares.GuestJWT(), suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist",
				middlewares.LimitPerRoute("60-H"), middlewares.GuestJWT(), suc.IsEmailExist)
			authGroup.POST("/signup/using-phone",
				middlewares.GuestJWT(), suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email",
				middlewares.GuestJWT(), suc.SignupUsingEmail)

			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-email", middlewares.GuestJWT(), lgc.LoginByEmail)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)
		}
	}
}