# InternalApiDtoExamplePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **string** | AppID 应用id | 
**CVer** | **string** | CVer 客户端协议版本号 | 
**CmdId** | **int32** | CmdId 功能类型枚举: 1 AUTH鉴权; 2 QA 问答； 4 总结SUMMARY; 5 KEYWORDS 关键词； 6 CLASSIFICATION 文档分类； 7 RENAME  重命名； 8 mindmap 思维导图； | 
**Fid** | **string** | Fid： 文档id | 
**IgnoreFeaturePrompt** | Pointer to **int32** | IgnoreFeaturePrompt 是否忽略功能检查  1 是； 0 否； | [optional] 
**Lang** | **string** | Lang 语言 | 
**Question** | **string** | Question 问题 | 
**ShowGpt2** | Pointer to **string** | show_gpt_2 | [optional] 
**Token** | **string** | Token 鉴权token | 

## Methods

### NewInternalApiDtoExamplePostRequest

`func NewInternalApiDtoExamplePostRequest(appId string, cVer string, cmdId int32, fid string, lang string, question string, token string, ) *InternalApiDtoExamplePostRequest`

NewInternalApiDtoExamplePostRequest instantiates a new InternalApiDtoExamplePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalApiDtoExamplePostRequestWithDefaults

`func NewInternalApiDtoExamplePostRequestWithDefaults() *InternalApiDtoExamplePostRequest`

NewInternalApiDtoExamplePostRequestWithDefaults instantiates a new InternalApiDtoExamplePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *InternalApiDtoExamplePostRequest) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *InternalApiDtoExamplePostRequest) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *InternalApiDtoExamplePostRequest) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetCVer

`func (o *InternalApiDtoExamplePostRequest) GetCVer() string`

GetCVer returns the CVer field if non-nil, zero value otherwise.

### GetCVerOk

`func (o *InternalApiDtoExamplePostRequest) GetCVerOk() (*string, bool)`

GetCVerOk returns a tuple with the CVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCVer

`func (o *InternalApiDtoExamplePostRequest) SetCVer(v string)`

SetCVer sets CVer field to given value.


### GetCmdId

`func (o *InternalApiDtoExamplePostRequest) GetCmdId() int32`

GetCmdId returns the CmdId field if non-nil, zero value otherwise.

### GetCmdIdOk

`func (o *InternalApiDtoExamplePostRequest) GetCmdIdOk() (*int32, bool)`

GetCmdIdOk returns a tuple with the CmdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdId

`func (o *InternalApiDtoExamplePostRequest) SetCmdId(v int32)`

SetCmdId sets CmdId field to given value.


### GetFid

`func (o *InternalApiDtoExamplePostRequest) GetFid() string`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *InternalApiDtoExamplePostRequest) GetFidOk() (*string, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *InternalApiDtoExamplePostRequest) SetFid(v string)`

SetFid sets Fid field to given value.


### GetIgnoreFeaturePrompt

`func (o *InternalApiDtoExamplePostRequest) GetIgnoreFeaturePrompt() int32`

GetIgnoreFeaturePrompt returns the IgnoreFeaturePrompt field if non-nil, zero value otherwise.

### GetIgnoreFeaturePromptOk

`func (o *InternalApiDtoExamplePostRequest) GetIgnoreFeaturePromptOk() (*int32, bool)`

GetIgnoreFeaturePromptOk returns a tuple with the IgnoreFeaturePrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreFeaturePrompt

`func (o *InternalApiDtoExamplePostRequest) SetIgnoreFeaturePrompt(v int32)`

SetIgnoreFeaturePrompt sets IgnoreFeaturePrompt field to given value.

### HasIgnoreFeaturePrompt

`func (o *InternalApiDtoExamplePostRequest) HasIgnoreFeaturePrompt() bool`

HasIgnoreFeaturePrompt returns a boolean if a field has been set.

### GetLang

`func (o *InternalApiDtoExamplePostRequest) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *InternalApiDtoExamplePostRequest) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *InternalApiDtoExamplePostRequest) SetLang(v string)`

SetLang sets Lang field to given value.


### GetQuestion

`func (o *InternalApiDtoExamplePostRequest) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *InternalApiDtoExamplePostRequest) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *InternalApiDtoExamplePostRequest) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetShowGpt2

`func (o *InternalApiDtoExamplePostRequest) GetShowGpt2() string`

GetShowGpt2 returns the ShowGpt2 field if non-nil, zero value otherwise.

### GetShowGpt2Ok

`func (o *InternalApiDtoExamplePostRequest) GetShowGpt2Ok() (*string, bool)`

GetShowGpt2Ok returns a tuple with the ShowGpt2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowGpt2

`func (o *InternalApiDtoExamplePostRequest) SetShowGpt2(v string)`

SetShowGpt2 sets ShowGpt2 field to given value.

### HasShowGpt2

`func (o *InternalApiDtoExamplePostRequest) HasShowGpt2() bool`

HasShowGpt2 returns a boolean if a field has been set.

### GetToken

`func (o *InternalApiDtoExamplePostRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *InternalApiDtoExamplePostRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *InternalApiDtoExamplePostRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


