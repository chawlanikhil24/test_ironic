// Copyright © 2019 Nikhli Chawla chawlanikhil24@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// nodeDetailCmd represents the nodeDetail command
var nodeDetailCmd = &cobra.Command{
	Use:   "nodeDetail",
	Short: "Run benchmark for Baremetal detailing APIs",
	Long: `nodeDetail Command is used to hit the APIs which provides details for the IRONIC provisioned bare metals.
	List, Searching, Creating, Updating, and Deleting of bare metal Node resources are done through the /v1/nodes resource.
	There are also several sub-resources, which allow further actions to be performed on a bare metal Node.

	Following are the type of requests that are handled by this API:


GET	/v1/nodes			List Nodes
GET	/v1/nodes/detail		List Nodes Detailed
GET	/v1/nodes/node_identity	Show Node Detailed
POST	/v1/nodes			Create Node --NOT SUPPORTED FOR NOW--
PATCH	/v1/nodes/{node_identity}	Update Node --NOT SUPPORTED FOR NOW--
DELETE	/v1/nodes/{node_identity}	Delete Node --NOT SUPPORTED FOR NOW--`,

	Run: func(cmd *cobra.Command, args []string) {
		api := cmd.Flag("api").Value.String()
		host := cmd.Flag("host").Value.String()
		port := cmd.Flag("port").Value.String()
		request := cmd.Flag("request").Value.String()
		csv := cmd.Flag("csv").Value.String()
		repeat, err := strconv.Atoi(cmd.Flag("repeat").Value.String())
		checkError("Error in coverting String to Int in nodeDetailCmd for repeat variable", err)
		threads, err := strconv.Atoi(cmd.Flag("threads").Value.String())
		checkError("Error in coverting String to Int in nodeDetailCmd for threads variable", err)
		timer, err := time.ParseDuration((cmd.Flag("timer").Value.String()) + "s") // "s" represents 'seconds' unit of time, by appending "s" to the timer string, will create a time duration object which represents "timer" seconds
		checkError("Error in coverting String to time Duration in nodeDetailCmd for timer", err)

		var ironic_server string = "http://" + host + ":" + port + api
		if request == "GET" {
			for iterations := 1; iterations <= repeat; iterations++ {
				for i := 1; i <= threads; i++ {
					go getRequest(ironic_server, iterations, i)
				}
				time.Sleep(timer)
				fmt.Println("SleepOver")
			}
		} else {
			log.Fatal("We are not supporting other requests than GET requests")
			return
		}
		if csv == "true" {
			if len(RequestMetricsArray) > 0 {
				writeMetricstoCSV()
			} else {
				log.Fatal("No Data Collected in RequestMetricsArray")
				return
			}
		}
	},
}

type RecordRequestData struct {
	Iteration int
	Thread    int
	Latency   float64
}

var RequestMetricsArray []RecordRequestData

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
		data := []string{Iteration, Thread, Latency}
		arr = append(arr, data)
	}
	return arr
}

func init() {
	runBenchmarkCmd.AddCommand(nodeDetailCmd)

	var request_type string
	var API string
	var repeat int = 1
	var port int = 6385
	var n_threads int = 10
	var timer int = 5
	var csv bool = false

	nodeDetailCmd.Flags().String("host", "localhost", "Used to specify the ironic server address")
	nodeDetailCmd.Flags().StringVarP(&request_type, "request", "r", "", "Used to specify the request type that should hit ironic server API(GET, POST, PATCH, DELETE etc.)")
	nodeDetailCmd.Flags().StringVarP(&API, "api", "a", "", "Used to specify the API URL")
	nodeDetailCmd.Flags().IntVar(&port, "port", port, "Used to specify the ironic port address")
	nodeDetailCmd.Flags().IntVar(&n_threads, "threads", n_threads, "Used to specify the number of worker threads to be created at the runtime")
	nodeDetailCmd.Flags().IntVar(&timer, "timer", timer, "Used to specify the number of seconds for which the program should wait to execute worker threads")
	nodeDetailCmd.Flags().IntVar(&repeat, "repeat", repeat, "Used to specify the number of iterations for which the API hits should happen")
	nodeDetailCmd.Flags().BoolVar(&csv, "csv", csv, "Used to specify the output should be in CSV file")

	nodeDetailCmd.MarkFlagRequired("request")
	nodeDetailCmd.MarkFlagRequired("api")
}

func getRequest(URL string, iteration_count int, count int) {
	t1 := time.Now()
	_, err := http.Get(URL)
	t2 := time.Now()
	checkError("looks like an error has occur", err)
	time_diff := t2.Sub(t1)
	fmt.Println("Request Count: ", count, " ,Iteration: ", iteration_count, " ,Latency: ", time_diff.Seconds())
	metrics := RecordRequestData{}
	metrics.Iteration = iteration_count
	metrics.Thread = count
	metrics.Latency = time_diff.Seconds()
	RequestMetricsArray = append(RequestMetricsArray, metrics)
}

func writeMetricstoCSV() {
	convertedMetrics := convertMetricstoString()
	file, err := os.Create("results.csv")
	checkError("Err in creating results.csv", err)
	defer file.Close() //At the end of the execution of this function close this File
	writer := csv.NewWriter(file)
	defer writer.Flush()
	headers := []string{"Iteration", "Thread", "Latency(in Seconds)"}
	err = writer.Write(headers)
	checkError("Error in writing metrics of CSV", err)
	for _, metrics := range convertedMetrics {
		err = writer.Write(metrics)
		checkError("Error in writing metrics of CSV", err)
	}
}
