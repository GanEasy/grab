package api

import (
	"net/http"
	"strings"

	"github.com/GanEasy/grab"
	cpi "github.com/GanEasy/grab/core"
	"github.com/GanEasy/grab/reader"
	"github.com/labstack/echo"
)

// Carousel 小程序首页轮播内容(作为专题广告或其它的东西使用)
type Carousel struct {
	URL    string `json:"url"`
	Type   string `json:"type"` // 期望可以同时支持视频播放(虽然很不现实)
	WxTo   string `json:"wxto"` // 点击后跳转地址
	Event  string `json:"event"`
	Poster string `json:"poster"` //type 为 poster 时生效，打开一个图片
	AppID  string `json:"appid"`  //type 为 jumpapp 时生效，跳转到另一个小程序
}

// GetCarousels 获取首页走马灯数据
func GetCarousels(c echo.Context) error {
	var carousels []Carousel
	// carousels = append(
	// 	carousels,
	// Carousel{
	// 	URL:    `https://aireadhelper.github.io/static/images/mfqz.jpg`,
	// 	Type:   `image`,
	// 	Event:  `poster`,
	// 	WxTo:   ``,
	// 	Poster: `https://aireadhelper.github.io/static/images/qrcode.png`,
	// })
	// carousels = append(
	// 	carousels,
	// 	Carousel{
	// 		URL:    `https://aireadhelper.github.io/static/images/boyslist.jpg`,
	// 		Type:   `image`,
	// 		Event:  `link`,
	// 		WxTo:   `/pages/list?drive=qidian&url=` + grab.EncodeURL(`https://www.qidian.com/all?orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`),
	// 		Poster: ``,
	// 	})
	// carousels = append(
	// 	carousels,
	// 	Carousel{
	// 		URL:    `https://aireadhelper.github.io/static/images/girlslist.jpg`,
	// 		Type:   `image`,
	// 		Event:  `link`,
	// 		WxTo:   `/pages/list?drive=xxsy&url=` + grab.EncodeURL(`https://www.xxsy.net/search?s_wd=&channel=2&sort=9&pn=1`),
	// 		Poster: ``,
	// 	})

	// carousels = append(
	// 	carousels,
	// 	Carousel{
	// 		URL:  `https://ossweb-img.qq.com/images/lol/web201310/skin/big84000.jpg`,
	// 		Type: `image`,
	// 		WxTo: ``,
	// 	})

	// carousels = append(
	// 	carousels,
	// 	Carousel{
	// 		URL:  `https://ossweb-img.qq.com/images/lol/web201310/skin/big37006.jpg`,
	// 		Type: `image`,
	// 		WxTo: ``,
	// 	})

	// carousels = append(
	// 	carousels,
	// 	Carousel{
	// 		URL:    `https://luck.wechatrank.com/images/adyyy.jpg`,
	// 		Type:   `image`,
	// 		Event:  `poster`,
	// 		WxTo:   ``,
	// 		Poster: `https://luck.wechatrank.com/images/adyyy.jpg`,
	// 	})

	return c.JSON(http.StatusOK, carousels)
}

// Link 小程序首页、用户页、看单页结构
type Link struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
	Type  string `json:"type"` // 打开页面，展示海报图片
	Image string `json:"image"`
	WxTo  string `json:"wxto"`  // 点击后跳转地址
	Style string `json:"style"` //特定样式
	Appid string `json:"appid"` //跳转小程序
}

// GetUserLinks 获取用户列表内容
func GetUserLinks(c echo.Context) error {

	cf := cpi.GetConf()
	version := c.QueryParam("version")

	if cf.Search.LimitLevel || version == cf.Search.DevVersion { // 开启严格检查

		return c.JSON(http.StatusOK, []Link{
			Link{
				Title: `浏览记录`,
				Icon:  `cuIcon-time`,
				Type:  `link`,
				Image: ``,
				WxTo:  `/pages/logs`,
				Style: `arrow`,
			},
		})
	}

	var links = []Link{
		Link{
			Title: `创建源`,
			Icon:  `cuIcon-new`,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/newCatalog`,
			Style: `arrow`,
		},
		Link{
			Title: `浏览记录`,
			Icon:  `cuIcon-time`,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/logs`,
			Style: `arrow`,
		},

		// Link{
		// 	Title: `主线路(笔趣阁plus)`,
		// 	Icon:  `cuIcon-move`,
		// 	Type:  `image`,
		// 	Image: `https://iblog.wechatrank.com/images/r1.jpg`,
		// 	WxTo:  ``,
		// 	Style: ``,
		// },
		// Link{
		// 	Title: `备用线路(无广告)`,
		// 	Icon:  `cuIcon-move`,
		// 	Type:  `image`,
		// 	Image: `https://iblog.wechatrank.com/images/r2.jpg`,
		// 	WxTo:  ``,
		// 	Style: ``,
		// },

		// Link{
		// 	Title: `广告策略与用户组`,
		// 	Icon:  `cuIcon-discover`,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/article?drive=blog&url=` + grab.EncodeURL(`https://aireadhelper.github.io/doc/v4/ads.html`),
		// 	Style: `arrow text-red`,
		// },

		// Link{
		// 	Title: `使用介绍`,
		// 	Icon:  `cuIcon-question`,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/article?drive=blog&url=` + grab.EncodeURL(`https://aireadhelper.github.io/doc/v2/about.html`),
		// 	Style: `arrow`,
		// },
		// Link{
		// 	Title: `免责声明`,
		// 	Icon:  `cuIcon-lightforbid`,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/article?drive=blog&url=` + grab.EncodeURL(`https://aireadhelper.github.io/doc/v2/exemption.html`),
		// 	Style: `arrow`,
		// },
		// Link{
		// 	Title: `交流群反馈问题`,
		// 	Icon:  `cuIcon-group`,
		// 	Type:  `image`,
		// 	Image: `https://aireadhelper.github.io/static/images/group.png`,
		// 	WxTo:  ``,
		// 	Style: ``,
		// },
		// Link{
		// 	Title: `我的推荐`,
		// 	Icon:  `cuIcon-activity`,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/activities`,
		// 	Style: `arrow`,
		// },
		// Link{
		// 	Title: `使用说明`,
		// 	Icon:  `cuIcon-question`,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/newCreate`,
		// 	Style: `arrow`,
		// },
		// Link{
		// 	Title: `免责声明`,
		// 	Icon:  `cuIcon-command`,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/newCreate`,
		// 	Style: `arrow`,
		// },

		// Link{
		// 	Title: `用户协议`,
		// 	Icon:  `cuIcon-squarecheck`,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/newCreate`,
		// 	Style: `arrow`,
		// },
		// Link{
		// 	Title: `公告信息`,
		// 	Icon:  `cuIcon-notification`,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/newCreate`,
		// 	Style: `arrow`,
		// },

		// Link{
		// 	Title: `交流群`,
		// 	Icon:  `cuIcon-group`,
		// 	Type:  `image`,
		// 	Image: `https://aireadhelper.github.io/static/images/group.png`,
		// 	WxTo:  ``,
		// 	Style: ``,
		// },
	}
	return c.JSON(http.StatusOK, links)
}

