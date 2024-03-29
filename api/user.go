package api

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/GanEasy/grab"
	cpi "github.com/GanEasy/grab/core"
	"github.com/GanEasy/grab/db"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// JwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	FansID  uint   `json:"fans_id"`
	OpenID  string `json:"open_id"`
	Code    string `json:"code"`
	Session string `json:"session"`
	jwt.StandardClaims
}

// GetToken 获取 jwt token
func GetToken(c echo.Context) error {
	cf := cpi.GetConf()
	code := c.QueryParam("code")
	version := c.QueryParam("version")

	if true {
		// return GetOpenToken(c)
	}
	fromid, _ := strconv.Atoi(c.QueryParam("fromid"))
	ret, _ := cpi.GetOpenID(code)
	if code != "" && ret.OpenID != "" {
		fans, err := cpi.GetFansByOpenID(ret.OpenID)

		// 增加用户被邀请次数
		if fromid > 0 && uint(fromid) != fans.ID {
			fans.InvitationTotal++
		}

		fans.LoginTotal++ // 增加一次访问登录次数
		fans.Save()
		claims := &JwtCustomClaims{
			fans.ID,
			ret.OpenID,
			code,
			ret.SessionKey,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		var jumpappid = `` // 强制跳
		var bookjumpappid = `` //wx359657b0849ee636 驴友记
		var articlejumpappid = `` //
		var jumpwebpage = ``      //
		var canCreate = 0

		if version != `` && version == cf.Search.DevVersion {
			jumpappid = ``
			bookjumpappid = ``
			articlejumpappid = ``
		} else if fans.LoginTotal < 5 && fans.Level < 3 { // 如果访问次数少于3次，等级小于3，强制跳转到其它小程序阅读(测试下)
			// day := time.Now().Day()
			if false { // 要不要新用户强制跳转
				var juid = int(fans.ID)
				var janum = juid % 2 //不同用户控制不同 转 不同小程序 （分流）
				if janum == 0 {
					if juid > 0 { // 待post新版本后，隔离掉老用户
						jumpappid = `wx359657b0849ee636` // 强制去 
					} 
				} else {
					jumpappid = `wx359657b0849ee636` // 强制去 
				}
			}

		} else if fans.LoginTotal > 5 { // 大于10次，强制跳转
			jumpappid = `` // 强制跳去 wx359657b0849ee636
		}

		// if fans.LoginTotal < 3 {
		// 	jumpappid = ``
		// 	bookjumpappid = ``
		// 	articlejumpappid = ``
		// }

		// if fans.Level > 2 {
		// 	jumpappid = ``
		// 	bookjumpappid = ``
		// 	articlejumpappid = ``
		// 	canCreate = 1
		// }

		// 蜘蛛来的，给采集相关内容
		var req = c.Request()
		if version != cf.Search.DevVersion && strings.Contains(req.Header.Get("User-Agent"), `mpcrawler`) { // 获取通用 token  新推荐阅读
			jumpappid = ``
			jumpwebpage = ``
			bookjumpappid = ``
			articlejumpappid = ``
		}

		var infoTipsBanner, infoTipsCustom string

		if fans.LoginTotal > 0 { // 大于x（随机给广告点击）
			rand.Seed(time.Now().UnixNano())
			inum := rand.Intn(3) // 先搞低些广告出现机率
			// if inum==1 {
			// 	infoTipsBanner =  cf.Ad.InfoBanner
			// }else if inum==2{
			// 	info_tips_grid =  cf.Ad.InfoGrid
			// }

			// day := time.Now().Day()
			// var uid = int(fans.ID)
			// var inum = (day + uid) % 3 //机率控制 2/3 banner
			if inum == 0 { // 日期加uid求余 为0 给banner 为 1 给grid
				infoTipsBanner = cf.Ad.InfoBanner
			} else if inum == 1 {
				// info_tips_grid = cf.Ad.InfoGrid
				// infoTipsCustom = `` // adunit-9bb55eb7ddd541d4
				infoTipsBanner = cf.Ad.InfoBanner
			} else if inum == 2 {
				infoTipsBanner = cf.Ad.InfoBanner
			}
		}
		if jumpappid == `` {
		}

		return c.JSON(http.StatusOK, echo.Map{
			"jumpappid":        jumpappid,        // 强制跳转其它小程序
			"bookjumpappid":    bookjumpappid,    // 文本强制跳转其它小程序
			"articlejumpappid": articlejumpappid, // 文章强制跳转其它小程序
			"jumpwebpage":      jumpwebpage,      // 强制跳转网站阅读
			"jumpwebtips":      `已复制网址，请使用浏览器访问`, // 强制跳转网站阅读
			"report_api":       `https://tongji.readfollow.com/api/report`,
			"token":            t,
			"uid":              fans.ID,
			"level":            0,
			"can_create":       canCreate, // 允许创建内容
			"ismini":           0,
			"hiderec":          0,
			"hidelog":          0,
			// "home_screen_adid": cf.Ad.InfoScreen, // 给个首页插屏试试
			"info_screen":      cf.Ad.InfoScreen, //插屏
			"info_banner":      cf.Ad.InfoBanner,
			"info_tips_banner": infoTipsBanner, // 点击广告开启自动加载更多功能
			"info_tips_custom": infoTipsCustom, // 详细页格子广告
			// "info_tips_banner": cf.Ad.InfoBanner, // 点击广告开启自动加载更多功能
			// "info_tips_custom": cf.Ad.InfoGrid, // 详细页格子广告
			"autoload_tips": `观看广告加载下一章(并关闭弹窗)`,
			// "top_home_video": cf.Ad.TopHomeVideo,
			"top_home_custom": `adunit-04fe1b3d519b9299`,
			// "top_list_video": cf.Ad.HomeVideo,
			// "home_video":     cf.Ad.HomeVideo,
			"list_video":  cf.Ad.ListVideo,
			"cata_video":  cf.Ad.CataVideo,
			"info_video":  cf.Ad.InfoVideo,
			"info_custom": `adunit-107ff5514ca1654a`, // 详细页格子广告
			"info_reward": cf.Ad.Reward,
			// 定义首页分享标题
			"share_title": cf.ReaderMinApp.AppTitle,
			// 定义首页分享图片
			"share_cover":       cf.ReaderMinApp.AppCover,
			"placeholder":       cf.ReaderMinApp.AppSearch, // 小说名
			"online_service":    true,
			"info_force_reward": true, // 强制广告
			"info_video_adlt":   4,    //详情页面视频轮循总数
			"info_video_adlm":   1,    //详情页面视频轮循开始余量
			"info_custom_adlt":  2,    //详情页面格子广告轮循总数
			"info_custom_adlm":  0,    //详情页面格子广告轮循开始余量
			"info_banner_adlt":  4,    //详情页面Banner轮循总数
			"info_banner_adlm":  3,    //详情页面Banner轮循开始余量
			"info_screen_adlt":  5,    //详情页面插屏广告轮循总数
			"info_screen_adlm":  3,    //详情页面插屏广告轮循开始余量

			"advert_txt":`18元红包(每天重复领)`,
			"advert_type":`jumpapp`,
			"advert_appid":`wxde8ac0a21135c07d`,
			"advert_path":`/index/pages/h5/h5?f_userId=1&f_token=1&s_cps=1&weburl=https%3A%2F%2Fdpurl.cn%2FwPpyUGtz`,
		})

	}

	return echo.ErrUnauthorized
}

// GetOpenToken 获取 开放的 token
func GetOpenToken(c echo.Context) error {
	claims := &JwtCustomClaims{
		1,
		`visitor.OpenID`,
		``,
		``,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token":     t,
		"uid":       -1,
		"level":     0,
		"jumpappid": `wx8664d56a896e375b`, // 强制跳转其它小程序
	})
}

