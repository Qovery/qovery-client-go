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

// JobEnvironmentVariableApiService JobEnvironmentVariableApi service
type JobEnvironmentVariableApiService service

type ApiCreateJobEnvironmentVariableRequest struct {
	ctx                        context.Context
	ApiService                 *JobEnvironmentVariableApiService
	environmentVariableRequest *EnvironmentVariableRequest
}

func (r ApiCreateJobEnvironmentVariableRequest) EnvironmentVariableRequest(environmentVariableRequest EnvironmentVariableRequest) ApiCreateJobEnvironmentVariableRequest {
	r.environmentVariableRequest = &environmentVariableRequest
	return r
}

func (r ApiCreateJobEnvironmentVariableRequest) Execute() (*EnvironmentVariable, *http.Response, error) {
	return r.ApiService.CreateJobEnvironmentVariableExecute(r)
}

/*
CreateJobEnvironmentVariable Add an environment variable to the job

- Add an environment variable to the job.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateJobEnvironmentVariableRequest
*/
func (a *JobEnvironmentVariableApiService) CreateJobEnvironmentVariable(ctx context.Context) ApiCreateJobEnvironmentVariableRequest {
	return ApiCreateJobEnvironmentVariableRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return EnvironmentVariable
func (a *JobEnvironmentVariableApiService) CreateJobEnvironmentVariableExecute(r ApiCreateJobEnvironmentVariableRequest) (*EnvironmentVariable, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EnvironmentVariable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobEnvironmentVariableApiService.CreateJobEnvironmentVariable")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/job/{jobId}/environmentVariable"

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
	localVarPostBody = r.environmentVariableRequest
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

type ApiCreateJobEnvironmentVariableAliasRequest struct {
	ctx                   context.Context
	ApiService            *JobEnvironmentVariableApiService
	environmentVariableId string
	key                   *Key
}

func (r ApiCreateJobEnvironmentVariableAliasRequest) Key(key Key) ApiCreateJobEnvironmentVariableAliasRequest {
	r.key = &key
	return r
}

func (r ApiCreateJobEnvironmentVariableAliasRequest) Execute() (*EnvironmentVariable, *http.Response, error) {
	return r.ApiService.CreateJobEnvironmentVariableAliasExecute(r)
}

/*
CreateJobEnvironmentVariableAlias Create an environment variable alias at the job level

- Allows you to add an alias at job level on an existing environment variable having higher scope, in order to customize its key.
- You only have to specify a key in the request body
- The system will create a new environment variable at job level with the same value as the one corresponding to the variable id in the path
- The response body will contain the newly created variable
- Information regarding the aliased_variable will be exposed in the "aliased_variable" field of the newly created variable
- Only 1 alias level is allowed. You can't create an alias on an alias


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentVariableId Environment Variable ID
 @return ApiCreateJobEnvironmentVariableAliasRequest
*/
func (a *JobEnvironmentVariableApiService) CreateJobEnvironmentVariableAlias(ctx context.Context, environmentVariableId string) ApiCreateJobEnvironmentVariableAliasRequest {
	return ApiCreateJobEnvironmentVariableAliasRequest{
		ApiService:            a,
		ctx:                   ctx,
		environmentVariableId: environmentVariableId,
	}
}

// Execute executes the request
//  @return EnvironmentVariable
func (a *JobEnvironmentVariableApiService) CreateJobEnvironmentVariableAliasExecute(r ApiCreateJobEnvironmentVariableAliasRequest) (*EnvironmentVariable, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EnvironmentVariable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobEnvironmentVariableApiService.CreateJobEnvironmentVariableAlias")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/job/{jobId}/environmentVariable/{environmentVariableId}/alias"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentVariableId"+"}", url.PathEscape(parameterToString(r.environmentVariableId, "")), -1)

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
	localVarPostBody = r.key
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

type ApiCreateJobEnvironmentVariableOverrideRequest struct {
	ctx                   context.Context
	ApiService            *JobEnvironmentVariableApiService
	environmentVariableId string
	value                 *Value
}

func (r ApiCreateJobEnvironmentVariableOverrideRequest) Value(value Value) ApiCreateJobEnvironmentVariableOverrideRequest {
	r.value = &value
	return r
}

func (r ApiCreateJobEnvironmentVariableOverrideRequest) Execute() (*EnvironmentVariable, *http.Response, error) {
	return r.ApiService.CreateJobEnvironmentVariableOverrideExecute(r)
}

/*
CreateJobEnvironmentVariableOverride Create an environment variable override at the job level

- Allows you to override at job level an environment variable that has a higher scope.
- You only have to specify a value in the request body
- The system will create a new environment variable at job level with the same key as the one corresponding to the variable id in the path
- The response body will contain the newly created variable
- Information regarding the overridden_variable will be exposed in the "overridden_variable" field of the newly created variable


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentVariableId Environment Variable ID
 @return ApiCreateJobEnvironmentVariableOverrideRequest
*/
func (a *JobEnvironmentVariableApiService) CreateJobEnvironmentVariableOverride(ctx context.Context, environmentVariableId string) ApiCreateJobEnvironmentVariableOverrideRequest {
	return ApiCreateJobEnvironmentVariableOverrideRequest{
		ApiService:            a,
		ctx:                   ctx,
		environmentVariableId: environmentVariableId,
	}
}

// Execute executes the request
//  @return EnvironmentVariable
func (a *JobEnvironmentVariableApiService) CreateJobEnvironmentVariableOverrideExecute(r ApiCreateJobEnvironmentVariableOverrideRequest) (*EnvironmentVariable, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EnvironmentVariable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobEnvironmentVariableApiService.CreateJobEnvironmentVariableOverride")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/job/{jobId}/environmentVariable/{environmentVariableId}/override"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentVariableId"+"}", url.PathEscape(parameterToString(r.environmentVariableId, "")), -1)

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
	localVarPostBody = r.value
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

type ApiDeleteJobEnvironmentVariableRequest struct {
	ctx                   context.Context
	ApiService            *JobEnvironmentVariableApiService
	environmentVariableId string
}

func (r ApiDeleteJobEnvironmentVariableRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteJobEnvironmentVariableExecute(r)
}

/*
DeleteJobEnvironmentVariable Delete an environment variable from a job

- To delete an environment variable from an job you must have the project user permission
- You can't delete a BUILT_IN variable
- If you delete a variable having override or alias, the associated override/alias will be deleted as well


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentVariableId Environment Variable ID
 @return ApiDeleteJobEnvironmentVariableRequest
*/
func (a *JobEnvironmentVariableApiService) DeleteJobEnvironmentVariable(ctx context.Context, environmentVariableId string) ApiDeleteJobEnvironmentVariableRequest {
	return ApiDeleteJobEnvironmentVariableRequest{
		ApiService:            a,
		ctx:                   ctx,
		environmentVariableId: environmentVariableId,
	}
}

// Execute executes the request
func (a *JobEnvironmentVariableApiService) DeleteJobEnvironmentVariableExecute(r ApiDeleteJobEnvironmentVariableRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobEnvironmentVariableApiService.DeleteJobEnvironmentVariable")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/job/{jobId}/environmentVariable/{environmentVariableId}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentVariableId"+"}", url.PathEscape(parameterToString(r.environmentVariableId, "")), -1)

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

type ApiEditJobEnvironmentVariableRequest struct {
	ctx                            context.Context
	ApiService                     *JobEnvironmentVariableApiService
	environmentVariableId          string
	environmentVariableEditRequest *EnvironmentVariableEditRequest
}

func (r ApiEditJobEnvironmentVariableRequest) EnvironmentVariableEditRequest(environmentVariableEditRequest EnvironmentVariableEditRequest) ApiEditJobEnvironmentVariableRequest {
	r.environmentVariableEditRequest = &environmentVariableEditRequest
	return r
}

func (r ApiEditJobEnvironmentVariableRequest) Execute() (*EnvironmentVariable, *http.Response, error) {
	return r.ApiService.EditJobEnvironmentVariableExecute(r)
}

/*
EditJobEnvironmentVariable Edit an environment variable belonging to the job

- You can't edit a BUILT_IN variable
- For an override, you can't edit the key
- For an alias, you can't edit the value
- An override can only have a scope lower to the variable it is overriding (hierarchy is BUILT_IN > PROJECT > ENVIRONMENT > CONTAINER)


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentVariableId Environment Variable ID
 @return ApiEditJobEnvironmentVariableRequest
*/
func (a *JobEnvironmentVariableApiService) EditJobEnvironmentVariable(ctx context.Context, environmentVariableId string) ApiEditJobEnvironmentVariableRequest {
	return ApiEditJobEnvironmentVariableRequest{
		ApiService:            a,
		ctx:                   ctx,
		environmentVariableId: environmentVariableId,
	}
}

// Execute executes the request
//  @return EnvironmentVariable
func (a *JobEnvironmentVariableApiService) EditJobEnvironmentVariableExecute(r ApiEditJobEnvironmentVariableRequest) (*EnvironmentVariable, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EnvironmentVariable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobEnvironmentVariableApiService.EditJobEnvironmentVariable")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/job/{jobId}/environmentVariable/{environmentVariableId}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentVariableId"+"}", url.PathEscape(parameterToString(r.environmentVariableId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.environmentVariableEditRequest == nil {
		return localVarReturnValue, nil, reportError("environmentVariableEditRequest is required and must be specified")
	}

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
	localVarPostBody = r.environmentVariableEditRequest
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

type ApiImportJobEnvironmentVariableRequest struct {
	ctx                   context.Context
	ApiService            *JobEnvironmentVariableApiService
	variableImportRequest *VariableImportRequest
}

func (r ApiImportJobEnvironmentVariableRequest) VariableImportRequest(variableImportRequest VariableImportRequest) ApiImportJobEnvironmentVariableRequest {
	r.variableImportRequest = &variableImportRequest
	return r
}

func (r ApiImportJobEnvironmentVariableRequest) Execute() (*VariableImport, *http.Response, error) {
	return r.ApiService.ImportJobEnvironmentVariableExecute(r)
}

/*
ImportJobEnvironmentVariable Import variables

Import environment variables in a defined scope, with a defined visibility.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiImportJobEnvironmentVariableRequest
*/
func (a *JobEnvironmentVariableApiService) ImportJobEnvironmentVariable(ctx context.Context) ApiImportJobEnvironmentVariableRequest {
	return ApiImportJobEnvironmentVariableRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return VariableImport
func (a *JobEnvironmentVariableApiService) ImportJobEnvironmentVariableExecute(r ApiImportJobEnvironmentVariableRequest) (*VariableImport, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VariableImport
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobEnvironmentVariableApiService.ImportJobEnvironmentVariable")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/job/{jobId}/environmentVariable/import"

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
	localVarPostBody = r.variableImportRequest
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

type ApiListJobEnvironmentVariableRequest struct {
	ctx        context.Context
	ApiService *JobEnvironmentVariableApiService
}

func (r ApiListJobEnvironmentVariableRequest) Execute() (*EnvironmentVariableResponseList, *http.Response, error) {
	return r.ApiService.ListJobEnvironmentVariableExecute(r)
}

/*
ListJobEnvironmentVariable List environment variables

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListJobEnvironmentVariableRequest
*/
func (a *JobEnvironmentVariableApiService) ListJobEnvironmentVariable(ctx context.Context) ApiListJobEnvironmentVariableRequest {
	return ApiListJobEnvironmentVariableRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return EnvironmentVariableResponseList
func (a *JobEnvironmentVariableApiService) ListJobEnvironmentVariableExecute(r ApiListJobEnvironmentVariableRequest) (*EnvironmentVariableResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EnvironmentVariableResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobEnvironmentVariableApiService.ListJobEnvironmentVariable")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/job/{jobId}/environmentVariable"

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
