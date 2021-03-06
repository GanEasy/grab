package reader

import (
	"math"
	"regexp"
	"testing"
)

func Test_MathSqrt(t *testing.T) {
	t.Fatal(math.Sqrt(10))
}

func Test_UxiaoshuoGetBooks(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	urlStr := "https://m.uxiaoshuo.com/sort/3/1.html"
	urlStr = "https://m.uxiaoshuo.com/sort/1/1.html"
	reader := UxiaoshuoReader{}
	list, err := reader.GetList(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}
func Test_UxiaoshuoGetChapters(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	// http://book.zongheng.com/chapter/777234/43415281.html
	urlStr := "https://m.uxiaoshuo.com/282/282134/all.html"
	urlStr = "https://m.uxiaoshuo.com/281/281973/all.html"
	reader := UxiaoshuoReader{}
	list, err := reader.GetCatalog(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_UxiaoshuoGetChapter(t *testing.T) {
	urlStr := "https://m.uxiaoshuo.com/282/282134/1460954.html"
	urlStr = "https://m.uxiaoshuo.com/140/140420/7432883.html"
	urlStr = "https://m.uxiaoshuo.com/278/278598/1741780.html"
	urlStr = "https://m.uxiaoshuo.com/281/281973/1795867.html"
	urlStr = "https://m.uxiaoshuo.com/281/281973/1796673.html"
	urlStr = "https://m.uxiaoshuo.com/281/281973/1798980.html"
	reader := UxiaoshuoReader{}
	list, err := reader.GetInfo(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_UxiaoshuoMustCompile(t *testing.T) {
	str := `aa<a href="javascript:postErrorChapter(1429556,282134);" style="text-align:center;color:red;">『章节错误,点此举报』</a>ss`

	reg := regexp.MustCompile(`<a(.+)>([^<]+)<\/a>`)

	str = reg.ReplaceAllString(str, "")

	t.Fatal(str)
}
