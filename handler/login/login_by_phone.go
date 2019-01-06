package login

import "github.com/pkg/errors"

func DealLoginByPhone(phone, verifyCode string) (token string, err error) {
	//TODO
	if verifyCode == "1111" {
		return "lyl_token", nil
	} else {
		err = errors.New("验证码错误")
		return
	}
	return
}
