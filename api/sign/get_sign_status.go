package sign

import (
	"easy-forum/auth"
	"easy-forum/common"
	"easy-forum/handler/sign"
	"github.com/labstack/echo"
	"net/http"
)

type GetSignStatusReply struct {
	SignStatus   bool `json:"sign_status"`
	ContinueDays int  `json:"continue_days"`
}

//用户签到接口
func GetSignStatus(c echo.Context) (err error) {

	data := new(GetSignStatusReply)
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
	status, continueDays, err := sign.DealGetSignStatus(userId)
	if err != nil {
		return
	}
	data.SignStatus = status
	data.ContinueDays = int(continueDays)
	return
}
