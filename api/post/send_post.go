package post

import (
	"easy-forum/api/verify"
	"easy-forum/handler/post"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type SendPostParams struct {
	Title   string `json:"title`
	Content string `json:"content"`
}

func checkSendPost(w http.ResponseWriter, r *http.Request) (info *SendPostParams, err error) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(fmt.Sprint("不支持%s方法", r.Method)))
		err = errors.Errorf("不支持%s方法", r.Method)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	info = new(SendPostParams)
	if err = json.Unmarshal(body, info); err != nil {
		err = errors.Wrap(err, "json解析错误")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("body中不是一个合法的json"))
		return
	}
	if info.Title == "" || info.Content == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("body中参数非法"))
		err = errors.New("body中参数非法")
		return
	}
	return
}

func (p Post) SendPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-----new request-----")
	var err error
	defer func() {
		if err != nil {
			fmt.Println("api层SendPost err:", err)
		}
		fmt.Println("-----request end-----")
	}()
	info, err := checkSendPost(w, r)
	if err != nil {
		return
	}
	var token string
	fmt.Println(r.Header)
	if val, found := r.Header["Token"]; !found || len(val) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("header中无Token"))
		return
	} else {
		token = val[0]
	}
	userId, err := verify.VerifyToken(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("token失效"))
		return
	}
	if err = post.DealSendPost(userId, info.Title, info.Content); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("服务内部错误"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("send-post success"))

}
