package util

import "github.com/jinzhu/gorm"

func QueryForUpdate(tx *gorm.DB) (lockedTx *gorm.DB) {
	return tx.Set("gorm:query_option", " FOR UPDATE ")
}
