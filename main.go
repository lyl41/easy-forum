package main

import (
	"easy-forum/api/post"
	"log"
	"net/http"
)

func main() {
	post := post.Post{}
	http.HandleFunc("/post", post.SendPost)
	http.HandleFunc("/reply", post.ReplyPost)
	log.Fatal(http.ListenAndServe(":9001", nil))
}
