package grab

// GetListReader 获取列表解释器
func GetListReader(urlstr, interpreter string) ListReader {
	return &DefaultListReader{}
}

// GetInfoReader 获取内容解释器
func GetInfoReader(urlstr, interpreter string) InfoReader {
	// todo 根据 interpreter 返回不同的解释器
	return &DefaultInfoReader{}
}
