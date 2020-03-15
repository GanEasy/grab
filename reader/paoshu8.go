package reader

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//Paoshu8Reader 顶点小说 (盗版小说网站)
type Paoshu8Reader struct {
}

// GetCategories 获取所有分类
func (r Paoshu8Reader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://m.paoshu8.com/`

	list.Title = `分类-笔趣阁paoshu8`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`全部小说`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/top-lastupdate-1/`), "", `link`, ``, nil, ``},
		Card{`玄幻小说`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/sort-1-1/`), "", `link`, ``, nil, ``},
		Card{`仙侠小说`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/sort-2-1/`), "", `link`, ``, nil, ``},
		Card{`都市小说`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/sort-3-1/`), "", `link`, ``, nil, ``},
		Card{`历史小说`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/sort-4-1/`), "", `link`, ``, nil, ``},
		Card{`游戏小说`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/sort-5-1/`), "", `link`, ``, nil, ``},
		Card{`科幻小说`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/sort-6-1/`), "", `link`, ``, nil, ``},
		Card{`言情小说`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/sort-7-1/`), "", `link`, ``, nil, ``},
		Card{`同人小说`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/sort-8-1/`), "", `link`, ``, nil, ``},
		Card{`灵异小说`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/sort-9-1/`), "", `link`, ``, nil, ``},
		Card{`奇幻小说`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/sort-10-1/`), "", `link`, ``, nil, ``},
		Card{`竞技小说`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/sort-11-1/`), "", `link`, ``, nil, ``},
		Card{`武侠小说`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/sort-12-1/`), "", `link`, ``, nil, ``},
		Card{`军事小说`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/sort-13-1/`), "", `link`, ``, nil, ``},
		Card{`校园小说`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/sort-14-1/`), "", `link`, ``, nil, ``},
		Card{`官场小说`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/sort-15-1/`), "", `link`, ``, nil, ``},
		Card{`↘↘↘排行榜↙↙↙`, ``, "", `link`, ``, nil, ``},
		Card{`日点击榜`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/top-dayvisit-1/`), "", `link`, ``, nil, ``},
		Card{`周点击榜`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/top-weekvisit-1/`), "", `link`, ``, nil, ``},
		Card{`月点击榜`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/top-monthvisit-1/`), "", `link`, ``, nil, ``},
		Card{`总点击榜`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`http://m.paoshu8.com/top-allvisit-1/`), "", `link`, ``, nil, ``},
		// Card{`全部↘↘↘`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`https://m.paoshu8.com/waptop/1.html`), "", `link`, ``, nil, ``},
		// Card{`玄幻↘↘↘`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`https://m.paoshu8.com/waptop/1_1.html`), "", `link`, ``, nil, ``},
		// Card{`修真↘↘↘`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`https://m.paoshu8.com/waptop/2_1.html`), "", `link`, ``, nil, ``},
		// Card{`都市↘↘↘`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`https://m.paoshu8.com/waptop/3_1.html`), "", `link`, ``, nil, ``},
		// Card{`穿越↘↘↘`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`https://m.paoshu8.com/waptop/4_1.html`), "", `link`, ``, nil, ``},
		// Card{`网游↘↘↘`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`https://m.paoshu8.com/waptop/5_1.html`), "", `link`, ``, nil, ``},
		// Card{`科幻↘↘↘`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`https://m.paoshu8.com/waptop/6_1.html`), "", `link`, ``, nil, ``},
		// Card{`其他↘↘↘`, `/pages/list?action=list&drive=paoshu8&url=` + EncodeURL(`https://m.paoshu8.com/waptop/7_1.html`), "", `link`, ``, nil, ``},
	}
	list.SearchSupport = true
	return list, nil
}

// GetList 获取书籍列表列表
func (r Paoshu8Reader) GetList(urlStr string) (list Catalog, err error) {

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

	list.Title = FindString(`(?P<title>(.)+)-笔趣阁`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	// log.Println(links)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURLMore(l.URL, `http://m.paoshu8.com/info-135411/`, `http://m.paoshu8.com/info-139226/`, `http://www.paoshu8.com/0_135411/`) //`http://m.paoshu8.com/wapbook-135411_1/`
		if state {
			l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `paoshu8`)

	list.SourceURL = urlStr

	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// Search 搜索资源
func (r Paoshu8Reader) Search(keyword string) (list Catalog, err error) {
	urlStr := `http://www.paoshu8.com/modules/article/search.php?searchkey=` + keyword
	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, `.grid`)
	if err != nil {
		return
	}

	g, e := goquery.NewDocumentFromReader(strings.NewReader(html))

	if e != nil {
		return list, e
	}

	list.Title = fmt.Sprintf(`%v - 搜索结果-paoshu8.com`, keyword)

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
		l.URL, state = JaccardMateGetURL(l.URL, `http://www.paoshu8.com/88_88522/`, `http://www.paoshu8.com/888_888522/`, ``)
		if state {
			l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `paoshu8`)

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	return list, nil
}

// GetCatalog 获取章节列表
func (r Paoshu8Reader) GetCatalog(urlStr string) (list Catalog, err error) {

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

	list.Title = FindString(`(?P<title>(.)+)无弹窗`, g.Find("title").Text(), "title")
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
		l.URL, state = JaccardMateGetURL(l.URL, `http://www.paoshu8.com/132_132325/170746501.html`, `http://www.paoshu8.com/42_42714/16445540.html`, ``)
		// l.URL, state = JaccardMateGetURL(l.URL, `http://m.paoshu8.com/wapbook-135411-170696662/`, `http://m.paoshu8.com/wapbook-1011-783829/`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	needLinks = CleaningFrontRepeat(needLinks)

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `paoshu8`)

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
func (r Paoshu8Reader) GetInfo(urlStr string) (ret Content, err error) {

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

	reg3 := regexp.MustCompile(`try{content1\(\);} catch\(ex\){}`)

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
