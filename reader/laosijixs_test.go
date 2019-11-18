package reader

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

func Test_LaosijixsGetInfo(t *testing.T) {
	urlStr := `http://m.laosijixs.com/20/20961/546056.html`
	urlStr = `http://m.laosijixs.com/20/20961/546056_5.html`
	// urlStr = `http://m.laosijixs.com/20/20961/546056_6.html`
	// urlStr = `https://m.35xs.com/book/237551/51896850.html`
	reader := LaosijixsReader{}
	list, err := reader.GetInfo(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}
func Test_LaosijixsGetCatalog(t *testing.T) {
	urlStr := `http://m.laosijixs.com/80/80896/`
	urlStr = `http://m.laosijixs.com/79/79531/`
	reader := LaosijixsReader{}
	list, err := reader.GetCatalog(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_LaosijixsGetList(t *testing.T) {
	urlStr := `https://m.booktxt.net/wapsort/1_1.html`
	urlStr = `http://m.laosijixs.com/shuku/`
	reader := LaosijixsReader{}
	list, err := reader.GetList(urlStr)
	if err != nil {

	}
	t.Fatal(list)
}

func Test_LaosijixsGetInfoBoby(t *testing.T) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var title, res, res2 string
	var jres []string
	// var res2 []string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`http://m.laosijixs.com/20/20961/546056_5.html`),
		// chromedp.Reload(),
		// chromedp.WaitVisible("#content"),
		// chromedp.Title(&title),
		chromedp.Sleep(time.Second*6),
		// chromedp.Body(`html`, &res, chromedp.NodeVisible, chromedp.ByQuery),
		// chromedp.Evaluate(`function() {$('#content').find('span').remove();return { body: $('#content').innerText};}`, &jres),
		chromedp.Text(`#content`, &res2, chromedp.NodeVisible, chromedp.ByID),
		// chromedp.OuterHTML("#content", &res),
		// chromedp.OuterHTML(`#content`, &res, chromedp.NodeVisible, chromedp.ByQuery),
	)
	if err != nil {
		t.Fatal(err)
	}

	// err = chromedp.Run(ctx,
	// 	chromedp.Sleep(time.Second*3),
	// 	chromedp.Evaluate(`function() {$('#content').find('span').remove();return { body: $('#content').innerText};}`, &jres),
	// 	chromedp.OuterHTML("#html", &res),
	// ) //
	// chromedp.Body(`html`, &res, chromedp.NodeVisible, chromedp.ByQuery),

	if err != nil {
		t.Fatal(err)
	}
	// log.Println(strings.TrimSpace(res))
	t.Fatal(title, jres, strings.TrimSpace(res2), strings.TrimSpace(res))
}

func writeHTML(content string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, strings.TrimSpace(content))
	})
}
func Test_LaosijixsGetInfoClick(t *testing.T) {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ts := httptest.NewServer(writeHTML(`
<body>
<p id="content" onclick="changeText()">Original content.</p>
<script>
function changeText() {
	document.getElementById("content").textContent = "New content!"
}
</script>
</body>
	`))
	defer ts.Close()

	var outerBefore, outerAfter string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(ts.URL),
		chromedp.OuterHTML("#content", &outerBefore),
		chromedp.Click("#content", chromedp.ByID),
		chromedp.OuterHTML("#content", &outerAfter),
	); err != nil {
		panic(err)
	}
	t.Fatal("OuterHTML before clicking:", outerBefore, "OuterHTML after clicking:", outerAfter)

}

func Test_LaosijixsGetInfoJS(t *testing.T) {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ts := httptest.NewServer(writeHTML(`<!doctype html>
<html>
<body>
  <div id="content">the content</div>
</body>
</html>`))
	defer ts.Close()

	const expr = `(function(d, id, v) {
		var b = d.querySelector('body');
		var el = d.createElement('div');
		el.id = id;
		el.innerText = v;
		b.insertBefore(el, b.childNodes[0]);
	})(document, %q, %q);`

	var nodes []*cdp.Node
	if err := chromedp.Run(ctx,
		chromedp.Navigate(ts.URL),
		chromedp.Nodes(`document`, &nodes, chromedp.ByJSPath),
		chromedp.WaitVisible(`#content`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			s := fmt.Sprintf(expr, "thing", "a new thing!")
			_, exp, err := runtime.Evaluate(s).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.WaitVisible(`#thing`),
	); err != nil {
		panic(err)
	}
	t.Fatal(nodes[0].Dump("  ", "  ", false))

}
