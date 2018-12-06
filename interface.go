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
