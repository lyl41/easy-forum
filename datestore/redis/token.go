package redis

import (
	"encoding/json"

	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
)

const keyPreFix = `token:`

const tokenTimeout = 60 * 60 //1小时过期

type SessionInfo struct {
	UserId int64 `json:"user_id"`
}

func SetTokenValue(token string, sess *SessionInfo) (err error) {
	str, err := json.Marshal(sess)
	if err != nil {
		err = errors.Wrap(err, "redis SetTokenValue, json非法")
		return
	}
	key := keyPreFix + token
	conn := pool.Get()
	defer conn.Close()
	defer conn.Close()
	if _, err = conn.Do("setex", key, tokenTimeout, string(str)); err != nil {
		err = errors.Wrap(err, "redis token setex fail.")
		return
	}
	return
}

//NOTE: 如果redis中不存在key，返回nil, nil, 根据sess判断，外面就不用引入redigo/redis包了
func GetTokenValue(token string) (sess *SessionInfo, err error) {
	key := keyPreFix + token
	conn := pool.Get()
	defer conn.Close()
	reply, err := redis.String(conn.Do("get", key))
	if err != nil {
		if err == redis.ErrNil {
			return nil, nil
		}
		err = errors.Wrap(err, "redis token get fail.")
		return
	}
	sess = new(SessionInfo) //记得new
	if err = json.Unmarshal([]byte(reply), sess); err != nil {
		err = errors.Wrap(err, "redis token get unmarshal fail.")
		return
	}
	return
}
