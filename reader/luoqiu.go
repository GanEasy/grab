package reader

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//LuoqiuReader 落秋中文网 (盗版小说网站)
type LuoqiuReader struct {
}

// GetCategories 获取所有分类
func (r LuoqiuReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://www.luoqiu.com/`

	list.Title = `落秋中文网(免)`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`玄幻魔法`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://www.luoqiu.com/class/1_1.html`), "", `link`, ``, nil},
		Card{`武侠修真`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://www.luoqiu.com/class/2_1.html`), "", `link`, ``, nil},
		Card{`都市言情`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://www.luoqiu.com/class/3_1.html`), "", `link`, ``, nil},
		Card{`历史军事`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://www.luoqiu.com/class/4_1.html`), "", `link`, ``, nil},
		Card{`游戏竞技`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://www.luoqiu.com/class/5_1.html`), "", `link`, ``, nil},
		Card{`科幻小说`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://www.luoqiu.com/class/6_1.html`), "", `link`, ``, nil},
		Card{`恐怖灵异`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://www.luoqiu.com/class/7_1.html`), "", `link`, ``, nil},
		Card{`同人小说`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://www.luoqiu.com/class/8_1.html`), "", `link`, ``, nil},
		Card{`商战职场`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://www.luoqiu.com/class/9_1.html`), "", `link`, ``, nil},
		Card{`文学美文`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://www.luoqiu.com/class/10_1.html`), "", `link`, ``, nil},
		Card{`女生小说`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://www.luoqiu.com/class/11_1.html`), "", `link`, ``, nil},
	}
	return list, nil
}

// GetBooks 获取书籍列表列表
func (r LuoqiuReader) GetBooks(urlStr string) (list Catalog, err error) {

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
		l.URL, state = JaccardMateGetURL(l.URL, `http://www.luoqiu.com/book/35175.html`, `http://www.luoqiu.com/book/26145.html`, `http://www.luoqiu.com/read/35175/`)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/chapter/get`, `luoqiu`)

	list.SourceURL = urlStr

	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetChapters 获取章节列表
func (r LuoqiuReader) GetChapters(urlStr string) (list Catalog, err error) {

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
		l.URL, state = JaccardMateGetURL(l.URL, `http://www.luoqiu.com/read/35175/9654580.html`, `http://www.luoqiu.com/read/270270/41741931.html`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/chapter/info`, `luoqiu`)

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetChapter 获取详细内容
func (r LuoqiuReader) GetChapter(urlStr string) (ret TextContent, err error) {

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

	ret.Content = GetSectionByContent(c)

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
