package fetch

import (
	"log"
	"strings"
	"time"

	"github.com/LeKSuS-04/find-my-oracle/db"
)

type checkerFunc func(string) (bool, error)

func Scan(region string, useCache bool, workerAmount int, checker checkerFunc) {
	ipMasks, err := FetchIPMasks(region)
	if err != nil {
		log.Fatalln(err)
	}

	scanner := scanWithoutCache
	if useCache {
		scanner = scanWithCache
	}

	start := time.Now()
	log.Printf("%d subnets to scan: %s\n", len(ipMasks), strings.Join(ipMasks, ", "))
	scanner(ipMasks, workerAmount, checker)
	log.Printf("Scanned all subnets in %s\n", time.Since(start).String())
}

func scanWithCache(ipMasks []string, workerAmount int, checker checkerFunc) {
	cache, err := db.InitDB(ipMasks)
	if err != nil {
		log.Fatalln(err)
	}
	defer cache.Close()

	for _, mask := range ipMasks {
		if db.IsIPMaskChecked(cache, mask) {
			log.Printf("Skipping subnet %s because it's marked as checked in cache", mask)
		} else {
			scanIPMask(mask, workerAmount, checker)
			db.SetIPMaskChecked(cache, mask)
		}
	}
}

func scanWithoutCache(ipMasks []string, workerAmount int, checker checkerFunc) {
	for _, mask := range ipMasks {
		scanIPMask(mask, workerAmount, checker)
	}
}

func scanIPMask(mask string, workerAmount int, checker checkerFunc) {
	ips, err := GetIpsFromMask(mask)
	if err != nil {
		log.Printf("Invalid subnet %s: %s\n", mask, err.Error())
		return
	}
	log.Printf("Scanning subnet %s (%d ips)...\n", mask, len(ips))

	results := make(chan WorkResult, len(ips))
	jobs := make(chan string, len(ips))
	for i := 0; i < workerAmount; i++ {
		go Worker(checker, jobs, results)
	}

	for _, ip := range ips {
		jobs <- ip
	}
	close(jobs)

	for i := 0; i < len(ips); i++ {
		result := <-results
		if result.success {
			log.Printf("FOUND LOST SERVER UNDER IP %s!!! UNBELIEVABLE!!!\n", result.ip)
			log.Fatalln("Terminating...")
		}
	}
}
