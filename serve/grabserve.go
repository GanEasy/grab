package main

import (
	"net/http"

	"github.com/GanEasy/grab"
	a "github.com/GanEasy/grab/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK,
			`https://open.readfollow.com build by golang, author yizenghui. 
/gethtml?url=&find=
/gethtmllist?url=&find=
/getbookchapters?url=
/getbookchapterinfo?url=
/getrsslist?url=
/getarticlelist?url=
/getarticle?url=
`)
	})

	// 获取用户签名
	e.GET("/gettoken", a.GetToken)
	// 解密数据内容(保存数据到库)
	e.GET("/crypt", a.Crypt)

	// 获取二维码(图片资源)
	e.GET("/qrcode", a.GetQrcode)

	// 从二维码进来跳哪里去
	e.GET("/qrcodejump", a.GetQrcodeWxto)

	// Restricted group
	api := e.Group("/api")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &a.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	api.Use(middleware.JWTWithConfig(config))

	api.GET("/checkopenid", a.CheckOpenID)
	// r.Use(middleware.JWT([]byte("secret")))

	//  粉丝关注列表
	api.GET("/follows", a.GetUserFollows)

	// 粉丝添加关注
	api.POST("/follows", a.CreateUserFollow)

	//  粉丝自定义源
	api.GET("/sources", a.GetUserSources)

	api.POST("/urlencode", func(c echo.Context) (err error) {
		type Data struct {
			URL string `json:"url" validate:"required"`
		}
		u := new(Data)
		if err = c.Bind(u); err != nil {
			return
		}
		// url := c.FormValue("url")
		return c.JSON(http.StatusOK, grab.EncodeURL(u.URL))
	})

	api.POST("/getdrive", func(c echo.Context) (err error) {
		type Data struct {
			URL string `json:"url" validate:"required"`
		}
		u := new(Data)
		if err = c.Bind(u); err != nil {
			return
		}
		type Ret struct {
			Key   string `json:"key"`
			Drive string `json:"drive"`
			Page  string `json:"page"`
		}
		ret := Ret{
			grab.EncodeURL(u.URL),
			``,
			``,
		}
		// url := c.FormValue("url")
		return c.JSON(http.StatusOK, ret)
	})

	e.GET("/urlencode", func(c echo.Context) error {
		url := c.QueryParam("url")
		return c.JSON(http.StatusOK, grab.EncodeURL(url))
	})

	// 粉丝添加关注
	api.POST("/sources", a.CreateUserSource)

	// 搜索
	api.GET("/search", a.SearchPosts)

	//  自定义分类
	api.GET("/classify", func(c echo.Context) error {
		// list := grab.GetClassify()
		// list := grab.GetWaitExamineClassify()
		list := grab.GetResources()
		return c.JSON(http.StatusOK, list)
	})
	//  自定义资源列表
	api.GET("/resources", func(c echo.Context) error {
		list := grab.GetResources()
		return c.JSON(http.StatusOK, list)
	})

	//  自定义专题列表
	api.GET("/topics", func(c echo.Context) error {
		list := grab.GetTopics()
		return c.JSON(http.StatusOK, list)
	})

	//  获取第三方资源的分类(按驱动)
	api.GET("/categories/:url", func(c echo.Context) error {
		urlStr, _ := grab.DecodeURL(c.Param("url"))
		guide := grab.GetGuide(c.QueryParam("drive"))
		list, _ := guide.GetCategories(urlStr)
		return c.JSON(http.StatusOK, list)
	})

	//  获取可订阅目录列表
	api.GET("/list/:url", func(c echo.Context) error {
		urlStr, _ := grab.DecodeURL(c.Param("url"))
		drive := c.QueryParam("drive")
		guide := grab.GetGuide(drive)
		list, _ := guide.GetList(urlStr)

		if drive == `qidian` || drive == `zongheng` || drive == `17k` || drive == `luoqiu` || drive == `booktxt` || drive == `bxwx` || drive == `uxiaoshuo` {
			go a.SyncPosts(list, 1)
		} else if drive == `manhwa` || drive == `r2hm` {
			go a.SyncPosts(list, 2)
		}

		return c.JSON(http.StatusOK, list)
	})

	// 获取目录(订阅目录内容)
	api.GET("/catalog/:url", func(c echo.Context) error {
		urlStr, _ := grab.DecodeURL(c.Param("url"))
		drive := c.QueryParam("drive")
		reader := grab.GetReader(drive)
		list, _ := reader.GetCatalog(urlStr)
		return c.JSON(http.StatusOK, list)
	})

	// 获取页面正文内容
	api.GET("/info/:url", func(c echo.Context) error {
		urlStr, _ := grab.DecodeURL(c.Param("url"))
		drive := c.QueryParam("drive")
		reader := grab.GetReader(drive)
		list, _ := reader.GetInfo(urlStr)
		return c.JSON(http.StatusOK, list)
	})
	//  get book demo
	api.GET("/about", func(c echo.Context) error {
		list := grab.GetAbouts()
		return c.JSON(http.StatusOK, list)
	})
	//  get book demo
	api.GET("/help", func(c echo.Context) error {
		list := grab.GetHelps()
		return c.JSON(http.StatusOK, list)
	})
	//  get book demo
	api.GET("/alldrive", func(c echo.Context) error {
		type Item struct {
			Name  string `json:"name" `
			Drive string `json:"drive" `
			Page  string `json:"page"`
		}
		var list = []Item{
			Item{`文本类`, `text`, `/pages/catalog`},
			Item{`图文类`, `article`, `/pages/catalog`},
			Item{`Rss`, `rss`, `/pages/rss`},
		}
		return c.JSON(http.StatusOK, list)
	})

	// 图标
	e.File("favicon.ico", "images/favicon.ico")
	e.Logger.Fatal(e.Start(":8041"))
	// e.Logger.Fatal(e.StartAutoTLS(":443"))

}
