package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

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
	// provider := c.QueryParam("provider")
	version := c.QueryParam("version")

	var level = 5 // 4已经支持所有了(小说和漫画) 3支持小说，2什么都不支持

	cerr := cpi.MSGSecCHECK(name)
	if cerr != nil { //&& cerr.Message == `87014`
		catelog.Title = fmt.Sprintf(`暂不支持该关键字搜索`)
		return c.JSON(http.StatusOK, catelog)
	}

	catelog.Title = fmt.Sprintf(`%v - 搜索结果`, name)

	// 对受限制的应用进行过滤
	cf := cpi.GetConf()
	if cf.Search.LimitInvitation {
		openID := getOpenID(c)
		if openID == `` {
			return c.JSON(http.StatusOK, catelog)
		}
		user, _ := getUser(openID)
		level = int(user.Level)

		if cf.Search.InvitationCode != `` && name != `` && name == cf.Search.InvitationCode { // 输入邀请密令，解锁
			user.Level = 5
			user.Save()

			catelog.Cards = append(
				catelog.Cards,
				reader.Card{
					Title:  `成功解锁`,
					WxTo:   ``,
					Intro:  `请重新加载小程序！`,
					Type:   `card`,
					Cover:  ``,
					Images: nil,
					From:   `admin`,
				})
				return c.JSON(http.StatusOK, catelog)
		}
		
		if  name != `` && name == `000000` { // 固定输入6个0加锁
			user.Level = 1
			user.Save()

			catelog.Cards = append(
				catelog.Cards,
				reader.Card{
					Title:  `成功上锁`,
					WxTo:   ``,
					Intro:  `请重新加载小程序！`,
					Type:   `card`,
					Cover:  ``,
					Images: nil,
					From:   `admin`,
				})
				return c.JSON(http.StatusOK, catelog)
		}

		if user.Level<=2 && user.LoginTotal >= 10 {
			user.Level = 3
		}
	}
	


	// 临时
	var req = c.Request()
	
	
	if strings.Contains( req.Referer(),  cf.ReaderMinAppTwo.AppID ) { //  && strings.Contains( req.Referer(),  cf.ReaderMinAppThree.AppID )
		version = `1.0.0` // 暂时给搜索搜书大师和笔趣阁在线
		level = 4
	}

	var posts []db.Post
	if version != `` && version == cf.Search.DevVersion { // 开启严格检查 || 审核版本
		posts = cpi.GetPostsByNameLimitLevel(name, 2)
		// posts = cpi.GetPostsByNameLimitLevel(name, level)
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

	if len(catelog.Cards) == 0 {
		//

	}
	// if version != cf.Search.DevVersion && level <3 {
	// 	catelog.Cards = append(
	// 		catelog.Cards,
	// 		reader.Card{
	// 			Title:  `╅╅╅>>本程序暂不支持搜索，立即前往分支搜索<<╅╆╆`,
	// 			WxTo:   `/pages/search`,
	// 			Type:   `jumpapp`,
	// 			Appid: `wx7543142ce921d8e3`,
	// 		})
	// }
	if version != cf.Search.DevVersion && level > 2 {

		catelog.Cards = append(
			catelog.Cards,
			reader.Card{
				Title:  `╅╅╅︺ 找不到想要的？您可以 ︺╅╆╆`,
				WxTo:   ``,
				Intro:  ``,
				Type:   `link`,
				Cover:  ``,
				Images: nil,
				From:   ``,
			})

		catelog.Cards = append(
			catelog.Cards,
			reader.Card{
				Title:  `1. 使用第三方平台搜索`,
				WxTo:   `/pages/transfer?action=allsearchdrives&url=&drive=` + name,
				Intro:  `>>点击前往第三方平台搜索“` + name + `<<`,
				Type:   `card`,
				Cover:  ``,
				Images: nil,
				From:   ``,
			})

		catelog.Cards = append(
			catelog.Cards,
			reader.Card{
				Title:  `2. 更换搜索关键字`,
				WxTo:   ``,
				Intro:  `请使用书名搜索，宁可少字也不要错字。例：输入“三生三”搜索“三生三世十里桃花”`,
				Type:   `card`,
				Cover:  ``,
				Images: nil,
				From:   ``,
			})

		
		catelog.Cards = append(
			catelog.Cards,
			reader.Card{
				Title:  `3. 联系我们获取帮助`,
				WxTo:   ``, //reader.EncodeURL(name),
				Intro:  `在我的>“在线客服”或“问题反馈”联系我们。`,
				Type:   `card`,
				Cover:  ``,
				Images: nil,
				From:   ``,
			})
		catelog.Cards = append(
			catelog.Cards,
			reader.Card{
				Title:  `请直接告诉我们，您遇到什么问题，需要我们做什么。`,
				WxTo:   ``,
				Intro:  ``,
				Type:   `link`,
				Cover:  ``,
				Images: nil,
				From:   ``,
			})
		catelog.Cards = append(
			catelog.Cards,
			reader.Card{
				Title:  `请不要问“在吗？”，因为我们答“在呢。”`,
				WxTo:   ``,
				Intro:  ``,
				Type:   `link`,
				Cover:  ``,
				Images: nil,
				From:   ``,
			})
	}
	return c.JSON(http.StatusOK, catelog)
}

