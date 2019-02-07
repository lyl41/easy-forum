package redis

import (
	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
)

const smsPrefix = `smscode`

const smsCodeTimeout = 60 * 15 //15分钟

func SetSmsCode(phone, smsCode string) (err error) {
	key := smsPrefix + phone
	conn := pool.Get()
	defer conn.Close()
	if _, err = conn.Do("setex", key, smsCodeTimeout, smsCode); err != nil {
		err = errors.Wrap(err, "setSmsCode redis setex fail.")
		return
	}
	return
}

//NOTE：从redis中未找到key时，err返回nil，但是smsCode为空
func GetSmsCode(phone string) (smsCode string, err error) {
	key := smsPrefix + phone
	conn := pool.Get()
	defer conn.Close()
	smsCode, err = redis.String(conn.Do("get", key))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
			smsCode = ""
			return
		}
		err = errors.Wrap(err, "GetSmsCode redis get fail.")
		return
	}
	return
}
