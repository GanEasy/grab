package grab

import (
	"fmt"
)

//GetWaitExamineClassify 获取做给审核检查的class
func GetWaitExamineClassify() []Item {
	var list = []Item{
		// Item{`从学徒到工匠精校版`, `/pages/catalog?drive=blog&url=` + EncodeURL(`https://laravelacademy.org/laravel-from-appreciate-to-artisan`), "", "card"},
		// Item{`微信小程序开发入门系列教程`, `/pages/catalog?drive=blog&url=` + EncodeURL(`https://laravelacademy.org/wechat-miniprogram-tutorial`), "", "card"},
		// Item{`从入门到精通系列教程`, `/pages/catalog?drive=blog&url=` + EncodeURL(`https://laravelacademy.org/laravel-tutorial-5_7`), "", "card"},
		// Item{`Go语言入门教程`, `/pages/catalog?drive=blog&url=` + EncodeURL(`https://laravelacademy.org/golang-tutorials`), "", "card"},
		// Item{`德哥博客`, `/pages/catalog?drive=blog&url=` + EncodeURL(`https://github.com/digoal/blog`), "", "card"},
		Item{`德哥博客-最佳实践`, `/pages/catalog?drive=github&url=` + EncodeURL(`https://github.com/digoal/blog/blob/master/class/24.md`), "", "card"},
		Item{`德哥博客-经典案例`, `/pages/catalog?drive=github&url=` + EncodeURL(`https://github.com/digoal/blog/blob/master/class/15.md`), "", "card"},
		Item{`Laravel 项目开发规范`, `/pages/catalog?drive=learnku&url=` + EncodeURL(`https://learnku.com/docs/laravel-specification/5.5`), "", "card"},
		Item{`Laravel5.5开发文档`, `/pages/catalog?drive=learnku&url=` + EncodeURL(`https://learnku.com/docs/laravel/5.5`), "", "card"},
		Item{`Dingo API 2.0.0 中文文档`, `/pages/catalog?drive=learnku&url=` + EncodeURL(`https://learnku.com/docs/dingo-api/2.0.0`), "", "card"},
	}
	return list
}

//GetClassify 发现分类
func GetClassify() []Item {

	var list = []Item{
		Item{`资源库`, `/pages/resources/index`, "收集整理常用平台资源，方便快速阅读。", "card"},
		Item{`各类榜单`, `/pages/rank/index`, "或许感兴趣的榜单，随时可以了解。", "card"},
		Item{`专题集`, `/pages/topic/index`, "由小易同学维护的专题包", "card"},
	}
	return list
}

//GetDrives 获取所有解释引擎
func GetDrives() []Item {
	var list = []Item{
		Item{`泛文本类`, `/pages/user/createSource?drive=text`, "转码目录所有链接，并将目标详情页转换成文本详情。", "card"},
		Item{`泛资讯类`, `/pages/user/createSource?drive=news`, "转码目录所有链接，并将目标详情页转换成图文详情。", "card"},
		Item{`起点小说`, `/pages/user/createSource?drive=qidian`, "精准解释起点小说免费章节", "card"},
	}
	return list
}

//GetAbouts 获得关于我们的介绍
func GetAbouts() []Item {
	var list = []Item{
		Item{`关于我们`, ``, "提供第三方web目录内容实时转码阅读服务。", "card"},
		Item{`数据缓存`, ``, "用户数据使用storage缓存在用户手机。", "card"},
		Item{`免责声明`, ``, "通过转码技术展现第三方内容，内容版权及相关连带责任归源站所有。", "card"},
		Item{`来源声明`, ``, "页面中“域名、原网页”等信息，都是转码数据的来源页，可点击复制。", "card"},
	}
	return list
}

