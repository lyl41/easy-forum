package mysql

import (
	"time"

	"github.com/pkg/errors"
)

//通过如下命令自动生成:
///var/folders/bm/5pxrd6855nj81734_7dx3jn80000gn/T/go-build273618223/b001/exe/main -t users -db root:1234567890@tcp(localhost:3306)/forum

type Users struct {
	ID          int64     `sql:"primary_key;column:id" json:"id,omitempty"`         //
	Name        string    `sql:"column:name" json:"name,omitempty"`                 //昵称
	Avatar      string    `sql:"column:avatar" json:"avatar,omitempty"`             //头像
	Level       int64     `sql:"column:level" json:"level,omitempty"`               //用户经验等级
	PostCount   int64     `sql:"column:post_count" json:"post_count,omitempty"`     //发帖数量
	Phone       string    `sql:"primary_key;column:phone" json:"phone,omitempty"`   //手机号
	Password    string    `sql:"column:password" json:"password,omitempty"`         //密码 md5加密
	PostLike    int64     `sql:"column:post_like" json:"post_like,omitempty"`       //所有帖子总集赞数
	Status      int64     `sql:"column:status" json:"status,omitempty"`             //用户状态，0表明正常，1表明人为注销或者拉黑
	Email       string    `sql:"column:email" json:"email,omitempty"`               //邮箱
	EmailVerify int64     `sql:"column:email_verify" json:"email_verify,omitempty"` //邮箱是否认证，0表示没有，1表示认证通过
	CreatedAt   time.Time `sql:"column:created_at" json:"created_at,omitempty"`     //
	UpdatedAt   time.Time `sql:"column:updated_at" json:"updated_at,omitempty"`     //

}

func (Users) TableName() string {
	return "users"
}

func FirstOrCreateUserByPhone(phone string, data *Users) (err error) {
	err = db.Where("phone=?", phone).FirstOrCreate(data).Error
	if err != nil {
		err = errors.Wrap(err, "mysql FirstOrCreate user fail.")
		return
	}
	return
}
