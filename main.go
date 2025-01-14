package main

import (
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {

	bin := ""
	// print args
	// print all args

	if len(os.Args) == 4 {
		bin = os.Args[3]
	} else {
		bin = "/usr/bin/chromium-browser"
	}
	u := launcher.New().Bin(bin).
		Headless(true).NoSandbox(false).
		Leakless(false).MustLaunch()

	rod.New().ControlURL(u).MustConnect().MustPage(os.Args[1]).
		MustWaitLoad().MustPDF(os.Args[2])

	// clean up
	rod.New().ControlURL(u).MustConnect().MustClose()

	return
}
