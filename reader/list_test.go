package reader

import (
	"path/filepath"
	"testing"
)

func Test_GetHTMLList(t *testing.T) {
	// // urlStr := "http://feeds.twit.tv/twit.xml"
	// // urlStr := "http://www.williamlong.info/"
	// // list, err := GetHTMLList(urlStr, `#divMain`)
	// urlStr := `http://www.xinshubao.net/18/18685/2515188.html`
	// list, err := GetHTMLList(urlStr, ``)
	// if err != nil {

	// }
	// l := Cleaning(list.Links)
	// t.Fatal(l)
}

func Test_GetHTMLListSplit(t *testing.T) {

	p1, _ := filepath.Split(`/link/a/123.html`)
	t.Fatal(p1)
}
