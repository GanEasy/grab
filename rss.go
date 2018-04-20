package grab

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"

	"github.com/GanEasy/grab/core"
	"github.com/mmcdole/gofeed"
)

// GetRssList 获取rss链接地址中的链接
func GetRssList(urlStr string) (list List, err error) {
	html, err := core.GetHTML(urlStr)
	if err != nil {
		return
	}
	fp := gofeed.NewParser()
	feed, err := fp.ParseString(html)
	if err != nil {
		return
	}
	list.Title = feed.Title
	for _, item := range feed.Items {
		list.Links = append(list.Links, Link{item.Title, item.Link})
	}
	list.SourceURL = urlStr

	// hashByte := uintptr(unsafe.Pointer(list))

	list.Hash = GetListHash(list)
	return
}

//GetHash 获取hash
func GetHash(s string) string {
	h := sha1.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

//GetListHash get
func GetListHash(list List) string {
	var buf bytes.Buffer
	buf.WriteString(list.Title)
	for _, link := range list.Links {
		buf.WriteString(link.Title)
		buf.WriteString(link.URL)
	}
	buf.WriteString(list.SourceURL)
	return GetHash(buf.String())
}
