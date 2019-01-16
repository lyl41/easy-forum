package main

import (
	"easy-forum/api/login"
	"easy-forum/api/post"
	"log"
	"net/http"
)

type common struct {
	post.Post
	login.Login
}

func main() {
	c := common{}
	http.HandleFunc("/post", c.SendPost)
	http.HandleFunc("/reply", c.ReplyPost)
	http.HandleFunc("/login-by-phone", c.LoginByPhone)
	http.HandleFunc("/like-post", c.LikePost)
	http.HandleFunc("/cancel-like-post", c.CancelLikePost)
	log.Fatal(http.ListenAndServe(":9001", nil))
}
