package grab

import "github.com/mmcdole/gofeed"

// GetRssList 获取rss链接地址中的链接
func GetRssList(urlStr string) (list List, err error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(urlStr)
	if err != nil {
		return
	}
	list.Title = feed.Title
	for _, item := range feed.Items {
		list.Links = append(list.Links, Link{item.Title, item.Link})
	}
	list.SourceURL = urlStr
	return
}
