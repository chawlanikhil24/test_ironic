# test_ironic
This project aims to do performance benchmarking of Openstack Ironic, just like Rally, but in non-Openstack environment.

# Installation

* Golang must be installed and GOPATH must be set.
* To Install golang on your choice of OS, follow this guide https://golang.org/doc/install
* export GOPATH=$HOME/go/
* go get github.com/chawlanikhil24/test_ironic
* cd $HOME/go/src/chawlanikhil24/test_ironic
* go build .
* sudo mv test_ironic /usr/local/bin/

# How to Run the tests ?

* Basic Test
```shell
$test_ironic runBenchmark nodeDetail --api 1 --host localhost --port 6385
```
```shell
Request Count:  4  ,Iteration:  1  ,Latency:  0.135166189 URL /v1/nodes/
Request Count:  2  ,Iteration:  1  ,Latency:  0.138374844 URL /v1/nodes/
Request Count:  10  ,Iteration:  1  ,Latency:  0.141671513 URL /v1/nodes/
Request Count:  8  ,Iteration:  1  ,Latency:  0.143757008 URL /v1/nodes/
Request Count:  3  ,Iteration:  1  ,Latency:  0.144646097 URL /v1/nodes/
Request Count:  7  ,Iteration:  1  ,Latency:  0.152624421 URL /v1/nodes/
Request Count:  1  ,Iteration:  1  ,Latency:  0.157211707 URL /v1/nodes/
Request Count:  5  ,Iteration:  1  ,Latency:  0.208825654 URL /v1/nodes/
Request Count:  9  ,Iteration:  1  ,Latency:  0.261059241 URL /v1/nodes/
Request Count:  6  ,Iteration:  1  ,Latency:  0.318185217 URL /v1/nodes/
Test Over
```
* Increase the concurrent threads to stress the API
```shell
$test_ironic runBenchmark nodeDetail --api 1 --host localhost --port 6385 --threads 20
```
```shell
Request Count:  14  ,Iteration:  1  ,Latency:  0.146125595 URL /v1/nodes/
Request Count:  2  ,Iteration:  1  ,Latency:  0.151007213 URL /v1/nodes/
Request Count:  20  ,Iteration:  1  ,Latency:  0.152508637 URL /v1/nodes/
Request Count:  3  ,Iteration:  1  ,Latency:  0.155545361 URL /v1/nodes/
Request Count:  19  ,Iteration:  1  ,Latency:  0.158863992 URL /v1/nodes/
Request Count:  16  ,Iteration:  1  ,Latency:  0.163825582 URL /v1/nodes/
Request Count:  5  ,Iteration:  1  ,Latency:  0.168267048 URL /v1/nodes/
Request Count:  17  ,Iteration:  1  ,Latency:  0.17936613 URL /v1/nodes/
Request Count:  12  ,Iteration:  1  ,Latency:  0.182062817 URL /v1/nodes/
Request Count:  4  ,Iteration:  1  ,Latency:  0.19772769 URL /v1/nodes/
Request Count:  9  ,Iteration:  1  ,Latency:  0.201890178 URL /v1/nodes/
Request Count:  1  ,Iteration:  1  ,Latency:  0.20790699 URL /v1/nodes/
Request Count:  15  ,Iteration:  1  ,Latency:  0.235892375 URL /v1/nodes/
Request Count:  11  ,Iteration:  1  ,Latency:  0.241949631 URL /v1/nodes/
Request Count:  7  ,Iteration:  1  ,Latency:  0.2451876 URL /v1/nodes/
Request Count:  18  ,Iteration:  1  ,Latency:  0.245899378 URL /v1/nodes/
Request Count:  13  ,Iteration:  1  ,Latency:  0.258001168 URL /v1/nodes/
Request Count:  10  ,Iteration:  1  ,Latency:  0.269836494 URL /v1/nodes/
Request Count:  8  ,Iteration:  1  ,Latency:  0.277604564 URL /v1/nodes/
Request Count:  6  ,Iteration:  1  ,Latency:  0.326116188 URL /v1/nodes/
Test Over
```
* If you need to repeat the test then you also add the count for Iterations
```shell
$test_ironic runBenchmark nodeDetail --api 1 --host localhost --port 6385 --threads 3 --repeat 3
```
```shell
Request Count:  2  ,Iteration:  1  ,Latency:  0.136736491 URL /v1/nodes/
Request Count:  3  ,Iteration:  1  ,Latency:  0.146211329 URL /v1/nodes/
Request Count:  1  ,Iteration:  1  ,Latency:  0.150101849 URL /v1/nodes/
Request Count:  1  ,Iteration:  2  ,Latency:  0.136395718 URL /v1/nodes/
Request Count:  3  ,Iteration:  2  ,Latency:  0.139686455 URL /v1/nodes/
Request Count:  2  ,Iteration:  2  ,Latency:  0.239952914 URL /v1/nodes/
Request Count:  3  ,Iteration:  3  ,Latency:  0.13201399 URL /v1/nodes/
Request Count:  1  ,Iteration:  3  ,Latency:  0.135316495 URL /v1/nodes/
Request Count:  2  ,Iteration:  3  ,Latency:  0.139297851 URL /v1/nodes/
Test Over
```
* If you want to save the results in a CSV file, set the "csv" flag to "true". After the test, you can find a csv file named "results.csv" in the current directory.
```shell
test_ironic runBenchmark nodeDetail --api 2 --host localhost --port 6385 --threads 3 --csv true
```
```shell
Request Count:  1  ,Iteration:  1  ,Latency:  0.015317947 URL /v1/nodes/details
Request Count:  3  ,Iteration:  1  ,Latency:  0.01585443 URL /v1/nodes/details
Request Count:  2  ,Iteration:  1  ,Latency:  0.016079593 URL /v1/nodes/details
Test Over
```
* Sample test on node details using node IDs is executed the following way:
```shell
$test_ironic runBenchmark nodeDetailByNodeID --host localhost --port 6385 --threads 1
```
```shell
Request Count:  1  ,Iteration:  1  ,Latency:  0.024087921 URL /v1/nodes/88778d85-1a0f-411d-b561-aaa022e0ccf3
Request Count:  1  ,Iteration:  1  ,Latency:  0.023576536 URL /v1/nodes/d27e9136-7dd1-42f8-8dc2-49f80cf96c79
Request Count:  1  ,Iteration:  1  ,Latency:  0.040186732 URL /v1/nodes/449b51bc-71b2-45a3-a6d0-58afceafd9e8
Request Count:  1  ,Iteration:  1  ,Latency:  0.04002651 URL /v1/nodes/8e75f9fe-8a9d-4207-8b30-916b1209736c
Request Count:  1  ,Iteration:  1  ,Latency:  0.040262503 URL /v1/nodes/48e089c8-a5df-4dbe-bf3c-d94a92b0bd18
Request Count:  1  ,Iteration:  1  ,Latency:  0.040366061 URL /v1/nodes/d3665405-5991-4ea5-8614-d00cf4eba080
Request Count:  1  ,Iteration:  1  ,Latency:  0.042277041 URL /v1/nodes/0d8c9345-5ed8-4b5f-8eb8-9d02953a6bcf
Request Count:  1  ,Iteration:  1  ,Latency:  0.043447799 URL /v1/nodes/b110c1ed-c5d3-4584-a11f-9a19c613370d
Request Count:  1  ,Iteration:  1  ,Latency:  0.043578547 URL /v1/nodes/393e2451-84e5-4cb7-9604-c52233ce1e99
Request Count:  1  ,Iteration:  1  ,Latency:  0.043904937 URL /v1/nodes/227ed679-2e8c-4d69-a42f-421fc952076d
Request Count:  1  ,Iteration:  1  ,Latency:  0.056136588 URL /v1/nodes/c990308e-680d-4449-8931-acc956adb3fb
Request Count:  1  ,Iteration:  1  ,Latency:  0.057517573 URL /v1/nodes/c404f72d-1d17-4a8d-9de3-19eb7b0bae5d
Request Count:  1  ,Iteration:  1  ,Latency:  0.058327982 URL /v1/nodes/a1bfe9f2-aa3b-4bb5-b6b8-e4f626b0bce2
Request Count:  1  ,Iteration:  1  ,Latency:  0.062722994 URL /v1/nodes/0385f0f7-adc8-442f-b1ef-4abdf83c9741
Request Count:  1  ,Iteration:  1  ,Latency:  0.072984908 URL /v1/nodes/93bd4d5a-354a-4dfe-b9aa-8094887b3c93
Request Count:  1  ,Iteration:  1  ,Latency:  0.079884761 URL /v1/nodes/834fea2e-a578-4475-9b34-cce6f02d5f94
```
* Concurrent request count on node details using node IDs can be increased by the following way:
```shell
$test_ironic runBenchmark nodeDetailByNodeID --host localhost --port 6385 --threads 2
```
```shell
Request Count:  2  ,Iteration:  1  ,Latency:  0.016597804 URL /v1/nodes/88778d85-1a0f-411d-b561-aaa022e0ccf3
Request Count:  1  ,Iteration:  1  ,Latency:  0.036245964 URL /v1/nodes/0d8c9345-5ed8-4b5f-8eb8-9d02953a6bcf
Request Count:  2  ,Iteration:  1  ,Latency:  0.040291379 URL /v1/nodes/c404f72d-1d17-4a8d-9de3-19eb7b0bae5d
Request Count:  1  ,Iteration:  1  ,Latency:  0.041434361 URL /v1/nodes/d3665405-5991-4ea5-8614-d00cf4eba080
Request Count:  1  ,Iteration:  1  ,Latency:  0.042619265 URL /v1/nodes/c990308e-680d-4449-8931-acc956adb3fb
Request Count:  2  ,Iteration:  1  ,Latency:  0.043260403 URL /v1/nodes/93bd4d5a-354a-4dfe-b9aa-8094887b3c93
Request Count:  1  ,Iteration:  1  ,Latency:  0.043377293 URL /v1/nodes/8e75f9fe-8a9d-4207-8b30-916b1209736c
Request Count:  2  ,Iteration:  1  ,Latency:  0.044661072 URL /v1/nodes/449b51bc-71b2-45a3-a6d0-58afceafd9e8
Request Count:  1  ,Iteration:  1  ,Latency:  0.045447544 URL /v1/nodes/c404f72d-1d17-4a8d-9de3-19eb7b0bae5d
Request Count:  1  ,Iteration:  1  ,Latency:  0.045533261 URL /v1/nodes/449b51bc-71b2-45a3-a6d0-58afceafd9e8
Request Count:  2  ,Iteration:  1  ,Latency:  0.045765135 URL /v1/nodes/393e2451-84e5-4cb7-9604-c52233ce1e99
Request Count:  2  ,Iteration:  1  ,Latency:  0.045564276 URL /v1/nodes/0385f0f7-adc8-442f-b1ef-4abdf83c9741
Request Count:  1  ,Iteration:  1  ,Latency:  0.052899285 URL /v1/nodes/a1bfe9f2-aa3b-4bb5-b6b8-e4f626b0bce2
Request Count:  1  ,Iteration:  1  ,Latency:  0.052515124 URL /v1/nodes/0385f0f7-adc8-442f-b1ef-4abdf83c9741
Request Count:  2  ,Iteration:  1  ,Latency:  0.054854017 URL /v1/nodes/d3665405-5991-4ea5-8614-d00cf4eba080
Request Count:  1  ,Iteration:  1  ,Latency:  0.055533025 URL /v1/nodes/48e089c8-a5df-4dbe-bf3c-d94a92b0bd18
Request Count:  2  ,Iteration:  1  ,Latency:  0.055668465 URL /v1/nodes/d27e9136-7dd1-42f8-8dc2-49f80cf96c79
Request Count:  1  ,Iteration:  1  ,Latency:  0.055889696 URL /v1/nodes/88778d85-1a0f-411d-b561-aaa022e0ccf3
Request Count:  1  ,Iteration:  1  ,Latency:  0.059093759 URL /v1/nodes/393e2451-84e5-4cb7-9604-c52233ce1e99
Request Count:  2  ,Iteration:  1  ,Latency:  0.059949655 URL /v1/nodes/a1bfe9f2-aa3b-4bb5-b6b8-e4f626b0bce2
Request Count:  2  ,Iteration:  1  ,Latency:  0.062079402 URL /v1/nodes/0d8c9345-5ed8-4b5f-8eb8-9d02953a6bcf
Request Count:  2  ,Iteration:  1  ,Latency:  0.061231926 URL /v1/nodes/227ed679-2e8c-4d69-a42f-421fc952076d
Request Count:  2  ,Iteration:  1  ,Latency:  0.062429513 URL /v1/nodes/b110c1ed-c5d3-4584-a11f-9a19c613370d
Request Count:  2  ,Iteration:  1  ,Latency:  0.06332732 URL /v1/nodes/8e75f9fe-8a9d-4207-8b30-916b1209736c
Request Count:  1  ,Iteration:  1  ,Latency:  0.067519355 URL /v1/nodes/227ed679-2e8c-4d69-a42f-421fc952076d
Request Count:  2  ,Iteration:  1  ,Latency:  0.068834299 URL /v1/nodes/c990308e-680d-4449-8931-acc956adb3fb
Request Count:  1  ,Iteration:  1  ,Latency:  0.068628175 URL /v1/nodes/93bd4d5a-354a-4dfe-b9aa-8094887b3c93
Request Count:  2  ,Iteration:  1  ,Latency:  0.069744372 URL /v1/nodes/834fea2e-a578-4475-9b34-cce6f02d5f94
Request Count:  2  ,Iteration:  1  ,Latency:  0.074248384 URL /v1/nodes/48e089c8-a5df-4dbe-bf3c-d94a92b0bd18
Request Count:  1  ,Iteration:  1  ,Latency:  0.080184906 URL /v1/nodes/834fea2e-a578-4475-9b34-cce6f02d5f94
Request Count:  1  ,Iteration:  1  ,Latency:  0.082131393 URL /v1/nodes/b110c1ed-c5d3-4584-a11f-9a19c613370d
Request Count:  1  ,Iteration:  1  ,Latency:  0.083529974 URL /v1/nodes/d27e9136-7dd1-42f8-8dc2-49f80cf96c79
```
