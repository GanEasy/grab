package grab

import (
	"crypto/sha1"
	"fmt"
	"io"
	"testing"

	"github.com/mmcdole/gofeed"
)

func Test_GetRssListReader(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	urlStr := "https://rsshub.app/douyin/user/93610979153"
	reader := RssListReader{}
	list, err := reader.GetList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_GetRssListReader22(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	urlStr := "https://rsshub.app/douyin/user/93610979153"

	fp := gofeed.NewParser()

	// feed, err := fp.ParseString(html)
	feed, err := fp.ParseURL(urlStr)
	if err != nil {

	}
	t.Fatal(feed.Items)
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
