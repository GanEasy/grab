package reader

import (
	"io/ioutil"
	"net/http"

	"github.com/GanEasy/html2article"
)

//GetHTML 获取html
func GetHTML(urlStr string) (htmlStr string, err error) {
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
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)
	htmlStr = string(bs)
	htmlStr = html2article.DecodeHtml(resp.Header, htmlStr, htmlStr)
	return htmlStr, err
}
