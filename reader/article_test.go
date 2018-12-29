package reader

import (
	"testing"
)

func Test_GetArticleInfo(t *testing.T) {
	// urlStr := `http://www.zjl88.com/article-show-id-709430.html`
	urlStr := `https://wechatrank.com/`

	reader := ArticleReader{}
	ret, _ := reader.GetInfo(urlStr)
	t.Fatal(ret)

}
