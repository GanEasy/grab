package grab

import (
	"net/url"
	"strings"

	"github.com/GanEasy/grab/core"
	"github.com/PuerkitoBio/goquery"
)

// GetHTMLList 获取指定URL指定元素里面的URL(所有)
func GetHTMLList(urlStr, find string) (list List, err error) {
	html, err := core.GetHTML(urlStr)
	if err != nil {
		return
	}

	html, err = FindContentHTML(html, find)

	if err != nil {
		return
	}
	// log.Println(html)

	g, e := goquery.NewDocumentFromReader(strings.NewReader(html))

	list.Title = g.Find("title").Eq(0).Text()
	if e != nil {
		return
	}

	link, _ := url.Parse(urlStr)

	// if err != nil {
	// 	return
	// }

	links := core.GetLinks(g, link)

	for _, l := range links {
		list.Links = append(list.Links, Link{
			l.Title,
			l.URL,
		})
	}

	list.Hash = GetListHash(list)
	return
}
