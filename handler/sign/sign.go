package sign

import (
	"easy-forum/common/util"
	"easy-forum/datestore/mysql"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"time"
)

func DealSign(userId int64) (err error) {
	tx := mysql.GetDB().Begin()
	if err = tx.Error; err != nil {
		err = errors.Wrap(err, "获取数据库事务失败")
		return
	}
	ok := false
	defer func() {
		if !ok {
			fmt.Println("database rollback")
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	d := time.Now()
	//有记录则更新，无记录则创建
	data := new(mysql.Sign)
	where := &mysql.Sign{
		UserID:    userId,
		DateMonth: util.GetFirstDateOfMonth(d), //每月第一天表示这个月份
	}
	if err = util.QueryForUpdate(tx).Select("mask").Where(where).First(data).Error; err != nil && err != gorm.ErrRecordNotFound {
		err = errors.Wrap(err, "查询签到信息失败")
		return
	}
	// err == nil || err == gorm.ErrRecordNotFound
	if err == nil { //说明找到了数据，本月已经有签到信息，需要更新
		data.Mask, data.ContinueSignMonth, err = AddMask(data.Mask, d) //添加签到信息
		if err != nil {
			return
		}
		//这里需要更新mask
		err = tx.Model(data).Where(where).Updates(data).Error
		if err != nil {
			err = errors.Wrap(err, "签到失败，请稍后再试")
			return
		}
	} else {
		err = nil
		//需要创建信息
		data.Mask, data.ContinueSignMonth, err = AddMask(0, d)
		if err != nil {
			err = errors.Wrap(err, "服务器内部错误err578157")
			return
		}
		data.UserID = userId
		data.DateMonth = util.GetFirstDateOfMonth(d)
		if err = tx.Create(data).Error; err != nil {
			err = errors.Wrap(err, "本月新签到失败，请稍后再试")
			return
		}
	}

	ok = true //这个容易漏，如果中途return
	return
}

func AddMask(mask int64, d time.Time) (newMask, continueDays int64, err error) {
	day := d.Day() - 1 //减1是为了合理利用mask的比特位
	if mask&(1<<uint32(day)) > 0 {
		err = errors.Wrap(errors.New("用户已经签到"), "您今天已经签到过啦")
		return
	}
	newMask = mask | (1 << uint32(day))
	continueDays = getContinueDays(mask, day)
	return
}

func getContinueDays(mask int64, tail int) (continueDays int64) {
	for i := tail; i >= 0; i-- { //TODO 这是通用的写法，补签也可以复用这个。优化：根据昨天的签到状态来更新continueDays，但是与补签要分开。
		if mask&(1<<uint32(i)) > 0 {
			continueDays++
		} else {
			break
		}
	}
	return
}

func QueryMask(mask int64, d time.Time) (SignStatus bool) {
	day := d.Day() - 1
	return mask&(1<<uint32(day)) > 0
}

func DealGetSignStatus(userId int64) (SignStatus bool, continueDays int64, err error) {
	d := time.Now()
	data := new(mysql.Sign)
	where := &mysql.Sign{
		UserID:    userId,
		DateMonth: util.GetFirstDateOfMonth(d), //每月第一天表示这个月份
	}
	if err = mysql.GetDB().Select("mask, continue_sign_month").Where(where).First(data).Error; err != nil {
		err = errors.Wrap(err, "查询签到信息失败")
		return
	}
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	SignStatus = QueryMask(data.Mask, d)
	continueDays = getContinueDays(data.Mask, d.Day()-1) //TODO 数据库中continue字段目测无用。
	return
}
