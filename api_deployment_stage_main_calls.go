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

// DeploymentStageMainCallsAPIService DeploymentStageMainCallsAPI service
type DeploymentStageMainCallsAPIService service

type ApiAttachServiceToDeploymentStageRequest struct {
	ctx               context.Context
	ApiService        *DeploymentStageMainCallsAPIService
	deploymentStageId string
	serviceId         string
}

func (r ApiAttachServiceToDeploymentStageRequest) Execute() (*DeploymentStageResponseList, *http.Response, error) {
	return r.ApiService.AttachServiceToDeploymentStageExecute(r)
}

/*
AttachServiceToDeploymentStage Attach service to deployment stage

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param deploymentStageId Deployment Stage ID
	@param serviceId Service ID of an application/job/container/database
	@return ApiAttachServiceToDeploymentStageRequest
*/
func (a *DeploymentStageMainCallsAPIService) AttachServiceToDeploymentStage(ctx context.Context, deploymentStageId string, serviceId string) ApiAttachServiceToDeploymentStageRequest {
	return ApiAttachServiceToDeploymentStageRequest{
		ApiService:        a,
		ctx:               ctx,
		deploymentStageId: deploymentStageId,
		serviceId:         serviceId,
	}
}

// Execute executes the request
//
//	@return DeploymentStageResponseList
func (a *DeploymentStageMainCallsAPIService) AttachServiceToDeploymentStageExecute(r ApiAttachServiceToDeploymentStageRequest) (*DeploymentStageResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeploymentStageResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentStageMainCallsAPIService.AttachServiceToDeploymentStage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deploymentStage/{deploymentStageId}/service/{serviceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"deploymentStageId"+"}", url.PathEscape(parameterValueToString(r.deploymentStageId, "deploymentStageId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

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

type ApiCreateEnvironmentDeploymentStageRequest struct {
	ctx                    context.Context
	ApiService             *DeploymentStageMainCallsAPIService
	environmentId          string
	deploymentStageRequest *DeploymentStageRequest
}

func (r ApiCreateEnvironmentDeploymentStageRequest) DeploymentStageRequest(deploymentStageRequest DeploymentStageRequest) ApiCreateEnvironmentDeploymentStageRequest {
	r.deploymentStageRequest = &deploymentStageRequest
	return r
}

func (r ApiCreateEnvironmentDeploymentStageRequest) Execute() (*DeploymentStageResponse, *http.Response, error) {
	return r.ApiService.CreateEnvironmentDeploymentStageExecute(r)
}

/*
CreateEnvironmentDeploymentStage Create environment deployment stage

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentId Environment ID
	@return ApiCreateEnvironmentDeploymentStageRequest
*/
func (a *DeploymentStageMainCallsAPIService) CreateEnvironmentDeploymentStage(ctx context.Context, environmentId string) ApiCreateEnvironmentDeploymentStageRequest {
	return ApiCreateEnvironmentDeploymentStageRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//
//	@return DeploymentStageResponse
func (a *DeploymentStageMainCallsAPIService) CreateEnvironmentDeploymentStageExecute(r ApiCreateEnvironmentDeploymentStageRequest) (*DeploymentStageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeploymentStageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentStageMainCallsAPIService.CreateEnvironmentDeploymentStage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/deploymentStage"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterValueToString(r.environmentId, "environmentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.deploymentStageRequest
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

type ApiDeleteDeploymentStageRequest struct {
	ctx               context.Context
	ApiService        *DeploymentStageMainCallsAPIService
	deploymentStageId string
}

func (r ApiDeleteDeploymentStageRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteDeploymentStageExecute(r)
}

/*
DeleteDeploymentStage Delete deployment stage

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param deploymentStageId Deployment Stage ID
	@return ApiDeleteDeploymentStageRequest
*/
func (a *DeploymentStageMainCallsAPIService) DeleteDeploymentStage(ctx context.Context, deploymentStageId string) ApiDeleteDeploymentStageRequest {
	return ApiDeleteDeploymentStageRequest{
		ApiService:        a,
		ctx:               ctx,
		deploymentStageId: deploymentStageId,
	}
}

// Execute executes the request
func (a *DeploymentStageMainCallsAPIService) DeleteDeploymentStageExecute(r ApiDeleteDeploymentStageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentStageMainCallsAPIService.DeleteDeploymentStage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deploymentStage/{deploymentStageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"deploymentStageId"+"}", url.PathEscape(parameterValueToString(r.deploymentStageId, "deploymentStageId")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiEditDeploymentStageRequest struct {
	ctx                    context.Context
	ApiService             *DeploymentStageMainCallsAPIService
	deploymentStageId      string
	deploymentStageRequest *DeploymentStageRequest
}

func (r ApiEditDeploymentStageRequest) DeploymentStageRequest(deploymentStageRequest DeploymentStageRequest) ApiEditDeploymentStageRequest {
	r.deploymentStageRequest = &deploymentStageRequest
	return r
}

func (r ApiEditDeploymentStageRequest) Execute() (*DeploymentStageResponse, *http.Response, error) {
	return r.ApiService.EditDeploymentStageExecute(r)
}

/*
EditDeploymentStage Edit deployment stage

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param deploymentStageId Deployment Stage ID
	@return ApiEditDeploymentStageRequest
*/
func (a *DeploymentStageMainCallsAPIService) EditDeploymentStage(ctx context.Context, deploymentStageId string) ApiEditDeploymentStageRequest {
	return ApiEditDeploymentStageRequest{
		ApiService:        a,
		ctx:               ctx,
		deploymentStageId: deploymentStageId,
	}
}

// Execute executes the request
//
//	@return DeploymentStageResponse
func (a *DeploymentStageMainCallsAPIService) EditDeploymentStageExecute(r ApiEditDeploymentStageRequest) (*DeploymentStageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeploymentStageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentStageMainCallsAPIService.EditDeploymentStage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deploymentStage/{deploymentStageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"deploymentStageId"+"}", url.PathEscape(parameterValueToString(r.deploymentStageId, "deploymentStageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.deploymentStageRequest
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

type ApiGetDeploymentStageRequest struct {
	ctx               context.Context
	ApiService        *DeploymentStageMainCallsAPIService
	deploymentStageId string
}

func (r ApiGetDeploymentStageRequest) Execute() (*DeploymentStageResponse, *http.Response, error) {
	return r.ApiService.GetDeploymentStageExecute(r)
}

/*
GetDeploymentStage Get Deployment Stage

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param deploymentStageId Deployment Stage ID
	@return ApiGetDeploymentStageRequest
*/
func (a *DeploymentStageMainCallsAPIService) GetDeploymentStage(ctx context.Context, deploymentStageId string) ApiGetDeploymentStageRequest {
	return ApiGetDeploymentStageRequest{
		ApiService:        a,
		ctx:               ctx,
		deploymentStageId: deploymentStageId,
	}
}

// Execute executes the request
//
//	@return DeploymentStageResponse
func (a *DeploymentStageMainCallsAPIService) GetDeploymentStageExecute(r ApiGetDeploymentStageRequest) (*DeploymentStageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeploymentStageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentStageMainCallsAPIService.GetDeploymentStage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deploymentStage/{deploymentStageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"deploymentStageId"+"}", url.PathEscape(parameterValueToString(r.deploymentStageId, "deploymentStageId")), -1)

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

type ApiGetServiceDeploymentStageRequest struct {
	ctx        context.Context
	ApiService *DeploymentStageMainCallsAPIService
	serviceId  string
}

func (r ApiGetServiceDeploymentStageRequest) Execute() (*DeploymentStageResponse, *http.Response, error) {
	return r.ApiService.GetServiceDeploymentStageExecute(r)
}

/*
GetServiceDeploymentStage Get Service Deployment Stage

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param serviceId Service ID of an application/job/container/database
	@return ApiGetServiceDeploymentStageRequest
*/
func (a *DeploymentStageMainCallsAPIService) GetServiceDeploymentStage(ctx context.Context, serviceId string) ApiGetServiceDeploymentStageRequest {
	return ApiGetServiceDeploymentStageRequest{
		ApiService: a,
		ctx:        ctx,
		serviceId:  serviceId,
	}
}

// Execute executes the request
//
//	@return DeploymentStageResponse
func (a *DeploymentStageMainCallsAPIService) GetServiceDeploymentStageExecute(r ApiGetServiceDeploymentStageRequest) (*DeploymentStageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeploymentStageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentStageMainCallsAPIService.GetServiceDeploymentStage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/{serviceId}/deploymentStage"
	localVarPath = strings.Replace(localVarPath, "{"+"serviceId"+"}", url.PathEscape(parameterValueToString(r.serviceId, "serviceId")), -1)

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

type ApiListEnvironmentDeploymentStageRequest struct {
	ctx           context.Context
	ApiService    *DeploymentStageMainCallsAPIService
	environmentId string
}

func (r ApiListEnvironmentDeploymentStageRequest) Execute() (*DeploymentStageResponseList, *http.Response, error) {
	return r.ApiService.ListEnvironmentDeploymentStageExecute(r)
}

/*
ListEnvironmentDeploymentStage List environment deployment stage

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentId Environment ID
	@return ApiListEnvironmentDeploymentStageRequest
*/
func (a *DeploymentStageMainCallsAPIService) ListEnvironmentDeploymentStage(ctx context.Context, environmentId string) ApiListEnvironmentDeploymentStageRequest {
	return ApiListEnvironmentDeploymentStageRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//
//	@return DeploymentStageResponseList
func (a *DeploymentStageMainCallsAPIService) ListEnvironmentDeploymentStageExecute(r ApiListEnvironmentDeploymentStageRequest) (*DeploymentStageResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeploymentStageResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentStageMainCallsAPIService.ListEnvironmentDeploymentStage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/deploymentStage"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterValueToString(r.environmentId, "environmentId")), -1)

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

type ApiMoveAfterDeploymentStageRequest struct {
	ctx               context.Context
	ApiService        *DeploymentStageMainCallsAPIService
	deploymentStageId string
	stageId           string
}

func (r ApiMoveAfterDeploymentStageRequest) Execute() (*DeploymentStageResponseList, *http.Response, error) {
	return r.ApiService.MoveAfterDeploymentStageExecute(r)
}

/*
MoveAfterDeploymentStage Move deployment stage after requested stage

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param deploymentStageId Deployment Stage ID
	@param stageId Deployment Stage ID
	@return ApiMoveAfterDeploymentStageRequest
*/
func (a *DeploymentStageMainCallsAPIService) MoveAfterDeploymentStage(ctx context.Context, deploymentStageId string, stageId string) ApiMoveAfterDeploymentStageRequest {
	return ApiMoveAfterDeploymentStageRequest{
		ApiService:        a,
		ctx:               ctx,
		deploymentStageId: deploymentStageId,
		stageId:           stageId,
	}
}

// Execute executes the request
//
//	@return DeploymentStageResponseList
func (a *DeploymentStageMainCallsAPIService) MoveAfterDeploymentStageExecute(r ApiMoveAfterDeploymentStageRequest) (*DeploymentStageResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeploymentStageResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentStageMainCallsAPIService.MoveAfterDeploymentStage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deploymentStage/{deploymentStageId}/moveAfter/{stageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"deploymentStageId"+"}", url.PathEscape(parameterValueToString(r.deploymentStageId, "deploymentStageId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stageId"+"}", url.PathEscape(parameterValueToString(r.stageId, "stageId")), -1)

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

type ApiMoveBeforeDeploymentStageRequest struct {
	ctx               context.Context
	ApiService        *DeploymentStageMainCallsAPIService
	deploymentStageId string
	stageId           string
}

func (r ApiMoveBeforeDeploymentStageRequest) Execute() (*DeploymentStageResponseList, *http.Response, error) {
	return r.ApiService.MoveBeforeDeploymentStageExecute(r)
}

/*
MoveBeforeDeploymentStage Move deployment stage before requested stage

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param deploymentStageId Deployment Stage ID
	@param stageId Deployment Stage ID
	@return ApiMoveBeforeDeploymentStageRequest
*/
func (a *DeploymentStageMainCallsAPIService) MoveBeforeDeploymentStage(ctx context.Context, deploymentStageId string, stageId string) ApiMoveBeforeDeploymentStageRequest {
	return ApiMoveBeforeDeploymentStageRequest{
		ApiService:        a,
		ctx:               ctx,
		deploymentStageId: deploymentStageId,
		stageId:           stageId,
	}
}

// Execute executes the request
//
//	@return DeploymentStageResponseList
func (a *DeploymentStageMainCallsAPIService) MoveBeforeDeploymentStageExecute(r ApiMoveBeforeDeploymentStageRequest) (*DeploymentStageResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DeploymentStageResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeploymentStageMainCallsAPIService.MoveBeforeDeploymentStage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/deploymentStage/{deploymentStageId}/moveBefore/{stageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"deploymentStageId"+"}", url.PathEscape(parameterValueToString(r.deploymentStageId, "deploymentStageId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stageId"+"}", url.PathEscape(parameterValueToString(r.stageId, "stageId")), -1)

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
