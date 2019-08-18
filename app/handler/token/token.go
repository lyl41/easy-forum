package token

import (
	"crypto/md5"
	"easy-forum/app/common/util"
	"easy-forum/app/datestore/redis"
	"encoding/hex"
	"strconv"

	"github.com/pkg/errors"
)

func genTokenByUserID(userId int) (token string, err error) {
	//TODO 临时先随机生成，后面考虑生成规则
	tmp := util.RandomStr(13) + ":" + strconv.Itoa(userId)
	md5Byte := md5.Sum([]byte(tmp))
	return hex.EncodeToString(md5Byte[:]), nil
}

func DealToken(userId int) (token string, err error) {
	if userId <= 0 {
		err = errors.New("DealToken, userid <= 0")
		return
	}
	tmpToken, err := genTokenByUserID(userId)
	if err != nil {
		err = errors.Wrap(err, "DealToken生成token失败")
		return
	}
	sessInfo := &redis.SessionInfo{
		UserId: int64(userId),
	}
	if err = redis.SetTokenValue(tmpToken, sessInfo); err != nil {
		err = errors.Wrap(err, "redis set token失败")
		return
	}
	token = tmpToken
	return
}
