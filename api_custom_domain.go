/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// CustomDomainApiService CustomDomainApi service
type CustomDomainApiService service

type ApiCreateApplicationCustomDomainRequest struct {
	ctx                 context.Context
	ApiService          *CustomDomainApiService
	applicationId       string
	customDomainRequest *CustomDomainRequest
}

func (r ApiCreateApplicationCustomDomainRequest) CustomDomainRequest(customDomainRequest CustomDomainRequest) ApiCreateApplicationCustomDomainRequest {
	r.customDomainRequest = &customDomainRequest
	return r
}

func (r ApiCreateApplicationCustomDomainRequest) Execute() (*CustomDomain, *http.Response, error) {
	return r.ApiService.CreateApplicationCustomDomainExecute(r)
}

/*
CreateApplicationCustomDomain Add custom domain to the application.

Add a custom domain to this application in order not to use qovery autogenerated domain

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiCreateApplicationCustomDomainRequest
*/
func (a *CustomDomainApiService) CreateApplicationCustomDomain(ctx context.Context, applicationId string) ApiCreateApplicationCustomDomainRequest {
	return ApiCreateApplicationCustomDomainRequest{
		ApiService:    a,
		ctx:           ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return CustomDomain
func (a *CustomDomainApiService) CreateApplicationCustomDomainExecute(r ApiCreateApplicationCustomDomainRequest) (*CustomDomain, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomDomain
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomDomainApiService.CreateApplicationCustomDomain")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/customDomain"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterToString(r.applicationId, "")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiDeleteCustomDomainRequest struct {
	ctx            context.Context
	ApiService     *CustomDomainApiService
	applicationId  string
	customDomainId string
}

func (r ApiDeleteCustomDomainRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCustomDomainExecute(r)
}

/*
DeleteCustomDomain Delete a Custom Domain

To delete an CustomDomain you must have the project user permission

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @param customDomainId Custom Domain ID
 @return ApiDeleteCustomDomainRequest
*/
func (a *CustomDomainApiService) DeleteCustomDomain(ctx context.Context, applicationId string, customDomainId string) ApiDeleteCustomDomainRequest {
	return ApiDeleteCustomDomainRequest{
		ApiService:     a,
		ctx:            ctx,
		applicationId:  applicationId,
		customDomainId: customDomainId,
	}
}

// Execute executes the request
func (a *CustomDomainApiService) DeleteCustomDomainExecute(r ApiDeleteCustomDomainRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomDomainApiService.DeleteCustomDomain")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/customDomain/{customDomainId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterToString(r.applicationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"customDomainId"+"}", url.PathEscape(parameterToString(r.customDomainId, "")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiEditCustomDomainRequest struct {
	ctx                 context.Context
	ApiService          *CustomDomainApiService
	applicationId       string
	customDomainId      string
	customDomainRequest *CustomDomainRequest
}

func (r ApiEditCustomDomainRequest) CustomDomainRequest(customDomainRequest CustomDomainRequest) ApiEditCustomDomainRequest {
	r.customDomainRequest = &customDomainRequest
	return r
}

func (r ApiEditCustomDomainRequest) Execute() (*CustomDomain, *http.Response, error) {
	return r.ApiService.EditCustomDomainExecute(r)
}

/*
EditCustomDomain Edit a Custom Domain

To edit a Custom Domain  you must have the project user permission

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @param customDomainId Custom Domain ID
 @return ApiEditCustomDomainRequest
*/
func (a *CustomDomainApiService) EditCustomDomain(ctx context.Context, applicationId string, customDomainId string) ApiEditCustomDomainRequest {
	return ApiEditCustomDomainRequest{
		ApiService:     a,
		ctx:            ctx,
		applicationId:  applicationId,
		customDomainId: customDomainId,
	}
}

// Execute executes the request
//  @return CustomDomain
func (a *CustomDomainApiService) EditCustomDomainExecute(r ApiEditCustomDomainRequest) (*CustomDomain, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomDomain
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomDomainApiService.EditCustomDomain")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/customDomain/{customDomainId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterToString(r.applicationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"customDomainId"+"}", url.PathEscape(parameterToString(r.customDomainId, "")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetCustomDomainStatusRequest struct {
	ctx            context.Context
	ApiService     *CustomDomainApiService
	applicationId  string
	customDomainId string
}

func (r ApiGetCustomDomainStatusRequest) Execute() (*CustomDomain, *http.Response, error) {
	return r.ApiService.GetCustomDomainStatusExecute(r)
}

/*
GetCustomDomainStatus Get Custom Domain status

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @param customDomainId Custom Domain ID
 @return ApiGetCustomDomainStatusRequest
*/
func (a *CustomDomainApiService) GetCustomDomainStatus(ctx context.Context, applicationId string, customDomainId string) ApiGetCustomDomainStatusRequest {
	return ApiGetCustomDomainStatusRequest{
		ApiService:     a,
		ctx:            ctx,
		applicationId:  applicationId,
		customDomainId: customDomainId,
	}
}

// Execute executes the request
//  @return CustomDomain
func (a *CustomDomainApiService) GetCustomDomainStatusExecute(r ApiGetCustomDomainStatusRequest) (*CustomDomain, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomDomain
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomDomainApiService.GetCustomDomainStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/customDomain/{customDomainId}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterToString(r.applicationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"customDomainId"+"}", url.PathEscape(parameterToString(r.customDomainId, "")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiListApplicationCustomDomainRequest struct {
	ctx           context.Context
	ApiService    *CustomDomainApiService
	applicationId string
}

func (r ApiListApplicationCustomDomainRequest) Execute() (*CustomDomainResponseList, *http.Response, error) {
	return r.ApiService.ListApplicationCustomDomainExecute(r)
}

/*
ListApplicationCustomDomain List application custom domains

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiListApplicationCustomDomainRequest
*/
func (a *CustomDomainApiService) ListApplicationCustomDomain(ctx context.Context, applicationId string) ApiListApplicationCustomDomainRequest {
	return ApiListApplicationCustomDomainRequest{
		ApiService:    a,
		ctx:           ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return CustomDomainResponseList
func (a *CustomDomainApiService) ListApplicationCustomDomainExecute(r ApiListApplicationCustomDomainRequest) (*CustomDomainResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CustomDomainResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomDomainApiService.ListApplicationCustomDomain")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/customDomain"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterToString(r.applicationId, "")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
