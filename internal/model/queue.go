package model

type JobQueue chan PageContent

//func NewJobQueue(bufferSize int) *JobQueue {
//	return &JobQueue{
//		Jobs: make(chan *PageContent, bufferSize),
//	}
//}
//
//func (jq *JobQueue) AddJob(page *PageContent) {
//	jq.Jobs <- page
//}
//
//func (jq *JobQueue) Close() {
//	close(jq.Jobs)
//}
