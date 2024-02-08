package spider

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"hao6v/internal/model"
	"hao6v/pkg/global"
	"hao6v/pkg/utils"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"
)

type Scape struct {
	model.Queue
	wg sync.WaitGroup
}

const r1 = `https://www\.hao6v\.tv/[\w\W]+/(\d{4}-\d{2}-\d{2})/(\d+|[\w-]+)\.html`

var urlPattern = regexp.MustCompile(r1)

func NewScrape(q model.Queue) *Scape {
	return &Scape{
		Queue: q,
		wg:    sync.WaitGroup{},
	}
}

func (s *Scape) handle(targetUrl string) {
	// 在方法开始使用defer来设置recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in handle. Error:", r)
			// 这里可以添加更多的错误处理逻辑，比如记录到日志文件
		}
	}()
	// 初始化Collector
	c := colly.NewCollector(
		colly.URLFilters(
			regexp.MustCompile(`https://www\.hao6v\.tv/dy/`),
		),
		colly.AllowedDomains("www.hao6v.tv"),
		// 可以在这里设置Collector的选项，如User-Agent、并发数等
	)
	extensions.RandomUserAgent(c)
	extensions.RandomMobileUserAgent(c)

	//	这一步匹配合适的链接放入处理队列
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		err := c.Visit(e.Request.AbsoluteURL(link))
		if err != nil {
			fmt.Println("访问", e.Request.URL, "出错", err.Error())
		}
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
			fmt.Println("recover", err.Error())
		}
	})

	c.OnResponse(func(r *colly.Response) {
		var resources []model.Resource

		pattern := `https:\/\/www\.hao6v\.tv\/dy\/\d{4}-\d{2}-\d{2}\/\d+\.html`
		reg := regexp.MustCompile(pattern)

		// url不匹配的话
		if !reg.MatchString(r.Request.AbsoluteURL(r.Request.URL.String())) {
			fmt.Println(r.Request.AbsoluteURL(r.Request.URL.String()), "不满足数据匹配要求，跳过处理")
			return
		}
		// 初始化goquery
		byteReader := bytes.NewReader(r.Body)
		doc, err := goquery.NewDocumentFromReader(byteReader)
		if err != nil {
			panic("解析文档错误")
		}

		var description string
		var isDetail bool
		doc.Find("#endText p").Siblings().Each(func(i int, selection *goquery.Selection) {
			tag := goquery.NodeName(selection)

			// 不是p标签不处理
			if tag == "p" {
				n := strings.Index(selection.Text(), "◎")
				if n != -1 {
					isDetail = true
					description = selection.Text()
				}
			}
		})

		if !isDetail {
			return
		}

		title := doc.Find(`#main > div.col6 > div > h1`).Text()
		fmt.Println(title)

		doc.Find("#endText > table > tbody td  a").Each(func(i int, selection *goquery.Selection) {

			href, exists := selection.Attr("href")
			if !exists {
				href = ""
			}
			title := selection.Text()
			resource := model.Resource{
				ResourceTitle: title,
				DownloadUrl:   href,
				Type:          utils.IdentifyLinkType(href),
			}
			// 写到一起去
			resources = append(resources, resource)
		})

		// 获取海报图
		posterNode := doc.Find("#endText p img")
		poster, _ := posterNode.Attr("src")

		page := &model.Page{
			AbsoluteUrl: r.Request.AbsoluteURL(r.Request.URL.String()),
			Status:      0,
			Title:       title,
			Poster:      poster,
			Description: description,
			Resources:   resources,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		}

		result := global.DB.Where(&model.Page{AbsoluteUrl: page.AbsoluteUrl}).FirstOrCreate(page)
		fmt.Println(result, page.AbsoluteUrl, "写入成功")
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("error of colly:", err.Error())
	})

	// 启动爬虫
	err := c.Visit(targetUrl)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Scape) Start(targetUrl string) {
	s.handle(targetUrl)
}

func (s *Scape) Wait() {
	s.wg.Wait()
}
