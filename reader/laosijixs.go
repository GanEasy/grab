package reader

import (
	"context"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/yizenghui/chromedp"
)

//LaosijixsReader 顶点小说 (盗版小说网站)
type LaosijixsReader struct {
}

// GetCategories 获取所有分类
func (r LaosijixsReader) GetCategories(urlStr string) (list Catalog, err error) {

	// urlStr := `http://m.laosijixs.com/`

	list.Title = `分类-老司机小说`

	list.SourceURL = urlStr

	list.Hash = GetCatalogHash(list)

	list.Cards = []Card{
		Card{`全部`, `/pages/list?action=list&drive=laosijixs&url=` + EncodeURL(`http://m.laosijixs.com/shuku/`), "", `link`, ``, nil, ``},
	}
	return list, nil
}

// GetList 获取书籍列表列表
func (r LaosijixsReader) GetList(urlStr string) (list Catalog, err error) {

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

	list.Title = `资源列表-老司机小说`

	link, _ := url.Parse(urlStr)

	var links = GetLinks(g, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `http://m.laosijixs.com/20/20961/`, `http://www.laosijixs.com/19/19634/`, ``)
		if state {
			// l.Title = FindString(`(?P<title>(.)+)`, l.Title, "title")
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/catalog`, `laosijixs`)

	list.SourceURL = urlStr

	list.Next = GetNextLink(links)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}

	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetCatalog 获取章节列表
func (r LaosijixsReader) GetCatalog(urlStr string) (list Catalog, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
	html, err := GetHTML(urlStr, ``)
	// html, err := GetHTMLByChromedp(urlStr)
	if err != nil {
		return
	}

	g, e := goquery.NewDocumentFromReader(strings.NewReader(html))

	if e != nil {
		return list, e
	}

	list.Title = FindString(`(?P<title>(.)+)_全文阅读`, g.Find("title").Text(), "title")
	if list.Title == `` {
		list.Title = g.Find("title").Text()
	}

	link, _ := url.Parse(urlStr)

	html2, _ := g.Find(`.chapter-list`).Eq(1).Html()

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html2))

	var links = GetLinks(g2, link)

	var needLinks []Link
	var state bool
	for _, l := range links {
		l.URL, state = JaccardMateGetURL(l.URL, `http://m.laosijixs.com/20/20961/546047.html`, `http://m.laosijixs.com/79/79525/5713401.html`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}

	list.Cards = LinksToCards(Cleaning(needLinks), `/pages/book`, `laosijixs`)

	list.SourceURL = urlStr

	var links2 = GetLinks(g, link)

	list.Next = GetNextLink(links2)
	if list.Next.URL != `` {
		list.Next.URL = EncodeURL(list.Next.URL)
	}
	list.Hash = GetCatalogHash(list)

	return list, nil

}

// GetInfox 获取详细内容
func (r LaosijixsReader) GetInfox(urlStr string) (ret Content, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var title, res string
	var jres []string
	// var res2 []string
	err = chromedp.Run(ctx,
		chromedp.Navigate(urlStr),
		chromedp.Title(&title),
		chromedp.Sleep(time.Second*6),
		chromedp.Evaluate(`function() {$('#content').find('span').remove();return { body: $('#content').innerText};}`, &jres),
		chromedp.Text(`#content'`, &res, chromedp.NodeVisible, chromedp.ByQuery),
	)

	if err != nil {
		return ret, err
	}
	ret.Title = title

	ret.SourceURL = urlStr

	c := MarkDownFormatContent(res)

	c = BookContReplace(c)

	ret.Contents = GetSectionByContent(c)

	return ret, nil

}

// GetInfoBodyText 获取详细内容
func (r LaosijixsReader) GetInfoBodyText(urlStr string) (html, body string, err error) {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	// var res, res1 string
	var res2 []string
	err = chromedp.Run(ctx,
		chromedp.Navigate(urlStr),
		chromedp.Sleep(time.Second*2),
		//$('#content').find('span').remove();
		chromedp.Evaluate(`Object.keys(window);`, &res2),
		// chromedp.Body(`html`, &res),

		// chromedp.Text(`html`, &res1, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#content`, &body, chromedp.NodeVisible, chromedp.ByID),
		chromedp.OuterHTML("html", &html),
	)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(strings.TrimSpace(res))
	// log.Println(res1)

	return html, body, err

	// ctx, cancel := chromedp.NewContext(context.Background())
	// defer cancel()
	// err = chromedp.Run(ctx,
	// 	chromedp.Navigate(urlStr),
	// 	chromedp.Sleep(time.Second*3),
	// 	chromedp.OuterHTML("html", &html),
	// 	chromedp.Sleep(time.Second*2),
	// 	chromedp.Text(`#content'`, &body, chromedp.NodeVisible, chromedp.ByID),
	// )
	// return html, body, err
}

// GetInfo 获取详细内容
func (r LaosijixsReader) GetInfo(urlStr string) (ret Content, err error) {

	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}

	log.Println(`GetInfo`, urlStr)
	html, body, err := r.GetInfoBodyText(urlStr)
	// html, err := GetHTML(urlStr, ``)
	// log.Println(`GetInfoBodyText`, body, html, err)
	// html, err := GetHTMLByChromedp(urlStr)
	if err != nil {
		return ret, err
	}
	article, err := GetActicleByHTML(html)
	if err != nil {
		return ret, err
	}
	article.Readable(urlStr)

	regc1 := regexp.MustCompile(`<span([^>]*)>([^<]+)<\/span>`)
	article.ReadContent = regc1.ReplaceAllString(article.ReadContent, "")

	c1 := MarkDownFormatContent(article.ReadContent)

	c1 = BookContReplace(c1)
	// log.Println(`article.ReadContent`, article.ReadContent)

	c1Contents := GetSectionByContent(c1)

	// var edu = map[string]int{}
	var c2Contents = map[string]int{}
	for _, v := range c1Contents { //所有内容
		if v != `` {
			c2Contents[v] = 1
			// log.Println(`c1Contents`, v)
		}
	}

	// log.Println(`c1Contents,c2Contents`, c2Contents)

	ret.Title = FindString(`(?P<title>(.)+)_(?P<bookname>(.)+)_(?P<category>(.)+)_`, article.Title, "title")
	if ret.Title == `` {
		ret.Title = article.Title
	}

	if body != `` {
		article.ReadContent = body
	}
	if err != nil {
		return ret, err
	}
	reg := regexp.MustCompile(`<span([^>]+)>([^<]+)<\/span>`)
	article.ReadContent = reg.ReplaceAllString(article.ReadContent, "")

	reg2 := regexp.MustCompile(`努力加载中...超过5秒钟未打开,请刷新一下！`)

	article.ReadContent = reg2.ReplaceAllString(article.ReadContent, "")

	ret.SourceURL = urlStr

	c := MarkDownFormatContent(article.ReadContent)

	c = BookContReplace(c)

	ret.Contents = GetSectionByContent(c)

	var retContents []string
	for _, v := range ret.Contents {
		// if i%2 == 0 {
		// 	retContents = append(retContents, v)
		// 	log.Println(`Contents`, i, v)
		// 	// delete(ret.Contents, i)
		// }

		if _, ok := c2Contents[v]; ok && v != `` {
			retContents = append(retContents, v)
			// log.Println(`Contents`, i, v)
		} else {
			// log.Println(`dd Contents`, i, v)
		}

		// if v != `` {
		// 	_, ok := c2Contents[v]
		// 	if ok {
		// 		log.Println(`c2Contents`, i, v)
		// 	} else {
		// 		log.Println(`dd c2Contents`, i, v)
		// 	}
		// }
	}
	ret.Contents = retContents

	links, _ := GetLinkByHTML(urlStr, html)
	ret.Previous = GetPreviousLink(links)
	if ret.Previous.URL != `` {
		ret.Previous.URL = EncodeURL(ret.Previous.URL)
	}

	g, e := goquery.NewDocumentFromReader(strings.NewReader(html))
	if e != nil {
		//
	}
	html2, _ := g.Find(`.chapterPages`).Html()

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html2))

	link, _ := url.Parse(urlStr)
	var links2 = GetLinks(g2, link)

	var needLinks []Link
	var state bool
	for _, l := range links2 {
		l.URL, state = JaccardMateGetURL(l.URL, `http://m.laosijixs.com/20/20961/546056_1.html`, `http://m.laosijixs.com/80/80894/5905659_2.html`, ``)
		if state {
			needLinks = append(needLinks, l)
		}
	}
	var thisPage = 0
	if len(needLinks) > 1 {
		for i, l := range needLinks {
			if thisPage == 0 && l.URL == urlStr {
				thisPage = i
				// log.Println(`thisPage`, thisPage)
			}
		}
	}

	if len(needLinks) > (thisPage + 1) {
		ret.Next.URL = EncodeURL(needLinks[thisPage+1].URL)
	} else {
		ret.Next = GetNextLink(links)
		if ret.Next.URL != `` {
			ret.Next.URL = EncodeURL(ret.Next.URL)
		}
	}

	return ret, nil

}
