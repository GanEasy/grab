package grab

// DemoItem 示例详细
type DemoItem struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	Category string `json:"category"`
}

//RssDemoList Rss示例列表
func RssDemoList() []DemoItem {
	var list = []DemoItem{
		DemoItem{`月光博客`, `http://feed.williamlong.info/`, `Rss`},
		// DemoItem{`知乎每日精选`, `https://www.zhihu.com/rss`, `Rss`},
		// DemoItem{`cnbeta`, `https://www.cnbeta.com/backend.php`, `Rss`},
		// DemoItem{`国内新闻-腾讯`, `http://news.qq.com/newsgn/rss_newsgn.xml`, `Rss`},
	}
	return list
}

//ArticleDemoList 文章示例列表
func ArticleDemoList() []DemoItem {
	var list = []DemoItem{
		DemoItem{`wechatRank.com`, `https://wechatrank.com/getlist`, `Article`},
	}
	return list
}

//BookDemoList 小说示例列表
func BookDemoList() []DemoItem {
	var list = []DemoItem{
		DemoItem{`laravel5.6`, `http://laravelacademy.org/laravel-docs-5_6`, `Book`},
		// DemoItem{`点道为止`, `http://book.zongheng.com/showchapter/730066.html`, `Book`},
		// DemoItem{`修罗武神`, `http://www.17k.com/list/493239.html`, `Book`},
		// DemoItem{`万古仙穹`, `https://www.cangqionglongqi.com/wanguxianzuo/`, `Book`},
		// DemoItem{`斗罗大陆`, `http://www.biquge.info/10_10218/`, `Book`},
		// DemoItem{`圣墟`, `http://www.biqiuge.com/book/4772/`, `Book`},
	}
	return list
}
