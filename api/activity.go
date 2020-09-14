package api

import (
	"net/http"

	cpi "github.com/GanEasy/grab/core"
	"github.com/GanEasy/grab/reader"
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

	provider := c.QueryParam("provider")

	var level int32
	level = 3 // 4已经支持所有了(小说和漫画) 3支持小说，2什么都不支持

	// 对受限制的应用进行过滤
	cf := cpi.GetConf()

	if cf.Search.LimitLevel || version == cf.Search.DevVersion { // 开启严格检查或者当前版本在审核模式
		if provider == `weixin` {
			level = 2
		} else if provider == `qq` {
			level = 2
		} else if provider == `toutiao` {
			level = 2
		} else if provider == `web` {
			level = 5
		}
	}

	// 推荐内容，如果是隐藏邀请模式的，严格隐藏
	if cf.Search.LimitInvitation {
		openID := getOpenID(c)
		if openID == `` {
			level = 2
		} else {
			user, _ := getUser(openID)
			level = user.Level
			if level < 2 {
				level = 2
			}
			if user.Level<=2 && user.LoginTotal >= 10 {
				user.Level = 3
			}
		}
		
	}

	var rows = cpi.GetActivities()
	if len(rows) > 0 {
		var rp = map[string]int{}
		var itemlevel int32
		for _, v := range rows {
			// 只显示拥有权限的级别
			itemlevel = reader.GetPathLevel(v.WxTo)

			// 过滤掉做审核的内容
			if level > 2 && itemlevel==1 {
				itemlevel = 0
			}

			if level > itemlevel && itemlevel > 0 {
				//  过滤掉相同 title 的资源（重复的只显示最新一个）
				if _, ok := rp[v.Title]; !ok {
					rp[v.Title] = 1
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
		}
	}

	// if len(links) == 0 {
	// 	links = append(links,
	// 		Link{
	// 			Title: `Laravel 项目开发规范`,
	// 			Icon:  ``, // cuIcon-new
	// 			Type:  `link`,
	// 			Image: ``,
	// 			WxTo:  `/pages/catalog?drive=learnku&url=` + grab.EncodeURL(`https://learnku.com/docs/laravel-specification/5.5`),
	// 			Style: `arrow`,
	// 		})
	// 	links = append(links,
	// 		Link{
	// 			Title: `Dingo API 2.0.0 中文文档`,
	// 			Icon:  ``, // cuIcon-new
	// 			Type:  `link`,
	// 			Image: ``,
	// 			WxTo:  `/pages/catalog?drive=learnku&url=` + grab.EncodeURL(`https://learnku.com/docs/dingo-api/2.0.0`),
	// 			Style: `arrow`,
	// 		})
	// }

	return c.JSON(http.StatusOK, links)
}
