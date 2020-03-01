package reader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"golang.org/x/net/html/charset"
)

func Test_XxsyGetBooks(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	urlStr := "https://www.xxsy.net/search?s_wd=&channel=2&sort=9&pn=1"
	reader := XxsyReader{}
	list, err := reader.GetList(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_XxsySearch(t *testing.T) {
	reader := XxsyReader{}
	list, err := reader.Search(`山河`)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_XxsyGetInfo(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	urlStr := "https://www.xxsy.net/chapter/79588269.html"
	urlStr = "https://www.xxsy.net/chapter/34109344.html"
	reader := XxsyReader{}
	list, err := reader.GetInfo(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}
func Test_XxsyGetChapters(t *testing.T) {
	urlStr := "https://www.xxsy.net/info/959418.html"
	// urlStr = "https://www.xxsy.net/info/1079349.html"

	reader := XxsyReader{}
	list, err := reader.GetCatalog(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

// func GetMapByMap(key string, m map[string]interface{}) (ret map[string]interface{}, err error) {

// 	var m2 map[string]interface{}
// 	if data, ok := m["data"]; ok {
// 		err = json.Unmarshal([]byte(m["data"]), &m2)

// 	}
// }

func Test_XxsyGetChaptersByJson(t *testing.T) {

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
	bookID := `1004608738`
	urlStr := `https://book.qidian.com/ajax/book/category?_csrfToken=&bookId=1004608738`
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

	var m QiChaptersJson
	err = json.Unmarshal(bs, &m)

	var links []Link

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
	t.Fatal(links)

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
