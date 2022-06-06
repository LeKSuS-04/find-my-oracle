package main

import (
	"flag"
	"fmt"

	"github.com/LeKSuS-04/find-my-oracle/fetch"
)

func main() {
	region := flag.String("region", "", "region to search server in")
	// doNotUseCache := flag.Bool("no-cache", false, "don't use local sqlite database to cache requests")
	flag.Parse()

	ipsMasks, err := fetch.FetchIPMasks(*region)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ipsMasks)
}
