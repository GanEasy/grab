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
}

//GetClassify 发现分类
func GetClassify() []Item {
	var list = []Item{
		Item{`资源库`, `/pages/resources/index`, "收集整理常用平台资源，方便快速阅读。"},
		Item{`各类榜单`, `/pages/rank/index`, "或许感兴趣的榜单，随时可以了解。"},
		Item{`专题集`, `/pages/topic/index`, "由小易同学维护的专题包"},
	}
	return list
}

//GetRanks 各类榜单
func GetRanks() []Item {
	var list = []Item{
		// uDec, _ := base64.URLEncoding.DecodeString(input)
		// eDec := base64.StdEncoding.EncodeToString(uDec)
		Item{`纵横小说榜`, `/pages/book/get?interpreter=default&url=` + EncodeURL(`http://www.zongheng.com/rank/details.html?rt=1&d=1`), ""},
		Item{`起点月票榜`, `/pages/book/get?interpreter=default&url=` + EncodeURL(`https://www.qidian.com/rank/hotsales`), ""},
		// Item{`动漫之家漫画排行榜`, `/pages/article/get?url=` + base64.URLEncoding.EncodeToString([]byte(`http://www.dm5.com/manhua-yaoshenji/`),""},
	}
	return list
}

//GetResources 自定义资源列表(支持平台目录)
func GetResources() []Item {
	var list = []Item{
		Item{`起点小说网`, `/pages/resources/get?interpreter=default&url=` + EncodeURL(`www.qidian.com`), ""},
		Item{`纵横小说网`, `/pages/resources/get?interpreter=default&url=` + EncodeURL(`www.zongheng.com`), ""},
		Item{`17K文学`, `/pages/resources/get?interpreter=default&url=` + EncodeURL(`www.17k.com`), ""},
	}
	return list
}

//GetResource 获得资源详细(分类)
func GetResource(url string) []Item {
	fmt.Println(`GetResource`, url)
	// todo 根据相应的url获得分类详细 后台控制
	var list = []Item{
		Item{`奇幻玄幻`, `/pages/book/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/store/c1/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), ""},
		Item{`武侠仙侠`, `/pages/book/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/store/c3/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), ""},
		Item{`历史军事`, `/pages/book/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/store/c6/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), ""},
		Item{`都市娱乐`, `/pages/book/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/store/c9/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), ""},
		Item{`科幻游戏`, `/pages/book/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/store/c15/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), ""},
		Item{`悬疑灵异`, `/pages/book/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/store/c18/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), ""},
		Item{`竞技同人`, `/pages/book/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/store/c21/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), ""},
		Item{`评论文集`, `/pages/book/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/store/c24/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), ""},
	}
	return list
}

//GetTopics 获取专题列表
func GetTopics() []Item {
	// todo 根据相应的url获得分类详细 后台控制
	var list = []Item{
		Item{`奇幻玄幻`, `/pages/book/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/store/c1/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), ""},
		Item{`武侠仙侠`, `/pages/book/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/store/c3/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), ""},
		Item{`历史军事`, `/pages/book/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/store/c6/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), ""},
		Item{`都市娱乐`, `/pages/book/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/store/c9/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), ""},
		Item{`科幻游戏`, `/pages/book/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/store/c15/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), ""},
		Item{`悬疑灵异`, `/pages/book/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/store/c18/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), ""},
		Item{`竞技同人`, `/pages/book/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/store/c21/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), ""},
		Item{`评论文集`, `/pages/book/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/store/c24/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), ""},
	}
	return list
}

//GetBooks 小说目录列表
func GetBooks(url string) []Item {
	// todo 目标返回
	fmt.Println(`GetBooks`, url)
	var list = []Item{
		Item{`元尊`, `/pages/chapter/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/showchapter/685640.html`), ""},
		Item{`圣武星辰`, `/pages/chapter/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/showchapter/682920.html`), ""},
		Item{`祭炼山河`, `/pages/chapter/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/showchapter/309318.html`), ""},
		Item{`逆天邪神`, `/pages/chapter/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/showchapter/408586.html`), ""},
		Item{`一剑独尊`, `/pages/chapter/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/showchapter/777234.html`), ""},
		Item{`魔域`, `/pages/chapter/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/showchapter/568980.html`), ""},
		Item{`永夜君王`, `/pages/chapter/get?interpreter=default&url=` + EncodeURL(`http://book.zongheng.com/showchapter/342974.html`), ""},
	}
	return list
}

//GetChapters 小说章节列表
func GetChapters(url string) []Item {

	fmt.Println(`GetChapters`, url)
	var list = []Item{
		Item{`月光博客`, `/pages/chapter/info?interpreter=default&url=` + EncodeURL(`http://feed.williamlong.info/`), ""},
		Item{`修罗武神`, `/pages/chapter/info?interpreter=default&url=` + EncodeURL(`http://www.17k.com/list/493239.html`), ""},
	}
	return list
}

//GetChapter 小说章节详细
func GetChapter(url string) []Item {

	fmt.Println(`GetChapters`, url)
	var list = []Item{
		Item{`月光博客`, `/pages/chapter/info?interpreter=default&url=` + EncodeURL(`http://feed.williamlong.info/`), ""},
		Item{`修罗武神`, `/pages/chapter/info?interpreter=default&url=` + EncodeURL(`http://www.17k.com/list/493239.html`), ""},
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
