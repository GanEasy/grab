package api

import (
	"fmt"
	"net/http"
	"net/url"

	cpi "github.com/GanEasy/grab/core"
	"github.com/GanEasy/grab/db"
	"github.com/GanEasy/grab/reader"
	"github.com/labstack/echo"
)

// SyncPosts 添加源
func SyncPosts(list reader.Catalog, cate int32) {
	if len(list.Cards) > 0 {
		for _, v := range list.Cards {
			// fmt.Println(`tv`, v, list.Title)
			cpi.SyncPost(v.Title, v.WxTo, v.From, cate)
		}
	}
}

// SearchPosts 搜索资源
func SearchPosts(c echo.Context) error {
	var catelog reader.Catalog
	name := c.QueryParam("name")
	provider := c.QueryParam("provider")
	version := c.QueryParam("version")
	// openID := getOpenID(c)
	// if openID == `` {
	// 	return c.HTML(http.StatusOK, "openid empty")
	// }

	var level = 5 // 4已经支持所有了(小说和漫画) 3支持小说，2什么都不支持
	if provider == `weixin` {
		level = 4
		cerr := cpi.MSGSecCHECK(name)
		if cerr != nil { //&& cerr.Message == `87014`
			catelog.Title = fmt.Sprintf(`暂不支持该关键字搜索`)
			return c.JSON(http.StatusOK, catelog)
		}

	} else if provider == `qq` {
		level = 2
	} else if provider == `web` {
		level = 4
	}
	catelog.Title = fmt.Sprintf(`%v - 搜索结果`, name)
	// fmt.Println(`Title`, catelog.Title)
	// user, _ := getUser(openID)
	cf := cpi.GetConf()
	var posts []db.Post
	if version != `` && version == cf.Search.DevVersion { // 开启严格检查 || 审核版本
		// posts = cpi.GetPostsByNameLimitLevel(name, 2)
		posts = cpi.GetPostsByNameLimitLevel(name, level)
	} else if cf.Search.LimitLevel {
		posts = cpi.GetPostsByNameLimitLevel(name, level)
		// posts = cpi.GetPostsByNameLimitLevel(name, int(user.Level))
	} else {
		posts = cpi.GetPostsByName(name)
	}
	var intro string
	if len(posts) > 0 {
		var itemlevel int32
		for _, v := range posts {

			itemlevel = reader.GetPathLevel(v.WxTo)
			if level >= int(itemlevel) {
				// intro :=
				link, err := url.Parse(v.From)

				if err == nil && link.Host != "" {
					intro = link.Host
				} else {
					intro = ``
				}

				catelog.Cards = append(
					catelog.Cards,
					reader.Card{
						Title:  v.Name,
						WxTo:   v.WxTo,
						Intro:  intro,
						Type:   `card`,
						Cover:  ``,
						Images: nil,
						From:   v.From,
					})
			}
		}
	}
	return c.JSON(http.StatusOK, catelog)
}