//GetHelps 获取帮助(常见问题)
func GetHelps() []Item {
	var list = []Item{
		Item{`Q:有什么用？`, ``, "A:工具通过转码技术实现第三方资源阅读", "card"},
		Item{`Q:能看什么？`, ``, "A:对于用户感兴趣的网站内容进行转码，以供阅读", "card"},
		Item{`Q:资源免费看？`, ``, "A:免费提供对第三方网站数据转码，但不提供对方需要授权才能阅读的内容(我们不破解第三方需要付费才能阅读的内容，你可以通过准确关键字搜索看看有哪些平台有相应免费资源)。", "card"},
		Item{`Q:资源找不到？`, ``, "A:您可以通过百度搜索引擎找到资源复制其目录地址，通过我的》创建源》粘目录链接地址》点击确认进行转码(纯文字选择文本类解释器，有图片有文本选择图文类，RSS链接请选择Rss)，也可以直接联系客服免费订制资源转码服务", "card"},
		Item{`Q:如何收藏？`, ``, "A:在目录级别，点击“加入看单”，即可在“看单”快速进入目录查看资源", "card"},
		Item{`Q:安全和隐私`, ``, "A:服务器不收集用户任何数据，看单缓存在本地，通过我的》清空数据即可清除用户所有数据(部分手机删除最近使用也会清空数据)", "card"},
	}
	return list
}

//GetRanks 各类榜单
func GetRanks() []Item {
	var list = []Item{
		// uDec, _ := base64.URLEncoding.DecodeString(input)
		// eDec := base64.StdEncoding.EncodeToString(uDec)
		Item{`纵横小说榜`, `/pages/book/get?drive=default&url=` + EncodeURL(`http://www.zongheng.com/rank/details.html?rt=1&d=1`), "", "link"},
		Item{`起点月票榜`, `/pages/book/get?drive=default&url=` + EncodeURL(`https://www.qidian.com/rank/hotsales`), "", "link"},
		// Item{`动漫之家漫画排行榜`, `/pages/article/get?url=` + base64.URLEncoding.EncodeToString([]byte(`http://www.dm5.com/manhua-yaoshenji/`),""},
	}
	return list
}

