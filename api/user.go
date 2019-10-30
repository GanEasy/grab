package api

import (
	"fmt"
	"log"
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
			"token": t,
			"uid":   fans.ID,
			"level": fans.Level,
			"score": fans.Score,
			"total": fans.Total,
			// "home_screen": cf.Ad.Screen,
			// "list_screen": cf.Ad.Screen,
			// "info_screen": cf.Ad.Screen,
			"screen":      cf.Ad.Screen,
			"reward":      cf.Ad.Reward,
			"pre_video":   cf.Ad.PreVideo,
			"home_banner": cf.Ad.HomeBanner,
			"list_banner": cf.Ad.ListBanner,
			"info_banner": cf.Ad.InfoBanner,

			"home_video": cf.Ad.HomeVideo,
			"list_video": cf.Ad.ListVideo,
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
	code := c.QueryParam("code")
	provider := c.QueryParam("provider")
	if provider == `weixin` {
		ret, _ := cpi.GetOpenID(code)
		if code != "" && ret.OpenID != "" {
			fans, err := cpi.GetFansByOpenID(ret.OpenID)
			if fans.Provider == `` {
				fans.Provider = provider
			}
			fans.LoginTotal++ // 增加一次授权访问次数
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
				"token":       t,
				"uid":         fans.ID,
				"level":       fans.Level,
				"score":       fans.Score,
				"total":       fans.Total,
				"home_screen": cf.Ad.Screen,
				"list_screen": cf.Ad.Screen,
				"info_screen": cf.Ad.Screen,
				"screen":      cf.Ad.Screen,
				"reward":      cf.Ad.Reward,
				"pre_video":   cf.Ad.PreVideo,
				"home_banner": cf.Ad.HomeBanner,
				"list_banner": cf.Ad.ListBanner,
				"info_banner": cf.Ad.InfoBanner,

				"home_video": cf.Ad.HomeVideo,
				"list_video": cf.Ad.ListVideo,
				"info_video": cf.Ad.InfoVideo,
				// "home_pre_video": cf.Ad.PreVideo,
				// "list_pre_video": cf.Ad.PreVideo,
				// "info_pre_video": cf.Ad.PreVideo,

				"home_reward": cf.Ad.Reward,
				"list_reward": cf.Ad.Reward,
				"info_reward": cf.Ad.Reward,

				// 定义首页分享标题
				"share_title": `全网小说资源免费转码阅读 | 笔趣阁小说转码工具 | 全本免费小说阅读`,
				// 定义首页分享图片
				"share_cover": ``,
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
