package slb

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

// CreateVServerGroup invokes the slb.CreateVServerGroup API synchronously
// api document: https://help.aliyun.com/api/slb/createvservergroup.html
func (client *Client) CreateVServerGroup(request *CreateVServerGroupRequest) (response *CreateVServerGroupResponse, err error) {
	response = CreateCreateVServerGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVServerGroupWithChan invokes the slb.CreateVServerGroup API asynchronously
// api document: https://help.aliyun.com/api/slb/createvservergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateVServerGroupWithChan(request *CreateVServerGroupRequest) (<-chan *CreateVServerGroupResponse, <-chan error) {
	responseChan := make(chan *CreateVServerGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVServerGroup(request)
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

// CreateVServerGroupWithCallback invokes the slb.CreateVServerGroup API asynchronously
// api document: https://help.aliyun.com/api/slb/createvservergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateVServerGroupWithCallback(request *CreateVServerGroupRequest, callback func(response *CreateVServerGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVServerGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateVServerGroup(request)
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

// CreateVServerGroupRequest is the request struct for api CreateVServerGroup
type CreateVServerGroupRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	Tags                 string           `position:"Query" name:"Tags"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
	VServerGroupName     string           `position:"Query" name:"VServerGroupName"`
	BackendServers       string           `position:"Query" name:"BackendServers"`
}

// CreateVServerGroupResponse is the response struct for api CreateVServerGroup
type CreateVServerGroupResponse struct {
	*responses.BaseResponse
	RequestId      string                             `json:"RequestId" xml:"RequestId"`
	VServerGroupId string                             `json:"VServerGroupId" xml:"VServerGroupId"`
	BackendServers BackendServersInCreateVServerGroup `json:"BackendServers" xml:"BackendServers"`
}

// CreateCreateVServerGroupRequest creates a request to invoke CreateVServerGroup API
func CreateCreateVServerGroupRequest() (request *CreateVServerGroupRequest) {
	request = &CreateVServerGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "CreateVServerGroup", "slb", "openAPI")
	return
}

// CreateCreateVServerGroupResponse creates a response to parse from CreateVServerGroup response
func CreateCreateVServerGroupResponse() (response *CreateVServerGroupResponse) {
	response = &CreateVServerGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}