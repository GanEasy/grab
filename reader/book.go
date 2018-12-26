package reader

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/GanEasy/html2article"
	"github.com/PuerkitoBio/goquery"
	"github.com/lunny/html2md"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

//GetLinkByHTML 获取网页内容所有链接
func GetLinkByHTML(urlStr, html string) (links []Link, err error) {
	// 没有 html标签 或者 body 标签可能出现文档解释异常
	if !strings.Contains(html, `</html>`) || !strings.Contains(html, `</body>`) {
		html = fmt.Sprintf(`
			<html>
			<head>
			<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
			<title>%v</title>
			<body>
			%v
			</body>
			</html>
			`, `NONE TITLE`, html)
	}

	c := strings.NewReader(html)

	g, err := goquery.NewDocumentFromReader(c)

	if err != nil {
		return
	}
	link, err := url.Parse(urlStr)
	if err != nil {
		return
	}
	links = GetLinks(g, link)
	return
}

//CheckStrIsLink 检查字符串是否支持的链接
func CheckStrIsLink(urlStr string) error {

	link, err := url.Parse(urlStr)

	if err != nil {
		return err
	}

	if link.Scheme == "" {
		return errors.New("Scheme Fatal")
	}

	if link.Host == "" {
		return errors.New("Host Fatal")
	}
	return nil
}

//GetPreviousLink 获取上一页或者上一章
func GetPreviousLink(links []Link) Link {
	for _, link := range links {
		if strings.Contains(link.Title, `上一页`) || strings.Contains(link.Title, `上一章`) || strings.Contains(link.Title, `上一`) {
			return Link{Title: "previous", URL: link.URL}
		}
	}
	return Link{}
}

//GetNextLink 获取下一页或者下一章
func GetNextLink(links []Link) Link {
	for _, link := range links {
		if strings.Contains(link.Title, `下一页`) || strings.Contains(link.Title, `下页`) || strings.Contains(link.Title, `下一章`) || strings.Contains(link.Title, `下章`) {
			return Link{Title: "next", URL: link.URL}
		}
	}
	return Link{}
}

//GetActicleByHTML 由Html返回*html2article.Article
func GetActicleByHTML(html string) (article *html2article.Article, err error) {
	ext, err := html2article.NewFromHtml(html)
	if err != nil {
		return
	}

	return ext.ToArticle()
}

// GetSectionByContent 通过正文获取段落拆分
func GetSectionByContent(content string) (sec []string) {
	// 替换换行符
	content = BookContReplace(content)
	// 拆分换行符
	arr := strings.Split(content, "</p>")
	if len(arr) > 1 {
		for _, v := range arr {
			text := strings.TrimSpace(v)
			if text != "" {
				// 不为空时组装段落
				sec = append(sec, text)
			}
		}
	}
	return
}

//MarkDownFormatContent 通过markdown语法格式化内容
func MarkDownFormatContent(content string) string {
	md := html2md.Convert(content)
	input := []byte(md)
	unsafe := blackfriday.MarkdownCommon(input)
	contentBytes := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return strings.TrimSpace(fmt.Sprintf(`%v`, string(contentBytes[:])))
}

// BookContReplace 小说内容正文替换标签
func BookContReplace(html string) string {
	c := strings.Replace(html, `<p>`, ``, -1)
	c = strings.Replace(c, `<code>`, ``, -1)
	c = strings.Replace(c, `</code>`, ``, -1)
	c = strings.Replace(c, `<pre>`, ``, -1)
	c = strings.Replace(c, `</pre>`, ``, -1)

	c = strings.Replace(c, `<br/>`, `</p>`, -1)
	c = strings.Replace(c, `<br />`, `</p>`, -1)
	c = strings.Replace(c, `<br>`, `</p>`, -1)
	c = strings.Replace(c, "\n", `</p>`, -1)
	return c
}

// GetBookChapters 获取目录列表
func GetBookChapters(urlStr string) (list Catalog, err error) {
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
	// log.Println(links)
	list.Links = Cleaning(links)

	list.SourceURL = urlStr

	list.Hash = GetListHash(list)

	return list, nil

}

// // GetBookMenu 获取小说目录
// func GetBookMenu(urlStr string) (data Data, err error) {
// 	return GetList(urlStr)
// }

//BookListReader 默认列表匹配器
type BookListReader struct {
}

// GetList 获取列表
func (r BookListReader) GetList(urlStr string) (list Catalog, err error) {

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

//BookInfoReader 默认详细页匹配器
type BookInfoReader struct {
}

// GetInfo 获取详细内容
func (r BookInfoReader) GetInfo(urlStr string) (ret Content, err error) {

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

//DefaultBookReader 默认详细页匹配器
type DefaultBookReader struct {
}

// GetCategories 获取列表
func (r DefaultBookReader) GetCategories(urlStr string) (list Catalog, err error) {

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

	list.Cards = LinksToCards(Cleaning(links), `/pages/book/get`, `book`)

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetBooks 获取列表
func (r DefaultBookReader) GetBooks(urlStr string) (list Catalog, err error) {

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

	list.Cards = LinksToCards(Cleaning(links), `/pages/chapter/get`, `book`)

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetChapters 获取列表
func (r DefaultBookReader) GetChapters(urlStr string) (list Catalog, err error) {

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

// GetChapter 获取详细内容
func (r DefaultBookReader) GetChapter(urlStr string) (ret Content, err error) {

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
