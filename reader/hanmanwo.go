package reader

import (
	"errors"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/lunny/html2md"
)

//HanmanwoReader 绅士漫画网 (盗版漫画网站)
type HanmanwoReader struct {
}

// GetCategories 获取所有分类
func (r HanmanwoReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://m.booktxt.com/`

	list.Title = `分类-韩漫窝`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`全部`, `/pages/list?action=list&drive=hanmanwo&url=` + EncodeURL(`http://www.hanmanwo.com/booklist`), "", `link`, ``, nil, ``, ``},
	}
	return list, nil
}

// GetList 获取书籍列表列表
func (r HanmanwoReader) GetList(urlStr string) (list Catalog, err error) {

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

	list.Title = FindString(`(?P<title>(.)+)-`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `http://www.hanmanwo.com/book/1343`, `http://www.hanmanwo.com/book/1332`, ``)
		if state {
			l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `hanmanwo`)

	list.SourceURL = urlStr

	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// Search 搜索资源
func (r HanmanwoReader) Search(keyword string) (list Catalog, err error) {
	return
}

// GetCatalog 获取章节列表
func (r HanmanwoReader) GetCatalog(urlStr string) (list Catalog, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, `#detail-list-select`)
	if err != nil {
		return
	}

	g, e := goquery.NewDocumentFromReader(strings.NewReader(html))

	if e != nil {
		return list, e
	}

	list.Title = FindString(`-(?P<title>(.)+)漫画全集免费`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	// html2, _ := g.Find(`#detail-list-select`).Eq(1).Html()

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html))

	var links = GetLinks(g2, link)

	var needLinks Links
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `http://www.hanmanwo.com/chapter/101669`, `http://www.hanmanwo.com/chapter/101667`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	// sort.Stable(needLinks)
	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/article`, `hanmanwo`)

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
func (r HanmanwoReader) GetInfo(urlStr string) (ret Content, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, `.comicpage`)
	if err != nil {
		return ret, err
	}
	// log.Println(html)

	// html2, err := FindContentHTML(html, `#content`)
	// g2, e2 := goquery.NewDocumentFromReader(strings.NewReader(html2))
	// html2article
	article, err := GetActicleForHTML(html)
	if err != nil {
		return ret, err
	}

	article.Readable(urlStr)
	if CheckStrIsLink(urlStr) != nil {
		return ret, errors.New(`url error`)
	}

	ret.Title = FindString(`(?P<title>(.)+)在线阅读-韩漫窝`, article.Title, "title")

	// ret.Content = article.ReadContent

	ret.Content = ImagesBuildHTML(article.Images)
	ret.Content = html2md.Convert(ret.Content)
	// ret.PubAt = string(article.Publishtime)
	ret.SourceURL = urlStr

	links, _ := GetLinkByHTML(urlStr, html)
	// log.Println(`article.Images`, links, ImagesBuildHTML(article.Images))
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
