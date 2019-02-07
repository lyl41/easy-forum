package auth

import (
	"easy-forum/common"
	"net/http"

	"github.com/labstack/echo"
)

type SendPostParams struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func HookErr(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		err := next(ctx)
		if ctx.Response().Committed {
			return err
		}
		if err != nil {
			reply := common.StdReply{
				Result: common.ResultFail,
				ErrMsg: err.Error(),
			}
			ctx.JSON(http.StatusOK, reply)
		}
		return err
	}
}
