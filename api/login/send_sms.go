package login

import (
	"easy-forum/handler/login"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type sendSms struct {
	Phone   string `json:"phone"`
	PicCode string `json:"pic_code"`
}

type sendSmsReply struct {
	SmsCode string `json:"sms_code"`
	PicCode string `json:"pic_code"`
	Err     error  `json:"error"`
}

func checkSendSms(w http.ResponseWriter, r *http.Request) (info *sendSms, err error) {
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
	info = new(sendSms)
	if err = json.Unmarshal(body, info); err != nil {
		err = errors.Wrap(err, "json解析错误")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("body中不是一个合法的json"))
		return
	}
	fmt.Println(string(body))
	if info.Phone == ""{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("body中参数非法"))
		err = errors.New("body中参数非法")
		return
	}
	return
}

func (Login) SendSms(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-----new request-----")
	var err error
	reply := new(sendSmsReply)
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
	info, err := checkSendSms(w, r)
	if err != nil {
		return
	}
	smsCode, picCode, err := login.DealSendSms(info.Phone, info.PicCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	reply.SmsCode = smsCode
	reply.PicCode = picCode
}
