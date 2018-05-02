package grab

import (
	"errors"

	"github.com/GanEasy/html2article"
)

// Content 正文
type Content struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	PubAt     string `json:"pub_at"`
	SourceURL string `json:"source_url"` // 数据抓取时，统一声明数据来源
}

// GetContent 读url中的正文 解释返回 markdown 格式正文
func GetContent(urlStr string) (cont Content, err error) {

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
