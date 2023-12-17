package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"
)

// JSON input and output format
type sortInput struct {
	ToSort [][]int `json:"to_sort"`
}
type sortOutput struct {
	SortedArrays [][]int `json:"sorted_arrays"`
	TimeNS       int64   `json:"time_ns"`
}

// Function which sort each sub-array sequentially
func handleSingleSort(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()
	var input sortInput

	// Parse JSON request
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// sort each sub-array sequentially
	var output sortOutput
	output.SortedArrays = make([][]int, len(input.ToSort))
	for i, arr := range input.ToSort {
		sort.Ints(arr)
		output.SortedArrays[i] = arr
	}

	endTime := time.Now()
	// calculating time
	elapsedTime := endTime.Sub(startTime).Nanoseconds()
	output.TimeNS = elapsedTime

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}

// sorting sort each sub-array concurrently using Go's concurrency features
func sortSubArray(subArray []int, ch chan []int) {
	sort.Ints(subArray)
	ch <- subArray
}

func handleConcurrentSort(w http.ResponseWriter, r *http.Request) {
	var input sortInput
	var output sortOutput

	// Parse JSON request
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	startTime := time.Now()

	sortedArrays := make([][]int, len(input.ToSort))
	ch := make(chan []int, len(input.ToSort))

	for i, subArray := range input.ToSort {
		go sortSubArray(subArray, ch)
		sortedArrays[i] = <-ch
	}

	close(ch)

	output.SortedArrays = sortedArrays
	output.TimeNS = time.Since(startTime).Nanoseconds()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}

func main() {
	// // API routes
	http.HandleFunc("/process_single", handleSingleSort)
	http.HandleFunc("/process_concurrent", handleConcurrentSort)

	port := ":8000"
	fmt.Println("Server is running on port" + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))
}
