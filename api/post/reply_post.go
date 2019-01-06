package post

import "net/http"

func (p Post) ReplyPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!@!"))
}
