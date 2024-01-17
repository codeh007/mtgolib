package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	l := launcher.New().RemoteDebuggingPort(9222).Headless(true).NoSandbox(true)

	// For more info: https://pkg.go.dev/github.com/go-rod/rod/lib/launcher
	u := l.MustLaunch()

	// browser := rod.New().ControlURL(u).MustConnect()

	// page := browser.MustPage("http://example.com").MustWaitStable()

	// fmt.Println(page.MustInfo().Title)

	// 范例： 连接现有 chrome 实例
	// u := launcher.MustResolveURL("172.17.0.1:9222")

	browser := rod.New().ControlURL(u).MustConnect()

	fmt.Println(
		browser.MustPage("https://www.bing.com/").MustEval("() => document.title"),
	)
	// browser.Sleeper()
	fmt.Println("休眠等等")
	time.Sleep(time.Second * 30000)
	fmt.Println("结束")
}
