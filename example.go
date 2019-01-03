package grab

import (
	"fmt"
)

// DemoItem 示例详细
type DemoItem struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	Category string `json:"category"`
}

// Item 小程序授受参数明细
type Item struct {
	Title string `json:"title"`
	WxTo  string `json:"wxto"`
	Intro string `json:"intro"`
	Type  string `json:"type"`
}

//GetWaitExamineClassify 获取做给审核检查的class
func GetWaitExamineClassify() []Item {
	var list = []Item{
		Item{`从学徒到工匠精校版`, `/pages/catalog?drive=blog&url=` + EncodeURL(`https://laravelacademy.org/laravel-from-appreciate-to-artisan`), "", "card"},
		Item{`微信小程序开发入门系列教程`, `/pages/catalog?drive=blog&url=` + EncodeURL(`https://laravelacademy.org/wechat-miniprogram-tutorial`), "", "card"},
		Item{`从入门到精通系列教程`, `/pages/catalog?drive=blog&url=` + EncodeURL(`https://laravelacademy.org/laravel-tutorial-5_7`), "", "card"},
		Item{`德哥博客`, `/pages/catalog?drive=blog&url=` + EncodeURL(`https://github.com/digoal/blog`), "", "card"},
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
		Item{`关于我们`, ``, "阅读助手提供第三方web目录内容实时转码阅读服务。", "card"},
		Item{`数据缓存`, ``, "用户数据使用本地数据缓存，阅读助手不对数据安全做任何保证。", "card"},
		Item{`免责声明`, ``, "阅读助手实时转码(服务器不缓存)第三方内容，内容版权归源站所有。", "card"},
		Item{`来源声明`, ``, "阅读助手页面中“域名、原网页”等信息，都是转码数据的来源页，可点击复制。", "card"},
	}
	return list
}

//GetHelps 获取帮助(常见问题)
func GetHelps() []Item {
	var list = []Item{
		Item{`Q:如何订阅`, ``, "A:复制您感兴趣的网页链接地址，在“创建源”中提交即可！", "card"},
		Item{`Q:选择类型`, ``, "A:纯文字选择文本类解释器，有图片有文本选择图文类，RSS链接请选择Rss", "card"},
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
		Item{`微信精选`, `/pages/catalog?drive=article&url=` + EncodeURL(`https://wechatrank.com/`), "", "link"},
		Item{`妹子图`, `/pages/rss?drive=rss&url=` + EncodeURL(`https://rsshub.app/mzitu/home`), "", "link"},

		Item{`起点小说网`, `/pages/categories?drive=qidian&url=` + EncodeURL(`https://www.qidian.com`), "", "link"},
		Item{`纵横小说网`, `/pages/categories?drive=zongheng&url=` + EncodeURL(`http://book.zongheng.com`), "", "link"},
		Item{`17K文学`, `/pages/categories?drive=17k&url=` + EncodeURL(`http://www.17k.com`), "", "link"},
		Item{`顶点小说`, `/pages/categories?drive=booktxt&url=` + EncodeURL(`http://www.booktxt.net`), "", "link"},
		// Item{`落秋中文`, `/pages/categories/get?drive=luoqiu&url=` + EncodeURL(`http://www.luoqiu.com`), "", "link"},
		Item{`7878小说`, `/pages/categories?drive=7878xs&url=` + EncodeURL(`http://www.7878xs.com`), "", "link"},
		Item{`笔下文学`, `/pages/categories?drive=bxwx&url=` + EncodeURL(`https://www.bxwx.la`), "", "link"},
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
