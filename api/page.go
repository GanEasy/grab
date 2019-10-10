package api

import (
	"net/http"

	"github.com/GanEasy/grab"
	cpi "github.com/GanEasy/grab/core"
	"github.com/GanEasy/grab/reader"
	"github.com/labstack/echo"
)

// Carousel 小程序首页轮播内容(作为专题广告或其它的东西使用)
type Carousel struct {
	URL  string `json:"url"`
	Type string `json:"type"` // 期望可以同时支持视频播放(虽然很不现实)
	WxTo string `json:"wxto"` // 点击后跳转地址
}

// GetCarousels 获取首页走马灯数据
func GetCarousels(c echo.Context) error {
	var carousels []Carousel
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
}

// GetUserLinks 获取用户列表内容
func GetUserLinks(c echo.Context) error {
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
		// 	Image: `https://ossweb-img.qq.com/images/lol/web201310/skin/big37006.jpg`,
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
			Title: `功能介绍`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/article?drive=blog&url=` + grab.EncodeURL(`https://github.com/GanEasy/grab/blob/master/doc/newCatalog.md`),
			Style: `arrow`,
		},
		// Link{
		// 	Title: `使用说明`,
		// 	Icon:  ``,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/newCreate`,
		// 	Style: `arrow`,
		// },

		// Link{
		// 	Title: `异常说明`,
		// 	Icon:  ``,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/newCreate`,
		// 	Style: `arrow`,
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
	cf := cpi.GetConf()

	version := c.QueryParam("version")

	if cf.Search.LimitLevel || version == cf.Search.DevVersion { // 开启严格检查
		return c.JSON(http.StatusOK, GetWaitExamineExplore())
	}
	return c.JSON(http.StatusOK, GetPublishExploreLinks())
}

// GetWaitExamineExplore 用于审核的内容列表
func GetWaitExamineExplore() []Link {

	var links = []Link{
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
			Title: `Go语言入门教程`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/catalog?drive=blog&url=` + grab.EncodeURL(`https://xueyuanjun.com/golang-tutorials`),
			Style: `arrow`,
		},
		Link{
			Title: `德哥博客-最佳实践`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/catalog?drive=github&url=` + grab.EncodeURL(`https://github.com/digoal/blog/blob/master/class/24.md`),
			Style: `arrow`,
		},

		Link{
			Title: `德哥博客-经典案例`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/catalog?drive=github&url=` + grab.EncodeURL(`https://github.com/digoal/blog/blob/master/class/15.md`),
			Style: `arrow`,
		},

		Link{
			Title: `Laravel 项目开发规范`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/catalog?drive=github&url=` + grab.EncodeURL(`https://github.com/digoal/blog/blob/master/class/15.md`),
			Style: `arrow`,
		},
		Link{
			Title: `Laravel5.5开发文档`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/catalog?drive=learnku&url=` + grab.EncodeURL(`https://learnku.com/docs/laravel-specification/5.5`),
			Style: `arrow`,
		},
		Link{
			Title: `Laravel 5.5 中文文档`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/catalog?drive=learnku&url=` + grab.EncodeURL(`https://learnku.com/docs/laravel/5.5`),
			Style: `arrow`,
		},
		Link{
			Title: `Dingo API 2.0.0 中文文档`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/catalog?drive=learnku&url=` + grab.EncodeURL(`https://learnku.com/docs/dingo-api/2.0.0`),
			Style: `arrow`,
		},
	}
	return links

}

