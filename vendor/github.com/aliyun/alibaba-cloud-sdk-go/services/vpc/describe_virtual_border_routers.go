package vpc

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

// DescribeVirtualBorderRouters invokes the vpc.DescribeVirtualBorderRouters API synchronously
// api document: https://help.aliyun.com/api/vpc/describevirtualborderrouters.html
func (client *Client) DescribeVirtualBorderRouters(request *DescribeVirtualBorderRoutersRequest) (response *DescribeVirtualBorderRoutersResponse, err error) {
	response = CreateDescribeVirtualBorderRoutersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVirtualBorderRoutersWithChan invokes the vpc.DescribeVirtualBorderRouters API asynchronously
// api document: https://help.aliyun.com/api/vpc/describevirtualborderrouters.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVirtualBorderRoutersWithChan(request *DescribeVirtualBorderRoutersRequest) (<-chan *DescribeVirtualBorderRoutersResponse, <-chan error) {
	responseChan := make(chan *DescribeVirtualBorderRoutersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVirtualBorderRouters(request)
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

// DescribeVirtualBorderRoutersWithCallback invokes the vpc.DescribeVirtualBorderRouters API asynchronously
// api document: https://help.aliyun.com/api/vpc/describevirtualborderrouters.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVirtualBorderRoutersWithCallback(request *DescribeVirtualBorderRoutersRequest, callback func(response *DescribeVirtualBorderRoutersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVirtualBorderRoutersResponse
		var err error
		defer close(result)
		response, err = client.DescribeVirtualBorderRouters(request)
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

// DescribeVirtualBorderRoutersRequest is the request struct for api DescribeVirtualBorderRouters
type DescribeVirtualBorderRoutersRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer                      `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string                                `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer                      `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer                      `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer                      `position:"Query" name:"PageSize"`
	Filter               *[]DescribeVirtualBorderRoutersFilter `position:"Query" name:"Filter"  type:"Repeated"`
}

// DescribeVirtualBorderRoutersFilter is a repeated param struct in DescribeVirtualBorderRoutersRequest
type DescribeVirtualBorderRoutersFilter struct {
	Key   string    `name:"Key"`
	Value *[]string `name:"Value" type:"Repeated"`
}

// DescribeVirtualBorderRoutersResponse is the response struct for api DescribeVirtualBorderRouters
type DescribeVirtualBorderRoutersResponse struct {
	*responses.BaseResponse
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	PageNumber             int                    `json:"PageNumber" xml:"PageNumber"`
	PageSize               int                    `json:"PageSize" xml:"PageSize"`
	TotalCount             int                    `json:"TotalCount" xml:"TotalCount"`
	VirtualBorderRouterSet VirtualBorderRouterSet `json:"VirtualBorderRouterSet" xml:"VirtualBorderRouterSet"`
}

// CreateDescribeVirtualBorderRoutersRequest creates a request to invoke DescribeVirtualBorderRouters API
func CreateDescribeVirtualBorderRoutersRequest() (request *DescribeVirtualBorderRoutersRequest) {
	request = &DescribeVirtualBorderRoutersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVirtualBorderRouters", "vpc", "openAPI")
	return
}

// CreateDescribeVirtualBorderRoutersResponse creates a response to parse from DescribeVirtualBorderRouters response
func CreateDescribeVirtualBorderRoutersResponse() (response *DescribeVirtualBorderRoutersResponse) {
	response = &DescribeVirtualBorderRoutersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
