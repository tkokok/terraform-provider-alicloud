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

// UnassociateHaVip invokes the vpc.UnassociateHaVip API synchronously
// api document: https://help.aliyun.com/api/vpc/unassociatehavip.html
func (client *Client) UnassociateHaVip(request *UnassociateHaVipRequest) (response *UnassociateHaVipResponse, err error) {
	response = CreateUnassociateHaVipResponse()
	err = client.DoAction(request, response)
	return
}

// UnassociateHaVipWithChan invokes the vpc.UnassociateHaVip API asynchronously
// api document: https://help.aliyun.com/api/vpc/unassociatehavip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnassociateHaVipWithChan(request *UnassociateHaVipRequest) (<-chan *UnassociateHaVipResponse, <-chan error) {
	responseChan := make(chan *UnassociateHaVipResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnassociateHaVip(request)
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

// UnassociateHaVipWithCallback invokes the vpc.UnassociateHaVip API asynchronously
// api document: https://help.aliyun.com/api/vpc/unassociatehavip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnassociateHaVipWithCallback(request *UnassociateHaVipRequest, callback func(response *UnassociateHaVipResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnassociateHaVipResponse
		var err error
		defer close(result)
		response, err = client.UnassociateHaVip(request)
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

// UnassociateHaVipRequest is the request struct for api UnassociateHaVip
type UnassociateHaVipRequest struct {
	*requests.RpcRequest
	HaVipId              string           `position:"Query" name:"HaVipId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Force                string           `position:"Query" name:"Force"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// UnassociateHaVipResponse is the response struct for api UnassociateHaVip
type UnassociateHaVipResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnassociateHaVipRequest creates a request to invoke UnassociateHaVip API
func CreateUnassociateHaVipRequest() (request *UnassociateHaVipRequest) {
	request = &UnassociateHaVipRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "UnassociateHaVip", "vpc", "openAPI")
	return
}

// CreateUnassociateHaVipResponse creates a response to parse from UnassociateHaVip response
func CreateUnassociateHaVipResponse() (response *UnassociateHaVipResponse) {
	response = &UnassociateHaVipResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
