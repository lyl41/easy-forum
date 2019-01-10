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

type SendPostParams struct {
	Title   string `json:"title"`
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
	reply := new(HttpReply)
	defer func() {
		if err != nil {
			fmt.Println("api层SendPost err:", err)
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
	//参数检查
	info, err := checkSendPost(w, r)
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
	//handler
	if err = post.DealSendPost(userId, info.Title, info.Content); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func getTokenFromHeader(r *http.Request)(token string, err error)  {
	if val, found := r.Header["Token"]; !found || len(val) == 0 {
		err = errors.New("未找到token")
		return
	} else {
		token = val[0]
	}
	return
}