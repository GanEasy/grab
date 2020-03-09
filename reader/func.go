package reader

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"strings"

	"github.com/GanEasy/html2article"
	"github.com/PuerkitoBio/goquery"
	"github.com/lunny/html2md"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

//ImagesBuildHTML 图片组转html
func ImagesBuildHTML(images []string) string {

	var buffer bytes.Buffer
	for _, v := range images {

		buffer.WriteString(`<img src="`)
		buffer.WriteString(v)
		buffer.WriteString(`">`)
	}

	return buffer.String()
}

// 函数包

//GetLinkByHTML 获取网页内容所有链接
func GetLinkByHTML(urlStr, html string) (links []Link, err error) {
	// 没有 html标签 或者 body 标签可能出现文档解释异常
	if !strings.Contains(html, `</html>`) || !strings.Contains(html, `</body>`) {
		html = fmt.Sprintf(`
			<html>
			<head>
			<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
			<title>%v</title>
			<body>
			%v
			</body>
			</html>
			`, `NONE TITLE`, html)
	}

	c := strings.NewReader(html)

	g, err := goquery.NewDocumentFromReader(c)

	if err != nil {
		return
	}
	link, err := url.Parse(urlStr)
	if err != nil {
		return
	}
	links = GetLinks(g, link)
	return
}

//CheckStrIsLink 检查字符串是否支持的链接
func CheckStrIsLink(urlStr string) error {

	link, err := url.Parse(urlStr)

	if err != nil {
		return err
	}

	if link.Scheme == "" {
		return errors.New("Scheme Fatal")
	}

	if link.Host == "" {
		return errors.New("Host Fatal")
	}
	return nil
}

//GetPreviousLink 获取上一页或者上一章
func GetPreviousLink(links []Link) Link {
	for _, link := range links {
		if strings.Contains(link.Title, `上一页`) || strings.Contains(link.Title, `上一章`) || strings.Contains(link.Title, `上一`) {
			return Link{Title: "previous", URL: link.URL}
		}
	}
	return Link{}
}

//GetNextLink 获取下一页或者下一章
func GetNextLink(links []Link) Link {
	for _, link := range links {
		if strings.Contains(link.Title, `下一页`) || strings.Contains(link.Title, `下页`) || strings.Contains(link.Title, `下一`) || strings.Contains(link.Title, `下一章`) || strings.Contains(link.Title, `下章`) || link.Title == `>` {
			return Link{Title: "next", URL: link.URL}
		}
	}
	return Link{}
}

//GetActicleByHTML 由Html返回*html2article.Article
func GetActicleByHTML(html string) (article *html2article.Article, err error) {
	ext, err := html2article.NewFromHtml(html)
	if err != nil {
		return
	}

	return ext.ToArticle()
}

//GetActicleForHTML 由Html返回 *html2article.Article (纯原文)
func GetActicleForHTML(html string) (article *html2article.Article, err error) {
	ext, err := html2article.NewFromHtml(html)
	if err != nil {
		return
	}
	return ext.HTMLToArticle()
}

//GetActicleByContent 由Html返回*html2article.Article
func GetActicleByContent(html string) (article *html2article.Article, err error) {

	if !strings.Contains(html, `</html>`) || !strings.Contains(html, `</body>`) {
		html = fmt.Sprintf(`
			<html>
			<head>
			<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
			<title>%v</title>
			<body>
			<div>
			%v
			</div>
			</body>
			</html>
			`, `NONE TITLE`, html)
	}

	return GetActicleByHTML(html)
}

// GetSectionByContent 通过正文获取段落拆分
func GetSectionByContent(content string) (sec []string) {
	// 替换换行符
	content = BookContReplace(content)
	// 拆分换行符
	arr := strings.Split(content, "</p>")
	if len(arr) > 1 {
		for _, v := range arr {
			text := strings.TrimSpace(v)
			if text != "" {
				// 不为空时组装段落
				sec = append(sec, text)
			}
		}
	}
	return
}

