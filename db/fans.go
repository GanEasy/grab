package db

import (
	"time"
)

// Fans 粉丝数据信息
type Fans struct {
	ID         uint   `gorm:"primary_key"`
	OpenID     string `gorm:"type:varchar(255);unique_index"`
	NickName   string
	Gender     int
	City       string
	Province   string
	Country    string
	AvatarURL  string
	Language   string
	Timestamp  int64
	Level      int32 // 用户等级
	Score      int64 // 积分
	Total      int64 // 总分
	AppID      string
	SessionKey string // 粉丝上次的session key 如果有变化，同步一次粉丝数据
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
}

// GetFansByOpenID  通过openID获取粉丝信息如果没有的话进行初始化
func (fans *Fans) GetFansByOpenID(openID string) {
	DB().Where(Fans{OpenID: openID}).FirstOrCreate(&fans)
}

// Save 保存粉丝信息
func (fans *Fans) Save() {
	DB().Save(&fans)
}

// Follow 关注列表
type Follow struct {
	ID        uint   `gorm:"primary_key"`
	OpenID    string `gorm:"type:varchar(255);index"`
	Title     string
	URL       string
	Drive     string
	Page      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Source 用户自定义的数据源
type Source struct {
	ID        uint   `gorm:"primary_key"`
	OpenID    string `gorm:"type:varchar(255);index"`
	Title     string
	URL       string
	Drive     string
	Page      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Log 粉丝浏览记录
type Log struct {
	ID        uint   `gorm:"primary_key"`
	OpenID    string `gorm:"type:varchar(255);index"`
	Title     string
	URL       string
	Drive     string
	Page      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// GetFansAllFollows  获取用户关注的源
func (fans *Fans) GetFansAllFollows(offser, limit int) (follows []Follow) {
	DB().Where(Follow{OpenID: fans.OpenID}).Limit(limit).Offset(offser).Find(&follows)
	return
}

// GetFansAllSources  获取粉丝添加的源
func (fans *Fans) GetFansAllSources(offser, limit int) (sources []Source) {
	DB().Where(Source{OpenID: fans.OpenID}).Limit(limit).Offset(offser).Find(&sources)
	return
}
