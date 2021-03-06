package core

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/GanEasy/grab/db"
	"github.com/GanEasy/grab/reader"
	wxbizdatacrypt "github.com/yilee/wx-biz-data-crypt"
)

//TokenServe token 服务器
var TokenServe *DefaultAccessTokenServer

func init() {

	TokenServe = NewDefaultAccessTokenServer(config.ReaderMinApp.AppID, config.ReaderMinApp.AppSecret)

}

// OpenIDData 开放数据 openID
type OpenIDData struct {
	ErrCode    int64  `json:"errcode"`
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

//GetOpenID 获取微信小程序上报的openid 此ID暂不加密处理
func GetOpenID(code string) (OpenIDData, error) {
	//
	type Ret struct {
		ErrCode    int64  `json:"errcode"`
		ErrMSG     string `json:"errmsg"`
		SessionKey string `json:"session_key"`
		ExpiresIn  int64  `json:"expires_in"`
		OpenID     string `json:"openid"`
	}
	var ret Ret

	url := fmt.Sprintf(`https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code`,
		config.ReaderMinApp.AppID,
		config.ReaderMinApp.AppSecret,
		code,
	)

	HTTPGetJSON(url, &ret)
	var err error

	if ret.ErrCode != 0 {
		err = errors.New(string(ret.ErrCode))
	}

	return OpenIDData{ret.ErrCode, ret.OpenID, ret.SessionKey}, err
}

// SubmitPage 提交页面数据
type SubmitPage struct {
	Path  string `json:"path"`  //
	Query string `json:"query"` //
}

//WxAppSubmitPage 提交页面单个页面（望收录）
func WxAppSubmitPage(wxto string) error {
	type Data struct {
		Pages []SubmitPage `json:"pages"`
	}
	//
	type Ret struct {
		ErrCode int64  `json:"errcode"` //errCode
		ErrMSG  string `json:"errmsg"`  //errMsg
	}
	var ret Ret

	var data = Data{}

	link, err3 := url.Parse(wxto)

	if err3 != nil {
		return err3
	}
	data.Pages = append(data.Pages,
		SubmitPage{
			Path:  link.Path,
			Query: link.RawQuery,
		})
	token, err2 := TokenServe.Token()
	if err2 != nil {
		return err2
	}
	// token = `27_EFpACLm1qpGcK8p_xEnZPnowJGKKEfWzy7500PLAR7Ek-8UaooSW-HTteSCfM2_r2f3zkKTcCgLFYvE094UNzXhZyv3KbZqAk_D8USQGFeYqklXrC6UVBIZfO0oAI2yB63nI0-cAsHjksNcAOPNjAEACDB`
	url := fmt.Sprintf(`https://api.weixin.qq.com/wxa/search/wxaapi_submitpages?access_token=%v`, token)

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	HTTPPostJSON(url, b, &ret)

	// log.Println(`xx`, ret)
	if ret.ErrCode != 0 {
		// err = errors.New(string(ret.ErrMSG))
		err = errors.New(strconv.FormatInt(ret.ErrCode, 10))
	}

	return err
}

//WxAppSubmitPages 提交页面（望收录）
func WxAppSubmitPages(list reader.Catalog) (err error) {
	type Data struct {
		Pages []SubmitPage `json:"pages"`
	}
	//
	type Ret struct {
		ErrCode int64  `json:"errcode"` //errCode
		ErrMSG  string `json:"errmsg"`  //errMsg
	}
	var ret Ret

	var data = Data{}

	if len(list.Cards) > 0 {
		var link *url.URL
		for _, v := range list.Cards {

			link, err = url.Parse(v.WxTo)

			if err != nil {
				return err
			}
			data.Pages = append(data.Pages,
				SubmitPage{
					Path:  link.Path,
					Query: link.RawQuery,
				})
			// fmt.Println(link.RawQuery)
		}

		token, err := TokenServe.Token()
		if err != nil {
			return err
		}

		// token = `27_EFpACLm1qpGcK8p_xEnZPnowJGKKEfWzy7500PLAR7Ek-8UaooSW-HTteSCfM2_r2f3zkKTcCgLFYvE094UNzXhZyv3KbZqAk_D8USQGFeYqklXrC6UVBIZfO0oAI2yB63nI0-cAsHjksNcAOPNjAEACDB`
		url := fmt.Sprintf(`https://api.weixin.qq.com/wxa/search/wxaapi_submitpages?access_token=%v`, token)

		b, err := json.Marshal(data)
		if err != nil {
			return err
		}

		HTTPPostJSON(url, b, &ret)

		// log.Println(`xx`, ret)
		if ret.ErrCode != 0 {
			// err = errors.New(string(ret.ErrMSG))
			err = errors.New(strconv.FormatInt(ret.ErrCode, 10))
		}

		return err
	}
	return nil
}

// //SendPostUpdateMSG 发送更新通知
// func SendPostUpdateMSG(openID, formID, title, page string) error {
// 	//
// 	type Ret struct {
// 		ErrCode int64  `json:"errcode"`
// 		ErrMSG  string `json:"errmsg"`
// 	}
// 	var ret Ret

// 	type DataItem struct {
// 		Value string `json:"value"`
// 		Color string `json:"color"`
// 	}

// 	type Template struct {
// 		Touser          string      `json:"touser"`
// 		TemplateID      string      `json:"template_id"`
// 		Page            string      `json:"page"`
// 		FormID          string      `json:"form_id"`
// 		Data            interface{} `json:"data"`
// 		EmphasisKeyword string      `json:"emphasis_keyword"`
// 	}

// 	//MSG 关注通知消息结构
// 	type MSG struct {
// 		Title    template.DataItem `json:"keyword1"`
// 		CATEGORY template.DataItem `json:"keyword2"`
// 	}

// 	data := Template{
// 		Touser:     openID,
// 		TemplateID: "QEhBZIivAI5x0hbWEp4IqMKAb3RhLXCl3eBr1GC_7FE",
// 		Page:       page,
// 		FormID:     formID,
// 		Data: MSG{
// 			Title:    template.DataItem{Value: title, Color: ""},
// 			CATEGORY: template.DataItem{Value: "文章", Color: ""},
// 		},
// 		EmphasisKeyword: "",
// 	}

// 	token, err2 := TokenServe.Token()
// 	if err2 != nil {

// 		return err2
// 	}
// 	url := fmt.Sprintf(`https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=%v`, token)

// 	b, err := json.Marshal(data)
// 	if err != nil {
// 		return err
// 	}

// 	HTTPPostJSON(url, b, &ret)

// 	if ret.ErrCode != 0 {
// 		err = errors.New(string(ret.ErrCode))
// 	}

// 	return err
// }

//MSGSecCHECK 文本安全检查
func MSGSecCHECK(text string) error {
	//
	type Ret struct {
		ErrCode int64  `json:"errcode"`
		ErrMSG  string `json:"errmsg"`
	}
	var ret Ret

	type Data struct {
		Content string `json:"content"`
	}

	var data = Data{text}
	token, err2 := TokenServe.Token()
	if err2 != nil {
		return err2
	}
	// token = `27_EFpACLm1qpGcK8p_xEnZPnowJGKKEfWzy7500PLAR7Ek-8UaooSW-HTteSCfM2_r2f3zkKTcCgLFYvE094UNzXhZyv3KbZqAk_D8USQGFeYqklXrC6UVBIZfO0oAI2yB63nI0-cAsHjksNcAOPNjAEACDB`
	url := fmt.Sprintf(`https://api.weixin.qq.com/wxa/msg_sec_check?access_token=%v`, token)

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	HTTPPostJSON(url, b, &ret)

	if ret.ErrCode != 0 {
		// err = errors.New(string(ret.ErrMSG))
		err = errors.New(strconv.FormatInt(ret.ErrCode, 10))
	}

	return err
}

//GetwxCodeUnlimit 获取无数量限制的微信二维码
func GetwxCodeUnlimit(scene, page string) (file string, err error) {

	// name := GetMd5String(fmt.Sprintf(`%v%v`, scene, page))
	file = fmt.Sprintf(`file/%v.jpg`, scene)

	_, err2 := os.Stat(file)
	if os.IsNotExist(err2) {

		type Template struct {
			Scene     string      `json:"scene"`
			Page      string      `json:"page"`
			Width     int         `json:"width"`
			AutoColor bool        `json:"auto_color"`
			LineColor interface{} `json:"line_color"`
		}

		data := Template{
			Scene: scene,
			Page:  page,
			Width: 280,
		}

		token, err2 := TokenServe.Token()
		if err2 != nil {
			return "", err2
		}
		url := fmt.Sprintf(`https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%v`, token)

		b, err := json.Marshal(data)
		if err != nil {
			return "", err
		}
		_, err = SaveQrcodeImg(url, file, b)
	}
	return file, err
}

//GetMd5String 获取MD5加密字符串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// SaveQrcodeImg 保存图片到本地
func SaveQrcodeImg(imageURL, saveName string, body []byte) (n int64, err error) {
	out, err := os.Create(saveName)
	defer out.Close()
	if err != nil {
		return
	}
	// text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
	// application/json; charset=utf-8
	//
	resp, err := httpClient.Post(imageURL, "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8", bytes.NewReader(body))

	// resp, err := httpClient.Post(imageURL, "application/json; charset=utf-8", bytes.NewReader(body))
	// resp, err := httpClient.Post(imageURL, "application/json;q=0.9,image/webp,*/*;q=0.8", bytes.NewReader(body))

	if err != nil {
		return
	}
	pix, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	if err != nil {
		return
	}
	//errcode errmsg {"errcode":40001,"errmsg":"invalid credential, access_token is invalid or not latest hint: [blkOua0570b105!]"}
	type ErrRet struct {
		Errcode int  `json:"errcode"`
		Errmsg  bool `json:"errmsg"`
	}
	var errret ErrRet
	err = json.Unmarshal(pix, errret)
	if err == nil {
		log.Println(`build wechat minapp qrcode error`, errret.Errcode, errret.Errmsg)
		return
	}

	if len(pix) > 2048 {
		n, err = io.Copy(out, bytes.NewReader(pix))

		if err != nil {
			return
		}
	}
	// todo 获取图片类型
	// fmt.Println(resp.Header.Get("Content-Type"))
	return 0, errors.New(`build wechat minapp qrcode error`)
}

// GetCryptData 解密数据
func GetCryptData(sessionKey, encryptedData, iv string) (*db.Fans, error) {

	// log.Println(config.ReaderMinApp.AppID, sessionKey, encryptedData, iv)
	pc := wxbizdatacrypt.NewWXBizDataCrypt(config.ReaderMinApp.AppID, sessionKey)
	userInfo, err := pc.Decrypt(encryptedData, iv)
	log.Println(err)
	if err != nil {
		return &db.Fans{}, err
	}
	fans, err := GetFansByOpenID(userInfo.OpenID)
	if err != nil {
		return &db.Fans{}, err
	}
	if fans.SessionKey != sessionKey {
		fans.OpenID = userInfo.OpenID
		// fans.UnionID = userInfo.UnionID
		fans.NickName = userInfo.NickName
		fans.Gender = userInfo.Gender
		fans.City = userInfo.City
		fans.Province = userInfo.Province
		fans.Country = userInfo.Country
		fans.AvatarURL = userInfo.AvatarURL
		fans.Language = userInfo.Language
		fans.Timestamp = userInfo.Watermark.Timestamp
		fans.AppID = userInfo.Watermark.AppID
		fans.SessionKey = sessionKey //
		fans.Save()
	}
	return fans, err
}

// GetFansByOpenID 解密数据
func GetFansByOpenID(openID string) (*db.Fans, error) {
	var err error
	var fans db.Fans
	if openID != "" {
		fans.GetFansByOpenID(openID)
	} else {
		err = errors.New(string(`openID is empty!!!`))
	}
	return &fans, err
}

// SyncPost 保存数据
func SyncPost(name, wxto, from string, cate int32) {
	var post db.Post
	if wxto != "" {
		post.GetPostByWxto(wxto)
		post.Cate = cate
		post.Name = name
		post.From = from
		post.Level = cate + 1
		post.Save()
	}
}

// GetPostsByName 搜索post
func GetPostsByName(name string) (posts []db.Post) {
	var post db.Post
	posts = post.GetPostsByName(name)
	return posts
}

// GetPostsByNameLimitLevel 搜索post
func GetPostsByNameLimitLevel(name string, level int) (posts []db.Post) {
	var post db.Post
	posts = post.GetPostsByNameLimitLevel(name, level)
	return posts
}

// AddActivity 新增号召
func AddActivity(title, wxto string) (activity db.Activity) {
	if wxto != `` {
		// 修正已推荐数据使用不同解释器问题
		url, err := reader.GetURLStringParam(wxto, `url`)
		if err == nil {
			activity.GetActivityByResource(url)
			activity.Title = title
			activity.Resource = url
			activity.WxTo = wxto
			activity.Level = reader.GetPathLevel(wxto)
			activity.Total++
			activity.Save()
		} else {
			activity.GetActivityByWxto(wxto)
			activity.Title = title
			activity.WxTo = wxto
			activity.Level = reader.GetPathLevel(wxto)
			activity.Total++
			activity.Save()
		}
	}
	return activity
}

// GetActivities 搜索post
func GetActivities() (activities []db.Activity) {
	var activity db.Activity
	activities = activity.GetActivities()
	// for _, v := range activities {
	// 	if v.Resource == `` {
	// 		url, err := reader.GetURLStringParam(v.WxTo, `url`)
	// 		if err == nil {
	// 			v.Resource = url
	// 			v.Save()
	// 		}
	// 	}
	// }
	return activities
}
