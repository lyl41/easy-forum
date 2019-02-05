package mysql

import (
	"time"
)

type Sign struct {
	ID                int64     `sql:"primary_key;column:id" json:"id,omitempty"`                       //
	UserID            int64     `sql:"column:user_id" json:"user_id,omitempty"`                         //
	DateMonth         time.Time `sql:"column:date_month" json:"date_month,omitempty"`                   //月份
	Mask              int64     `sql:"column:mask" json:"mask,omitempty"`                               //
	ContinueSignMonth int64     `sql:"column:continue_sign_month" json:"continue_sign_month,omitempty"` //本月连续登录的天数

}

func (Sign) TableName() string {
	return "sign"
}
