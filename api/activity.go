package api

import (
	"net/http"

	"github.com/GanEasy/grab"
	cpi "github.com/GanEasy/grab/core"
	"github.com/labstack/echo"
)

//NewActivity 新推荐
func NewActivity(c echo.Context) error {
	title := c.QueryParam("title")
	wxto := c.QueryParam("wxto")

	activity := cpi.AddActivity(title, wxto)
	return c.JSON(http.StatusOK, activity)
}

//GetActivities 获取100条推荐
func GetActivities(c echo.Context) error {
	var links = []Link{}

	version := c.QueryParam("version")
	cf := cpi.GetConf()
	if cf.Search.LimitLevel || version == cf.Search.DevVersion { // 开启严格检查

		links = append(links,
			Link{
				Title: `Laravel 项目开发规范`,
				Icon:  ``, // cuIcon-new
				Type:  `link`,
				Image: ``,
				WxTo:  `/pages/catalog?drive=learnku&url=` + grab.EncodeURL(`https://learnku.com/docs/laravel-specification/5.5`),
				Style: `arrow`,
			})
		links = append(links,
			Link{
				Title: `Dingo API 2.0.0 中文文档`,
				Icon:  ``, // cuIcon-new
				Type:  `link`,
				Image: ``,
				WxTo:  `/pages/catalog?drive=learnku&url=` + grab.EncodeURL(`https://learnku.com/docs/dingo-api/2.0.0`),
				Style: `arrow`,
			})
		return c.JSON(http.StatusOK, links)
	}
	var rows = cpi.GetActivities()
	if len(rows) > 0 {
		for _, v := range rows {
			links = append(links,
				Link{
					Title: v.Title,
					Icon:  ``, // cuIcon-new
					Type:  `link`,
					Image: ``,
					WxTo:  v.WxTo,
					Style: `arrow`,
				})
		}
	}
	return c.JSON(http.StatusOK, links)
}
