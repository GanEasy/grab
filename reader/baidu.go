package reader

import (
	"fmt"
	"net/url"
	"regexp"
	// "strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//BaiduReader U小说阅读网 (盗版小说网站)
type BaiduReader struct {
}

// GetCategories 获取所有分类
func (r BaiduReader) GetCategories(urlStr string) (list Catalog, err error) {
	// urlStr := `http://m.soe8.com/`
	list.Title = `分类-热搜榜`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`全部类型`, `/pages/transfer?action=list&url=` + EncodeURL(`https://top.baidu.com/board?platform=pc&tab=novel`) + `&drive=baidu`, "", `link`, ``, nil, ``, ``},
		Card{`都市`, `/pages/transfer?action=list&url=` + EncodeURL(`https://top.baidu.com/board?platform=pc&tab=novel&tag={"category":"都市"}`) + `&drive=baidu`, "", `link`, ``, nil, ``, ``},
		Card{`玄幻`, `/pages/transfer?action=list&url=` + EncodeURL(`https://top.baidu.com/board?platform=pc&tab=novel&tag={"category":"玄幻"}`) + `&drive=baidu`, "", `link`, ``, nil, ``, ``},
		Card{`历史`, `/pages/transfer?action=list&url=` + EncodeURL(`https://top.baidu.com/board?platform=pc&tab=novel&tag={"category":"历史"}`) + `&drive=baidu`, "", `link`, ``, nil, ``, ``},
		Card{`武侠`, `/pages/transfer?action=list&url=` + EncodeURL(`https://top.baidu.com/board?platform=pc&tab=novel&tag={"category":"武侠"}`) + `&drive=baidu`, "", `link`, ``, nil, ``, ``},
		Card{`现代言情`, `/pages/transfer?action=list&url=` + EncodeURL(`https://top.baidu.com/board?platform=pc&tab=novel&tag={"category":"现代言情"}`) + `&drive=baidu`, "", `link`, ``, nil, ``, ``},
		Card{`古代言情`, `/pages/transfer?action=list&url=` + EncodeURL(`https://top.baidu.com/board?platform=pc&tab=novel&tag={"category":"古代言情"}`) + `&drive=baidu`, "", `link`, ``, nil, ``, ``},
		Card{`青春`, `/pages/transfer?action=list&url=` + EncodeURL(`https://top.baidu.com/board?platform=pc&tab=novel&tag={"category":"青春"}`) + `&drive=baidu`, "", `link`, ``, nil, ``, ``},
	}
	return list, nil
}

// GetList 获取书籍列表列表
func (r BaiduReader) GetList(urlStr string) (list Catalog, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, `.container-bg_lQ801`)
	if err != nil {
		return
	}

	link, _ := url.Parse(urlStr)

	// html2, _ := g.Find(`#detail-list-select`).Eq(1).Html()

	g2, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	var links = GetLinks(g2, link)
	if err != nil {
		return
	}
	// t.Fatal(links)
	var needLinks []Link
	pattern := "\\d+" //反斜杠要转义
	for _, l := range links {
		if l.Title != `查看更多>` {
			match, _ := regexp.MatchString(pattern, l.Title)
			if !match {
				l.URL = fmt.Sprintf(`/pages/search?name=%v`, l.Title)

				needLinks = append(needLinks, l)
			}
		}
	}

	for _, link := range needLinks { //所有链接
		list.Cards = append(list.Cards, Card{link.Title, link.URL, ``, `link`, ``, nil, link.URL, ``})
	}

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// Search 搜索资源
func (r BaiduReader) Search(keyword string) (list Catalog, err error) {
	return
}

// GetCatalog 获取章节列表
func (r BaiduReader) GetCatalog(urlStr string) (list Catalog, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	// html, err := GetHTML(urlStr, ``)
	// if err != nil {
	// 	return
	// }

	return list, err

}

// GetInfo 获取详细内容
func (r BaiduReader) GetInfo(urlStr string) (ret Content, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	// html, err := GetHTML(urlStr, ``)
	// if err != nil {
	// 	return ret, err
	// }
	// // log.Println(html)
	// article, err := GetActicleByHTML(html)
	// if err != nil {
	// 	return ret, err
	// }

	return ret, nil

}
