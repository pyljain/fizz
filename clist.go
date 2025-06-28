package main

import "sync"

type cList struct {
	lock  *sync.Mutex
	files []string
}

func newcList() *cList {

	return &cList{
		lock:  &sync.Mutex{},
		files: make([]string, 0),
	}
}

func (cl *cList) add(fileName string) {
	cl.lock.Lock()
	cl.files = append(cl.files, fileName)
	cl.lock.Unlock()
}
