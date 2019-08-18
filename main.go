package main

import (
	"easy-forum/app/api/login"
	"easy-forum/app/api/post"
	"easy-forum/app/api/sign"
	"easy-forum/app/auth"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(auth.HookErr)

	e.POST("/post", post.SendPost, auth.VerifyAuth)
	e.POST("/reply-post", post.ReplyPost, auth.VerifyAuth)
	e.POST("/like-post", post.LikePost, auth.VerifyAuth)
	e.POST("/cancel-like-post", post.CancelLikePost, auth.VerifyAuth)
	e.POST("/query-like-status", post.QueryLikeStatus, auth.VerifyAuth)

	e.POST("/send-sms", login.SendSms)
	e.POST("login-by-phone", login.LoginByPhone)

	e.POST("/sign", sign.Sign, auth.VerifyAuth)
	e.POST("/get-sign-status", sign.GetSignStatus, auth.VerifyAuth)

	e.Start("0.0.0.0:9001") //TODO
}
