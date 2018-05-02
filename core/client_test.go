package core

import (
	"testing"
)

func Test_GetHTML(t *testing.T) {

	url := "http://book.zongheng.com/showchapter/730066.html"

	// url := "http://www.longfu8.com/417.html"
	html, _ := GetHTML(url)
	// t.Fatal(article.Content)
	t.Fatal(html)

}
