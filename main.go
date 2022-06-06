package main

import (
	"flag"
)

func checker(ip string) bool {
	return ip == ""
}

func main() {
	region := flag.String("region", "", "region to search server in")
	doNotUseCache := flag.Bool("no-cache", false, "don't use local sqlite database to cache requests")
	flag.Parse()
}
