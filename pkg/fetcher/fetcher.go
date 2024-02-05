package fetcher

import (
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"hao6v/handler"
	"log"
)

func Run() {
	// 初始化Collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.hao6v.tv"),
		// 可以在这里设置Collector的选项，如User-Agent、并发数等
	)
	extensions.RandomUserAgent(c)
	extensions.RandomMobileUserAgent(c)

	p := handler.GetProcessorInstance()
	// 注册响应后的回调函数
	// 用于保证持续访问
	//c.OnHTML("a[href]", p.LinkCallback)
	c.OnHTML("title", p.BodyCallback)
	//c.OnHTML("a[href]", p.)
	// 注册请求前的回调函数
	c.OnRequest(p.RequestCallback)
	// 注册响应后的回调函数
	c.OnResponse(p.ResponseCallback)
	// 启动爬虫
	err := c.Visit("https://www.hao6v.tv/jddy/2008-01-06/780.html")
	if err != nil {
		log.Fatal(err)
	}

}