//MarkDownFormatContent 通过markdown语法格式化内容
func MarkDownFormatContent(content string) string {
	md := html2md.Convert(content)
	input := []byte(md)
	unsafe := blackfriday.MarkdownCommon(input)
	contentBytes := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return strings.TrimSpace(fmt.Sprintf(`%v`, string(contentBytes[:])))
}

// BookContReplace 小说内容正文替换标签
func BookContReplace(html string) string {
	c := strings.Replace(html, `<p>`, ``, -1)
	c = strings.Replace(c, `<code>`, ``, -1)
	c = strings.Replace(c, `</code>`, ``, -1)
	c = strings.Replace(c, `<pre>`, ``, -1)
	c = strings.Replace(c, `</pre>`, ``, -1)

	c = strings.Replace(c, `<br/>`, `</p>`, -1)
	c = strings.Replace(c, `<br />`, `</p>`, -1)
	c = strings.Replace(c, `<br>`, `</p>`, -1)
	c = strings.Replace(c, "\n", `</p>`, -1)
	return c
}

/*******https://github.com/astaxie/beego/blob/master/utils/slice.go start*********/

// InSliceIface checks given interface in interface slice.
func InSliceIface(v interface{}, sl []interface{}) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

// SliceUnique cleans repeated values in slice. 去重
func SliceUnique(slice []interface{}) (uniqueslice []interface{}) {
	for _, v := range slice {
		if !InSliceIface(v, uniqueslice) {
			uniqueslice = append(uniqueslice, v)
		}
	}
	return
}

// SliceShuffle shuffles a slice. 重组
func SliceShuffle(slice []interface{}) []interface{} {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
	return slice
}

/*******https://github.com/astaxie/beego/blob/master/utils/slice.go end*********/

// InArr 字符串被包含
func InArr(str string, arr []string) bool {
	for _, vv := range arr {
		if vv == str {
			return true
		}
	}
	return false
}

//Intersection 字符串交集
func Intersection(arr1 []string, arr2 []string) []string {
	for _, v := range arr1 {
		if !InArr(v, arr2) {
			arr2 = append(arr2, v)
		}
	}
	return arr2
}

//Union 字符串并集
func Union(arr1 []string, arr2 []string) (ret []string) {
	for _, v := range arr1 {
		if InArr(v, arr2) {
			ret = append(ret, v)
		}
	}
	return ret
}

//Nonion 字符串非集
func Nonion(arr1 []string, arr2 []string) (ret []string) {
	arr3 := Intersection(arr1, arr2)
	arr4 := Union(arr1, arr2)
	for _, v := range arr3 {
		if !InArr(v, arr4) {
			ret = append(ret, v)
		}
	}
	return ret
}

//Jaccard  杰卡德（Jaccard）相似系数
func Jaccard(arr1 []string, arr2 []string) float64 {
	return float64(len(Union(arr1, arr2))) / float64(len(Intersection(arr1, arr2)))
}

//JaccardMateGetURL  杰卡德（Jaccard）相似系数 匹配出目标url
/**
快速获取固定结构动态链接
url 为要验证参考的链接地址
demo1 有效的学习链接地址1，与url具有相同结构
demo2 有效的学习链接地址2，与url具有相同结构
to1   有效果的目标链接地址1， 将url变量替换到to1结构中 (to1为空时，保持原有结构)

todo 	t.Fatal(JaccardMateGetURL(`http://book.zongheng.com/book/book/658887.html`, `http://book.zongheng.com/book/769150.html`, `http://book.zongheng.com/book/316562.html`, `http://book.zongheng.com/showchapter/769150.html`))
相同值不同位置时抽风了。。
*/
func JaccardMateGetURL(url, demo1, demo2, to1 string) (string, bool) {
	// demo1,2的标签
	demotag1 := strings.Split(GetTag(demo1), ",")
	demotag2 := strings.Split(GetTag(demo2), ",")

	// url的标签
	urltag := strings.Split(GetTag(url), ",")

	if len(demotag1) != len(urltag) {
		return url, false
	}

	eduUn := Union(demotag1, demotag2)
	eduIn := Intersection(demotag1, demotag2)

	demoParamUn := Union(demotag1, Nonion(eduIn, eduUn))

	paramRp := Nonion(urltag, eduUn)

	if len(demoParamUn) != len(paramRp) {
		return url, false
	}
	// log.Println(demotag1, demotag2, urltag, demoParamUn, paramRp)
	if to1 != `` {
		to := to1
		for i, val := range demoParamUn {
			// t.Fatal(i, val)
			// log.Println(to, val, paramRp[i])
			to = strings.Replace(to, val, paramRp[i], 1)
		}
		return to, true
	}
	return url, true

}

