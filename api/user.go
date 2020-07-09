package api

import (
	"fmt"
	"log"
	// "math/rand"
	"net/http"
	"strconv"
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
	code := c.QueryParam("code")
	ret, _ := cpi.GetOpenID(code)
	if code != "" && ret.OpenID != "" {
		fans, err := cpi.GetFansByOpenID(ret.OpenID)
		// if err != nil {
		// 	return err
		// }
		log.Println(ret)
		// Set custom claims
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
			"token":       t,
			"uid":         fans.ID,
			"level":       fans.Level,
			"score":       fans.Score,
			"total":       fans.Total,
			"home_screen": cf.Ad.HomeScreen,
			"list_screen": cf.Ad.ListScreen,
			"info_screen": cf.Ad.InfoScreen,

			"screen":      cf.Ad.Screen,
			"reward":      cf.Ad.Reward,
			"pre_video":   cf.Ad.PreVideo,
			"home_banner": cf.Ad.HomeBanner,
			"list_banner": cf.Ad.ListBanner,
			"info_banner": cf.Ad.InfoBanner,

			"home_video": cf.Ad.HomeVideo,
			// "list_video": cf.Ad.ListVideo,
			"info_video": cf.Ad.InfoVideo,
			// "home_pre_video": cf.Ad.PreVideo,
			// "list_pre_video": cf.Ad.PreVideo,
			// "info_pre_video": cf.Ad.PreVideo,

			// "home_reward": cf.Ad.Reward,
			// "list_reward": cf.Ad.Reward,
			// "info_reward": cf.Ad.Reward,
		})
	}

	return echo.ErrUnauthorized
}

//GetAPIToken 获取 jwt token
func GetAPIToken(c echo.Context) error {

	fromid, _ := strconv.Atoi(c.QueryParam("fromid"))
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

		// rand.Seed(time.Now().UnixNano())
		// inum := rand.Intn(10) // 先搞低些广告出现机率

		// var info_tips_banner,info_tips_grid string
		// if inum==1 {
		// 	info_tips_banner = cf.Ad.InfoBanner
		// }else if inum==2{
		// 	info_tips_grid =  cf.Ad.InfoGrid
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
			// "info_banner": cf.Ad.InfoBanner,
			// "info_tips_banner": info_tips_banner, // 点击广告开启自动加载更多功能
			// "info_tips_grid": info_tips_grid, // 详细页格子广告
			"info_tips_banner": cf.Ad.InfoBanner, // 点击广告开启自动加载更多功能
			// "info_tips_grid": cf.Ad.InfoGrid, // 详细页格子广告
			"autoload_tips": `体验广告6~15秒，开启自动加载+免打扰模式`,

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
			"info_force_reward": true, // 强制广告
			"info_video_adlt":   2,    //详情页面视频轮循总数
			"info_video_adlm":   0,    //详情页面视频轮循开始余量
			"info_grid_adlt":    2,    //详情页面格子广告轮循总数
			"info_grid_adlm":    1,    //详情页面格子广告轮循开始余量
			// "info_banner_adlt":  2,    //详情页面Banner轮循总数
			// "info_banner_adlm":  1,    //详情页面Banner轮循开始余量
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

			// rand.Seed(time.Now().UnixNano())
			// inum := rand.Intn(3) // 先搞低些广告出现机率

			// var info_tips_banner,info_tips_grid string
			// if inum==1 {
			// 	info_tips_banner = cf.Ad.InfoBanner
			// }else if inum==2{
			// 	info_tips_grid =  cf.Ad.InfoGrid
			// }

			// 用户登录次数大于5并且不是从分享页面来的
			if fans.LoginTotal > 5 && fromid < 1 {

				return c.JSON(http.StatusOK, echo.Map{
					"token":      t,
					"uid":        fans.ID,
					"level":      fans.Level,
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
					"info_banner":      cf.Ad.InfoBanner,
					"info_tips_banner": cf.Ad.InfoBanner, // 点击广告开启自动加载更多功能
					// "info_tips_banner": info_tips_banner, // 点击广告开启自动加载更多功能
					// "info_tips_grid": info_tips_grid, // 格子广告
					"autoload_tips": `体验广告6~15秒，开启自动加载+免打扰模式`,

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
					"info_force_reward": true, // 老人强制广告
					"info_video_adlt":   2,    //详情页面视频轮循总数
					"info_video_adlm":   0,    //详情页面视频轮循开始余量
					"info_banner_adlt":  2,    //详情页面Banner轮循总数
					"info_banner_adlm":  1,    //详情页面Banner轮循开始余量
					// "info_grid_adlt":    4,    //详情页面格子广告轮循总数
					// "info_grid_adlm":    1,    //详情页面格子广告轮循开始余量
					"info_screen_adlt": 10, //详情页面插屏广告轮循总数
					"info_screen_adlm": 8,  //详情页面插屏广告轮循开始余量
				})
			}
			// 新人访问体验要好些
			return c.JSON(http.StatusOK, echo.Map{
				"token":      t,
				"uid":        fans.ID,
				"level":      fans.Level,
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
				// "info_tips_banner": cf.Ad.InfoBanner, // 点击广告开启自动加载更多功能
				// "info_tips_grid": cf.Ad.InfoGrid, // 详细页格子广告
				"autoload_tips": `体验广告6~15秒，解锁自动加载功能`,

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
				"info_force_reward": false, // 新人不强制广告
				"info_video_adlt":   2,     //详情页面视频轮循总数
				"info_video_adlm":   0,     //详情页面视频轮循开始余量
				"info_banner_adlt":  2,     //详情页面Banner轮循总数
				"info_banner_adlm":  1,     //详情页面Banner轮循开始余量
				// "info_grid_adlt":    6,                    //详情页面格子广告轮循总数
				// "info_grid_adlm":    4,                    //详情页面格子广告轮循开始余量
				"info_screen_adlt": 10, //详情页面插屏广告轮循总数
				"info_screen_adlm": 7,  //详情页面插屏广告轮循开始余量
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
			"can_create":     1, // 允许创建内容

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
