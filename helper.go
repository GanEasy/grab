package grab

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//Cleaning 清洗数据
func Cleaning(links []Link) (newlinks []Link) {
	// 拆分链接字符占比重
	var edu = map[string]int{}
	for _, link := range links { //所有链接
		s := GetTag(link.URL)
		for _, k := range strings.Split(s, ",") { //链接分拆统计
			if v, ok := edu[k]; ok && k != "" && k != " " {
				v++
				edu[k] = v
			} else {
				edu[k] = 1
			}
		}
	}

	var mw = 0
	var maxWeight = 0.0

	// log.Fatal(edu)
	for _, v := range edu {
		if v > 10 {
			v = 10
		}
		mw += v
	}

	// 找出最大重量
	for _, link := range links {
		s := GetTag(link.URL)
		w := 0
		for _, k := range strings.Split(s, ",") {
			if v, ok := edu[k]; ok {
				w += v
			}
		}
		if (float64(w) / float64(mw)) > maxWeight {
			maxWeight = float64(w) / float64(mw)
		}
		// wg[link.URL] = w
	}
	var pro = maxWeight * 0.30
	// 这个链接的重量
	var wg = map[string]int{}
	for _, link := range links {
		s := GetTag(link.URL)
		w := 0
		for _, k := range strings.Split(s, ",") {
			if v, ok := edu[k]; ok {
				w += v
			}
		}
		if float64(w) > (float64(mw) * float64(pro)) {
			wg[link.URL] = w
		}
		// wg[link.URL] = w
	}

	// log.Fatal(links)
	var crp = map[string]int{}
	for _, link := range links {
		if _, ok := crp[link.URL]; !ok && link.Title != "" {
			crp[link.URL] = 1
			if _, ok := wg[link.URL]; ok && link.Title != "" {
				newlinks = append(newlinks, link)
			}
		}

	}
	return newlinks
}

//GetTag 获取特点
func GetTag(urlStr string) string {

	var exp = []string{
		`?`,
		`&`,
		`#`,
		`/`,
		`=`,
		// `.`, // 不能把这个点去掉
	}

	link, _ := url.Parse(urlStr)

	// link.Path =
	for _, t := range exp {
		// u := fmt.Sprintf(`%v`, link.Path)
		link.Path = strings.Replace(link.Path, t, ",", -1)
	}
	return link.Path
}

//GetHash 获取hash
func GetHash(s string) string {
	h := sha1.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

//GetListHash get
func GetListHash(list List) string {
	var buf bytes.Buffer
	buf.WriteString(list.Title)
	for _, link := range list.Links {
		buf.WriteString(link.Title)
		buf.WriteString(link.URL)
	}
	buf.WriteString(list.SourceURL)
	return GetHash(buf.String())
}

//GetLinks 获取链接地址 参考链接
func GetLinks(g *goquery.Document, link *url.URL) (links []Link) {
	g.Find("a").Each(func(i int, content *goquery.Selection) {
		n := strings.TrimSpace(content.Text())
		u, _ := content.Attr("href")
		if strings.Index(u, "java") != 0 {
			if strings.Index(u, "//") == 0 {
				u = fmt.Sprintf(`%v:%v`, link.Scheme, u)
			} else if strings.Index(u, "/") == 0 {
				u = fmt.Sprintf(`%v://%v%v`, link.Scheme, link.Host, u)
			} else if strings.Index(u, "#") != 0 && strings.Index(u, "http") != 0 {
				//todo   link.Path 获取目录
				p1, _ := filepath.Split(link.Path)
				u = fmt.Sprintf(`%v://%v%v%v`, link.Scheme, link.Host, p1, u)
			}
			u = strings.Replace(u, " ", "", -1)
			u = strings.Replace(u, "　", "", -1)
			// 去除换行符
			u = strings.Replace(u, "\n", "", -1)
			u = strings.Replace(u, "\t", "", -1)

			links = append(links, Link{
				n,
				u,
			})
		}
	})
	return links
}
