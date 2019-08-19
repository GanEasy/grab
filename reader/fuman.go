package reader

import (
	"errors"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//FumanReader 顶点小说 (盗版小说网站)
type FumanReader struct {
}

// GetCategories 获取所有分类
func (r FumanReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://m.booktxt.com/`

	list.Title = `分类-我爱妹子漫画`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`全部`, `/pages/list?action=list&drive=fuman&url=` + EncodeURL(`https://5aimeizi.com/booklist`), "", `link`, ``, nil, ``},
		Card{`伦理`, `/pages/list?action=list&drive=fuman&url=` + EncodeURL(`https://www.fuman.cc/booklist?tag=%E4%BC%A6%E7%90%86&area=-1&end=-1`), "", `link`, ``, nil, ``},
		Card{`恋爱`, `/pages/list?action=list&drive=fuman&url=` + EncodeURL(`https://www.fuman.cc/booklist?tag=%E6%81%8B%E7%88%B1&area=-1&end=-1`), "", `link`, ``, nil, ``},
		Card{`情感`, `/pages/list?action=list&drive=fuman&url=` + EncodeURL(`https://www.fuman.cc/booklist?tag=%E6%83%85%E6%84%9F&area=-1&end=-1`), "", `link`, ``, nil, ``},
		Card{`OL`, `/pages/list?action=list&drive=fuman&url=` + EncodeURL(`https://www.fuman.cc/booklist?tag=OL&area=-1&end=-1`), "", `link`, ``, nil, ``},
		Card{`暧昧`, `/pages/list?action=list&drive=fuman&url=` + EncodeURL(`https://www.fuman.cc/booklist?tag=%E6%9A%A7%E6%98%A7&area=-1&end=-1`), "", `link`, ``, nil, ``},
		Card{`韩国`, `/pages/list?action=list&drive=fuman&url=` + EncodeURL(`https://www.fuman.cc/booklist?tag=%E5%85%A8%E9%83%A8&area=1&end=-1`), "", `link`, ``, nil, ``},
	}
	return list, nil
}

// GetList 获取书籍列表列表
func (r FumanReader) GetList(urlStr string) (list Catalog, err error) {

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

	list.Title = FindString(`免费韩漫列表-(?P<title>(.)+)`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://www.fuman.cc/book/400`, `https://www.fuman.cc/book/25`, ``)
		if state {
			l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `fuman`)

	list.SourceURL = urlStr

	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetCatalog 获取章节列表
func (r FumanReader) GetCatalog(urlStr string) (list Catalog, err error) {

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

	list.Title = FindString(`(?P<title>(.)+)全集`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	// html2, _ := g.Find(`#detail-list-select`).Eq(1).Html()

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html))

	var links = GetLinks(g2, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://www.fuman.cc/chapter/1011`, `https://www.fuman.cc/chapter/1018`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/article`, `fuman`)

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
func (r FumanReader) GetInfo(urlStr string) (ret Content, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, ``)
	if err != nil {
		return ret, err
	}
	// log.Println(html)

	html2, err := FindContentHTML(html, `.comicpage`)
	// g2, e2 := goquery.NewDocumentFromReader(strings.NewReader(html2))
	// html2article
	article, err := GetActicleForHTML(html2)
	if err != nil {
		return ret, err
	}

	article.Readable(urlStr)
	if CheckStrIsLink(urlStr) != nil {
		return ret, errors.New(`url error`)
	}

	ret.Title = FindString(`(?P<title>(.)+)免费阅读-腐漫漫画`, article.Title, "title")

	// ret.Content = article.ReadContent

	ret.Content = ImagesBuildHTML(article.Images)
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
