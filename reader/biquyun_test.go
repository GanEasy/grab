package reader

import (
	"regexp"
	"testing"
)

func Test_BiquyunGetBooks(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	urlStr := "https://m.uxiaoshuo.com/sort/3/1.html"
	urlStr = "https://m.biquyun.com/wapsort/1_1.html"
	reader := BiquyunReader{}
	list, err := reader.GetList(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}
func Test_BiquyunGetChapters(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	// http://book.zongheng.com/chapter/777234/43415281.html
	urlStr := "https://m.uxiaoshuo.com/282/282134/all.html"
	urlStr = "https://m.biquyun.com/20_20760_1_1.html"
	reader := BiquyunReader{}
	list, err := reader.GetCatalog(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_BiquyunGetChapter(t *testing.T) {
	urlStr := "https://m.uxiaoshuo.com/282/282134/1460954.html"
	urlStr = "https://m.biquyun.com/20_20760/9866494.html"
	reader := BiquyunReader{}
	list, err := reader.GetInfo(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_BiquyunMustCompile(t *testing.T) {
	str := `aa<a href="javascript:postErrorChapter(1429556,282134);" style="text-align:center;color:red;">『章节错误,点此举报』</a>ss`

	reg := regexp.MustCompile(`<a(.+)>([^<]+)<\/a>`)

	str = reg.ReplaceAllString(str, "")

	t.Fatal(str)
}
