package monitor

import (
	"net/http"
	"time"

	model "github.com/Calgorr/IE_Backend_Fall/Model"
)

type Result struct {
	httpStatusCode int
	description    string
}

func sendHTTPResp(url model.URL) Result {

	resp, err := http.Get(url.Address)
	result := Result{}
	if resp.StatusCode/100 != 2 {
		result.httpStatusCode = 0
		result.description = err.Error()
	} else {
		result.httpStatusCode = resp.StatusCode
		result.description = ""
	}
	return result
}

func Worker(jobs <-chan model.URL, results chan<- Result) {
	for job := range jobs {
		results <- sendHTTPResp(job)
	}
}

func DoEveryPeriod(d time.Duration) {
	jobs, results := make(chan model.URL, 100), make(chan Result, 100)
	ticker := time.NewTicker(d)
	for {
		select {
		case <-ticker.C:
			for i := 0; i < 100; i++ {
				go Worker(jobs, results)
			}
			//need to get all the urls send them to the jobs chanel and then get the results andd update the database

		}
	}
}
