package reader

import (
	"testing"
)

func Test_GetGithubInfoReader(t *testing.T) {
	urlStr := `https://github.com/digoal/blog/blob/master/201212/20121218_03.md`
	reader := GithubReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
func Test_GithubGetCatalog(t *testing.T) {
	urlStr := `http://www.longfu8.com/`
	urlStr = `https://github.com/digoal/blog/blob/master/class/24.md`
	reader := GithubReader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
