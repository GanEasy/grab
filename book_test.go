package grab

import (
	"log"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func Test_GetBookInfo(t *testing.T) {
	urlStr := `http://www.longfu8.com/417.html`
	urlStr = `http://www.xinshubao.net/18/18685/2515188_2.html`
	urlStr = `https://m.35xs.com/book/237551/51896850.html`
	info, err := GetBookInfo(urlStr)
	if err != nil {

	}
	t.Fatal(info)

}

func Test_GetBookInfoReader(t *testing.T) {
	urlStr := `http://www.longfu8.com/417.html`
	// urlStr = `http://www.xinshubao.net/18/18685/2515188_2.html`
	// urlStr = `https://m.35xs.com/book/237551/51896850.html`
	reader := BookInfoReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
func Test_GetBookListReader(t *testing.T) {
	urlStr := `http://www.longfu8.com/`
	urlStr = `http://www.xinshubao.net/18/18685/`
	reader := BookListReader{}
	list, err := reader.GetList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_GetBookChapters(t *testing.T) {
	urlStr := `http://www.longfu8.com/`
	urlStr = `http://www.xinshubao.net/18/18685/`

	// urlStr := `http://www.longfuxiaoshuo.com/`
	info, err := GetBookChapters(urlStr)
	if err != nil {

	}
	t.Fatal(info)

}
func Test_GetBookList(t *testing.T) {

	g, err := goquery.NewDocument(`http://www.xinshubao.net/18/18685/2515188.html`)

	if err != nil {
		return
	}

	// fmt.Println(g.Text())
	g.Find("a").Each(func(i int, content *goquery.Selection) {
		n := strings.TrimSpace(content.Text())
		u, _ := content.Attr("href")
		if err := CheckStrIsLink(u); err == nil {
			log.Println(u, n)
		}
	})

	t.Fatal(g.Html())

}
