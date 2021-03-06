package airec

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

// ListDiversify invokes the airec.ListDiversify API synchronously
// api document: https://help.aliyun.com/api/airec/listdiversify.html
func (client *Client) ListDiversify(request *ListDiversifyRequest) (response *ListDiversifyResponse, err error) {
	response = CreateListDiversifyResponse()
	err = client.DoAction(request, response)
	return
}

// ListDiversifyWithChan invokes the airec.ListDiversify API asynchronously
// api document: https://help.aliyun.com/api/airec/listdiversify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListDiversifyWithChan(request *ListDiversifyRequest) (<-chan *ListDiversifyResponse, <-chan error) {
	responseChan := make(chan *ListDiversifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDiversify(request)
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

// ListDiversifyWithCallback invokes the airec.ListDiversify API asynchronously
// api document: https://help.aliyun.com/api/airec/listdiversify.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListDiversifyWithCallback(request *ListDiversifyRequest, callback func(response *ListDiversifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDiversifyResponse
		var err error
		defer close(result)
		response, err = client.ListDiversify(request)
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

// ListDiversifyRequest is the request struct for api ListDiversify
type ListDiversifyRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"InstanceId"`
}

// ListDiversifyResponse is the response struct for api ListDiversify
type ListDiversifyResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	Code      string       `json:"Code" xml:"Code"`
	Message   string       `json:"Message" xml:"Message"`
	Result    []ResultItem `json:"Result" xml:"Result"`
}

// CreateListDiversifyRequest creates a request to invoke ListDiversify API
func CreateListDiversifyRequest() (request *ListDiversifyRequest) {
	request = &ListDiversifyRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2018-10-12", "ListDiversify", "/openapi/instances/[InstanceId]/diversifies", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListDiversifyResponse creates a response to parse from ListDiversify response
func CreateListDiversifyResponse() (response *ListDiversifyResponse) {
	response = &ListDiversifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
