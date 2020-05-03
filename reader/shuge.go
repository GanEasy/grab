package reader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//ShugeReader U小说阅读网 (盗版小说网站)
type ShugeReader struct {
}

// GetCategories 获取所有分类
func (r ShugeReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://m.xbiquge.com/`

	list.Title = `分类-shuge.la`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`玄幻奇幻`, `/pages/list?action=book&drive=shuge&url=` + EncodeURL(`https://m.shuge.la/sort/1_1/`), "", `link`, ``, nil, ``},
		Card{`武侠仙侠`, `/pages/list?action=book&drive=shuge&url=` + EncodeURL(`https://m.shuge.la/sort/2_1/`), "", `link`, ``, nil, ``},
		Card{`都市异能`, `/pages/list?action=book&drive=shuge&url=` + EncodeURL(`https://m.shuge.la/sort/3_1/`), "", `link`, ``, nil, ``},
		Card{`军事历史`, `/pages/list?action=book&drive=shuge&url=` + EncodeURL(`https://m.shuge.la/sort/4_1/`), "", `link`, ``, nil, ``},
		Card{`游戏竞技`, `/pages/list?action=book&drive=shuge&url=` + EncodeURL(`https://m.shuge.la/sort/5_1/`), "", `link`, ``, nil, ``},
		Card{`科幻世界`, `/pages/list?action=book&drive=shuge&url=` + EncodeURL(`https://m.shuge.la/sort/6_1/`), "", `link`, ``, nil, ``},
		Card{`灵异悬疑`, `/pages/list?action=book&drive=shuge&url=` + EncodeURL(`https://m.shuge.la/sort/7_1/`), "", `link`, ``, nil, ``},
		Card{`耽美同人`, `/pages/list?action=book&drive=shuge&url=` + EncodeURL(`https://m.shuge.la/sort/8_1/`), "", `link`, ``, nil, ``},
		Card{`女生言情`, `/pages/list?action=book&drive=shuge&url=` + EncodeURL(`https://m.shuge.la/sort/9_1/`), "", `link`, ``, nil, ``},
	}
	list.SearchSupport = true
	return list, nil
}

// GetList 获取书籍列表列表
func (r ShugeReader) GetList(urlStr string) (list Catalog, err error) {

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
	list.Title = FindString(`(?P<title>(.)+)-书阁小说网`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://m.shuge.la/read/7/7523/`, `https://m.shuge.la/read/0/657/`, ``)
		if state {
			l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `shuge`)

	list.SourceURL = urlStr

	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// Search 搜索资源
func (r ShugeReader) Search(keyword string) (list Catalog, err error) {

	urlStr := `https://m.shuge.la/s.php`

	//搜索关键字 utf8转gbk (源站gbk)
	gbkkeyword := ConvertStrEncode(keyword, "utf-8", "gbk")

	client := http.Client{}
	data := make(url.Values)
	data["type"] = []string{`articlename`}
	data["s"] = []string{gbkkeyword}
	data["submit"] = []string{``}
	// 提交表单数据
	resp, err := client.PostForm("https://m.shuge.la/s.php", data)
	if err != nil {
		return
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return
	}
	//byte数组直接转成string，优化内存
	html := ConvertStrEncode(string(respBytes), "gbk", "utf-8")
	// log.Println(html)

	g, e := goquery.NewDocumentFromReader(strings.NewReader(html))

	if e != nil {
		return list, e
	}

	list.Title = fmt.Sprintf(`%v - 搜索结果 - shuge.la`, keyword)

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)
	// log.Println(links)
	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://m.shuge.la/read/7/7523/`, `https://m.shuge.la/read/0/657/`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `shuge`)

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	return list, nil
}

// GetCatalog 获取章节列表
func (r ShugeReader) GetCatalog(urlStr string) (list Catalog, err error) {

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
	list.Title = FindString(`(?P<title>(.)+)小说最新章节_`, g.Find("title").Text(), "title")

	if list.Title == `` {
		list.Title = FindString(`(?P<title>(.)+)最新章节_`, g.Find("title").Text(), "title")
	}
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	html2, _ := FindContentHTML(html, `#chapterList`)

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html2))

	var links = GetLinks(g2, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `https://m.shuge.la/read/5/5804/2841238.html`, `https://m.shuge.la/read/0/657/`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `shuge`)

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
func (r ShugeReader) GetInfo(urlStr string) (ret Content, err error) {

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

	ret.Title = FindString(`(?P<bookname>(.)+) (?P<title>(.)+)-书阁小说网`, article.Title, "title")
	if ret.Title == `` {
		ret.Title = article.Title
	}

	reg := regexp.MustCompile(`<a([^<]+)<\/a>`)

	article.ReadContent = reg.ReplaceAllString(article.ReadContent, "")

	reg2 := regexp.MustCompile(`（本章未完，请点击下一页继续阅读）`)

	article.ReadContent = reg2.ReplaceAllString(article.ReadContent, "")

	// reg3 := regexp.MustCompile(`&gt;&gt;`)

	// article.ReadContent = reg3.ReplaceAllString(article.ReadContent, "")

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
