package fetch

import "log"

type WorkResult struct {
	ip      string
	success bool
}

func Worker(checker checkerFunc, jobs <-chan string, results chan<- WorkResult) {
	for ip := range jobs {
		checkResult, err := checker(ip)
		var workResult WorkResult
		if err != nil {
			log.Printf("Error with ip %s: %s\n", ip, err.Error())
			workResult = WorkResult{ip: "", success: false}
		} else {
			workResult = WorkResult{
				ip:      ip,
				success: checkResult,
			}
		}
		results <- workResult
	}
}
