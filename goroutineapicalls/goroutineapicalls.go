package main

import (
	"strconv"
	"sync"
	"time"
)

type Input struct {
	ID   int
	Text string
}

func main() {
	// generate input
	bigInput := make([]Input, 0)
	for i := 0; i < 1000; i++ {
		bigInput = append(bigInput, Input{
			ID:   i,
			Text: "text-" + strconv.Itoa(i),
		})
	}
	process(bigInput)
}

// process all input
func process(records []Input) {
	totalRecords := len(records)
	maxParallel := 50
	batchSize := 100

	wg := &sync.WaitGroup{}
	sem := make(chan struct{}, maxParallel)

	routineNum := 0
	for i := 0; i < totalRecords; i += batchSize {
		end := i + batchSize
		if end > totalRecords {
			end = totalRecords
		}

		// generate batch of 100 records
		batch := records[i:end]
		wg.Add(1)
		sem <- struct{}{}

		routineNum++
		go makeAPICall(wg, sem, batch, routineNum)

	}

	wg.Wait()
	println("Done processing all records.")
}

func makeAPICall(wg *sync.WaitGroup, sem chan struct{}, batch []Input, routineNum int) {
	defer wg.Done()
	defer func() { <-sem }()

	// makes batch API call (max 100 in a batch)

	for i, record := range batch {
		println("Routine Num:", routineNum, "(", "ID:", record.ID, "Text:", record.Text, ")", "Batch Size:", len(batch))
		if i%10 == 0 {
			time.Sleep(1 * time.Second)
		}
	}
	println("Done processing this batch.")
}
