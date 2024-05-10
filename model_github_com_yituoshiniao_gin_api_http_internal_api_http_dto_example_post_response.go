/*
gin-http API

gin-http服务文档

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse{}

// GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse struct for GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse
type GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse struct {
	// code:  0 成功; 非0失败;
	Code int32                                                             `json:"code"`
	Data GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData `json:"data"`
	// 错误消息
	Msg string `json:"msg"`
	// traceId
	TraceId string `json:"traceId"`
}

type _GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse

// NewGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse instantiates a new GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse(code int32, data GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData, msg string, traceId string) *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse {
	this := GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse{}
	this.Code = code
	this.Data = data
	this.Msg = msg
	this.TraceId = traceId
	return &this
}

// NewGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponseWithDefaults instantiates a new GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponseWithDefaults() *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse {
	this := GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse{}
	return &this
}

// GetCode returns the Code field value
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) SetCode(v int32) {
	o.Code = v
}

// GetData returns the Data field value
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) GetData() GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData {
	if o == nil {
		var ret GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) GetDataOk() (*GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) SetData(v GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoUserPortraitData) {
	o.Data = v
}

// GetMsg returns the Msg field value
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) GetMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msg
}

// GetMsgOk returns a tuple with the Msg field value
// and a boolean to check if the value has been set.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) GetMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msg, true
}

// SetMsg sets field value
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) SetMsg(v string) {
	o.Msg = v
}

// GetTraceId returns the TraceId field value
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) GetTraceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value
// and a boolean to check if the value has been set.
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) GetTraceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TraceId, true
}

// SetTraceId sets field value
func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) SetTraceId(v string) {
	o.TraceId = v
}

func (o GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["data"] = o.Data
	toSerialize["msg"] = o.Msg
	toSerialize["traceId"] = o.TraceId
	return toSerialize, nil
}

func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"data",
		"msg",
		"traceId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse := _GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse)

	if err != nil {
		return err
	}

	*o = GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse(varGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse)

	return err
}

type NullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse struct {
	value *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse
	isSet bool
}

func (v NullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) Get() *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse {
	return v.value
}

func (v *NullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) Set(val *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse(val *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) *NullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse {
	return &NullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse{value: val, isSet: true}
}

func (v NullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
