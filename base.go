package grab

// List 列表数据
type List struct {
	// Basic
	Title     string `json:"title"`
	Links     []Link
	SourceURL string `json:"source_url"` // 数据抓取时，统一声明数据来源
}

// Link 链接
type Link struct {
	// Basic
	Title string `json:"title"`
	URL   string `json:"url"`
}
