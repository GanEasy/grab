package grab

// List 列表数据
type List struct {
	// Basic
	Title     string `json:"title"`
	Links     []Link //`json:"links"`
	SourceURL string `json:"source_url"` // 数据抓取时，统一声明数据来源
	Hash      string `json:"hash"`
	Previous  Link   `json:"previous"`
	Next      Link   `json:"next"`
}

// BookSection 小说段落 字数
type BookSection struct {
	Text string `json:"text"`
}

// Book 小说详细
type Book struct {
	URL      string        `json:"url"`
	Title    string        `json:"title"`
	Content  []BookSection `json:"content"`
	PubAt    string        `json:"pub_at"`
	Previous Link          `json:"previous"`
	Next     Link          `json:"next"`
}

// Catalog 第三方资源目录(用户可任意订阅此目录)
type Catalog struct {
	// Basic
	Title     string `json:"title"`
	Links     []Link //`json:"links"`
	SourceURL string `json:"source_url"` // 数据抓取时，统一声明数据来源
	Hash      string `json:"hash"`       // 当前目录 Hash
	Previous  Link   `json:"previous"`   // 如果有上一页
	Next      Link   `json:"next"`       // 如果有下一页
}

// Link 链接
type Link struct {
	// Basic
	Title string `json:"title"`
	URL   string `json:"url"`
	WxTo  string `json:"wxto"` // 小程序跳转到目标页
}

// TextContent 文本类内容正文
type TextContent struct {
	URL      string        `json:"url"`
	Title    string        `json:"title"`
	Content  []BookSection `json:"content"`
	PubAt    string        `json:"pub_at"`
	Previous Link          `json:"previous"`
	Next     Link          `json:"next"`
}
