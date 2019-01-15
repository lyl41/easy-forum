package post

import (
	"easy-forum/handler/post"
	"easy-forum/handler/verify"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type CancelLikePostParams struct {
	PostId  int `json:"post_id"`
}

func checkCancelLikePost(w http.ResponseWriter, r *http.Request) (info *CancelLikePostParams, err error) {
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
	info = new(CancelLikePostParams)
	if err = json.Unmarshal(body, info); err != nil {
		err = errors.Wrap(err, "json解析错误")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("body中不是一个合法的json"))
		return
	}
	if info.PostId <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("body中参数非法"))
		err = errors.New("body中参数非法")
		return
	}
	return
}
//取消点赞帖子
func (p Post) CancelLikePost(w http.ResponseWriter, r *http.Request) {
	var err error
	reply := new(HttpReply)
	defer func() {
		if err != nil {
			fmt.Println("api层CancelLikePost err:", err)
			w.Write([]byte(err.Error()))
		} else {
			reply.Msg = "请求成功"
			ret, err := json.Marshal(reply)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				fmt.Println("json marshal fail.")
				return
			}
			w.WriteHeader(http.StatusOK)
			fmt.Println(string(ret))
			w.Write(ret)
		}
		fmt.Println("-----request end-----")
	}()
	info, err := checkCancelLikePost(w, r)
	if err != nil {
		return
	}
	//取出token
	token, err := getTokenFromHeader(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//根据token获取userid，根据userid操作数据库
	userId, err := verify.VerifyToken(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err = post.DealCancelLikePost(userId, info.PostId)
	if err != nil {
		return
	}
}