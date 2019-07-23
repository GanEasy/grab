package reader

import (
	"regexp"
	"testing"
)

func Test_Soe8GetBooks(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	urlStr := "http://m.soe8.com/sort/2_1/"
	// urlStr = "http://m.soe8.com/sort/1_1/"
	reader := Soe8Reader{}
	list, err := reader.GetList(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}
func Test_Soe8GetChapters(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	// http://book.zongheng.com/chapter/777234/43415281.html
	urlStr := "https://m.uxiaoshuo.com/282/282134/all.html"
	urlStr = "http://m.soe8.com/86_86451_1/"
	reader := Soe8Reader{}
	list, err := reader.GetCatalog(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_Soe8GetChapter(t *testing.T) {
	urlStr := "http://m.soe8.com/0_2/15978818.html"
	urlStr = "http://m.soe8.com/0_313/573670.html"
	// urlStr = "http://m.soe8.com/1_1505/26862795.html"
	// urlStr = "http://m.soe8.com/86_86451/27934236.html"
	reader := Soe8Reader{}
	list, err := reader.GetInfo(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_Soe8MustCompile(t *testing.T) {
	str := `aa<a href="javascript:postErrorChapter(1429556,282134);" style="text-align:center;color:red;">『章节错误,点此举报』</a>ss`

	reg := regexp.MustCompile(`<a(.+)>([^<]+)<\/a>`)

	str = reg.ReplaceAllString(str, "")

	t.Fatal(str)
}
