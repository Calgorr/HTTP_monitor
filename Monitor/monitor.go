package monitor

import (
	"fmt"
	"net/http"

	model "github.com/Calgorr/IE_Backend_Fall/Model"
	"github.com/Calgorr/IE_Backend_Fall/database"
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

func (mnt *Monitor) LoadFromDatabase() error {
	var urls []*model.URL
	var err error
	if urls, err = database.GetAllURLs(); err != nil {
		return err
	}
	mnt.urls = urls
	return nil
}

func (mnt *Monitor) Do() {
	for _, url := range mnt.urls {
		mnt.wp.Submit(func() {
			mnt.monitorURL(url)
		})
	}
}

func (mnt *Monitor) monitorURL(url *model.URL) {
	req, err := url.SendRequest()
	if err != nil {
		fmt.Println(err)
		req = new(model.Request)
		req.URLID = url.ID
		req.StatusCode = http.StatusBadRequest
	}
	if err := database.AddRequest(req); err != nil {
		fmt.Println(err)
	}
	if req.StatusCode/100 != 2 {

		if err := database.IncrementFailedByOne(url); err != nil {
			fmt.Println(err)
		}
	}
}
