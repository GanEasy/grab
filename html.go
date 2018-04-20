package grab

import (
	"github.com/GanEasy/grab/core"
)

// GetHTML 获取rss链接地址中的链接
func GetHTML(urlStr, find string) (html string, err error) {
	html, err = core.GetHTML(urlStr)
	return
}
