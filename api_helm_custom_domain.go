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

// HelmCustomDomainAPIService HelmCustomDomainAPI service
type HelmCustomDomainAPIService service

type ApiCheckHelmCustomDomainRequest struct {
	ctx        context.Context
	ApiService *HelmCustomDomainAPIService
	helmId     string
}

func (r ApiCheckHelmCustomDomainRequest) Execute() (*CheckedCustomDomainsResponse, *http.Response, error) {
	return r.ApiService.CheckHelmCustomDomainExecute(r)
}

/*
CheckHelmCustomDomain Check Helm Custom Domain

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param helmId Helm ID
	@return ApiCheckHelmCustomDomainRequest
*/
func (a *HelmCustomDomainAPIService) CheckHelmCustomDomain(ctx context.Context, helmId string) ApiCheckHelmCustomDomainRequest {
	return ApiCheckHelmCustomDomainRequest{
		ApiService: a,
		ctx:        ctx,
		helmId:     helmId,
	}
}

// Execute executes the request
//
//	@return CheckedCustomDomainsResponse
func (a *HelmCustomDomainAPIService) CheckHelmCustomDomainExecute(r ApiCheckHelmCustomDomainRequest) (*CheckedCustomDomainsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CheckedCustomDomainsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HelmCustomDomainAPIService.CheckHelmCustomDomain")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/helm/{helmId}/checkCustomDomain"
	localVarPath = strings.Replace(localVarPath, "{"+"helmId"+"}", url.PathEscape(parameterValueToString(r.helmId, "helmId")), -1)

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

type ApiCreateHelmCustomDomainRequest struct {
	ctx                 context.Context
	ApiService          *HelmCustomDomainAPIService
	helmId              string
	customDomainRequest *CustomDomainRequest
}

func (r ApiCreateHelmCustomDomainRequest) CustomDomainRequest(customDomainRequest CustomDomainRequest) ApiCreateHelmCustomDomainRequest {
	r.customDomainRequest = &customDomainRequest
	return r
}

func (r ApiCreateHelmCustomDomainRequest) Execute() (*CustomDomain, *http.Response, error) {
	return r.ApiService.CreateHelmCustomDomainExecute(r)
}

/*
CreateHelmCustomDomain Add custom domain to the helm.

Add a custom domain to this helm in order not to use qovery autogenerated domain

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param helmId Helm ID
	@return ApiCreateHelmCustomDomainRequest
*/
func (a *HelmCustomDomainAPIService) CreateHelmCustomDomain(ctx context.Context, helmId string) ApiCreateHelmCustomDomainRequest {
	return ApiCreateHelmCustomDomainRequest{
		ApiService: a,
		ctx:        ctx,
		helmId:     helmId,
	}
}

// Execute executes the request
//
//	@return CustomDomain
func (a *HelmCustomDomainAPIService) CreateHelmCustomDomainExecute(r ApiCreateHelmCustomDomainRequest) (*CustomDomain, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomDomain
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HelmCustomDomainAPIService.CreateHelmCustomDomain")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/helm/{helmId}/customDomain"
	localVarPath = strings.Replace(localVarPath, "{"+"helmId"+"}", url.PathEscape(parameterValueToString(r.helmId, "helmId")), -1)

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
	localVarPostBody = r.customDomainRequest
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

type ApiDeleteHelmCustomDomainRequest struct {
	ctx            context.Context
	ApiService     *HelmCustomDomainAPIService
	helmId         string
	customDomainId string
}

func (r ApiDeleteHelmCustomDomainRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteHelmCustomDomainExecute(r)
}

/*
DeleteHelmCustomDomain Delete a Custom Domain

To delete an CustomDomain you must have the project user permission

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param helmId Helm ID
	@param customDomainId Custom Domain ID
	@return ApiDeleteHelmCustomDomainRequest
*/
func (a *HelmCustomDomainAPIService) DeleteHelmCustomDomain(ctx context.Context, helmId string, customDomainId string) ApiDeleteHelmCustomDomainRequest {
	return ApiDeleteHelmCustomDomainRequest{
		ApiService:     a,
		ctx:            ctx,
		helmId:         helmId,
		customDomainId: customDomainId,
	}
}

// Execute executes the request
func (a *HelmCustomDomainAPIService) DeleteHelmCustomDomainExecute(r ApiDeleteHelmCustomDomainRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HelmCustomDomainAPIService.DeleteHelmCustomDomain")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/helm/{helmId}/customDomain/{customDomainId}"
	localVarPath = strings.Replace(localVarPath, "{"+"helmId"+"}", url.PathEscape(parameterValueToString(r.helmId, "helmId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"customDomainId"+"}", url.PathEscape(parameterValueToString(r.customDomainId, "customDomainId")), -1)

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

type ApiEditHelmCustomDomainRequest struct {
	ctx                 context.Context
	ApiService          *HelmCustomDomainAPIService
	helmId              string
	customDomainId      string
	customDomainRequest *CustomDomainRequest
}

func (r ApiEditHelmCustomDomainRequest) CustomDomainRequest(customDomainRequest CustomDomainRequest) ApiEditHelmCustomDomainRequest {
	r.customDomainRequest = &customDomainRequest
	return r
}

func (r ApiEditHelmCustomDomainRequest) Execute() (*CustomDomain, *http.Response, error) {
	return r.ApiService.EditHelmCustomDomainExecute(r)
}

/*
EditHelmCustomDomain Edit a Custom Domain

To edit a Custom Domain you must have the project user permission

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param helmId Helm ID
	@param customDomainId Custom Domain ID
	@return ApiEditHelmCustomDomainRequest
*/
func (a *HelmCustomDomainAPIService) EditHelmCustomDomain(ctx context.Context, helmId string, customDomainId string) ApiEditHelmCustomDomainRequest {
	return ApiEditHelmCustomDomainRequest{
		ApiService:     a,
		ctx:            ctx,
		helmId:         helmId,
		customDomainId: customDomainId,
	}
}

// Execute executes the request
//
//	@return CustomDomain
func (a *HelmCustomDomainAPIService) EditHelmCustomDomainExecute(r ApiEditHelmCustomDomainRequest) (*CustomDomain, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomDomain
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HelmCustomDomainAPIService.EditHelmCustomDomain")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/helm/{helmId}/customDomain/{customDomainId}"
	localVarPath = strings.Replace(localVarPath, "{"+"helmId"+"}", url.PathEscape(parameterValueToString(r.helmId, "helmId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"customDomainId"+"}", url.PathEscape(parameterValueToString(r.customDomainId, "customDomainId")), -1)

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
	localVarPostBody = r.customDomainRequest
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

type ApiGetHelmCustomDomainRequest struct {
	ctx            context.Context
	ApiService     *HelmCustomDomainAPIService
	helmId         string
	customDomainId string
}

func (r ApiGetHelmCustomDomainRequest) Execute() (*CustomDomain, *http.Response, error) {
	return r.ApiService.GetHelmCustomDomainExecute(r)
}

/*
GetHelmCustomDomain Get a Custom Domain

Get a custom domain

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param helmId Helm ID
	@param customDomainId Custom Domain ID
	@return ApiGetHelmCustomDomainRequest
*/
func (a *HelmCustomDomainAPIService) GetHelmCustomDomain(ctx context.Context, helmId string, customDomainId string) ApiGetHelmCustomDomainRequest {
	return ApiGetHelmCustomDomainRequest{
		ApiService:     a,
		ctx:            ctx,
		helmId:         helmId,
		customDomainId: customDomainId,
	}
}

// Execute executes the request
//
//	@return CustomDomain
func (a *HelmCustomDomainAPIService) GetHelmCustomDomainExecute(r ApiGetHelmCustomDomainRequest) (*CustomDomain, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomDomain
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HelmCustomDomainAPIService.GetHelmCustomDomain")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/helm/{helmId}/customDomain/{customDomainId}"
	localVarPath = strings.Replace(localVarPath, "{"+"helmId"+"}", url.PathEscape(parameterValueToString(r.helmId, "helmId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"customDomainId"+"}", url.PathEscape(parameterValueToString(r.customDomainId, "customDomainId")), -1)

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

type ApiListHelmCustomDomainRequest struct {
	ctx        context.Context
	ApiService *HelmCustomDomainAPIService
	helmId     string
}

func (r ApiListHelmCustomDomainRequest) Execute() (*CustomDomainResponseList, *http.Response, error) {
	return r.ApiService.ListHelmCustomDomainExecute(r)
}

/*
ListHelmCustomDomain List helm custom domains

List the custom domains of this helm

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param helmId Helm ID
	@return ApiListHelmCustomDomainRequest
*/
func (a *HelmCustomDomainAPIService) ListHelmCustomDomain(ctx context.Context, helmId string) ApiListHelmCustomDomainRequest {
	return ApiListHelmCustomDomainRequest{
		ApiService: a,
		ctx:        ctx,
		helmId:     helmId,
	}
}

// Execute executes the request
//
//	@return CustomDomainResponseList
func (a *HelmCustomDomainAPIService) ListHelmCustomDomainExecute(r ApiListHelmCustomDomainRequest) (*CustomDomainResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomDomainResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HelmCustomDomainAPIService.ListHelmCustomDomain")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/helm/{helmId}/customDomain"
	localVarPath = strings.Replace(localVarPath, "{"+"helmId"+"}", url.PathEscape(parameterValueToString(r.helmId, "helmId")), -1)

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
