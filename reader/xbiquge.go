package reader

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//XbiqugeReader U小说阅读网 (盗版小说网站)
type XbiqugeReader struct {
}

// GetCategories 获取所有分类
func (r XbiqugeReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://m.xbiquge.com/`

	list.Title = `分类-xbiquge.la`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`玄幻·奇幻`, `/pages/list?action=book&drive=xbiquge&url=` + EncodeURL(`http://www.xbiquge.la/xuanhuanxiaoshuo/`), "", `link`, ``, nil, ``, ``},
		Card{`修真·仙侠`, `/pages/list?action=book&drive=xbiquge&url=` + EncodeURL(`http://www.xbiquge.la/xiuzhenxiaoshuo/`), "", `link`, ``, nil, ``, ``},
		Card{`都市·青春`, `/pages/list?action=book&drive=xbiquge&url=` + EncodeURL(`http://www.xbiquge.la/dushixiaoshuo/`), "", `link`, ``, nil, ``, ``},
		Card{`历史·穿越`, `/pages/list?action=book&drive=xbiquge&url=` + EncodeURL(`http://www.xbiquge.la/chuanyuexiaoshuo/`), "", `link`, ``, nil, ``, ``},
		Card{`网游·竞技`, `/pages/list?action=book&drive=xbiquge&url=` + EncodeURL(`http://www.xbiquge.la/wangyouxiaoshuo/`), "", `link`, ``, nil, ``, ``},
		Card{`科幻·灵异`, `/pages/list?action=book&drive=xbiquge&url=` + EncodeURL(`http://www.xbiquge.la/kehuanxiaoshuo/`), "", `link`, ``, nil, ``, ``},
	}
	return list, nil
}

// GetList 获取书籍列表列表
func (r XbiqugeReader) GetList(urlStr string) (list Catalog, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, `.l`)
	if err != nil {
		return
	}

	g, e := goquery.NewDocumentFromReader(strings.NewReader(html))

	if e != nil {
		return list, e
	}
	list.Title = FindString(`(?P<title>(.)+)_好看的`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://www.xbiquge.la/39/39551/`, `https://www.xbiquge.la/27/27807/`, ``)
		if state {
			l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `xbiquge`)

	list.SourceURL = urlStr

	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// Search 搜索资源
func (r XbiqugeReader) Search(keyword string) (list Catalog, err error) {
	return
}

// GetCatalog 获取章节列表
func (r XbiqugeReader) GetCatalog(urlStr string) (list Catalog, err error) {

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

	// 偷心透视小村医最新章节,
	list.Title = FindString(`(?P<title>(.)+)小说_`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	html2, _ := FindContentHTML(html, `#list`)

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html2))

	var links = GetLinks(g2, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://www.xbiquge.la/39/39551/18263995.html`, `https://www.xbiquge.la/7/7004/3246402.html`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `xbiquge`)

	list.SourceURL = urlStr

	var links2 = GetLinks(g, link)

	list.Next = GetNextLink(links2)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}
	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetInfo 获取详细内容
func (r XbiqugeReader) GetInfo(urlStr string) (ret Content, err error) {

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
	ret.SourceURL = urlStr

	ret.Title = FindString(`(?P<title>(.)+)_(?P<bookname>(.)+)_新笔趣阁`, article.Title, "title")
	if ret.Title == `` {
		ret.Title = article.Title
	}

	reg := regexp.MustCompile(`<a([^<]+)<\/a>`)

	article.ReadContent = reg.ReplaceAllString(article.ReadContent, "")

	reg2 := regexp.MustCompile(`(手机站全新改版升级地址：|http://m.xbiquge.la，数据和书签与电脑站同步，无广告清新阅读！)`)

	article.ReadContent = reg2.ReplaceAllString(article.ReadContent, "")

	c := MarkDownFormatContent(article.ReadContent)

	c = BookContReplace(c)

	ret.Contents = GetSectionByContent(c)

	links, _ := GetLinkByHTML(urlStr, html)
	ret.Previous = GetPreviousLink(links)
	if ret.Previous.URL != `` {
		ret.Previous.URL = EncodeURL(ret.Previous.URL)
	}
	//todo 现在不支持下一页 参数写在JS文件里面用脚本跳转的 (坑爹)
	ret.Next = GetNextLink(links)
	if ret.Next.URL != `` {
		ret.Next.URL = EncodeURL(ret.Next.URL)
	}
	return ret, nil

}
