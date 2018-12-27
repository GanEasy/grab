package reader

//Reader 目录资源阅读器
/**输出第三方平台资源
 */
type Reader interface {
	// 获取目录
	GetCatalog(string) (Catalog, error)
	// 获取详情
	GetInfo(string) (Content, error)
}

//Guide 第三方平台资源引导
/**
 */
type Guide interface {
	// 第三方平台分类列表(多是我们自行定义)
	GetCategories(string) (Catalog, error)
	// 获取第三方目录列表
	GetList(string) (Catalog, error)
}
