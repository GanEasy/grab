package reader

import (
	"testing"
)

func Test_GetBlogInfoReader(t *testing.T) {
	urlStr := `https://learnku.com/docs/laravel-specification/5.5/data-filling/516`
	reader := BlogReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
func Test_BlogGetCatalog(t *testing.T) {
	urlStr := `http://www.longfu8.com/`
	urlStr = `https://learnku.com/docs/laravel-specification/5.5`
	reader := BlogReader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
