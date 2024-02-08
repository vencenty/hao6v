package model

type Queue chan *PageContent

func NewQueue(bufferSize int) Queue {
	q := make(Queue, bufferSize)
	return q
}
