package db

import (
	"time"
)

//Sign 签到记录
type Sign struct {
	ID        uint `gorm:"primary_key"`
	FansID    uint `gorm:"index:user_id"` // 谁在收集助力
	CreatedAt time.Time
}

//Fail 签到记录
type Fail struct {
	ID        uint `gorm:"primary_key"`
	FansID    uint `gorm:"index:user_id"` // 谁在收集助力
	CreatedAt time.Time
}
