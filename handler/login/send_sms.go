package login

import (
	"easy-forum/datestore/redis"

	"github.com/pkg/errors"
)

func DealSendSms(phone, picCode string) (smsCode, picCodeReply string, err error) {
	//TODO
	smsCode = "1111"
	err = redis.SetSmsCode(phone, smsCode)
	if err != nil {
		err = errors.Wrap(err, "redis SetSmsCode failed")
		return
	}
	return
}
