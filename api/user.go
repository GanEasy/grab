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

// GetToken 获取 jwt token 2020新版接口 笔趣阁Pro最最最稳定通道
func GetToken(c echo.Context) error {

	cf := cpi.GetConf()
	var req = c.Request()

	version := c.QueryParam("version")
	if version == cf.Search.DevVersion { // 开启严格检查
		return GetCheckModeToken(c)
	}
	if cf.Search.LimitLevel { // 开启安全检查
		return GetSafeToken(c)
	}

	if strings.Contains(req.Referer(), cf.ReaderMinApp.AppID) { // 获取通用 token  Pro
		return GetAPIToken8(c)
	}
	if strings.Contains(req.Referer(), `wx8ffa5a58c0bb3589`) { // 获取通用 token  新推荐阅读
		return GetAPIToken7(c)
	}
	if strings.Contains(req.Referer(), `wx331f3c3e2761f080`) { // 获取 token plus版
		return GetAPIToken4(c)
	}
	if strings.Contains(req.Referer(), `wx8664d56a896e375b`) { // 获取通用 token 免版本图
		return GetAPIToken6(c)
	}
	if strings.Contains(req.Referer(), `wxe70eee58e64c7ac7`) { // 获取通用 token 搜书大师
		return GetAPIToken2(c)
	}
	if strings.Contains(req.Referer(), `wx359657b0849ee636`) { // 获取通用 token 驴友记
		return GetAPIToken9(c)
	}
	if strings.Contains(req.Referer(), cf.ReaderMinAppFour.AppID) { // 获取 token 笔趣阁在线 wx7c30b98c7f42f651
		return GetAPIToken3(c)
	}
	if !strings.Contains(req.Referer(), cf.ReaderMinApp.AppID) { // 获取通用 token 非笔趣阁Pro
		return GetOpenToken(c)
	}
	return GetOpenToken(c)

}

// GetOpenToken 获取对外开的接口
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
	cf := cpi.GetConf()

	version := c.QueryParam("version")
	var ismini = 0
	if cf.Search.LimitLevel || version == cf.Search.DevVersion { // 开启严格检查
		ismini = 1
	}

	return c.JSON(http.StatusOK, echo.Map{
		"jumpappid":  cf.ReaderMinApp.JumpAppID, // 强制跳转其它小程序
		"token":      t,
		"uid":        -1,
		"level":      0,
		"can_create": 1, // 允许创建内容
		"ismini":     ismini,
		// 定义首页分享标题
		"share_title": cf.ReaderMinApp.AppTitle,
		// 定义首页分享图片
		"share_cover":    cf.ReaderMinApp.AppCover,
		"placeholder":    cf.ReaderMinApp.AppSearch, // 小说名
		"online_service": true,
	})
}

// GetCheckModeToken 获取审核模式的token
func GetCheckModeToken(c echo.Context) error {

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
		"jumpappid":   ``,               // 强制跳转其它小程序
		"jumpwebpage": ``,               //
		"jumpwebtips": `已复制网址，请使用浏览器访问`, //
		"token":       t,
		"uid":         -1,
		"level":       0,
		"can_create":  0, // 允许创建内容
		"ismini":           0,
		"hiderec":          1,
		"hidelog":          1,
		// 定义首页分享标题
		"share_title": ``,
		// 定义首页分享图片
		"share_cover":     ``,
		"placeholder":     `请输入关键字搜索`, // 小说名
		"online_service":  true,
		"top_home_custom": `adunit-44122f4a8ef3d7d0`,
	})
}

// GetSafeToken 获取安全的token
func GetSafeToken(c echo.Context) error {

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
		"jumpappid":   ``,               // 强制跳转其它小程序
		"jumpwebpage": ``,               //
		"jumpwebtips": `已复制网址，请使用浏览器访问`, //
		"token":       t,
		"uid":         -1,
		"level":       0,
		"can_create":  0, // 允许创建内容
		"ismini":      0,
		// 定义首页分享标题
		"share_title": ``,
		// 定义首页分享图片
		"share_cover":    ``,
		"placeholder":    `请输入关键字搜索`, // 小说名
		"online_service": true,
	})
}

