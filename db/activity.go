package db

import (
	"time"
)

// Activity 号召阅读
type Activity struct {
	ID        uint   `gorm:"primary_key"`
	Title     string `gorm:"type:varchar(255)"`       // 资源名称	名称搜索
	WxTo      string `gorm:"type:varchar(255);index"` // 目标地址 	小程序往哪里走
	Total     int64  // 次数统计
	Level     int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Save Feedback
func (activity *Activity) Save() {
	DB().Save(&activity)
}

//GetActivityByWxto  想去哪里 创建一个记录
func (activity *Activity) GetActivityByWxto(wxto string) {
	DB().Where(Activity{WxTo: wxto}).FirstOrCreate(&activity)
}

//GetActivityByID 获取一个资源
func (activity *Activity) GetActivityByID(id int) {
	DB().First(&activity, id)
}

//GetActivities  获取最近100条号召令
func (activity *Activity) GetActivities() (activities []Activity) {
	DB().Order(`updated_at desc`).Limit(200).Find(&activities)
	return
}
