package reader

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//JxReader 笔趣阁qula (盗版小说网站)
type JxReader struct {
}

// GetCategories 获取所有分类
func (r JxReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `https://m.jx.la/`

	list.Title = `分类-笔趣阁jxla`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`-全部分类`, `/pages/list?action=book&drive=jx&url=` + EncodeURL(`https://m.jx.la/wapsort/0_1.html`), "", `link`, ``, nil, ``},
		Card{`-玄幻奇幻`, `/pages/list?action=book&drive=jx&url=` + EncodeURL(`https://m.jx.la/wapsort/1_1.html`), "", `link`, ``, nil, ``},
		Card{`-武侠仙侠`, `/pages/list?action=book&drive=jx&url=` + EncodeURL(`https://m.jx.la/wapsort/2_1.html`), "", `link`, ``, nil, ``},
		Card{`-都市言情`, `/pages/list?action=book&drive=jx&url=` + EncodeURL(`https://m.jx.la/wapsort/3_1.html`), "", `link`, ``, nil, ``},
		Card{`-历史军事`, `/pages/list?action=book&drive=jx&url=` + EncodeURL(`https://m.jx.la/wapsort/4_1.html`), "", `link`, ``, nil, ``},
		Card{`-科幻灵异`, `/pages/list?action=book&drive=jx&url=` + EncodeURL(`https://m.jx.la/wapsort/5_1.html`), "", `link`, ``, nil, ``},
		Card{`-网游竞技`, `/pages/list?action=book&drive=jx&url=` + EncodeURL(`https://m.jx.la/wapsort/6_1.html`), "", `link`, ``, nil, ``},
		Card{`-女生频道`, `/pages/list?action=book&drive=jx&url=` + EncodeURL(`https://m.jx.la/wapsort/7_1.html`), "", `link`, ``, nil, ``},

		Card{`\全部排行`, `/pages/list?action=book&drive=jx&url=` + EncodeURL(`https://m.jx.la/waptop/month.html`), "", `link`, ``, nil, ``},
		Card{`\玄幻奇幻`, `/pages/list?action=book&drive=jx&url=` + EncodeURL(`https://m.jx.la/waptop/month1.html`), "", `link`, ``, nil, ``},
		Card{`\武侠仙侠`, `/pages/list?action=book&drive=jx&url=` + EncodeURL(`https://m.jx.la/waptop/month2.html`), "", `link`, ``, nil, ``},
		Card{`\都市言情`, `/pages/list?action=book&drive=jx&url=` + EncodeURL(`https://m.jx.la/waptop/month3.html`), "", `link`, ``, nil, ``},
		Card{`\历史军事`, `/pages/list?action=book&drive=jx&url=` + EncodeURL(`https://m.jx.la/waptop/month4.html`), "", `link`, ``, nil, ``},
		Card{`\科幻灵异`, `/pages/list?action=book&drive=jx&url=` + EncodeURL(`https://m.jx.la/waptop/month5.html`), "", `link`, ``, nil, ``},
		Card{`\网游竞技`, `/pages/list?action=book&drive=jx&url=` + EncodeURL(`https://m.jx.la/waptop/month6.html`), "", `link`, ``, nil, ``},
		Card{`\女生频道`, `/pages/list?action=book&drive=jx&url=` + EncodeURL(`https://m.jx.la/waptop/month7.html`), "", `link`, ``, nil, ``},
	}
	list.SearchSupport = true
	return list, nil
}

// GetList 获取书籍列表列表
func (r JxReader) GetList(urlStr string) (list Catalog, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, `.recommend`)
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
		l.URL, state = JaccardMateGetURL(l.URL, `https://m.jx.la/book/175443/`, `https://m.jx.la/book/142095/`, `https://m.jx.la/booklist/175443.html`)
		if state {
			l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `jx`)

	list.SourceURL = urlStr

	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// Search 搜索资源
func (r JxReader) Search(keyword string) (list Catalog, err error) {

	urlStr := `https://sou.xanbhx.com/search?q=` + keyword + `&s=&t=m&siteid=qula`

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, `.recommend`)
	if err != nil {
		return
	}

	g, e := goquery.NewDocumentFromReader(strings.NewReader(html))

	if e != nil {
		return list, e
	}

	list.Title = fmt.Sprintf(`%v - 搜索结果-qu.la`, keyword)

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	if len(links) == 0 {

		g2, e2 := goquery.NewDocumentFromReader(strings.NewReader(html))

		if e2 != nil {
			return list, e2
		}

		links = GetLinks(g2, link)
	}

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://m.jx.la/book/175443/`, `https://m.jx.la/book/142095/`, `https://m.jx.la/booklist/175443.html`)
		if state {
			l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			l.Title = strings.Replace(l.Title, " ", "", -1)
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `jx`)

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	return list, nil
}

// GetCatalog 获取章节列表
func (r JxReader) GetCatalog(urlStr string) (list Catalog, err error) {

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
	list.Title = FindString(`(?P<title>(.)+)全文阅读_`, g.Find("title").Text(), "title")
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
		l.URL, state = JaccardMateGetURL(l.URL, `https://m.jx.la/book/175443/9124417.html`, `https://m.jx.la/book/142095/7545899.html`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `jx`)

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
func (r JxReader) GetInfo(urlStr string) (ret Content, err error) {

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

	reg3 := regexp.MustCompile(`<u([^>]*)>([^<]*)<\/u>`)

	article.ReadContent = reg3.ReplaceAllString(article.ReadContent, "")

	reg2 := regexp.MustCompile(`<span([^>]*)>([^<]*)<\/span>`)

	article.ReadContent = reg2.ReplaceAllString(article.ReadContent, "")

	reg := regexp.MustCompile(`<a([^>]*)>([^<]*)<\/a>`)

	article.ReadContent = reg.ReplaceAllString(article.ReadContent, "")

	reg4 := regexp.MustCompile(`<([^>]+)>([^<]+)<\/([^>]+)>`)

	article.ReadContent = reg4.ReplaceAllString(article.ReadContent, "")

	// log.Println(article.ReadContent)
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
