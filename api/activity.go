package api

import (
	"net/http"
	"strings"

	cpi "github.com/GanEasy/grab/core"
	"github.com/GanEasy/grab/reader"
	"github.com/labstack/echo"
)

//NewActivity 新推荐
func NewActivity(c echo.Context) error {
	title := c.QueryParam("title")
	wxto := c.QueryParam("wxto")

	activity := cpi.AddActivity(title, wxto)

	cpi.WxAppSubmitPage(wxto) //提交页面（望收录）

	return c.JSON(http.StatusOK, activity)
}

//RemoveActivity 新推荐
func RemoveActivity(c echo.Context) error {
	url := c.QueryParam("url")
	activity := cpi.RemoveActivity(url)
	return c.JSON(http.StatusOK, activity)
}

//GetActivities 获取100条推荐
func GetActivities(c echo.Context) error {
	var links = []Link{}

	var req = c.Request()

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
			// level = user.Level
			if strings.Contains(req.Referer(), cf.ReaderMinApp.AppID) { // VIP稳定通道 笔趣阁Pro，必须邀请用户才能访问，才有推荐。
				level = user.Level

				level = 3
				if level > 3 { // 暂时过滤动漫
				}
			}
			if level < 2 {
				level = 2
			}
			if user.Level <= 2 && user.LoginTotal >= 10 {
				user.Level = 3
			}
		}

	}
	if strings.Contains(req.Referer(), `wx359657b0849ee636`) { // 驴友记
		// level = 3
	}
	var rows = cpi.GetActivities()
	if len(rows) > 0 {
		var rp = map[string]int{}
		var itemlevel int32
		for _, v := range rows {
			// 只显示拥有权限的级别
			itemlevel = reader.GetPathLevel(v.WxTo)
			if v.Level > 3 {
				itemlevel = v.Level
			}
			// 过滤掉做审核的内容
			if level > 2 && itemlevel == 1 {
				itemlevel = 0
			}

			var linkType = `link`
			var appid = ``
			if strings.Contains(req.Referer(), `wxe70eee58e64c7ac7`) { //sodu 去 新推荐阅读
				// 不是新推荐阅读的，全部推荐跳新推荐阅读去
				// linkType = `jumpapp`
				// appid = `wx8ffa5a58c0bb3589`
			}

			if level > itemlevel && itemlevel > 0 {
				//  过滤掉相同 title 的资源（重复的只显示最新一个）
				if _, ok := rp[v.Title]; !ok {
					rp[v.Title] = 1
					links = append(links,
						Link{
							Title: v.Title,
							Icon:  ``,       // cuIcon-new
							Type:  linkType, //  link
							Image: ``,
							WxTo:  v.WxTo,
							Style: `arrow`,
							Appid: appid, //推荐阅读内容全部跳走
						})
				}

			}
		}
	}

	// if strings.Contains( req.Referer(), `wxe70eee58e64c7ac7` ){ //sodu 去 新推荐阅读
	// 	// 不是新推荐阅读的，全部推荐跳新推荐阅读去
	// 	var links2 = []Link{}
	// 	links = append(links2,
	// 	Link{
	// 		Title: `推荐数据正在维护中`,
	// 		Icon:  ``, // cuIcon-new
	// 		Type:  ``,
	// 		Image: ``,
	// 		WxTo:  ``,
	// 		Style: ``,
	// 	})
	// 	return c.JSON(http.StatusOK, links)
	// }

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
