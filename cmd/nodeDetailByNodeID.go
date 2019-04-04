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
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// nodeDetailByNodeIDCmd represents the nodeDetailByNodeID command
var nodeDetailByNodeIDCmd = &cobra.Command{
	Use:   "nodeDetailByNodeID",
	Short: "Run benchmark for Baremetal Information by using Node IDs specifically",
	Long: `nodeDetailByNodeIDCmd Command is used to hit the node detailing APIs,  which provides complete details for the IRONIC provisioned bare metals.
Following  is therequest that is handled by this command:

GET	/v1/nodes/<node_identity>	Show complete node details `,
	Run: func(cmd *cobra.Command, args []string) {
		host := cmd.Flag("host").Value.String()
		port := cmd.Flag("port").Value.String()
		csv := cmd.Flag("csv").Value.String()
		repeat, err := strconv.Atoi(cmd.Flag("repeat").Value.String())
		checkError("Error in coverting String to Int in nodeDetailCmd for repeat variable", err)
		threadsPerRequest, err := strconv.Atoi(cmd.Flag("threads").Value.String())
		checkError("Error in coverting String to Int in nodeDetailCmd for threads variable", err)
		timer, err := time.ParseDuration((cmd.Flag("timer").Value.String()) + "s") // "s" represents 'seconds' unit of time, by appending "s" to the timer string, will create a time duration object which represents "timer" seconds
		checkError("Error in coverting String to time Duration in nodeDetailCmd for timer", err)

		initiateTestingNodeDetailByNodeID(host, port, repeat, threadsPerRequest, timer)

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

type node struct {
	InstanceUUID   string `json:"instance_uuid"`
	Maintainence   bool   `json:"maintenance"`
	UUID           string `json:"uuid"`
	ProvisionState string `json:"provision_state"`
	PowerState     string `json:"power_state"`
	Links          []map[string]string
}

func getNodeUUIDs(server string) []string {
	var ListUUIDs []string
	response, err := http.Get(server)
	checkError("Error in GET Request to Ironic Server while retrieving Nodes: ", err)
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(response.Body)
		checkError("Error in reading HTTP Response", err2)
		var jsonResp map[string][]node
		err3 := json.Unmarshal(bodyBytes, &jsonResp)
		checkError("Error in Unmarshalling JSON Response", err3)
		// fmt.Println(jsonResp["nodes"])
		for _, item := range jsonResp["nodes"] {
			ListUUIDs = append(ListUUIDs, item.UUID)
		}
	}
	return ListUUIDs
}

func init() {
	runBenchmarkCmd.AddCommand(nodeDetailByNodeIDCmd)

	var repeat int = 1
	var port int = 6385
	var n_threads int = 1
	var timer int = 5
	var csv bool = false

	nodeDetailByNodeIDCmd.Flags().String("host", "localhost", "Used to specify the ironic server address")
	nodeDetailByNodeIDCmd.Flags().IntVar(&port, "port", port, "Used to specify the ironic port address")
	nodeDetailByNodeIDCmd.Flags().IntVar(&n_threads, "threads", n_threads, "Used to specify the number of worker threads per node request to be created at the runtime")
	nodeDetailByNodeIDCmd.Flags().IntVar(&timer, "timer", timer, "Used to specify the number of seconds for which the program should wait to execute worker threads")
	nodeDetailByNodeIDCmd.Flags().IntVar(&repeat, "repeat", repeat, "Used to specify the number of iterations to repeat the test")
	nodeDetailByNodeIDCmd.Flags().BoolVar(&csv, "csv", csv, "Used to specify the output should be in CSV file")
}

func initiateTestingNodeDetailByNodeID(host string, port string, repeat int, threads int, timer time.Duration) {
	serverAddress := "http://" + host + ":" + port + "/v1/nodes/"
	UUIDs := getNodeUUIDs(serverAddress)
	for iteration_count := 1; iteration_count <= repeat; iteration_count++ {
		for _, node_id := range UUIDs {
			nodeAPI := "http://" + host + ":" + port + "/v1/nodes/" + node_id
			for i := 1; i <= threads; i++ {
				go getRequest(nodeAPI, iteration_count, i)
			}
		}
		time.Sleep(timer)
	}
}
