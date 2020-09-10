package core

import "github.com/BurntSushi/toml"

// Config 配置
type Config struct {
	ReaderMinApp ReaderMinApp
	ReaderMinAppTwo ReaderMinAppTwo
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
	JumpAppID string `toml:"jump_app_id"`
}

//ReaderMinAppTwo 配置
type ReaderMinAppTwo struct {
	AppID     string `toml:"app_id"`
	AppSecret string `toml:"app_secret"`
	AppTitle  string `toml:"app_title"`
	AppCover  string `toml:"app_cover"`
	AppSearch string `toml:"app_search"`
	JumpAppID string `toml:"jump_app_id"`
}

//ReaderMinAppThree 配置
type ReaderMinAppThree struct {
	AppID     string `toml:"app_id"`
	AppSecret string `toml:"app_secret"`
	AppTitle  string `toml:"app_title"`
	AppCover  string `toml:"app_cover"`
	AppSearch string `toml:"app_search"`
	JumpAppID string `toml:"jump_app_id"`
}

//Search 配置
type Search struct {
	LimitLevel       bool   `toml:"limit_level"`
	LimiterStageShow bool   `toml:"limiter_stage_show"`
	DevVersion       string `toml:"dev_version"`
	LimitInvitation  bool   `toml:"limit_invitation"` // 必须受邀请用户才能搜索和查看目录
	InvitationCode   string `toml:"invitation_code"`  // 邀请密令，输入密令解锁邀请
	InvitationNo     int    `toml:"invitation_no"`    // 邀请编号，从分享页面进来，必须为指定编号
}

//Ad 配置
type Ad struct {
	ForceReward    bool   `toml:"force_reward"`     // 强制激励广告看下一章
	Screen         string `toml:"screen"`           // 弹窗视频广告
	Reward         string `toml:"reward"`           // 激励式视频广告
	TopHomeBanner  string `toml:"top_home_banner"`  // 首页顶部banner广告(与视频广告两者用其一)
	HomeBanner     string `toml:"home_banner"`      // 首页banner广告(与视频广告两者用其一)
	TopListBanner  string `toml:"top_list_banner"`  // 列表页顶部banner广告
	ListBanner     string `toml:"list_banner"`      // 列表页banner广告
	InfoBanner     string `toml:"info_banner"`      // 详情页banner广告
	CataBanner     string `toml:"cata_banner"`      // 文章目录banner广告
	HomeScreen     string `toml:"home_screen"`      // 首页弹窗广告
	ListScreen     string `toml:"list_screen"`      // 列表页弹窗广告
	InfoScreen     string `toml:"info_screen"`      // 详情页弹窗广告
	CataScreen     string `toml:"cata_screen"`      // 文章目录弹窗广告
	TopHomeVideo   string `toml:"top_home_video"`   // 首页顶部视频广告
	HomeVideo      string `toml:"home_video"`       // 首页视频广告
	TopListVideo   string `toml:"top_list_video"`   // 列表页顶部视频广告
	ListVideo      string `toml:"list_video"`       // 列表页视频广告
	InfoVideo      string `toml:"info_video"`       // 详情页视频广告
	CataVideo      string `toml:"cata_video"`       // 文章目录视频广告
	PreVideo       string `toml:"pre_video"`        // 视频播放前广告
	TopHomeGrid    string `toml:"top_home_grid"`    // 首页顶部格子广告
	HomeGrid       string `toml:"home_grid"`        // 首页格子广告
	TopListGrid    string `toml:"top_list_grid"`    // 列表页顶部格子广告
	ListGrid       string `toml:"list_grid"`        // 列表页格子广告
	InfoGrid       string `toml:"info_grid"`        // 详细页格子广告
	CataGrid       string `toml:"cata_grid"`        // 文章目录格子广告
	InfoVideoAdlt  int32  `toml:"info_video_adlt"`  //详情页面视频轮循总数
	InfoVideoAdlm  int32  `toml:"info_video_adlm"`  //详情页面视频轮循开始余量
	InfoBannerAdlt int32  `toml:"info_banner_adlt"` //详情页面Banner轮循总数
	InfoBannerAdlm int32  `toml:"info_banner_adlm"` //详情页面Banner轮循开始余量
	InfoGridAdlt   int32  `toml:"info_grid_adlt"`   //详情页面格子广告轮循总数
	InfoGridAdlm   int32  `toml:"info_grid_adlm"`   //详情页面格子广告轮循开始余量
	InfoScreenAdlt int32  `toml:"info_screen_adlt"` //详情页面插屏广告轮循总数
	InfoScreenAdlm int32  `toml:"info_screen_adlm"` //详情页面插屏广告轮循开始余量

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
