package reader

// GetReader 获取阅读器
func GetReader(drive string) Reader {
	//
	if drive == `qidian` {
		// return &QidianReader{}
	} else if drive == `zongheng` {
		// return &ZonghengReader{}
	} else if drive == `17k` {
		// return &SeventeenKReader{}
	} else if drive == `luoqiu` {
		// return &MLuoqiuReader{}
	} else if drive == `booktxt` {
		// return &BooktxtReader{}
	} else if drive == `7878xs` {
		// return &Xs7878Reader{}
	}
	// todo 根据 drive 返回不同的解释器
	return &DefaultReader{}
}
