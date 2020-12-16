package main

import (
	"bufio"
	"flag"
	"fmt"
	//	"io/ioutil"
	"log"
	//	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

func main() {
	var concurrency int
	var payloads string
	var match string
	var appendMode bool
	flag.IntVar(&concurrency, "t", 30, "Set Concurrency")
	flag.StringVar(&payloads, "p", "", "payload list")
	flag.StringVar(&match, "m", "", "match pattern")
	flag.BoolVar(&appendMode, "a", false, "append a payload to certain parameter")
	flag.Parse()

	if payloads != "" && match != "" {
		var wg sync.WaitGroup
		for i := 0; i <= concurrency; i++ {
			wg.Add(1)
			go func() {
				test_ssrf(payloads, match, appendMode)
				wg.Done()
			}()
			wg.Wait()
		}
	}
}

func test_ssrf(payloads string, match string, appendMode bool) {

	file, err := os.Open(payloads)

	if err != nil {
		log.Fatal("[!] error reading payloads file")
	}

	defer file.Close()
	time.Sleep(time.Millisecond * 10)
	scanner := bufio.NewScanner(os.Stdin)

	pScanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for pScanner.Scan() {
			link := scanner.Text()
			payload := pScanner.Text()
			qs := url.Values{}
			u, err := url.Parse(link)
			if err != nil {
				log.Fatal("[!] Error parsing the url")
			}
			for param, vv := range u.Query() {
				if appendMode {
					qs.Set(param, vv[0]+payload)
				} else {
					qs.Set(param, payload)
				}

				u.RawQuery = qs.Encode()
				fmt.Printf("[?] Testing %s", u)
			}
		}

	}
}
