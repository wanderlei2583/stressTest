package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	Duration    time.Duration
	StatusCode  int
	StatusError error
}

type Report struct {
	TotalTime     time.Duration
	TotalRequests int
	StatusCodes   map[int]int
	Errors        int
}

func main() {
	url := flag.String("url", "", "URL do serviço para o teste.")
	requests := flag.Int("requests", 0, "Quantidade de requisições.")
	concurrency := flag.Int(
		"concurrency",
		1,
		"Número de chamadas simultâneas.",
	)
	flag.Parse()

	if *url == "" || *requests == 0 {
		fmt.Println("URL e número de requisições devem ser informados.")
		flag.PrintDefaults()
		return
	}

	fmt.Printf("Iniciando teste de carga para %s\n", *url)
	fmt.Printf("Quantidade de requisições: %d\n", *requests)
	fmt.Printf("Número de chamadas simultâneas: %d\n\n", *concurrency)

	results := make(chan Result, *requests)
	start := time.Now()

	var wg sync.WaitGroup
	requestsPerWorker := *requests / *concurrency
	remainingRequests := *requests % *concurrency

	for i := 0; i < *concurrency; i++ {
		requestsForThisWorker := requestsPerWorker
		if remainingRequests > 0 {
			requestsForThisWorker++
			remainingRequests--
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			makeRequests(*url, requestsForThisWorker, results)
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	report := Report{
		StatusCodes: make(map[int]int),
	}

	for result := range results {
		report.TotalRequests++
		if result.StatusError != nil {
			report.Errors++
		} else {
			report.StatusCodes[result.StatusCode]++
		}
	}

	report.TotalTime = time.Since(start)
	printReport(report)
}

func makeRequests(url string, count int, results chan<- Result) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	for i := 0; i < count; i++ {
		start := time.Now()
		resp, err := client.Get(url)
		duration := time.Since(start)

		if err != nil {
			results <- Result{
				Duration:    duration,
				StatusError: err,
			}
			continue
		}
		results <- Result{
			Duration:   duration,
			StatusCode: resp.StatusCode,
		}
		resp.Body.Close()
	}
}

func printReport(report Report) {
	fmt.Printf("\nRelatório de Teste de Carga\n")
	fmt.Printf("===========================")

	fmt.Printf("\nDuração total: %v\n", report.TotalTime)
	fmt.Printf("\nQuantidade de requisições: %d\n", report.TotalRequests)
	fmt.Printf("\nQuantidade de erros: %d\n\n", report.Errors)

	fmt.Println("Distribuição de Status HTTP:")
	fmt.Println("----------------------------")
	for code, count := range report.StatusCodes {
		fmt.Printf(
			"HTTP %d: %d requests (%.1f%%)\n",
			code,
			count,
			float64(count)/float64(report.TotalRequests)*100,
		)
	}
}
