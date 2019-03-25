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
		threads, err := strconv.Atoi(cmd.Flag("threads").Value.String())
		if err != nil {
			fmt.Println("Error in coverting String to Int in nodeDetailCmd: ", "threads", err)
		}
		timer, err := time.ParseDuration((cmd.Flag("timer").Value.String()) + "s") // "s" represents 'seconds' unit of time, by appending "s" to the timer string, will create a time duration object which represents "timer" seconds
		if err != nil {
			fmt.Println("Error in coverting String to time Duration in nodeDetailCmd: ", "timer", err)
		}

		var ironic_server string = "http://" + host + ":" + port + api

		fmt.Println(request, threads, timer)
		fmt.Println(ironic_server)
		if request == "GET" {
			for i := 1; i <= threads; i++ {
				go getRequest(ironic_server)
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
		time.Sleep(timer)
	},
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

func getRequest(URL string) (float64, error) {
	t1 := time.Now()
	res, err := http.Get(URL)
	t2 := time.Now()
	if err != nil {
		fmt.Println("looks like an error has occr", err)
		return 0.0000000000, err
	} else {
		time_diff := t2.Sub(t1)
		fmt.Println(res.StatusCode, time_diff.Seconds())
		return time_diff.Seconds(), nil
	}
}

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

	var port int = 6385
	var n_threads int = 10
	var timer int = 5
	var request_type string
	var API string
	nodeDetailCmd.Flags().String("host", "localhost", "Used to specify the ironic server address")
	nodeDetailCmd.Flags().StringVarP(&request_type, "request", "r", "", "Used to specify the request type that should hit ironic server API(GET, POST, PATCH, DELETE etc.)")
	nodeDetailCmd.Flags().StringVarP(&API, "api", "a", "", "Used to specify the API URL")
	nodeDetailCmd.Flags().IntVar(&port, "port", port, "Used to specify the ironic port address")
	nodeDetailCmd.Flags().IntVar(&n_threads, "threads", n_threads, "Used to specify the number of worker threads to be created at the runtime")
	nodeDetailCmd.Flags().IntVar(&timer, "timer", timer, "Used to specify the number of seconds for which the program should wait to execute worker threads")

	nodeDetailCmd.MarkFlagRequired("request")
	nodeDetailCmd.MarkFlagRequired("api")
}

// body = {
// 	"name": "test_node_dynamic",
// 	"driver": "ipmi",
// 	"driver_info": {
// 			"ipmi_username": "ADMIN",
// 			"ipmi_password": "password"
// 	},
// 	"power_interface": "ipmitool",
// 	"resource_class": "bm-large"
// }
