package reader

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//XxsyReader 纵横小说网
type XxsyReader struct {
}

// GetCategories 获取所有自定义分类(写死)
func (r XxsyReader) GetCategories(urlStr string) (list Catalog, err error) {

	list.Title = `分类-潇湘书院`

	list.SourceURL = urlStr

	list.Cards = []Card{
		Card{`全部`, `/pages/list?drive=xxsy&url=` + EncodeURL(`https://www.xxsy.net/search?s_wd=&channel=2&sort=9&pn=1`), "", `link`, ``, nil, ``},
		Card{`玄幻言情`, `/pages/list?drive=xxsy&url=` + EncodeURL(`https://www.xxsy.net/search?s_wd=&s_type=3&channel=2&sort=9&pn=1`), "", `link`, ``, nil, ``},
		Card{`仙侠奇缘`, `/pages/list?drive=xxsy&url=` + EncodeURL(`https://www.xxsy.net/search?s_wd=&s_type=4&channel=2&sort=9&pn=1`), "", `link`, ``, nil, ``},
		Card{`古代言情`, `/pages/list?drive=xxsy&url=` + EncodeURL(`https://www.xxsy.net/search?s_wd=&s_type=1&channel=2&sort=9&pn=1`), "", `link`, ``, nil, ``},
		Card{`现代言情`, `/pages/list?drive=xxsy&url=` + EncodeURL(`https://www.xxsy.net/search?s_wd=&s_type=2&channel=2&sort=9&pn=1`), "", `link`, ``, nil, ``},
		Card{`浪漫青春`, `/pages/list?drive=xxsy&url=` + EncodeURL(`https://www.xxsy.net/search?s_wd=&s_type=6&channel=2&sort=9&pn=1`), "", `link`, ``, nil, ``},
		Card{`悬疑`, `/pages/list?drive=xxsy&url=` + EncodeURL(`https://www.xxsy.net/search?s_wd=&s_type=5&channel=2&sort=9&pn=1`), "", `link`, ``, nil, ``},
		Card{`科幻空间`, `/pages/list?drive=xxsy&url=` + EncodeURL(`https://www.xxsy.net/search?s_wd=&s_type=7&channel=2&sort=9&pn=1`), "", `link`, ``, nil, ``},
		Card{`游戏竞技`, `/pages/list?drive=xxsy&url=` + EncodeURL(`https://www.xxsy.net/search?s_wd=&s_type=8&channel=2&sort=9&pn=1`), "", `link`, ``, nil, ``},
		Card{`轻小说`, `/pages/list?drive=xxsy&url=` + EncodeURL(`https://www.xxsy.net/search?s_wd=&s_type=9&channel=2&sort=9&pn=1`), "", `link`, ``, nil, ``},
		Card{`短篇`, `/pages/list?drive=xxsy&url=` + EncodeURL(`https://www.xxsy.net/search?s_wd=&s_type=11&channel=2&sort=9&pn=1`), "", `link`, ``, nil, ``},
	}
	list.Hash = GetCatalogHash(list)
	return list, nil
}

// GetList 获取分类书籍列表
func (r XxsyReader) GetList(urlStr string) (list Catalog, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, ``)
	if err != nil {
		return
	}

	g, e := goquery.NewDocumentFromReader(strings.NewReader(html))

	if e != nil {
		return list, e
	}

	list.Title = `书籍列表-潇湘书院`

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://www.xxsy.net/info/1079349.html`, `https://www.xxsy.net/info/1104267.html`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `xxsy`)

	list.SourceURL = urlStr

	// html := `{"I":"5333","V":"马经理"},`
	// page := FindString(`/p(?P<page>[^"]+)/`, html, "page")
	// class='page-next' href="javascript:" onclick="setPage(3)"
	page := FindString(`class='page-next' href="javascript:" onclick="setPage((?P<page>(\d)+))"`, urlStr, "page")
	p, err := strconv.Atoi(page)
	if p > 1 && p < 6 && err == nil {
		// 已经组装url
		nextURL := strings.Replace(urlStr, fmt.Sprintf(`&pn=%v`, p-1), fmt.Sprintf(`&pn=%v`, p), -1)
		list.Next = Link{`下一页`, EncodeURL(nextURL), ``}
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetCatalog 获取章节列表
func (r XxsyReader) GetCatalog(urlStr string) (list Catalog, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, ``)
	if err != nil {
		return
	}

	g, e := goquery.NewDocumentFromReader(strings.NewReader(html))

	if e != nil {
		return list, e
	}

	list.Title = FindString(`,(?P<title>(.)+)全文阅读,`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)
	var needLinks []Link
	if len(links) > 0 {

		bookID := FindString(`/info/(?P<id>(\d)+)`, urlStr, "id")

		if bookID != `` {
			links, _ = r.GetChaptersLinksByHTML(bookID)
			needLinks = links
		}

	} else {

		var state bool
		for _, l := range links { //
			l.URL, state = JaccardMateGetURL(l.URL, `https://www.xxsy.net/chapter/80232498.html`, `https://www.xxsy.net/chapter/80342483.html`, ``)
			if state {
				needLinks = append(needLinks, l)
			}
		}

	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `xxsy`)

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetInfo 获取章节正文内容
func (r XxsyReader) GetInfo(urlStr string) (ret Content, err error) {

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

	ret.Title = article.Title

	ret.Title = FindString(`_(?P<title>(.)+)_全文阅读`, article.Title, "title")
	if ret.Title == `` {
		ret.Title = article.Title
	}

	ret.SourceURL = urlStr

	c := MarkDownFormatContent(article.ReadContent)

	c = BookContReplace(c)

	ret.Contents = GetSectionByContent(c)

	links, _ := GetLinkByHTML(urlStr, html)
	ret.Previous = GetPreviousLink(links)
	if ret.Previous.URL != `` {
		ret.Previous.URL = EncodeURL(ret.Previous.URL)
	}
	ret.Next = GetNextLink(links)
	if ret.Next.URL != `` {
		ret.Next.URL = EncodeURL(ret.Next.URL)
	}
	return ret, nil

}

// GetChaptersLinksByHTML 获取章节链接列表
func (r XxsyReader) GetChaptersLinksByHTML(bookID string) (links []Link, err error) {

	//
	urlStr := fmt.Sprintf(`https://www.xxsy.net/partview/GetChapterList?bookid=%v&noNeedBuy=0&special=0&maxFreeChapterId=0&isMonthly=0`, bookID)

	html, err := GetHTML(urlStr, ``)
	if err != nil {
		return links, err
	}
	html2 := fmt.Sprintf(`
		<html>
		<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
		<title>chapter list</title>
		<body>
		%v
		</body>
		</html>
		`, html)

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html2))
	if e != nil {
		return links, e
	}
	link, _ := url.Parse(urlStr)
	links = GetLinks(g2, link)
	return links, nil

}
