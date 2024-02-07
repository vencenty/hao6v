package crawler

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"gorm.io/gorm"
	"hao6v/internal/model"
	"hao6v/pkg/utils"
	"log"
	"regexp"
)

type URLService struct {
	Dao       *gorm.DB
	TargetUrl string
}

func NewURLService(dao *gorm.DB) *URLService {
	return &URLService{
		Dao: dao,
	}
}

const r1 = `https://www\.hao6v\.tv/[\w\W]+/(\d{4}-\d{2}-\d{2})/(\d+|[\w-]+)\.html`

var urlPattern = regexp.MustCompile(r1)

func (s *URLService) SetTargetURL(url string) {
	s.TargetUrl = url
}

func (s *URLService) Start(targetUrl string) {
	// 初始化Collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.hao6v.tv"),
		// 可以在这里设置Collector的选项，如User-Agent、并发数等
	)
	extensions.RandomUserAgent(c)
	extensions.RandomMobileUserAgent(c)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if urlPattern.MatchString(e.Request.AbsoluteURL(e.Request.URL.String())) {
			pageContent := &model.PageContent{
				URL:  e.Request.AbsoluteURL(e.Request.URL.String()),
				HTML: e.Response.Body,
			}
			// 写入队列
			p.Jobs <- pageContent
		}

		c.Visit(e.Attr("href"))
	})
	// 注册请求前的回调函数
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.AbsoluteURL(r.URL.String()))

	})
	// 注册响应后的回调函数
	c.OnResponse(func(r *colly.Response) {
		var err error
		// 1. 编码转换
		r.Body, err = utils.ConvertEncoding(r.Body)
		if err != nil {
			panic(err)
		}
	})
	// 启动爬虫
	err := c.Visit(targetUrl)
	if err != nil {
		log.Fatal(err)
	}
}
