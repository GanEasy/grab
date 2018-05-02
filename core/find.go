package reader

import (
	"fmt"
	"strings"

	"github.com/GanEasy/html2article"
	"github.com/PuerkitoBio/goquery"
)

// GetFindArticle 读url中的正文 解释返回 markdown 格式正文
func GetFindArticle(urlStr, findStr string) (article Article, err error) {
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

	if findStr != `` {
		g, err := goquery.NewDocumentFromReader(strings.NewReader(html))

		if err != nil {
			// return
		}
		findHtml, err := g.Find(findStr).Html()
		if err != nil {
			// return
		}

		bh := fmt.Sprintf(`
		<html>
		<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
		<title>%v</title>
		<body>
		%v
		</body>
		</html>
		`, article.Title, findHtml)

		ext2, err := html2article.NewFromReader(strings.NewReader(bh))
		if err != nil {
			// return
		}
		art2, err := ext2.ToArticle()
		if err != nil {
			// return
		}
		art2.Readable(urlStr)
		article.Content = art2.Content

	}

	return article, nil
}
