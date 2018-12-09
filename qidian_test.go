package grab

import (
	"testing"
)

func Test_QidianGetBooks(t *testing.T) {
	// urlStr := "http://feeds.twit.tv/twit.xml"
	// urlStr := "http://feed.williamlong.info/"
	urlStr := "https://www.qidian.com/all?orderId=&page=1&style=1&pageSize=20&siteid=1&pubflag=0&hiddenField=0"
	reader := QidianReader{}
	list, err := reader.GetBooks(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}

func Test_QidianGetChapters(t *testing.T) {
	urlStr := "https://book.qidian.com/info/1010734492"
	// urlStr = "https://book.qidian.com/info/1004608738"

	reader := QidianReader{}
	list, err := reader.GetChapters(urlStr)
	// list, err = reader.GetCategories()
	if err != nil {

	}
	t.Fatal(list)
}
