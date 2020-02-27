package reader

import (
	"regexp"
	"testing"
)

func Test_McmsscSearch(t *testing.T) {
	reader := McmsscReader{}
	list, err := reader.Search(`点道`)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_McmsscGetBooks(t *testing.T) {
	urlStr := "https://www.mcmssc.com/xuanhuanxiaoshuo/"
	urlStr = "https://www.mcmssc.com/chuanyuexiaoshuo/"
	urlStr = "https://www.mcmssc.com/paihangbang/"
	reader := McmsscReader{}
	list, err := reader.GetList(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_McmsscGetChapters(t *testing.T) {
	urlStr := "http://www.xbiquge.la/39/39720/"
	urlStr = "https://www.mcmssc.com/44_44569/"
	// urlStr = "http://www.xbiquge.la/15/15484/"
	reader := McmsscReader{}
	list, err := reader.GetCatalog(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_McmsscGetChapter(t *testing.T) {
	urlStr := "http://www.xbiquge.la/29/29211/18370058.html"
	urlStr = "https://www.mcmssc.com/44_44569/21647159.html"
	reader := McmsscReader{}
	list, err := reader.GetInfo(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_McmsscMustCompile(t *testing.T) {
	str := `window.location.href = "/chuanyuexiaoshuo/4_" + obj.curr + ".html";`

	reg := regexp.MustCompile(`window.location.href = "(?P<pagemod>[^"]+)"`)

	str = reg.ReplaceAllString(str, "")

	t.Fatal(str)
}
