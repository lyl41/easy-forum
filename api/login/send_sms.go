package login

import (
	"easy-forum/common"
	"easy-forum/handler/login"
	"fmt"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"net/http"
)

type SendSmsParams struct {
	Phone   string `json:"phone"`
	PicCode string `json:"pic_code"`
}

type SendSmsReply struct {
	SmsCode string `json:"sms_code"`
	PicCode string `json:"pic_code"`
}

func checkSendSms(info *SendSmsParams) (err error) {
	if info.Phone == "" {
		err = errors.New("body中参数非法")
		return
	}
	return
}

func SendSms(c echo.Context) (err error) {
	req := new(SendSmsParams)
	err = c.Bind(req)
	if err != nil {
		fmt.Println("bind err")
		return err
	}
	data := new(SendSmsReply)
	reply := common.StdReply{
		Result: common.ResultFail,
	}
	defer func() {
		if err != nil {
			reply.ErrMsg = err.Error()
		} else {
			reply.Result = common.ResultSuccess
			reply.Data = data
		}
		c.JSON(http.StatusOK, reply)
	}()
	err = checkSendSms(req)
	if err != nil {
		return
	}
	//handler
	smsCode, picCode, err := login.DealSendSms(req.Phone, req.PicCode)
	if err != nil {
		return
	}
	data.PicCode = picCode
	data.SmsCode = smsCode
	return
}
