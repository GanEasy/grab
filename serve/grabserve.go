package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/GanEasy/grab"
	a "github.com/GanEasy/grab/api"
	cpi "github.com/GanEasy/grab/core"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//删除目录下的文件信息
//dirpath 目录路径
func delDirFile(dirpath string) {
	//读取目录信息
	dir, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return
	}

	//当前系统的时间
	ct := int32(time.Now().Unix())
	//12小时
	spt := int32(12 * 3600)

	for _, file := range dir {
		//读取到的是目录
		if file.IsDir() {
			// subDir := dirpath + `/` + file.Name()
			subDir := fmt.Sprintf("%s/%s", dirpath, file.Name())
			// dir2, _ := ioutil.ReadDir(subDir)
			// fmt.Println(subDir, dir2)
			delDirFile(subDir)
			// continue
		}
		//文件的最后修改时间
		tdate := file.ModTime()
		ft := int32(tdate.Unix())

		//3天前的文件就删除
		if ft < ct-spt {
			os.Remove(dirpath + "/" + file.Name())
			// fmt.Println("del file:", dirpath+"/"+file.Name())
		}
	}
}

func delDirTask() {
	for {
		//读取目录配置文件信息
		delDirFile(`cache/`)

		//休眠1小时
		time.Sleep(time.Hour)
		// time.Sleep(3600e9)
	}
}

