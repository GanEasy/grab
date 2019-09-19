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
	} else if drive == `uxiaoshuo` { // uxiaoshuo
		return &reader.UxiaoshuoReader{}
	} else if drive == `soe8` {
		return &reader.Soe8Reader{}
	} else if drive == `xbiquge` {
		return &reader.XbiqugeReader{}
	} else if drive == `biquyun` {
		return &reader.BiquyunReader{}
	} else if drive == `r2hm` {
		return &reader.R2hmReader{}
	} else if drive == `manwuyu` {
		return &reader.ManwuyuReader{}
	} else if drive == `manhwa` {
		return &reader.ManhwaReader{}
	} else if drive == `aimeizi5` {
		return &reader.Aimeizi5Reader{}
	} else if drive == `kanmeizi` {
		return &reader.KanmeiziReader{}
	} else if drive == `fuman` {
		return &reader.FumanReader{}
	} else if drive == `weijiaoshou` {
		return &reader.WeijiaoshouReader{}
	} else if drive == `haimaoba` {
		return &reader.HaimaobaReader{}
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
	} else if drive == `uxiaoshuo` {
		return &reader.UxiaoshuoReader{}
	} else if drive == `soe8` {
		return &reader.Soe8Reader{}
	} else if drive == `bxwx` {
		return &reader.BxwxReader{}
	} else if drive == `xbiquge` {
		return &reader.XbiqugeReader{}
	} else if drive == `biquyun` {
		return &reader.BiquyunReader{}
	} else if drive == `book` {
		return &reader.BookReader{}
	} else if drive == `article` {
		return &reader.ArticleReader{}
	} else if drive == `rss` {
		return &reader.RssReader{}
	} else if drive == `blog` {
		return &reader.BlogReader{}
	} else if drive == `learnku` {
		return &reader.LearnkuReader{}
	} else if drive == `github` {
		return &reader.GithubReader{}
	} else if drive == `r2hm` {
		return &reader.R2hmReader{}
	} else if drive == `manwuyu` {
		return &reader.ManwuyuReader{}
	} else if drive == `manhwa` {
		return &reader.ManhwaReader{}
	} else if drive == `aimeizi5` {
		return &reader.Aimeizi5Reader{}
	} else if drive == `kanmeizi` {
		return &reader.KanmeiziReader{}
	} else if drive == `fuman` {
		return &reader.FumanReader{}
	} else if drive == `weijiaoshou` {
		return &reader.WeijiaoshouReader{}
	} else if drive == `haimaoba` {
		return &reader.HaimaobaReader{}
	}
	// todo 根据 drive 返回不同的解释器
	return &reader.DefaultReader{}
}
