package reader

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
	From   string   `json:"from"`
}

// Link 链接
type Link struct {
	// Basic
	Title string `json:"title"`
	URL   string `json:"url"`
	WxTo  string `json:"wxto"` // 小程序跳转到目标页
}

//Content reader 返回正文
type Content struct {
	Title     string   `json:"title"`      // 内容标题
	SourceURL string   `json:"source_url"` // 数据来源
	Author    string   `json:"author"`
	PubAt     string   `json:"pub_at"`   //发布时间
	Previous  Link     `json:"previous"` // 上一章
	Next      Link     `json:"next"`     // 下一章
	Contents  []string `json:"contents"` // text正文
	Content   string   `json:"content"`  //新闻(图文)类内容正文
	Images    string   `json:"images"`
	SRC       string   `json:"src"`
	Typw      string   `json:"type"`
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
