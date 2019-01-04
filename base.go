package grab

import (
	"encoding/base64"
	"net/url"
)

/**
* base.go 放一些基础数据结构类,用于制定内外数据结构
 */
// Catalog 第三方资源目录(用户可任意订阅此目录)
type Catalog struct {
	Title     string `json:"title"`
	Cards     []Card //`json:"links"`
	SourceURL string `json:"source_url"` // 数据抓取时，统一声明数据来源
	Hash      string `json:"hash"`       // 当前目录 Hash
	Previous  Link   `json:"previous"`   // 如果有上一页
	Next      Link   `json:"next"`       // 如果有下一页
}

// Card 使用卡片代替链接
type Card struct {
	Title  string   `json:"title"`  // 标题
	WxTo   string   `json:"wxto"`   // 小程序跳转到目标页
	Intro  string   `json:"intro"`  //介绍
	Type   string   `json:"type"`   // card展示形式 media card text image images link
	Cover  string   `json:"cover"`  // 封面图片
	Images []string `json:"images"` // 图片组效果时图片列表
}

// Link 链接
type Link struct {
	// Basic
	Title string `json:"title"`
	URL   string `json:"url"`
	WxTo  string `json:"wxto"` // 小程序跳转到目标页
}

// DemoItem 示例详细
type DemoItem struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	Category string `json:"category"`
}

// Item 小程序授受参数明细
type Item struct {
	Title string `json:"title"`
	WxTo  string `json:"wxto"`
	Intro string `json:"intro"`
	Type  string `json:"type"`
}

//EncodeURL 把url encode
func EncodeURL(str string) string {
	// es := base64.URLEncoding.EncodeToString([]byte(str))
	return url.QueryEscape(base64.URLEncoding.EncodeToString([]byte(str)))
	// return encodeURIComponent(es)
}

//DecodeURL 把url decode
func DecodeURL(str string) (string, error) {
	es, err := url.QueryUnescape(str)
	strByte, err := base64.URLEncoding.DecodeString(es)
	return string(strByte), err
}
