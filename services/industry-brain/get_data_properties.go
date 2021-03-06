package industry_brain

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

// GetDataProperties invokes the industry_brain.GetDataProperties API synchronously
// api document: https://help.aliyun.com/api/industry-brain/getdataproperties.html
func (client *Client) GetDataProperties(request *GetDataPropertiesRequest) (response *GetDataPropertiesResponse, err error) {
	response = CreateGetDataPropertiesResponse()
	err = client.DoAction(request, response)
	return
}

// GetDataPropertiesWithChan invokes the industry_brain.GetDataProperties API asynchronously
// api document: https://help.aliyun.com/api/industry-brain/getdataproperties.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDataPropertiesWithChan(request *GetDataPropertiesRequest) (<-chan *GetDataPropertiesResponse, <-chan error) {
	responseChan := make(chan *GetDataPropertiesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDataProperties(request)
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

// GetDataPropertiesWithCallback invokes the industry_brain.GetDataProperties API asynchronously
// api document: https://help.aliyun.com/api/industry-brain/getdataproperties.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDataPropertiesWithCallback(request *GetDataPropertiesRequest, callback func(response *GetDataPropertiesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDataPropertiesResponse
		var err error
		defer close(result)
		response, err = client.GetDataProperties(request)
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

// GetDataPropertiesRequest is the request struct for api GetDataProperties
type GetDataPropertiesRequest struct {
	*requests.RpcRequest
	ServiceId string `position:"Query" name:"ServiceId"`
}

// GetDataPropertiesResponse is the response struct for api GetDataProperties
type GetDataPropertiesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetDataPropertiesRequest creates a request to invoke GetDataProperties API
func CreateGetDataPropertiesRequest() (request *GetDataPropertiesRequest) {
	request = &GetDataPropertiesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("industry-brain", "2018-07-12", "GetDataProperties", "", "")
	return
}

// CreateGetDataPropertiesResponse creates a response to parse from GetDataProperties response
func CreateGetDataPropertiesResponse() (response *GetDataPropertiesResponse) {
	response = &GetDataPropertiesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
