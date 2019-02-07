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

type QueryLikeStatusParams struct {
	PostId int `json:"post_id"`
}

func checkQueryLikeStatus(info *QueryLikeStatusParams) (err error) {
	if info.PostId <= 0 {
		err = errors.New("body中参数非法")
		return
	}
	return
}

type QueryLikeStatusReply struct {
	Status int `json:"status"`
}

//点赞帖子
func QueryLikeStatus(c echo.Context) (err error) {
	req := new(QueryLikeStatusParams)
	err = c.Bind(req)
	if err != nil {
		fmt.Println("bind err")
		return err
	}
	data := new(QueryLikeStatusReply)
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
	err = checkQueryLikeStatus(req)
	if err != nil {
		return
	}
	userId := auth.GetSessionInfo(c).UserId
	//handler
	d, err := post.DealQueryLikeStatus(int(userId), req.PostId)
	if err != nil {
		return
	}
	data.Status = d
	return
}
