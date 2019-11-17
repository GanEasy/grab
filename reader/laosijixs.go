package reader

import (
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//LaosijixsReader 顶点小说 (盗版小说网站)
type LaosijixsReader struct {
}

// GetCategories 获取所有分类
func (r LaosijixsReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://m.laosijixs.com/`

	list.Title = `分类-老司机小说`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`全部`, `/pages/list?action=list&drive=laosijixs&url=` + EncodeURL(`http://m.laosijixs.com/shuku/`), "", `link`, ``, nil, ``},
	}
	return list, nil
}

// GetList 获取书籍列表列表
func (r LaosijixsReader) GetList(urlStr string) (list Catalog, err error) {

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

	list.Title = `资源列表-老司机小说`

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `http://m.laosijixs.com/20/20961/`, `http://www.laosijixs.com/19/19634/`, ``)
		if state {
			// l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `laosijixs`)

	list.SourceURL = urlStr

	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetCatalog 获取章节列表
func (r LaosijixsReader) GetCatalog(urlStr string) (list Catalog, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, ``)
	// html, err := GetHTMLByChromedp(urlStr)
	if err != nil {
		return
	}

	g, e := goquery.NewDocumentFromReader(strings.NewReader(html))

	if e != nil {
		return list, e
	}

	list.Title = FindString(`(?P<title>(.)+)_全文阅读`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	html2, _ := g.Find(`.chapter-list`).Eq(1).Html()

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html2))

	var links = GetLinks(g2, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `http://m.laosijixs.com/20/20961/546047.html`, `http://m.laosijixs.com/79/79525/5713401.html`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `laosijixs`)

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
func (r LaosijixsReader) GetInfo(urlStr string) (ret Content, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, ``)
	// html, err := GetHTMLByChromedp(urlStr)
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
	//content

	content3, err := FindContentForHTML(html, `#content`)
	// g2, e2 := goquery.NewDocumentFromReader(strings.NewReader(html2))
	// html2article
	// log.Println(`content3`, content3)

	if content3 != `` {
		article.ReadContent = content3
	}

	reg := regexp.MustCompile(`<span([^>]+)>([^<]+)<\/span>`)
	article.ReadContent = reg.ReplaceAllString(article.ReadContent, "")

	reg2 := regexp.MustCompile(`努力加载中...超过5秒钟未打开,请刷新一下！`)

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

	g, e := goquery.NewDocumentFromReader(strings.NewReader(html))
	if e != nil {
		//
	}
	html2, _ := g.Find(`.chapterPages`).Html()

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html2))

	link, _ := url.Parse(urlStr)
	var links2 = GetLinks(g2, link)

	var needLinks []Link
	var state bool
	for _, l := range links2 {
		l.URL, state = JaccardMateGetURL(l.URL, `http://m.laosijixs.com/20/20961/546056_1.html`, `http://m.laosijixs.com/80/80894/5905659_2.html`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}
	var thisPage = 0
	if len(needLinks) > 1 {
		for i, l := range needLinks {
			if thisPage == 0 && l.URL == urlStr {
				thisPage = i
				log.Println(`thisPage`, thisPage)
			}
		}
	}

	if len(needLinks) > (thisPage + 1) {
		ret.Next.URL = EncodeURL(needLinks[thisPage+1].URL)
	} else {
		ret.Next = GetNextLink(links)
		if ret.Next.URL != `` {
			ret.Next.URL = EncodeURL(ret.Next.URL)
		}
	}

	return ret, nil

}
