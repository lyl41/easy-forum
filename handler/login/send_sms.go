package login

import "easy-forum/datestore/redis"

func DealSendSms(phone, picCode string) (smsCode, picCodeReply string, err error) {
	//TODO
	smsCode = "1111"
	redis.SetSmsCode(phone, smsCode)
	return
}
