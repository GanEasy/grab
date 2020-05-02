package reader

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//Qkshu6Reader 去看书 (盗版小说网站)
type Qkshu6Reader struct {
}

// GetCategories 获取所有分类
func (r Qkshu6Reader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://m.qkshu6.com/`

	list.Title = `分类-去看书`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`都市小说`, `/pages/list?action=list&drive=qkshu6&url=` + EncodeURL(`https://m.qkshu6.com/dushi/`), "", `link`, ``, nil, ``},
		Card{`玄幻小说`, `/pages/list?action=list&drive=qkshu6&url=` + EncodeURL(`https://m.qkshu6.com/xuanhuan/`), "", `link`, ``, nil, ``},
		Card{`仙侠小说`, `/pages/list?action=list&drive=qkshu6&url=` + EncodeURL(`https://m.qkshu6.com/xianxia/`), "", `link`, ``, nil, ``},
		Card{`历史小说`, `/pages/list?action=list&drive=qkshu6&url=` + EncodeURL(`https://m.qkshu6.com/lishi/`), "", `link`, ``, nil, ``},
		Card{`科幻小说`, `/pages/list?action=list&drive=qkshu6&url=` + EncodeURL(`https://m.qkshu6.com/kehuan/`), "", `link`, ``, nil, ``},
		Card{`悬疑小说`, `/pages/list?action=list&drive=qkshu6&url=` + EncodeURL(`https://m.qkshu6.com/xuanyi/`), "", `link`, ``, nil, ``},
		Card{`其他小说`, `/pages/list?action=list&drive=qkshu6&url=` + EncodeURL(`https://m.qkshu6.com/qita/`), "", `link`, ``, nil, ``},
		Card{`全本小说`, `/pages/list?action=list&drive=qkshu6&url=` + EncodeURL(`https://m.qkshu6.com/quanben/`), "", `link`, ``, nil, ``},
		Card{`-----------------`, ``, "", `link`, ``, nil, ``},
		Card{`↘↘↘排行榜↙↙↙`, `/pages/list?action=list&drive=qkshu6&url=` + EncodeURL(`https://m.qkshu6.com/top.php`), "", `link`, ``, nil, ``},
	}
	list.SearchSupport = false
	return list, nil
}

// GetList 获取书籍列表列表
func (r Qkshu6Reader) GetList(urlStr string) (list Catalog, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, ``) //cover
	if err != nil {
		return
	}

	g, e := goquery.NewDocumentFromReader(strings.NewReader(html))

	if e != nil {
		return list, e
	}

	list.Title = FindString(`(?P<title>(.)+)_好看的`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	// log.Println(links)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURLMore(l.URL, `https://m.qkshu6.com/book/haosexiaoyi/`, `https://m.qkshu6.com/book/haoseyanfu/`, ``) //
		if state {
			l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `qkshu6`)

	list.SourceURL = urlStr

	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// Search 搜索资源
func (r Qkshu6Reader) Search(keyword string) (list Catalog, err error) {

	return list, nil
}

// GetCatalog 获取章节列表
func (r Qkshu6Reader) GetCatalog(urlStr string) (list Catalog, err error) {

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

	list.Title = FindString(`(?P<title>(.)+)全文阅读_`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = FindString(`(?P<title>(.)+)最新章节`, g.Find("title").Text(), "title")
		if list.Title == `` {
			list.Title = g.Find("title").Text()
		}
	}

	link, _ := url.Parse(urlStr)

	// html2, _ := g.Find(`.chapter`).Eq(0).Html()

	// g2, e := goquery.NewDocumentFromReader(strings.NewReader(html2))

	// var links = GetLinks(g2, link)
	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://m.qkshu6.com/book/ylshzwlxs/10930.html`, `https://m.qkshu6.com/book/haoseyanfu/10020.html`, ``)
		// l.URL, state = JaccardMateGetURL(l.URL, `http://m.qkshu6.com/wapbook-135411-170696662/`, `http://m.qkshu6.com/wapbook-1011-783829/`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	needLinks = CleaningFrontRepeat(needLinks)

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `qkshu6`)

	list.SourceURL = urlStr

	var links2 = GetLinks(g, link)
	list.Next = GetNextLink(links2)

	// list.Next = GetNextLink(links)

	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}
	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetInfo 获取详细内容
func (r Qkshu6Reader) GetInfo(urlStr string) (ret Content, err error) {

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

	ret.Title = FindString(`(?P<title>(.)+)_(?P<bookname>(.)+)_笔趣阁`, article.Title, "title")
	if ret.Title == `` {
		ret.Title = article.Title
	}

	reg := regexp.MustCompile(`<a([^>]+)>([^<]+)<\/a>`)

	article.ReadContent = reg.ReplaceAllString(article.ReadContent, "")

	reg2 := regexp.MustCompile(`（本章未完，请点击下一页继续阅读）`)

	article.ReadContent = reg2.ReplaceAllString(article.ReadContent, "")

	reg3 := regexp.MustCompile(`try{([^<]+)} catch\(ex\){}`)

	article.ReadContent = reg3.ReplaceAllString(article.ReadContent, "")

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
