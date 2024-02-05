package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"hao6v/handler"
	"hao6v/internal/dao"
	"hao6v/internal/model"
	"log"
	"regexp"
)

const r1 = `https://www\.hao6v\.tv/[\w\W]+/(\d{4}-\d{2}-\d{2})/(\d+|[\w-]+)\.html`

var urlPattern = regexp.MustCompile(r1)

func main() {
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
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if urlPattern.MatchString(e.Request.AbsoluteURL(e.Request.URL.String())) {
			page := &model.Page{
				AbsoluteUrl: e.Request.AbsoluteURL(e.Request.URL.String()),
			}
			pageDao := dao.NewUrlDao()
			_, _, err := pageDao.FirstOrCreate(page)
			if err != nil {
				fmt.Println(err)
			}
		}

		link := e.Attr("href")
		// 访问下个链接
		e.Request.Visit(link)
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {

	})

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
