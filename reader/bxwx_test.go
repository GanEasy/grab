package reader

import (
	"testing"
)

func Test_BxwxJaccardMateGetURL(t *testing.T) {
	urlStr := `https://m.bxwx.la/b/2/2218/`
	urlStr2, state := JaccardMateGetURL(urlStr, `https://m.bxwx.la/b/246/246596/`, `https://m.bxwx.la/b/287/287378/`, `https://m.bxwx.la/binfo/246/246596.htm`)

	t.Fatal(urlStr2, state)
}

func Test_BxwxGetInfo(t *testing.T) {
	urlStr := `https://m.bxwx.la/b/316/316850/1684236.html`
	// urlStr = `http://www.xinshubao.net/18/18685/2515188_2.html`
	// urlStr = `https://m.35xs.com/book/237551/51896850.html`
	reader := BxwxReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
func Test_BxwxGetCatalog(t *testing.T) {
	urlStr := `https://m.bxwx.la/binfo/316/316850.htm`
	urlStr = `https://m.bxwx.la/binfo/246/246596.htm`
	reader := BxwxReader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_BxwxGetList(t *testing.T) {
	urlStr := `https://m.bxwx.la/binfo/316/316850.htm`
	urlStr = `https://m.bxwx.la/bsort1/0/1.htm`
	urlStr = `https://m.bxwx.la/bsort1/0/8.htm`
	reader := BxwxReader{}
	list, err := reader.GetList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
