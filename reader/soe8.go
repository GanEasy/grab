package reader

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//Soe8Reader U小说阅读网 (盗版小说网站)
type Soe8Reader struct {
}

// GetCategories 获取所有分类
func (r Soe8Reader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://m.soe8.com/`

	list.Title = `分类-笔趣阁soe8`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`全部`, `/pages/list?action=list&drive=soe8&url=` + EncodeURL(`http://m.soe8.com/sort/0/1.html`), "", `link`, ``, nil, ``},
		Card{`玄幻奇幻`, `/pages/list?action=list&drive=soe8&url=` + EncodeURL(`http://m.soe8.com/sort/1_1/`), "", `link`, ``, nil, ``},
		Card{`武侠仙侠`, `/pages/list?action=list&drive=soe8&url=` + EncodeURL(`http://m.soe8.com/sort/2_1/`), "", `link`, ``, nil, ``},
		Card{`都市言情`, `/pages/list?action=list&drive=soe8&url=` + EncodeURL(`http://m.soe8.com/sort/3_1/`), "", `link`, ``, nil, ``},
		Card{`历史军事`, `/pages/list?action=list&drive=soe8&url=` + EncodeURL(`http://m.soe8.com/sort/4_1/`), "", `link`, ``, nil, ``},
		Card{`游戏体育`, `/pages/list?action=list&drive=soe8&url=` + EncodeURL(`http://m.soe8.com/sort/5_1/`), "", `link`, ``, nil, ``},
		Card{`科幻灵异`, `/pages/list?action=list&drive=soe8&url=` + EncodeURL(`http://m.soe8.com/sort/6_1/`), "", `link`, ``, nil, ``},

		Card{`\月点击`, `/pages/list?action=list&drive=soe8&url=` + EncodeURL(`http://m.soe8.com/top/monthvisit_1/`), "", `link`, ``, nil, ``},
		Card{`\周点击`, `/pages/list?action=list&drive=soe8&url=` + EncodeURL(`http://m.soe8.com/top/weekvisit_1/`), "", `link`, ``, nil, ``},
		Card{`\总点击`, `/pages/list?action=list&drive=soe8&url=` + EncodeURL(`http://m.soe8.com/top/allvisit_1/`), "", `link`, ``, nil, ``},
	}
	return list, nil
}

// GetList 获取书籍列表列表
func (r Soe8Reader) GetList(urlStr string) (list Catalog, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, ``)
	if err != nil {
		return
	}

	html2, err := FindContentHTML(html, `.list`)

	g, e := goquery.NewDocumentFromReader(strings.NewReader(html2))

	if e != nil {
		return list, e
	}
	list.Title = FindString(`(?P<title>(.)+)_`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	g2, e2 := goquery.NewDocumentFromReader(strings.NewReader(html))

	if e2 != nil {
		return list, e2
	}
	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `http://m.soe8.com/76_76207/`, `http://m.soe8.com/80_80569/`, `http://m.soe8.com/76_76207_1/`)
		if state {
			l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `soe8`)

	list.SourceURL = urlStr

	var links2 = GetLinks(g2, link)
	list.Next = GetNextLink(links2)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetCatalog 获取章节列表
func (r Soe8Reader) GetCatalog(urlStr string) (list Catalog, err error) {

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
	list.Title = FindString(`(?P<title>(.)+)_(?P<title2>(.)+)_`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	html2, _ := FindContentHTML(html, `.fk`)

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html2))

	var links = GetLinks(g2, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `http://m.soe8.com/0_2/101956.html`, `http://m.soe8.com/80_80569/27942619.html`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `soe8`)

	list.SourceURL = urlStr

	// var links2 = GetLinks(g, link)

	mnpage := SelectString(`(?P<cate>(\d)+)_(?P<bookid>(\d)+)_(?P<page>(\d)+)`, urlStr)

	page, _ := strconv.Atoi(mnpage[`page`])

	nextURL := fmt.Sprintf(`http://m.soe8.com/%v_%v_%v/`, mnpage[`cate`], mnpage[`bookid`], page+1)

	list.Next = Link{Title: "next", URL: nextURL}

	// list.Next = GetNextLink(links2)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}
	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetInfo 获取详细内容
func (r Soe8Reader) GetInfo(urlStr string) (ret Content, err error) {

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

	ret.Title = FindString(`_(?P<title>(.)+)-`, article.Title, "title")
	if ret.Title == `` {
		ret.Title = article.Title
	}

	reg := regexp.MustCompile(`<a([^>]+)>([^<]+)<\/a>`)

	article.ReadContent = reg.ReplaceAllString(article.ReadContent, "")

	reg2 := regexp.MustCompile(`最快更新([^<]+)章节！`)

	article.ReadContent = reg2.ReplaceAllString(article.ReadContent, "")

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
