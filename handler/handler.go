package handler

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"hao6v/internal/dao"
	"hao6v/internal/model"
	"hao6v/pkg/utils"
	"regexp"
	"sync"
)

type Processor struct {
	sync.Once
}

// 匹配网站
const r1 = `https://www\.hao6v\.tv/[\w\W]+/(\d{4}-\d{2}-\d{2})/(\d+|[\w-]+)\.html`

var urlPattern = regexp.MustCompile(r1)

var (
	once     sync.Once
	instance *Processor
)

func GetProcessorInstance() *Processor {
	once.Do(func() {
		instance = &Processor{}
	})
	return instance
}

func (p *Processor) LinkCallback(e *colly.HTMLElement) {
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
}

func (p *Processor) BodyCallback(e *colly.HTMLElement) {
	r := e.DOM.Text()
	fmt.Println(utils.ConvertGBKToUTF8(r))
}

func (p *Processor) ParseHTML(e *colly.HTMLElement) {
	// 如果不匹配的网站，那么不解析
	if !urlPattern.MatchString(e.Request.AbsoluteURL(e.Request.URL.String())) {
		fmt.Println("不解析这个网站，不匹配规则", e.Request.URL.String())
	} else {
		fmt.Println("Found", e.Request.URL.String())
	}
}

func (p *Processor) RequestCallback(r *colly.Request) {
	fmt.Println("Visiting", r.AbsoluteURL(r.URL.String()))
}
func (p *Processor) ResponseCallback(r *colly.Response) {
	// 1. 编码转换
	_, err := utils.ConvertEncoding(r.Body)
	if err != nil {
		panic(err)
	}
}
