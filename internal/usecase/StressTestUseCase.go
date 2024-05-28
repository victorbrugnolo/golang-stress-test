package usecase

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

func Execute(url string, requests, concurrency int) error {
	t := time.Now()
	concurrencyGroupSize := requests / concurrency
	concurrencyGroupSizeRemainder := requests % concurrency

	totalByStatus := sync.Map{}

	for i := 0; i < concurrencyGroupSize; i++ {
		wg := &sync.WaitGroup{}

		for j := 0; j < concurrency; j++ {
			makeConcurrentRequest(wg, url, &totalByStatus)
		}

		wg.Wait()
	}

	if concurrencyGroupSizeRemainder > 0 {
		wg := &sync.WaitGroup{}

		for j := 0; j < concurrencyGroupSizeRemainder; j++ {
			makeConcurrentRequest(wg, url, &totalByStatus)
		}

		wg.Wait()
	}

	fmt.Println()
	fmt.Printf("Time elapsed: %s\n", time.Since(t))
	fmt.Printf("Total requests: %d\n", requests)

	totalByStatus.Range(func(key, value interface{}) bool {
		status := key.(int)
		count := value.(*int64)

		if status == 0 {
			fmt.Printf("Error: %d\n", *count)
		}

		fmt.Printf("Status %d: %d\n", status, *count)
		return true
	})

	return nil
}

func makeConcurrentRequest(wg *sync.WaitGroup, url string, totalByStatus *sync.Map) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		statusCode := makeRequest(url)
		val, _ := totalByStatus.LoadOrStore(statusCode, new(int64))
		ptr := val.(*int64)
		atomic.AddInt64(ptr, 1)
	}()
}

func makeRequest(url string) int {
	fmt.Printf("Requesting %s\n", url)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Request error: %s\n", err.Error())
		return 0
	}

	defer resp.Body.Close()

	return resp.StatusCode
}