// GetNewCatelogLinks 获取用户列表内容 newcateloghelps
func GetNewCatelogLinks(c echo.Context) error {
	var links = []Link{
		Link{
			Title: `如何添加转码数据源`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/article?drive=blog&url=` + grab.EncodeURL(`https://aireadhelper.github.io/doc/v2/newguide.html`),
			Style: `arrow`,
		},
		Link{
			Title: `免责声明`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/article?drive=blog&url=` + grab.EncodeURL(`https://aireadhelper.github.io/doc/v2/exemption.html`),
			Style: `arrow`,
		},

		// Link{
		// 	Title: `加入交流群获得帮助`,
		// 	Icon:  `cuIcon-group`,
		// 	Type:  `image`,
		// 	Image: `https://aireadhelper.github.io/static/images/group.png`,
		// 	WxTo:  ``,
		// 	Style: ``,
		// },
		// Link{
		// 	Title: `加入交流群获得帮助`,
		// 	Icon:  ``,
		// 	Type:  `image`,
		// 	Image: `https://ossweb-img.qq.com/images/lol/web201310/skin/big37006.jpg`,
		// 	WxTo:  ``,
		// 	Style: `text-red`,
		// },
	}
	return c.JSON(http.StatusOK, links)
}

// GetExploreLinks 获取首页(广场)列表内容
func GetExploreLinks(c echo.Context) error {

	openID := getOpenID(c)
	if openID == `` {
		return c.HTML(http.StatusOK, "wxto empty")
	}
	user, _ := getUser(openID)
	// fmt.Println(user)
	if user.LoginTotal >= 10 { //老用户，使用超过20次了
		// return c.JSON(http.StatusOK, GetGuideExploreLinksAndJumpVipApp())
	}

	cf := cpi.GetConf()

	version := c.QueryParam("version")

	var req = c.Request()

	if cf.Search.LimitLevel || version == cf.Search.DevVersion { // 开启严格检查

		return c.JSON(http.StatusOK, GetWaitExamineExplore())
	}

	//
	if strings.Contains(req.Referer(), `wx72a2a482881fd47a`) && strings.Contains(req.Header.Get("User-Agent"), `mpcrawler`) {
		// return c.JSON(http.StatusOK, GetGuideExploreJumpLinks())
		// return c.JSON(http.StatusOK, SingleMenu())
		return c.JSON(http.StatusOK, GetGuideExploreLinks())

	}

	if strings.Contains(req.Referer(), `wx72a2a482881fd47a`) && !strings.Contains(req.Header.Get("User-Agent"), `mpcrawler`) { // plus
		// return c.JSON(http.StatusOK, GetGuideExploreJumpLinks())
		return c.JSON(http.StatusOK, GetGuideExploreLinks())
		// return c.JSON(http.StatusOK, SingleMenu())
	}

	// if strings.Contains(req.Header.Get("User-Agent"), `mpcrawler`) { // plus
	// 	// return c.JSON(http.StatusOK, GetGuideExploreJumpLinks())
	// 	return c.JSON(http.StatusOK, GetPublishExploreLinks())
	// }
	return c.JSON(http.StatusOK, SingleMenu()) // 2019年12月26日 09:02:19 放到列表试试
	// return c.JSON(http.StatusOK, GetGuideExploreLinks())
}

// SingleMenu 伪装给单一入口
func SingleMenu() []Link {

	var links = []Link{
		Link{
			Title: `免费小说`,
			Icon:  ``,
			Type:  `jumpapp`,
			Image: ``,
			WxTo:  `/pages/transfer?action=allbookroesoures&drive=&url=`,
			// WxTo:  `/pages/transfer?action=alllearnresources&drive=&url=`, // allbookroesoures
			Style: `arrow`,
			Appid: `wx96830e80b331c267`, // wx7543142ce921d8e3
		},

		Link{
			Title: `免责声明`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/article?drive=blog&url=` + grab.EncodeURL(`https://aireadhelper.github.io/doc/v2/exemption.html`),
			Style: ``,
		},
	}
	return links

}

