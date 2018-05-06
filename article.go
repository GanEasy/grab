package grab

import (
	"errors"

	"github.com/GanEasy/html2article"
)

// Article 正文
type Article struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	PubAt     string `json:"pub_at"`
	SourceURL string `json:"source_url"` // 数据抓取时，统一声明数据来源
}

// GetArticle 读url中的正文 解释返回 markdown 格式正文
func GetArticle(urlStr string) (cont Content, err error) {

	if CheckStrIsLink(urlStr) != nil {
		return cont, errors.New(`url error`)
	}

	ext, err := html2article.NewFromUrl(urlStr)
	if err != nil {
		return
	}
	article, err := ext.ToArticle()
	if err != nil {
		return
	}
	// fmt.Println(article)

	//parse the article to be readability
	article.Readable(urlStr)

	// fmt.Println(article.Title, article.Publishtime)
	// md = html2md.Convert(article.ReadContent)

	cont.Title = article.Title
	// info.Content = article.ReadContent
	cont.Content = article.ReadContent
	cont.PubAt = string(article.Publishtime)
	cont.SourceURL = urlStr
	return

}

// GetArticleList 获取目录列表
func GetArticleList(urlStr string) (list List, err error) {

	html, err := GetHTML(urlStr, ``)
	if err != nil {
		return
	}

	article, err := GetActicleByHTML(html)
	if err != nil {
		return
	}

	article.Readable(urlStr)

	list.Title = article.Title
	// log.Println(article.ReadContent)
	list.Links, err = GetLinkByHTML(urlStr, article.ReadContent)
	if err != nil {
		return
	}

	list.SourceURL = urlStr

	list.Hash = GetListHash(list)

	return list, nil

}
