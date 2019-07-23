package reader

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func Test_GetHTML(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	urlStr := "http://feed.williamlong.info/"
	list, err := GetHTML(urlStr, ``)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_GetHTMLOrCache(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	urlStr := "http://feed.williamlong.info/"
	list, err := GetHTMLOrCache(urlStr, ``)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_GetURLHost(t *testing.T) {
	urlStr := `http://feeds.twit.tv/twit.xml`
	urlStr = "https://m.uxiaoshuo.com/281/281973/1798980.html"
	link, err := url.Parse(urlStr)

	if err != nil {
	}

	if link.Scheme == "" {
	}

	if link.Host == "" {
	}

	t.Fatal(link.Host)
}
func Test_ReplaceImageServe(t *testing.T) {

	body := `<img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c07.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c01.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c02.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c03.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c04.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c05.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c06.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c07.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c08.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c09.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c10.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c11.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c12.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c13.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c14.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c15.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c16.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c17.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c18.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c19.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c20.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c21.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c22.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c23.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c24.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c25.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c26.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c27.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c28.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c29.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c30.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c31.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c32.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c33.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c34.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c35.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c36.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c37.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c38.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c39.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c40.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c41.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c42.jpg"><br />`
	// log.Println(html)

	list, err := ReplaceImageServe(body)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_ReplaceImageServeDev(t *testing.T) {
	// list, err := ReplaceImageServe(``)
	body := `<img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c07.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c01.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c02.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c03.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c04.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c05.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c06.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c07.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c08.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c09.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c10.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c11.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c12.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c13.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c14.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c15.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c16.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c17.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c18.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c19.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c20.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c21.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c22.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c23.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c24.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c25.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c26.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c27.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c28.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c29.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c30.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c31.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c32.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c33.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c34.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c35.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c36.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c37.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c38.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c39.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c40.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c41.jpg"><br /><img referrerpolicy="no-referrer" src="https://i2.meizitu.net/2018/12/27c42.jpg"><br />`
	// log.Println(html)
	article, err := GetActicleByContent(body)
	if err != nil {

		t.Fatal(err)
	}
	for _, i := range article.Images {
		body = strings.Replace(body, i, fmt.Sprintf(`https://img.readfollow.com/file?url=%v`, i), -1)
	}
	t.Fatal(body)
}
func Test_SimilarText(t *testing.T) {
	var percent float64

	// tSimilarText := SimilarText("golang", "google", &percent)
	tSimilarText := SimilarText(`https://blog.csdn.net/shijing_0214/article/details/53100992`, "https://blog.csdn.net/u010480899", &percent)
	t.Fatal(tSimilarText, percent)
	equal(t, 3, tSimilarText)
	equal(t, float64(50), percent)
}

// Expected to be equal.
func equal(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual))
	}
}

// Expected to be unequal.
func unequal(t *testing.T, expected, actual interface{}) {
	if reflect.DeepEqual(expected, actual) {
		t.Errorf("Did not expect %v (type %v) - Got %v (type %v)", expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual))
	}
}

// Expect a greater than b.
func gt(t *testing.T, a, b float64) {
	if a <= b {
		t.Errorf("Expected %v (type %v) > Got %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
}

// Expect a greater than or equal to b.
func gte(t *testing.T, a, b float64) {
	if a < b {
		t.Errorf("Expected %v (type %v) > Got %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
}

// Expected value needs to be within range.
func rangeValue(t *testing.T, min, max, actual float64) {
	if actual < min || actual > max {
		t.Errorf("Expected range of %v-%v (type %v) > Got %v (type %v)", min, max, reflect.TypeOf(min), actual, reflect.TypeOf(actual))
	}
}