//GetResources 自定义资源列表(支持平台目录)
func GetResources() []Item {
	var list = []Item{
		// Item{`起点小说网`, `/pages/transfer/list?action=resource&drive=qidian&url=` + EncodeURL(`https://www.qidian.com`), "", "link"},
		// Item{`纵横小说网`, `/pages/transfer/list?action=resource&drive=zongheng&url=` + EncodeURL(`http://book.zongheng.com`), "", "link"},
		// Item{`17K文学`, `/pages/transfer/list?action=resource&drive=17k&url=` + EncodeURL(`http://www.17k.com`), "", "link"},
		// Item{`落秋中文`, `/pages/transfer/list?action=resource&drive=luoqiu&url=` + EncodeURL(`http://www.luoqiu.com`), "", "link"},

		// drive=article&url=aHR0cHM6Ly93ZWNoYXRyYW5rLmNvbS9nZXRsaXN0
		// Item{`微信精选`, `/pages/catalog?drive=article&url=` + EncodeURL(`https://wechatrank.com/getlist`), "", "link"},
		// Item{`妹子图`, `/pages/rss?drive=rss&url=` + EncodeURL(`https://rsshub.app/mzitu/home`), "", "link"},

		Item{`起点小说网`, `/pages/categories?drive=qidian&url=` + EncodeURL(`https://www.qidian.com`), "", "link"},
		Item{`纵横小说网`, `/pages/categories?drive=zongheng&url=` + EncodeURL(`http://book.zongheng.com`), "", "link"},
		Item{`顶点小说`, `/pages/categories?drive=booktxt&url=` + EncodeURL(`http://www.booktxt.net`), "", "link"},
		// Item{`落秋中文`, `/pages/categories/get?drive=luoqiu&url=` + EncodeURL(`http://www.luoqiu.com`), "", "link"},
		// Item{`7878小说`, `/pages/categories?drive=7878xs&url=` + EncodeURL(`http://www.7878xs.com`), "", "link"},
		Item{`笔下文学`, `/pages/categories?drive=bxwx&url=` + EncodeURL(`https://www.bxwx.la`), "", "link"},
		Item{`U小说阅读网`, `/pages/categories?drive=uxiaoshuo&url=` + EncodeURL(`https://m.uxiaoshuo.com/`), "", "link"},
		Item{`笔趣阁biquyun`, `/pages/categories?drive=biquyun&url=` + EncodeURL(`https://m.biquyun.com`), "", "link"},
		Item{`笔趣阁soe8`, `/pages/categories?drive=soe8&url=` + EncodeURL(`http://m.soe8.com/`), "", "link"},
		Item{`笔趣阁xbiquge`, `/pages/categories?drive=xbiquge&url=` + EncodeURL(`http://www.xbiquge.la/`), "", "link"},
		Item{`17K文学`, `/pages/categories?drive=17k&url=` + EncodeURL(`http://www.17k.com`), "", "link"},

		Item{`╅╅╅︺未满18岁禁止观看︺╅╆╆`, ``, "", "link"},
		Item{`海猫吧(18禁)`, `/pages/list?action=list&drive=haimaoba&url=` + EncodeURL(`http://www.haimaoba.com/list/0/`), "", "link"},
		Item{`我爱妹子漫画(18禁)`, `/pages/list?action=list&drive=aimeizi5&url=` + EncodeURL(`https://5aimeizi.com/booklist`), "", "link"},
		// Item{`我爱妹子漫画(18禁)`, `/pages/categories?drive=aimeizi5&url=` + EncodeURL(`https://www.fuman.cc/`), "", "link"},
		Item{`腐漫漫画(18禁)`, `/pages/categories?drive=fuman&url=` + EncodeURL(`https://www.5aimeizi.com/`), "", "link"},
		Item{`漫画台(18禁)`, `/pages/categories?drive=manhwa&url=` + EncodeURL(`https://www.manhwa.cc/`), "", "link"},

		Item{`看妹子漫画(18禁)`, `/pages/list?action=list&drive=kanmeizi&url=` + EncodeURL(`https://www.kanmeizi.cc/booklist`), "", "link"},
		Item{`伟叫兽漫画网(18禁)`, `/pages/categories?action=list&drive=weijiaoshou&url=` + EncodeURL(`http://www.weijiaoshou.cn`), "", "link"},

		// Item{`看妹子漫画(18禁)`, `/pages/categories?drive=kanmeizi&url=` + EncodeURL(`https://www.kanmeizi.cc/`), "", "link"},
		// Item{`无双漫画(18禁)`, `/pages/categories?drive=r2hm&url=` + EncodeURL(`https://r2hm.com/`), "", "link"},
		Item{`漫物语(18禁)`, `/pages/categories?drive=manwuyu&url=` + EncodeURL(`http://www.manwuyu.com/`), "", "link"},
		// SeventeenKReader
		// Item{`17K文学`, `/pages/book/get?drive=book&url=` + EncodeURL(`http://www.17k.com`), "", "link"},
	}
	return list
}

