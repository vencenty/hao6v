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

	//c := colly.NewCollector()
	//f, _ := os.Open("./test.html")
	// 目标页面URL，需要替换成你的实际页面
	targetURL := "https://www.hao6v.tv/dy/2024-02-02/43708.html"
	// 输出文件路径
	outputFilePath := "./iframe_content.html"

	// 创建Chrome实例
	options := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true), // 设置为无头模式
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// 设置超时时间
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// 运行任务
	var iframeContent string
	var htmlContent string

	err := chromedp.Run(ctx,
		chromedp.Navigate(targetURL),
		chromedp.WaitReady(`iframe#ifc`, chromedp.ByQuery), // 等待iframe加载完成
		chromedp.Sleep(1*time.Second),
		chromedp.EvaluateAsDevTools("document.querySelector('#ifc').contentWindow.document.body.outerHTML", &iframeContent),
		chromedp.OuterHTML(`html`, &htmlContent, chromedp.ByQuery),
	)

	fmt.Println(iframeContent)

	if err != nil {
		log.Fatalf("Failed to execute chromedp tasks: %v", err)
	}

	// 将获取的内容写入文件
	err = os.WriteFile(outputFilePath, []byte(iframeContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write iframe content to file: %v", err)
	}

	log.Printf("Iframe content successfully written to %s\n", outputFilePath)
}
