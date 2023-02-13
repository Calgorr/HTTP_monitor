package monitor

import (
	"net/http"
	"strings"
	"time"

	model "github.com/Calgorr/IE_Backend_Fall/Model"
	"github.com/Calgorr/IE_Backend_Fall/database"
)

type Result struct {
	httpStatusCode int
	description    string
	url            *model.URL
}

func sendHTTPResp(url *model.URL) *Result {

	resp, err := http.Get(url.Address)
	result := &Result{}
	if resp.StatusCode/100 != 2 {
		result.httpStatusCode = 0
		result.description = err.Error()
		result.url = url
	} else {
		result.httpStatusCode = resp.StatusCode
		result.description = ""
		result.url = url
	}
	return result
}

func Worker(jobs <-chan *model.URL, results chan<- *Result) {
	for job := range jobs {
		results <- sendHTTPResp(job)
	}
}

func DoEveryPeriod(d time.Duration) {
	jobs, results := make(chan *model.URL, 100), make(chan *Result, 100)
	ticker := time.NewTicker(d)
	for {
		select {
		case <-ticker.C:
			for i := 0; i < 100; i++ {
				go Worker(jobs, results)
			}
			urls := database.GetAllURLs()
			for _, url := range urls {
				jobs <- url
			}
			for res := range results {
				if res.httpStatusCode/100 != 2 && strings.Compare(res.description, "") == 0 {
					database.IncrementFailedByOne(res.url)
				}
				else if
			}

		}
	}
}
