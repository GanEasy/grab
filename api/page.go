package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// Carousel 小程序首页轮播内容(作为专题广告或其它的东西使用)
type Carousel struct {
	URL  string `json:"url"`
	Type string `json:"type"` // 期望可以同时支持视频播放(虽然很不现实)
	WxTo string `json:"wxto"` // 点击后跳转地址
}

// GetCarousels 获取首页走马灯数据
func GetCarousels(c echo.Context) error {
	var carousels []Carousel
	carousels = append(
		carousels,
		Carousel{
			URL:  `https://ossweb-img.qq.com/images/lol/web201310/skin/big84000.jpg`,
			Type: `image`,
			WxTo: ``,
		})

	carousels = append(
		carousels,
		Carousel{
			URL:  `https://ossweb-img.qq.com/images/lol/web201310/skin/big37006.jpg`,
			Type: `image`,
			WxTo: ``,
		})

	return c.JSON(http.StatusOK, carousels)
}
