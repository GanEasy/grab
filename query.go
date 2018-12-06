package grab

/**
* 类似JQ用法的开放性转码器
* 用于实现通过#id或.class获取特定数据服务
* 未想好具体需要实现哪些属性和层级
 */
import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

//QueryListReader Rss列表匹配器
type QueryListReader struct {
}

// GetList 获取Rss订阅接口文章列表
func (r QueryListReader) GetList(urlStr string) (list Catalog, err error) {
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

//QueryInfoReader 默认详细页匹配器
type QueryInfoReader struct {
}

// GetInfo 获取详细内容
func (r QueryInfoReader) GetInfo() {
	fmt.Print(`a read`)
}
