package reader

import (
	"testing"
)

func Test_Paoshu8GetInfo(t *testing.T) {
	urlStr := `http://m.paoshu8.com/wapbook-1011-783829/`
	// urlStr = `http://www.xinshubao.com/18/18685/2515188_2.html`
	urlStr = `http://www.paoshu8.com/132_132325/170731522.html`
	reader := Paoshu8Reader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
func Test_Paoshu8GetCatalog(t *testing.T) {
	urlStr := `https://m.Paoshu8.com/wapbook/4891.html`
	urlStr = `http://www.paoshu8.com/0_132325/`
	reader := Paoshu8Reader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_Paoshu8Search(t *testing.T) {
	reader := Paoshu8Reader{}
	list, err := reader.Search(`圣王`)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_Paoshu8GetList(t *testing.T) {
	urlStr := `http://m.paoshu8.com/sort-2-1/`
	urlStr = `http://m.paoshu8.com/sort-1-1/`
	reader := Paoshu8Reader{}
	list, err := reader.GetList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
