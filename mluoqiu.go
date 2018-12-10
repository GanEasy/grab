package grab

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//MLuoqiuReader 落秋中文网手机版 (盗版小说网站)
type MLuoqiuReader struct {
}

// GetCategories 获取所有分类
func (r MLuoqiuReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://m.luoqiu.com/`

	list.Title = `落秋中文网(免)`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`玄幻魔法`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://m.luoqiu.com/sort-1-1/`), "", `link`, ``, nil},
		Card{`武侠修真`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://m.luoqiu.com/sort-2-1/`), "", `link`, ``, nil},
		Card{`都市言情`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://m.luoqiu.com/sort-3-1/`), "", `link`, ``, nil},
		Card{`历史军事`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://m.luoqiu.com/sort-4-1/`), "", `link`, ``, nil},
		Card{`游戏竞技`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://m.luoqiu.com/sort-5-1/`), "", `link`, ``, nil},
		Card{`科幻小说`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://m.luoqiu.com/sort-6-1/`), "", `link`, ``, nil},
		Card{`恐怖灵异`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://m.luoqiu.com/sort-7-1/`), "", `link`, ``, nil},
		Card{`同人小说`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://m.luoqiu.com/sort-8-1/`), "", `link`, ``, nil},
		Card{`商战职场`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://m.luoqiu.com/sort-9-1/`), "", `link`, ``, nil},
		Card{`文学美文`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://m.luoqiu.com/sort-10-1/`), "", `link`, ``, nil},
		Card{`女生小说`, `/pages/transfer/list?action=book&drive=luoqiu&url=` + EncodeURL(`http://m.luoqiu.com/sort-11-1/`), "", `link`, ``, nil},
	}
	return list, nil
}

// GetBooks 获取书籍列表列表
func (r MLuoqiuReader) GetBooks(urlStr string) (list Catalog, err error) {

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
		l.URL, state = JaccardMateGetURL(l.URL, `http://m.luoqiu.com/info-22132/`, `http://m.luoqiu.com/info-909/`, `http://m.luoqiu.com/wapbook-22132/`)
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
func (r MLuoqiuReader) GetChapters(urlStr string) (list Catalog, err error) {

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
		l.URL, state = JaccardMateGetURL(l.URL, `http://m.luoqiu.com/wapbook-22132-5789517/`, `http://m.luoqiu.com/wapbook-38949-40993071/`, ``)
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
func (r MLuoqiuReader) GetChapter(urlStr string) (ret TextContent, err error) {

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