// SafeMenu 霸道总裁专题小说 推荐页
func SafeMenu() []Link {

	var links = []Link{
		Link{
			Title: `发现更多免费小说资源`,
			Icon:  ``,
			Type:  `jumpapp`,
			Image: ``,
			WxTo:  `/pages/index`,
			Style: ``,
			Appid: `wx7543142ce921d8e3`,
		},

		Link{
			Title: `免责声明`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/article?drive=blog&url=` + grab.EncodeURL(`https://aireadhelper.github.io/doc/v2/exemption.html`),
			Style: `arrow`,
		},
	}
	return links

}

// GetWaitExamineExplore 用于审核的内容列表
func GetWaitExamineExplore() []Link {

	var links = []Link{

		Link{
			Title: `全部编程学习资料`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/transfer?action=alllearnresources&drive=&url=`,
			Style: `arrow`,
		},
		Link{
			Title: `微信小程序开发入门系列教程`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/catalog?drive=blog&url=` + grab.EncodeURL(`https://xueyuanjun.com/laravel-from-appreciate-to-artisan`),
			Style: `arrow`,
		},

		Link{
			Title: `从学徒到工匠精校版`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/catalog?drive=blog&url=` + grab.EncodeURL(`https://xueyuanjun.com/wechat-miniprogram-tutorial`),
			Style: `arrow`,
		},

		Link{
			Title: `从入门到精通系列教程`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/catalog?drive=blog&url=` + grab.EncodeURL(`https://xueyuanjun.com/laravel-tutorial-5_7`),
			Style: `arrow`,
		},

		Link{
			Title: `学习心得`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/catalog?drive=github&url=` + grab.EncodeURL(`https://github.com/aireadhelper/aireadhelper.github.io/blob/master/edudoc/index.md`),
			Style: `arrow`,
		},
		// Link{
		// 	Title: `免责声明`,
		// 	Icon:  ``,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/article?drive=blog&url=` + grab.EncodeURL(`https://aireadhelper.github.io/doc/v2/exemption.html`),
		// 	Style: `arrow`,
		// },
	}
	return links

}

// GetWaitExamineExploreWx2 微信小程序临时应付一下
func GetWaitExamineExploreWx2() []Link {

	var links = []Link{

		Link{
			Title: `编程学习资料`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/transfer?action=alllearnresources&drive=&url=`,
			Style: `arrow`,
		},

		Link{
			Title: `使用教程`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/article?drive=blog&url=` + grab.EncodeURL(`https://aireadhelper.github.io/doc/guide.html?`),
			Style: `arrow`,
		},

		Link{
			Title: `关于`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/article?drive=blog&url=` + grab.EncodeURL(`https://aireadhelper.github.io/doc/about.html`),
			Style: `arrow`,
		},

		Link{
			Title: `帮助`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/article?drive=blog&url=` + grab.EncodeURL(`https://aireadhelper.github.io/doc/help.html`),
			Style: `arrow`,
		},
		// Link{
		// 	Title: `从学徒到工匠精校版`,
		// 	Icon:  ``,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/catalog?drive=blog&url=` + grab.EncodeURL(`https://xueyuanjun.com/wechat-miniprogram-tutorial`),
		// 	Style: `arrow`,
		// },

		// Link{
		// 	Title: `从入门到精通系列教程`,
		// 	Icon:  ``,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/catalog?drive=blog&url=` + grab.EncodeURL(`https://xueyuanjun.com/laravel-tutorial-5_7`),
		// 	Style: `arrow`,
		// },
		Link{
			Title: `Go语言入门教程`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/catalog?drive=blog&url=` + grab.EncodeURL(`https://xueyuanjun.com/golang-tutorials`),
			Style: `arrow`,
		},
		// Link{
		// 	Title: `德哥博客-最佳实践`,
		// 	Icon:  ``,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/catalog?drive=github&url=` + grab.EncodeURL(`https://github.com/digoal/blog/blob/master/class/24.md`),
		// 	Style: `arrow`,
		// },

		// Link{
		// 	Title: `德哥博客-经典案例`,
		// 	Icon:  ``,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/catalog?drive=github&url=` + grab.EncodeURL(`https://github.com/digoal/blog/blob/master/class/15.md`),
		// 	Style: `arrow`,
		// },

		Link{
			Title: `Laravel 项目开发规范`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/catalog?drive=learnku&url=` + grab.EncodeURL(`https://learnku.com/docs/laravel-specification/5.5`),
			Style: `arrow`,
		},
		// Link{
		// 	Title: `Laravel5.5开发文档`,
		// 	Icon:  ``,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/catalog?drive=learnku&url=` + grab.EncodeURL(`https://learnku.com/docs/laravel-specification/5.5`),
		// 	Style: `arrow`,
		// },
		// Link{
		// 	Title: `Laravel 5.5 中文文档`,
		// 	Icon:  ``,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/catalog?drive=learnku&url=` + grab.EncodeURL(`https://learnku.com/docs/laravel/5.5`),
		// 	Style: `arrow`,
		// },
		// Link{
		// 	Title: `Dingo API 2.0.0 中文文档`,
		// 	Icon:  ``,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/catalog?drive=learnku&url=` + grab.EncodeURL(`https://learnku.com/docs/dingo-api/2.0.0`),
		// 	Style: `arrow`,
		// },

		Link{
			Title: `所有资源`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/transfer?action=allroesoures&drive=&url=`,
			Style: `arrow`,
		},
	}
	return links

}

//GetPublishExploreLinks 获取公开发布的内容
func GetPublishExploreLinks() []Link {
	var links = []Link{

		Link{
			Title: `小说资源`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/transfer?action=allbookroesoures&drive=&url=`,
			Style: `arrow`,
		},
		Link{
			Title: `漫画资源`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/transfer?action=allcartoonroesoures&drive=&url=`,
			Style: `arrow`,
		},
		Link{
			Title: `编程学习资料`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/transfer?action=alllearnresources&drive=&url=`,
			Style: `arrow`,
		},

		Link{
			Title: `使用介绍`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/article?drive=blog&url=` + grab.EncodeURL(`https://aireadhelper.github.io/doc/v2/about.html`),
			Style: `arrow`,
		},

		Link{
			Title: `免责声明`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/article?drive=blog&url=` + grab.EncodeURL(`https://aireadhelper.github.io/doc/v2/exemption.html`),
			Style: `arrow`,
		},

		// Link{
		// 	Title: `阅读交流群`,
		// 	Icon:  `cuIcon-group`,
		// 	Type:  `image`,
		// 	Image: `https://aireadhelper.github.io/static/images/group.png`,
		// 	WxTo:  ``,
		// 	Style: ``,
		// },
	}

	return links
}

