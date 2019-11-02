package reader

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//BxwxReader 笔下文学 (盗版小说网站)
type BxwxReader struct {
}

// GetCategories 获取所有分类
func (r BxwxReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://m.booktxt.com/`

	list.Title = `分类-笔下文学`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{

		// https://m.bxwx.la/bxwx/week.html    https://m.bxwx.la/bsort0/0/1.htm
		Card{`-全部分类`, `/pages/list?action=list&drive=bxwx&url=` + EncodeURL(`https://m.bxwx.la/bsort0/0/1.htm`), "", `link`, ``, nil, ``},
		Card{`-玄幻奇幻`, `/pages/list?action=list&drive=bxwx&url=` + EncodeURL(`https://m.bxwx.la/bsort1/0/1.htm`), "", `link`, ``, nil, ``},
		Card{`-武侠仙侠`, `/pages/list?action=list&drive=bxwx&url=` + EncodeURL(`https://m.bxwx.la/bsort2/0/1.htm`), "", `link`, ``, nil, ``},
		Card{`-都市言情`, `/pages/list?action=list&drive=bxwx&url=` + EncodeURL(`https://m.bxwx.la/bsort3/0/1.htm`), "", `link`, ``, nil, ``},
		Card{`-历史军事`, `/pages/list?action=list&drive=bxwx&url=` + EncodeURL(`https://m.bxwx.la/bsort4/0/1.htm`), "", `link`, ``, nil, ``},
		Card{`-科幻灵异`, `/pages/list?action=list&drive=bxwx&url=` + EncodeURL(`https://m.bxwx.la/bsort5/0/1.htm`), "", `link`, ``, nil, ``},
		Card{`-网游竞技`, `/pages/list?action=list&drive=bxwx&url=` + EncodeURL(`https://m.bxwx.la/bsort6/0/1.htm`), "", `link`, ``, nil, ``},
		Card{`-女生频道`, `/pages/list?action=list&drive=bxwx&url=` + EncodeURL(`https://m.bxwx.la/bsort7/0/1.htm`), "", `link`, ``, nil, ``},

		Card{`↘↘↘排行榜↙↙↙`, ``, "", `link`, ``, nil, ``},
		Card{`全部热门↘↘↘`, `/pages/list?action=list&drive=bxwx&url=` + EncodeURL(`https://m.bxwx.la/bxwx/week.html`), "", `link`, ``, nil, ``},
		Card{`玄幻奇幻↘↘↘`, `/pages/list?action=list&drive=bxwx&url=` + EncodeURL(`https://m.bxwx.la/bxwx/week1.html`), "", `link`, ``, nil, ``},
		Card{`武侠仙侠↘↘↘`, `/pages/list?action=list&drive=bxwx&url=` + EncodeURL(`https://m.bxwx.la/bxwx/week2.html`), "", `link`, ``, nil, ``},
		Card{`都市言情↘↘↘`, `/pages/list?action=list&drive=bxwx&url=` + EncodeURL(`https://m.bxwx.la/bxwx/week3.html`), "", `link`, ``, nil, ``},
		Card{`历史军事↘↘↘`, `/pages/list?action=list&drive=bxwx&url=` + EncodeURL(`https://m.bxwx.la/bxwx/week4.html`), "", `link`, ``, nil, ``},
		Card{`科幻灵异↘↘↘`, `/pages/list?action=list&drive=bxwx&url=` + EncodeURL(`https://m.bxwx.la/bxwx/week5.html`), "", `link`, ``, nil, ``},
		Card{`游戏竞技↘↘↘`, `/pages/list?action=list&drive=bxwx&url=` + EncodeURL(`https://m.bxwx.la/bxwx/week6.html`), "", `link`, ``, nil, ``},
		Card{`女生频道↘↘↘`, `/pages/list?action=list&drive=bxwx&url=` + EncodeURL(`https://m.bxwx.la/bxwx/week7.html`), "", `link`, ``, nil, ``},
	}
	return list, nil
}

// GetList 获取书籍列表列表
func (r BxwxReader) GetList(urlStr string) (list Catalog, err error) {

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

	list.Title = `分类目录`
	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://m.bxwx.la/b/246/246596/`, `https://m.bxwx.la/b/287/287378/`, `https://m.bxwx.la/binfo/246/246596.htm`)
		if state {
			l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `bxwx`)

	list.SourceURL = urlStr

	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetCatalog 获取章节列表
func (r BxwxReader) GetCatalog(urlStr string) (list Catalog, err error) {

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

	list.Title = FindString(`(?P<title>(.)+)最新章节`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://m.bxwx.la/b/187/187253/9475107.html`, `https://m.bxwx.la/b/53/53693/3388765.html`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `bxwx`)

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
func (r BxwxReader) GetInfo(urlStr string) (ret Content, err error) {

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

	ret.Title = FindString(`(?P<title>(.)+)_(?P<bookname>(.)+)_`, article.Title, "title")
	if ret.Title == `` {
		ret.Title = article.Title
	}

	reg := regexp.MustCompile(`<a([^>]+)>([^，]+)<\/a>`)

	article.ReadContent = reg.ReplaceAllString(article.ReadContent, "")

	reg2 := regexp.MustCompile(`还在用浏览器看([^<]+)立即下载&gt;&gt;&gt;`)

	article.ReadContent = reg2.ReplaceAllString(article.ReadContent, "")

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
