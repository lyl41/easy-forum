package post

import (
	"easy-forum/datestore/mysql"
	"github.com/pkg/errors"
)

func DealReplyPost(userId, postId int, content string) (err error) {
	//先在post表中找帖子
	tx := mysql.GetDB().Begin()
	if err = tx.Error; err != nil {
		err = errors.Wrap(err, "获取数据库事务失败")
		return
	}
	ok := false
	defer func() {
		if !ok {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	postInfo, err := mysql.FindPostByID(tx, userId)
	if err != nil {
		return
	}
	if postInfo != nil && postInfo.ID <= 0 {
		err = errors.Wrap(errors.New("postId not found"), "帖子不存在或已经删除")
		return
	}


}
