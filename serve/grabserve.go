package main

import (
	"net/http"

	"github.com/GanEasy/grab"
	c "github.com/GanEasy/grab/api"
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
	e.GET("/gettoken", c.GetToken)
	// 解密数据内容(保存数据到库)
	e.GET("/crypt", c.Crypt)

	// 获取二维码(图片资源)
	e.GET("/qrcode", c.GetQrcode)

	// Restricted group
	api := e.Group("/api")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &c.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	api.Use(middleware.JWTWithConfig(config))

	api.GET("/checkopenid", c.CheckOpenID)
	// r.Use(middleware.JWT([]byte("secret")))

	//  粉丝关注列表
	api.GET("/follows", c.GetUserFollows)

	// 粉丝添加关注
	api.POST("/follows", c.CreateUserFollow)

	//  粉丝自定义源
	api.GET("/sources", c.GetUserSources)

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

	e.GET("/urlencode", func(c echo.Context) error {
		url := c.QueryParam("url")
		return c.JSON(http.StatusOK, grab.EncodeURL(url))
	})

	// 粉丝添加关注
	api.POST("/sources", c.CreateUserSource)

	// 获取html gethtml
	e.GET("/gethtml", func(c echo.Context) error {
		urlStr := c.QueryParam("url")
		find := c.QueryParam("find")
		html, _ := grab.GetHTMLContent(urlStr, find)
		return c.JSON(http.StatusOK, html)
	})

	// 获取获取文章 getarticle
	e.GET("/getarticle", func(c echo.Context) error {
		urlStr := c.QueryParam("url")
		ret, _ := grab.GetArticle(urlStr)
		return c.JSON(http.StatusOK, ret)
	})

	// 获取获取文章列表 getarticlelist
	e.GET("/getarticlelist", func(c echo.Context) error {
		urlStr := c.QueryParam("url")
		ret, _ := grab.GetArticleList(urlStr)
		return c.JSON(http.StatusOK, ret)
	})

	// 获取小说目录正文
	e.GET("/getbookchapters", func(c echo.Context) error {
		urlStr := c.QueryParam("url")
		ret, _ := grab.GetBookChapters(urlStr)
		return c.JSON(http.StatusOK, ret)
	})
	// 获取小说章节内容 getbookchapterinfo
	e.GET("/getbookchapterinfo", func(c echo.Context) error {
		urlStr := c.QueryParam("url")
		info, _ := grab.GetBookInfo(urlStr)
		return c.JSON(http.StatusOK, info)
	})

	// 获取Rss列表 getrsslist
	e.GET("/getrsslist", func(c echo.Context) error {
		urlStr := c.QueryParam("url")

		reader := grab.RssListReader{}
		list, err := reader.GetList(urlStr)
		if err == nil {
			return c.JSON(http.StatusOK, list)
		}
		return c.JSON(http.StatusOK, list)
	})

	// // 获取html列表
	// e.GET("/gethtmllist", func(c echo.Context) error {
	// 	urlStr := c.QueryParam("url")
	// 	find := c.QueryParam("find")
	// 	list, err := grab.GetHTMLList(urlStr, find)
	// 	if err == nil {
	// 		return c.JSON(http.StatusOK, list)
	// 	}
	// 	return c.JSON(http.StatusOK, list)
	// })

	//  自定义分类
	e.GET("/classify", func(c echo.Context) error {
		// list := grab.GetClassify()
		list := grab.GetResources()
		return c.JSON(http.StatusOK, list)
	})
	//  自定义资源列表
	e.GET("/resources", func(c echo.Context) error {
		list := grab.GetResources()
		return c.JSON(http.StatusOK, list)
	})

	//  获取自定义平台资源列表
	e.GET("/resource/:url", func(c echo.Context) error {
		urlStr, _ := grab.DecodeURL(c.Param("url"))
		drive := c.QueryParam("drive")
		reader := grab.GetBookReader(drive)
		list, _ := reader.GetCategories(urlStr)
		return c.JSON(http.StatusOK, list)
	})
	//  自定义专题列表
	e.GET("/topics", func(c echo.Context) error {
		list := grab.GetTopics()
		return c.JSON(http.StatusOK, list)
	})

	//  获取小说资源列表
	e.GET("/book/:url", func(c echo.Context) error {
		urlStr, _ := grab.DecodeURL(c.Param("url"))
		drive := c.QueryParam("drive")
		reader := grab.GetBookReader(drive)
		list, _ := reader.GetBooks(urlStr)
		return c.JSON(http.StatusOK, list)
	})

	// 获取章节列表
	e.GET("/chapters/:url", func(c echo.Context) error {
		urlStr, _ := grab.DecodeURL(c.Param("url"))
		drive := c.QueryParam("drive")
		reader := grab.GetBookReader(drive)
		list, _ := reader.GetChapters(urlStr)
		return c.JSON(http.StatusOK, list)
	})

	// 获取章节详细内容
	e.GET("/chapter/:url", func(c echo.Context) error {
		urlStr, _ := grab.DecodeURL(c.Param("url"))
		drive := c.QueryParam("drive")
		reader := grab.GetBookReader(drive)
		list, _ := reader.GetChapter(urlStr)
		return c.JSON(http.StatusOK, list)
	})

	//  get rss demo
	e.GET("/getrssdemo", func(c echo.Context) error {
		list := grab.RssDemoList()
		return c.JSON(http.StatusOK, list)
	})
	//  get article demo
	e.GET("/getarticledemo", func(c echo.Context) error {
		list := grab.ArticleDemoList()
		return c.JSON(http.StatusOK, list)
	})
	//  get book demo
	e.GET("/getbookdemo", func(c echo.Context) error {
		list := grab.BookDemoList()
		return c.JSON(http.StatusOK, list)
	})
	//  get book demo
	e.GET("/about", func(c echo.Context) error {
		list := grab.GetAbouts()
		return c.JSON(http.StatusOK, list)
	})

	// 图标
	e.File("favicon.ico", "images/favicon.ico")
	e.Logger.Fatal(e.Start(":8041"))
	// e.Logger.Fatal(e.StartAutoTLS(":443"))

}
