package main

import (
	"easy-forum/api/login"
	"easy-forum/api/post"
	"easy-forum/api/sign"
	"easy-forum/auth"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	//c := common{}
	//http.HandleFunc("/post", c.SendPost)
	//http.HandleFunc("/reply", c.ReplyPost)
	//http.HandleFunc("/send-sms", c.SendSms)
	//http.HandleFunc("/login-by-phone", c.LoginByPhone)
	//http.HandleFunc("/like-post", c.LikePost)
	//http.HandleFunc("/cancel-like-post", c.CancelLikePost)
	//http.HandleFunc("/query-like-status", c.QueryLikeStatus)
	//log.Fatal(http.ListenAndServe(":9001", nil))

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

	e.Start(":9001")
}
