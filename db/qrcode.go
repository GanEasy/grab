package db

import (
	"time"
)

// Qrcode 小程序二维码
type Qrcode struct {
	ID        uint   `gorm:"primary_key"`
	WxTo      string `gorm:"type:varchar(255);index"` // 目标地址
	Total     int64  //次数统计
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Save Feedback
func (qrcode *Qrcode) Save() {
	DB().Save(&qrcode)
}

// GetQrcodeID 生成一个二维码ID
func (qrcode *Qrcode) GetQrcodeID(wxto string) {
	DB().Where(Qrcode{WxTo: wxto}).FirstOrCreate(&qrcode)
}

// GetQrcodeByID 获取一个二维码
func (qrcode *Qrcode) GetQrcodeByID(id int) {
	DB().First(&qrcode, id)
}
