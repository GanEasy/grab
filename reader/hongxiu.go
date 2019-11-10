package reader

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//HongxiuReader 红袖添香
type HongxiuReader struct {
}

// GetCategories 获取所有自定义分类(写死)
func (r HongxiuReader) GetCategories(urlStr string) (list Catalog, err error) {

	list.Title = `分类-红袖添香`

	list.SourceURL = urlStr

	list.Cards = []Card{
		Card{`全部`, `/pages/list?drive=hongxiu&url=` + EncodeURL(`https://www.hongxiu.com/all?pageNum=1&pageSize=10&gender=2&catId=-1&isFinish=-1&isVip=-1&size=-1&updT=-1&orderBy=0&pageNum=1`), "", `link`, ``, nil, ``},
		Card{`现代言情`, `/pages/list?drive=hongxiu&url=` + EncodeURL(`https://www.hongxiu.com/all?pageNum=1&pageSize=10&gender=2&catId=30020&isFinish=-1&isVip=-1&size=-1&updT=-1&orderBy=0&pageNum=1`), "", `link`, ``, nil, ``},
		Card{`古代言情`, `/pages/list?drive=hongxiu&url=` + EncodeURL(`https://www.hongxiu.com/all?pageNum=1&pageSize=10&gender=2&catId=30013&isFinish=-1&isVip=-1&size=-1&updT=-1&orderBy=0&pageNum=1`), "", `link`, ``, nil, ``},
		Card{`浪漫青春`, `/pages/list?drive=hongxiu&url=` + EncodeURL(`https://www.hongxiu.com/all?pageNum=1&pageSize=10&gender=2&catId=30031&isFinish=-1&isVip=-1&size=-1&updT=-1&orderBy=0&pageNum=1`), "", `link`, ``, nil, ``},
		Card{`玄幻言情`, `/pages/list?drive=hongxiu&url=` + EncodeURL(`https://www.hongxiu.com/all?pageNum=1&pageSize=10&gender=2&catId=30001&isFinish=-1&isVip=-1&size=-1&updT=-1&orderBy=0&pageNum=1`), "", `link`, ``, nil, ``},
		Card{`仙侠奇缘`, `/pages/list?drive=hongxiu&url=` + EncodeURL(`https://www.hongxiu.com/all?pageNum=1&pageSize=10&gender=2&catId=30008&isFinish=-1&isVip=-1&size=-1&updT=-1&orderBy=0&pageNum=1`), "", `link`, ``, nil, ``},
		Card{`悬疑`, `/pages/list?drive=hongxiu&url=` + EncodeURL(`https://www.hongxiu.com/all?pageNum=1&pageSize=10&gender=2&catId=30036&isFinish=-1&isVip=-1&size=-1&updT=-1&orderBy=0&pageNum=1`), "", `link`, ``, nil, ``},
		Card{`科幻空间`, `/pages/list?drive=hongxiu&url=` + EncodeURL(`https://www.hongxiu.com/all?pageNum=1&pageSize=10&gender=2&catId=30042&isFinish=-1&isVip=-1&size=-1&updT=-1&orderBy=0&pageNum=1`), "", `link`, ``, nil, ``},
		Card{`游戏竞技`, `/pages/list?drive=hongxiu&url=` + EncodeURL(`https://www.hongxiu.com/all?pageNum=1&pageSize=10&gender=2&catId=30050&isFinish=-1&isVip=-1&size=-1&updT=-1&orderBy=0&pageNum=1`), "", `link`, ``, nil, ``},
		Card{`短篇小说`, `/pages/list?drive=hongxiu&url=` + EncodeURL(`https://www.hongxiu.com/all?pageNum=1&pageSize=10&gender=2&catId=30083&isFinish=-1&isVip=-1&size=-1&updT=-1&orderBy=0&pageNum=1`), "", `link`, ``, nil, ``},
		Card{`轻小说`, `/pages/list?drive=hongxiu&url=` + EncodeURL(`https://www.hongxiu.com/all?pageNum=1&pageSize=10&gender=2&catId=30055&isFinish=-1&isVip=-1&size=-1&updT=-1&orderBy=0&pageNum=1`), "", `link`, ``, nil, ``},
	}
	list.Hash = GetCatalogHash(list)
	return list, nil
}

