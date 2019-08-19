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
	LimitLevel bool `toml:"limit_level"`
}

//Ad 配置
type Ad struct {
	Screen     string `toml:"screen"`
	Reward     string `toml:"reward"`
	ListBanner string `toml:"list_banner"`
	InfoBanner string `toml:"info_banner"`
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
