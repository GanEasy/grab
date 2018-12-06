package grab

import (
	"testing"
)

func Test_GetArticleList(t *testing.T) {
	urlStr := `http://www.longfu8.com/417.html`
	urlStr = `http://www.xinshubao.net/18/18685/2515188_2.html`
	urlStr = `http://news.qq.com/`
	urlStr = `https://wechatrank.com/`
	info, err := GetArticleList(urlStr)
	if err != nil {

	}
	t.Fatal(info)

}

func Test_GetArticleListReader(t *testing.T) {
	urlStr := `http://www.longfu8.com/417.html`
	urlStr = `http://www.xinshubao.net/18/18685/2515188_2.html`
	// urlStr = `http://news.qq.com/`
	// urlStr = `https://wechatrank.com/`
	reader := ArticleListReader{}
	list, err := reader.GetList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_GetArticleInfoReader(t *testing.T) {
	urlStr := `http://www.longfu8.com/417.html`
	urlStr = `http://www.xinshubao.net/18/18685/2515188_2.html`
	urlStr = `https://mp.weixin.qq.com/s?__biz=MzU0NzY5NjQyNQ==&mid=2247489969&idx=4&sn=c28f89d891ac1895af4ddbe9954b0b43&chksm=fb4b23f7cc3caae175a87d9e0297744658dabfe3a6367cfee6892e4b58e3186b709f395f9636&scene=27#wechat_redirect`

	// urlStr = `http://news.qq.com/`
	// urlStr = `https://wechatrank.com/`
	reader := ArticleInfoReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_GetArticle(t *testing.T) {
	urlStr := `http://www.longfu8.com/417.html`
	urlStr = `https://mp.weixin.qq.com/s?__biz=MzU0NzY5NjQyNQ==&mid=2247489969&idx=4&sn=c28f89d891ac1895af4ddbe9954b0b43&chksm=fb4b23f7cc3caae175a87d9e0297744658dabfe3a6367cfee6892e4b58e3186b709f395f9636&scene=27#wechat_redirect`
	info, err := GetArticle(urlStr)
	if err != nil {

	}
	t.Fatal(info)

}
