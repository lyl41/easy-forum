package login

import (
	"easy-forum/app/common"
	"easy-forum/app/handler/login"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type loginByPhone struct {
	Phone      string `json:"phone"`
	VerifyCode string `json:"verify_code"`
}

type loginReply struct {
	Token string `json:"token"`
}

func checkLoginByPhone(info *loginByPhone) (err error) {
	if info.Phone == "" || info.VerifyCode == "" {
		err = errors.New("body中参数非法")
		return
	}
	return
}

func LoginByPhone(c echo.Context) (err error) {
	req := new(loginByPhone)
	err = c.Bind(req)
	if err != nil {
		return
	}
	data := new(loginReply)
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
	err = checkLoginByPhone(req)
	if err != nil {
		return
	}
	token, err := login.DealLoginByPhone(req.Phone, req.VerifyCode)
	if err != nil {
		return
	}
	data.Token = token
	return
}
