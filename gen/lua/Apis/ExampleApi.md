# ExampleApi

All URIs are relative to *http://127.0.0.1:3013/goodsCenterLogic*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**v1ExampleGetGet**](ExampleApi.md#v1ExampleGetGet) | **GET** /v1/exampleGet | get接口示例 |
| [**v1ExampleGetOneGet**](ExampleApi.md#v1ExampleGetOneGet) | **GET** /v1/exampleGetOne | getOne接口示例 |
| [**v1ExamplePostPost**](ExampleApi.md#v1ExamplePostPost) | **POST** /v1/examplePost | post 接口 示例 |


<a name="v1ExampleGetGet"></a>
# **v1ExampleGetGet**
> internal_api_http_dto.ExampleGetResponse v1ExampleGetGet(create\_time, query\_id, user\_id)

get接口示例

    get接口示例

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **create\_time** | **Integer**| create_time | [default to null] |
| **query\_id** | **String**| query_id | [default to null] |
| **user\_id** | **String**| user_id 用户id | [default to null] |

### Return type

[**internal_api_http_dto.ExampleGetResponse**](../Models/internal_api_http_dto.ExampleGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

<a name="v1ExampleGetOneGet"></a>
# **v1ExampleGetOneGet**
> internal_api_http_dto.ExampleGetOneResponse v1ExampleGetOneGet(create\_time, query\_id, user\_id)

getOne接口示例

    getOne接口示例

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **create\_time** | **Integer**| create_time | [default to null] |
| **query\_id** | **String**| query_id | [default to null] |
| **user\_id** | **String**| user_id 用户id | [default to null] |

### Return type

[**internal_api_http_dto.ExampleGetOneResponse**](../Models/internal_api_http_dto.ExampleGetOneResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

<a name="v1ExamplePostPost"></a>
# **v1ExamplePostPost**
> internal_api_http_dto.ExamplePostResponse v1ExamplePostPost(body)

post 接口 示例

     ios购买项类型 &lt;a href&#x3D;\&quot;https://developer.apple.com/documentation/appstoreconnectapi/list_all_in-app_purchases_for_an_app\&quot;&gt; 详情&lt;/a&gt; &lt;/br&gt;    android订阅 &lt;a href&#x3D;\&quot;https://developers.google.com/android-publisher/api-ref/rest/v3/monetization.subscriptions/list?hl&#x3D;zh-cn\&quot;&gt; 详情&lt;/a&gt; &lt;/br&gt;   android非订阅&lt;a href&#x3D;\&quot;https://developers.google.com/android-publisher/api-ref/rest/v3/inappproducts/list?hl&#x3D;zh-cn\&quot;&gt; 详情&lt;/a&gt; &lt;/br&gt;   android订阅产品的类型&lt;a href&#x3D;\&quot;https://developers.google.com/android-publisher/api-ref/rest/v3/inappproducts?hl&#x3D;zh-cn#PurchaseType\&quot;&gt; 详情&lt;/a&gt; &lt;/br&gt;

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | [**internal_api_http_dto.ExamplePostRequest**](../Models/internal_api_http_dto.ExamplePostRequest.md)| 请求参数 | |

### Return type

[**internal_api_http_dto.ExamplePostResponse**](../Models/internal_api_http_dto.ExamplePostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

