package grab

import (
	"testing"
)

func Test_ExplainLinkQidian(t *testing.T) {
	urlStr := `http://m.qidian.com/book/1004608738/catalog`
	address, drive, page := ExplainLink(urlStr)
	t.Fatal(address, drive, page)
}

func Test_ExplainLinkBiqugeinfo(t *testing.T) {
	urlStr := `https://m.biquge.info/list/5_7.html`
	// urlStr = `https://m.biquge.info/51_51498/`
	address, drive, page := ExplainLink(urlStr)
	t.Fatal(address, drive, page)
}

func Test_ExplainLinkBiquyun(t *testing.T) {
	urlStr := `https://m.biquyun.com/16_16635/10124285.html`
	// urlStr = `https://m.biquyun.com/16_16635_1_1.html`
	// urlStr = `https://m.biquyun.com/16_16635/`
	urlStr = `https://m.biquyun.com/`
	address, drive, page := ExplainLink(urlStr)
	t.Fatal(address, drive, page)
}
