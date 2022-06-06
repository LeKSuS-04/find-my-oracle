package main

import (
	"flag"

	"github.com/LeKSuS-04/find-my-oracle/checker"
	"github.com/LeKSuS-04/find-my-oracle/fetch"
)

func main() {
	region := *flag.String("region", "", "region to search server in")
	workerAmount := *flag.Int("threads", 20, "amount of threads")
	timeout := *flag.Int("timeout", 10000, "timeout in milliseconds to use in checker function")
	noCache := *flag.Bool("no-cache", false, "don't use local sqlite database to cache requests")
	flag.Parse()

	fetch.Scan(region, !noCache, workerAmount, checker.GetChecker(timeout))
}
