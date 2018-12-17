package db

import (
	"time"
)

// Catalog 目录
type Catalog struct {
	ID        uint   `gorm:"primary_key"`
	OpenID    string `gorm:"type:varchar(255);index"` // 微信文章地址
	Title     string `gorm:"type:varchar(255);"`      //订阅formID，一次订阅只能推送一次通知
	URL       string `gorm:"type:varchar(2048);"`     //订阅formID，一次订阅只能推送一次通知
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Save Feedback
func (feedback *Catalog) Save() {
	DB().Save(&feedback)
}
