package grab

// List 列表数据
type List struct {
	// Basic
	Title     string `json:"title"`
	Links     []Link //`json:"links"`
	SourceURL string `json:"source_url"` // 数据抓取时，统一声明数据来源
	Hash      string `json:"hash"`
}

// Link 链接
type Link struct {
	// Basic
	Title string `json:"title"`
	URL   string `json:"url"`
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

// // Article 文章
// type Article struct {
// 	Title   string `json:"title"`
// 	Content string `json:"content"`
// 	PubAt   string `json:"pub_at"`
// }

//RssList RssList
func RssList(urlStr string) {

}

//HTMLList html list
func HTMLList(urlStr string) {

}

//ArticleList 文章列表(博客文章类)
func ArticleList(urlStr string) {

}

//BookList 小说类列表 (相似度高)
func BookList(urlStr string) {

}

//CustomList 自定义 匹配规则列表
func CustomList(urlStr string) {

}

//LearnList 学习匹配列表
func LearnList(urlStr string) {

}

//HTMLInfo html list
func HTMLInfo(urlStr string) {

}

//BookInfo 小说类列表 (相似度高)
func BookInfo(urlStr string) {

}
