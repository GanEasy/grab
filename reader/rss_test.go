package reader

import (
	"crypto/sha1"
	"fmt"
	"io"
	"testing"
)

func Test_GetRssListReader(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	// urlStr := "http://news.qq.com/newsgn/rss_newsgn.xml"
	urlStr := "https://rsshub.app/mzitu/home"
	// urlStr := "https://rsshub.app/douyin/user/93610979153"
	reader := RssReader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}
func Test_Hash(t *testing.T) {
	s := "sha1"

	h := sha1.New()

	// h.Write([]byte(s))

	// bs := h.Sum(nil)
	// t.Fatal(byteString(bs))

	io.WriteString(h, s)
	a := fmt.Sprintf("%x", h.Sum(nil))
	t.Fatal(a)
}
