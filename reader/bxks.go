package reader

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//BxksReader www.jininggeyin.com (盗版小说网站)
type BxksReader struct {
}

// GetCategories 获取所有分类
func (r BxksReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://www.jininggeyin.com/`

	list.Title = `分类-笔下看书阁`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`玄幻奇幻`, `/pages/list?action=book&drive=bxks&url=` + EncodeURL(`https://www.jininggeyin.com/ji/18.html`), "", `link`, ``, nil, ``, ``},
		Card{`武侠仙侠`, `/pages/list?action=book&drive=bxks&url=` + EncodeURL(`https://www.jininggeyin.com/ji/19.html`), "", `link`, ``, nil, ``, ``},
		Card{`历史军事`, `/pages/list?action=book&drive=bxks&url=` + EncodeURL(`https://www.jininggeyin.com/ji/20.html`), "", `link`, ``, nil, ``, ``},
		Card{`都市娱乐`, `/pages/list?action=book&drive=bxks&url=` + EncodeURL(`https://www.jininggeyin.com/ji/21.html`), "", `link`, ``, nil, ``, ``},
		Card{`科幻末日`, `/pages/list?action=book&drive=bxks&url=` + EncodeURL(`https://www.jininggeyin.com/ji/22.html`), "", `link`, ``, nil, ``, ``},
		Card{`悬疑灵异`, `/pages/list?action=book&drive=bxks&url=` + EncodeURL(`https://www.jininggeyin.com/ji/23.html`), "", `link`, ``, nil, ``, ``},
		Card{`游戏竞技`, `/pages/list?action=book&drive=bxks&url=` + EncodeURL(`https://www.jininggeyin.com/ji/34.html`), "", `link`, ``, nil, ``, ``},
		Card{`耽美百合`, `/pages/list?action=book&drive=bxks&url=` + EncodeURL(`https://www.jininggeyin.com/ji/40.html`), "", `link`, ``, nil, ``, ``},
		Card{`幻想言情`, `/pages/list?action=book&drive=bxks&url=` + EncodeURL(`https://www.jininggeyin.com/ji/28.html`), "", `link`, ``, nil, ``, ``},
		Card{`其他`, `/pages/list?action=book&drive=bxks&url=` + EncodeURL(`https://www.jininggeyin.com/ji/35.html`), "", `link`, ``, nil, ``, ``},
	}
	list.SearchSupport = true
	return list, nil
}

// Search 搜索资源
func (r BxksReader) Search(keyword string) (list Catalog, err error) {
	urlStr := `https://www.jininggeyin.com/search.html?keyword=` + keyword + `&t_btnsearch=`
	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, `.content`)
	if err != nil {
		return
	}

	g, e := goquery.NewDocumentFromReader(strings.NewReader(html))

	if e != nil {
		return list, e
	}

	list.Title = fmt.Sprintf(`%v - 搜索-笔下看书阁`, keyword)

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
		l.URL, state = JaccardMateGetURL(l.URL, `https://www.jininggeyin.com/jin/50142.html`, `https://www.jininggeyin.com/jin/49294.html`, ``)
		if state {
			l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `bxks`)

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetList 获取书籍列表列表
func (r BxksReader) GetList(urlStr string) (list Catalog, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, ``)
	if err != nil {
		return
	}

	html2, err2 := FindContentHTML(html, `.list-type`)
	if err2 != nil {
		return
	}
	g, e := goquery.NewDocumentFromReader(strings.NewReader(html2))

	if e != nil {
		return list, e
	}
	list.Title = FindString(`(?P<title>(.)+)-笔下看书阁`, g.Find("title").Text(), "title")
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
		l.URL, state = JaccardMateGetURL(l.URL, `https://www.jininggeyin.com/jin/52006.html`, `https://www.jininggeyin.com/jin/51702.html`, ``)
		if state {
			l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `bxks`)

	list.SourceURL = urlStr
	// log.Println(links)
	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetCatalog 获取章节列表
func (r BxksReader) GetCatalog(urlStr string) (list Catalog, err error) {

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
	list.Title = FindString(`(?P<title>(.)+)最新章节阅读`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	html2, _ := FindContentHTML(html, `.chapterlist`)

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html2))

	var links = GetLinks(g2, link)

	if len(links) == 0 {
		links = GetLinks(g, link)
	}

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://www.jininggeyin.com/jin1/50725/SURG9JT8HSn9W.html`, `https://www.jininggeyin.com/jin1/50749/ecOkp0Ccya3Ky.html`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	needLinks = CleaningFrontRepeat(needLinks)

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `bxks`)

	list.SourceURL = urlStr

	// var links2 = GetLinks(g, link)

	// list.Next = GetNextLink(links2)
	// if list.Next.URL != `` {
	// 	list.Next.URL = EncodeURL(list.Next.URL)
	// }
	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetInfo 获取详细内容
func (r BxksReader) GetInfo(urlStr string) (ret Content, err error) {

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

	ret.Title = FindString(`(?P<title>(.)+)_(?P<bookname>(.)+)-笔下看书阁`, article.Title, "title")
	if ret.Title == `` {
		ret.Title = article.Title
	}

	// reg := regexp.MustCompile(`<a([^<]+)<\/a>`)

	// article.ReadContent = reg.ReplaceAllString(article.ReadContent, "")

	reg2 := regexp.MustCompile(`<h1([^<]*)<\/h1>`)

	article.ReadContent = reg2.ReplaceAllString(article.ReadContent, "")

	reg4 := regexp.MustCompile(`天才一秒钟记住本网站([^<]*)小说网站!`)

	article.ReadContent = reg4.ReplaceAllString(article.ReadContent, "")

	c := MarkDownFormatContent(article.ReadContent)
	c = BookContReplace(c)

	ret.Contents = GetSectionByContent(c)
	// log.Println(len(ret.Contents))

	links, _ := GetLinkByHTML(urlStr, html)
	// log.Println(links)
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
