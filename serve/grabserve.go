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

	//  自定义专题列表
	e.GET("/topics", func(c echo.Context) error {
		list := grab.GetTopics()
		return c.JSON(http.StatusOK, list)
	})

	//  获取第三方资源的分类(按驱动)
	e.GET("/categories/:url", func(c echo.Context) error {
		urlStr, _ := grab.DecodeURL(c.Param("url"))
		guide := grab.GetGuide(c.QueryParam("drive"))
		list, _ := guide.GetCategories(urlStr)
		return c.JSON(http.StatusOK, list)
	})

	//  获取可订阅目录列表
	e.GET("/list/:url", func(c echo.Context) error {
		urlStr, _ := grab.DecodeURL(c.Param("url"))
		drive := c.QueryParam("drive")
		guide := grab.GetGuide(drive)
		list, _ := guide.GetList(urlStr)
		return c.JSON(http.StatusOK, list)
	})

	// 获取目录(订阅目录内容)
	e.GET("/catalog/:url", func(c echo.Context) error {
		urlStr, _ := grab.DecodeURL(c.Param("url"))
		drive := c.QueryParam("drive")
		reader := grab.GetReader(drive)
		list, _ := reader.GetCatalog(urlStr)
		return c.JSON(http.StatusOK, list)
	})

	// 获取页面正文内容
	e.GET("/info/:url", func(c echo.Context) error {
		urlStr, _ := grab.DecodeURL(c.Param("url"))
		drive := c.QueryParam("drive")
		reader := grab.GetReader(drive)
		list, _ := reader.GetInfo(urlStr)
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
