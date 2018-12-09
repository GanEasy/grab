package grab

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//QidianReader 纵横小说网
type QidianReader struct {
}

// GetCategories 获取所有自定义分类(写死)
func (r QidianReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://book.qidian.com`

	list.Title = `分类-起点中文网`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`全部`, `/pages/book/get?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`玄幻`, `/pages/book/get?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=21&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`奇幻`, `/pages/book/get?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=1&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`武侠`, `/pages/book/get?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=2&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`仙侠`, `/pages/book/get?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=22&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`都市`, `/pages/book/get?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=4&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`现实`, `/pages/book/get?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=15&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`军事`, `/pages/book/get?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=6&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`历史`, `/pages/book/get?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=5&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`游戏`, `/pages/book/get?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=7&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`体育`, `/pages/book/get?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=8&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`科幻`, `/pages/book/get?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=9&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`灵异`, `/pages/book/get?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=10&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`二次元`, `/pages/book/get?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=12&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`短篇`, `/pages/book/get?drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=20076&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
	}
	return list, nil
}

// GetBooks 获取分类书籍列表
func (r QidianReader) GetBooks(urlStr string) (list Catalog, err error) {

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
		l.URL, state = JaccardMateGetURL(l.URL, `https://book.qidian.com/info/1010734492`, `https://book.qidian.com/info/1010868264`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/chapter/get`, `qidian`)

	list.SourceURL = urlStr

	// html := `{"I":"5333","V":"马经理"},`
	// page := FindString(`/p(?P<page>[^"]+)/`, html, "page")

	page := FindString(`&page=(?P<page>(\d)+)&`, html, "page")
	p, err := strconv.Atoi(page)
	if p > 0 && err == nil {
		// 已经组装url
		nextURL := strings.Replace(urlStr, fmt.Sprintf(`/p%v/`, p), fmt.Sprintf(`/p%v/`, p+1), -1)
		list.Next = Link{`下一页`, EncodeURL(nextURL), ``}
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetChapter 获取章节正文内容
func (r QidianReader) GetChapter(urlStr string) (ret TextContent, err error) {

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
	ret.Next = GetNextLink(links)
	if ret.Next.URL != `` {
		ret.Next.URL = EncodeURL(ret.Next.URL)
	}
	return ret, nil

}

// GetChapters 获取章节列表
func (r QidianReader) GetChapters(urlStr string) (list Catalog, err error) {

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

	catalogMsg := g.Find("#J-catalogCount").Text()
	if catalogMsg == `` { //todo 从 https://book.qidian.com/ajax/book/category?_csrfToken=&bookId=1004608738 中获取章节列表(要解释json)
		// panic(`catalogMsg`)
	}

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links { //起点普通和VIP章节不同地址
		l.URL, state = JaccardMateGetURL(l.URL, `https://read.qidian.com/chapter/ORlSeSgZ6E_MQzCecGvf7A2/DKk0ho2xSYTM5j8_3RRvhw2`, `https://read.qidian.com/chapter/_AaqI-dPJJ4uTkiRw_sFYA2/_4Wioy7TTQD6ItTi_ILQ7A2`, ``)
		if state {
			needLinks = append(needLinks, l)
		} else {
			l.URL, state = JaccardMateGetURL(l.URL, `https://vipreader.qidian.com/chapter/1004608738/347194141`, `https://vipreader.qidian.com/chapter/1010734492/399246504`, ``)
			if state {
				needLinks = append(needLinks, l)
			}
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/chapter/info`, `qidian`)

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	return list, nil

}
