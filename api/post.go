package api

import (
	"fmt"
	"net/http"

	cpi "github.com/GanEasy/grab/core"
	"github.com/GanEasy/grab/reader"
	"github.com/labstack/echo"
)

// SyncPosts 添加源
func SyncPosts(list reader.Catalog) {
	if len(list.Cards) > 0 {
		for _, v := range list.Cards {
			// fmt.Println(`tv`, v, list.Title)
			cpi.SyncPost(v.Title, v.WxTo, v.From, 1)
		}
	}
}

// SelectPosts 搜索资源
func SelectPosts(c echo.Context) error {
	var catelog reader.Catalog
	name := c.QueryParam("name")
	openID := getOpenID(c)
	if openID == `` {
		return c.HTML(http.StatusOK, "openid empty")
	}
	catelog.Title = fmt.Sprintf(`%v - 搜索结果`, name)
	// fmt.Println(`Title`, catelog.Title)
	user, _ := getUser(openID)
	posts := cpi.GetPostsByName(name)

	if len(posts) > 0 {
		for _, v := range posts {
			if user.Level >= v.Level {
				catelog.Cards = append(
					catelog.Cards,
					reader.Card{
						v.Name,
						v.WxTo,
						``,
						`link`,
						``,
						nil,
						v.From,
					})
			}
		}
	}
	return c.JSON(http.StatusOK, catelog)
}
