package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"time"
)

//通过如下命令自动生成:
///var/folders/bm/5pxrd6855nj81734_7dx3jn80000gn/T/go-build734321937/b001/exe/main -t user_like_post -db root:1234567890@tcp(localhost:3306)/forum

type UserLikePost struct {
	ID        int64     `sql:"primary_key;column:id" json:"id,omitempty"`     //
	UserID    int64     `sql:"column:user_id" json:"user_id,omitempty"`       //
	PostID    int64     `sql:"column:post_id" json:"post_id,omitempty"`       //
	Like      int64     `sql:"column:like" json:"like,omitempty"`             //0保留，1表示点赞，2表示取消点赞
	CreatedAt time.Time `sql:"column:created_at" json:"created_at,omitempty"` //
	UpdatedAt time.Time `sql:"column:updated_at" json:"updated_at,omitempty"` //

}

const (
	UserLikePostStatus       = 1
	UserCancelLikePostStatus = 2
)

func (UserLikePost) TableName() string {
	return "user_like_post"
}

func AddNewLikeRecord(tx *gorm.DB, userId, postId int64) (err error) {
	data := new(UserLikePost)

	where := &UserLikePost{
		UserID: userId,
		PostID: postId,
	}
	if err = tx.Set("gorm:query_option", " FOR UPDATE ").Where(where).First(data).Error; err != nil && err != gorm.ErrRecordNotFound {
		err = errors.Wrap(err, "点赞数据库查找错误")
		return
	}
	//err == gorm.ErrRecordNotFound || err == nil
	if err == gorm.ErrRecordNotFound {
		data.UserID = userId
		data.PostID = postId
		data.Like = UserLikePostStatus
		err = tx.Create(data).Error
		if err != nil {
			err = errors.Wrap(err, "创建点赞记录失败")
			return
		}
	} else {
		if data.Like == UserLikePostStatus {
			err = errors.Wrap(errors.New("重复点赞"), "您已经点过赞啦")
			return
		}
		data.UserID = userId
		data.PostID = postId
		data.Like = UserLikePostStatus
		err = tx.Model(data).Where(where).Updates(data).Error
		if err != nil {
			err = errors.Wrap(err, "更新点赞记录失败")
			return
		}
	}
	return
}
func CancelLikeRecord(tx *gorm.DB, userId, postId int64) (err error) {
	data := new(UserLikePost)

	where := &UserLikePost{
		UserID: userId,
		PostID: postId,
	}
	if err = tx.Set("gorm:query_option", " FOR UPDATE ").Where(where).First(data).Error; err != nil && err != gorm.ErrRecordNotFound {
		err = errors.Wrap(err, "点赞数据库查找错误")
		return
	}
	//err == gorm.ErrRecordNotFound || err == nil
	if err == gorm.ErrRecordNotFound {
		data.UserID = userId
		data.PostID = postId
		data.Like = UserCancelLikePostStatus
		err = tx.Create(data).Error
		if err != nil {
			err = errors.Wrap(err, "创建取消点赞记录失败")
			return
		}
	} else {
		if data.Like == UserCancelLikePostStatus { //TODO
			err = errors.Wrap(errors.New("重复取消点赞"), "您已经取消点赞啦")
			return
		}
		data.UserID = userId
		data.PostID = postId
		data.Like = UserCancelLikePostStatus
		err = tx.Model(data).Where(where).Updates(data).Error
		if err != nil {
			err = errors.Wrap(err, "更新取消点赞记录失败")
			return
		}
	}
	return
}
