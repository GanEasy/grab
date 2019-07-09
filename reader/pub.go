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

// Score2Level 积分阶梯
func Score2Level(x int) int {
	if x < 10 {
		return 0
	} else if x < 100 {
		return 1
	} else if x < 1000 {
		return 2
	} else if x < 10000 {
		return 3
	} else if x < 100000 {
		return 4
	} else if x < 1000000 {
		return 5
	} else {
		return 0
	}
}
