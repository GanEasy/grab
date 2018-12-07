package grab

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//ZonghengReader 纵横小说网
type ZonghengReader struct {
}

// GetCategories 获取所有分类
func (r ZonghengReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://book.zongheng.com`

	list.Title = `纵横中文网`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`奇幻玄幻`, `/pages/book/get?drive=zongheng&url=` + EncodeURL(`http://book.zongheng.com/store/c1/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", `link`, ``, nil},
		Card{`武侠仙侠`, `/pages/book/get?drive=zongheng&url=` + EncodeURL(`http://book.zongheng.com/store/c3/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", `link`, ``, nil},
		Card{`历史军事`, `/pages/book/get?drive=zongheng&url=` + EncodeURL(`http://book.zongheng.com/store/c6/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", `link`, ``, nil},
		Card{`都市娱乐`, `/pages/book/get?drive=zongheng&url=` + EncodeURL(`http://book.zongheng.com/store/c9/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", `link`, ``, nil},
		Card{`科幻游戏`, `/pages/book/get?drive=zongheng&url=` + EncodeURL(`http://book.zongheng.com/store/c15/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", `link`, ``, nil},
		Card{`悬疑灵异`, `/pages/book/get?drive=zongheng&url=` + EncodeURL(`http://book.zongheng.com/store/c18/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", `link`, ``, nil},
		Card{`竞技同人`, `/pages/book/get?drive=zongheng&url=` + EncodeURL(`http://book.zongheng.com/store/c21/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", `link`, ``, nil},
		Card{`评论文集`, `/pages/book/get?drive=zongheng&url=` + EncodeURL(`http://book.zongheng.com/store/c24/c0/b0/u6/p1/v9/s9/t0/u0/i1/ALL.html`), "", `link`, ``, nil},
		Card{`二次元`, `/pages/book/get?drive=zongheng&url=` + EncodeURL(`http://book.zongheng.com/store/c40/c0/b0/u0/p1/v0/s9/t0/u0/i1/ALL.html`), "", `link`, ``, nil},
	}
	return list, nil
}

// GetBooks 获取书籍列表列表
func (r ZonghengReader) GetBooks(urlStr string) (list Catalog, err error) {

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
		l.URL, state = JaccardMateGetURL(l.URL, `http://book.zongheng.com/book/769150.html`, `http://book.zongheng.com/book/316562.html`, `http://book.zongheng.com/showchapter/769150.html`)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/chapter/get`, `book`)

	list.SourceURL = urlStr

	// html := `{"I":"5333","V":"马经理"},`
	// page := FindString(`/p(?P<page>[^"]+)/`, html, "page")

	page := FindString(`/p(?P<page>(\d)+)/`, html, "page")
	p, err := strconv.Atoi(page)
	if p > 0 && err == nil {
		// 已经组装url
		nextURL := strings.Replace(urlStr, fmt.Sprintf(`/p%v/`, p), fmt.Sprintf(`/p%v/`, p+1), -1)
		list.Next = Link{`下一页`, EncodeURL(nextURL), ``}
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetList 获取列表
func (r ZonghengReader) GetList(urlStr string) (list Catalog, err error) {

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

	list.Cards = LinksToCards(Cleaning(links), `/pages/chapter/info`, `book`)

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetInfo 获取详细内容
func (r ZonghengReader) GetInfo(urlStr string) (ret TextContent, err error) {

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
	ret.SourceURL = urlStr

	c := MarkDownFormatContent(article.ReadContent)

	c = BookContReplace(c)

	ret.Content = GetSectionByContent(c)

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