//GetAPIToken 获取 jwt token
func GetAPIToken(c echo.Context) error {

	fromid, _ := strconv.Atoi(c.QueryParam("fromid"))
	code := c.QueryParam("code")
	provider := c.QueryParam("provider")
	ret, _ := cpi.GetOpenIDForApp(code, ``, ``)
	if code != "" && ret.OpenID != "" {
		cf := cpi.GetConf()
		fans, err := cpi.GetFansByOpenID(ret.OpenID)
		if fans.Provider == `` {
			fans.Provider = provider
		}

		// 增加用户被邀请次数
		if fromid > 0 && uint(fromid) != fans.ID {
			fans.InvitationTotal++
			if cf.Search.InvitationNo > 0 && fromid == cf.Search.InvitationNo { //特邀人员(设置邀请暗号)
				fans.Level = 5
			}
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

		return c.JSON(http.StatusOK, echo.Map{
			"jumpappid":   "", // 强制跳转其它小程序
			"token":       t,
			"uid":         fans.ID,
			"level":       fans.Level,
			"ismini":      0,
			"can_create":  fans.Level, // 允许创建内容
			"list_screen": cf.Ad.ListScreen,
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
			// "info_tips_banner": infoTipsBanner, // 点击广告开启自动加载更多功能
			// "info_tips_grid": infoTipsGrid, // 详细页格子广告
			"info_tips_banner": cf.Ad.InfoBanner, // 点击广告开启自动加载更多功能
			// "info_tips_grid": cf.Ad.InfoGrid, // 详细页格子广告
			"autoload_tips": `观看6~15秒视频广告开启无弹窗自动加载功能`,

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
			"online_service":    false,
			"info_force_reward": false, // 强制广告
			"info_video_adlt":   2,     //详情页面视频轮循总数
			"info_video_adlm":   0,     //详情页面视频轮循开始余量
			// "info_grid_adlt":    2,    //详情页面格子广告轮循总数
			// "info_grid_adlm":    1,    //详情页面格子广告轮循开始余量
			"info_banner_adlt": 2, //详情页面Banner轮循总数
			"info_banner_adlm": 1, //详情页面Banner轮循开始余量
			"info_screen_adlt": 5, //详情页面插屏广告轮循总数
			"info_screen_adlm": 3, //详情页面插屏广告轮循开始余量
		})

	}

	return echo.ErrUnauthorized
}

//GetAPIToken8 获取 jwt token PRO  被举报了没广告了
func GetAPIToken8(c echo.Context) error {
	fromid, _ := strconv.Atoi(c.QueryParam("fromid"))
	code := c.QueryParam("code")
	provider := c.QueryParam("provider")
	cf := cpi.GetConf()

	ret, _ := cpi.GetOpenID(code)
	// ret, _ := cpi.GetOpenIDForApp(code, cf.ReaderMinAppTwo.AppID, cf.ReaderMinAppTwo.AppSecret)
	if code != "" && ret.OpenID != "" {
		fans, err := cpi.GetFansByOpenID(ret.OpenID)
		if fans.Provider == `` {
			fans.Provider = provider
		}

		// 增加用户被邀请次数
		if fromid > 0 && uint(fromid) != fans.ID {
			fans.InvitationTotal++
			if cf.Search.InvitationNo > 0 && fromid == cf.Search.InvitationNo { //特邀人员(设置邀请暗号)
				fans.Level = 5
			}
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

		var infoTipsBanner, infoTipsCustom string
		day := time.Now().Day()
		var uid = int(fans.ID)
		var inum = (day + uid) % 3 //机率控制
		if inum == 0 {             // 日期加uid求余 为0 给banner 为 1 给grid
			// infoTipsCustom = `adunit-eb46e70f80c319ff`
		} else if inum == 1 {
			// infoTipsBanner = `adunit-8ff0e12978abbb22`
		} else if inum == 2 {
			// infoTipsBanner = `adunit-8ff0e12978abbb22`
		}

		var canCreate = 1
		if fans.Level > 2 {
			// canCreate = 1
		}

		version := c.QueryParam("version")
		var ismini = 0
		if cf.Search.LimitLevel || version == cf.Search.DevVersion { // 开启严格检查
			if fans.LoginTotal < 20 {
				ismini = 1
				canCreate = 0
			}
		}

		var jumpappid = `wx8664d56a896e375b` // mbqt wx8664d56a896e375b

		// 蜘蛛来的，给采集相关内容
		var req = c.Request()
		if version != cf.Search.DevVersion && strings.Contains(req.Header.Get("User-Agent"), `mpcrawler`) { // 获取通用 token  新推荐阅读
			jumpappid = `` // 蜘蛛给访问所有数据
		}
		//
		return c.JSON(http.StatusOK, echo.Map{
			"jumpappid":        jumpappid,        // 强制跳转其它小程序 lyj wx359657b0849ee636
			"jumpwebpage":      ``,               // 强制跳转网站阅读 http://r.1x7q.cn/#
			"jumpwebtips":      `已复制网址，请使用浏览器访问`, // 强制跳转网站阅读
			"token":            t,
			"uid":              fans.ID,
			"level":            0,
			"ismini":           ismini,
			"show_tips_next":   0,
			"can_create":       canCreate, // 允许创建内容
			"info_screen":      ``,
			"info_banner":      `adunit-8ff0e12978abbb22`,
			"info_custom":      `adunit-eb46e70f80c319ff`,
			"info_tips_banner": infoTipsBanner, // 点击广告开启自动加载更多功能
			"info_tips_custom": infoTipsCustom, // 详细页格子广告
			"autoload_tips":    `观看视频开启自动加载无弹窗模式`,
			// "autoload_tips": `体验广告6秒开启自动加载无弹窗模式`,
			// "top_home_video": `adunit-6a6203ae9a1f4252`,
			// "list_video": `adunit-4d779b9509cfa7a8`,
			// "cata_video": `adunit-61660192b3436fe7`,
			"info_video": `adunit-e57f377ec7c59a5d`,
			// "info_reward": `adunit-37d73c4714563ea5`,
			"top_home_custom": `adunit-29d916b6d72e9aef`,
			"list_custom":     `adunit-18a4d557507b274b`,
			"cata_custom":     `adunit-18a4d557507b274b`,
			"info_reward":     `adunit-c57dc483b3e62ce1`,
			// 定义首页分享标题
			"share_title": cf.ReaderMinAppThree.AppTitle,
			// 定义首页分享图片
			"share_cover":       cf.ReaderMinAppThree.AppCover,
			"placeholder":       cf.ReaderMinAppThree.AppSearch, // 小说名
			"online_service":    false,
			"info_force_reward": true, // 强制广告
			"info_video_adlt":   4,    //详情页面视频轮循总数
			"info_video_adlm":   1,    //详情页面视频轮循开始余量
			"info_custom_adlt":  2,    //详情页面格子广告轮循总数
			"info_custom_adlm":  0,    //详情页面格子广告轮循开始余量
			"info_banner_adlt":  4,    //详情页面Banner轮循总数
			"info_banner_adlm":  3,    //详情页面Banner轮循开始余量
			"info_screen_adlt":  5,    //详情页面插屏广告轮循总数
			"info_screen_adlm":  3,    //详情页面插屏广告轮循开始余量
		})
	}

	return echo.ErrUnauthorized
}

//GetAPIToken2 获取 jwt token 搜书大师
func GetAPIToken2(c echo.Context) error {
	fromid, _ := strconv.Atoi(c.QueryParam("fromid"))
	code := c.QueryParam("code")
	provider := c.QueryParam("provider")
	cf := cpi.GetConf()
	ret, _ := cpi.GetOpenIDForApp(code, cf.ReaderMinAppTwo.AppID, cf.ReaderMinAppTwo.AppSecret)
	if code != "" && ret.OpenID != "" {
		fans, err := cpi.GetFansByOpenID(ret.OpenID)
		if fans.Provider == `` {
			fans.Provider = provider
		}

		// 增加用户被邀请次数
		if fromid > 0 && uint(fromid) != fans.ID {
			fans.InvitationTotal++
			if cf.Search.InvitationNo > 0 && fromid == cf.Search.InvitationNo { //特邀人员(设置邀请暗号)
				fans.Level = 5
			}
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

		var jumpappid = `wx331f3c3e2761f080`
		if fans.LoginTotal > 3 {
			// jumpappid = `wx8ffa5a58c0bb3589`
		}

		var infoTipsBanner, infoTipsCustom string
		// if fans.LoginTotal > 3 { // 大于3（老用户了，随机给广告点击）
		day := time.Now().Day()
		var uid = int(fans.ID)
		var inum = (day + uid) % 3 //机率控制
		if inum == 0 {             // 日期加uid求余 为0 给banner 为 1 给grid
			infoTipsCustom = `adunit-c0a4c9c06c1bfb27`
		} else if inum == 1 {
			infoTipsBanner = `adunit-80ab5cf805e61964`
		} else if inum == 2 {
			infoTipsBanner = `adunit-80ab5cf805e61964`
		}
		// }

		// 蜘蛛来的，给采集相关内容
		version := c.QueryParam("version")
		var req = c.Request()
		if version != cf.Search.DevVersion && strings.Contains(req.Header.Get("User-Agent"), `mpcrawler`) { // 获取通用 token  新推荐阅读
			jumpappid = `` // 蜘蛛给访问所有数据
		}

		return c.JSON(http.StatusOK, echo.Map{
			"jumpappid":        jumpappid, // 强制跳转其它小程序
			"token":            t,
			"jumpwebpage":      ``,               // 强制跳转网站阅读
			"jumpwebtips":      `已复制网址，请使用浏览器访问`, // 强制跳转网站阅读
			"uid":              fans.ID,
			"level":            0,
			"ismini":           0,
			"show_tips_next":   0,
			"can_create":       1, // 允许创建内容
			"info_screen":      `adunit-0118779b141995e4`,
			"info_banner":      `adunit-80ab5cf805e61964`,
			"info_custom":      `adunit-c0a4c9c06c1bfb27`,
			"info_tips_banner": infoTipsBanner, // 点击广告开启自动加载更多功能
			"info_tips_custom": infoTipsCustom, // 详细页格子广告
			"autoload_tips":    `观看视频开启自动加载无弹窗模式`,
			// "autoload_tips": `体验广告6秒开启自动加载无弹窗模式`,
			// "top_home_video": `adunit-6a6203ae9a1f4252`,
			// "list_video": `adunit-4d779b9509cfa7a8`,
			// "cata_video": `adunit-61660192b3436fe7`,
			"info_video": `adunit-e21a2857faff7fba`,
			// "info_reward": `adunit-37d73c4714563ea5`,
			"top_home_custom": `adunit-6b3c3877de16d635`,
			"list_custom":     `adunit-c0a4c9c06c1bfb27`,
			"cata_custom":     `adunit-c0a4c9c06c1bfb27`,
			"info_reward":     `adunit-790a8d650d5c71b2`,
			// 定义首页分享标题
			"share_title": cf.ReaderMinAppThree.AppTitle,
			// 定义首页分享图片
			"share_cover":       cf.ReaderMinAppThree.AppCover,
			"placeholder":       cf.ReaderMinAppThree.AppSearch, // 小说名
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
		})
	}

	return echo.ErrUnauthorized
}

//GetAPIToken3 获取 jwt token 笔趣阁在线
func GetAPIToken3(c echo.Context) error {
	fromid, _ := strconv.Atoi(c.QueryParam("fromid"))
	code := c.QueryParam("code")
	provider := c.QueryParam("provider")
	version := c.QueryParam("version")
	cf := cpi.GetConf()
	ret, _ := cpi.GetOpenIDForApp(code, cf.ReaderMinAppFour.AppID, cf.ReaderMinAppFour.AppSecret)
	// fmt.Println(err, code, cf.ReaderMinAppThree.AppID, cf.ReaderMinAppThree.AppSecret)
	if code != "" && ret.OpenID != "" {
		fans, err := cpi.GetFansByOpenID(ret.OpenID)
		if fans.Provider == `` {
			fans.Provider = provider
		}

		// 增加用户被邀请次数
		if fromid > 0 && uint(fromid) != fans.ID {
			fans.InvitationTotal++
			if cf.Search.InvitationNo > 0 && fromid == cf.Search.InvitationNo { //特邀人员(设置邀请暗号)
				fans.Level = 5
			}
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
		var canCreate = 1

		var ismini = 0
		if cf.Search.LimitLevel || version == cf.Search.DevVersion { // 开启严格检查
			if fans.LoginTotal < 10 {
				canCreate = 0
			}
			ismini = 1
		} else {
			// .Header.Get("User-Agent")
		}

		var jumpappid = ``        // wx90dee998347266dd 新推荐阅读
		if fans.LoginTotal > 10 { // 访问次数大于5去Pro
			// ismini = 1
			// jumpappid = `wx359657b0849ee636` //驴友网 wx359657b0849ee636  免费版权图 wx8664d56a896e375b  强制跳转 搜书大师 wxe70eee58e64c7ac7
		}

		var infoTipsBanner, infoTipsGrid string

		if fans.LoginTotal > 0 { // 大于3（老用户了，随机给广告点击）
			day := time.Now().Day()
			var uid = int(fans.ID)
			var inum = (day + uid) % 3 //机率控制
			if inum == 0 {             // 日期加uid求余 为0 给banner 为 1 给grid
				infoTipsBanner = cf.Ad.InfoBanner
			} else if inum == 1 {
				// infoTipsGrid = cf.Ad.InfoGrid
			} else if inum == 2 {
				infoTipsBanner = cf.Ad.InfoBanner
			}
		}

		// 蜘蛛来的，给采集相关内容
		version := c.QueryParam("version")
		var req = c.Request()
		if version != cf.Search.DevVersion && strings.Contains(req.Header.Get("User-Agent"), `mpcrawler`) { // 获取通用 token  新推荐阅读
			jumpappid = `` // 蜘蛛给访问所有数据
		}

		return c.JSON(http.StatusOK, echo.Map{
			"jumpappid":   jumpappid,        // cf.ReaderMinAppThree.JumpAppID, // 强制跳转其它小程序
			"jumpwebpage": ``,               // 强制跳转网站阅读
			"jumpwebtips": `已复制网址，请使用浏览器访问`, // 强制跳转网站阅读
			"token":       t,
			"uid":         fans.ID,
			"level":       fans.Level,
			"ismini":      ismini,
			"can_create":  canCreate, // 允许创建内容
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
			"info_banner":      cf.Ad.InfoBanner,
			"info_tips_banner": infoTipsBanner, // 点击广告开启自动加载更多功能
			"info_tips_grid":   infoTipsGrid,   // 详细页格子广告
			// "info_tips_banner": cf.Ad.InfoBanner, // 点击广告开启自动加载更多功能
			// "info_tips_grid": cf.Ad.InfoGrid, // 详细页格子广告
			"autoload_tips": `观看视频开启自动加载无弹窗模式`,
			// "autoload_tips": `体验广告6秒开启自动加载无弹窗模式`,

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
			"share_title": cf.ReaderMinAppFour.AppTitle,
			// 定义首页分享图片
			"share_cover":       cf.ReaderMinAppFour.AppCover,
			"placeholder":       cf.ReaderMinAppFour.AppSearch, // 小说名
			"online_service":    true,
			"info_force_reward": true, // 强制广告
			"info_video_adlt":   2,    //详情页面视频轮循总数
			"info_video_adlm":   0,    //详情页面视频轮循开始余量
			// "info_grid_adlt":    4,    //详情页面格子广告轮循总数
			// "info_grid_adlm":    3,    //详情页面格子广告轮循开始余量
			"info_banner_adlt": 2, //详情页面Banner轮循总数
			"info_banner_adlm": 1, //详情页面Banner轮循开始余量
			"info_screen_adlt": 5, //详情页面插屏广告轮循总数
			"info_screen_adlm": 3, //详情页面插屏广告轮循开始余量
		})
	}
	return echo.ErrUnauthorized
}

//GetAPIToken4 获取 jwt token 笔趣阁plus 未接入的
func GetAPIToken4(c echo.Context) error {
	fromid, _ := strconv.Atoi(c.QueryParam("fromid"))
	code := c.QueryParam("code")
	provider := c.QueryParam("provider")
	version := c.QueryParam("version")
	cf := cpi.GetConf()
	ret, _ := cpi.GetOpenIDForApp(code, cf.ReaderMinAppFour.AppID, cf.ReaderMinAppThree.AppSecret)
	// fmt.Println(err, code, cf.ReaderMinAppThree.AppID, cf.ReaderMinAppThree.AppSecret)
	if code != "" && ret.OpenID != "" {
		fans, err := cpi.GetFansByOpenID(ret.OpenID)
		if fans.Provider == `` {
			fans.Provider = provider
		}

		// 增加用户被邀请次数
		if fromid > 0 && uint(fromid) != fans.ID {
			fans.InvitationTotal++
			if cf.Search.InvitationNo > 0 && fromid == cf.Search.InvitationNo { //特邀人员(设置邀请暗号)
				fans.Level = 5
			}
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
		var canCreate = 0
		if fans.LoginTotal > 5 {
			canCreate = 1
		}

		var ismini = 0
		if cf.Search.LimitLevel || version == cf.Search.DevVersion { // 开启严格检查
			if fans.LoginTotal < 10 {
				ismini = 1
				canCreate = 0
			}
		}

		var jumpappid = `wx8ffa5a58c0bb3589`
		if fans.LoginTotal < 3 {
			jumpappid = `wx8ffa5a58c0bb3589`
		}

		return c.JSON(http.StatusOK, echo.Map{
			"jumpappid":  jumpappid, // cf.ReaderMinAppThree.JumpAppID, // 强制跳转其它小程序
			"token":      t,
			"uid":        fans.ID,
			"level":      fans.Level,
			"ismini":     ismini,
			"can_create": canCreate, // 允许创建内容
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
			// "info_tips_banner": infoTipsBanner, // 点击广告开启自动加载更多功能
			// "info_tips_grid": infoTipsGrid, // 详细页格子广告
			"info_tips_banner": cf.Ad.InfoBanner, // 点击广告开启自动加载更多功能
			// "info_tips_grid": cf.Ad.InfoGrid, // 详细页格子广告
			"autoload_tips": `观看视频开启自动加载无弹窗模式`,
			// "autoload_tips": `体验广告6秒开启自动加载无弹窗模式`,

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
			"share_title": cf.ReaderMinAppThree.AppTitle,
			// 定义首页分享图片
			"share_cover":       cf.ReaderMinAppThree.AppCover,
			"placeholder":       cf.ReaderMinAppThree.AppSearch, // 小说名
			"online_service":    true,
			"info_force_reward": true, // 强制广告
			"info_video_adlt":   2,    //详情页面视频轮循总数
			"info_video_adlm":   0,    //详情页面视频轮循开始余量
			// "info_grid_adlt":    2,    //详情页面格子广告轮循总数
			// "info_grid_adlm":    1,    //详情页面格子广告轮循开始余量
			"info_banner_adlt": 2, //详情页面Banner轮循总数
			"info_banner_adlm": 1, //详情页面Banner轮循开始余量
			"info_screen_adlt": 5, //详情页面插屏广告轮循总数
			"info_screen_adlm": 3, //详情页面插屏广告轮循开始余量
		})
	}
	return echo.ErrUnauthorized
}

//GetAPIToken6 获取 jwt token 免版权图，暂时做个中转试试
func GetAPIToken6(c echo.Context) error {

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

	rand.Seed(time.Now().UnixNano())
	inum := rand.Intn(3) // 先搞低些广告出现机率

	var infoTipsBanner, infoTipsCustom string
	if true {
		if inum == 1 {
			infoTipsBanner = `adunit-a237f95dd4ce9ae7`
		} else if inum == 2 {
			infoTipsBanner = `adunit-a237f95dd4ce9ae7`
		} else if inum == 3 {
			infoTipsCustom = ``
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		// wx8ffa5a58c0bb3589 推荐阅读
		"jumpappid":        ``, // wxe70eee58e64c7ac7  // 强制跳转搜书大师  // 这个准备不做了，怕被抓鸡脚
		"bookjumpappid":    ``,
		"articlejumpappid": ``, //
		"token":            t,
		"uid":              -1,
		"level":            0,
		"ismini":           0,
		"can_create":       1, // 允许创建内容
		"home_screen_adid": `adunit-44763f52c54f72f9`,
		"info_screen":      `adunit-44763f52c54f72f9`,
		"info_banner":      `adunit-a237f95dd4ce9ae7`,
		"info_custom":      `adunit-ade0b17378833a01`,
		"info_tips_banner": infoTipsBanner, // 点击广告开启自动加载更多功能
		"info_tips_custom": infoTipsCustom, // 详细页格子广告
		"autoload_tips":    `观看视频开启自动加载无弹窗模式`,
		// "autoload_tips": `体验广告6秒开启自动加载无弹窗模式`,
		// "top_home_video": `adunit-8d6906f779544df6`,
		// "list_video": `adunit-8d6906f779544df6`,
		// "cata_video": `adunit-8d6906f779544df6`,
		"info_video":      `adunit-8d6906f779544df6`,
		"top_home_custom": `adunit-44122f4a8ef3d7d0`,
		"list_custom":     `adunit-ade0b17378833a01`,
		"cata_custom":     `adunit-ade0b17378833a01`,
		"info_reward":     `adunit-37d73c4714563ea5`,
		// 定义首页分享标题
		"share_title": cf.ReaderMinAppThree.AppTitle,
		// 定义首页分享图片
		"share_cover":       cf.ReaderMinAppThree.AppCover,
		"placeholder":       `请输入关键字搜索`, // 小说名
		"online_service":    false,
		"info_force_reward": true, // 强制广告
		"info_video_adlt":   4,    //详情页面视频轮循总数
		"info_video_adlm":   3,    //详情页面视频轮循开始余量
		"info_custom_adlt":  2,    //详情页面格子广告轮循总数
		"info_custom_adlm":  0,    //详情页面格子广告轮循开始余量
		"info_banner_adlt":  4,    //详情页面Banner轮循总数
		"info_banner_adlm":  1,    //详情页面Banner轮循开始余量
		"info_screen_adlt":  5,    //详情页面插屏广告轮循总数
		"info_screen_adlm":  3,    //详情页面插屏广告轮循开始余量

	})

}

//GetAPIToken7 获取 jwt token 新推荐阅读
func GetAPIToken7(c echo.Context) error {

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

	rand.Seed(time.Now().UnixNano())
	inum := rand.Intn(3) // 先搞低些广告出现机率

	var infoTipsBanner, infoTipsCustom string
	infoTipsBanner = ``
	if inum == 1 {
		infoTipsBanner = `adunit-0d62bae54bcefd36`
	} else if inum == 2 {
		infoTipsCustom = `adunit-6b354d2130f204aa`
	}

	return c.JSON(http.StatusOK, echo.Map{
		"jumpappid":        ``,               //
		"jumpwebpage":      ``,               // 强制跳转网站阅读
		"jumpwebtips":      `已复制网址，请使用浏览器访问`, // 强制跳转网站阅读
		"token":            t,
		"uid":              -1,
		"level":            0,
		"ismini":           0,
		"can_create":       1, // 允许创建内容
		"info_screen":      `adunit-f2f43997333bd86d`,
		"info_banner":      `adunit-0d62bae54bcefd36`,
		"info_custom":      `adunit-6b354d2130f204aa`,
		"info_tips_banner": infoTipsBanner, // 点击广告开启自动加载更多功能
		"info_tips_custom": infoTipsCustom, // 详细页格子广告
		"autoload_tips":    `观看视频开启自动加载无弹窗模式`,
		// "autoload_tips": `体验广告6秒开启自动加载无弹窗模式`,
		// "top_home_video": `adunit-997349cedbfe172f`,
		// "list_video": `adunit-997349cedbfe172f`,
		// "cata_video": `adunit-997349cedbfe172f`,
		"info_video": `adunit-b528ceb7836c247f`,
		// "info_reward": `adunit-37d73c4714563ea5`,
		"top_home_custom": `adunit-7931b9985beaf4db`,
		"list_custom":     `adunit-6b354d2130f204aa`,
		"cata_custom":     `adunit-6b354d2130f204aa`,
		"info_reward":     `adunit-756e936e72536645`,
		// 定义首页分享标题
		"share_title": cf.ReaderMinAppThree.AppTitle,
		// 定义首页分享图片
		"share_cover":       cf.ReaderMinAppThree.AppCover,
		"placeholder":       cf.ReaderMinAppThree.AppSearch, // 小说名
		"online_service":    false,
		"info_force_reward": true, // 强制广告
		"info_video_adlt":   4,    //详情页面视频轮循总数
		"info_video_adlm":   1,    //详情页面视频轮循开始余量
		"info_custom_adlt":  2,    //详情页面格子广告轮循总数
		"info_custom_adlm":  0,    //详情页面格子广告轮循开始余量
		"info_banner_adlt":  4,    //详情页面Banner轮循总数
		"info_banner_adlm":  3,    //详情页面Banner轮循开始余量
		"info_screen_adlt":  5,    //详情页面插屏广告轮循总数
		"info_screen_adlm":  4,    //详情页面插屏广告轮循开始余量（进去给就插屏）

	})

}

//GetAPIToken9 获取 jwt token 驴友记
func GetAPIToken9(c echo.Context) error {

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

	rand.Seed(time.Now().UnixNano())
	inum := rand.Intn(3) // 先搞低些广告出现机率

	var infoTipsBanner, infoTipsCustom string
	infoTipsBanner = ``
	if inum == 1 {
		infoTipsBanner = `adunit-c0d2320d02a94006`
	} else if inum == 2 {
		infoTipsCustom = `adunit-c9618bd19a0ed146`
	}
	return c.JSON(http.StatusOK, echo.Map{

		"jumpappid":        ``,   
		"bookjumpappid":    `wx8664d56a896e375b`,
		"articlejumpappid": `wx8664d56a896e375b`, //
		"jumpwebpage":      ``,               // 强制跳转网站阅读
		"jumpwebtips":      `已复制网址，请使用浏览器访问`, // 强制跳转网站阅读
		"token":            t,
		"uid":              -1,
		"level":            0,
		"ismini":           0,
		"hiderec":          1,
		"hidelog":          1,
		"can_create":       1, // 允许创建内容
		"info_screen":      `adunit-6584f905ac888622`,
		"info_banner":      `adunit-c0d2320d02a94006`,
		"info_custom":      `adunit-c9618bd19a0ed146`,
		"info_tips_banner": infoTipsBanner, // 点击广告开启自动加载更多功能
		"info_tips_custom": infoTipsCustom, // 详细页格子广告
		"autoload_tips":    `观看视频开启自动加载无弹窗模式`,
		// "autoload_tips": `体验广告6秒开启自动加载无弹窗模式`,
		// "top_home_video": `adunit-cc2f19cdc09c7a48`,
		// "list_video": `adunit-cc2f19cdc09c7a48`,
		// "cata_video": `adunit-cc2f19cdc09c7a48`,
		"info_video":      `adunit-a842a36d2700a76c`,
		"info_reward":     `adunit-70cea938ef5025dc`,
		"top_home_custom": `adunit-c9618bd19a0ed146`,
		"list_custom":     `adunit-c9618bd19a0ed146`,
		"cata_custom":     `adunit-c9618bd19a0ed146`,
		// "info_reward": `adunit-756e936e72536645`,
		// 定义首页分享标题
		"share_title": cf.ReaderMinAppThree.AppTitle,
		// 定义首页分享图片
		"share_cover":       cf.ReaderMinAppThree.AppCover,
		"placeholder":       cf.ReaderMinAppThree.AppSearch, // 小说名
		"online_service":    true,
		"info_force_reward": true, // 强制广告
		"info_video_adlt":   4,    //详情页面视频轮循总数
		"info_video_adlm":   1,    //详情页面视频轮循开始余量
		"info_custom_adlt":  2,    //详情页面格子广告轮循总数
		"info_custom_adlm":  0,    //详情页面格子广告轮循开始余量
		"info_banner_adlt":  4,    //详情页面Banner轮循总数
		"info_banner_adlm":  3,    //详情页面Banner轮循开始余量
		"info_screen_adlt":  5,    //详情页面插屏广告轮循总数
		"info_screen_adlm":  4,    //详情页面插屏广告轮循开始余量

	})

}

//GetAPIToken11  看书助手
func GetAPIToken11(c echo.Context) error {

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

	rand.Seed(time.Now().UnixNano())
	inum := rand.Intn(3) // 先搞低些广告出现机率

	var infoTipsBanner, infoTipsCustom string
	infoTipsBanner = ``
	if inum == 1 {
		infoTipsBanner = `adunit-c0d2320d02a94006`
	} else if inum == 2 {
		infoTipsCustom = `adunit-c9618bd19a0ed146`
	}
	return c.JSON(http.StatusOK, echo.Map{

		"jumpappid":        ``,               //
		"jumpwebpage":      ``,               // 强制跳转网站阅读
		"jumpwebtips":      `已复制网址，请使用浏览器访问`, // 强制跳转网站阅读
		"token":            t,
		"uid":              -1,
		"level":            0,
		"ismini":           0,
		"can_create":       1, // 允许创建内容
		"info_screen":      `adunit-6584f905ac888622`,
		"info_banner":      `adunit-c0d2320d02a94006`,
		"info_custom":      `adunit-c9618bd19a0ed146`,
		"info_tips_banner": infoTipsBanner, // 点击广告开启自动加载更多功能
		"info_tips_custom": infoTipsCustom, // 详细页格子广告
		"autoload_tips":    `观看视频开启自动加载无弹窗模式`,
		// "autoload_tips": `体验广告6秒开启自动加载无弹窗模式`,
		// "top_home_video": `adunit-cc2f19cdc09c7a48`,
		// "list_video": `adunit-cc2f19cdc09c7a48`,
		// "cata_video": `adunit-cc2f19cdc09c7a48`,
		"info_video":      `adunit-a842a36d2700a76c`,
		"info_reward":     `adunit-70cea938ef5025dc`,
		"top_home_custom": `adunit-c9618bd19a0ed146`,
		"list_custom":     `adunit-c9618bd19a0ed146`,
		"cata_custom":     `adunit-c9618bd19a0ed146`,
		// "info_reward": `adunit-756e936e72536645`,
		// 定义首页分享标题
		"share_title": cf.ReaderMinAppThree.AppTitle,
		// 定义首页分享图片
		"share_cover":       cf.ReaderMinAppThree.AppCover,
		"placeholder":       cf.ReaderMinAppThree.AppSearch, // 小说名
		"online_service":    true,
		"info_force_reward": true, // 强制广告
		"info_video_adlt":   4,    //详情页面视频轮循总数
		"info_video_adlm":   1,    //详情页面视频轮循开始余量
		"info_custom_adlt":  2,    //详情页面格子广告轮循总数
		"info_custom_adlm":  0,    //详情页面格子广告轮循开始余量
		"info_banner_adlt":  4,    //详情页面Banner轮循总数
		"info_banner_adlm":  3,    //详情页面Banner轮循开始余量
		"info_screen_adlt":  5,    //详情页面插屏广告轮循总数
		"info_screen_adlm":  4,    //详情页面插屏广告轮循开始余量

	})

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
				wxto,
				``,
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
