package main

import (
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	router := gin.Default()
	router.GET("/pdf", func(ctx *gin.Context) {
		url := ctx.Query("url")

		path := "result/pdf_" + strconv.Itoa(rand.Int()) + ".pdf"
		generatePdf(url, path)

		ctx.File(path)
	})

	router.Run("localhost:3000")
}

func generatePdf(url string, path string) {
	bin := "C:/Users/dell/AppData/Local/Google/Chrome/Application/chrome.exe"

	u := launcher.New().Bin(bin).
		Headless(true).NoSandbox(false).
		Set("--database-path", "/tmp/rod").
		Leakless(false).MustLaunch()

	page := rod.New().ControlURL(u).MustConnect().MustPage(url).
		MustWaitLoad()

	page.MustPDF(path)
	// clean up
	page.MustClose()
}
