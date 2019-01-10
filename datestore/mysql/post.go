package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"time"
)

type Post struct {
	ID         int64     `sql:"primary_key;column:id" json:"id,omitempty"`       //
	Title      string    `sql:"column:title" json:"title,omitempty"`             //帖子标题
	UserID     int64     `sql:"column:user_id" json:"user_id,omitempty"`         //作者的id
	Detail     string    `sql:"column:detail" json:"detail,omitempty"`           //帖子内容，一楼内容
	Like       int64     `sql:"column:like" json:"like,omitempty"`               //帖子点赞数
	ReplyCount int64     `sql:"column:reply_count" json:"reply_count,omitempty"` //帖子回复数量/热度
	CreatedAt  time.Time `sql:"column:created_at" json:"created_at,omitempty"`   //
	UpdatedAt  time.Time `sql:"column:updated_at" json:"updated_at,omitempty"`   //

}

func (Post) TableName() string {
	return "post"
}

//根据userid或者title找获取帖子数据，title可以为空, 如果未找到数据，err为nil，同时data中 id == 0
func FindPostByUserIdAndTitle(tx *gorm.DB, userId int, title string) (data *Post, err error) {
	where := &Post{
		UserID: int64(userId),
	}
	if title != "" {
		where.Title = title
	}
	data = new(Post)
	if err = tx.Where(where).First(data).Error; err != nil {
		if err == gorm.ErrRecordNotFound { //处理了未找到的情况，err = nil
			err = nil
		}
		return
		err = errors.Wrap(err, "数据库查找错误")
	}
	return
}

func FindPostByID(tx *gorm.DB, postId int64) (data *Post, err error) {
	where := &Post{
		ID: int64(postId),
	}
	data = new(Post)
	//加行锁
	if err = tx.Set("gorm:query_option", " FOR UPDATE ").Where(where).First(data).Error; err != nil {
		if err == gorm.ErrRecordNotFound { //处理了未找到的情况，err = nil, 但是data.ID==0
			err = nil
		}
		return
		err = errors.Wrap(err, "数据库查找错误")
	}
	return
}

func AddNewPostRecord(userId int64, title, content string) (err error) {
	data := &Post{
		UserID: int64(userId),
		Title:  title,
		Detail: content,
		ReplyCount: 1,
	}
	if err = db.Create(data).Error; err != nil {
		return
	}
	return
}

func UpdatePostReplyCount(db *gorm.DB, postId, floor int64) (err error){
	data := &Post{
		ReplyCount: int64(floor),
	}
	if err = db.Table(data.TableName()).Where("id=?", postId).Updates(data).Error; err != nil {
		err = errors.Wrap(err, "更新数据库reply_count失败")
		return
	}
	return
}
