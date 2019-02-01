package post

import (
	"easy-forum/datestore/mysql"
	"fmt"
	"github.com/pkg/errors"
)

//参数检查要在api层做好，这里保证参数不为空
func DealSendPost(userId int64, title, content string) (err error) {
	defer func() {
		if err != nil {
			fmt.Println("handler层：DealSendPost err:", err)
		}
	}()
	if info, _ := mysql.FindPostByUserIdAndTitle(mysql.GetDB(), userId, title); info != nil && info.ID > 0 { //找到记录的话，ID应该是 > 0
		err = errors.New("您先前已经发布了一篇相同标题的帖子")
		return
	}
	if err = mysql.AddNewPostRecord(userId, title, content); err != nil {
		err = errors.Wrap(err, "发布帖子失败，请稍后重试")
		return
	}
	return
}
