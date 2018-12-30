package reader

import (
	"errors"
	"fmt"

	"github.com/mmcdole/gofeed"
)

//RssReader Rss 订阅工具
type RssReader struct {
}

// GetCatalog 获取章节列表
func (r RssReader) GetCatalog(urlStr string) (list Catalog, err error) {
	fp := gofeed.NewParser()

	// feed, err := fp.ParseString(html)
	feed, err := fp.ParseURL(urlStr)
	if err != nil {
		return
	}
	list.Title = feed.Title
	var wxto, cover, describe string
	drive := `article`
	page := `/pages/article`
	for _, item := range feed.Items {
		// log.Println(item.Description)
		wxto = fmt.Sprintf(`%v?drive=%v&url=%v`, page, drive, EncodeURL(item.Link))
		cover = ``
		if item.Image != nil {
			cover = item.Image.URL
		}
		describe, _ = ReplaceImageServe(item.Description)
		// 小程序里面不支持含有html标签内容，一些rss的描述直接是文章正文的，现在暂时不需
		list.Cards = append(list.Cards, Card{item.Title, wxto, describe, `card`, cover, nil})
		// list.Cards = append(list.Cards, Card{item.Title, wxto, item.Description, `card`, item.Image.URL, nil})
	}
	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)
	return

}

// GetInfo 获取详细内容
func (r RssReader) GetInfo(urlStr string) (ret Content, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, ``)
	if err != nil {
		return ret, err
	}
	// log.Println(html)
	article, err := GetActicleByHTML(html)
	if err != nil {
		return ret, err
	}

	article.Readable(urlStr)
	if CheckStrIsLink(urlStr) != nil {
		return ret, errors.New(`url error`)
	}

	ret.Title = article.Title
	ret.Content = article.ReadContent
	ret.PubAt = string(article.Publishtime)
	ret.SourceURL = urlStr

	links, _ := GetLinkByHTML(urlStr, html)
	ret.Previous = GetPreviousLink(links)
	if ret.Previous.URL != `` {
		ret.Previous.URL = EncodeURL(ret.Previous.URL)
	}
	//todo 现在不支持下一页 参数写在JS文件里面用脚本跳转的 (坑爹)
	ret.Next = GetNextLink(links)
	if ret.Next.URL != `` {
		ret.Next.URL = EncodeURL(ret.Next.URL)
	}
	return ret, nil

}
