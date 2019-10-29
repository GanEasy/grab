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
			return fmt.Sprintf("http://book.qidian.com/info/%v", Map["book_id"]), `qidian`, `/page/catelog`
		}

		// 起点手机详细页
		//http://m.qidian.com/book/1004608738
		MobileBook := `m.qidian.com\/book\/(?P<book_id>\d+)`
		if b, _ := regexp.MatchString(MobileBook, url); b {
			Map := reader.SelectString(MobileBook, url)
			return fmt.Sprintf("http://book.qidian.com/info/%v", Map["book_id"]), `qidian`, `/page/catelog`
		}

		// 起点手机章节列表页
		//http://m.qidian.com/book/1004608738/catalog
		MobileBookChapterMenu := `m.qidian.com\/book\/(?P<book_id>\d+)\/catalog`
		if b, _ := regexp.MatchString(MobileBookChapterMenu, url); b {
			Map := reader.SelectString(MobileBookChapterMenu, url)
			return fmt.Sprintf("http://book.qidian.com/info/%v", Map["book_id"]), `qidian`, `/page/catelog`
		}

		// 手机章节详细页
		//http://m.qidian.com/book/1004608738/342363924
		MobileBookChapterInfo := `m.qidian.com\/book\/(?P<book_id>\d+)\/(?P<chapter_id>\d+)`
		if b, _ := regexp.MatchString(MobileBookChapterInfo, url); b {
			Map := reader.SelectString(MobileBookChapterInfo, url)
			return fmt.Sprintf("http://book.qidian.com/info/%v", Map["book_id"]), `qidian`, `/page/catelog`
		}

		BookVIPChapterInfo := `vipreader.qidian.com\/chapter\/(?P<book_id>\d+)\/(?P<chapter_id>\d+)`
		if b, _ := regexp.MatchString(BookVIPChapterInfo, url); b {
			Map := reader.SelectString(BookVIPChapterInfo, url)
			return fmt.Sprintf("http://book.qidian.com/info/%v", Map["book_id"]), `qidian`, `/page/catelog`
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
			return fmt.Sprintf("http://book.zongheng.com/showchapter/%v.html", Map["book_id"]), `zongheng`, `/page/catelog`
		}

		// http://book.zongheng.com/chapter/672340/38144043.html
		ChapterInfo := `book.zongheng.com\/chapter\/(?P<book_id>\d+)\/(?P<chapter_id>\d+).html`
		if b, _ := regexp.MatchString(ChapterInfo, url); b {
			Map := reader.SelectString(ChapterInfo, url)
			return fmt.Sprintf("http://book.zongheng.com/showchapter/%v.html", Map["book_id"]), `zongheng`, `/page/catelog`

		}

		// http://book.zongheng.com/showchapter/672340.html
		BookChapterMenu := `book.zongheng.com\/showchapter\/(?P<book_id>\d+).html`
		if b, _ := regexp.MatchString(BookChapterMenu, url); b {
			Map := reader.SelectString(BookChapterMenu, url)
			return fmt.Sprintf("http://book.zongheng.com/showchapter/%v.html", Map["book_id"]), `zongheng`, `/page/catelog`

		}
		// 纵横手机详细页
		// http://m.zongheng.com/h5/book?bookid=490607
		MobileBook := `m.zongheng.com\/h5\/book\?bookid=(?P<book_id>\d+)`
		if b, _ := regexp.MatchString(MobileBook, url); b {
			Map := reader.SelectString(MobileBook, url)
			return fmt.Sprintf("http://book.zongheng.com/showchapter/%v.html", Map["book_id"]), `zongheng`, `/page/catelog`

		}

		// 纵横手机章节列表页
		// http://m.zongheng.com/h5/chapter/list?bookid=490607
		MobileBookChapterMenu := `m.zongheng.com\/h5\/chapter\/list\?bookid=(?P<book_id>\d+)`
		if b, _ := regexp.MatchString(MobileBookChapterMenu, url); b {
			Map := reader.SelectString(MobileBookChapterMenu, url)
			return fmt.Sprintf("http://book.zongheng.com/showchapter/%v.html", Map["book_id"]), `zongheng`, `/page/catelog`

		}

		// 起点手机章节列表页
		//http://m.qidian.com/book/1004608738/342363924
		// http://m.zongheng.com/h5/chapter?bookid=490607&cid=8134632
		MobileBookChapterInfo := `m.zongheng.com\/h5\/chapter\?bookid=(?P<book_id>\d+)&cid=(?P<chapter_id>\d+)`
		if b, _ := regexp.MatchString(MobileBookChapterInfo, url); b {
			Map := reader.SelectString(MobileBookChapterInfo, url)
			return fmt.Sprintf("http://book.zongheng.com/showchapter/%v.html", Map["book_id"]), `zongheng`, `/page/catelog`

		}
	}

	// 检查是不是17k地址
	if checkLinkIsSeventeenK, _ := regexp.MatchString(`17k.com`, url); checkLinkIsSeventeenK {
		// 17k详细页
		InfoBook := `17k.com\/book\/(?P<book_id>\d+).html`
		if b, _ := regexp.MatchString(InfoBook, url); b {
			Map := reader.SelectString(InfoBook, url)
			return fmt.Sprintf("http://www.17k.com/book/%v.html", Map["book_id"]), `17k`, `/page/catelog`
		}

		// 章节列表
		// 17k.com/book/2317974.html
		BookChapterMenu := `17k.com\/list\/(?P<book_id>\d+).html`
		if b, _ := regexp.MatchString(BookChapterMenu, url); b {
			Map := reader.SelectString(BookChapterMenu, url)
			return fmt.Sprintf("http://www.17k.com/book/%v.html", Map["book_id"]), `17k`, `/page/catelog`
		}

		// 章节详细
		MobileBookChapterInfo := `17k.com\/chapter\/(?P<book_id>\d+)/(?P<chapter_id>\d+).html`
		if b, _ := regexp.MatchString(MobileBookChapterInfo, url); b {
			Map := reader.SelectString(MobileBookChapterInfo, url)
			return fmt.Sprintf("http://www.17k.com/book/%v.html", Map["book_id"]), `17k`, `/page/catelog`
		}
	}

	// 检查是不是biquge.info
	if checkLinkIsBiqugeinfo, _ := regexp.MatchString(`biquge.info`, url); checkLinkIsBiqugeinfo {

		// 章节详细 https://m.biquge.info/10_10218/5002113.html
		MobileBookChapterInfo := `m.biquge.info\/(?P<cate_id>\d+)_(?P<book_id>\d+)/(?P<chapter_id>\d+).html`
		if b, _ := regexp.MatchString(MobileBookChapterInfo, url); b {
			Map := reader.SelectString(MobileBookChapterInfo, url)
			return fmt.Sprintf("https://m.biquge.info/%v_%v/%v.html", Map["cate_id"], Map["book_id"], Map["chapter_id"]), `biqugeinfo`, `/page/book`
		}
		// 章节列表
		// https://m.biquge.info/10_10218/
		BookChapterMenu := `m.biquge.info\/(?P<cate_id>\d+)_(?P<book_id>\d+)\/`
		if b, _ := regexp.MatchString(BookChapterMenu, url); b {
			// Map := reader.SelectString(BookChapterMenu, url)
			return url, `biqugeinfo`, `/page/catelog`
		}
		// 其它的当作列表页
		BookList := `m.biquge.info/(?P<path>.*)`
		if b, _ := regexp.MatchString(BookList, url); b {
			return url, `biqugeinfo`, `/page/list`
		}
	}

	// 检查是不是biquyun.com
	if checkLinkIsBiquyun, _ := regexp.MatchString(`biquyun.com`, url); checkLinkIsBiquyun {

		// 章节详细 https://m.biquyun.com/16_16635/10124285.html
		BookChapterInfo := `m.biquyun.com\/(?P<cate_id>\d+)_(?P<book_id>\d+)/(?P<chapter_id>\d+).html`
		if b, _ := regexp.MatchString(BookChapterInfo, url); b {
			Map := reader.SelectString(BookChapterInfo, url)
			return fmt.Sprintf("https://m.biquyun.com/%v_%v/%v.html", Map["cate_id"], Map["book_id"], Map["chapter_id"]), `biqugeinfo`, `/page/book`
		}
		// 章节列表
		// https://m.biquyun.com/16_16635/
		BookChapterMenu := `m.biquyun.com\/(?P<cate_id>\d+)_(?P<book_id>\d+)\/`
		if b, _ := regexp.MatchString(BookChapterMenu, url); b {
			Map := reader.SelectString(BookChapterMenu, url)
			return fmt.Sprintf("https://m.biquyun.com/%v_%v_1_1.html", Map["cate_id"], Map["book_id"]), `biquyun`, `/page/catelog`

		}
		BookChapterMenu2 := `m.biquyun.com\/(?P<cate_id>\d+)_(?P<book_id>\d+)_(?P<page>\d+)_1.html`
		if b, _ := regexp.MatchString(BookChapterMenu2, url); b {
			Map := reader.SelectString(BookChapterMenu2, url)
			return fmt.Sprintf("https://m.biquyun.com/%v_%v_%v_1.html", Map["cate_id"], Map["book_id"], Map["page"]), `biquyun`, `/page/catelog`

		}
		// 其它的当作列表页
		BookList := `m.biquyun.com/(?P<path>.*)`
		if b, _ := regexp.MatchString(BookList, url); b {
			return url, `biquyun`, `/page/list`
		}
	}
	return ``, ``, ``
}
