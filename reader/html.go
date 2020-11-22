package reader

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/axgle/mahonia"
	"github.com/chromedp/chromedp"
	"golang.org/x/net/html/charset"
)

//ReplaceImageServe 替换图片服务地址(由服务器转抓取)
func ReplaceImageServe(body string) (string, error) {
	article, err := GetActicleByContent(body)
	if err != nil {
		return body, err
	}
	for _, i := range article.Images {
		body = strings.Replace(body, i, fmt.Sprintf(`https://img.readfollow.com/file?url=%v`, i), -1)
	}
	return body, nil
}

// GetHTMLContent 获取html链接地址中的链接
func GetHTMLContent(urlStr, find string) (cont FetchContent, err error) {
	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
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
	htmlStr := string(bs)

	htmlStr, err = FindContentHTML(htmlStr, find)

	cont.Title = `HTML.Title`
	// info.Content = article.ReadContent
	cont.Content = htmlStr
	cont.PubAt = ""
	cont.SourceURL = urlStr
	return cont, err
}

// GetHTMLByChromedp 通过运行浏览器代理的方式获取被渲染后的网站内容
func GetHTMLByChromedp(urlStr string) (htmlStr string, err error) {

	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err = chromedp.Run(ctx,
		// chromedp.Navigate(`https://github.com/chromedp/examples`),
		// chromedp.Text(`.Box-body`, &htmlStr, chromedp.NodeVisible, chromedp.ByQuery),
		// chromedp.Navigate(`https://m.138txt.com/193/193028/`),
		// chromedp.Text(`body`, &htmlStr, chromedp.NodeVisible, chromedp.ByQuery),
		// chromedp.OuterHTML(`body`, &htmlStr, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Navigate(urlStr),
		// chromedp.WaitReady(`html`, chromedp.ByQuery),
		chromedp.Sleep(time.Second*3),
		chromedp.OuterHTML(`html`, &htmlStr),
		// chromedp.Body(`html`, &htmlStr, chromedp.NodeVisible, chromedp.ByQuery),

		// chromedp.Navigate(`http://m.laosijixs.com/20/20961/546056_5.html`),
		// chromedp.Text(`#content`, &htmlStr, chromedp.NodeVisible, chromedp.ByID),
		// chromedp.Text(`#pkg-overview`, &htmlStr, chromedp.NodeVisible, chromedp.ByID),
	)
	return htmlStr, err

}

// GetHTML 获取html链接地址中的内容
func GetHTML(urlStr, find string) (htmlStr string, err error) {
	err = CheckStrIsLink(urlStr)
	if err != nil {
		return
	}
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

	htmlStr, err = FindContentHTML(htmlStr, find)
	return htmlStr, err
}

// GetHTMLOrCache 获取html链接地址中的内容(进行本地化数据缓存)
func GetHTMLOrCache(urlStr, find string) (string, error) {
	err := CheckStrIsLink(urlStr)
	if err != nil {
		return ``, err
	}

	htmlpath := fmt.Sprintf("cache/%v.cache", GetMd5String(urlStr))

	_, err = os.Stat(htmlpath)
	if os.IsNotExist(err) {
		_, err2 := SaveHTML(urlStr, htmlpath)
		if err2 != nil {
			return ``, err2
		}
	}
	f, e := os.Open(htmlpath)
	if e != nil {
		return ``, e
	}
	defer f.Close()

	if htmlStr, err := ioutil.ReadAll(f); err == nil {
		return FindContentHTML(string(htmlStr), find)
	}

	return ``, err
}

//Substr 截取字符串 start 起点下标 end 终点下标(不包括)
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return `nil`
	}

	if end < 0 || end > length {
		return `nil`
	}

	return string(rs[start:end])
}

//GetMd5String 生成32位md5字串路径
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	// return hex.EncodeToString(h.Sum(nil))
	mddir := hex.EncodeToString(h.Sum(nil))
	dir := Substr(mddir, 0, 3) + `/` + Substr(mddir, 3, 6) + `/` + Substr(mddir, 6, 32)
	return dir
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// SaveHTML 保存HTML到本地
func SaveHTML(urlStr, saveName string) (n int64, err error) {
	filePath := path.Dir(saveName)
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return 0, err
		}
	}
	out, err := os.Create(saveName)
	defer out.Close()
	if err != nil {
		return
	}

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

	n, err = io.Copy(out, bytes.NewReader(bs))

	if err != nil {
		return
	}
	return
}

// 字符串转编码
func ConvertStrEncode(inStr, inCharset, outCharset string) string {
	if outCharset == "" {
		outCharset = inCharset
	}

	inCharset = strings.ToLower(inCharset)
	outCharset = strings.ToLower(outCharset)

	if inCharset == outCharset {
		return inStr
	}

	if inCharset == "gbk" || inCharset == "gb2312" {
		inCharset = "gb18030"
	}

	// 输入字符串解码为utf-8
	var destr string
	if inCharset != "utf8" && inCharset != "utf-8" {
		destr = mahonia.NewDecoder(inCharset).ConvertString(inStr)
	} else {
		destr = inStr
	}

	if outCharset == "utf8" || outCharset == "utf-8" {
		return destr
	}
	// 转换为 outCharset
	return mahonia.NewEncoder(outCharset).ConvertString(destr)
}
