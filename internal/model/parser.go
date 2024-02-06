package model

import (
	"hao6v/internal/dao"
	"sync"
)

type Parser struct {
	Jobs     chan *PageContent
	wg       sync.WaitGroup
	Parallel int
}

// NewParser 初始化一个Parser
func NewParser() *Parser {
	return &Parser{
		Jobs:     make(chan *PageContent, 10),
		wg:       sync.WaitGroup{},
		Parallel: 2,
	}
}

// SetParallel 设置并行数量
func (p *Parser) SetParallel(num int) {
	p.Parallel = num
}

// ParallelProcess 并行处理
func (p *Parser) ParallelProcess(page []*PageContent) {

}

func (p *Parser) Start() {
	p.wg.Add(1)
	go p.fetchJob()
}

func (p *Parser) process(page *PageContent) {
	// 写入数据哭
	pageDao := dao.NewPageDao()
	pageModel := &Page{
		AbsoluteUrl: page.URL,
	}
	pageDao.FirstOrCreate(pageModel)
}

func (p *Parser) fetchJob() {
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

func (p *Parser) Wait() {
	p.wg.Wait()
}
