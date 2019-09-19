package reader

import (
	"testing"
)

func Test_SsmhReaderGetInfo(t *testing.T) {
	urlStr := `https://r2hm.com/chapter/25486`
	urlStr = `http://ssmh.cc/chapter/231508`
	// urlStr = `https://m.35xs.com/book/237551/51896850.html`
	reader := SsmhReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_SsmhReaderGetCatalog(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `http://www.ssmh.cc/book/4729`
	reader := SsmhReader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_SsmhReaderGetList(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `http://www.ssmh.cc/booklist`
	reader := SsmhReader{}
	list, err := reader.GetList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
