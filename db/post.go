package db

import (
	"time"
)

// Post 小程序二维码
type Post struct {
	ID        uint   `gorm:"primary_key"`
	Cate      int32  // 类型  1文章  2小说
	Name      string `gorm:"type:varchar(255);index"` // 资源名称	名称搜索
	WxTo      string `gorm:"type:varchar(255);index"` // 目标地址 	小程序往哪里走
	From      string // 数据来源
	Total     int64  // 次数统计
	Level     int32  // 级别限制(不加入搜索条件限制)
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Save Feedback
func (post *Post) Save() {
	DB().Save(&post)
}

//GetPostByWxto  想去哪里 创建一个记录
func (post *Post) GetPostByWxto(wxto string) {
	DB().Where(Post{WxTo: wxto}).FirstOrCreate(&post)
}

// GetPostByID 获取一个资源
func (post *Post) GetPostByID(id int) {
	DB().First(&post, id)
}

//GetPostsByName  通过名字获得查询记录
func (post *Post) GetPostsByName(name string) (posts []Post) {
	DB().Where("name LIKE %?%", name).Find(&posts)
	// DB().Where("name = ?", name).Find(&posts)
	return
}
