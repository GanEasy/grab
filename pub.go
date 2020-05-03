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
	} else if drive == `xxsy` {
		return &reader.XxsyReader{}
	} else if drive == `hongxiu` {
		return &reader.HongxiuReader{}
	} else if drive == `luoqiu` {
		return &reader.MLuoqiuReader{}
	} else if drive == `booktxt` {
		return &reader.BooktxtReader{}
	} else if drive == `paoshu8` {
		return &reader.Paoshu8Reader{}
	} else if drive == `shuge` {
		return &reader.ShugeReader{}
	} else if drive == `qkshu6` {
		return &reader.Qkshu6Reader{}
	} else if drive == `bxwx` {
		return &reader.BxwxReader{}
	} else if drive == `biqugeinfo` {
		return &reader.BiqugeinfoReader{}
	} else if drive == `qu` { // 笔趣阁qula  改域名 jx 了
		return &reader.QuReader{}
	} else if drive == `jx` { // 笔趣阁 jx
		return &reader.JxReader{}
	} else if drive == `uxiaoshuo` { // uxiaoshuo
		return &reader.UxiaoshuoReader{}
	} else if drive == `soe8` {
		return &reader.Soe8Reader{}
	} else if drive == `mcmssc` {
		return &reader.McmsscReader{}
	} else if drive == `xs280` {
		return &reader.Xs280Reader{}
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
	} else if drive == `ssmh` {
		return &reader.SsmhReader{}
	} else if drive == `hanmanku` {
		return &reader.HanmankuReader{}
	} else if drive == `hanmanwo` {
		return &reader.HanmanwoReader{}
	} else if drive == `laosijixs` {
		return &reader.LaosijixsReader{}
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
	} else if drive == `xxsy` {
		return &reader.XxsyReader{}
	} else if drive == `hongxiu` {
		return &reader.HongxiuReader{}
	} else if drive == `luoqiu` {
		return &reader.MLuoqiuReader{}
	} else if drive == `booktxt` {
		return &reader.BooktxtReader{}
	} else if drive == `booktxtnet` {
		return &reader.BooktxtnetReader{}
	} else if drive == `paoshu8` {
		return &reader.Paoshu8Reader{}
	} else if drive == `qkshu6` {
		return &reader.Qkshu6Reader{}
	} else if drive == `shuge` {
		return &reader.ShugeReader{}
	} else if drive == `qu` { // 笔趣阁qula 改域名为 jx.la 了
		return &reader.QuReader{}
	} else if drive == `jx` { // 笔趣阁qula 改域名为 jx.la 了
		return &reader.JxReader{}
	} else if drive == `uxiaoshuo` {
		return &reader.UxiaoshuoReader{}
	} else if drive == `soe8` {
		return &reader.Soe8Reader{}
	} else if drive == `bxwx` {
		return &reader.BxwxReader{}
	} else if drive == `biqugeinfo` {
		return &reader.BiqugeinfoReader{}
	} else if drive == `mcmssc` {
		return &reader.McmsscReader{}
	} else if drive == `xs280` {
		return &reader.Xs280Reader{}
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
	} else if drive == `ssmh` {
		return &reader.SsmhReader{}
	} else if drive == `hanmanku` {
		return &reader.HanmankuReader{}
	} else if drive == `hanmanwo` {
		return &reader.HanmanwoReader{}
	} else if drive == `laosijixs` {
		return &reader.LaosijixsReader{}
	}
	// todo 根据 drive 返回不同的解释器
	return &reader.DefaultReader{}
}
