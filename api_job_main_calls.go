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

// JobMainCallsApiService JobMainCallsApi service
type JobMainCallsApiService service

type ApiDeleteJobRequest struct {
	ctx        context.Context
	ApiService *JobMainCallsApiService
	jobId      string
}

func (r ApiDeleteJobRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteJobExecute(r)
}

/*
DeleteJob Delete job

To delete the job you must have the admin permission

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param jobId Job ID
 @return ApiDeleteJobRequest
*/
func (a *JobMainCallsApiService) DeleteJob(ctx context.Context, jobId string) ApiDeleteJobRequest {
	return ApiDeleteJobRequest{
		ApiService: a,
		ctx:        ctx,
		jobId:      jobId,
	}
}

// Execute executes the request
func (a *JobMainCallsApiService) DeleteJobExecute(r ApiDeleteJobRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobMainCallsApiService.DeleteJob")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/job/{jobId}"
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterToString(r.jobId, "")), -1)

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

type ApiEditJobRequest struct {
	ctx        context.Context
	ApiService *JobMainCallsApiService
	jobId      string
	jobRequest *JobRequest
}

func (r ApiEditJobRequest) JobRequest(jobRequest JobRequest) ApiEditJobRequest {
	r.jobRequest = &jobRequest
	return r
}

func (r ApiEditJobRequest) Execute() (*JobResponse, *http.Response, error) {
	return r.ApiService.EditJobExecute(r)
}

/*
EditJob Edit job

- To edit the job you must have the admin permission.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param jobId Job ID
 @return ApiEditJobRequest
*/
func (a *JobMainCallsApiService) EditJob(ctx context.Context, jobId string) ApiEditJobRequest {
	return ApiEditJobRequest{
		ApiService: a,
		ctx:        ctx,
		jobId:      jobId,
	}
}

// Execute executes the request
//  @return JobResponse
func (a *JobMainCallsApiService) EditJobExecute(r ApiEditJobRequest) (*JobResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *JobResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobMainCallsApiService.EditJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/job/{jobId}"
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterToString(r.jobId, "")), -1)

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
	localVarPostBody = r.jobRequest
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

type ApiGetJobRequest struct {
	ctx        context.Context
	ApiService *JobMainCallsApiService
	jobId      string
}

func (r ApiGetJobRequest) Execute() (*JobResponse, *http.Response, error) {
	return r.ApiService.GetJobExecute(r)
}

/*
GetJob Get job by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param jobId Job ID
 @return ApiGetJobRequest
*/
func (a *JobMainCallsApiService) GetJob(ctx context.Context, jobId string) ApiGetJobRequest {
	return ApiGetJobRequest{
		ApiService: a,
		ctx:        ctx,
		jobId:      jobId,
	}
}

// Execute executes the request
//  @return JobResponse
func (a *JobMainCallsApiService) GetJobExecute(r ApiGetJobRequest) (*JobResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *JobResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobMainCallsApiService.GetJob")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/job/{jobId}"
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterToString(r.jobId, "")), -1)

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

type ApiGetJobStatusRequest struct {
	ctx        context.Context
	ApiService *JobMainCallsApiService
	jobId      string
}

func (r ApiGetJobStatusRequest) Execute() (*Status, *http.Response, error) {
	return r.ApiService.GetJobStatusExecute(r)
}

/*
GetJobStatus Get job status

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param jobId Job ID
 @return ApiGetJobStatusRequest
*/
func (a *JobMainCallsApiService) GetJobStatus(ctx context.Context, jobId string) ApiGetJobStatusRequest {
	return ApiGetJobStatusRequest{
		ApiService: a,
		ctx:        ctx,
		jobId:      jobId,
	}
}

// Execute executes the request
//  @return Status
func (a *JobMainCallsApiService) GetJobStatusExecute(r ApiGetJobStatusRequest) (*Status, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Status
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobMainCallsApiService.GetJobStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/job/{jobId}/status"
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterToString(r.jobId, "")), -1)

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

type ApiListJobCommitRequest struct {
	ctx         context.Context
	ApiService  *JobMainCallsApiService
	jobId       string
	startId     *string
	gitCommitId *string
}

// Starting point after which to return results
func (r ApiListJobCommitRequest) StartId(startId string) ApiListJobCommitRequest {
	r.startId = &startId
	return r
}

// Git Commit ID
func (r ApiListJobCommitRequest) GitCommitId(gitCommitId string) ApiListJobCommitRequest {
	r.gitCommitId = &gitCommitId
	return r
}

func (r ApiListJobCommitRequest) Execute() (*CommitResponseList, *http.Response, error) {
	return r.ApiService.ListJobCommitExecute(r)
}

/*
ListJobCommit List last job commits

Returns list of the last 100 commits made on the repository linked to the job

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param jobId Job ID
 @return ApiListJobCommitRequest
*/
func (a *JobMainCallsApiService) ListJobCommit(ctx context.Context, jobId string) ApiListJobCommitRequest {
	return ApiListJobCommitRequest{
		ApiService: a,
		ctx:        ctx,
		jobId:      jobId,
	}
}

// Execute executes the request
//  @return CommitResponseList
func (a *JobMainCallsApiService) ListJobCommitExecute(r ApiListJobCommitRequest) (*CommitResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CommitResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobMainCallsApiService.ListJobCommit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/job/{jobId}/commit"
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterToString(r.jobId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startId != nil {
		localVarQueryParams.Add("startId", parameterToString(*r.startId, ""))
	}
	if r.gitCommitId != nil {
		localVarQueryParams.Add("gitCommitId", parameterToString(*r.gitCommitId, ""))
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