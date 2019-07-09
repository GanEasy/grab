package reader

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//LinksToCards 链接组转换成卡片组
func LinksToCards(links []Link, page, drive string) (cards []Card) {
	for _, link := range links { //所有链接
		// todo 合成链接
		wxto := fmt.Sprintf(`%v?drive=%v&url=%v`, page, drive, EncodeURL(link.URL))
		cards = append(cards, Card{link.Title, wxto, ``, `link`, ``, nil, link.URL})
	}
	return
}

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
		// `-`,
		`.`,
	}

	link, _ := url.Parse(urlStr)

	// link.Path =
	for _, t := range exp {
		// u := fmt.Sprintf(`%v`, link.Path)
		link.Path = strings.Replace(link.Path, t, ",", -1)
	}

	// 补丁，替换掉 .htm .html .shtml 静态后缀
	link.Path = strings.Replace(link.Path, ".htm", ",htm", -1)
	link.Path = strings.Replace(link.Path, ".sht", ",sht", -1)

	return fmt.Sprintf(`%v,%v%v`, link.Scheme, link.Host, link.Path)
}

//GetHash 获取字符串hash
func GetHash(s string) string {
	h := sha1.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

//GetCatalogHash get 获取目录hash
func GetCatalogHash(list Catalog) string {
	var buf bytes.Buffer
	buf.WriteString(list.Title)
	for _, card := range list.Cards {
		buf.WriteString(card.Title)
		buf.WriteString(card.WxTo)
	}
	buf.WriteString(list.SourceURL)
	return GetHash(buf.String())
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
		if strings.Index(u, "java") != 0 && n != "" {
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
				"",
			})
		}
	})
	return links
}

//EncodeURL 把url encode
func EncodeURL(str string) string {
	// es := base64.URLEncoding.EncodeToString([]byte(str))
	return url.QueryEscape(base64.URLEncoding.EncodeToString([]byte(str)))
	// return encodeURIComponent(es)
}

//DecodeURL 把url decode
func DecodeURL(str string) (string, error) {
	es, err := url.QueryUnescape(str)
	strByte, err := base64.URLEncoding.DecodeString(es)
	return string(strByte), err
}

// 可以通过修改底层url.QueryEscape代码获得更高的效率，很简单
func encodeURIComponent(str string) string {
	r := url.QueryEscape(str)
	r = strings.Replace(r, "+", "%20", -1)
	return r
}

//SimilarText 函数来自 http://www.syyong.com/Go/Go-implements-the-string-similarity-calculation-function-Levenshtein-and-SimilarText.html
// similar_text()
func SimilarText(first, second string, percent *float64) int {
	var similarText func(string, string, int, int) int
	similarText = func(str1, str2 string, len1, len2 int) int {
		var sum, max int
		pos1, pos2 := 0, 0

		// Find the longest segment of the same section in two strings
		for i := 0; i < len1; i++ {
			for j := 0; j < len2; j++ {
				for l := 0; (i+l < len1) && (j+l < len2) && (str1[i+l] == str2[j+l]); l++ {
					if l+1 > max {
						max = l + 1
						pos1 = i
						pos2 = j
					}
				}
			}
		}

		if sum = max; sum > 0 {
			if pos1 > 0 && pos2 > 0 {
				sum += similarText(str1, str2, pos1, pos2)
			}
			if (pos1+max < len1) && (pos2+max < len2) {
				s1 := []byte(str1)
				s2 := []byte(str2)
				sum += similarText(string(s1[pos1+max:]), string(s2[pos2+max:]), len1-pos1-max, len2-pos2-max)
			}
		}

		return sum
	}

	l1, l2 := len(first), len(second)
	if l1+l2 == 0 {
		return 0
	}
	sim := similarText(first, second, l1, l2)
	if percent != nil {
		*percent = float64(sim*200) / float64(l1+l2)
	}
	return sim
}
