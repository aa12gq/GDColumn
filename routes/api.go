package routes

import (
	controllers "GDColumn/app/http/controllers/api/v1"
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

		uc := new(controllers.UsersController)
		usersGroup := v1.Group("/users")
		v1.GET("/current", middlewares.AuthJWT(),middlewares.Cors, uc.CurrentUser)
		{
			usersGroup.GET("", uc.Index)
			usersGroup.PUT("", middlewares.AuthJWT(), uc.UpdateProfile)
			usersGroup.PUT("/email", middlewares.AuthJWT(), uc.UpdateEmail)
			usersGroup.PUT("/phone", middlewares.AuthJWT(), uc.UpdatePhone)
			usersGroup.PUT("/password", middlewares.AuthJWT(), uc.UpdatePassword)
		}
		imc := new(controllers.ImagesController)
		imcGroup := v1.Group("/upload")
		{
			imcGroup.POST("",middlewares.LimitPerRoute("20-H"),imc.Upload)
		}
		clc := new(controllers.ColumnsController)
		clcGroup := v1.Group("/columns")
		{
			clcGroup.GET("", clc.Index)
			clcGroup.GET("/:id", middlewares.AuthJWT(), clc.CurrentColumn)
			clcGroup.POST("", middlewares.AuthJWT(), clc.Store)
			clcGroup.PUT("", middlewares.AuthJWT(), clc.Update)
			clcGroup.DELETE("/:id", middlewares.AuthJWT(), clc.Delete)
		}
		poc := new(controllers.PostsController)
		pocGroup := v1.Group("/posts")
		{
			pocGroup.POST("", middlewares.AuthJWT(), poc.Store)
			pocGroup.PUT("/:id",middlewares.AuthJWT(),poc.Update)
			pocGroup.DELETE("/:id", middlewares.AuthJWT(), poc.Delete)
			pocGroup.GET("/:id", poc.Show)
		}

	}
}