// Copyright © 2019 NAME HERE chawlanikhil24@gmail.com
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
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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
GET	/v1/nodes/{node_identity}	Show Node Detailed
POST	/v1/nodes			Create Node
PATCH	/v1/nodes/{node_identity}	Update Node
DELETE	/v1/nodes/{node_identity}	Delete Node`,

	Run: func(cmd *cobra.Command, args []string) {
		api := cmd.Flag("api").Value.String()
		host := cmd.Flag("host").Value.String()
		port := cmd.Flag("port").Value.String()
		request := cmd.Flag("request").Value.String()
		repeat, err := strconv.Atoi(cmd.Flag("repeat").Value.String())
		if err != nil {
			fmt.Println("Error in coverting String to Int in nodeDetailCmd: ", "repeat", err)
		}
		threads, err := strconv.Atoi(cmd.Flag("threads").Value.String())
		if err != nil {
			fmt.Println("Error in coverting String to Int in nodeDetailCmd: ", "threads", err)
		}
		timer, err := time.ParseDuration((cmd.Flag("timer").Value.String()) + "s") // "s" represents 'seconds' unit of time, by appending "s" to the timer string, will create a time duration object which represents "timer" seconds
		if err != nil {
			fmt.Println("Error in coverting String to time Duration in nodeDetailCmd: ", "timer", err)
		}

		var ironic_server string = "http://" + host + ":" + port + api
		// channel := make(chan RecordRequestData)
		if request == "GET" {
			for iterations := 1; iterations <= repeat; iterations++ {
				for i := 1; i <= threads; i++ {
					// go getRequest(ironic_server, channel)
					// fmt.Println("CHANNEL: ", <-channel)
					go getRequest(ironic_server, iterations, i)
					// close(channel)
				}
				time.Sleep(timer)
				fmt.Println("SleepOver")
			}
		} else {
			for i := 1; i <= threads; i++ {
				jsonBody := jsonBodyForPostRequest{
					Name:            "test_node_dynamic",
					Driver:          "ipmi",
					Power_Interface: "ipmitool",
					Resource_Class:  "bm-large",
					DriverInfo: IPMICredentials{
						IPMI_Username: "ADMIN",
						IPMI_Password: "password",
					},
				}
				body, err := json.Marshal(jsonBody)
				if err != nil {
					panic(err)
				}
				go postRequest(ironic_server, body)
			}
		}
		// time.Sleep(timer)
	},
}

type RecordRequestData struct {
	Iteration int
	Thread    int
	latency   float64
}

type IPMICredentials struct {
	IPMI_Username string `json:"ipmi_username"`
	IPMI_Password string `json:"ipmi_password"`
}

type jsonBodyForPostRequest struct {
	Name            string `json:"name"`
	Driver          string `json:"driver"`
	Power_Interface string `json:"power_interface"`
	Resource_Class  string `json:"resource_class"`
	DriverInfo      IPMICredentials
}

func getRequest(URL string, iteration_count int, count int) {
	t1 := time.Now()
	_, err := http.Get(URL)
	t2 := time.Now()
	if err != nil {
		fmt.Println("looks like an error has occr", err)
		// return 0.0000000000, err
	} else {
		time_diff := t2.Sub(t1)
		fmt.Println("Request Count: ", count, " ,Iteration: ", iteration_count, " ,Latency: ", time_diff.Seconds())
		// return time_diff.Seconds(), nil
	}
}

// func getRequest(URL string, channel chan<- string) (float64, error) {
// 	t1 := time.Now()
// 	res, err := http.Get(URL)
// 	t2 := time.Now()
// 	if err != nil {
// 		fmt.Println("looks like an error has occr", err)
// 		return 0.0000000000, err
// 	} else {
// 		time_diff := t2.Sub(t1)
// 		response_body, err := ioutil.ReadAll(res.Body)
// 		if err != nil {
// 			panic(err)
// 		}
// 		channel <- string(response_body)
// 		fmt.Println(time_diff.Seconds())
// 		res.Body.Close()
// 		return time_diff.Seconds(), nil
// 	}
// }

func postRequest(URL string, body []byte) (float64, error) {

	t1 := time.Now()
	res, err := http.NewRequest("POST", URL, bytes.NewBuffer(body))
	t2 := time.Now()
	if err != nil {
		fmt.Println("looks like an error has occur", err)
		return 0.0000000000, err
	}
	fmt.Println(res.Response)
	time_diff := t2.Sub(t1)
	fmt.Println(time_diff.Seconds(), time_diff.Nanoseconds())
	return time_diff.Seconds(), nil
}

func init() {
	runBenchmarkCmd.AddCommand(nodeDetailCmd)

	var request_type string
	var API string
	var repeat int = 1
	var port int = 6385
	var n_threads int = 10
	var timer int = 5
	nodeDetailCmd.Flags().String("host", "localhost", "Used to specify the ironic server address")
	nodeDetailCmd.Flags().StringVarP(&request_type, "request", "r", "", "Used to specify the request type that should hit ironic server API(GET, POST, PATCH, DELETE etc.)")
	nodeDetailCmd.Flags().StringVarP(&API, "api", "a", "", "Used to specify the API URL")
	nodeDetailCmd.Flags().IntVar(&port, "port", port, "Used to specify the ironic port address")
	nodeDetailCmd.Flags().IntVar(&n_threads, "threads", n_threads, "Used to specify the number of worker threads to be created at the runtime")
	nodeDetailCmd.Flags().IntVar(&timer, "timer", timer, "Used to specify the number of seconds for which the program should wait to execute worker threads")
	nodeDetailCmd.Flags().IntVar(&repeat, "repeat", repeat, "Used to specify the number of iterations for which the API hits should happen")

	nodeDetailCmd.MarkFlagRequired("request")
	nodeDetailCmd.MarkFlagRequired("api")
}
