package login

import (
	"easy-forum/common/util"
	"easy-forum/datestore/mysql"
	"easy-forum/datestore/redis"
	"easy-forum/handler/token"
	"github.com/pkg/errors"
)

func DealLoginByPhone(phone, verifyCode string) (replyToken string, err error) {
	smsCode, err := redis.GetSmsCode(phone)
	if err != nil {
		return
	}
	if smsCode != verifyCode || verifyCode == "" {
		err = errors.New("短信验证码错误")
		return
	}
	//验证通过，判断是否是新用户，新用户就新增users记录
	user := &mysql.Users{
		Phone: phone,
		Name:  "手机用户" + util.RandomStr(10),
	}
	//存在就查询，不存在就创建
	if err = mysql.FirstOrCreateUserByPhone(phone, user); err != nil {
		err = errors.Wrap(err, "mysql.FirstOrCreateUserByPhone err")
		return
	}
	replyToken, err = token.DealToken(int(user.ID))
	if err != nil {
		err = errors.Wrap(err, "DealLoginByPhone token生成失败")
		return
	}
	return
}
