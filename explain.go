package grab

import (
	"fmt"
	"regexp"

	"github.com/GanEasy/grab/reader"
)

// ExplainLink 解释链接地址
func ExplainLink(url string) (address, drive, page string) {

	// 检查是不是起点地址
	if checkLinkIsQiDian, _ := regexp.MatchString(`qidian.com`, url); checkLinkIsQiDian {
		// 起点详细页
		//http://book.qidian.com/info/1004608738
		InfoBook := `book.qidian.com\/info\/(?P<book_id>\d+)`
		if b, _ := regexp.MatchString(InfoBook, url); b {
			Map := reader.SelectString(InfoBook, url)
			return fmt.Sprintf("http://book.qidian.com/info/%v", Map["book_id"]), `qidian`, `/pages/catalog`
		}

		// 起点手机详细页
		//http://m.qidian.com/book/1004608738
		MobileBook := `m.qidian.com\/book\/(?P<book_id>\d+)`
		if b, _ := regexp.MatchString(MobileBook, url); b {
			Map := reader.SelectString(MobileBook, url)
			return fmt.Sprintf("http://book.qidian.com/info/%v", Map["book_id"]), `qidian`, `/pages/catalog`
		}

		// 起点手机章节列表页
		//http://m.qidian.com/book/1004608738/catalog
		MobileBookChapterMenu := `m.qidian.com\/book\/(?P<book_id>\d+)\/catalog`
		if b, _ := regexp.MatchString(MobileBookChapterMenu, url); b {
			Map := reader.SelectString(MobileBookChapterMenu, url)
			return fmt.Sprintf("http://book.qidian.com/info/%v", Map["book_id"]), `qidian`, `/pages/catalog`
		}

		// 手机章节详细页
		//http://m.qidian.com/book/1004608738/342363924
		MobileBookChapterInfo := `m.qidian.com\/book\/(?P<book_id>\d+)\/(?P<chapter_id>\d+)`
		if b, _ := regexp.MatchString(MobileBookChapterInfo, url); b {
			Map := reader.SelectString(MobileBookChapterInfo, url)
			return fmt.Sprintf("http://book.qidian.com/info/%v", Map["book_id"]), `qidian`, `/pages/catalog`
		}

		BookVIPChapterInfo := `vipreader.qidian.com\/chapter\/(?P<book_id>\d+)\/(?P<chapter_id>\d+)`
		if b, _ := regexp.MatchString(BookVIPChapterInfo, url); b {
			Map := reader.SelectString(BookVIPChapterInfo, url)
			return fmt.Sprintf("http://book.qidian.com/info/%v", Map["book_id"]), `qidian`, `/pages/catalog`
		}

		// todo http://read.qidian.com/chapter/_AaqI-dPJJ4uTkiRw_sFYA2/-Yjl2ADCXQvM5j8_3RRvhw2
	}

	// 检查是不是纵横地址
	if checkLinkIsZongHeng, _ := regexp.MatchString(`zongheng.com`, url); checkLinkIsZongHeng {
		// 纵横详细页
		// http://book.zongheng.com/book/490607.html
		InfoBook := `book.zongheng.com\/book\/(?P<book_id>\d+).html`
		if b, _ := regexp.MatchString(InfoBook, url); b {
			Map := reader.SelectString(InfoBook, url)
			return fmt.Sprintf("http://book.zongheng.com/showchapter/%v.html", Map["book_id"]), `zongheng`, `/pages/catalog`
		}

		// http://book.zongheng.com/chapter/672340/38144043.html
		ChapterInfo := `book.zongheng.com\/chapter\/(?P<book_id>\d+)\/(?P<chapter_id>\d+).html`
		if b, _ := regexp.MatchString(ChapterInfo, url); b {
			Map := reader.SelectString(ChapterInfo, url)
			return fmt.Sprintf("http://book.zongheng.com/showchapter/%v.html", Map["book_id"]), `zongheng`, `/pages/catalog`

		}

		// http://book.zongheng.com/showchapter/672340.html
		BookChapterMenu := `book.zongheng.com\/showchapter\/(?P<book_id>\d+).html`
		if b, _ := regexp.MatchString(BookChapterMenu, url); b {
			Map := reader.SelectString(BookChapterMenu, url)
			return fmt.Sprintf("http://book.zongheng.com/showchapter/%v.html", Map["book_id"]), `zongheng`, `/pages/catalog`

		}
		// 纵横手机详细页
		// http://m.zongheng.com/h5/book?bookid=490607
		MobileBook := `m.zongheng.com\/h5\/book\?bookid=(?P<book_id>\d+)`
		if b, _ := regexp.MatchString(MobileBook, url); b {
			Map := reader.SelectString(MobileBook, url)
			return fmt.Sprintf("http://book.zongheng.com/showchapter/%v.html", Map["book_id"]), `zongheng`, `/pages/catalog`

		}

		// 纵横手机章节列表页
		// http://m.zongheng.com/h5/chapter/list?bookid=490607
		MobileBookChapterMenu := `m.zongheng.com\/h5\/chapter\/list\?bookid=(?P<book_id>\d+)`
		if b, _ := regexp.MatchString(MobileBookChapterMenu, url); b {
			Map := reader.SelectString(MobileBookChapterMenu, url)
			return fmt.Sprintf("http://book.zongheng.com/showchapter/%v.html", Map["book_id"]), `zongheng`, `/pages/catalog`

		}

		// 起点手机章节列表页
		//http://m.qidian.com/book/1004608738/342363924
		// http://m.zongheng.com/h5/chapter?bookid=490607&cid=8134632
		MobileBookChapterInfo := `m.zongheng.com\/h5\/chapter\?bookid=(?P<book_id>\d+)&cid=(?P<chapter_id>\d+)`
		if b, _ := regexp.MatchString(MobileBookChapterInfo, url); b {
			Map := reader.SelectString(MobileBookChapterInfo, url)
			return fmt.Sprintf("http://book.zongheng.com/showchapter/%v.html", Map["book_id"]), `zongheng`, `/pages/catalog`

		}
	}

	// 检查是不是17k地址
	if checkLinkIsSeventeenK, _ := regexp.MatchString(`17k.com`, url); checkLinkIsSeventeenK {
		// 17k详细页
		InfoBook := `17k.com\/book\/(?P<book_id>\d+).html`
		if b, _ := regexp.MatchString(InfoBook, url); b {
			Map := reader.SelectString(InfoBook, url)
			return fmt.Sprintf("http://www.17k.com/book/%v.html", Map["book_id"]), `17k`, `/pages/catalog`
		}

		// 章节列表
		// 17k.com/book/2317974.html
		BookChapterMenu := `17k.com\/list\/(?P<book_id>\d+).html`
		if b, _ := regexp.MatchString(BookChapterMenu, url); b {
			Map := reader.SelectString(BookChapterMenu, url)
			return fmt.Sprintf("http://www.17k.com/book/%v.html", Map["book_id"]), `17k`, `/pages/catalog`
		}

		// 章节详细
		MobileBookChapterInfo := `17k.com\/chapter\/(?P<book_id>\d+)/(?P<chapter_id>\d+).html`
		if b, _ := regexp.MatchString(MobileBookChapterInfo, url); b {
			Map := reader.SelectString(MobileBookChapterInfo, url)
			return fmt.Sprintf("http://www.17k.com/book/%v.html", Map["book_id"]), `17k`, `/pages/catalog`
		}
	}

	// 检查是不是biquge.info
	if checkLinkIsBiqugeinfo, _ := regexp.MatchString(`biquge.info`, url); checkLinkIsBiqugeinfo {

		// 章节详细 https://m.biquge.info/10_10218/5002113.html
		MobileBookChapterInfo := `m.biquge.info\/(?P<cate_id>\d+)_(?P<book_id>\d+)/(?P<chapter_id>\d+).html`
		if b, _ := regexp.MatchString(MobileBookChapterInfo, url); b {
			Map := reader.SelectString(MobileBookChapterInfo, url)
			return fmt.Sprintf("http://www.biquge.info/%v_%v/%v.html", Map["cate_id"], Map["book_id"], Map["chapter_id"]), `biqugeinfo`, `/pages/book`
		}
		// 章节列表
		// https://m.biquge.info/10_10218/
		BookChapterMenu := `m.biquge.info\/(?P<cate_id>\d+)_(?P<book_id>\d+)\/`
		if b, _ := regexp.MatchString(BookChapterMenu, url); b {
			Map := reader.SelectString(BookChapterMenu, url)
			return fmt.Sprintf("http://www.biquge.info/%v_%v/", Map["cate_id"], Map["book_id"]), `biqugeinfo`, `/pages/catalog`
		}
		// 其它的当作列表页
		BookList := `m.biquge.info/(?P<path>.*)`
		if b, _ := regexp.MatchString(BookList, url); b {
			return url, `biqugeinfo`, `/pages/list`
		}
	}

	// 检查是不是biquyun.com
	if checkLinkIsBiquyun, _ := regexp.MatchString(`biquyun.com`, url); checkLinkIsBiquyun {

		// 章节详细 https://m.biquyun.com/16_16635/10124285.html
		BookChapterInfo := `m.biquyun.com\/(?P<cate_id>\d+)_(?P<book_id>\d+)/(?P<chapter_id>\d+).html`
		if b, _ := regexp.MatchString(BookChapterInfo, url); b {
			Map := reader.SelectString(BookChapterInfo, url)
			return fmt.Sprintf("https://m.biquyun.com/%v_%v/%v.html", Map["cate_id"], Map["book_id"], Map["chapter_id"]), `biqugeinfo`, `/pages/book`
		}
		// 章节列表
		// https://m.biquyun.com/16_16635/
		BookChapterMenu := `m.biquyun.com\/(?P<cate_id>\d+)_(?P<book_id>\d+)\/`
		if b, _ := regexp.MatchString(BookChapterMenu, url); b {
			Map := reader.SelectString(BookChapterMenu, url)
			return fmt.Sprintf("https://m.biquyun.com/%v_%v_1_1.html", Map["cate_id"], Map["book_id"]), `biquyun`, `/pages/catalog`

		}
		BookChapterMenu2 := `m.biquyun.com\/(?P<cate_id>\d+)_(?P<book_id>\d+)_(?P<page>\d+)_1.html`
		if b, _ := regexp.MatchString(BookChapterMenu2, url); b {
			Map := reader.SelectString(BookChapterMenu2, url)
			return fmt.Sprintf("https://m.biquyun.com/%v_%v_%v_1.html", Map["cate_id"], Map["book_id"], Map["page"]), `biquyun`, `/pages/catalog`

		}
		// 其它的当作列表页
		BookList := `m.biquyun.com/(?P<path>.*)`
		if b, _ := regexp.MatchString(BookList, url); b {
			return url, `biquyun`, `/pages/list`
		}
	}

	// 检查是不是booktxt.net
	if checkLinkIsBooktxt, _ := regexp.MatchString(`booktxt.net`, url); checkLinkIsBooktxt {

		// 章节详细 https://m.booktxt.net/wapbook/4891_4943641.html
		BookChapterInfo := `m.booktxt.net\/wapbook\/(?P<book_id>\d+)_(?P<chapter_id>\d+).html`
		if b, _ := regexp.MatchString(BookChapterInfo, url); b {
			return url, `booktxt`, `/pages/book`
		}
		// 章节列表
		// https://m.booktxt.net/wapbook/6454.html
		BookChapterMenu := `m.booktxt.net\/wapbook\/(?P<book_id>\d+).html`
		if b, _ := regexp.MatchString(BookChapterMenu, url); b {
			return url, `booktxt`, `/pages/catalog`
		}
		// 其它的当作列表页
		BookList := `m.booktxt.net/(?P<path>.*)`
		if b, _ := regexp.MatchString(BookList, url); b {
			return url, `booktxt`, `/pages/list`
		}
	}

	// 检查是不是bxwx.la
	if checkLinkIsBxwx, _ := regexp.MatchString(`bxwx.la`, url); checkLinkIsBxwx {

		// 章节详细 https://m.bxwx.la/b/316/316850/1684236.html
		BookChapterInfo := `m.bxwx.la\/b\/(?P<cate_id>\d+)\/(?P<book_id>\d+)\/(?P<chapter_id>\d+).html`
		if b, _ := regexp.MatchString(BookChapterInfo, url); b {
			return url, `bxwx`, `/pages/book`
		}
		// 章节列表
		// https://m.bxwx.la/binfo/246/246596.htm
		BookChapterMenu := `m.bxwx.la\/binfo\/(?P<cate_id>\d+)\/(?P<book_id>\d+).htm`
		if b, _ := regexp.MatchString(BookChapterMenu, url); b {
			return url, `bxwx`, `/pages/catalog`
		}
		// 其它的当作列表页
		BookList := `m.bxwx.la/(?P<path>.*)`
		if b, _ := regexp.MatchString(BookList, url); b {
			return url, `bxwx`, `/pages/list`
		}
	}

	// 检查是不是mcmssc.com
	if checkLinkIsBxwx, _ := regexp.MatchString(`mcmssc.com`, url); checkLinkIsBxwx {

		// 章节详细 https://www.mcmssc.com/44_44569/21647159.html
		BookChapterInfo := `mcmssc.com\/(?P<cate_id>\d+)_(?P<book_id>\d+)\/(?P<chapter_id>\d+).html`
		if b, _ := regexp.MatchString(BookChapterInfo, url); b {
			Map := reader.SelectString(BookChapterInfo, url)
			return fmt.Sprintf("https://www.mcmssc.com/%v_%v/%v.html", Map["cate_id"], Map["book_id"], Map["chapter_id"]), `mcmssc`, `/pages/book`

		}
		// 章节列表
		// https://www.mcmssc.com/44_44569/
		BookChapterMenu := `mcmssc.com\/(?P<cate_id>\d+)_(?P<book_id>\d+)\/`
		if b, _ := regexp.MatchString(BookChapterMenu, url); b {
			Map := reader.SelectString(BookChapterMenu, url)
			return fmt.Sprintf("https://www.mcmssc.com/%v_%v/", Map["cate_id"], Map["book_id"]), `mcmssc`, `/pages/catalog`
		}
		// https://m.mcmssc.com/44_44569/
		BookChapterMenu2 := `mcmssc.com\/(?P<cate_id>\d+)_(?P<book_id>\d+)\/all.html`
		if b, _ := regexp.MatchString(BookChapterMenu2, url); b {
			Map := reader.SelectString(BookChapterMenu2, url)
			return fmt.Sprintf("https://www.mcmssc.com/%v_%v/", Map["cate_id"], Map["book_id"]), `mcmssc`, `/pages/catalog`
		}
		// 其它的当作列表页
		BookList := `mcmssc.com/(?P<path>.*)`
		if b, _ := regexp.MatchString(BookList, url); b {
			return url, `mcmssc`, `/pages/list`
		}
	}
	// 检查是不是uxiaoshuo.com
	if checkLinkIsUxs, _ := regexp.MatchString(`uxiaoshuo.com`, url); checkLinkIsUxs {

		// 章节详细 https://m.uxiaoshuo.com/40/40423/2193344.html
		BookChapterInfo := `m.uxiaoshuo.com\/(?P<cate_id>\d+)\/(?P<book_id>\d+)\/(?P<chapter_id>\d+)\/.html`
		if b, _ := regexp.MatchString(BookChapterInfo, url); b {
			Map := reader.SelectString(BookChapterInfo, url)
			return fmt.Sprintf("https://m.uxiaoshuo.com/%v/%v/%v.html", Map["cate_id"], Map["book_id"], Map["chapter_id"]), `uxiaoshuo`, `/pages/book`

		}
		// 章节列表
		// https://m.uxiaoshuo.com/40/40423/all.html
		BookChapterMenu := `m.uxiaoshuo.com\/(?P<cate_id>\d+)\/(?P<book_id>\d+)\/all.html`
		if b, _ := regexp.MatchString(BookChapterMenu, url); b {
			Map := reader.SelectString(BookChapterMenu, url)
			return fmt.Sprintf("https://m.uxiaoshuo.com/%v/%v/all.html", Map["cate_id"], Map["book_id"]), `uxiaoshuo`, `/pages/catalog`
		}
		// https://m.uxiaoshuo.com/40/40423/
		BookChapterMenu2 := `m.uxiaoshuo.com\/(?P<cate_id>\d+)\/(?P<book_id>\d+)\/`
		if b, _ := regexp.MatchString(BookChapterMenu2, url); b {
			Map := reader.SelectString(BookChapterMenu2, url)
			return fmt.Sprintf("https://m.uxiaoshuo.com/%v/%v/all.html", Map["cate_id"], Map["book_id"]), `uxiaoshuo`, `/pages/catalog`
		}
		// 其它的当作列表页
		BookList := `uxiaoshuo.com/(?P<path>.*)`
		if b, _ := regexp.MatchString(BookList, url); b {
			return url, `uxiaoshuo`, `/pages/list`
		}
	}

	// 检查是不是 laosijixs.com
	if checkLinkIsBiquyun, _ := regexp.MatchString(`laosijixs.com`, url); checkLinkIsBiquyun {

		// 章节详细 http://m.laosijixs.com/81/81228/5940566.html
		BookChapterInfo := `laosijixs.com\/(?P<cate_id>\d+)\/(?P<book_id>\d+)\/(?P<chapter_id>\d+).html`
		if b, _ := regexp.MatchString(BookChapterInfo, url); b {
			Map := reader.SelectString(BookChapterInfo, url)
			return fmt.Sprintf("http://m.laosijixs.com/%v/%v/%v.html", Map["cate_id"], Map["book_id"], Map["chapter_id"]), `laosijixs`, `/pages/book`
		}
		// 章节详细 http://m.laosijixs.com/81/81228/5940566_1.html
		BookChapterInfo2 := `laosijixs.com\/(?P<cate_id>\d+)\/(?P<book_id>\d+)\/(?P<chapter_id>\d+)_(?P<page>\d+).html`
		if b, _ := regexp.MatchString(BookChapterInfo2, url); b {
			Map := reader.SelectString(BookChapterInfo2, url)
			return fmt.Sprintf("http://m.laosijixs.com/%v/%v/%v_%v.html", Map["cate_id"], Map["book_id"], Map["chapter_id"], Map["page"]), `laosijixs`, `/pages/book`
		}
		// 章节列表
		// https://m.biquyun.com/16_16635/
		BookChapterMenu := `laosijixs.com\/(?P<cate_id>\d+)\/(?P<book_id>\d+)\/`
		if b, _ := regexp.MatchString(BookChapterMenu, url); b {
			Map := reader.SelectString(BookChapterMenu, url)
			return fmt.Sprintf("http://m.laosijixs.com/%v/%v/", Map["cate_id"], Map["book_id"]), `laosijixs`, `/pages/catalog`

		}
		BookChapterMenu2 := `laosijixs.com\/(?P<cate_id>\d+)\/(?P<book_id>\d+)_(?P<page>\d+)\/`
		if b, _ := regexp.MatchString(BookChapterMenu2, url); b {
			Map := reader.SelectString(BookChapterMenu2, url)
			return fmt.Sprintf("http://m.laosijixs.com/%v/%v_%v/", Map["cate_id"], Map["book_id"], Map["page"]), `laosijixs`, `/pages/catalog`

		}
		// 其它的当作列表页
		BookList := `laosijixs.com/(?P<path>.*)`
		if b, _ := regexp.MatchString(BookList, url); b {
			return url, `laosijixs`, `/pages/list`
		}
	}

	return ``, ``, ``
}