//GetGuideExploreLinks  新版，引导转化
func GetGuideExploreLinks() []Link {
	var links = []Link{
		Link{
			Title: `编程学习资料`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/transfer?action=alllearnresources&drive=&url=`,
			Style: `arrow`,
		},
		// Link{
		// 	Title: `小说资源`,
		// 	Icon:  ``,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/transfer?action=allbookroesoures&drive=&url=`,
		// 	Style: `arrow`,
		// },
		// Link{
		// 	Title: `漫画资源`,
		// 	Icon:  ``,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/transfer?action=allcartoonroesoures&drive=&url=`,
		// 	Style: `arrow`,
		// },
		// Link{
		// 	Title: `广告策略与用户组`,
		// 	Icon:  ``,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/article?drive=blog&url=` + grab.EncodeURL(`https://aireadhelper.github.io/doc/v4/ads.html`),
		// 	Style: `arrow text-red`,
		// },

		//

		// Link{
		// 	Title: `评书、推书、同人作品征稿`,
		// 	Icon:  ``,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/article?drive=blog&url=` + grab.EncodeURL(`https://mp.weixin.qq.com/s/RszJ-ibBDR9Np7GJkRKyjQ`),
		// 	Style: `arrow`,
		// },
		// Link{
		// 	Title: `请同学们以学业、工作、生活为重!`,
		// 	Type:  `link`,
		// 	WxTo:  ``,
		// 	Style: ``,
		// },
		// Link{
		// 	Title: `备用线路(无广告)`,
		// 	Type:  `image`,
		// 	Image: `https://iblog.wechatrank.com/images/r2.jpg`,
		// 	WxTo:  ``,
		// 	Style: ``,
		// },
		// Link{
		// 	Title: `-----------------------------`,
		// 	Type:  `link`,
		// 	WxTo:  ``,
		// 	Style: ``,
		// },

		Link{
			Title: `百度热搜榜`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=baidu&url=` + grab.EncodeURL(`https://www.baidu.com`),
			Style: `arrow`,
		},

		Link{
			Title: `起点小说网`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=qidian&url=` + grab.EncodeURL(`https://www.qidian.com`),
			Style: `arrow`,
		},
		// Link{
		// 	Title: `纵横小说网`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/categories?drive=zongheng&url=` + grab.EncodeURL(`http://book.zongheng.com`),
		// 	Style: `arrow`,
		// },
		Link{
			Title: `17K文学`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=17k&url=` + grab.EncodeURL(`http://www.17k.com`),
			Style: `arrow`,
		},
		Link{
			Title: `红袖添香`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=hongxiu&url=` + grab.EncodeURL(`https://www.hongxiu.com`),
			Style: `arrow`,
		},
		Link{
			Title: `潇湘书院`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=xxsy&url=` + grab.EncodeURL(`https://www.xxsy.net`),
			Style: `arrow`,
		},

		// Link{
		// 	Title: `更多小说资源`,
		// 	Icon:  ``,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/transfer?action=allbookroesoures&drive=&url=`,
		// 	Style: `arrow`,
		// },

		Link{
			Title: `笔趣阁jxla`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=jx&url=` + grab.EncodeURL(`https://m.jx.la/`),
			Style: `arrow`,
		},
		Link{
			Title: `笔趣阁mcmssc`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=mcmssc&url=` + grab.EncodeURL(`https://www.mcmssc.com/`),
			Style: `arrow`,
		},

		// Link{
		// 	Title: `笔趣阁paoshu8`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/categories?drive=paoshu8&url=` + grab.EncodeURL(`http://www.paoshu8.com/`),
		// 	Style: `arrow`,
		// },

		Link{
			Title: `顶点小说280xs`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=xs280&url=` + grab.EncodeURL(`https://www.280xs.com/`),
			Style: `arrow`,
		},
		Link{
			Title: `笔趣阁xbiquge`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=xbiquge&url=` + grab.EncodeURL(`http://www.xbiquge.la/`),
			Style: `arrow`,
		},
		// Link{
		// 	Title: `笔趣阁biqugeinfo`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/categories?drive=biqugeinfo&url=` + grab.EncodeURL(`https://m.biquge.info/`),
		// 	Style: `arrow`,
		// },
		Link{
			Title: `顶点小说booktxt`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=booktxt&url=` + grab.EncodeURL(`http://www.booktxt.net`),
			Style: `arrow`,
		},
		Link{
			Title: `书阁小说网shugela`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=shuge&url=` + grab.EncodeURL(`https://m.shuge.la/`),
			Style: `arrow`,
		},

		// Link{
		// 	Title: `笔下看书阁jininggeyin`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/categories?drive=bxks&url=` + grab.EncodeURL(`https://www.jininggeyin.com/`),
		// 	Style: `arrow`,
		// },

		// Link{
		// 	Title: `新18小说网0335jjlm`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/categories?drive=xin18&url=` + grab.EncodeURL(`https://www.0335jjlm.com/`),
		// 	Style: `arrow`,
		// },
		// Link{
		// 	Title: `去看书qkshu6`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/categories?drive=qkshu6&url=` + grab.EncodeURL(`https://www.qkshu6.com/`),
		// 	Style: `arrow`,
		// },
		// Link{
		// 	Title: `笔趣阁soe8`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/categories?drive=soe8&url=` + grab.EncodeURL(`http://m.soe8.com/`),
		// 	Style: `arrow`,
		// },

		// Link{
		// 	Title: `U小说阅读网`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/categories?drive=uxiaoshuo&url=` + grab.EncodeURL(`https://m.uxiaoshuo.com/`),
		// 	Style: `arrow`,
		// },
		// Link{
		// 	Title: `老司机小说`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/categories?drive=laosijixs&url=` + grab.EncodeURL(`http://m.laosijixs.com/`),
		// },

		// Link{
		// 	Title: `漫画资源`,
		// 	Icon:  ``,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/transfer?action=allcartoonroesoures&drive=&url=`,
		// 	Style: `arrow`,
		// },

		// Link{
		// 	Title: `╅╅╅︺未满18岁禁止观看︺╅╆╆`,
		// 	Type:  `link`,
		// 	WxTo:  ``,
		// },

		// // Link{
		// // 	Title: `韩漫窝(18禁)`,
		// // 	Type:  `link`,
		// // 	WxTo:  `/pages/list?action=list&drive=hanmanwo&url=` + grab.EncodeURL(`http://www.hanmanwo.com/booklist`),
		// // },

		// Link{
		// 	Title: `韩漫库(18禁)`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/list?action=list&drive=hanmanku&url=` + grab.EncodeURL(`http://www.hanmanku.com/booklist`),
		// 	Style: `arrow`,
		// },

		// Link{
		// 	Title: `海猫吧(18禁)`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/list?action=list&drive=haimaoba&url=` + grab.EncodeURL(`http://www.haimaoba.com/list/0/`),
		// 	Style: `arrow`,
		// },

		// // reader.Card{
		// // 	Title: `我爱妹子漫画(18禁)`,
		// // 	Type:  `link`,
		// // 	WxTo:  `/pages/list?action=list&drive=aimeizi5&url=` + grab.EncodeURL(`https://5aimeizi.com/booklist`),
		// // },
		// Link{
		// 	Title: `腐漫漫画(18禁)`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/categories?drive=fuman&url=` + grab.EncodeURL(`https://www.5aimeizi.com/`),
		// 	Style: `arrow`,
		// },
		// Link{
		// 	Title: `漫画台(18禁)`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/categories?drive=manhwa&url=` + grab.EncodeURL(`https://www.manhwa.cc/`),
		// 	Style: `arrow`,
		// },
		// Link{
		// 	Title: `看妹子漫画(18禁)`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/list?action=list&drive=kanmeizi&url=` + grab.EncodeURL(`https://www.kanmeizi.cc/booklist`),
		// 	Style: `arrow`,
		// },
		// Link{
		// 	Title: `伟叫兽漫画网(18禁)`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/categories?action=list&drive=weijiaoshou&url=` + grab.EncodeURL(`http://www.weijiaoshou.cn`),
		// 	Style: `arrow`,
		// },
		// Link{
		// 	Title: `漫物语(18禁)`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/categories?drive=manwuyu&url=` + grab.EncodeURL(`http://www.manwuyu.com/`),
		// 	Style: `arrow`,
		// },


		Link{
			Title: `关于&声明`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/article?drive=blog&url=` + grab.EncodeURL(`https://aireadhelper.github.io/doc/v3/about.html`),
			Style: `arrow`,
		},
	}

	return links
}

//GetGuideExploreJumpLinks  改，跳转列表
func GetGuideExploreJumpLinks() []Link {
	var links = []Link{

		Link{
			Title: `起点小说网`,
			Type:  `jumpapp`,
			WxTo:  `/pages/categories?drive=qidian&url=` + grab.EncodeURL(`https://www.qidian.com`),
			Style: `arrow`,
			Appid: `wxa94ddd94358b2d1d`,
		},
		// Link{
		// 	Title: `纵横小说网`,
		// 	Type:  `jumpapp`,
		// 	WxTo:  `/pages/categories?drive=zongheng&url=` + grab.EncodeURL(`http://book.zongheng.com`),
		// 	Style: `arrow`,
		// 	Appid: `wxa94ddd94358b2d1d`,
		// },
		Link{
			Title: `17K文学`,
			Type:  `jumpapp`,
			WxTo:  `/pages/categories?drive=17k&url=` + grab.EncodeURL(`http://www.17k.com`),
			Style: `arrow`,
			Appid: `wxa94ddd94358b2d1d`,
		},
		Link{
			Title: `红袖添香`,
			Type:  `jumpapp`,
			WxTo:  `/pages/categories?drive=hongxiu&url=` + grab.EncodeURL(`https://www.hongxiu.com`),
			Style: `arrow`,
			Appid: `wxa94ddd94358b2d1d`,
		},
		Link{
			Title: `潇湘书院`,
			Type:  `jumpapp`,
			WxTo:  `/pages/categories?drive=xxsy&url=` + grab.EncodeURL(`https://www.xxsy.net`),
			Style: `arrow`,
			Appid: `wxa94ddd94358b2d1d`,
		},

		Link{
			Title: `笔趣阁jxla`,
			Type:  `jumpapp`,
			WxTo:  `/pages/categories?drive=jx&url=` + grab.EncodeURL(`https://m.jx.la/`),
			Style: `arrow`,
			Appid: `wxa94ddd94358b2d1d`,
		},
		Link{
			Title: `笔趣阁mcmssc`,
			Type:  `jumpapp`,
			WxTo:  `/pages/categories?drive=mcmssc&url=` + grab.EncodeURL(`https://www.mcmssc.com/`),
			Style: `arrow`,
			Appid: `wxa94ddd94358b2d1d`,
		},

		Link{
			Title: `笔趣阁paoshu8`,
			Type:  `jumpapp`,
			WxTo:  `/pages/categories?drive=paoshu8&url=` + grab.EncodeURL(`http://www.paoshu8.com/`),
			Style: `arrow`,
			Appid: `wxa94ddd94358b2d1d`,
		},

		Link{
			Title: `顶点小说280xs`,
			Type:  `jumpapp`,
			WxTo:  `/pages/categories?drive=xs280&url=` + grab.EncodeURL(`https://www.280xs.com/`),
			Style: `arrow`,
			Appid: `wxa94ddd94358b2d1d`,
		},
		Link{
			Title: `笔趣阁xbiquge`,
			Type:  `jumpapp`,
			WxTo:  `/pages/categories?drive=xbiquge&url=` + grab.EncodeURL(`http://www.xbiquge.la/`),
			Style: `arrow`,
			Appid: `wxa94ddd94358b2d1d`,
		},

		Link{
			Title: `顶点小说booktxt`,
			Type:  `jumpapp`,
			WxTo:  `/pages/categories?drive=booktxt&url=` + grab.EncodeURL(`http://www.booktxt.net`),
			Style: `arrow`,
			Appid: `wxa94ddd94358b2d1d`,
		},
		Link{
			Title: `书阁小说网shugela`,
			Type:  `jumpapp`,
			WxTo:  `/pages/categories?drive=shuge&url=` + grab.EncodeURL(`https://m.shuge.la/`),
			Style: `arrow`,
			Appid: `wxa94ddd94358b2d1d`,
		},

		Link{
			Title: `编程学习资料`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/transfer?action=alllearnresources&drive=&url=`,
			Style: `arrow`,
		},

		Link{
			Title: `关于&声明`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/article?drive=blog&url=` + grab.EncodeURL(`https://aireadhelper.github.io/doc/v3/about.html`),
			Style: `arrow`,
		},
	}

	return links
}

//GetGuideExploreLinksAndJumpVipApp  新版，引导转化
func GetGuideExploreLinksAndJumpVipApp() []Link {
	if true {
		return GetGuideExploreLinks()
	}
	var links = []Link{
		Link{
			Title: `老铁进VIP通道搜“112233”解锁`,
			Icon:  `notice`,
			Type:  `jumpapp`,
			Image: ``,
			WxTo:  `/pages/index?uid=123`,
			Style: `arrow`,
			Appid: `wxa94ddd94358b2d1d`,
		},
	}
	var links2 = GetGuideExploreLinks()

	return append(links, links2...)
}

//GetAllResources 获取所有资源
func GetAllResources(c echo.Context) error {
	var list = reader.Catalog{}
	list.Title = `全部资源`

	list.SourceURL = ``

	list.Hash = ``

	list.Cards = []reader.Card{

		reader.Card{
			Title: `起点小说网`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=qidian&url=` + grab.EncodeURL(`https://www.qidian.com`),
		},
		reader.Card{
			Title: `纵横小说网`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=zongheng&url=` + grab.EncodeURL(`http://book.zongheng.com`),
		},
		reader.Card{
			Title: `17K文学`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=17k&url=` + grab.EncodeURL(`http://www.17k.com`),
		},
		// reader.Card{
		// 	Title: `笔下文学`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/categories?drive=bxwx&url=` + grab.EncodeURL(`https://www.bxwx.la`),
		// },
		reader.Card{
			Title: `U小说阅读网`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=uxiaoshuo&url=` + grab.EncodeURL(`https://m.uxiaoshuo.com/`),
		},
		// reader.Card{
		// 	Title: `笔趣阁biquyun`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/categories?drive=biquyun&url=` + grab.EncodeURL(`https://m.biquyun.com`),
		// },

		reader.Card{
			Title: `顶点小说booktxt`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=booktxt&url=` + grab.EncodeURL(`http://www.booktxt.net`),
		},

		reader.Card{
			Title: `笔趣阁soe8`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=soe8&url=` + grab.EncodeURL(`http://m.soe8.com/`),
		},

		reader.Card{
			Title: `笔趣阁xbiquge`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=xbiquge&url=` + grab.EncodeURL(`http://www.xbiquge.la/`),
		},

		reader.Card{
			Title: `笔趣阁qula`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=qu&url=` + grab.EncodeURL(`https://m.qu.la/`),
		},
		reader.Card{
			Title: `笔趣阁biqugeinfo`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=biqugeinfo&url=` + grab.EncodeURL(`https://m.biquge.info/`),
		},
		reader.Card{
			Title: `笔趣阁mcmssc`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=mcmssc&url=` + grab.EncodeURL(`https://www.mcmssc.com/`),
		},

		reader.Card{
			Title: `╅╅╅︺未满18岁禁止观看︺╅╆╆`,
			Type:  `link`,
			WxTo:  ``,
		},

		// reader.Card{
		// 	Title: `韩漫窝(18禁)`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/list?action=list&drive=hanmanwo&url=` + grab.EncodeURL(`http://www.hanmanwo.com/booklist`),
		// },

		reader.Card{
			Title: `韩漫库(18禁)`,
			Type:  `link`,
			WxTo:  `/pages/list?action=list&drive=hanmanku&url=` + grab.EncodeURL(`http://www.hanmanku.com/booklist`),
		},

		reader.Card{
			Title: `海猫吧(18禁)`,
			Type:  `link`,
			WxTo:  `/pages/list?action=list&drive=haimaoba&url=` + grab.EncodeURL(`http://www.haimaoba.com/list/0/`),
		},

		// reader.Card{
		// 	Title: `我爱妹子漫画(18禁)`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/list?action=list&drive=aimeizi5&url=` + grab.EncodeURL(`https://5aimeizi.com/booklist`),
		// },
		reader.Card{
			Title: `腐漫漫画(18禁)`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=fuman&url=` + grab.EncodeURL(`https://www.5aimeizi.com/`),
		},
		reader.Card{
			Title: `漫画台(18禁)`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=manhwa&url=` + grab.EncodeURL(`https://www.manhwa.cc/`),
		},
		reader.Card{
			Title: `看妹子漫画(18禁)`,
			Type:  `link`,
			WxTo:  `/pages/list?action=list&drive=kanmeizi&url=` + grab.EncodeURL(`https://www.kanmeizi.cc/booklist`),
		},
		reader.Card{
			Title: `伟叫兽漫画网(18禁)`,
			Type:  `link`,
			WxTo:  `/pages/categories?action=list&drive=weijiaoshou&url=` + grab.EncodeURL(`http://www.weijiaoshou.cn`),
		},
		reader.Card{
			Title: `漫物语(18禁)`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=manwuyu&url=` + grab.EncodeURL(`http://www.manwuyu.com/`),
		},
	}
	return c.JSON(http.StatusOK, list)
}

//GetAllBookResources 所有小说资源
func GetAllBookResources(c echo.Context) error {
	var list = reader.Catalog{}
	list.Title = `小说资源`

	list.SourceURL = ``

	list.Hash = ``

	list.Cards = []reader.Card{
		// reader.Card{
		// 	Title: `备用线路`,
		// 	Type:  `jumpapp`,
		// 	WxTo:  `/pages/index`,
		// 	AppID: `wx7c30b98c7f42f651`,
		// },
		reader.Card{
			Title: `起点小说网`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=qidian&url=` + grab.EncodeURL(`https://www.qidian.com`),
			// Appid: `wxa94ddd94358b2d1d`,
		},
		reader.Card{
			Title: `17K文学`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=17k&url=` + grab.EncodeURL(`http://www.17k.com`),
			// Appid: `wxa94ddd94358b2d1d`,
		},
		reader.Card{
			Title: `红袖添香`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=hongxiu&url=` + grab.EncodeURL(`https://www.hongxiu.com`),
			// Appid: `wxa94ddd94358b2d1d`,
		},
		reader.Card{
			Title: `潇湘书院`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=xxsy&url=` + grab.EncodeURL(`https://www.xxsy.net`),
			// Appid: `wxa94ddd94358b2d1d`,
		},
		reader.Card{
			Title: `笔趣阁jxla`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=jx&url=` + grab.EncodeURL(`https://m.jx.la/`),
			// Appid: `wxa94ddd94358b2d1d`,
		},
		reader.Card{
			Title: `笔趣阁mcmssc`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=mcmssc&url=` + grab.EncodeURL(`https://www.mcmssc.com/`),
			// Appid: `wxa94ddd94358b2d1d`,
		},
		reader.Card{
			Title: `笔趣阁paoshu8`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=paoshu8&url=` + grab.EncodeURL(`http://www.paoshu8.com/`),
			// Appid: `wxa94ddd94358b2d1d`,
		},

		reader.Card{
			Title: `顶点小说280xs`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=xs280&url=` + grab.EncodeURL(`https://www.280xs.com/`),
			// Appid: `wxa94ddd94358b2d1d`,
		},

		reader.Card{
			Title: `笔趣阁xbiquge`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=xbiquge&url=` + grab.EncodeURL(`http://www.xbiquge.la/`),
			// Appid: `wxa94ddd94358b2d1d`,
		},

		reader.Card{
			Title: `顶点小说booktxt`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=booktxt&url=` + grab.EncodeURL(`http://www.booktxt.net`),
			// Appid: `wxa94ddd94358b2d1d`,
		},

		reader.Card{
			Title: `笔趣阁qula`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=qu&url=` + grab.EncodeURL(`https://m.qu.la/`),
		},
		reader.Card{
			Title: `书阁小说网shugela`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=shuge&url=` + grab.EncodeURL(`https://m.shuge.la/`),
			// Appid: `wxa94ddd94358b2d1d`,
		},
		reader.Card{
			Title: `笔趣阁biqugeinfo`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=biqugeinfo&url=` + grab.EncodeURL(`https://m.biquge.info/`),
		},
		reader.Card{
			Title: `纵横小说网`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=zongheng&url=` + grab.EncodeURL(`http://book.zongheng.com`),
			// Appid: `wxa94ddd94358b2d1d`,
		},
		// reader.Card{
		// 	Title: `老司机小说`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/categories?drive=laosijixs&url=` + grab.EncodeURL(`http://m.laosijixs.com/`),
		// },
	}

	cf := cpi.GetConf()
	if cf.Search.LimiterStageShow && false { //受显示限制给资源控制
		list.Cards = append(list.Cards,
			reader.Card{
				Title: `老司机小说`,
				Type:  `link`,
				WxTo:  `/pages/categories?drive=laosijixs&url=` + grab.EncodeURL(`http://m.laosijixs.com/`),
			})
	}

	return c.JSON(http.StatusOK, list)
}

//GetAllCartoonResources 获取所有漫画资源
func GetAllCartoonResources(c echo.Context) error {
	var list = reader.Catalog{}
	list.Title = `漫画资源`

	list.SourceURL = ``

	list.Hash = ``

	list.Cards = []reader.Card{

		reader.Card{
			Title: `╅╅╅︺未满18岁禁止观看︺╅╆╆`,
			Type:  `link`,
			WxTo:  ``,
		},

		// reader.Card{
		// 	Title: `韩漫窝(18禁)`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/list?action=list&drive=hanmanwo&url=` + grab.EncodeURL(`http://www.hanmanwo.com/booklist`),
		// },

		reader.Card{
			Title: `韩漫库(18禁)`,
			Type:  `link`,
			WxTo:  `/pages/list?action=list&drive=hanmanku&url=` + grab.EncodeURL(`http://www.hanmanku.com/booklist`),
		},

		reader.Card{
			Title: `海猫吧(18禁)`,
			Type:  `link`,
			WxTo:  `/pages/list?action=list&drive=haimaoba&url=` + grab.EncodeURL(`http://www.haimaoba.com/list/0/`),
		},

		// reader.Card{
		// 	Title: `我爱妹子漫画(18禁)`,
		// 	Type:  `link`,
		// 	WxTo:  `/pages/list?action=list&drive=aimeizi5&url=` + grab.EncodeURL(`https://5aimeizi.com/booklist`),
		// },
		reader.Card{
			Title: `腐漫漫画(18禁)`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=fuman&url=` + grab.EncodeURL(`https://www.5aimeizi.com/`),
		},
		reader.Card{
			Title: `漫画台(18禁)`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=manhwa&url=` + grab.EncodeURL(`https://www.manhwa.cc/`),
		},
		reader.Card{
			Title: `看妹子漫画(18禁)`,
			Type:  `link`,
			WxTo:  `/pages/list?action=list&drive=kanmeizi&url=` + grab.EncodeURL(`https://www.kanmeizi.cc/booklist`),
		},
		reader.Card{
			Title: `伟叫兽漫画网(18禁)`,
			Type:  `link`,
			WxTo:  `/pages/categories?action=list&drive=weijiaoshou&url=` + grab.EncodeURL(`http://www.weijiaoshou.cn`),
		},
		reader.Card{
			Title: `漫物语(18禁)`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=manwuyu&url=` + grab.EncodeURL(`http://www.manwuyu.com/`),
		},
	}
	return c.JSON(http.StatusOK, list)
}

//GetAllLearnResources 获取所有编程资源
func GetAllLearnResources(c echo.Context) error {
	var list = reader.Catalog{}
	list.Title = `编程学习资料`

	list.SourceURL = ``

	list.Hash = ``

	list.Cards = []reader.Card{

		reader.Card{
			Title: `微信小程序开发入门系列教程`,
			Type:  `link`,
			WxTo:  `/pages/catalog?drive=blog&url=` + grab.EncodeURL(`https://xueyuanjun.com/laravel-from-appreciate-to-artisan`),
		},
		reader.Card{
			Title: `从学徒到工匠精校版`,
			Type:  `link`,
			WxTo:  `/pages/catalog?drive=blog&url=` + grab.EncodeURL(`https://xueyuanjun.com/wechat-miniprogram-tutorial`),
		},

		reader.Card{
			Title: `从入门到精通系列教程`,
			Type:  `link`,
			WxTo:  `/pages/catalog?drive=blog&url=` + grab.EncodeURL(`https://xueyuanjun.com/laravel-tutorial-5_7`),
		},
		reader.Card{
			Title: `Go语言入门教程`,
			Type:  `link`,
			WxTo:  `/pages/catalog?drive=blog&url=` + grab.EncodeURL(`https://xueyuanjun.com/golang-tutorials`),
		},
		reader.Card{
			Title: `德哥博客-最佳实践`,
			Type:  `link`,
			WxTo:  `/pages/catalog?drive=github&url=` + grab.EncodeURL(`https://github.com/digoal/blog/blob/master/class/24.md`),
		},

		reader.Card{
			Title: `德哥博客-经典案例`,
			Type:  `link`,
			WxTo:  `/pages/catalog?drive=github&url=` + grab.EncodeURL(`https://github.com/digoal/blog/blob/master/class/15.md`),
		},

		reader.Card{
			Title: `Laravel 项目开发规范`,
			Type:  `link`,
			WxTo:  `/pages/catalog?drive=learnku&url=` + grab.EncodeURL(`https://learnku.com/docs/laravel-specification/5.5`),
		},
		reader.Card{
			Title: `Laravel5.5开发文档`,
			Type:  `link`,
			WxTo:  `/pages/catalog?drive=learnku&url=` + grab.EncodeURL(`https://learnku.com/docs/laravel-specification/5.5`),
		},
		reader.Card{
			Title: `Laravel 5.5 中文文档`,
			Type:  `link`,
			WxTo:  `/pages/catalog?drive=learnku&url=` + grab.EncodeURL(`https://learnku.com/docs/laravel/5.5`),
		},
		reader.Card{
			Title: `Dingo API 2.0.0 中文文档`,
			Type:  `link`,
			WxTo:  `/pages/catalog?drive=learnku&url=` + grab.EncodeURL(`https://learnku.com/docs/dingo-api/2.0.0`),
		},
	}
	return c.JSON(http.StatusOK, list)
}
