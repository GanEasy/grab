package grab

import (
	"testing"
)

func Test_GetQueryInfoReader(t *testing.T) {
	urlStr := `http://www.longfu8.com/417.html`
	// urlStr = `http://www.xinshubao.net/18/18685/2515188_2.html`
	// urlStr = `https://m.35xs.com/book/237551/51896850.html`
	reader := QueryInfoReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
func Test_GetQuerykListReader(t *testing.T) {
	urlStr := `http://www.longfu8.com/`
	urlStr = `http://www.xinshubao.net/18/18685/`
	reader := QueryListReader{}
	list, err := reader.GetList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
