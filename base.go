package grab

/**
* base.go 放一些基础数据结构类,用于制定内外数据结构
 */
// Catalog 第三方资源目录(用户可任意订阅此目录)
type Catalog struct {
	Title     string `json:"title"`
	Links     []Link //`json:"links"`
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

// TextContent 文本类内容正文
type TextContent struct {
	Title     string        `json:"title"`
	Content   []BookSection `json:"content"`
	SourceURL string        `json:"source_url"`
	PubAt     string        `json:"pub_at"`
	Previous  Link          `json:"previous"`
	Next      Link          `json:"next"`
}

// NewsContent 新闻类内容正文
type NewsContent struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	PubAt     string `json:"pub_at"`
	SourceURL string `json:"source_url"` // 数据抓取时，统一声明数据来源
	Previous  Link   `json:"previous"`
	Next      Link   `json:"next"`
}

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

// Link 链接
type Link struct {
	// Basic
	Title string `json:"title"`
	URL   string `json:"url"`
	WxTo  string `json:"wxto"` // 小程序跳转到目标页
}
