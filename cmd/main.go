package main

import (
	"hao6v/internal/model"
	"hao6v/pkg/spider"
)

func main() {
	// 初始化队列
	q := model.NewQueue(10)
	// 初始化爬虫
	scrape := spider.NewScrape(q)
	scrape.Start("https://www.hao6v.tv/dy/2024-02-06/43744.html")
}
