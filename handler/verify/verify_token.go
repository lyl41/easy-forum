package verify

import (
	"easy-forum/datestore/redis"
	"github.com/pkg/errors"
)

//验证接口token，返回userid
func VerifyToken(token string) (userId int, err error) {
	sess, err := redis.GetTokenValue(token)
	if err != nil {
		err = errors.Wrap(err, "verify token failed.")
		return
	}
	if sess == nil {
		err = errors.New("登录信息过期，请重新登录^ ^")
		return
	}
	userId = int(sess.UserId)
	return
}
