package reader

import (
	"errors"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/lunny/html2md"
)

//LearnkuReader 常规博客(使用) html2article算法
type LearnkuReader struct {
}

// GetCatalog 获取章节列表
func (r LearnkuReader) GetCatalog(urlStr string) (list Catalog, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, ``)
	if err != nil {
		return
	}

	article, err := GetActicleByHTML(html)
	if err != nil {
		return
	}

	article.Readable(urlStr)

	g, e := goquery.NewDocumentFromReader(strings.NewReader(article.ReadContent))

	if e != nil {
		return list, e
	}

	list.Title = article.Title

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	list.Cards = LinksToCards(Cleaning(links), `/pages/article`, `blog`)

	if len(list.Cards) == 0 {

		html2, _ := FindContentHTML(html, `.tree`)

		g2, e2 := goquery.NewDocumentFromReader(strings.NewReader(html2))

		if e2 != nil {
			return list, e2
		}
		var links2 = GetLinks(g2, link)

		list.Cards = LinksToCards(Cleaning(links2), `/pages/article`, `blog`)

	}

	list.SourceURL = urlStr

	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}
	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetInfo 获取详细内容
func (r LearnkuReader) GetInfo(urlStr string) (ret Content, err error) {

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
	if CheckStrIsLink(urlStr) != nil {
		return ret, errors.New(`url error`)
	}

	ret.Title = article.Title
	ret.Content = article.ReadContent
	ret.Content = html2md.Convert(ret.Content)
	ret.PubAt = string(article.Publishtime)
	ret.SourceURL = urlStr

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
