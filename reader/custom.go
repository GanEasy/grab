package reader

/**
* 完全开放自定义的数据匹配器 (todo)
* 用于解决特殊网站数据转码服务
* 现在未想清楚怎么用
 */
import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

//CustomReader 自定义阅读器
type CustomReader struct {
}

// GetCatalog 获取目录
func (r CustomReader) GetCatalog(urlStr string) (list Catalog, err error) {
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
		// 小程序里面不支持含有html标签内容，一些rss的描述直接是文章正文的，现在暂时不需
		list.Cards = append(list.Cards, Card{item.Title, wxto, ``, `card`, cover, nil})
		// list.Cards = append(list.Cards, Card{item.Title, wxto, item.Description, `card`, item.Image.URL, nil})
	}
	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)
	return
}

// GetInfo 获取详情
func (r CustomReader) GetInfo() {
	fmt.Print(`a read`)
}
