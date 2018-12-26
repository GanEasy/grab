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
func GetArticle(urlStr string) (cont FetchContent, err error) {

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

//ArticleListReader 默认列表匹配器
type ArticleListReader struct {
}

// GetList 获取列表
func (r ArticleListReader) GetList(urlStr string) (list Catalog, err error) {
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
	links, err := GetLinkByHTML(urlStr, html)

	list.Cards = LinksToCards(links, ``, `article`)

	if err != nil {
		return
	}

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)
	return
}

//ArticleInfoReader 默认详细页匹配器
type ArticleInfoReader struct {
}

// GetInfo 获取详细内容
func (r ArticleInfoReader) GetInfo(urlStr string) (ret ReaderContent, err error) {

	if CheckStrIsLink(urlStr) != nil {
		return ret, errors.New(`url error`)
	}

	ext, err := html2article.NewFromUrl(urlStr)
	if err != nil {
		return
	}
	article, err := ext.ToArticle()
	if err != nil {
		return
	}

	article.Readable(urlStr)

	ret.Title = article.Title
	ret.Content = article.ReadContent
	ret.PubAt = string(article.Publishtime)
	ret.SourceURL = urlStr
	return
}
