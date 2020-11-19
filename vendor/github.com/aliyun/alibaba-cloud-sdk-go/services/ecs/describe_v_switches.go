package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeVSwitches invokes the ecs.DescribeVSwitches API synchronously
func (client *Client) DescribeVSwitches(request *DescribeVSwitchesRequest) (response *DescribeVSwitchesResponse, err error) {
	response = CreateDescribeVSwitchesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVSwitchesWithChan invokes the ecs.DescribeVSwitches API asynchronously
func (client *Client) DescribeVSwitchesWithChan(request *DescribeVSwitchesRequest) (<-chan *DescribeVSwitchesResponse, <-chan error) {
	responseChan := make(chan *DescribeVSwitchesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVSwitches(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeVSwitchesWithCallback invokes the ecs.DescribeVSwitches API asynchronously
func (client *Client) DescribeVSwitchesWithCallback(request *DescribeVSwitchesRequest, callback func(response *DescribeVSwitchesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVSwitchesResponse
		var err error
		defer close(result)
		response, err = client.DescribeVSwitches(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeVSwitchesRequest is the request struct for api DescribeVSwitches
type DescribeVSwitchesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	IsDefault            requests.Boolean `position:"Query" name:"IsDefault"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	VpcId                string           `position:"Query" name:"VpcId"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// DescribeVSwitchesResponse is the response struct for api DescribeVSwitches
type DescribeVSwitchesResponse struct {
	*responses.BaseResponse
	RequestId  string    `json:"RequestId" xml:"RequestId"`
	TotalCount int       `json:"TotalCount" xml:"TotalCount"`
	PageNumber int       `json:"PageNumber" xml:"PageNumber"`
	PageSize   int       `json:"PageSize" xml:"PageSize"`
	VSwitches  VSwitches `json:"VSwitches" xml:"VSwitches"`
}

// CreateDescribeVSwitchesRequest creates a request to invoke DescribeVSwitches API
func CreateDescribeVSwitchesRequest() (request *DescribeVSwitchesRequest) {
	request = &DescribeVSwitchesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeVSwitches", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeVSwitchesResponse creates a response to parse from DescribeVSwitches response
func CreateDescribeVSwitchesResponse() (response *DescribeVSwitchesResponse) {
	response = &DescribeVSwitchesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
