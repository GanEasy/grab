package reader

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//BiqugeinfoReader 顶点小说 (盗版小说网站)
type BiqugeinfoReader struct {
}

// GetCategories 获取所有分类
func (r BiqugeinfoReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://m.biqugeinfo.com/`

	list.Title = `分类-biquge.info`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`玄幻`, `/pages/list?action=list&drive=biqugeinfo&url=` + EncodeURL(`https://m.biquge.info/list/1_1.html`), "", `link`, ``, nil, ``},
		Card{`修真`, `/pages/list?action=list&drive=biqugeinfo&url=` + EncodeURL(`https://m.biquge.info/list/2_1.html`), "", `link`, ``, nil, ``},
		Card{`都市`, `/pages/list?action=list&drive=biqugeinfo&url=` + EncodeURL(`https://m.biquge.info/list/3_1.html`), "", `link`, ``, nil, ``},
		Card{`穿越`, `/pages/list?action=list&drive=biqugeinfo&url=` + EncodeURL(`https://m.biquge.info/list/4_1.html`), "", `link`, ``, nil, ``},
		Card{`网游`, `/pages/list?action=list&drive=biqugeinfo&url=` + EncodeURL(`https://m.biquge.info/list/5_1.html`), "", `link`, ``, nil, ``},
		Card{`女生`, `/pages/list?action=list&drive=biqugeinfo&url=` + EncodeURL(`https://m.biquge.info/list/6_1.html`), "", `link`, ``, nil, ``},

		Card{`\日点击榜`, `/pages/list?action=list&drive=biqugeinfo&url=` + EncodeURL(`https://m.biquge.info/paihangbang_dayvisit/1.html`), "", `link`, ``, nil, ``},
		Card{`\周点击榜`, `/pages/list?action=list&drive=biqugeinfo&url=` + EncodeURL(`https://m.biquge.info/paihangbang_weekvisit/1.html`), "", `link`, ``, nil, ``},
		Card{`\月点击榜`, `/pages/list?action=list&drive=biqugeinfo&url=` + EncodeURL(`https://m.biquge.info/paihangbang_monthvisit/1.html`), "", `link`, ``, nil, ``},
	}
	return list, nil
}

// GetList 获取书籍列表列表
func (r BiqugeinfoReader) GetList(urlStr string) (list Catalog, err error) {

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

	list.Title = FindString(`(?P<title>(.)+)_好看的`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://m.biquge.info/71_71389/`, `https://m.biquge.info/1_1760/`, ``)
		if state {
			// l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `biqugeinfo`)

	list.SourceURL = urlStr

	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetCatalog 获取章节列表
func (r BiqugeinfoReader) GetCatalog(urlStr string) (list Catalog, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTMLOrCache(urlStr, ``)
	if err != nil {
		return
	}

	g, e := goquery.NewDocumentFromReader(strings.NewReader(html))

	if e != nil {
		return list, e
	}

	list.Title = FindString(`(?P<title>(.)+)最新章节列表`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	html2, _ := g.Find(`.chapter`).Eq(1).Html()

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html2))

	var links = GetLinks(g2, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://m.biquge.info/10_10218/5001515.html`, `https://m.biquge.info/68_68619/12705323.html`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `biqugeinfo`)

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
func (r BiqugeinfoReader) GetInfo(urlStr string) (ret Content, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTMLOrCache(urlStr, ``)
	if err != nil {
		return ret, err
	}
	// log.Println(html)
	article, err := GetActicleByHTML(html)
	if err != nil {
		return ret, err
	}

	article.Readable(urlStr)

	ret.Title = FindString(`(?P<title>(.)+)_(?P<bookname>(.)+)_笔趣阁手机版`, article.Title, "title")
	if ret.Title == `` {
		ret.Title = article.Title
	}

	reg := regexp.MustCompile(`<a([^>]+)>([^<]+)<\/a>`)

	article.ReadContent = reg.ReplaceAllString(article.ReadContent, "")

	reg2 := regexp.MustCompile(`本章未完，请点击下一页继续阅读....`)

	article.ReadContent = reg2.ReplaceAllString(article.ReadContent, "")

	ret.SourceURL = urlStr

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
