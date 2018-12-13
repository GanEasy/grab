package grab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
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
		Card{`全部`, `/pages/transfer/list?action=book&drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`玄幻`, `/pages/transfer/list?action=book&drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=21&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`奇幻`, `/pages/transfer/list?action=book&drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=1&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`武侠`, `/pages/transfer/list?action=book&drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=2&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`仙侠`, `/pages/transfer/list?action=book&drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=22&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`都市`, `/pages/transfer/list?action=book&drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=4&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`现实`, `/pages/transfer/list?action=book&drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=15&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`军事`, `/pages/transfer/list?action=book&drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=6&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`历史`, `/pages/transfer/list?action=book&drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=5&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`游戏`, `/pages/transfer/list?action=book&drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=7&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`体育`, `/pages/transfer/list?action=book&drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=8&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`科幻`, `/pages/transfer/list?action=book&drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=9&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`灵异`, `/pages/transfer/list?action=book&drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=10&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`二次元`, `/pages/transfer/list?action=book&drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=12&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
		Card{`短篇`, `/pages/transfer/list?action=book&drive=qidian&url=` + EncodeURL(`https://www.qidian.com/all?chanId=20076&orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0`), "", `link`, ``, nil},
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

	page := FindString(`&page=(?P<page>(\d)+)&`, urlStr, "page")
	p, err := strconv.Atoi(page)
	if p > 0 && err == nil {
		// 已经组装url
		nextURL := strings.Replace(urlStr, fmt.Sprintf(`&page=%v&`, p), fmt.Sprintf(`&page=%v&`, p+1), -1)
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
	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)
	if catalogMsg == `` { //todo 从 https://book.qidian.com/ajax/book/category?_csrfToken=&bookId=1004608738 中获取章节列表(要解释json)
		// panic(`catalogMsg`)

		bookID := FindString(`/info/(?P<id>(\d)+)`, urlStr, "id")

		if bookID != `` {
			links, _ = r.GetChaptersLinksByJSON(bookID)
		}

	}

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

// GetChaptersLinksByJSON 获取章节链接列表
func (r QidianReader) GetChaptersLinksByJSON(bookID string) (links []Link, err error) {

	urlStr := fmt.Sprintf(`https://book.qidian.com/ajax/book/category?_csrfToken=&bookId=%v`, bookID)
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return
	}
	req.Header = make(http.Header)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	reader, err := charset.NewReader(resp.Body, strings.ToLower(resp.Header.Get("Content-Type")))
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(reader)

	type QiChaptersJsonDataCsChapter struct {
		UT          string `json:"uT"`
		ChapterName string `json:"cN"`
		ChapterURL  string `json:"cU"`
		UuID        int    `json:"uuid"`
		ID          int    `json:"id"`
		Ss          int    `json:"sS"`
	}

	type QiChaptersJsonDataCs struct {
		CCnt     int                           `json:"cCnt"`
		Chapters []QiChaptersJsonDataCsChapter `json:"cs"`
		// Chapters []map[string]interface{}      `json:"cs"`
		// Chapters map[int]interface{} `json:"cs"`
	}
	type QiChaptersJsonData struct {
		ChapterTotal int `json:"chapterTotalCnt"`
		// Vs           map[string]QiChaptersJsonDataCs `json:"vs"`
		Vs []QiChaptersJsonDataCs `json:"vs"`
		// Vs map[string]interface{} `json:"vs"`
		// Vs []interface{} `json:"vs"`
	}
	type QiChaptersJson struct {
		Code int `json:"code"`
		// Data map[string]interface{} `json:"data['vs']['cs']"`
		Data QiChaptersJsonData `json:"data"`
	}

	var m QiChaptersJson
	err = json.Unmarshal(bs, &m)

	if err == nil {
		for _, v := range m.Data.Vs {
			for _, vv := range v.Chapters {
				if vv.Ss == 1 {

					links = append(links, Link{
						vv.ChapterName,
						fmt.Sprintf(`https://read.qidian.com/chapter/%v`, vv.ChapterURL),
						``,
					})
				} else {

					links = append(links, Link{
						vv.ChapterName,
						fmt.Sprintf(`https://vipreader.qidian.com/chapter/%v/%v`, bookID, vv.ID),
						``,
					})
				}
			}
		}

	}
	return

}
