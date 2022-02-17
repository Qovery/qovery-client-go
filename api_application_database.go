/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.

API version: 1.0.0
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

// ApplicationDatabaseApiService ApplicationDatabaseApi service
type ApplicationDatabaseApiService service

type ApiAttachDatabasetoApplicationRequest struct {
	ctx              context.Context
	ApiService       *ApplicationDatabaseApiService
	applicationId    string
	targetDatabaseId string
}

func (r ApiAttachDatabasetoApplicationRequest) Execute() (*DatabaseResponse, *http.Response, error) {
	return r.ApiService.AttachDatabasetoApplicationExecute(r)
}

/*
AttachDatabasetoApplication Link a database to the application

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @param targetDatabaseId Target database ID
 @return ApiAttachDatabasetoApplicationRequest
*/
func (a *ApplicationDatabaseApiService) AttachDatabasetoApplication(ctx context.Context, applicationId string, targetDatabaseId string) ApiAttachDatabasetoApplicationRequest {
	return ApiAttachDatabasetoApplicationRequest{
		ApiService:       a,
		ctx:              ctx,
		applicationId:    applicationId,
		targetDatabaseId: targetDatabaseId,
	}
}

// Execute executes the request
//  @return DatabaseResponse
func (a *ApplicationDatabaseApiService) AttachDatabasetoApplicationExecute(r ApiAttachDatabasetoApplicationRequest) (*DatabaseResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationDatabaseApiService.AttachDatabasetoApplication")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/database/{targetDatabaseId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterToString(r.applicationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"targetDatabaseId"+"}", url.PathEscape(parameterToString(r.targetDatabaseId, "")), -1)

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

type ApiAttachLogicalDatabasetoApplicationRequest struct {
	ctx                     context.Context
	ApiService              *ApplicationDatabaseApiService
	applicationId           string
	targetLogicalDatabaseId string
}

func (r ApiAttachLogicalDatabasetoApplicationRequest) Execute() (*LogicalDatabaseResponse, *http.Response, error) {
	return r.ApiService.AttachLogicalDatabasetoApplicationExecute(r)
}

/*
AttachLogicalDatabasetoApplication Link a logical database to the application

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @param targetLogicalDatabaseId Target database ID
 @return ApiAttachLogicalDatabasetoApplicationRequest
*/
func (a *ApplicationDatabaseApiService) AttachLogicalDatabasetoApplication(ctx context.Context, applicationId string, targetLogicalDatabaseId string) ApiAttachLogicalDatabasetoApplicationRequest {
	return ApiAttachLogicalDatabasetoApplicationRequest{
		ApiService:              a,
		ctx:                     ctx,
		applicationId:           applicationId,
		targetLogicalDatabaseId: targetLogicalDatabaseId,
	}
}

// Execute executes the request
//  @return LogicalDatabaseResponse
func (a *ApplicationDatabaseApiService) AttachLogicalDatabasetoApplicationExecute(r ApiAttachLogicalDatabasetoApplicationRequest) (*LogicalDatabaseResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LogicalDatabaseResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationDatabaseApiService.AttachLogicalDatabasetoApplication")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/logicalDatabase/{targetLogicalDatabaseId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterToString(r.applicationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"targetLogicalDatabaseId"+"}", url.PathEscape(parameterToString(r.targetLogicalDatabaseId, "")), -1)

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

type ApiListApplicationDatabaseRequest struct {
	ctx           context.Context
	ApiService    *ApplicationDatabaseApiService
	applicationId string
}

func (r ApiListApplicationDatabaseRequest) Execute() (*DatabaseResponseList, *http.Response, error) {
	return r.ApiService.ListApplicationDatabaseExecute(r)
}

