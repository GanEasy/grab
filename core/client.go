package core

import (
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html/charset"
)

//GetHTML get请求 传入链接地址 获取utf-8格式的 html
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

	reader, err := charset.NewReader(resp.Body, strings.ToLower(resp.Header.Get("Content-Type")))
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(reader)
	htmlStr = string(bs)
	return htmlStr, err
}
