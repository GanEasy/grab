package reader

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//Xs7878Reader 7878小说 (盗版小说网站)
type Xs7878Reader struct {
}

// GetCategories 获取所有分类
func (r Xs7878Reader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://m.7878xs.com/`

	list.Title = `分类-7878小说`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`异界`, `/pages/transfer/list?action=book&drive=7878xs&url=` + EncodeURL(`http://m.7878xs.com/f/1/1.htm`), "", `link`, ``, nil},
		Card{`奇幻`, `/pages/transfer/list?action=book&drive=7878xs&url=` + EncodeURL(`http://m.7878xs.com/f/2/1.htm`), "", `link`, ``, nil},
		Card{`仙侠`, `/pages/transfer/list?action=book&drive=7878xs&url=` + EncodeURL(`http://m.7878xs.com/f/3/1.htm`), "", `link`, ``, nil},
		Card{`都市`, `/pages/transfer/list?action=book&drive=7878xs&url=` + EncodeURL(`http://m.7878xs.com/f/4/1.htm`), "", `link`, ``, nil},
		Card{`历史`, `/pages/transfer/list?action=book&drive=7878xs&url=` + EncodeURL(`http://m.7878xs.com/f/5/1.htm`), "", `link`, ``, nil},
		Card{`游戏`, `/pages/transfer/list?action=book&drive=7878xs&url=` + EncodeURL(`http://m.7878xs.com/f/6/1.htm`), "", `link`, ``, nil},
		Card{`竞技`, `/pages/transfer/list?action=book&drive=7878xs&url=` + EncodeURL(`http://m.7878xs.com/f/7/1.htm`), "", `link`, ``, nil},
		Card{`灵异`, `/pages/transfer/list?action=book&drive=7878xs&url=` + EncodeURL(`http://m.7878xs.com/f/8/1.htm`), "", `link`, ``, nil},
		Card{`其他`, `/pages/transfer/list?action=book&drive=7878xs&url=` + EncodeURL(`http://m.7878xs.com/f/9/1.htm`), "", `link`, ``, nil},
	}
	return list, nil
}

// GetList 获取书籍列表列表
func (r Xs7878Reader) GetList(urlStr string) (list Catalog, err error) {

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

	list.Title = FindString(`(?P<title>(.)+)小说第`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `http://m.7878xs.com/x/269052/`, `http://m.7878xs.com/x/182203/`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `7878xs`)

	list.SourceURL = urlStr

	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetCatalog 获取章节列表
func (r Xs7878Reader) GetCatalog(urlStr string) (list Catalog, err error) {

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

	list.Title = FindString(`(?P<title>(.)+)_无弹窗`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	html2, _ := FindContentHTML(html, ` #J-chapterlist`)

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html2))

	var links = GetLinks(g2, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `http://m.7878xs.com/x/269052/1149019.htm`, `http://m.7878xs.com/x/260546/1403938.htm`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `7878xs`)

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
func (r Xs7878Reader) GetInfo(urlStr string) (ret Content, err error) {

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
