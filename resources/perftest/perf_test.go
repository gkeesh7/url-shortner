package perftest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

func hitURLShortner() (interface{}, interface{}) {
	requestBody, _ := json.Marshal(map[string]string{
		"url":        "https://youtu.be/dQw4w9WgXcQ?t=43",
		"request_id": "asdlfjhlaksdjffsajkflkjghjasfflkg",
	})
	resp, err := http.Post("http://0.0.0.0:8080/shorten", "application/json", bytes.NewBuffer(requestBody))
	return resp, err
}

func serially(wg *sync.WaitGroup, n int) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		hitURLShortner()
	}
}

func PerfTest() {
	t1 := time.Now()

	numGoRoutines := 5            // Number of go routines
	maxQueriesPerGoRoutine := 200 //Number of request each go Routine is going to make Serially
	var wg sync.WaitGroup

	wg.Add(numGoRoutines)
	for i := 0; i < numGoRoutines; i++ {
		go serially(&wg, maxQueriesPerGoRoutine) //Fan Out
	}
	wg.Wait() //Fan In

	t2 := time.Now()
	fmt.Printf("Time taken to add %v records %v\n", numGoRoutines*maxQueriesPerGoRoutine, t2.Sub(t1))
	return
}

//Perf Testing
func TestAddingRecordsConcurrently(t *testing.T) {
	PerfTest()
}
