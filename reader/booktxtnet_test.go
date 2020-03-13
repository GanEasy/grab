package reader

import (
	"testing"
)

func Test_BooktxtnetGetInfo(t *testing.T) {
	urlStr := `https://www.booktxt.net/wapbook/4891_4943641.html`
	urlStr = `https://www.booktxt.net/0_10117/511717956.html`
	// urlStr = `https://m.35xs.com/book/237551/51896850.html`
	reader := BooktxtnetReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
func Test_BooktxtnetGetCatalog(t *testing.T) {
	urlStr := `https://www.booktxt.net/wapbook/4891.html`
	urlStr = `https://www.booktxt.net/0_10117/`
	reader := BooktxtnetReader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_BooktxtnetSearch(t *testing.T) {
	reader := BooktxtnetReader{}
	list, err := reader.Search(`圣王`)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_BooktxtnetGetList(t *testing.T) {
	urlStr := `https://www.booktxt.net/wapsort/1_1.html`
	urlStr = `https://www.booktxt.net/wapsort/5_1.html`
	reader := BooktxtnetReader{}
	list, err := reader.GetList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
