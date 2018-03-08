package worker

var jobs chan RefreshRequest

type RefreshRequest struct {
	serverId string
	result   chan int
}

func CreateRefreshWorker(workerCount int, jobCount int) {
	jobs = make(chan RefreshRequest, jobCount)
	for w := 1; w <= workerCount; w++ {
		go refreshWorker(w, jobs)
	}
}

func refreshWorker(workerId int, request <-chan RefreshRequest) {

}
