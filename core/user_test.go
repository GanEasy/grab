package core

import (
	"strconv"
	"testing"

	"github.com/GanEasy/grab/db"
)

func Test_GetOpenID(t *testing.T) {
	u2, _ := GetOpenID("00397Gwx0b39sj1Pm3wx0guNwx097GwP")
	t.Fatal(u2)
}

// func Test_SendPostUpdateMSG(t *testing.T) {
// 	u2 := SendPostUpdateMSG("oe9Hq0GwS4umXENTCn4lJgxaNVog", "af1a0cdf6adfbf4030358fc2b4264d24", "tttt", "")
// 	t.Fatal(u2)
// }

func Test_GetwxCodeUnlimit(t *testing.T) {
	u2, err := GetwxCodeUnlimit("123", "")
	t.Fatal(u2)
	t.Fatal(err)
}

func Test_MSGSecCHECK(t *testing.T) {
	err := MSGSecCHECK(`李克强强奸性爱游戏`) //这个是违禁词
	// err := MSGSecCHECK(`点道`) //
	t.Fatal(err)
}

func Test_GetToken(t *testing.T) {

	token, _ := TokenServe.Token()
	t.Fatal(token)
}

func Test_GetFromid(t *testing.T) {

	fromid, _ := strconv.Atoi(`-1`)
	t.Fatal(uint(fromid))
}

func Test_GetFansByID(t *testing.T) {
	var fans db.Fans
	fromid, _ := strconv.Atoi(`1`)
	fans.GetFansByID(uint(fromid))
	t.Fatal(fans)
}
