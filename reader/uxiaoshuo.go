package reader

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//UxiaoshuoReader U小说阅读网 (盗版小说网站)
type UxiaoshuoReader struct {
}

// GetCategories 获取所有分类
func (r UxiaoshuoReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://m.uxiaoshuo.com/`

	list.Title = `分类-U小说阅读网`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`全部`, `/pages/list?action=book&drive=uxiaoshuo&url=` + EncodeURL(`https://m.uxiaoshuo.com/sort/0/1.html`), "", `link`, ``, nil, ``, ``},
		Card{`玄幻奇幻`, `/pages/list?action=book&drive=uxiaoshuo&url=` + EncodeURL(`https://m.uxiaoshuo.com/sort/1/1.html`), "", `link`, ``, nil, ``, ``},
		Card{`武侠仙侠`, `/pages/list?action=book&drive=uxiaoshuo&url=` + EncodeURL(`https://m.uxiaoshuo.com/sort/2/1.html`), "", `link`, ``, nil, ``, ``},
		Card{`都市言情`, `/pages/list?action=book&drive=uxiaoshuo&url=` + EncodeURL(`https://m.uxiaoshuo.com/sort/3/1.html`), "", `link`, ``, nil, ``, ``},
		Card{`历史军事`, `/pages/list?action=book&drive=uxiaoshuo&url=` + EncodeURL(`https://m.uxiaoshuo.com/sort/4/1.html`), "", `link`, ``, nil, ``, ``},
		Card{`科幻灵异`, `/pages/list?action=book&drive=uxiaoshuo&url=` + EncodeURL(`https://m.uxiaoshuo.com/sort/5/1.html`), "", `link`, ``, nil, ``, ``},
		Card{`网游竞技`, `/pages/list?action=book&drive=uxiaoshuo&url=` + EncodeURL(`https://m.uxiaoshuo.com/sort/6/1.html`), "", `link`, ``, nil, ``, ``},
		Card{`女生频道`, `/pages/list?action=book&drive=uxiaoshuo&url=` + EncodeURL(`https://m.uxiaoshuo.com/sort/7/1.html`), "", `link`, ``, nil, ``, ``},

		Card{`↘↘↘排行榜↙↙↙`, ``, "", `link`, ``, nil, ``, ``},
		Card{`全部热门↘↘↘`, `/pages/list?action=list&drive=uxiaoshuo&url=` + EncodeURL(`https://m.uxiaoshuo.com/phb/week.html`), "", `link`, ``, nil, ``, ``},
		Card{`玄幻奇幻↘↘↘`, `/pages/list?action=list&drive=uxiaoshuo&url=` + EncodeURL(`https://m.uxiaoshuo.com/phb/week1.html`), "", `link`, ``, nil, ``, ``},
		Card{`武侠仙侠↘↘↘`, `/pages/list?action=list&drive=uxiaoshuo&url=` + EncodeURL(`https://m.uxiaoshuo.com/phb/week2.html`), "", `link`, ``, nil, ``, ``},
		Card{`都市言情↘↘↘`, `/pages/list?action=list&drive=uxiaoshuo&url=` + EncodeURL(`https://m.uxiaoshuo.com/phb/week3.html`), "", `link`, ``, nil, ``, ``},
		Card{`历史军事↘↘↘`, `/pages/list?action=list&drive=uxiaoshuo&url=` + EncodeURL(`https://m.uxiaoshuo.com/phb/week4.html`), "", `link`, ``, nil, ``, ``},
		Card{`科幻灵异↘↘↘`, `/pages/list?action=list&drive=uxiaoshuo&url=` + EncodeURL(`https://m.uxiaoshuo.com/phb/week5.html`), "", `link`, ``, nil, ``, ``},
		Card{`游戏竞技↘↘↘`, `/pages/list?action=list&drive=uxiaoshuo&url=` + EncodeURL(`https://m.uxiaoshuo.com/phb/week6.html`), "", `link`, ``, nil, ``, ``},
		Card{`女生频道↘↘↘`, `/pages/list?action=list&drive=uxiaoshuo&url=` + EncodeURL(`https://m.uxiaoshuo.com/phb/week7.html`), "", `link`, ``, nil, ``, ``},
	}
	return list, nil
}

// GetList 获取书籍列表列表
func (r UxiaoshuoReader) GetList(urlStr string) (list Catalog, err error) {

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
	list.Title = FindString(`(?P<title>(.)+),`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://m.uxiaoshuo.com/140/140420/`, `https://m.uxiaoshuo.com/238/238242/`, `https://m.uxiaoshuo.com/140/140420/all.html`)
		if state {
			l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `uxiaoshuo`)

	list.SourceURL = urlStr

	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// Search 搜索资源
func (r UxiaoshuoReader) Search(keyword string) (list Catalog, err error) {
	return
}

// GetCatalog 获取章节列表
func (r UxiaoshuoReader) GetCatalog(urlStr string) (list Catalog, err error) {

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
	list.Title = FindString(`(?P<title>(.)+)最新章节,`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	html2, _ := FindContentHTML(html, `#chapterlist`)

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html2))

	var links = GetLinks(g2, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://m.uxiaoshuo.com/238/238242/2093869.html`, `https://m.uxiaoshuo.com/291/291542/1840965.html`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `uxiaoshuo`)

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
func (r UxiaoshuoReader) GetInfo(urlStr string) (ret Content, err error) {

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

	ret.Title = FindString(`(?P<title>(.)+)_(?P<bookname>(.)+)_`, article.Title, "title")
	if ret.Title == `` {
		ret.Title = article.Title
	}

	reg := regexp.MustCompile(`<a([^>]+)>([^<]+)<\/a>`)

	article.ReadContent = reg.ReplaceAllString(article.ReadContent, "")

	reg2 := regexp.MustCompile(`温馨提示：([^“]+)阅读。`)

	article.ReadContent = reg2.ReplaceAllString(article.ReadContent, "")

	reg3 := regexp.MustCompile(`---([\-]+)\[\]`)

	article.ReadContent = reg3.ReplaceAllString(article.ReadContent, "")

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
