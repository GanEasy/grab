package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/GanEasy/grab"
	cpi "github.com/GanEasy/minappapi"
	"github.com/labstack/echo"
)

//CheckSubcribeUpdate  每天处理订阅更新
func CheckSubcribeUpdate() {
	ticker := time.NewTicker(time.Hour * 6)
	for _ = range ticker.C {
		go cpi.RunSubcribePostUpdateCheck()
	}
}
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, `open.readfollow.com!
			/gethtml?url=&find= <br>
			/getbookchapters?url= <br>
			/getbookchapterinfo?url= <br>
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
		id, _ := strconv.Atoi(c.QueryParam("id"))
		ret := cpi.GetPostByID(int64(id))
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

	// 获取列表
	e.GET("/getlist", func(c echo.Context) error {
		urlStr := c.QueryParam("url")
		if urlStr == "" {
			return c.JSON(http.StatusOK, "0")
		}
		ret, _ := cpi.GetBookMenu(urlStr)
		return c.JSON(http.StatusOK, ret)
	})

	// 获取正文
	e.GET("/getcontent", func(c echo.Context) error {
		urlStr := c.QueryParam("url")

		ret, _ := grab.GetContent(urlStr)
		return c.JSON(http.StatusOK, ret)
	})

	// 获取小说章节正文
	e.GET("/getbookcontent", func(c echo.Context) error {
		urlStr := c.QueryParam("url")
		if urlStr == "" {
			return c.JSON(http.StatusOK, "0")
		}
		ret, _ := cpi.GetBookContent(urlStr)
		return c.JSON(http.StatusOK, ret)
	})

	// 图标
	e.File("favicon.ico", "images/favicon.ico")
	e.Logger.Fatal(e.Start(":8009"))
	// e.Logger.Fatal(e.StartAutoTLS(":443"))

}