//GetPublishExploreLinks 获取公开发布的内容
func GetPublishExploreLinks() []Link {
	var links = []Link{
		// Link{
		// 	Title: `全部资源`,
		// 	Icon:  ``,
		// 	Type:  `link`,
		// 	Image: ``,
		// 	WxTo:  `/pages/transfer?action=allroesoures&drive=&url=`,
		// 	Style: `arrow`,
		// },

		Link{
			Title: `起点小说网`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/categories?drive=qidian&url=` + grab.EncodeURL(`https://www.qidian.com`),
			Style: `arrow`,
		},
		Link{
			Title: `纵横小说网`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/categories?drive=zongheng&url=` + grab.EncodeURL(`http://book.zongheng.com`),
			Style: `arrow`,
		},
		Link{
			Title: `17K文学`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/categories?drive=17k&url=` + grab.EncodeURL(`http://www.17k.com`),
			Style: `arrow`,
		},
		Link{
			Title: `笔下文学`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/categories?drive=bxwx&url=` + grab.EncodeURL(`https://www.bxwx.la`),
			Style: `arrow`,
		},

		Link{
			Title: `U小说阅读网`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/categories?drive=uxiaoshuo&url=` + grab.EncodeURL(`https://m.uxiaoshuo.com/`),
			Style: `arrow`,
		},
		Link{
			Title: `笔趣阁biquyun`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/categories?drive=biquyun&url=` + grab.EncodeURL(`https://m.biquyun.com`),
			Style: `arrow`,
		},

		Link{
			Title: `顶点小说`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/categories?drive=booktxt&url=` + grab.EncodeURL(`http://www.booktxt.net`),
			Style: `arrow`,
		},

		Link{
			Title: `笔趣阁soe8`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/categories?drive=soe8&url=` + grab.EncodeURL(`http://m.soe8.com/`),
			Style: `arrow`,
		},

		Link{
			Title: `笔趣阁xbiquge`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/categories?drive=xbiquge&url=` + grab.EncodeURL(`http://www.xbiquge.la/`),
			Style: `arrow`,
		},

		Link{
			Title: `笔趣阁qula`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/categories?drive=qu&url=` + grab.EncodeURL(`https://m.qu.la/`),
			Style: `arrow`,
		},

		Link{
			Title: `╅╅╅︺未满18岁禁止观看︺╅╆╆`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  ``,
			Style: ``,
		},

		Link{
			Title: `韩漫窝(18禁)`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/list?action=list&drive=hanmanwo&url=` + grab.EncodeURL(`http://www.hanmanwo.com/booklist`),
			Style: `arrow`,
		},

		Link{
			Title: `韩漫库(18禁)`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/list?action=list&drive=hanmanku&url=` + grab.EncodeURL(`http://www.hanmanku.com/booklist`),
			Style: `arrow`,
		},

		Link{
			Title: `海猫吧(18禁)`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/list?action=list&drive=haimaoba&url=` + grab.EncodeURL(`http://www.haimaoba.com/list/0/`),
			Style: `arrow`,
		},

		Link{
			Title: `我爱妹子漫画(18禁)`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/list?action=list&drive=aimeizi5&url=` + grab.EncodeURL(`https://5aimeizi.com/booklist`),
			Style: `arrow`,
		},
		Link{
			Title: `腐漫漫画(18禁)`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/categories?drive=fuman&url=` + grab.EncodeURL(`https://www.5aimeizi.com/`),
			Style: `arrow`,
		},
		Link{
			Title: `漫画台(18禁)`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/categories?drive=manhwa&url=` + grab.EncodeURL(`https://www.manhwa.cc/`),
			Style: `arrow`,
		},
		Link{
			Title: `看妹子漫画(18禁)`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/list?action=list&drive=kanmeizi&url=` + grab.EncodeURL(`https://www.kanmeizi.cc/booklist`),
			Style: `arrow`,
		},
		Link{
			Title: `伟叫兽漫画网(18禁)`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/categories?action=list&drive=weijiaoshou&url=` + grab.EncodeURL(`http://www.weijiaoshou.cn`),
			Style: `arrow`,
		},
		Link{
			Title: `漫物语(18禁)`,
			Icon:  ``,
			Type:  `link`,
			Image: ``,
			WxTo:  `/pages/categories?drive=manwuyu&url=` + grab.EncodeURL(`http://www.manwuyu.com/`),
			Style: `arrow`,
		},
		// Link{
		// 	Title: `交流群`,
		// 	Icon:  `cuIcon-group`,
		// 	Type:  `image`,
		// 	Image: `https://ossweb-img.qq.com/images/lol/web201310/skin/big37006.jpg`,
		// 	WxTo:  ``,
		// 	Style: ``,
		// },

	}

	return links
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
		reader.Card{
			Title: `笔下文学`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=bxwx&url=` + grab.EncodeURL(`https://www.bxwx.la`),
		},
		reader.Card{
			Title: `U小说阅读网`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=uxiaoshuo&url=` + grab.EncodeURL(`https://m.uxiaoshuo.com/`),
		},
		reader.Card{
			Title: `笔趣阁biquyun`,
			Type:  `link`,
			WxTo:  `/pages/categories?drive=biquyun&url=` + grab.EncodeURL(`https://m.biquyun.com`),
		},

		reader.Card{
			Title: `顶点小说`,
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
			Title: `╅╅╅︺未满18岁禁止观看︺╅╆╆`,
			Type:  `link`,
			WxTo:  ``,
		},

		reader.Card{
			Title: `韩漫窝(18禁)`,
			Type:  `link`,
			WxTo:  `/pages/list?action=list&drive=hanmanwo&url=` + grab.EncodeURL(`http://www.hanmanwo.com/booklist`),
		},

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

		reader.Card{
			Title: `我爱妹子漫画(18禁)`,
			Type:  `link`,
			WxTo:  `/pages/list?action=list&drive=aimeizi5&url=` + grab.EncodeURL(`https://5aimeizi.com/booklist`),
		},
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
