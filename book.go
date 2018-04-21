package grab

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/GanEasy/grab/core"
	"github.com/GanEasy/html2article"
	"github.com/PuerkitoBio/goquery"
	"github.com/lunny/html2md"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

// GetHTML 获取rss链接地址中的链接
func GetBookList(urlStr, find string) (html string, err error) {
	html, err = core.GetHTML(urlStr)
	return
}

// GetBookInfo 获取章节内容详细
func GetBookInfo(url string) (info Book, err error) {
	html, err := GetHTML(url, ``)
	if err != nil {
		return info, err
	}
	// log.Println(html)
	article, err := GetActicleByHTML(html)
	if err != nil {
		return info, err
	}

	article.Readable(url)

	info.Title = article.Title
	info.URL = url

	c := MarkDownFormatContent(article.ReadContent)

	c = BookContReplace(c)

	info.Content = GetSectionByContent(c)

	links, _ := GetLinkByHTML(html)
	info.Previous = GetPreviousLink(links)
	info.Next = GetNextLink(links)
	// info.PubAt = Publishtime
	return info, nil
}

//GetLinkByHTML 获取网页内容所有链接
func GetLinkByHTML(html string) (links []Link, err error) {
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

	// fmt.Println(g.Text())
	g.Find("a").Each(func(i int, content *goquery.Selection) {
		n := strings.TrimSpace(content.Text())
		u, _ := content.Attr("href")
		if err := CheckStrIsLink(u); err == nil {
			links = append(links, Link{
				n,
				u,
			})
		}
	})
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
		if strings.Contains(link.Title, `上一页`) || strings.Contains(link.Title, `上一章`) {
			return Link{Title: "previous", URL: link.URL}
		}
	}
	return Link{}
}

//GetNextLink 获取下一页或者下一章
func GetNextLink(links []Link) Link {
	for _, link := range links {
		if strings.Contains(link.Title, `下一页`) || strings.Contains(link.Title, `下一章`) {
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
func GetSectionByContent(content string) (sec []BookSection) {
	// 替换换行符
	content = BookContReplace(content)
	// 拆分换行符
	arr := strings.Split(content, "</p>")
	if len(arr) > 1 {
		for _, v := range arr {
			text := strings.TrimSpace(v)
			if text != "" {
				// 不为空时组装段落
				sec = append(sec, BookSection{
					Text: text,
				})
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
	c = strings.Replace(c, `\n`, `</p>`, -1)
	return c
}

// // GetBookMenu 获取小说目录
// func GetBookMenu(urlStr string) (data Data, err error) {
// 	return GetList(urlStr)
// }
