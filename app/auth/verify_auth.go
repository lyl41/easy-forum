package auth

import (
	"easy-forum/app/handler/verify"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

const SessionInfoKey = `session_info_key`

//校验header中token是否存在，校验token的合法性，将session信息存储到echo.Context中。
func VerifyAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Request().Header.Get("token")
		if token == "" {
			err := errors.New("token不能为空")
			return err
		}
		sessInfo, err := verify.VerifyToken(token)
		if err != nil {
			return err
		}
		//session已经确保不为nil
		ctx.Set(SessionInfoKey, sessInfo)
		return next(ctx)
	}
}
