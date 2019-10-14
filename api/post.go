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
	// openID := getOpenID(c)
	// if openID == `` {
	// 	return c.HTML(http.StatusOK, "openid empty")
	// }
	provider := c.QueryParam("provider")

	var level = 5 // 4已经支持所有了 3支持小说，2什么都不支持
	if provider == `weixin` {
		level = 4
	} else if provider == `qq` {
		level = 2
	} else if provider == `web` {
		level = 5
	}
	catelog.Title = fmt.Sprintf(`%v - 搜索结果`, name)
	// fmt.Println(`Title`, catelog.Title)
	// user, _ := getUser(openID)
	cf := cpi.GetConf()
	var posts []db.Post
	if cf.Search.LimitLevel { // 开启严格检查
		posts = cpi.GetPostsByNameLimitLevel(name, level)
		// posts = cpi.GetPostsByNameLimitLevel(name, int(user.Level))
	} else {
		posts = cpi.GetPostsByName(name)
	}
	var intro string
	if len(posts) > 0 {
		for _, v := range posts {
			// if user.Level >= v.Level {
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
			// }
		}
	}
	return c.JSON(http.StatusOK, catelog)
}
