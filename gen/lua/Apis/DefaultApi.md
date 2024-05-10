# DefaultApi

All URIs are relative to *http://127.0.0.1:3013/goodsCenterLogic*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**v1CommonGenerateIdGet**](DefaultApi.md#v1CommonGenerateIdGet) | **GET** /v1/common/generateId | 雪花ID生成 |


<a name="v1CommonGenerateIdGet"></a>
# **v1CommonGenerateIdGet**
> internal_api_http_servicev1.HttpGenerateIDResponse v1CommonGenerateIdGet(num, id, Authorization)

雪花ID生成

    生成id-描述

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **num** | **Integer**| 生成id数量 1-1000 | [default to null] |
| **id** | **Integer**| ID | [default to null] |
| **Authorization** | **String**| 认证信息 eg:xxxx-xxxx-xxxx-xxx | [optional] [default to null] |

### Return type

[**internal_api_http_servicev1.HttpGenerateIDResponse**](../Models/internal_api_http_servicev1.HttpGenerateIDResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