//GetResource 获得资源详细(分类)
func GetResource(url string) []Item {
	fmt.Println(`GetResource`, url)
	// todo 根据相应的url获得分类详细 后台控制
	var list = []Item{
		Item{`奇幻玄幻`, `/pages/book/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/store/c1/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", "link"},
		Item{`武侠仙侠`, `/pages/book/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/store/c3/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", "link"},
		Item{`历史军事`, `/pages/book/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/store/c6/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", "link"},
		Item{`都市娱乐`, `/pages/book/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/store/c9/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", "link"},
		Item{`科幻游戏`, `/pages/book/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/store/c15/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", "link"},
		Item{`悬疑灵异`, `/pages/book/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/store/c18/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", "link"},
		Item{`竞技同人`, `/pages/book/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/store/c21/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", "link"},
		Item{`评论文集`, `/pages/book/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/store/c24/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", "link"},
	}
	return list
}

//GetTopics 获取专题列表
func GetTopics() []Item {
	// todo 根据相应的url获得分类详细 后台控制
	var list = []Item{
		Item{`奇幻玄幻`, `/pages/book/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/store/c1/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", "link"},
		Item{`武侠仙侠`, `/pages/book/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/store/c3/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", "link"},
		Item{`历史军事`, `/pages/book/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/store/c6/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", "link"},
		Item{`都市娱乐`, `/pages/book/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/store/c9/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", "link"},
		Item{`科幻游戏`, `/pages/book/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/store/c15/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", "link"},
		Item{`悬疑灵异`, `/pages/book/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/store/c18/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", "link"},
		Item{`竞技同人`, `/pages/book/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/store/c21/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", "link"},
		Item{`评论文集`, `/pages/book/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/store/c24/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", "link"},
	}
	return list
}

//GetBooks 小说目录列表
func GetBooks(url string) []Item {
	// todo 目标返回
	fmt.Println(`GetBooks`, url)
	var list = []Item{
		Item{`元尊`, `/pages/chapter/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/showchapter/685640.html`), "", "link"},
		Item{`圣武星辰`, `/pages/chapter/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/showchapter/682920.html`), "", "link"},
		Item{`祭炼山河`, `/pages/chapter/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/showchapter/309318.html`), "", "link"},
		Item{`逆天邪神`, `/pages/chapter/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/showchapter/408586.html`), "", "link"},
		Item{`一剑独尊`, `/pages/chapter/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/showchapter/777234.html`), "", "link"},
		Item{`魔域`, `/pages/chapter/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/showchapter/568980.html`), "", "link"},
		Item{`永夜君王`, `/pages/chapter/get?drive=default&url=` + EncodeURL(`http://book.zongheng.com/showchapter/342974.html`), "", "link"},
	}
	return list
}

//GetChapters 小说章节列表
func GetChapters(url string) []Item {

	fmt.Println(`GetChapters`, url)
	var list = []Item{
		Item{`月光博客`, `/pages/chapter/info?drive=default&url=` + EncodeURL(`http://feed.williamlong.info/`), "", "link"},
		Item{`修罗武神`, `/pages/chapter/info?drive=default&url=` + EncodeURL(`http://www.17k.com/list/493239.html`), "", "link"},
	}
	return list
}

//GetChapter 小说章节详细
func GetChapter(url string) []Item {

	fmt.Println(`GetChapters`, url)
	var list = []Item{
		Item{`月光博客`, `/pages/chapter/info?drive=default&url=` + EncodeURL(`http://feed.williamlong.info/`), "", "link"},
		Item{`修罗武神`, `/pages/chapter/info?drive=default&url=` + EncodeURL(`http://www.17k.com/list/493239.html`), "", "link"},
	}
	return list
}

//RssDemoList Rss示例列表
func RssDemoList() []DemoItem {
	var list = []DemoItem{
		DemoItem{`月光博客`, `http://feed.williamlong.info/`, `Rss`},
		DemoItem{`知乎每日精选`, `https://www.zhihu.com/rss`, `Rss`},
		DemoItem{`cnbeta`, `https://www.cnbeta.com/backend.php`, `Rss`},
		DemoItem{`国内新闻-腾讯`, `http://news.qq.com/newsgn/rss_newsgn.xml`, `Rss`},
	}
	return list
}

//ArticleDemoList 文章示例列表
func ArticleDemoList() []DemoItem {
	var list = []DemoItem{
		// DemoItem{`laravel5.6`, `http://laravelacademy.org/laravel-docs-5_6`, `Article`},
		DemoItem{`wechatRank.com`, `https://wechatrank.com/getlist`, `Article`},
	}
	return list
}

//BookDemoList 小说示例列表
func BookDemoList() []DemoItem {
	var list = []DemoItem{
		// DemoItem{`laravel5.6`, `http://laravelacademy.org/laravel-docs-5_6`, `Book`},
		DemoItem{`点道为止`, `http://book.zongheng.com/showchapter/730066.html`, `Book`},
		DemoItem{`修罗武神`, `http://www.17k.com/list/493239.html`, `Book`},
		DemoItem{`万古仙穹`, `https://www.cangqionglongqi.com/wanguxianzuo/`, `Book`},
		DemoItem{`斗罗大陆`, `http://www.biquge.info/10_10218/`, `Book`},
		DemoItem{`圣墟`, `http://www.biqiuge.com/book/4772/`, `Book`},
	}
	return list
}
