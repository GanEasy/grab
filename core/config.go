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
}

//Search 配置
type Search struct {
	LimitLevel bool   `toml:"limit_level"`
	DevVersion string `toml:"dev_version"`
}

//Ad 配置
type Ad struct {
	Screen     string `toml:"screen"`      // 弹窗视频广告
	Reward     string `toml:"reward"`      // 激励式视频广告
	HomeBanner string `toml:"home_banner"` // 首页banner广告(与视频广告两者用其一)
	ListBanner string `toml:"list_banner"` // 列表页banner广告
	InfoBanner string `toml:"info_banner"` // 详情页banner广告
	HomeVideo  string `toml:"home_video"`  // 首页视频广告
	ListVideo  string `toml:"list_video"`  // 列表页视频广告
	InfoVideo  string `toml:"info_video"`  // 详情页视频广告
	PreVideo   string `toml:"pre_video"`   // 视频播放前广告
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
