package core

import "github.com/BurntSushi/toml"

// Config 配置
type Config struct {
	ReaderMinApp ReaderMinApp
	Database     Database
	Search       Search
	Ad           Ad
}

//ReaderMinApp 配置
type ReaderMinApp struct {
	AppID     string `toml:"app_id"`
	AppSecret string `toml:"app_secret"`
	AppTitle  string `toml:"app_title"`
	AppCover  string `toml:"app_cover"`
	AppSearch string `toml:"app_search"`
}

//Search 配置
type Search struct {
	LimitLevel       bool   `toml:"limit_level"`
	LimiterStageShow bool   `toml:"limiter_stage_show"`
	DevVersion       string `toml:"dev_version"`
}

//Ad 配置
type Ad struct {
	ForceReward bool   `toml:"force_reward"` // 强制激励广告看下一章
	Screen      string `toml:"screen"`       // 弹窗视频广告
	Reward      string `toml:"reward"`       // 激励式视频广告
	HomeBanner  string `toml:"home_banner"`  // 首页banner广告(与视频广告两者用其一)
	ListBanner  string `toml:"list_banner"`  // 列表页banner广告
	InfoBanner  string `toml:"info_banner"`  // 详情页banner广告
	CataBanner  string `toml:"cata_banner"`  // 文章目录banner广告
	HomeScreen  string `toml:"home_screen"`  // 首页弹窗广告
	ListScreen  string `toml:"list_screen"`  // 列表页弹窗广告
	InfoScreen  string `toml:"info_screen"`  // 详情页弹窗广告
	CataScreen  string `toml:"cata_screen"`  // 文章目录弹窗广告
	HomeVideo   string `toml:"home_video"`   // 首页视频广告
	ListVideo   string `toml:"list_video"`   // 列表页视频广告
	InfoVideo   string `toml:"info_video"`   // 详情页视频广告
	CataVideo   string `toml:"cata_video"`   // 文章目录视频广告
	PreVideo    string `toml:"pre_video"`    // 视频播放前广告
	HomeGrid    string `toml:"home_grid"`    // 首页格子广告
	ListGrid    string `toml:"list_grid"`    // 列表页格子广告
	InfoGrid    string `toml:"info_grid"`    // 详细页格子广告
	CataGrid    string `toml:"cata_grid"`    // 文章目录格子广告
}

//Database 配置
type Database struct {
	Type     string `toml:"type"`
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Sslmode  string `toml:"sslmode"`
	Dbname   string `toml:"dbname"`
	Port     int    `toml:"port"`
}

var confFile = "conf.toml"
var config Config

func init() {
	GetConf()
	// DB().AutoMigrate(&Subscribe{}, &Fans{}, &Post{}, &Feedback{})
}

//GetConf 获取config
func GetConf() Config {
	if config.ReaderMinApp.AppID == "" {
		if _, err := toml.DecodeFile(confFile, &config); err != nil {
			panic(err)
		}
	}
	return config
}
