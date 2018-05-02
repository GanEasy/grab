package core

import (
	"github.com/GanEasy/html2article"
)

// GetRssList 读url中的正文 解释返回 markdown 格式正文
func GetRssList(urlStr string) (article Article, err error) {
	html, err := GetHTML(urlStr)
	if err != nil {
		return
	}
	ext, err := html2article.NewFromHtml(html)
	if err != nil {
		return
	}
	art, err := ext.ToArticle()
	if err != nil {
		return
	}
	art.Readable(urlStr)
	article.Title = art.Title
	article.Content = art.ReadContent
	article.PubAt = string(art.Publishtime)
	return article, nil
}
