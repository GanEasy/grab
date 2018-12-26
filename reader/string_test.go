package reader

import (
	"strconv"
	"testing"
)

func Test_FindString(t *testing.T) {

	// html := `{"I":"5333","V":"马经理"},`
	// Linkman := FindString(`{"I":"5333","V":"(?P<value>[^"]+)"}`, html, "value")
	// t.Fatal(Linkman)
	html := `http://book.zongheng.com/store/c0/c0/b0/u0/p1/v9/s9/t0/u0/i0/ALL.html`
	page := FindString(`/p(?P<page>(\d)+)/`, html, "page")
	p, err := strconv.Atoi(page)
	// page := FindString(`zong(?P<page>(\w)+)`, html, "page")

	t.Fatal(`FatalFatal`, page, p, err)
}
