package reader

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//FindContentForHTML 从html字符串中获取内容 通过 .class 或 #id 或 name
func FindContentForHTML(htmlStr, find string) (string, error) {
	if find == `` {
		return htmlStr, errors.New(`empty find`)
	}
	g, e := goquery.NewDocumentFromReader(strings.NewReader(htmlStr))

	if e != nil {
		return htmlStr, e
	}
	return g.Find(find).Html()
}

//GetURLStringParam 从url字符串中获取指定问号后的参数 	urlStr := "/pages/catalog?drive=xbiquge&url=aHR0cDovL3d3dy54YmlxdWdlLmxhLzE1LzE1MDIxLw%3D%3D"
// 返回 aHR0cDovL3d3dy54YmlxdWdlLmxhLzE1LzE1MDIxLw==
func GetURLStringParam(urlStr, key string) (string, error) {
	u, e := url.Parse(urlStr)
	if e != nil {
		return ``, e
	}
	m, e2 := url.ParseQuery(u.RawQuery)
	if e2 != nil {
		return ``, e
	}
	if v, ok := m[key]; ok {
		return v[0], nil
	}
	return ``, e
}

//ContentBuildHTML 内容合成 html
func ContentBuildHTML(content, title string) string {
	return fmt.Sprintf(`
		<html>
		<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
		<title>%v</title>
		<body>
		%v
		</body>
		</html>
		`, title, content)
}

// FindContentHTML 获得纯正文html
func FindContentHTML(htmlStr, find string) (string, error) {

	if find == `` {
		return htmlStr, nil
	}
	g, e := goquery.NewDocumentFromReader(strings.NewReader(htmlStr))
	if e != nil {
		return htmlStr, e
	}
	title := g.Find("title").Eq(0).Text()

	html, err := g.Find(find).Html()
	if err != nil {
		return htmlStr, err
	}
	return fmt.Sprintf(`
		<html>
		<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
		<title>%v</title>
		<body>
		%v
		</body>
		</html>
		`, title, html), nil

}
