// Command text is a chromedp example demonstrating how to extract text from a
// specific element.
package main

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res, res1 string
	var res2 []string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`http://m.laosijixs.com/20/20961/546056_5.html`),
		chromedp.Sleep(time.Second*2),
		//$('#content').find('span').remove();
		chromedp.Evaluate(`Object.keys(window);`, &res2),
		// chromedp.Body(`html`, &res),

		// chromedp.Text(`html`, &res1, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Text(`#content`, &res, chromedp.NodeVisible, chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(strings.TrimSpace(res))
	log.Println(res1)

	// chromedp.Navigate(`https://github.com/chromedp/examples`),
	// chromedp.Text(`.Box-body`, &res, chromedp.NodeVisible, chromedp.ByQuery),
	// chromedp.Navigate(`https://m.138txt.com/193/193028/`),
	// chromedp.Text(`body`, &res, chromedp.NodeVisible, chromedp.ByQuery),
	// chromedp.OuterHTML(`body`, &res, chromedp.NodeVisible, chromedp.ByQuery),
	// chromedp.Reload(),
	// chromedp.WaitVisible(`#content`, chromedp.ByID),
	// chromedp.WaitVisible("#content"),
	// chromedp.Text(`html`, &res, chromedp.NodeVisible, chromedp.ByQuery),
	// $('#content').find('span').remove();
	// chromedp.Evaluate(`$('#content').find('span').remove();`, &res2),
	// chromedp.Evaluate(`$('#content').find('span').remove();`, &res2),
	// chromedp.Text(`html`, &res, chromedp.NodeVisible, chromedp.ByQuery),
	// chromedp.TextContent(`#content`, &res2, chromedp.NodeVisible, chromedp.ByID),
	// chromedp.OuterHTML("html", &res),
	// chromedp.Navigate(`http://m.laosijixs.com/20/20961/546056_5.html`),
	// chromedp.Text(`#content`, &res, chromedp.NodeVisible, chromedp.ByID),
	// chromedp.Text(`#pkg-overview`, &res, chromedp.NodeVisible, chromedp.ByID),

}

/*

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

*/
