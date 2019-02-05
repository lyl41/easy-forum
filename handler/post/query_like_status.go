package post

import "easy-forum/datestore/mysql"

func DealQueryLikeStatus(userId, postId int) (status int, err error) {
	status, err = mysql.QueryUserLikeStatus(int64(userId), int64(postId))
	if err != nil {
		return
	}
	return
}
