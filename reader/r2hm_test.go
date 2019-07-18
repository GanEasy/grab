package reader

import (
	"testing"
)

func Test_R2hmReaderGetInfo(t *testing.T) {
	urlStr := `https://r2hm.com/chapter/25486`
	// urlStr = `http://www.xinshubao.net/18/18685/2515188_2.html`
	// urlStr = `https://m.35xs.com/book/237551/51896850.html`
	reader := R2hmReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_R2hmReaderGetCatalog(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `https://r2hm.com/book/459`
	reader := R2hmReader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
