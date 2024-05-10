/*
gin-http API

gin-http服务文档

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiV1CommonGenerateIdGetRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	num *int32
	id int32
	authorization *string
}

// 生成id数量 1-1000
func (r ApiV1CommonGenerateIdGetRequest) Num(num int32) ApiV1CommonGenerateIdGetRequest {
	r.num = &num
	return r
}

// 认证信息 eg:xxxx-xxxx-xxxx-xxx
func (r ApiV1CommonGenerateIdGetRequest) Authorization(authorization string) ApiV1CommonGenerateIdGetRequest {
	r.authorization = &authorization
	return r
}

func (r ApiV1CommonGenerateIdGetRequest) Execute() (*InternalApiHttpServicev1HttpGenerateIDResponse, *http.Response, error) {
	return r.ApiService.V1CommonGenerateIdGetExecute(r)
}

/*
V1CommonGenerateIdGet 雪花ID生成

生成id-描述

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ID
 @return ApiV1CommonGenerateIdGetRequest
*/
func (a *DefaultApiService) V1CommonGenerateIdGet(ctx context.Context, id int32) ApiV1CommonGenerateIdGetRequest {
	return ApiV1CommonGenerateIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return InternalApiHttpServicev1HttpGenerateIDResponse
func (a *DefaultApiService) V1CommonGenerateIdGetExecute(r ApiV1CommonGenerateIdGetRequest) (*InternalApiHttpServicev1HttpGenerateIDResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InternalApiHttpServicev1HttpGenerateIDResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.V1CommonGenerateIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/common/generateId"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.num == nil {
		return localVarReturnValue, nil, reportError("num is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "num", r.num, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