// GetPathLevel 获取页面地址的等级
func GetPathLevel(wxto string) (level int32) {

	if b := strings.Contains(wxto, string("drive=qidian")); b == true {
		return 2
	}
	if b := strings.Contains(wxto, string("drive=zongheng")); b == true {
		return 2
	}
	if b := strings.Contains(wxto, string("drive=17k")); b == true {
		return 2
	}
	if b := strings.Contains(wxto, string("drive=qu")); b == true {
		return 9 //换域名jx了
	}
	if b := strings.Contains(wxto, string("drive=jx")); b == true {
		return 2
	}
	if b := strings.Contains(wxto, string("drive=mcmssc")); b == true {
		return 2
	}
	if b := strings.Contains(wxto, string("drive=xbiquge")); b == true {
		return 2
	}
	if b := strings.Contains(wxto, string("drive=luoqiu")); b == true {
		return 2
	}
	if b := strings.Contains(wxto, string("drive=booktxt")); b == true {
		return 2 //
	}
	if b := strings.Contains(wxto, string("drive=bxwx")); b == true {
		return 9 //网站打不开了
	}
	if b := strings.Contains(wxto, string("drive=uxiaoshuo")); b == true {
		return 2
	}
	if b := strings.Contains(wxto, string("drive=biquyun")); b == true {
		return 9 //网站打不开了
	}
	if b := strings.Contains(wxto, string("drive=soe8")); b == true {
		return 2
	}
	if b := strings.Contains(wxto, string("drive=soe8")); b == true {
		return 2
	}
	if b := strings.Contains(wxto, string("drive=xs280")); b == true {
		return 2
	}

	if b := strings.Contains(wxto, string("drive=xxsy")); b == true {
		return 2
	}
	if b := strings.Contains(wxto, string("drive=hongxiu")); b == true {
		return 2
	}
	if b := strings.Contains(wxto, string("drive=laosijixs")); b == true {
		return 2
	}
	if b := strings.Contains(wxto, string("drive=text")); b == true {
		return 2
	}

	if b := strings.Contains(wxto, string("drive=article")); b == true {
		return 2
	}

	if b := strings.Contains(wxto, string("drive=biqugeinfo")); b == true {
		return 3
	}
	if b := strings.Contains(wxto, string("drive=haimaoba")); b == true {
		return 3
	}
	if b := strings.Contains(wxto, string("drive=hanmanwo")); b == true {
		return 3
	}
	if b := strings.Contains(wxto, string("drive=hanmanku")); b == true {
		return 3
	}
	if b := strings.Contains(wxto, string("drive=manhwa")); b == true {
		return 3
	}
	if b := strings.Contains(wxto, string("drive=ssmh")); b == true {
		return 3
	}
	if b := strings.Contains(wxto, string("drive=fuman")); b == true {
		return 3
	}
	if b := strings.Contains(wxto, string("drive=aimeizi5")); b == true {
		return 3
	}
	if b := strings.Contains(wxto, string("drive=kanmeizi")); b == true {
		return 3
	}
	if b := strings.Contains(wxto, string("drive=manwuyu")); b == true {
		return 3
	}
	return 1
}

// GetURLDrive 获取链接地址的解释器
func GetURLDrive(url string) (drive, key string) {

	return ``, ``
}
