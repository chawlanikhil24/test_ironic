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

	"github.com/spf13/cobra"
)

// runBenchmarkCmd represents the runBenchmark command
var runBenchmarkCmd = &cobra.Command{
	Use:   "runBenchmark",
	Short: "Select the scenario of benchmarking",
	Long: `Following is the list of benchmarks scenarios which are currently available:

	1. nodeDetail: List, Searching, Creating, Updating, and Deleting of bare metal Node resources are done through the /v1/nodes resource. There are also several sub-resources, which allow further actions to be performed on a bare metal Node.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("runBenchmark called")
		fmt.Println("Following are the arguments: ", args)
	},
}

func init() {
	rootCmd.AddCommand(runBenchmarkCmd)
}
