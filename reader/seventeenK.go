package reader

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//SeventeenKReader 纵横小说网
type SeventeenKReader struct {
}

// GetCategories 获取所有分类
func (r SeventeenKReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://book.17k.com`

	list.Title = `17K小说网`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{

		Card{`奇幻玄幻`, `/pages/list?action=book&drive=17k&url=` + EncodeURL(`https://www.17k.com/all/book/2_21_0_0_0_0_0_0_1.html`), "", `link`, ``, nil, ``},
		Card{`武侠仙侠`, `/pages/list?action=book&drive=17k&url=` + EncodeURL(`https://www.17k.com/all/book/2_24_0_0_0_0_0_0_1.html`), "", `link`, ``, nil, ``},
		Card{`都市小说`, `/pages/list?action=book&drive=17k&url=` + EncodeURL(`https://www.17k.com/all/book/2_3_0_0_0_0_0_0_1.html`), "", `link`, ``, nil, ``},
		Card{`历史军事`, `/pages/list?action=book&drive=17k&url=` + EncodeURL(`https://www.17k.com/all/book/2_22_0_0_0_0_0_0_1.html`), "", `link`, ``, nil, ``},
		Card{`游戏竞技`, `/pages/list?action=book&drive=17k&url=` + EncodeURL(`https://www.17k.com/all/book/2_23_0_0_0_0_0_0_1.html`), "", `link`, ``, nil, ``},
		Card{`科幻末世`, `/pages/list?action=book&drive=17k&url=` + EncodeURL(`https://www.17k.com/all/book/2_14_0_0_0_0_0_0_1.html`), "", `link`, ``, nil, ``},
		Card{`古装言情`, `/pages/list?action=book&drive=17k&url=` + EncodeURL(`https://www.17k.com/all/book/3_5_0_0_0_0_0_0_1.html`), "", `link`, ``, nil, ``},
		Card{`都市言情`, `/pages/list?action=book&drive=17k&url=` + EncodeURL(`https://www.17k.com/all/book/3_17_0_0_0_0_0_0_1.html`), "", `link`, ``, nil, ``},
		Card{`浪漫青春`, `/pages/list?action=book&drive=17k&url=` + EncodeURL(`https://www.17k.com/all/book/3_20_0_0_0_0_0_0_1.html`), "", `link`, ``, nil, ``},
		Card{`幻想言情`, `/pages/list?action=book&drive=17k&url=` + EncodeURL(`https://www.17k.com/all/book/3_18_0_0_0_0_0_0_1.html`), "", `link`, ``, nil, ``},
	}
	return list, nil
}

// GetList 获取书籍列表列表
func (r SeventeenKReader) GetList(urlStr string) (list Catalog, err error) {

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

	list.Title = g.Find("title").Text()

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://www.17k.com/book/2897539.html`, `https://www.17k.com/book/2927482.html`, `https://www.17k.com/list/2897539.html`)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `17k`)

	list.SourceURL = urlStr

	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetCatalog 获取章节列表
func (r SeventeenKReader) GetCatalog(urlStr string) (list Catalog, err error) {

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
		l.URL, state = JaccardMateGetURL(l.URL, `https://www.17k.com/chapter/2927485/36602417.html`, `https://www.17k.com/chapter/493239/36611963.html`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `17k`)

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetChapter 获取详细内容
func (r SeventeenKReader) GetInfo(urlStr string) (ret Content, err error) {

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
	ret.SourceURL = urlStr

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
