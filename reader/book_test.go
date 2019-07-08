package reader

import (
	"testing"
)

func Test_GetBookInfoReader(t *testing.T) {
	urlStr := `http://www.longfu8.com/417.html`
	// urlStr = `http://www.xinshubao.net/18/18685/2515188_2.html`
	// urlStr = `https://m.35xs.com/book/237551/51896850.html`
	reader := BookReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
func Test_BookGetCatalog(t *testing.T) {
	urlStr := `http://www.longfu8.com/`
	urlStr = `http://www.xinshubao.net/18/18685/`
	urlStr = `https://m.uxiaoshuo.com/155/155018/all.html`
	reader := BookReader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
