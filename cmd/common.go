package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var RequestMetricsArray []RecordRequestData

type RecordRequestData struct {
	Iteration int
	Thread    int
	Latency   float64
	URL       string
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func convertMetricstoString() [][]string {
	var arr [][]string
	for _, item := range RequestMetricsArray {
		Iteration := strconv.Itoa(item.Iteration)
		Thread := strconv.Itoa(item.Thread)
		Latency := fmt.Sprintf("%f", item.Latency)
		URL := item.URL
		data := []string{Iteration, Thread, Latency, URL}
		arr = append(arr, data)
	}
	return arr
}

func getRequest(URL string, iteration_count int, request_count int) {
	t1 := time.Now()
	response, err := http.Get(URL)
	t2 := time.Now()
	defer response.Body.Close()
	checkError("looks like an error has occur", err)
	// bodyBytes, err2 := ioutil.ReadAll(response.Body)
	// checkError("Error in reading HTTP Response", err2)
	// var body interface{}
	// json.Unmarshal(bodyBytes, &body)
	// fmt.Println(body)
	time_diff := t2.Sub(t1)
	fmt.Println("Request Count: ", request_count, " ,Iteration: ", iteration_count, " ,Latency: ", time_diff.Seconds(), "URL", URL[21:])
	metrics := RecordRequestData{}
	metrics.Iteration = iteration_count
	metrics.Thread = request_count
	metrics.Latency = time_diff.Seconds()
	metrics.URL = URL[21:]
	RequestMetricsArray = append(RequestMetricsArray, metrics)
}

func writeMetricstoCSV() {
	convertedMetrics := convertMetricstoString()
	file, err := os.Create("results.csv")
	checkError("Err in creating results.csv", err)
	defer file.Close() //At the end of the execution of this function close this File
	writer := csv.NewWriter(file)
	defer writer.Flush()
	headers := []string{"Iteration", "Thread", "Latency(in Seconds)", "API"}
	err = writer.Write(headers)
	checkError("Error in writing metrics of CSV", err)
	for _, metrics := range convertedMetrics {
		err = writer.Write(metrics)
		checkError("Error in writing metrics of CSV", err)
	}
}
