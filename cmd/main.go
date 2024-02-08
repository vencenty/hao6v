package main

import (
	"hao6v/internal/model"
	"hao6v/pkg/spider"
)

func main() {

	//utils.DownloadDemoHTML("https://www.hao6v.tv/dy/2024-02-02/43708.html", "detail.html")
	//utils.DownloadDemoHTML("https://www.hao6v.tv/", "home.html")
	//utils.DownloadDemoHTML("https://www.hao6v.tv/s/gf/", "home.html")
	//return
	// 初始化队列
	q := model.NewQueue(10)

	// 初始化爬虫
	scrape := spider.NewScrape(q)
	scrape.Start("https://www.hao6v.tv/dy/2024-02-06/43744.html")

	//// 初始化解析器
	//parserService := service.NewParser(q)
	//
	//// 解析器开始工作
	//parserService.Start()

	//// 等待爬虫完成
	//scrape.Wait()
	//
	//// 等待解析器执行完成
	//parserService.Wait()

}
