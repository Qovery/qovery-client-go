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

// Linger please
var (
	_ context.Context
)

// ApplicationDeploymentRestrictionApiService ApplicationDeploymentRestrictionApi service
type ApplicationDeploymentRestrictionApiService service

type ApiCreateApplicationDeploymentRestrictionRequest struct {
	ctx                                     context.Context
	ApiService                              *ApplicationDeploymentRestrictionApiService
	applicationId                           string
	applicationDeploymentRestrictionRequest *ApplicationDeploymentRestrictionRequest
}

func (r ApiCreateApplicationDeploymentRestrictionRequest) ApplicationDeploymentRestrictionRequest(applicationDeploymentRestrictionRequest ApplicationDeploymentRestrictionRequest) ApiCreateApplicationDeploymentRestrictionRequest {
	r.applicationDeploymentRestrictionRequest = &applicationDeploymentRestrictionRequest
	return r
}

func (r ApiCreateApplicationDeploymentRestrictionRequest) Execute() (*ApplicationDeploymentRestriction, *http.Response, error) {
	return r.ApiService.CreateApplicationDeploymentRestrictionExecute(r)
}

/*
CreateApplicationDeploymentRestriction Create an application deployment restriction

Create an application deployment restriction

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiCreateApplicationDeploymentRestrictionRequest
*/
func (a *ApplicationDeploymentRestrictionApiService) CreateApplicationDeploymentRestriction(ctx context.Context, applicationId string) ApiCreateApplicationDeploymentRestrictionRequest {
	return ApiCreateApplicationDeploymentRestrictionRequest{
		ApiService:    a,
		ctx:           ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return ApplicationDeploymentRestriction
func (a *ApplicationDeploymentRestrictionApiService) CreateApplicationDeploymentRestrictionExecute(r ApiCreateApplicationDeploymentRestrictionRequest) (*ApplicationDeploymentRestriction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApplicationDeploymentRestriction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationDeploymentRestrictionApiService.CreateApplicationDeploymentRestriction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/deploymentRestriction"
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
	localVarPostBody = r.applicationDeploymentRestrictionRequest
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

type ApiDeleteApplicationDeploymentRestrictionRequest struct {
	ctx           context.Context
	ApiService    *ApplicationDeploymentRestrictionApiService
	applicationId string
}

func (r ApiDeleteApplicationDeploymentRestrictionRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteApplicationDeploymentRestrictionExecute(r)
}

/*
DeleteApplicationDeploymentRestriction Delete an application deployment restriction

Delete an application deployment restriction

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiDeleteApplicationDeploymentRestrictionRequest
*/
func (a *ApplicationDeploymentRestrictionApiService) DeleteApplicationDeploymentRestriction(ctx context.Context, applicationId string) ApiDeleteApplicationDeploymentRestrictionRequest {
	return ApiDeleteApplicationDeploymentRestrictionRequest{
		ApiService:    a,
		ctx:           ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
func (a *ApplicationDeploymentRestrictionApiService) DeleteApplicationDeploymentRestrictionExecute(r ApiDeleteApplicationDeploymentRestrictionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationDeploymentRestrictionApiService.DeleteApplicationDeploymentRestriction")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/deploymentRestriction/{deploymentRestrictionId}"
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

type ApiEditApplicationDeploymentRestrictionRequest struct {
	ctx                                     context.Context
	ApiService                              *ApplicationDeploymentRestrictionApiService
	applicationId                           string
	deploymentRestrictionId                 string
	applicationDeploymentRestrictionRequest *ApplicationDeploymentRestrictionRequest
}

func (r ApiEditApplicationDeploymentRestrictionRequest) ApplicationDeploymentRestrictionRequest(applicationDeploymentRestrictionRequest ApplicationDeploymentRestrictionRequest) ApiEditApplicationDeploymentRestrictionRequest {
	r.applicationDeploymentRestrictionRequest = &applicationDeploymentRestrictionRequest
	return r
}

func (r ApiEditApplicationDeploymentRestrictionRequest) Execute() (*ApplicationDeploymentRestriction, *http.Response, error) {
	return r.ApiService.EditApplicationDeploymentRestrictionExecute(r)
}

/*
EditApplicationDeploymentRestriction Edit an application deployment restriction

Edit an application deployment restriction

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @param deploymentRestrictionId Deployment Restriction ID
 @return ApiEditApplicationDeploymentRestrictionRequest
*/
func (a *ApplicationDeploymentRestrictionApiService) EditApplicationDeploymentRestriction(ctx context.Context, applicationId string, deploymentRestrictionId string) ApiEditApplicationDeploymentRestrictionRequest {
	return ApiEditApplicationDeploymentRestrictionRequest{
		ApiService:              a,
		ctx:                     ctx,
		applicationId:           applicationId,
		deploymentRestrictionId: deploymentRestrictionId,
	}
}

// Execute executes the request
//  @return ApplicationDeploymentRestriction
func (a *ApplicationDeploymentRestrictionApiService) EditApplicationDeploymentRestrictionExecute(r ApiEditApplicationDeploymentRestrictionRequest) (*ApplicationDeploymentRestriction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApplicationDeploymentRestriction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationDeploymentRestrictionApiService.EditApplicationDeploymentRestriction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/deploymentRestriction/{deploymentRestrictionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterToString(r.applicationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deploymentRestrictionId"+"}", url.PathEscape(parameterToString(r.deploymentRestrictionId, "")), -1)

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
	localVarPostBody = r.applicationDeploymentRestrictionRequest
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

type ApiGetApplicationDeploymentRestrictionsRequest struct {
	ctx           context.Context
	ApiService    *ApplicationDeploymentRestrictionApiService
	applicationId string
}

func (r ApiGetApplicationDeploymentRestrictionsRequest) Execute() (*ApplicationDeploymentRestrictionResponseList, *http.Response, error) {
	return r.ApiService.GetApplicationDeploymentRestrictionsExecute(r)
}

/*
GetApplicationDeploymentRestrictions Get application deployment restrictions

Get application deployment restrictions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiGetApplicationDeploymentRestrictionsRequest
*/
func (a *ApplicationDeploymentRestrictionApiService) GetApplicationDeploymentRestrictions(ctx context.Context, applicationId string) ApiGetApplicationDeploymentRestrictionsRequest {
	return ApiGetApplicationDeploymentRestrictionsRequest{
		ApiService:    a,
		ctx:           ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return ApplicationDeploymentRestrictionResponseList
func (a *ApplicationDeploymentRestrictionApiService) GetApplicationDeploymentRestrictionsExecute(r ApiGetApplicationDeploymentRestrictionsRequest) (*ApplicationDeploymentRestrictionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApplicationDeploymentRestrictionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationDeploymentRestrictionApiService.GetApplicationDeploymentRestrictions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/deploymentRestriction"
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
