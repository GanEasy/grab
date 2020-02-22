package reader

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//Xs280Reader www.280xs.com (盗版顶点小说网站)
type Xs280Reader struct {
}

// GetCategories 获取所有分类
func (r Xs280Reader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://m.Xs280.com/`

	list.Title = `分类-顶点小说280xs`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`玄幻奇幻`, `/pages/list?action=book&drive=xs280&url=` + EncodeURL(`https://www.280xs.com/book_1_1/`), "", `link`, ``, nil, ``},
		Card{`武侠修真`, `/pages/list?action=book&drive=xs280&url=` + EncodeURL(`https://www.280xs.com/book_2_1/`), "", `link`, ``, nil, ``},
		Card{`言情都市`, `/pages/list?action=book&drive=xs280&url=` + EncodeURL(`https://www.280xs.com/book_3_1/`), "", `link`, ``, nil, ``},
		Card{`历史穿越`, `/pages/list?action=book&drive=xs280&url=` + EncodeURL(`https://www.280xs.com/book_4_1/`), "", `link`, ``, nil, ``},
		Card{`网游动漫`, `/pages/list?action=book&drive=xs280&url=` + EncodeURL(`https://www.280xs.com/book_6_1/`), "", `link`, ``, nil, ``},
		Card{`科幻小说`, `/pages/list?action=book&drive=xs280&url=` + EncodeURL(`https://www.280xs.com/book_7_1/`), "", `link`, ``, nil, ``},
		Card{`恐怖灵异`, `/pages/list?action=book&drive=xs280&url=` + EncodeURL(`https://www.280xs.com/book_8_1/`), "", `link`, ``, nil, ``},
		Card{`其它小说`, `/pages/list?action=book&drive=xs280&url=` + EncodeURL(`https://www.280xs.com/book_10_1/`), "", `link`, ``, nil, ``},
		// Card{`排行榜单`, `/pages/list?action=book&drive=xs280&url=` + EncodeURL(`https://www.280xs.com/paihangbang/`), "", `link`, ``, nil, ``},
	}
	return list, nil
}

// GetList 获取书籍列表列表
func (r Xs280Reader) GetList(urlStr string) (list Catalog, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, ``)
	if err != nil {
		return
	}

	html2, err2 := FindContentHTML(html, `.l`)
	if err2 != nil {
		return
	}
	g, e := goquery.NewDocumentFromReader(strings.NewReader(html2))

	if e != nil {
		return list, e
	}
	list.Title = FindString(`(?P<title>(.)+)_好看的`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

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
		l.URL, state = JaccardMateGetURL(l.URL, `https://www.280xs.com/dingdian/10_10353/`, `https://www.280xs.com/dingdian/36_36734/`, ``)
		if state {
			l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `xs280`)

	list.SourceURL = urlStr

	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetCatalog 获取章节列表
func (r Xs280Reader) GetCatalog(urlStr string) (list Catalog, err error) {

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
	list.Title = FindString(`(?P<title>(.)+)无弹窗_`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	html2, _ := FindContentHTML(html, `.article`)

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html2))

	var links = GetLinks(g2, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://www.280xs.com/dingdian/34_34693/n9414141.html`, `https://www.280xs.com/dingdian/26_26019/n7169097.html`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	needLinks = CleaningFrontRepeat(needLinks)

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `xs280`)

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
func (r Xs280Reader) GetInfo(urlStr string) (ret Content, err error) {

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

	ret.Title = FindString(`(?P<title>(.)+) 无弹窗广告`, article.Title, "title")
	if ret.Title == `` {
		ret.Title = article.Title
	}

	reg := regexp.MustCompile(`<a([^<]+)<\/a>`)

	article.ReadContent = reg.ReplaceAllString(article.ReadContent, "")

	reg2 := regexp.MustCompile(`<div([^<]*)<\/div>`)

	article.ReadContent = reg2.ReplaceAllString(article.ReadContent, "")

	reg4 := regexp.MustCompile(`(天才一秒记住本站地址：([^<]*)最快更新！无广告！)`)

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
