package reader

import (
	"testing"
)

func Test_WeijiaoshouReaderGetInfo(t *testing.T) {
	urlStr := `http://www.weijiaoshou.cn/manhua/bobobo/1916.html`
	// urlStr = `http://www.xinshubao.net/18/18685/2515188_2.html`
	// urlStr = `https://m.35xs.com/book/237551/51896850.html`
	reader := WeijiaoshouReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_WeijiaoshouReaderGetCatalog(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `http://www.weijiaoshou.cn/manhua/shougandieji.html`
	reader := WeijiaoshouReader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_WeijiaoshouReaderGetList(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapbook/4891.html`
	urlStr = `http://www.weijiaoshou.cn/manhua/liebiao/hanguomanhua-2.html`
	reader := WeijiaoshouReader{}
	list, err := reader.GetList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
