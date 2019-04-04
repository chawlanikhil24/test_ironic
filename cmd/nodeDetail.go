// Copyright © 2019 Nikhil Chawla chawlanikhil24@gmail.com
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
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// nodeDetailCmd represents the nodeDetail command
var nodeDetailCmd = &cobra.Command{
	Use:   "nodeDetail",
	Short: "Run benchmark for Baremetal listing detailing APIs as a Summary",
	Long: `nodeDetail Command is used to hit the APIs which provides details for the IRONIC provisioned bare metals.
	List, Searching, Creating, Updating, and Deleting of bare metal Node resources are done through the /v1/nodes resource.
	There are also several sub-resources, which allow further actions to be performed on a bare metal Node.

	Following are the type of requests that are handled by this command:


GET	/v1/nodes			List Nodes --API -> 1
GET	/v1/nodes/detail		List Nodes Detailed  --API -> 2`,

	Run: func(cmd *cobra.Command, args []string) {
		api_choice := cmd.Flag("api").Value.String()
		host := cmd.Flag("host").Value.String()
		port := cmd.Flag("port").Value.String()
		csv := cmd.Flag("csv").Value.String()
		repeat, err := strconv.Atoi(cmd.Flag("repeat").Value.String())
		checkError("Error in coverting String to Int in nodeDetailCmd for repeat variable", err)
		threads, err := strconv.Atoi(cmd.Flag("threads").Value.String())
		checkError("Error in coverting String to Int in nodeDetailCmd for threads variable", err)
		timer, err := time.ParseDuration((cmd.Flag("timer").Value.String()) + "s") // "s" represents 'seconds' unit of time, by appending "s" to the timer string, will create a time duration object which represents "timer" seconds
		checkError("Error in coverting String to time Duration in nodeDetailCmd for timer", err)

		var api string
		switch api_choice {
		case "1":
			api = "http://" + host + ":" + port + "/v1/nodes/"
			initiateTestingNodeDetail(api, repeat, threads, timer)
		case "2":
			api = "http://" + host + ":" + port + "/v1/nodes/details"
			initiateTestingNodeDetail(api, repeat, threads, timer)
		default:
			log.Fatal("Wrong API choice. Check nodeDetailCmd --help for details of API Choices")
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

func init() {
	runBenchmarkCmd.AddCommand(nodeDetailCmd)

	var API int
	var repeat int = 1
	var port int = 6385
	var n_threads int = 10
	var timer int = 5
	var csv bool = false

	nodeDetailCmd.Flags().String("host", "localhost", "Used to specify the ironic server address")
	nodeDetailCmd.Flags().IntVarP(&API, "api", "a", 0, "Used to specify the API URL you want to test. For commplete list of URLs available with nodeDetail command, check nodeDetail --help")
	nodeDetailCmd.Flags().IntVar(&port, "port", port, "Used to specify the ironic port address")
	nodeDetailCmd.Flags().IntVar(&n_threads, "threads", n_threads, "Used to specify the number of worker threads to be created at the runtime")
	nodeDetailCmd.Flags().IntVar(&timer, "timer", timer, "Used to specify the number of seconds for which the program should wait to execute worker threads")
	nodeDetailCmd.Flags().IntVar(&repeat, "repeat", repeat, "Used to specify the number of iterations to repeat the test")
	nodeDetailCmd.Flags().BoolVar(&csv, "csv", csv, "Used to specify the output should be in CSV file")

	nodeDetailCmd.MarkFlagRequired("api")
}

func initiateTestingNodeDetail(API string, repeat int, threads int, timer time.Duration) {
	for iterations := 1; iterations <= repeat; iterations++ {
		for i := 1; i <= threads; i++ {
			go getRequest(API, iterations, i)
		}
		time.Sleep(timer)
	}
	fmt.Println("Test Over")
}