func main() {
	// go delDirTask()
	e := echo.New()

	s := NewStats()
	e.Use(s.Process)

	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK,
			`https://open.readfollow.com build by golang, author yizenghui.  use the api must jwt token
/api/list/:url?drive=
/api/catalog/:url?drive=
/api/info/:url?drive=
drive sup: qidian,zongheng,17k,luoqiu,booktxt,bxwx,uxiaoshuo,soe8,manhwa,r2hm,xbiquge,biquyun
`)
	})

	// 获取用户签名
	e.GET("/gettoken", a.GetToken)
	e.GET("/getapitoken", a.GetAPIToken)
	e.GET("/getapitoken2", a.GetAPIToken2)
	// 解密数据内容(保存数据到库)
	e.GET("/crypt", a.Crypt)

	// 获取二维码(图片资源)
	e.GET("/qrcode", a.GetQrcode)

	// 从二维码进来跳哪里去
	e.GET("/qrcodejump", a.GetQrcodeWxto)

	// Restricted group
	api := e.Group("/api")

	api.GET("/stats", s.Handle) // Endpoint to get stats

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &a.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	api.Use(middleware.JWTWithConfig(config))

	api.GET("/checkopenid", a.CheckOpenID)
	// r.Use(middleware.JWT([]byte("secret")))

	// //  粉丝关注列表
	// api.GET("/follows", a.GetUserFollows)

	// // 粉丝添加关注
	// api.POST("/follows", a.CreateUserFollow)

	//  粉丝自定义源
	// api.GET("/sources", a.GetUserSources)

	api.POST("/urlencode", func(c echo.Context) (err error) {
		type Data struct {
			URL string `json:"url" validate:"required"`
		}
		u := new(Data)
		if err = c.Bind(u); err != nil {
			return
		}
		// url := c.FormValue("url")
		return c.JSON(http.StatusOK, grab.EncodeURL(u.URL))
	})

	api.POST("/getdrive", func(c echo.Context) (err error) {
		type Data struct {
			URL string `json:"url" validate:"required"`
		}
		u := new(Data)
		if err = c.Bind(u); err != nil {
			return
		}
		type Ret struct {
			Key   string `json:"key"`
			Drive string `json:"drive"`
			Page  string `json:"page"`
		}
		ret := Ret{
			grab.EncodeURL(u.URL),
			``,
			``,
		}

		address, drive, page := grab.ExplainLink(u.URL)
		if address != `` && drive != `` && page != `` { // 有解释到什么是资源就返回解释到的资源
			ret.Key = grab.EncodeURL(address)
			ret.Drive = drive
			ret.Page = page
		}

		// url := c.FormValue("url")
		return c.JSON(http.StatusOK, ret)
	})

	e.GET("/urlencode", func(c echo.Context) error {
		url := c.QueryParam("url")
		return c.JSON(http.StatusOK, grab.EncodeURL(url))
	})

	// 粉丝添加关注
	// api.POST("/sources", a.CreateUserSource)

	// 搜索
	api.GET("/search", a.SearchPosts)

	//首页走马灯数据
	api.GET("/carousels", a.GetCarousels)
	api.GET("/userlinks", a.GetUserLinks)
	api.GET("/explorelinks", a.GetExploreLinks)
	api.GET("/newcateloghelps", a.GetNewCatelogLinks)
	api.GET("/allroesoures", a.GetAllResources)
	api.GET("/allbookroesoures", a.GetAllBookResources)
	api.GET("/allcartoonroesoures", a.GetAllCartoonResources)
	api.GET("/alllearnresources", a.GetAllLearnResources)

	api.GET("/newactivity", a.NewActivity)  //新号召令
	api.GET("/activities", a.GetActivities) //所有号召令(100条)

	//  自定义分类
	api.GET("/classify", func(c echo.Context) error {
		// list := grab.GetClassify()
		cf := cpi.GetConf()

		version := c.QueryParam("version")

		if cf.Search.LimitLevel || version == cf.Search.DevVersion { // 开启严格检查
			list := grab.GetWaitExamineClassify()
			return c.JSON(http.StatusOK, list)
		}
		list := grab.GetResources()
		//
		return c.JSON(http.StatusOK, list)
	})
	//  自定义资源列表
	api.GET("/resources", func(c echo.Context) error {
		list := grab.GetResources()
		return c.JSON(http.StatusOK, list)
	})

	//  自定义专题列表
	api.GET("/topics", func(c echo.Context) error {
		list := grab.GetTopics()
		return c.JSON(http.StatusOK, list)
	})

	//  获取第三方资源的分类(按驱动)
	api.GET("/categories/:url", func(c echo.Context) error {
		urlStr, _ := grab.DecodeURL(c.Param("url"))
		guide := grab.GetGuide(c.QueryParam("drive"))
		list, _ := guide.GetCategories(urlStr)
		return c.JSON(http.StatusOK, list)
	})
	//  在第三方平台进行资源搜索(按驱动)
	api.GET("/searchmore", func(c echo.Context) error {
		keyword := c.QueryParam("keyword") //grab.DecodeURL()
		drive := c.QueryParam("drive")
		guide := grab.GetGuide(drive)
		list, _ := guide.Search(keyword)

		if drive == `qidian` || drive == `zongheng` || drive == `jx` || drive == `paoshu8` || drive == `xs280` || drive == `hongxiu` || drive == `xxsy` || drive == `biqugeinfo` || drive == `mcmssc` || drive == `17k` || drive == `xbiquge` || drive == `luoqiu` || drive == `booktxt` || drive == `bxwx` || drive == `uxiaoshuo` || drive == `biquyun` || drive == `soe8` {
			go a.SyncPosts(list, 1)
		} else if drive == `manhwa` || drive == `kanmeizi` || drive == `haimaoba` || drive == `hanmanwo` || drive == `hanmanku` || drive == `ssmh` || drive == `fuman` || drive == `aimeizi5` {
			go a.SyncPosts(list, 2)
		}
		return c.JSON(http.StatusOK, list)
	})
	// 组装第三方平台搜索所需链接
	api.GET("/allsearchdrives", a.SearchMoreAction)

	//  获取可订阅目录列表
	api.GET("/list/:url", func(c echo.Context) error {
		urlStr, _ := grab.DecodeURL(c.Param("url"))
		drive := c.QueryParam("drive")
		guide := grab.GetGuide(drive)
		list, _ := guide.GetList(urlStr)

		if drive == `qidian` || drive == `zongheng` || drive == `jx` || drive == `paoshu8` || drive == `shuge` || drive == `qkshu6` || drive == `xs280` || drive == `hongxiu` || drive == `xxsy` || drive == `biqugeinfo` || drive == `mcmssc` || drive == `17k` || drive == `xbiquge` || drive == `luoqiu` || drive == `booktxt` || drive == `bxks` || drive == `xin18` || drive == `bxwx` || drive == `uxiaoshuo` || drive == `biquyun` || drive == `soe8` {
			go a.SyncPosts(list, 1)
		} else if drive == `manhwa` || drive == `kanmeizi` || drive == `haimaoba` || drive == `hanmanwo` || drive == `hanmanku` || drive == `ssmh` || drive == `fuman` || drive == `aimeizi5` {
			go a.SyncPosts(list, 2)
		}

		return c.JSON(http.StatusOK, list)
	})

	// 获取目录(订阅目录内容)
	api.GET("/catalog/:url", func(c echo.Context) error {
		urlStr, _ := grab.DecodeURL(c.Param("url"))
		drive := c.QueryParam("drive")
		reader := grab.GetReader(drive)
		list, _ := reader.GetCatalog(urlStr)
		return c.JSON(http.StatusOK, list)
	})

	// 获取页面正文内容
	api.GET("/info/:url", func(c echo.Context) error {
		urlStr, _ := grab.DecodeURL(c.Param("url"))
		drive := c.QueryParam("drive")
		reader := grab.GetReader(drive)
		list, _ := reader.GetInfo(urlStr)
		return c.JSON(http.StatusOK, list)
	})
	//  get book demo
	api.GET("/about", func(c echo.Context) error {
		list := grab.GetAbouts()
		return c.JSON(http.StatusOK, list)
	})
	//  get book demo
	api.GET("/help", func(c echo.Context) error {
		list := grab.GetHelps()
		return c.JSON(http.StatusOK, list)
	})
	//  get book demo
	api.GET("/alldrive", func(c echo.Context) error {
		type Item struct {
			Name  string `json:"name" `
			Drive string `json:"drive" `
			Page  string `json:"page"`
		}
		var list = []Item{
			Item{`小说目录`, `text`, `/pages/catalog`},
			Item{`小说单章`, `text`, `/pages/book`},
			Item{`文章目录`, `article`, `/pages/catalog`},
			Item{`文章详情`, `article`, `/pages/article`},
		}
		return c.JSON(http.StatusOK, list)
	})

	// 图标
	e.File("favicon.ico", "images/favicon.ico")
	e.Logger.Fatal(e.Start(":8041"))
	// e.Logger.Fatal(e.StartAutoTLS(":443"))

}

type (
	//Stats  struct
	Stats struct {
		Uptime       time.Time      `json:"uptime"`
		RequestCount uint64         `json:"requestCount"`
		Statuses     map[string]int `json:"statuses"`
		mutex        sync.RWMutex
	}
)

//NewStats create new stats
func NewStats() *Stats {
	return &Stats{
		Uptime:   time.Now(),
		Statuses: make(map[string]int),
	}
}

// Process is the middleware function.
func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.RequestCount++
		status := strconv.Itoa(c.Response().Status)
		s.Statuses[status]++
		return nil
	}
}

// Handle is the endpoint to get stats.
func (s *Stats) Handle(c echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return c.JSON(http.StatusOK, s)
}