/*
ListApplicationDatabase List linked databases

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiListApplicationDatabaseRequest
*/
func (a *ApplicationDatabaseApiService) ListApplicationDatabase(ctx context.Context, applicationId string) ApiListApplicationDatabaseRequest {
	return ApiListApplicationDatabaseRequest{
		ApiService:    a,
		ctx:           ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return DatabaseResponseList
func (a *ApplicationDatabaseApiService) ListApplicationDatabaseExecute(r ApiListApplicationDatabaseRequest) (*DatabaseResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationDatabaseApiService.ListApplicationDatabase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/database"
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

type ApiListApplicationLogicalDatabaseRequest struct {
	ctx           context.Context
	ApiService    *ApplicationDatabaseApiService
	applicationId string
}

func (r ApiListApplicationLogicalDatabaseRequest) Execute() (*LogicalDatabaseResponseList, *http.Response, error) {
	return r.ApiService.ListApplicationLogicalDatabaseExecute(r)
}

/*
ListApplicationLogicalDatabase List linked logical databases

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiListApplicationLogicalDatabaseRequest
*/
func (a *ApplicationDatabaseApiService) ListApplicationLogicalDatabase(ctx context.Context, applicationId string) ApiListApplicationLogicalDatabaseRequest {
	return ApiListApplicationLogicalDatabaseRequest{
		ApiService:    a,
		ctx:           ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return LogicalDatabaseResponseList
func (a *ApplicationDatabaseApiService) ListApplicationLogicalDatabaseExecute(r ApiListApplicationLogicalDatabaseRequest) (*LogicalDatabaseResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LogicalDatabaseResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationDatabaseApiService.ListApplicationLogicalDatabase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/logicalDatabase"
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

type ApiRemoveDatabaseFromApplicationRequest struct {
	ctx              context.Context
	ApiService       *ApplicationDatabaseApiService
	applicationId    string
	targetDatabaseId string
}

func (r ApiRemoveDatabaseFromApplicationRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveDatabaseFromApplicationExecute(r)
}

/*
RemoveDatabaseFromApplication Remove database link to this application.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @param targetDatabaseId Target database ID
 @return ApiRemoveDatabaseFromApplicationRequest
*/
func (a *ApplicationDatabaseApiService) RemoveDatabaseFromApplication(ctx context.Context, applicationId string, targetDatabaseId string) ApiRemoveDatabaseFromApplicationRequest {
	return ApiRemoveDatabaseFromApplicationRequest{
		ApiService:       a,
		ctx:              ctx,
		applicationId:    applicationId,
		targetDatabaseId: targetDatabaseId,
	}
}

// Execute executes the request
func (a *ApplicationDatabaseApiService) RemoveDatabaseFromApplicationExecute(r ApiRemoveDatabaseFromApplicationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationDatabaseApiService.RemoveDatabaseFromApplication")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/database/{targetDatabaseId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterToString(r.applicationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"targetDatabaseId"+"}", url.PathEscape(parameterToString(r.targetDatabaseId, "")), -1)

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

type ApiRemoveLogicalDatabaseFromApplicationRequest struct {
	ctx                     context.Context
	ApiService              *ApplicationDatabaseApiService
	applicationId           string
	targetLogicalDatabaseId string
}

func (r ApiRemoveLogicalDatabaseFromApplicationRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveLogicalDatabaseFromApplicationExecute(r)
}

/*
RemoveLogicalDatabaseFromApplication Remove logical database link to this application.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @param targetLogicalDatabaseId Target database ID
 @return ApiRemoveLogicalDatabaseFromApplicationRequest
*/
func (a *ApplicationDatabaseApiService) RemoveLogicalDatabaseFromApplication(ctx context.Context, applicationId string, targetLogicalDatabaseId string) ApiRemoveLogicalDatabaseFromApplicationRequest {
	return ApiRemoveLogicalDatabaseFromApplicationRequest{
		ApiService:              a,
		ctx:                     ctx,
		applicationId:           applicationId,
		targetLogicalDatabaseId: targetLogicalDatabaseId,
	}
}

// Execute executes the request
func (a *ApplicationDatabaseApiService) RemoveLogicalDatabaseFromApplicationExecute(r ApiRemoveLogicalDatabaseFromApplicationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationDatabaseApiService.RemoveLogicalDatabaseFromApplication")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/logicalDatabase/{targetLogicalDatabaseId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterToString(r.applicationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"targetLogicalDatabaseId"+"}", url.PathEscape(parameterToString(r.targetLogicalDatabaseId, "")), -1)

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
