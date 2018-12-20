package grab

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//BooktxtReader 顶点小说 (盗版小说网站)
type BooktxtReader struct {
}

// GetCategories 获取所有分类
func (r BooktxtReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://m.booktxt.com/`

	list.Title = `分类-顶点小说`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`全部`, `/pages/transfer/list?action=book&drive=booktxt&url=` + EncodeURL(`https://m.booktxt.net/wapsort/0_1.html`), "", `link`, ``, nil},
		Card{`玄幻`, `/pages/transfer/list?action=book&drive=booktxt&url=` + EncodeURL(`https://m.booktxt.net/wapsort/1_1.html`), "", `link`, ``, nil},
		Card{`修真`, `/pages/transfer/list?action=book&drive=booktxt&url=` + EncodeURL(`https://m.booktxt.net/wapsort/2_1.html`), "", `link`, ``, nil},
		Card{`都市`, `/pages/transfer/list?action=book&drive=booktxt&url=` + EncodeURL(`https://m.booktxt.net/wapsort/3_1.html`), "", `link`, ``, nil},
		Card{`穿越`, `/pages/transfer/list?action=book&drive=booktxt&url=` + EncodeURL(`https://m.booktxt.net/wapsort/4_1.html`), "", `link`, ``, nil},
		Card{`网游`, `/pages/transfer/list?action=book&drive=booktxt&url=` + EncodeURL(`https://m.booktxt.net/wapsort/5_1.html`), "", `link`, ``, nil},
		Card{`科幻`, `/pages/transfer/list?action=book&drive=booktxt&url=` + EncodeURL(`https://m.booktxt.net/wapsort/6_1.html`), "", `link`, ``, nil},
		Card{`其他`, `/pages/transfer/list?action=book&drive=booktxt&url=` + EncodeURL(`https://m.booktxt.net/wapsort/7_1.html`), "", `link`, ``, nil},
	}
	return list, nil
}

// GetBooks 获取书籍列表列表
func (r BooktxtReader) GetBooks(urlStr string) (list Catalog, err error) {

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

	list.Title = FindString(`(?P<title>(.)+)_顶点小说`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://m.booktxt.net/wapbook/10441.html`, `https://m.booktxt.net/wapbook/9643.html`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/chapter/get`, `booktxt`)

	list.SourceURL = urlStr

	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetChapters 获取章节列表
func (r BooktxtReader) GetChapters(urlStr string) (list Catalog, err error) {

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

	list.Title = FindString(`(?P<title>(.)+)最新章节`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	html2, _ := g.Find(`.directoryArea`).Eq(1).Html()

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html2))

	var links = GetLinks(g2, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://m.booktxt.net/wapbook/9643_3668840.html`, `https://m.booktxt.net/wapbook/10170_4878581.html`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/chapter/info`, `booktxt`)

	list.SourceURL = urlStr

	var links2 = GetLinks(g, link)

	list.Next = GetNextLink(links2)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}
	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetChapter 获取详细内容
func (r BooktxtReader) GetChapter(urlStr string) (ret TextContent, err error) {

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

	ret.Title = FindString(`(?P<title>(.)+)_(?P<bookname>(.)+)_(?P<category>(.)+)_`, article.Title, "title")
	if ret.Title == `` {
		ret.Title = article.Title
	}

	ret.SourceURL = urlStr

	c := MarkDownFormatContent(article.ReadContent)

	c = BookContReplace(c)

	ret.Content = GetSectionByContent(c)

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
