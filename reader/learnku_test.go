package reader

import (
	"testing"
)

func Test_GetLearnkuInfoReader(t *testing.T) {
	urlStr := `https://learnku.com/docs/laravel-specification/5.5/data-filling/516`
	reader := LearnkuReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
func Test_LearnkuGetCatalog(t *testing.T) {
	urlStr := `http://www.longfu8.com/`
	urlStr = `https://learnku.com/docs/laravel-specification/5.5`
	reader := LearnkuReader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
