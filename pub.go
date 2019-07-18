package grab

import (
	"github.com/GanEasy/grab/reader"
)

// GetGuide 获取引导
func GetGuide(drive string) reader.Guide {
	//
	if drive == `qidian` {
		return &reader.QidianReader{}
	} else if drive == `zongheng` {
		return &reader.ZonghengReader{}
	} else if drive == `17k` {
		return &reader.SeventeenKReader{}
	} else if drive == `luoqiu` {
		return &reader.MLuoqiuReader{}
	} else if drive == `booktxt` {
		return &reader.BooktxtReader{}
	} else if drive == `bxwx` {
		return &reader.BxwxReader{}
	} else if drive == `uxiaoshuo` {
		return &reader.UxiaoshuoReader{}
	} else if drive == `r2hm` {
		return &reader.R2hmReader{}
	}
	// todo 根据 drive 返回不同的解释器
	return &reader.DefaultGuide{}
}

// GetReader 获取阅读器
func GetReader(drive string) reader.Reader {
	//
	if drive == `qidian` {
		return &reader.QidianReader{}
	} else if drive == `zongheng` {
		return &reader.ZonghengReader{}
	} else if drive == `17k` {
		return &reader.SeventeenKReader{}
	} else if drive == `luoqiu` {
		return &reader.MLuoqiuReader{}
	} else if drive == `booktxt` {
		return &reader.BooktxtReader{}
	} else if drive == `7878xs` {
		return &reader.Xs7878Reader{}
	} else if drive == `bxwx` {
		return &reader.BxwxReader{}
	} else if drive == `book` {
		return &reader.BookReader{}
	} else if drive == `article` {
		return &reader.ArticleReader{}
	} else if drive == `rss` {
		return &reader.RssReader{}
	} else if drive == `blog` {
		return &reader.BlogReader{}
	} else if drive == `r2hm` {
		return &reader.R2hmReader{}
	}
	// todo 根据 drive 返回不同的解释器
	return &reader.DefaultReader{}
}
