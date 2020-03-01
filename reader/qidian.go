package reader

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/yizenghui/chromedp"
	"golang.org/x/net/html/charset"
)

//QidianReader 纵横小说网
type QidianReader struct {
}

// GetCategories 获取所有自定义分类(写死)
func (r QidianReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://book.qidian.com`

	list.Title = `分类-起点中文网`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`全部`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`玄幻`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=21&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-东方玄幻`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=21&subCateId=8&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-异世大陆`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=21&subCateId=73&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-王朝争霸`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=21&subCateId=58&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-高武世界`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=21&subCateId=78&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`奇幻`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=1&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-现代魔法`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=1&subCateId=38&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-剑与魔法`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=1&subCateId=62&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-史诗奇幻`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=1&subCateId=201&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-黑暗幻想`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=1&subCateId=202&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-历史神话`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=1&subCateId=20092&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-另类幻想`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=1&subCateId=20093&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`武侠`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=2&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-传统武侠`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=2&subCateId=5&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-武侠幻想`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=2&subCateId=30&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-国术无双`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=2&subCateId=206&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-古武未来`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=2&subCateId=20099&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-武侠同人`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=2&subCateId=20100&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`仙侠`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=22&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-修真文明`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=22&subCateId=18&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-幻想修仙`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=22&subCateId=44&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-现代修真`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=22&subCateId=64&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-神话修真`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=22&subCateId=207&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-古典仙侠`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=22&subCateId=20101&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`都市`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=4&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-都市生活`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=4&subCateId=12&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-都市异能`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=4&subCateId=16&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-异术超能`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=4&subCateId=74&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-青春校园`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=4&subCateId=130&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-娱乐明星`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=4&subCateId=151&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-商战职场`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=4&subCateId=153&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`军事`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=6&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-军旅生涯`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=6&subCateId=54&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-军事战争`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=6&subCateId=65&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-战争幻想`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=6&subCateId=80&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-抗战烽火`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=6&subCateId=230&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-谍战特工`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=6&subCateId=231&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`历史`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=5&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-架空历史`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=5&subCateId=22&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-秦汉三国`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=5&subCateId=48&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-上古先秦`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=5&subCateId=220&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-历史传记`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=5&subCateId=32&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-两晋隋唐`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=5&subCateId=222&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-五代十国`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=5&subCateId=223&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-两宋元明`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=5&subCateId=224&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-清史民国`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=5&subCateId=225&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-外国历史`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=5&subCateId=226&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-民间传说`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=5&subCateId=20094&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},

		Card{`游戏`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=7&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-电子竞技`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=7&subCateId=7&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-虚拟网游`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=7&subCateId=70&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-游戏异界`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=7&subCateId=240&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-游戏系统`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=7&subCateId=20102&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-游戏主播`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=7&subCateId=20103&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0 `), "", `link`, ``, nil, ``},

		Card{`体育`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=8&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-篮球运动`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=8&subCateId=28&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-体育赛事`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=8&subCateId=55&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`-足球运动`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=8&subCateId=82&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},

		Card{`科幻`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=9&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`现实`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=15&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		// Card{`-社会乡土`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=15&subCateId=20104&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		// Card{`-生活时尚`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=15&subCateId=20105&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		// Card{`-文学艺术`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=15&subCateId=20106&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		// Card{`-成功励志`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=15&subCateId=20107&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		// Card{`-青春文学`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=15&subCateId=20108&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		// Card{`-爱情婚姻`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=15&subCateId=6&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		// Card{`-现实百态`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=15&subCateId=209&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},

		Card{`灵异`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=10&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`二次元`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=12&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
		Card{`短篇`, `/pages/list?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=20076&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil, ``},
	}
	list.SearchSupport = true
	return list, nil
}

