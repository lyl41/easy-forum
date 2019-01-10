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
	postInfo, err := mysql.FindPostByID(tx, int64(postId))
	if err != nil {
		return
	}
	if postInfo != nil && postInfo.ID <= 0 {
		err = errors.Wrap(errors.New("postId not found"), "帖子不存在或已经删除")
		return
	}
	floor := postInfo.ReplyCount + 1 //所在楼层
	//创建帖子的回复记录
	err = mysql.CreateReplyPostRecord(tx, postInfo.ID, floor, postInfo.UserID, int64(userId), content)
	if err != nil {
		return
	}
	//更新回复数量
	err = mysql.UpdatePostReplyCount(tx, postInfo.ID, floor)
	if err != nil {
		return
	}

	ok = true
	return
}
