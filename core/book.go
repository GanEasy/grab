package core

import (
	"fmt"
	"strings"

	"github.com/GanEasy/html2article"
	"github.com/lunny/html2md"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

// BookSection 小说段落 字数
type BookSection struct {
	Text string `json:"text"`
}

// BookInfo 链接
type BookInfo struct {
	URL      string        `json:"url"`
	Title    string        `json:"title"`
	Content  []BookSection `json:"content"`
	PubAt    string        `json:"pub_at"`
	Previous Link          `json:"previous"`
	Next     Link          `json:"next"`
}

// GetBookInfo 获取章节内容详细
func GetBookInfo(url string) (info BookInfo, err error) {
	html, err := GetHTML(url)
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

// GetBookMenu 获取小说目录
func GetBookMenu(urlStr string) (data Data, err error) {
	return GetList(urlStr)
}
