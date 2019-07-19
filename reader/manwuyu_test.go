package reader

import (
	"testing"
)

func Test_ManwuyuReaderGetInfo(t *testing.T) {
	urlStr := `http://www.manwuyu.com/15396.html`
	urlStr = `http://www.manwuyu.com/11132.html`
	// urlStr = `http://www.manwuyu.com/2838.html`
	// urlStr = `https://m.35xs.com/book/237551/51896850.html`
	reader := ManwuyuReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_ManwuyuReaderGetCatalog(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `http://www.manwuyu.com/tag/%E4%B8%A7%E5%B0%B8%E9%81%BF%E9%9A%BE%E6%89%80`
	reader := ManwuyuReader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_ManwuyuReaderGetList(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `http://www.manwuyu.com/`
	reader := ManwuyuReader{}
	list, err := reader.GetList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
