package api

import (
	cpi "github.com/GanEasy/grab/core"
	"github.com/GanEasy/grab/reader"
)

// SyncPosts 添加源
func SyncPosts(list reader.Catalog) {
	if len(list.Cards) > 0 {
		for _, v := range list.Cards {
			// fmt.Println(`tv`, v, list.Title)
			cpi.SyncPost(v.Title, v.WxTo, v.From, 1)
		}
	}
}