//GetAPIToken 获取 jwt token
func GetAPIToken(c echo.Context) error {

	fromid, _ := strconv.Atoi(c.QueryParam("fromid"))
	// 直接给 -1(不经过验证用户openid)
	if false {
		claims := &JwtCustomClaims{
			1,
			`visitor.OpenID`,
			``,
			``,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		cf := cpi.GetConf()

		// rand.Seed(time.Now().UnixNano())
		// inum := rand.Intn(10) // 先搞低些广告出现机率

		// var info_tips_banner,info_tips_grid string
		// if inum==1 {
		// 	info_tips_banner = cf.Ad.InfoBanner
		// }else if inum==2{
		// 	info_tips_grid =  cf.Ad.InfoGrid
		// }
		// var canCreate int32
		// canCreate = 0
		// if fromid > 0 {
		// 	canCreate = 1
		// }

		return c.JSON(http.StatusOK, echo.Map{
			"token":      t,
			"uid":        -1,
			"level":      0,
			"can_create": 1, // 允许创建内容
			// "list_screen": cf.Ad.ListScreen,
			"info_screen": cf.Ad.InfoScreen,
			// "cata_screen": cf.Ad.CataScreen,
			// "screen":      cf.Ad.Screen,
			// "reward":      cf.Ad.Reward,
			// "pre_video":   cf.Ad.PreVideo,

			// "top_home_banner": cf.Ad.TopHomeBanner,
			// "top_list_banner": cf.Ad.HomeBanner,
			// "home_banner":     cf.Ad.HomeBanner,
			// "list_banner": cf.Ad.ListBanner,
			// "cata_banner": cf.Ad.CataBanner,
			"info_banner": cf.Ad.InfoBanner,
			// "info_tips_banner": info_tips_banner, // 点击广告开启自动加载更多功能
			// "info_tips_grid": info_tips_grid, // 详细页格子广告
			// "info_tips_banner": cf.Ad.InfoBanner, // 点击广告开启自动加载更多功能
			// "info_tips_grid": cf.Ad.InfoGrid, // 详细页格子广告
			"autoload_tips": `体验广告6秒开启自动加载无弹窗模式`,

			"top_home_video": cf.Ad.TopHomeVideo,
			// "top_list_video": cf.Ad.HomeVideo,
			// "home_video":     cf.Ad.HomeVideo,
			"list_video": cf.Ad.ListVideo,
			"cata_video": cf.Ad.CataVideo,
			"info_video": cf.Ad.InfoVideo,

			// "top_home_grid": cf.Ad.HomeGrid, // 首页格子广告
			// "top_list_grid": cf.Ad.HomeGrid, // 首页格子广告
			// "home_grid":     cf.Ad.HomeGrid, // 首页格子广告
			// "list_grid": cf.Ad.ListGrid, // 列表页格子广告
			// "cata_grid": cf.Ad.CataGrid, // 列表页格子广告
			// "info_grid": cf.Ad.InfoGrid, // 详细页格子广告
			// "home_pre_video": cf.Ad.PreVideo,
			// "list_pre_video": cf.Ad.PreVideo,
			// "info_pre_video": cf.Ad.PreVideo,

			// "home_reward": cf.Ad.Reward,
			// "list_reward": cf.Ad.Reward,
			"info_reward": cf.Ad.Reward,

			// 定义首页分享标题
			"share_title": cf.ReaderMinApp.AppTitle,
			// 定义首页分享图片
			"share_cover":       cf.ReaderMinApp.AppCover,
			"placeholder":       cf.ReaderMinApp.AppSearch, // 小说名
			"online_service":    true,
			"info_force_reward": true, // 强制广告
			"info_video_adlt":   2,    //详情页面视频轮循总数
			"info_video_adlm":   0,    //详情页面视频轮循开始余量
			// "info_grid_adlt":    2,    //详情页面格子广告轮循总数
			// "info_grid_adlm":    1,    //详情页面格子广告轮循开始余量
			"info_banner_adlt": 2, //详情页面Banner轮循总数
			"info_banner_adlm": 1, //详情页面Banner轮循开始余量
			"info_screen_adlt": 3, //详情页面插屏广告轮循总数
			"info_screen_adlm": 2, //详情页面插屏广告轮循开始余量
		})
	}

	code := c.QueryParam("code")
	provider := c.QueryParam("provider")
	if provider == `weixin` {
		ret, _ := cpi.GetOpenID(code)
		if code != "" && ret.OpenID != "" {
			fans, err := cpi.GetFansByOpenID(ret.OpenID)
			if fans.Provider == `` {
				fans.Provider = provider
			}

			// 增加用户被邀请次数
			if fromid > 0 && uint(fromid) != fans.ID {
				fans.InvitationTotal++
			}

			fans.LoginTotal++ // 增加一次访问登录次数
			fans.Save()
			claims := &JwtCustomClaims{
				fans.ID,
				ret.OpenID,
				code,
				ret.SessionKey,
				jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
				},
			}

			// Create token with claims
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			// Generate encoded token and send it as response.
			t, err := token.SignedString([]byte("secret"))
			if err != nil {
				return err
			}
			cf := cpi.GetConf()

			return c.JSON(http.StatusOK, echo.Map{
				"token": t,
				"uid":   fans.ID,

				"level":      0,
				"can_create": 1, // 允许创建内容
				// "list_screen": cf.Ad.ListScreen,
				"info_screen": cf.Ad.InfoScreen,
				// "cata_screen": cf.Ad.CataScreen,
				// "screen":      cf.Ad.Screen,
				// "reward":      cf.Ad.Reward,
				// "pre_video":   cf.Ad.PreVideo,

				// "top_home_banner": cf.Ad.TopHomeBanner,
				// "top_list_banner": cf.Ad.HomeBanner,
				// "home_banner":     cf.Ad.HomeBanner,
				// "list_banner": cf.Ad.ListBanner,
				// "cata_banner": cf.Ad.CataBanner,
				"info_banner": cf.Ad.InfoBanner,
				// "info_tips_banner": info_tips_banner, // 点击广告开启自动加载更多功能
				// "info_tips_grid": info_tips_grid, // 详细页格子广告
				// "info_tips_banner": cf.Ad.InfoBanner, // 点击广告开启自动加载更多功能
				// "info_tips_grid": cf.Ad.InfoGrid, // 详细页格子广告
				"autoload_tips": `体验广告6秒开启自动加载无弹窗模式`,

				"top_home_video": cf.Ad.TopHomeVideo,
				// "top_list_video": cf.Ad.HomeVideo,
				// "home_video":     cf.Ad.HomeVideo,
				"list_video": cf.Ad.ListVideo,
				"cata_video": cf.Ad.CataVideo,
				"info_video": cf.Ad.InfoVideo,

				// "top_home_grid": cf.Ad.HomeGrid, // 首页格子广告
				// "top_list_grid": cf.Ad.HomeGrid, // 首页格子广告
				// "home_grid":     cf.Ad.HomeGrid, // 首页格子广告
				// "list_grid": cf.Ad.ListGrid, // 列表页格子广告
				// "cata_grid": cf.Ad.CataGrid, // 列表页格子广告
				// "info_grid": cf.Ad.InfoGrid, // 详细页格子广告
				// "home_pre_video": cf.Ad.PreVideo,
				// "list_pre_video": cf.Ad.PreVideo,
				// "info_pre_video": cf.Ad.PreVideo,

				// "home_reward": cf.Ad.Reward,
				// "list_reward": cf.Ad.Reward,
				"info_reward": cf.Ad.Reward,

				// 定义首页分享标题
				"share_title": cf.ReaderMinApp.AppTitle,
				// 定义首页分享图片
				"share_cover":       cf.ReaderMinApp.AppCover,
				"placeholder":       cf.ReaderMinApp.AppSearch, // 小说名
				"online_service":    true,
				"info_force_reward": true, // 强制广告
				"info_video_adlt":   2,    //详情页面视频轮循总数
				"info_video_adlm":   0,    //详情页面视频轮循开始余量
				// "info_grid_adlt":    2,    //详情页面格子广告轮循总数
				// "info_grid_adlm":    1,    //详情页面格子广告轮循开始余量
				"info_banner_adlt": 2, //详情页面Banner轮循总数
				"info_banner_adlm": 1, //详情页面Banner轮循开始余量
				"info_screen_adlt": 3, //详情页面插屏广告轮循总数
				"info_screen_adlm": 2, //详情页面插屏广告轮循开始余量
			})

		}
	} else if provider == `h5` {

		claims := &JwtCustomClaims{
			1,
			`visitor.OpenID`,
			``,
			``,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
			"uid":   -1,
			"level": 0,
		})
	} else if provider == `qq` {

		claims := &JwtCustomClaims{
			1,
			`visitor.OpenID`,
			``,
			``,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
			"uid":   -1,
			"level": 0,
		})
	}

	// 如果上面没通过，给个 -1吧
	if true {
		claims := &JwtCustomClaims{
			1,
			`visitor.OpenID`,
			``,
			``,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		cf := cpi.GetConf()
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
			"uid":   -1,
			"level": 0,
			// "list_screen": cf.Ad.ListScreen,
			// "info_screen": cf.Ad.InfoScreen,
			// "cata_screen": cf.Ad.CataScreen,
			// "screen":      cf.Ad.Screen,
			// "reward":      cf.Ad.Reward,
			// "pre_video":   cf.Ad.PreVideo,

			// "top_home_banner": cf.Ad.TopHomeBanner,
			// "top_list_banner": cf.Ad.HomeBanner,
			// "home_banner":     cf.Ad.HomeBanner,
			// "list_banner": cf.Ad.ListBanner,
			// "cata_banner": cf.Ad.CataBanner,
			"info_banner": cf.Ad.InfoBanner,

			"top_home_video": cf.Ad.TopHomeVideo,
			// "top_list_video": cf.Ad.HomeVideo,
			// "home_video":     cf.Ad.HomeVideo,
			"list_video": cf.Ad.ListVideo,
			"cata_video": cf.Ad.CataVideo,
			"info_video": cf.Ad.InfoVideo,

			// "top_home_grid": cf.Ad.HomeGrid, // 首页格子广告
			// "top_list_grid": cf.Ad.HomeGrid, // 首页格子广告
			// "home_grid":     cf.Ad.HomeGrid, // 首页格子广告
			// "list_grid": cf.Ad.ListGrid, // 列表页格子广告
			// "cata_grid": cf.Ad.CataGrid, // 列表页格子广告
			"info_grid": cf.Ad.InfoGrid, // 详细页格子广告
			// "home_pre_video": cf.Ad.PreVideo,
			// "list_pre_video": cf.Ad.PreVideo,
			// "info_pre_video": cf.Ad.PreVideo,

			// "home_reward": cf.Ad.Reward,
			// "list_reward": cf.Ad.Reward,
			"info_reward": cf.Ad.Reward,

			// 定义首页分享标题
			"share_title": cf.ReaderMinApp.AppTitle,
			// 定义首页分享图片
			"share_cover":       cf.ReaderMinApp.AppCover,
			"placeholder":       cf.ReaderMinApp.AppSearch, // 小说名
			"online_service":    true,
			"info_force_reward": cf.Ad.ForceReward,    // 看小说下一章强制要点视频广告
			"info_video_adlt":   cf.Ad.InfoVideoAdlt,  //详情页面视频轮循总数
			"info_video_adlm":   cf.Ad.InfoVideoAdlm,  //详情页面视频轮循开始余量
			"info_banner_adlt":  cf.Ad.InfoBannerAdlt, //详情页面Banner轮循总数
			"info_banner_adlm":  cf.Ad.InfoBannerAdlm, //详情页面Banner轮循开始余量
			"info_grid_adlt":    cf.Ad.InfoGridAdlt,   //详情页面格子广告轮循总数
			"info_grid_adlm":    cf.Ad.InfoGridAdlm,   //详情页面格子广告轮循开始余量
			// "info_screen_adlt":  cf.Ad.InfoScreenAdlt, //详情页面插屏广告轮循总数
			// "info_screen_adlm":  cf.Ad.InfoScreenAdlm, //详情页面插屏广告轮循开始余量
		})
	}
	return echo.ErrUnauthorized
}

