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
)

// GitRepositoriesApiService GitRepositoriesApi service
type GitRepositoriesApiService service

type ApiGetBitbucketRepositoriesRequest struct {
	ctx        context.Context
	ApiService *GitRepositoriesApiService
}

func (r ApiGetBitbucketRepositoriesRequest) Execute() (*GitRepositoryResponseList, *http.Response, error) {
	return r.ApiService.GetBitbucketRepositoriesExecute(r)
}

/*
GetBitbucketRepositories Get bitbucket repositories of the connected user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetBitbucketRepositoriesRequest

Deprecated
*/
func (a *GitRepositoriesApiService) GetBitbucketRepositories(ctx context.Context) ApiGetBitbucketRepositoriesRequest {
	return ApiGetBitbucketRepositoriesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GitRepositoryResponseList
// Deprecated
func (a *GitRepositoriesApiService) GetBitbucketRepositoriesExecute(r ApiGetBitbucketRepositoriesRequest) (*GitRepositoryResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GitRepositoryResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GitRepositoriesApiService.GetBitbucketRepositories")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account/bitbucket/repository"

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

type ApiGetBitbucketRepositoryBranchesRequest struct {
	ctx        context.Context
	ApiService *GitRepositoriesApiService
	name       *string
}

// The name of the repository where to retrieve the branches
func (r ApiGetBitbucketRepositoryBranchesRequest) Name(name string) ApiGetBitbucketRepositoryBranchesRequest {
	r.name = &name
	return r
}

func (r ApiGetBitbucketRepositoryBranchesRequest) Execute() (*GitRepositoryBranchResponseList, *http.Response, error) {
	return r.ApiService.GetBitbucketRepositoryBranchesExecute(r)
}

/*
GetBitbucketRepositoryBranches Get bitbucket branches of the specified repository

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetBitbucketRepositoryBranchesRequest

Deprecated
*/
func (a *GitRepositoriesApiService) GetBitbucketRepositoryBranches(ctx context.Context) ApiGetBitbucketRepositoryBranchesRequest {
	return ApiGetBitbucketRepositoryBranchesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GitRepositoryBranchResponseList
// Deprecated
func (a *GitRepositoriesApiService) GetBitbucketRepositoryBranchesExecute(r ApiGetBitbucketRepositoryBranchesRequest) (*GitRepositoryBranchResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GitRepositoryBranchResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GitRepositoriesApiService.GetBitbucketRepositoryBranches")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account/bitbucket/repository/branch"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
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

type ApiGetGitProviderAccountRequest struct {
	ctx        context.Context
	ApiService *GitRepositoriesApiService
}

func (r ApiGetGitProviderAccountRequest) Execute() (*GitAuthProviderResponseList, *http.Response, error) {
	return r.ApiService.GetGitProviderAccountExecute(r)
}

/*
GetGitProviderAccount Get git provider accounts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetGitProviderAccountRequest

Deprecated
*/
func (a *GitRepositoriesApiService) GetGitProviderAccount(ctx context.Context) ApiGetGitProviderAccountRequest {
	return ApiGetGitProviderAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GitAuthProviderResponseList
// Deprecated
func (a *GitRepositoriesApiService) GetGitProviderAccountExecute(r ApiGetGitProviderAccountRequest) (*GitAuthProviderResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GitAuthProviderResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GitRepositoriesApiService.GetGitProviderAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account/gitAuthProvider"

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

type ApiGetGithubRepositoriesRequest struct {
	ctx        context.Context
	ApiService *GitRepositoriesApiService
}

func (r ApiGetGithubRepositoriesRequest) Execute() (*GitRepositoryResponseList, *http.Response, error) {
	return r.ApiService.GetGithubRepositoriesExecute(r)
}

/*
GetGithubRepositories Get github repositories of the connected user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetGithubRepositoriesRequest

Deprecated
*/
func (a *GitRepositoriesApiService) GetGithubRepositories(ctx context.Context) ApiGetGithubRepositoriesRequest {
	return ApiGetGithubRepositoriesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GitRepositoryResponseList
// Deprecated
func (a *GitRepositoriesApiService) GetGithubRepositoriesExecute(r ApiGetGithubRepositoriesRequest) (*GitRepositoryResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GitRepositoryResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GitRepositoriesApiService.GetGithubRepositories")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account/github/repository"

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

type ApiGetGithubRepositoryBranchesRequest struct {
	ctx        context.Context
	ApiService *GitRepositoriesApiService
	name       *string
}

// The name of the repository where to retrieve the branches
func (r ApiGetGithubRepositoryBranchesRequest) Name(name string) ApiGetGithubRepositoryBranchesRequest {
	r.name = &name
	return r
}

func (r ApiGetGithubRepositoryBranchesRequest) Execute() (*GitRepositoryBranchResponseList, *http.Response, error) {
	return r.ApiService.GetGithubRepositoryBranchesExecute(r)
}

/*
GetGithubRepositoryBranches Get github branches of the specified repository

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetGithubRepositoryBranchesRequest

Deprecated
*/
func (a *GitRepositoriesApiService) GetGithubRepositoryBranches(ctx context.Context) ApiGetGithubRepositoryBranchesRequest {
	return ApiGetGithubRepositoryBranchesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GitRepositoryBranchResponseList
// Deprecated
func (a *GitRepositoriesApiService) GetGithubRepositoryBranchesExecute(r ApiGetGithubRepositoryBranchesRequest) (*GitRepositoryBranchResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GitRepositoryBranchResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GitRepositoriesApiService.GetGithubRepositoryBranches")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account/github/repository/branch"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
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

type ApiGetGitlabRepositoriesRequest struct {
	ctx        context.Context
	ApiService *GitRepositoriesApiService
}

func (r ApiGetGitlabRepositoriesRequest) Execute() (*GitRepositoryResponseList, *http.Response, error) {
	return r.ApiService.GetGitlabRepositoriesExecute(r)
}

/*
GetGitlabRepositories Get gitlab repositories of the connected user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetGitlabRepositoriesRequest

Deprecated
*/
func (a *GitRepositoriesApiService) GetGitlabRepositories(ctx context.Context) ApiGetGitlabRepositoriesRequest {
	return ApiGetGitlabRepositoriesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GitRepositoryResponseList
// Deprecated
func (a *GitRepositoriesApiService) GetGitlabRepositoriesExecute(r ApiGetGitlabRepositoriesRequest) (*GitRepositoryResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GitRepositoryResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GitRepositoriesApiService.GetGitlabRepositories")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account/gitlab/repository"

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

type ApiGetGitlabRepositoryBranchesRequest struct {
	ctx        context.Context
	ApiService *GitRepositoriesApiService
	name       *string
}

// The name of the repository to retrieve the branches
func (r ApiGetGitlabRepositoryBranchesRequest) Name(name string) ApiGetGitlabRepositoryBranchesRequest {
	r.name = &name
	return r
}

func (r ApiGetGitlabRepositoryBranchesRequest) Execute() (*GitRepositoryBranchResponseList, *http.Response, error) {
	return r.ApiService.GetGitlabRepositoryBranchesExecute(r)
}

/*
GetGitlabRepositoryBranches Get gitlab branches of the specified repository

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetGitlabRepositoryBranchesRequest

Deprecated
*/
func (a *GitRepositoriesApiService) GetGitlabRepositoryBranches(ctx context.Context) ApiGetGitlabRepositoryBranchesRequest {
	return ApiGetGitlabRepositoryBranchesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return GitRepositoryBranchResponseList
// Deprecated
func (a *GitRepositoriesApiService) GetGitlabRepositoryBranchesExecute(r ApiGetGitlabRepositoryBranchesRequest) (*GitRepositoryBranchResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GitRepositoryBranchResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GitRepositoriesApiService.GetGitlabRepositoryBranches")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account/gitlab/repository/branch"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
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
