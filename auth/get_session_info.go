package auth

import (
	"easy-forum/datestore/redis"

	"github.com/labstack/echo"
)

func GetSessionInfo(ctx echo.Context) (sess *redis.SessionInfo) {
	sessTmp := ctx.Get(SessionInfoKey)
	if sessTmp == nil {
		return nil
	}
	sess, ok := sessTmp.(*redis.SessionInfo)
	if !ok {
		return nil
	}
	return
}