// GetList 获取分类书籍列表
func (r QidianReader) GetList(urlStr string) (list Catalog, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, ``)
	if err != nil {
		return
	}

	g, e := goquery.NewDocumentFromReader(strings.NewReader(html))

	if e != nil {
		return list, e
	}

	list.Title = g.Find("title").Text()

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://book.qidian.com/info/1010734492`, `https://book.qidian.com/info/1010868264`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `qidian`)

	list.SourceURL = urlStr

	// html := `{"I":"5333","V":"马经理"},`
	// page := FindString(`/p(?P<page>[^"]+)/`, html, "page")

	page := FindString(`&page=(?P<page>(\d)+)&`, urlStr, "page")
	p, err := strconv.Atoi(page)
	if p > 0 && p < 5 && err == nil {
		// 已经组装url
		nextURL := strings.Replace(urlStr, fmt.Sprintf(`&page=%v&`, p), fmt.Sprintf(`&page=%v&`, p+1), -1)
		list.Next = Link{`下一页`, EncodeURL(nextURL), ``}
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// Search 搜索资源
func (r QidianReader) Search(keyword string) (list Catalog, err error) {

	urlStr := `https://www.qidian.com/search?kw=` + keyword
	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}

	var html string

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err = chromedp.Run(ctx,
		chromedp.Navigate(urlStr),
		chromedp.Sleep(time.Second*2),
		chromedp.OuterHTML("html", &html),
	)
	if err != nil {
		// log.Fatal(err)
		return
	}

	html, err = FindContentHTML(html, `#result-list`)
	// html, err := GetHTML(urlStr, `#result-list`)
	// if err != nil {
	// 	return
	// }
	// log.Println(html)

	g, e := goquery.NewDocumentFromReader(strings.NewReader(html))

	if e != nil {
		return list, e
	}

	list.Title = fmt.Sprintf(`%v - 搜索结果 - 起点小说qidian.com`, keyword)

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)
	// log.Println(links)
	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://book.qidian.com/info/1010734492`, `https://book.qidian.com/info/1010868264`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `qidian`)

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetCatalog 获取章节列表
func (r QidianReader) GetCatalog(urlStr string) (list Catalog, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, ``)
	if err != nil {
		return
	}

	g, e := goquery.NewDocumentFromReader(strings.NewReader(html))

	if e != nil {
		return list, e
	}

	list.Title = FindString(`《(?P<title>(.)+)》_`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	catalogMsg := g.Find("#J-catalogCount").Text()
	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)
	var needLinks []Link
	if catalogMsg == `` { //todo 从 https://book.qidian.com/ajax/book/category?_csrfToken=&bookId=1004608738 中获取章节列表(要解释json)
		// panic(`catalogMsg`)

		bookID := FindString(`/info/(?P<id>(\d)+)`, urlStr, "id")

		if bookID != `` {
			links, _ = r.GetChaptersLinksByJSON(bookID)
			needLinks = links
		}

	} else {

		var state bool
		for _, l := range links { //起点普通和VIP章节不同地址
			l.URL, state = JaccardMateGetURL(l.URL, `https://read.qidian.com/chapter/ORlSeSgZ6E_MQzCecGvf7A2/DKk0ho2xSYTM5j8_3RRvhw2`, `https://read.qidian.com/chapter/_AaqI-dPJJ4uTkiRw_sFYA2/_4Wioy7TTQD6ItTi_ILQ7A2`, ``)
			if state {
				needLinks = append(needLinks, l)
			} else {
				l.URL, state = JaccardMateGetURL(l.URL, `https://vipreader.qidian.com/chapter/1004608738/347194141`, `https://vipreader.qidian.com/chapter/1010734492/399246504`, ``)
				if state {
					needLinks = append(needLinks, l)
				}
			}
		}

	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `qidian`)

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetInfo 获取章节正文内容
func (r QidianReader) GetInfo(urlStr string) (ret Content, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, ``)
	if err != nil {
		return ret, err
	}
	// log.Println(html)
	article, err := GetActicleByHTML(html)
	if err != nil {
		return ret, err
	}

	article.Readable(urlStr)

	ret.Title = article.Title

	ret.Title = FindString(`_(?P<title>(.)+)_起点中文网`, article.Title, "title")
	if ret.Title == `` {
		ret.Title = article.Title
	}

	ret.SourceURL = urlStr

	c := MarkDownFormatContent(article.ReadContent)

	c = BookContReplace(c)

	ret.Contents = GetSectionByContent(c)

	links, _ := GetLinkByHTML(urlStr, html)
	ret.Previous = GetPreviousLink(links)
	if ret.Previous.URL != `` {
		ret.Previous.URL = EncodeURL(ret.Previous.URL)
	}
	ret.Next = GetNextLink(links)
	if ret.Next.URL != `` {
		ret.Next.URL = EncodeURL(ret.Next.URL)
	}
	return ret, nil

}

// GetChaptersLinksByJSON 获取章节链接列表
func (r QidianReader) GetChaptersLinksByJSON(bookID string) (links []Link, err error) {

	urlStr := fmt.Sprintf(`https://book.qidian.com/ajax/book/category?_csrfToken=&bookId=%v`, bookID)
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return
	}
	req.Header = make(http.Header)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	reader, err := charset.NewReader(resp.Body, strings.ToLower(resp.Header.Get("Content-Type")))
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(reader)

	type QiChaptersJsonDataCsChapter struct {
		UT          string `json:"uT"`
		ChapterName string `json:"cN"`
		ChapterURL  string `json:"cU"`
		UuID        int    `json:"uuid"`
		ID          int    `json:"id"`
		Ss          int    `json:"sS"`
	}

	type QiChaptersJsonDataCs struct {
		CCnt     int                           `json:"cCnt"`
		Chapters []QiChaptersJsonDataCsChapter `json:"cs"`
		// Chapters []map[string]interface{}      `json:"cs"`
		// Chapters map[int]interface{} `json:"cs"`
	}
	type QiChaptersJsonData struct {
		ChapterTotal int `json:"chapterTotalCnt"`
		// Vs           map[string]QiChaptersJsonDataCs `json:"vs"`
		Vs []QiChaptersJsonDataCs `json:"vs"`
		// Vs map[string]interface{} `json:"vs"`
		// Vs []interface{} `json:"vs"`
	}
	type QiChaptersJson struct {
		Code int `json:"code"`
		// Data map[string]interface{} `json:"data['vs']['cs']"`
		Data QiChaptersJsonData `json:"data"`
	}

	var m QiChaptersJson
	err = json.Unmarshal(bs, &m)

	if err == nil {
		for _, v := range m.Data.Vs {
			for _, vv := range v.Chapters {
				if vv.Ss == 1 {

					links = append(links, Link{
						vv.ChapterName,
						fmt.Sprintf(`https://read.qidian.com/chapter/%v`, vv.ChapterURL),
						``,
					})
				} else {

					links = append(links, Link{
						vv.ChapterName,
						fmt.Sprintf(`https://vipreader.qidian.com/chapter/%v/%v`, bookID, vv.ID),
						``,
					})
				}
			}
		}

	}
	return

}
