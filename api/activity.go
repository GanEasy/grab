package api

import (
	"net/http"

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
