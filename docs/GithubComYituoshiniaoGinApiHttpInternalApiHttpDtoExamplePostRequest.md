# GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest

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

### NewGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest

`func NewGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest(appId string, cVer string, cmdId int32, fid string, lang string, question string, token string, ) *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest`

NewGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest instantiates a new GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequestWithDefaults

`func NewGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequestWithDefaults() *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest`

NewGithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequestWithDefaults instantiates a new GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetCVer

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) GetCVer() string`

GetCVer returns the CVer field if non-nil, zero value otherwise.

### GetCVerOk

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) GetCVerOk() (*string, bool)`

GetCVerOk returns a tuple with the CVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCVer

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) SetCVer(v string)`

SetCVer sets CVer field to given value.


### GetCmdId

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) GetCmdId() int32`

GetCmdId returns the CmdId field if non-nil, zero value otherwise.

### GetCmdIdOk

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) GetCmdIdOk() (*int32, bool)`

GetCmdIdOk returns a tuple with the CmdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdId

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) SetCmdId(v int32)`

SetCmdId sets CmdId field to given value.


### GetFid

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) GetFid() string`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) GetFidOk() (*string, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) SetFid(v string)`

SetFid sets Fid field to given value.


### GetIgnoreFeaturePrompt

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) GetIgnoreFeaturePrompt() int32`

GetIgnoreFeaturePrompt returns the IgnoreFeaturePrompt field if non-nil, zero value otherwise.

### GetIgnoreFeaturePromptOk

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) GetIgnoreFeaturePromptOk() (*int32, bool)`

GetIgnoreFeaturePromptOk returns a tuple with the IgnoreFeaturePrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreFeaturePrompt

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) SetIgnoreFeaturePrompt(v int32)`

SetIgnoreFeaturePrompt sets IgnoreFeaturePrompt field to given value.

### HasIgnoreFeaturePrompt

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) HasIgnoreFeaturePrompt() bool`

HasIgnoreFeaturePrompt returns a boolean if a field has been set.

### GetLang

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) SetLang(v string)`

SetLang sets Lang field to given value.


### GetQuestion

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) SetQuestion(v string)`

SetQuestion sets Question field to given value.


### GetShowGpt2

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) GetShowGpt2() string`

GetShowGpt2 returns the ShowGpt2 field if non-nil, zero value otherwise.

### GetShowGpt2Ok

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) GetShowGpt2Ok() (*string, bool)`

GetShowGpt2Ok returns a tuple with the ShowGpt2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowGpt2

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) SetShowGpt2(v string)`

SetShowGpt2 sets ShowGpt2 field to given value.

### HasShowGpt2

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) HasShowGpt2() bool`

HasShowGpt2 returns a boolean if a field has been set.

### GetToken

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GithubComYituoshiniaoGinApiHttpInternalApiHttpDtoExamplePostRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


