package reader

import (
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	// "net/http"
	"net/url"
	"regexp"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	// "golang.org/x/net/html/charset"
)

func Test_BaiduReaderGetGetList(t *testing.T) {
	// https://www.feixuemh.com/chapter/54380
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `https://top.baidu.com/board?platform=pc&tab=novel&tag={"category":"玄幻"}`
	reader := BaiduReader{}
	list, err := reader.GetList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

// 通过百度热搜来获取数据
func Test_BaiDuHotSearch(t *testing.T) {
	var urlStr = `https://top.baidu.com/board?tab=novel`
	// 全部类型 都市 玄幻 历史 武侠 现代言情 古代言情 青春    奇幻 科幻  军事 游戏 幻想言情
	urlStr = `https://top.baidu.com/board?platform=pc&tab=novel&tag={"category":"青春"}`
	err := CheckStrIsLink(urlStr)
	if err != nil {
		t.Fatal(err)
	}
	html, err := GetHTML(urlStr, `.container-bg_lQ801`)
	if err != nil {
		t.Fatal(err)
	}

	link, _ := url.Parse(urlStr)

	// html2, _ := g.Find(`#detail-list-select`).Eq(1).Html()

	g2, e := goquery.NewDocumentFromReader(strings.NewReader(html))

	var links = GetLinks(g2, link)
	if e != nil {
		t.Fatal(e)
	}
	// t.Fatal(links)
	var needLinks []Link
	pattern := "\\d+" //反斜杠要转义
	for _, l := range links {
		if l.Title != `查看更多>` {
			match, _ := regexp.MatchString(pattern, l.Title)
			if !match {
				l.URL = fmt.Sprintf(`/pages/search?keyword=%v`, l.Title)

				needLinks = append(needLinks, l)
			}
		}
	}

	t.Fatal(needLinks)
	t.Fatal(html)

	// type QiChaptersJsonDataCsChapter struct {
	// 	UT          string `json:"uT"`
	// 	ChapterName string `json:"cN"`
	// 	ChapterURL  string `json:"cU"`
	// 	UuID        int    `json:"uuid"`
	// 	ID          int    `json:"id"`
	// 	Ss          int    `json:"sS"`
	// }

	// type QiChaptersJsonDataCs struct {
	// 	CCnt     int                           `json:"cCnt"`
	// 	Chapters []QiChaptersJsonDataCsChapter `json:"cs"`
	// 	// Chapters []map[string]interface{}      `json:"cs"`
	// 	// Chapters map[int]interface{} `json:"cs"`
	// }
	// type QiChaptersJsonData struct {
	// 	ChapterTotal int `json:"chapterTotalCnt"`
	// 	// Vs           map[string]QiChaptersJsonDataCs `json:"vs"`
	// 	Vs []QiChaptersJsonDataCs `json:"vs"`
	// 	// Vs map[string]interface{} `json:"vs"`
	// 	// Vs []interface{} `json:"vs"`
	// }
	// type QiChaptersJson struct {
	// 	Code int `json:"code"`
	// 	// Data map[string]interface{} `json:"data['vs']['cs']"`
	// 	Data QiChaptersJsonData `json:"data"`
	// }
	// bookID := `1004608738`
	// urlStr := `https://book.qidian.com/ajax/book/category?_csrfToken=&bookId=1004608738`
	// req, err := http.NewRequest("GET", urlStr, nil)
	// if err != nil {
	// 	return
	// }
	// req.Header = make(http.Header)
	// req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36")
	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	return
	// }

	// reader, err := charset.NewReader(resp.Body, strings.ToLower(resp.Header.Get("Content-Type")))
	// defer resp.Body.Close()
	// bs, _ := ioutil.ReadAll(reader)

	// var m QiChaptersJson
	// err = json.Unmarshal(bs, &m)

	// var links []Link

	// if err == nil {
	// 	for _, v := range m.Data.Vs {
	// 		for _, vv := range v.Chapters {
	// 			if vv.Ss == 1 {

	// 				links = append(links, Link{
	// 					vv.ChapterName,
	// 					fmt.Sprintf(`https://read.qidian.com/chapter/%v`, vv.ChapterURL),
	// 					``,
	// 				})
	// 			} else {

	// 				links = append(links, Link{
	// 					vv.ChapterName,
	// 					fmt.Sprintf(`https://vipreader.qidian.com/chapter/%v/%v`, bookID, vv.ID),
	// 					``,
	// 				})
	// 			}
	// 		}
	// 	}

	// }
	// t.Fatal(links)

	// var dat map[string]interface{}
	// if err := json.Unmarshal(bs, &dat); err == nil {
	// 	fmt.Println("==============json str 转map=======================")
	// 	fmt.Println(dat)

	// 	mapTmp := dat["data"].(map[string]interface{})
	// 	fmt.Println(mapTmp["id"])
	// 	/*
	// 	   var dat2 map[string]interface{}
	// 	   if err := json.Unmarshal([]byte(jsonStr), &dat2); err == nil {
	// 	       fmt.Println( dat2["firstName"])
	// 	   }
	// 	*/

	// 	mapTmp2 := (dat["data"].([]interface{}))[0].(map[string]interface{})
	// 	//mapTmp3 := mapTmp2[0].(map[string]interface {})
	// 	fmt.Println(mapTmp2["id"])

	// }
}
