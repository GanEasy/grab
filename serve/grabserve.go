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
		list := grab.GetClassify()
		return c.JSON(http.StatusOK, list)
	})
	//  自定义资源列表
	e.GET("/resources", func(c echo.Context) error {
		list := grab.GetResources()
		return c.JSON(http.StatusOK, list)
	})

	//  获取自定义平台资源列表
	e.GET("/resource/:url", func(c echo.Context) error {
		url, _ := grab.DecodeURL(c.Param("url"))
		list := grab.GetResource(string(url))
		return c.JSON(http.StatusOK, list)
	})
	//  自定义专题列表
	e.GET("/topics", func(c echo.Context) error {
		list := grab.GetTopics()
		return c.JSON(http.StatusOK, list)
	})

	//  获取小说资源列表
	e.GET("/book/:url", func(c echo.Context) error {
		url, _ := grab.DecodeURL(c.Param("url"))
		list := grab.GetBooks(string(url))
		return c.JSON(http.StatusOK, list)
	})

	// 获取章节列表
	e.GET("/chapters/:url", func(c echo.Context) error {
		urlStr, _ := grab.DecodeURL(c.Param("url"))
		// list, _ := grab.GetBookChapters(string(urlStr))

		reader := grab.BookListReader{}
		list, _ := reader.GetList(urlStr)
		// list := grab.GetChapters(string(urlStr))
		return c.JSON(http.StatusOK, list)
	})

	// 获取章节详细内容
	e.GET("/chapter/:url", func(c echo.Context) error {
		// urlStr := c.QueryParam("url")
		urlStr, _ := grab.DecodeURL(c.Param("url"))
		reader := grab.BookInfoReader{}
		info, _ := reader.GetInfo(urlStr)
		// info, _ := grab.GetBookInfo(urlStr)
		return c.JSON(http.StatusOK, info)
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

	// 图标
	e.File("favicon.ico", "images/favicon.ico")
	e.Logger.Fatal(e.Start(":8041"))
	// e.Logger.Fatal(e.StartAutoTLS(":443"))

}
