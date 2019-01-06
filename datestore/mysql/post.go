package mysql

import (
	"github.com/jinzhu/gorm"
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

func FindPostByUserIdAndTitle(userId int, title string) (data *Post, err error) {
	paramsNeed := []string{"id"}
	where := &Post{
		UserID: int64(userId),
		Title:  title,
	}
	data = new(Post)
	if err = db.Select(paramsNeed).Where(where).First(data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return
	}
	return
}

func AddNewPost(userId int, title, content string) (err error) {
	data := &Post{
		UserID: int64(userId),
		Title:  title,
		Detail: content,
	}
	if err = db.Create(data).Error; err != nil {
		return
	}
	return
}
