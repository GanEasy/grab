package grab

import (
	"testing"
)

func Test_GetBookInfo(t *testing.T) {
	urlStr := `http://www.longfu8.com/417.html`
	info, err := GetBookInfo(urlStr)
	if err != nil {

	}
	t.Fatal(info)

}
