package reader

import (
	"testing"
)

func Test_HanmankuReaderGetInfo(t *testing.T) {
	urlStr := `https://r2hm.com/chapter/25486`
	urlStr = `http://www.hanmanku.com/chapter/24116`
	// urlStr = `https://m.35xs.com/book/237551/51896850.html`
	reader := HanmankuReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_HanmankuReaderGetCatalog(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `http://www.hanmanku.com/book/316`
	reader := HanmankuReader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_HanmankuReaderGetList(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `http://www.hanmanku.com/booklist`
	reader := HanmankuReader{}
	list, err := reader.GetList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
