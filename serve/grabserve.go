package main

import (
	"net/http"

	"github.com/GanEasy/grab"
	"github.com/labstack/echo"
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

	// 获取html gethtml
	e.GET("/gethtml", func(c echo.Context) error {
		urlStr := c.QueryParam("url")
		find := c.QueryParam("find")
		html, _ := grab.GetHTML(urlStr, find)
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
		list, err := grab.GetRssList(urlStr)
		if err == nil {
			return c.JSON(http.StatusOK, list)
		}
		return c.JSON(http.StatusFound, list)
	})

	// 获取html列表
	e.GET("/gethtmllist", func(c echo.Context) error {
		urlStr := c.QueryParam("url")
		find := c.QueryParam("find")
		list, err := grab.GetHTMLList(urlStr, find)
		if err == nil {
			return c.JSON(http.StatusOK, list)
		}
		return c.JSON(http.StatusFound, list)
	})

	// 图标
	e.File("favicon.ico", "images/favicon.ico")
	e.Logger.Fatal(e.Start(":8041"))
	// e.Logger.Fatal(e.StartAutoTLS(":443"))

}
