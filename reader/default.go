package reader

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//DefaultListReader 默认列表匹配器
type DefaultListReader struct {
}

// GetList 获取列表
func (r DefaultListReader) GetList() {
	fmt.Print(`a read`)
}

//DefaultInfoReader 默认详细页匹配器
type DefaultInfoReader struct {
}

// GetInfo 获取详细内容
func (r DefaultInfoReader) GetInfo() {
	fmt.Print(`a read`)
}

// // GetNextURL 获取详细内容
// func (r DefaultInfoReader) GetNextURL() string {
// 	return ``
// }

//DefaultReader 默认阅读器
type DefaultReader struct {
}

// Catalog 获取列表
func (r DefaultReader) Catalog(urlStr string) (list Catalog, err error) {

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

// Info 获取详细内容
func (r DefaultReader) Info(urlStr string) (ret ReaderContent, err error) {

	if err != nil {
		return ret, err
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
