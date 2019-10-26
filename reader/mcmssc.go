package reader

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//McmsscReader www.mcmssc.com (盗版小说网站)
type McmsscReader struct {
}

// GetCategories 获取所有分类
func (r McmsscReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://m.mcmssc.com/`

	list.Title = `分类-笔趣阁mcmssc`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`玄幻奇幻`, `/pages/list?action=book&drive=mcmssc&url=` + EncodeURL(`https://www.mcmssc.com/xuanhuanxiaoshuo/`), "", `link`, ``, nil, ``},
		Card{`修真仙侠`, `/pages/list?action=book&drive=mcmssc&url=` + EncodeURL(`https://www.mcmssc.com/xiuzhenxiaoshuo/`), "", `link`, ``, nil, ``},
		Card{`都市青春`, `/pages/list?action=book&drive=mcmssc&url=` + EncodeURL(`https://www.mcmssc.com/dushixiaoshuo/`), "", `link`, ``, nil, ``},
		Card{`历史穿越`, `/pages/list?action=book&drive=mcmssc&url=` + EncodeURL(`https://www.mcmssc.com/chuanyuexiaoshuo/`), "", `link`, ``, nil, ``},
		Card{`网游竞技`, `/pages/list?action=book&drive=mcmssc&url=` + EncodeURL(`https://www.mcmssc.com/wangyouxiaoshuo/`), "", `link`, ``, nil, ``},
		Card{`科幻灵异`, `/pages/list?action=book&drive=mcmssc&url=` + EncodeURL(`https://www.mcmssc.com/kehuanxiaoshuo/`), "", `link`, ``, nil, ``},
		Card{`其它小说`, `/pages/list?action=book&drive=mcmssc&url=` + EncodeURL(`https://www.mcmssc.com/qitaxiaoshuo/`), "", `link`, ``, nil, ``},
	}
	return list, nil
}

// GetList 获取书籍列表列表
func (r McmsscReader) GetList(urlStr string) (list Catalog, err error) {

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

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://www.mcmssc.com/90_90485/`, `https://www.mcmssc.com/87_87108/`, ``)
		if state {
			l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `mcmssc`)

	list.SourceURL = urlStr

	pagemod := FindString(`window.location.href = "(?P<pagemod>[^"]+)"`, html, "pagemod")
	page := FindString(`curr: (?P<page>(\d)+)`, html, "page")
	p, err := strconv.Atoi(page)
	if p > 0 && err == nil {
		// 已经组装url
		nextURL := fmt.Sprintf(`https://www.mcmssc.com%v%v.html`, pagemod, (p + 1))
		list.Next = Link{`下一页`, EncodeURL(nextURL), ``}
		log.Println(`pagemodpage`, nextURL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetCatalog 获取章节列表
func (r McmsscReader) GetCatalog(urlStr string) (list Catalog, err error) {

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
	list.Title = FindString(`(?P<title>(.)+)最新章节_`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	html2, _ := FindContentHTML(html, `#list`)

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html2))

	var links = GetLinks(g2, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://www.mcmssc.com/87_87108/39151754.html`, `https://www.mcmssc.com/79_79810/34434083.html`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	needLinks = CleaningFrontRepeat(needLinks)

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `mcmssc`)

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
func (r McmsscReader) GetInfo(urlStr string) (ret Content, err error) {

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
