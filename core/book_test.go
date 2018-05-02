package core

import "testing"

func Test_GetBookInfo(t *testing.T) {

	// urlStr := `http://www.longfu8.com/417.html`
	urlStr := `https://m.35xs.com/book/237551/51896850.html`
	info, err := GetBookInfo(urlStr)
	if err != nil {

	}
	t.Fatal(info)

}
