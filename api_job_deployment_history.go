/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.4
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// JobDeploymentHistoryAPIService JobDeploymentHistoryAPI service
type JobDeploymentHistoryAPIService service

type ApiListJobDeploymentHistoryRequest struct {
	ctx        context.Context
	ApiService *JobDeploymentHistoryAPIService
	jobId      string
}

func (r ApiListJobDeploymentHistoryRequest) Execute() (*ListJobDeploymentHistory200Response, *http.Response, error) {
	return r.ApiService.ListJobDeploymentHistoryExecute(r)
}

/*
ListJobDeploymentHistory List job deployments

Returns the 20 last job deployments

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param jobId Job ID
	@return ApiListJobDeploymentHistoryRequest
*/
func (a *JobDeploymentHistoryAPIService) ListJobDeploymentHistory(ctx context.Context, jobId string) ApiListJobDeploymentHistoryRequest {
	return ApiListJobDeploymentHistoryRequest{
		ApiService: a,
		ctx:        ctx,
		jobId:      jobId,
	}
}

// Execute executes the request
//
//	@return ListJobDeploymentHistory200Response
func (a *JobDeploymentHistoryAPIService) ListJobDeploymentHistoryExecute(r ApiListJobDeploymentHistoryRequest) (*ListJobDeploymentHistory200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListJobDeploymentHistory200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobDeploymentHistoryAPIService.ListJobDeploymentHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/job/{jobId}/deploymentHistory"
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterValueToString(r.jobId, "jobId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

type ApiListJobDeploymentHistoryV2Request struct {
	ctx        context.Context
	ApiService *JobDeploymentHistoryAPIService
	jobId      string
	pageSize   *float32
}

// The number of deployments to return in the current page
func (r ApiListJobDeploymentHistoryV2Request) PageSize(pageSize float32) ApiListJobDeploymentHistoryV2Request {
	r.pageSize = &pageSize
	return r
}

func (r ApiListJobDeploymentHistoryV2Request) Execute() (*DeploymentHistoryServicePaginatedResponseListV2, *http.Response, error) {
	return r.ApiService.ListJobDeploymentHistoryV2Execute(r)
}

/*
ListJobDeploymentHistoryV2 List job deployments

Returns the 20 last job deployments

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param jobId
	@return ApiListJobDeploymentHistoryV2Request
*/
func (a *JobDeploymentHistoryAPIService) ListJobDeploymentHistoryV2(ctx context.Context, jobId string) ApiListJobDeploymentHistoryV2Request {
	return ApiListJobDeploymentHistoryV2Request{
		ApiService: a,
		ctx:        ctx,
		jobId:      jobId,
	}
}

// Execute executes the request
//
//	@return DeploymentHistoryServicePaginatedResponseListV2
func (a *JobDeploymentHistoryAPIService) ListJobDeploymentHistoryV2Execute(r ApiListJobDeploymentHistoryV2Request) (*DeploymentHistoryServicePaginatedResponseListV2, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeploymentHistoryServicePaginatedResponseListV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobDeploymentHistoryAPIService.ListJobDeploymentHistoryV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/job/{jobId}/deploymentHistoryV2"
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterValueToString(r.jobId, "jobId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	} else {
		var defaultValue float32 = 20
		r.pageSize = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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
