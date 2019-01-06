package login

import (
	"easy-forum/handler/login"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type loginByPhone struct {
	Phone      string `json:"phone"`
	VerifyCode string `json:"verify_code"`
}

type loginReply struct {
	Token string `json:"token"`
	Err   error  `json:"error"`
}

func checkLoginByPhone(w http.ResponseWriter, r *http.Request) (info *loginByPhone, err error) {
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
	info = new(loginByPhone)
	if err = json.Unmarshal(body, info); err != nil {
		err = errors.Wrap(err, "json解析错误")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("body中不是一个合法的json"))
		return
	}
	fmt.Println(string(body))
	if info.Phone == "" || info.VerifyCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("body中参数非法"))
		err = errors.New("body中参数非法")
		return
	}
	return
}

func (Login) LoginByPhone(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-----new request-----")
	var err error
	reply := new(loginReply)
	defer func() {
		if err != nil {
			fmt.Println("api层SendPost err:", err)
			w.Write([]byte(err.Error()))
		} else {
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
	info, err := checkLoginByPhone(w, r)
	if err != nil {
		return
	}
	token, err := login.DealLoginByPhone(info.Phone, info.VerifyCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	reply.Token = token
}
