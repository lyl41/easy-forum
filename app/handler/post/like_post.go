package post

import (
	"easy-forum/app/datestore/mysql"

	"github.com/pkg/errors"
)

func DealLikePost(userId, postId int) (err error) {
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
		err = errors.Wrap(errors.New("postId not found"), "该帖子不存在或已经删除")
		return
	}
	if err = mysql.AddNewLikeRecord(tx, int64(userId), int64(postId)); err != nil {
		return
	}
	likeCount := postInfo.Like + 1
	if err = mysql.UpdatePostLikeCount(tx, int64(postId), likeCount); err != nil {
		return
	}

	ok = true
	return
}
