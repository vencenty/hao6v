package model

import (
	"sync"
)

type Parser struct {
	Jobs     chan *PageContent
	wg       sync.WaitGroup
	Parallel int
}
