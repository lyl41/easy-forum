package post

import (
	"easy-forum/auth"
	"easy-forum/common"
	"easy-forum/handler/post"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type CancelLikePostParams struct {
	PostId int `json:"post_id"`
}

func checkCancelLikePost(info *CancelLikePostParams) (err error) {
	if info.PostId <= 0 {
		err = errors.New("body中参数非法")
		return
	}
	return
}

func CancelLikePost(c echo.Context) (err error) {
	req := new(CancelLikePostParams)
	err = c.Bind(req)
	if err != nil {
		fmt.Println("bind err")
		return err
	}
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
	err = checkCancelLikePost(req)
	if err != nil {
		return
	}
	userId := auth.GetSessionInfo(c).UserId
	//handler
	if err = post.DealCancelLikePost(int(userId), req.PostId); err != nil {
		return
	}
	return
}
