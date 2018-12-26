package grab

//ListReader 列表获取器
type ListReader interface {
	GetList()
}

//InfoReader 内容获取器
type InfoReader interface {
	GetInfo()
	// GetNextURL() string
}

//BookReader 小说类网站阅读器
type BookReader interface {
	GetCategories(string) (Catalog, error)
	GetBooks(string) (Catalog, error)
	GetChapters(string) (Catalog, error)
	GetChapter(string) (TextContent, error)
}

//Reader 目录资源阅读器
/**输出第三方平台资源
 */
type Reader interface {
	Catalog(string) (Catalog, error)
	Info(string) (ReaderContent, error)
}
