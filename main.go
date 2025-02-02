package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	router := gin.Default()
	bin := "/usr/bin/chromium"
	// bin = "C:/Users/dell/AppData/Local/Google/Chrome/Application/chrome.exe"

	fmt.Println("Creating browser")
	u := launcher.New().Bin(bin).
		Headless(true).NoSandbox(false).
		Leakless(false)

	fmt.Println("Launching browser")
	defer u.Cleanup()
	launcher := u.MustLaunch()
	fmt.Println("Creating page")
	browser := rod.New().ControlURL(launcher).Trace(true).SlowMotion(2 * time.Second).MustConnect()
	defer browser.Close()

	router.GET("/pdf", func(ctx *gin.Context) {
		url := ctx.Query("url")
		path := "result/pdf_" + strconv.Itoa(rand.Int()) + ".pdf"
		generatePdf(*browser, url, path)
		ctx.File(path)
	})

	router.Run("localhost:3030")
}

func generatePdf(browser rod.Browser, url string, path string) {
	fmt.Println("start page")
	page := browser.MustPage(url)
	fmt.Println("defer page close")
	defer page.Close()
	fmt.Println("wait load")
	page.MustWaitLoad()
	fmt.Println("pdf")
	page.MustPDF(path)
}
