package mysql

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

//通过如下命令自动生成:
///var/folders/bm/5pxrd6855nj81734_7dx3jn80000gn/T/go-build020277320/b001/exe/main -t reply_post -db root:1234567890@tcp(localhost:3306)/forum

type ReplyPost struct {
	ID        int64     `sql:"primary_key;column:id" json:"id,omitempty"`     //
	PostID    int64     `sql:"column:post_id" json:"post_id,omitempty"`       //所属帖子
	To        int64     `sql:"column:to" json:"to,omitempty"`                 //回复给谁，userid
	UserID    int64     `sql:"column:user_id" json:"user_id,omitempty"`       //作者
	Floor     int64     `sql:"column:floor" json:"floor,omitempty"`           //回复所在楼层
	Detail    string    `sql:"column:detail" json:"detail,omitempty"`         //具体回帖内容
	CreatedAt time.Time `sql:"column:created_at" json:"created_at,omitempty"` //
	UpdatedAt time.Time `sql:"column:updated_at" json:"updated_at,omitempty"` //

}

func (ReplyPost) TableName() string {
	return "reply_post"
}

func CreateReplyPostRecord(db *gorm.DB, postId, floor, postUserId, replyUserId int64, content string) (err error) {
	where := &ReplyPost{ //postid和floor应该是联合唯一索引 TODO
		PostID: postId,
		Floor:  floor,
	}
	data := &ReplyPost{
		PostID: int64(postId),
		To:     int64(postUserId),
		UserID: int64(replyUserId),
		Floor:  int64(floor),
		Detail: content,
	}
	err = db.Set("gorm:query_option", " FOR UPDATE ").Where(where).First(data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		err = errors.Wrap(err, "postid and floor数据库查找错误")
		return
	}
	if err == nil {
		err = errors.Wrap(errors.New("数据库存在相同数据"), "发表评论失败，请您重试")
		return
	}
	err = db.Create(data).Error
	if err != nil {
		err = errors.Wrap(err, "发表评论失败，请稍后再试")
		return
	}
	return
}
