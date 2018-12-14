package grab

// GetListReader 获取列表解释器
func GetListReader(urlstr, drive string) ListReader {
	return &DefaultListReader{}
}

// GetInfoReader 获取内容解释器
func GetInfoReader(urlstr, drive string) InfoReader {
	// todo 根据 drive 返回不同的解释器
	return &DefaultInfoReader{}
}

// GetCategoryReader 在分类页获取小说目录列表
func GetCategoryReader(urlstr, drive string) InfoReader {
	// todo 根据 drive 返回不同的解释器
	return &DefaultInfoReader{}
}

// GetResourceReader 自定义不同平台有哪些资源
func GetResourceReader(urlstr, drive string) InfoReader {
	// todo 根据 drive 返回不同的解释器
	return &DefaultInfoReader{}
}

// GetBookReader 小说资源平台
func GetBookReader(drive string) BookReader {
	//
	if drive == `qidian` {
		return &QidianReader{}
	} else if drive == `zongheng` {
		return &ZonghengReader{}
	} else if drive == `17k` {
		return &SeventeenKReader{}
	} else if drive == `luoqiu` {
		return &MLuoqiuReader{}
	} else if drive == `booktxt` {
		return &BooktxtReader{}
	} else if drive == `7878xs` {
		return &Xs7878Reader{}
	}
	// todo 根据 drive 返回不同的解释器
	return &DefaultBookReader{}
}