// SearchMoreAction 更多搜索方法
func SearchMoreAction(c echo.Context) error {

	var catelog reader.Catalog
	name := c.QueryParam("drive") // 注： 小程序页面 pages/transfer 无法将name参数传上来

	catelog.Title = fmt.Sprintf(`更多“%v”搜索结果`, name)

	
	cf := cpi.GetConf()
	var req = c.Request()

	// 对受限制的应用进行过滤
	if cf.Search.LimitInvitation {
		// 先兼容一下，不过滤二号和三号小程序搜索功能先
		if !strings.Contains( req.Referer(),  cf.ReaderMinAppTwo.AppID ) && !strings.Contains( req.Referer(),  cf.ReaderMinAppThree.AppID ) {
			openID := getOpenID(c)
			if openID == `` {
				return c.JSON(http.StatusOK, catelog)
			}
			user, _ := getUser(openID)
			if user.LoginTotal < 10 || user.Level < 1 { //限制用户不返回搜索结果
				return c.JSON(http.StatusOK, catelog)
			}
		}
	}
	

	// catelog.Title = `更多相关搜索结果`

	catelog.Cards = append(
		catelog.Cards,
		reader.Card{
			Title:  `起点小说网 搜索“` + name + `”`,
			WxTo:   `/pages/searchmore?drive=qidian&name=` + name,
			Intro:  ``,
			Type:   `link`,
			Cover:  ``,
			Images: nil,
			From:   ``,
		})
	catelog.Cards = append(
		catelog.Cards,
		reader.Card{
			Title:  `在纵横小说网 搜索“` + name + `”`,
			WxTo:   `/pages/searchmore?drive=zongheng&name=` + name,
			Intro:  ``,
			Type:   `link`,
			Cover:  ``,
			Images: nil,
			From:   ``,
		})
	catelog.Cards = append(
		catelog.Cards,
		reader.Card{
			Title:  `在17K文学 搜索“` + name + `”`,
			WxTo:   `/pages/searchmore?drive=17k&name=` + name,
			Intro:  ``,
			Type:   `link`,
			Cover:  ``,
			Images: nil,
			From:   ``,
		})
	catelog.Cards = append(
		catelog.Cards,
		reader.Card{
			Title:  `在潇湘书院 搜索“` + name + `”`,
			WxTo:   `/pages/searchmore?drive=xxsy&name=` + name,
			Intro:  ``,
			Type:   `link`,
			Cover:  ``,
			Images: nil,
			From:   ``,
		})
	catelog.Cards = append(
		catelog.Cards,
		reader.Card{
			Title:  `在笔趣阁jxla 搜索“` + name + `”`,
			WxTo:   `/pages/searchmore?drive=jx&name=` + name,
			Intro:  ``,
			Type:   `link`,
			Cover:  ``,
			Images: nil,
			From:   ``,
		})
	catelog.Cards = append(
		catelog.Cards,
		reader.Card{
			Title:  `在笔趣阁paoshu8 搜索“` + name + `”`,
			WxTo:   `/pages/searchmore?drive=paoshu8&name=` + name,
			Intro:  ``,
			Type:   `link`,
			Cover:  ``,
			Images: nil,
			From:   ``,
		})
	catelog.Cards = append(
		catelog.Cards,
		reader.Card{
			Title:  `在笔趣阁mcmssc 搜索“` + name + `”`,
			WxTo:   `/pages/searchmore?drive=mcmssc&name=` + name,
			Intro:  ``,
			Type:   `link`,
			Cover:  ``,
			Images: nil,
			From:   ``,
		})
	catelog.Cards = append(
		catelog.Cards,
		reader.Card{
			Title:  `在顶点小说booktxt 搜索“` + name + `”`,
			WxTo:   `/pages/searchmore?drive=booktxt&name=` + name,
			Intro:  ``,
			Type:   `link`,
			Cover:  ``,
			Images: nil,
			From:   ``,
		})
	catelog.Cards = append(
		catelog.Cards,
		reader.Card{
			Title:  `在顶点小说280xs 搜索“` + name + `”`,
			WxTo:   `/pages/searchmore?drive=xs280&name=` + name,
			Intro:  ``,
			Type:   `link`,
			Cover:  ``,
			Images: nil,
			From:   ``,
		})

	catelog.Cards = append(
		catelog.Cards,
		reader.Card{
			Title:  `在书阁小说shuge 搜索“` + name + `”`,
			WxTo:   `/pages/searchmore?drive=shuge&name=` + name,
			Intro:  ``,
			Type:   `link`,
			Cover:  ``,
			Images: nil,
			From:   ``,
		})

	catelog.Cards = append(
		catelog.Cards,
		reader.Card{
			Title:  `笔下看书阁jininggeyin 搜索“` + name + `”`,
			WxTo:   `/pages/searchmore?drive=bxks&name=` + name,
			Intro:  ``,
			Type:   `link`,
			Cover:  ``,
			Images: nil,
			From:   ``,
		})

	return c.JSON(http.StatusOK, catelog)
}
