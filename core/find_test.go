package reader

import (
	"testing"
)

func Test_GetFindArticle(t *testing.T) {

	url := "http://book.zongheng.com/showchapter/730066.html"

	// url := "http://www.longfu8.com/417.html"
	html, _ := GetFindArticle(url, `.read_con`)
	// t.Fatal(article.Content)
	t.Fatal(html)

}
