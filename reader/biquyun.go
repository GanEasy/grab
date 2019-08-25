package reader

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//BiquyunReader U小说阅读网 (盗版小说网站)
type BiquyunReader struct {
}

// GetCategories 获取所有分类
func (r BiquyunReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://m.biquyun.com/`

	list.Title = `分类-笔趣阁biquyun`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`玄幻`, `/pages/list?action=book&drive=biquyun&url=` + EncodeURL(`https://m.biquyun.com/wapsort/1_1.html`), "", `link`, ``, nil, ``},
		Card{`修真`, `/pages/list?action=book&drive=biquyun&url=` + EncodeURL(`https://m.biquyun.com/wapsort/2_1.html`), "", `link`, ``, nil, ``},
		Card{`都市`, `/pages/list?action=book&drive=biquyun&url=` + EncodeURL(`https://m.biquyun.com/wapsort/3_1.html`), "", `link`, ``, nil, ``},
		Card{`历史`, `/pages/list?action=book&drive=biquyun&url=` + EncodeURL(`https://m.biquyun.com/wapsort/4_1.html`), "", `link`, ``, nil, ``},
		Card{`网游`, `/pages/list?action=book&drive=biquyun&url=` + EncodeURL(`https://m.biquyun.com/wapsort/5_1.html`), "", `link`, ``, nil, ``},
		Card{`科幻`, `/pages/list?action=book&drive=biquyun&url=` + EncodeURL(`https://m.biquyun.com/wapsort/6_1.html`), "", `link`, ``, nil, ``},
		Card{`恐怖`, `/pages/list?action=book&drive=biquyun&url=` + EncodeURL(`https://m.biquyun.com/wapsort/7_1.html`), "", `link`, ``, nil, ``},
		Card{`其它`, `/pages/list?action=book&drive=biquyun&url=` + EncodeURL(`https://m.biquyun.com/wapsort/8_1.html`), "", `link`, ``, nil, ``},
		Card{`↘↘↘排行榜↙↙`, ``, "", `link`, ``, nil, ``},
		Card{`总点击榜`, `/pages/list?action=book&drive=biquyun&url=` + EncodeURL(`https://m.biquyun.com/top/allvisit_1.html`), "", `link`, ``, nil, ``},
		Card{`月点击榜`, `/pages/list?action=book&drive=biquyun&url=` + EncodeURL(`https://m.biquyun.com/top/monthvisit_1.html`), "", `link`, ``, nil, ``},
		Card{`周点击榜`, `/pages/list?action=book&drive=biquyun&url=` + EncodeURL(`https://m.biquyun.com/top/weekvisit_1.html`), "", `link`, ``, nil, ``},
		Card{`日点击榜`, `/pages/list?action=book&drive=biquyun&url=` + EncodeURL(`https://m.biquyun.com/top/dayvisit_1.html`), "", `link`, ``, nil, ``},
	}
	return list, nil
}

// GetList 获取书籍列表列表
func (r BiquyunReader) GetList(urlStr string) (list Catalog, err error) {

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
	list.Title = FindString(`(?P<title>(.)+)_笔趣阁`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://m.biquyun.com/10_10267/`, `https://m.biquyun.com/20_20963/`, `https://m.biquyun.com/10_10267_1.html`)
		if state {
			l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `biquyun`)

	list.SourceURL = urlStr

	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetCatalog 获取章节列表
func (r BiquyunReader) GetCatalog(urlStr string) (list Catalog, err error) {

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

	// 偷心透视小村医最新章节,
	list.Title = FindString(`(?P<title>([^_])+)_`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	html2, _ := FindContentHTML(html, `.chapters`)

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html2))

	var links = GetLinks(g2, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://m.biquyun.com/10_10267/9947799.html`, `https://m.biquyun.com/20_20963/10035540.html`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `biquyun`)

	list.SourceURL = urlStr

	var links2 = GetLinks(g, link)

	list.Next = GetNextLink(links2)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}
	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetInfo 获取详细内容
func (r BiquyunReader) GetInfo(urlStr string) (ret Content, err error) {

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
	ret.SourceURL = urlStr

	ret.Title = FindString(`(?P<title>(.)+)_(?P<bookname>(.)+)_笔趣阁`, article.Title, "title")
	if ret.Title == `` {
		ret.Title = article.Title
	}

	reg := regexp.MustCompile(`<a([^<]+)<\/a>`)

	article.ReadContent = reg.ReplaceAllString(article.ReadContent, "")

	reg2 := regexp.MustCompile(`<br/>看(?P<title>([^<])+)最新章节请去 m.biquyun.com`)

	article.ReadContent = reg2.ReplaceAllString(article.ReadContent, "")

	reg3 := regexp.MustCompile(`<p([^<]+)<\/p>`)

	article.ReadContent = reg3.ReplaceAllString(article.ReadContent, "")

	reg4 := regexp.MustCompile(`<li([^<]+)<\/li>`)

	article.ReadContent = reg4.ReplaceAllString(article.ReadContent, "")
	c := MarkDownFormatContent(article.ReadContent)

	c = BookContReplace(c)

	ret.Contents = GetSectionByContent(c)

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
