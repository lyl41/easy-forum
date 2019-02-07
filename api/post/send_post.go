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

type SendPostParams struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func checkSendPost(info *SendPostParams) (err error) {
	if info.Title == "" || info.Content == "" {
		err = errors.New("body中参数非法")
		return
	}
	return
}

func SendPost(c echo.Context) (err error) {
	req := new(SendPostParams)
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
	err = checkSendPost(req)
	if err != nil {
		return
	}
	userId := auth.GetSessionInfo(c).UserId
	//handler
	if err = post.DealSendPost(userId, req.Title, req.Content); err != nil {
		return
	}
	return
}
