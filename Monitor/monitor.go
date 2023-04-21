package monitor

import (
	"net/http"
	"net/url"

	model "github.com/Calgorr/IE_Backend_Fall/Model"
	"github.com/gammazero/workerpool"
)

type Monitor struct {
	wp         workerpool.WorkerPool
	urls       []*model.URL
	workerSize int
}

type Resp struct {
	StatusCode int
	err        string
}

func NewMonitor(workerSize int) *Monitor {
	mnt := new(Monitor)
	mnt.urls = make([]*model.URL, 0)
	mnt.workerSize = workerSize
	mnt.wp = *workerpool.New(workerSize)
	return mnt
}

func sendRequest(urlString string, client http.Client, result chan<- *Resp) {
	var response *Resp
	req := http.Request{}
	req.URL, _ = url.Parse(urlString)
	req.Method = "GET"
	resp, err := client.Do(&req)
	if err != nil {
		response.StatusCode = 0
	}
	response.StatusCode = resp.StatusCode
	result <- response
}
