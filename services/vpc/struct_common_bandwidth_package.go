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

// CommonBandwidthPackage is a nested struct in vpc response
type CommonBandwidthPackage struct {
	BandwidthPackageId string                                             `json:"BandwidthPackageId" xml:"BandwidthPackageId"`
	RegionId           string                                             `json:"RegionId" xml:"RegionId"`
	Name               string                                             `json:"Name" xml:"Name"`
	Description        string                                             `json:"Description" xml:"Description"`
	Bandwidth          string                                             `json:"Bandwidth" xml:"Bandwidth"`
	InstanceChargeType string                                             `json:"InstanceChargeType" xml:"InstanceChargeType"`
	InternetChargeType string                                             `json:"InternetChargeType" xml:"InternetChargeType"`
	BusinessStatus     string                                             `json:"BusinessStatus" xml:"BusinessStatus"`
	CreationTime       string                                             `json:"CreationTime" xml:"CreationTime"`
	ExpiredTime        string                                             `json:"ExpiredTime" xml:"ExpiredTime"`
	Status             string                                             `json:"Status" xml:"Status"`
	Ratio              int                                                `json:"Ratio" xml:"Ratio"`
	PublicIpAddresses  PublicIpAddressesInDescribeCommonBandwidthPackages `json:"PublicIpAddresses" xml:"PublicIpAddresses"`
}