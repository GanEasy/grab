package grab

import (
	"github.com/mmcdole/gofeed"
)

//RssListReader Rss列表匹配器
type RssListReader struct {
}

// GetList 获取Rss订阅接口文章列表
func (r RssListReader) GetList(urlStr string) (list Catalog, err error) {
	fp := gofeed.NewParser()

	// feed, err := fp.ParseString(html)
	feed, err := fp.ParseURL(urlStr)
	if err != nil {
		return
	}
	list.Title = feed.Title
	var wxto, cover string
	for _, item := range feed.Items {
		// todo
		wxto = item.Link
		cover = ``
		if item.Image != nil {
			cover = item.Image.URL
		}
		// https://img.readfollow.com/file?url=
		// 小程序里面不支持含有html标签内容，一些rss的描述直接是文章正文的，现在暂时不需
		list.Cards = append(list.Cards, Card{item.Title, wxto, item.Description, `card`, cover, nil})
		// list.Cards = append(list.Cards, Card{item.Title, wxto, item.Description, `card`, item.Image.URL, nil})
	}
	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)
	return
}
