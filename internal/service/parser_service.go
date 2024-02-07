package service

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"
	"hao6v/internal/dao"
	"hao6v/internal/model"
	"sync"
)

type ParserService struct {
	Dao      *gorm.DB
	Jobs     chan *model.PageContent
	wg       sync.WaitGroup
	Parallel int
}

// NewParser 初始化一个Parser
func NewParser(dao *gorm.DB) *ParserService {
	return &ParserService{
		Dao:      dao,
		Jobs:     make(chan *model.PageContent, 10),
		wg:       sync.WaitGroup{},
		Parallel: 2,
	}
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
	if err != nil {
		fmt.Println("当前job处理错误", job.URL)
	}

	r := doc.Find(`document.querySelector("#main > div.col6 > div > h1").innerText`).Text()
	fmt.Println(r)
	return

	// 写入数据
	pageDao.FirstOrCreate(page)
}

func (p *ParserService) fetchJob() {
	defer p.wg.Done()
	for {
		select {
		case v, ok := <-p.Jobs:
			if !ok {
				break
			}
			p.process(v)
		}
	}
}

func (p *ParserService) Wait() {
	p.wg.Wait()
}
