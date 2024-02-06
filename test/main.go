package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"time"
)

func main() {

	//utils.DownloadDemoHTML("https://www.hao6v.tv/dy/2024-02-02/43708.html")
	//return
	//c := colly.NewCollector()
	//f, _ := os.Open("./test.html")
	// 目标页面URL，需要替换成你的实际页面
	//targetURL := "https://www.hao6v.tv/dy/2024-02-02/43708.html"
	//targetURL := "https://www.jd.com"
	targetURL := "https://www.hao6v.tv/dy/2024-02-04/43731.html"
	//targetURL := "https://www.baidu.com"
	// 输出文件路径
	outputFilePath := "./chromedp_demo.html"

	// 创建Chrome实例
	options := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true), // 设置为无头模式
		// 打开开发者工具，便于调试
		chromedp.Flag("auto-open-devtools-for-tabs", true),
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// 设置超时时间
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// 运行任务
	//var iframeContent string
	var htmlContent string
	var buf []byte

	var (
		height int64
	)
	err := chromedp.Run(ctx,
		chromedp.Navigate(targetURL),
		chromedp.WaitReady("body"),
		chromedp.EvaluateAsDevTools(`document.body.scrollHeight`, &height),
	)

	var res string
	chromedp.Run(ctx,
		// 根据页面尺寸设置视图区域大小
		chromedp.EmulateViewport(1920, height),
		chromedp.EvaluateAsDevTools(`
var xpath = "//strong/span/span[contains(text(), '【下载地址】')]";
var result = document.evaluate(xpath, document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue;
result ? result.innerText : '';
`, &res),
		//chromedp.CaptureScreenshot(&buf),
		chromedp.OuterHTML(`body`, &htmlContent, chromedp.ByQuery),
	)
	fmt.Println(res)

	if err != nil {
		log.Fatalf("Failed to execute chromedp tasks: %v", err)
	}

	os.WriteFile("example.png", buf, 0644)
	// 将获取的内容写入文件
	err = os.WriteFile("./htmlContent_demo.html", []byte(htmlContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write iframe content to file: %v", err)
	}

	log.Printf("Iframe content successfully written to %s\n", outputFilePath)
}
