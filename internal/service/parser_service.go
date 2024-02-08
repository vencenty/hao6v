package service

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"hao6v/internal/dao"
	"hao6v/internal/model"
	"sync"
)

type ParserService struct {
	Queue    model.Queue //队列
	wg       sync.WaitGroup
	Parallel int //并行处理的数量
}

type Parser interface {
	Parse(page *model.PageContent) error
}

type ParserOption func(s *ParserService)

func ParserParallelNum(num int) ParserOption {
	return func(s *ParserService) {
		s.Parallel = num
	}
}

// NewParser 初始化一个Parser
func NewParser(q model.Queue, opts ...ParserOption) *ParserService {
	s := &ParserService{
		Queue:    q,
		wg:       sync.WaitGroup{},
		Parallel: 2,
	}

	for _, f := range opts {
		f(s)
	}

	return s
}

// SetParallel 设置并行数量
func (p *ParserService) SetParallel(num int) {
	p.Parallel = num
}

// ParallelProcess 并行处理
func (p *ParserService) ParallelProcess(page []*model.PageContent) {

}

func (p *ParserService) Start() {
	p.wg.Add(1)
	go p.fetchJob()
}

func (p *ParserService) process(job *model.PageContent) {
	// 写入数据哭
	pageDao := dao.NewPageDao()
	page := &model.Page{
		AbsoluteUrl: job.URL,
	}

	byteReader := bytes.NewReader(job.HTML)
	doc, err := goquery.NewDocumentFromReader(byteReader)
	fmt.Println(doc.Find("title").Text())

	return

	if err != nil {
		fmt.Println("当前job处理错误", job.URL)
	}

	r := doc.Find(`document.querySelector("#main > div.col6 > div > h1").innerText`).Text()
	fmt.Println("r=", r, page.AbsoluteUrl)
	return

	// 写入数据
	pageDao.FirstOrCreate(page)
}

func (p *ParserService) fetchJob() {
	defer p.wg.Done()
	for {
		select {
		case v, ok := <-p.Queue:
			if !ok {
				break
			}
			fmt.Println("add to parser", v.URL)
			p.process(v)
		}
	}
}

func (p *ParserService) Wait() {
	p.wg.Wait()
}