//GetAPIToken2 获取 jwt token
func GetAPIToken2(c echo.Context) error {

	// fromid, _ := strconv.Atoi(c.QueryParam("fromid"))
	// 直接给 -1(不经过验证用户openid)
	if true {
		claims := &JwtCustomClaims{
			1,
			`visitor.OpenID`,
			``,
			``,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		cf := cpi.GetConf()
		return c.JSON(http.StatusOK, echo.Map{
			"token":          t,
			"uid":            -1,
			"level":          0,
			"ismini":         1,
			"show_tips_next": 0,
			"can_create":     0, // 允许创建内容

			// 定义首页分享标题
			"share_title": cf.ReaderMinApp.AppTitle,
			// 定义首页分享图片
			"share_cover":       cf.ReaderMinApp.AppCover,
			"placeholder":       cf.ReaderMinApp.AppSearch, // 小说名
			"online_service":    true,
			"info_force_reward": cf.Ad.ForceReward,    // 看小说下一章强制要点视频广告
			"info_video_adlt":   cf.Ad.InfoVideoAdlt,  //详情页面视频轮循总数
			"info_video_adlm":   cf.Ad.InfoVideoAdlm,  //详情页面视频轮循开始余量
			"info_banner_adlt":  cf.Ad.InfoBannerAdlt, //详情页面Banner轮循总数
			"info_banner_adlm":  cf.Ad.InfoBannerAdlm, //详情页面Banner轮循开始余量
			"info_grid_adlt":    cf.Ad.InfoGridAdlt,   //详情页面格子广告轮循总数
			"info_grid_adlm":    cf.Ad.InfoGridAdlm,   //详情页面格子广告轮循开始余量
			"info_screen_adlt":  cf.Ad.InfoScreenAdlt, //详情页面插屏广告轮循总数
			"info_screen_adlm":  cf.Ad.InfoScreenAdlm, //详情页面插屏广告轮循开始余量
		})
	}

	return echo.ErrUnauthorized
}

//CheckOpenID 获取签名里面的信息
func CheckOpenID(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return c.JSON(http.StatusOK, claims.OpenID)

}

// 获取签名里面的信息
func getOpenID(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims.OpenID
}

// 获取最后一个用户信息
func GetLastUser(c echo.Context) error {
	var fans db.Fans
	fans.GetLastUser()
	return c.JSON(http.StatusOK, fans)
}

// 获取用户信息
func getUser(openID string) (*db.Fans, error) {
	fans, err := cpi.GetFansByOpenID(openID)
	return fans, err
}

// Crypt 解密同步用户信息
func Crypt(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	// sessionKey := c.QueryParam("sk")
	encryptedData := c.QueryParam("ed")
	iv := c.QueryParam("iv")
	ret, _ := cpi.GetCryptData(claims.Session, encryptedData, iv)
	return c.JSON(http.StatusOK, ret)
}

// GetQrcode 获取二维码地址
func GetQrcode(c echo.Context) error {
	// 去哪里
	wxto := c.QueryParam("wxto")
	log.Println(`wxto`, wxto)
	if wxto == `` {
		return c.HTML(http.StatusOK, "wxto empty")
	}
	var qrcode = db.Qrcode{}
	qrcode.GetQrcodeID(wxto)

	page := `pages/index`
	fileName, err := cpi.GetwxCodeUnlimit(strconv.Itoa(int(qrcode.ID)), page)

	log.Println(`fileName`, fileName)
	if err == nil {
		http.ServeFile(c.Response().Writer, c.Request(), fileName)
	} else {
		return err
	}

	var err2 error
	return err2
}

// GetQrcodeWxto 二维码ID去哪里
func GetQrcodeWxto(c echo.Context) error {
	// 去哪里
	qrid := c.QueryParam("qrid")
	if qrid == `` {
		return c.HTML(http.StatusOK, `pages/index`)
	}
	var qrcode = db.Qrcode{}

	id, e := strconv.Atoi(qrid)
	if e != nil {
		return c.HTML(http.StatusOK, `pages/index`)
	}
	qrcode.GetQrcodeByID(id)

	if qrcode.WxTo == `` {
		return c.HTML(http.StatusOK, `pages/index`)
	}

	log.Println(`qrcode to `, qrcode.WxTo)
	return c.HTML(http.StatusOK, qrcode.WxTo)
}

// GetUserFollows 获取用户关注的
func GetUserFollows(c echo.Context) error {
	//
	openID := getOpenID(c)
	if openID == `` {
		return c.HTML(http.StatusOK, "wxto empty")
	}
	user, _ := getUser(openID)
	sources := user.GetFansAllSources(0, 100)

	var cards []grab.Card
	wxto := ``
	for _, l := range sources {
		wxto = fmt.Sprintf(`/pages/catalog/get?drive=%v&url=%v`, l.Drive, grab.EncodeURL(l.URL))
		cards = append(cards,
			grab.Card{
				l.Title,
				wxto, ``,
				`link`,
				``,
				nil,
			},
		)
	}

	return c.JSON(http.StatusOK, cards)
}

// CreateUserFollow 获取用户关注的
func CreateUserFollow(c echo.Context) error {
	//
	openID := getOpenID(c)
	if openID == `` {
		return c.HTML(http.StatusOK, "wxto empty")
	}
	user, _ := getUser(openID)
	sources := user.GetFansAllSources(0, 100)

	var cards []grab.Card
	wxto := ``
	for _, l := range sources {
		wxto = fmt.Sprintf(`/pages/catalog/get?drive=%v&url=%v`, l.Drive, grab.EncodeURL(l.URL))
		cards = append(cards,
			grab.Card{
				l.Title,
				wxto, ``,
				`link`,
				``,
				nil,
			},
		)
	}

	return c.JSON(http.StatusOK, cards)
}

// GetUserSources 获取用户自己添加的书籍源
func GetUserSources(c echo.Context) error {
	//
	openID := getOpenID(c)
	if openID == `` {
		return c.HTML(http.StatusOK, "wxto empty")
	}
	user, _ := getUser(openID)
	sources := user.GetFansAllSources(0, 100)

	var cards []grab.Card
	wxto := ``
	for _, l := range sources {
		wxto = fmt.Sprintf(`/pages/catalog/get?drive=%v&url=%v`, l.Drive, grab.EncodeURL(l.URL))
		cards = append(cards,
			grab.Card{
				l.Title,
				wxto, ``,
				`link`,
				``,
				nil,
			},
		)
	}

	return c.JSON(http.StatusOK, cards)
}

// CreateUserSource 添加源
func CreateUserSource(c echo.Context) error {
	//
	openID := getOpenID(c)
	if openID == `` {
		return c.HTML(http.StatusOK, "wxto empty")
	}
	user, _ := getUser(openID)
	sources := user.GetFansAllSources(0, 100)

	var cards []grab.Card
	wxto := ``
	for _, l := range sources {
		wxto = fmt.Sprintf(`/pages/catalog/get?drive=%v&url=%v`, l.Drive, grab.EncodeURL(l.URL))
		cards = append(cards,
			grab.Card{
				l.Title,
				wxto, ``,
				`link`,
				``,
				nil,
			},
		)
	}

	return c.JSON(http.StatusOK, cards)
}
