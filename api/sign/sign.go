package sign

import (
	"easy-forum/auth"
	"easy-forum/common"
	"easy-forum/handler/sign"
	"github.com/labstack/echo"
	"net/http"
)

//用户签到接口
func Sign(c echo.Context) (err error) {

	data := new(struct{})
	reply := common.StdReply{
		Result: common.ResultFail,
	}
	defer func() {
		if err != nil {
			reply.ErrMsg = err.Error()
		} else {
			reply.Result = common.ResultSuccess
			reply.Data = data
		}
		c.JSON(http.StatusOK, reply)
	}()

	userId := auth.GetSessionInfo(c).UserId
	//handler
	if err = sign.DealSign(userId); err != nil {
		return
	}
	return
}
