//

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
)

type Result struct {
	status int
	err    error
}

func worker(id int, url string, requests int, results chan<- Result, wg *sync.WaitGroup, enableLogs bool, logColor *color.Color) {
	defer wg.Done()
	for i := 0; i < requests; i++ {
		if enableLogs {
			logColor.Printf("Worker %d: Sending request %d\n", id, i+1)
		}
		resp, err := http.Get(url)
		if err != nil {
			results <- Result{status: 0, err: err}
			continue
		}
		results <- Result{status: resp.StatusCode, err: nil}
		if enableLogs {
			logColor.Printf("Worker %d: Received response %d with status %d\n", id, i+1, resp.StatusCode)
		}
		resp.Body.Close()
	}
}

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	totalRequests := flag.Int("requests", 100, "Número total de requests")
	concurrency := flag.Int("concurrency", 10, "Número de chamadas simultâneas")
	enableLogs := flag.Bool("logs", false, "Habilitar logs detalhados")
	flag.Parse()

	if *url == "" {
		fmt.Println("A URL do serviço deve ser fornecida.")
		return
	}

	// Verifica se o terminal suporta cores
	useColors := isatty.IsTerminal(os.Stdout.Fd())

	requestsPerWorker := *totalRequests / *concurrency
	results := make(chan Result, *totalRequests)
	var wg sync.WaitGroup

	startTime := time.Now()

	colors := []*color.Color{
		color.New(color.FgRed),
		color.New(color.FgGreen),
		color.New(color.FgYellow),
		color.New(color.FgBlue),
		color.New(color.FgMagenta),
		color.New(color.FgCyan),
		color.New(color.FgWhite),
	}

	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		var logColor *color.Color
		if useColors {
			logColor = colors[randGen.Intn(len(colors))]
		} else {
			logColor = color.New(color.FgWhite) // Default to white if colors are not supported
		}
		go worker(i, *url, requestsPerWorker, results, &wg, *enableLogs, logColor)
	}

	wg.Wait()
	close(results)

	endTime := time.Now()

	totalReqs := 0
	successRequests := 0
	statusCodes := make(map[int]int)

	for result := range results {
		totalReqs++
		if result.err == nil {
			statusCodes[result.status]++
			if result.status == http.StatusOK {
				successRequests++
			}
		}
	}

	fmt.Printf("Tempo total gasto na execução: %v\n", endTime.Sub(startTime))
	fmt.Printf("Quantidade total de requests realizados: %d\n", totalReqs)
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", successRequests)
	fmt.Printf("Distribuição de outros códigos de status HTTP:\n")
	for code, count := range statusCodes {
		if code != http.StatusOK {
			fmt.Printf("  %d: %d\n", code, count)
		}
	}
}