// GetList 获取分类书籍列表
func (r HongxiuReader) GetList(urlStr string) (list Catalog, err error) {

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

	list.Title = `书籍列表-潇湘书院`

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://www.hongxiu.com/book/3756981504436501`, `https://www.hongxiu.com/book/22351265000781002`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `hongxiu`)

	list.SourceURL = urlStr

	// html := `{"I":"5333","V":"马经理"},`
	// page := FindString(`/p(?P<page>[^"]+)/`, html, "page")
	// class='page-next' href="javascript:" onclick="setPage(3)"
	// <div class="pagination" id="page-container" data-total="2750" data-size="10" data-page="1" data-url="/all?pageNum=1&amp;pageSize=10&amp;gender=2&amp;catId=-1&amp;isFinish=-1&amp;isVip=-1&amp;size=-1&amp;updT=-1&amp;orderBy=0"></div>

	total := FindString(`id="page-container" data-total="(?P<total>(\d)+)" data-size="10"`, html, "total")
	t, err1 := strconv.Atoi(total)
	page := FindString(`id="page-container" data-total="(?P<total>(\d)+)" data-size="10" data-page="(?P<page>(\d)+)" `, html, "page")
	p, err2 := strconv.Atoi(page)

	// log.Println(`xxx`, total, page)
	if p > 0 && p < t && err1 == nil && err2 == nil {
		// 已经组装url
		nextURL := strings.Replace(urlStr, fmt.Sprintf(`pageNum=%v`, p), fmt.Sprintf(`pageNum=%v`, p+1), -1)
		list.Next = Link{`下一页`, EncodeURL(nextURL), ``}
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetCatalog 获取章节列表
func (r HongxiuReader) GetCatalog(urlStr string) (list Catalog, err error) {

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

	list.Title = FindString(`下载_(?P<title>(.)+)无弹窗`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)
	links = CleaningFrontLinkRepeat(links) //保留最后一次出现的目录列表
	// log.Println(links)
	var needLinks []Link
	// if len(links) == 0 {

	// 	bookID := FindString(`/info/(?P<id>(\d)+)`, urlStr, "id")

	// 	if bookID != `` {
	// 		links, _ = r.GetChaptersLinksByHTML(bookID)
	// 		needLinks = links
	// 	}

	// }
	var state bool
	for _, l := range links { //
		l.URL, state = JaccardMateGetURL(l.URL, `https://www.hongxiu.com/chapter/3756981504436501/20291346771700823`, `https://www.hongxiu.com/chapter/13287843305142504/35669283073749147`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `hongxiu`)

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetInfo 获取章节正文内容
func (r HongxiuReader) GetInfo(urlStr string) (ret Content, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, ``)
	if err != nil {
		return ret, err
	}
	article, err := GetActicleByHTML(html)
	if err != nil {
		return ret, err
	}

	article.Readable(urlStr)

	ret.Title = article.Title

	ret.Title = FindString(`(?P<title>(.)+)_(?P<auther>(.)+)著`, article.Title, "title")
	if ret.Title == `` {
		ret.Title = article.Title
	}

	ret.SourceURL = urlStr

	reg := regexp.MustCompile(`<a([^<]+)<\/a>`)

	article.ReadContent = reg.ReplaceAllString(article.ReadContent, "")

	c := MarkDownFormatContent(article.ReadContent)

	c = BookContReplace(c)

	ret.Contents = GetSectionByContent(c)

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

// GetChaptersLinksByHTML 获取章节链接列表
func (r HongxiuReader) GetChaptersLinksByHTML(bookID string) (links []Link, err error) {

	//
	urlStr := fmt.Sprintf(`https://www.hongxiu.net/partview/GetChapterList?bookid=%v&noNeedBuy=0&special=0&maxFreeChapterId=0&isMonthly=0`, bookID)

	html, err := GetHTML(urlStr, ``)
	if err != nil {
		return links, err
	}
	html2 := fmt.Sprintf(`
		<html>
		<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
		<title>chapter list</title>
		<body>
		%v
		</body>
		</html>
		`, html)

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html2))
	if e != nil {
		return links, e
	}
	link, _ := url.Parse(urlStr)
	links = GetLinks(g2, link)
	return links, nil

}
