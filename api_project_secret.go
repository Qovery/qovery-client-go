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

// ProjectSecretAPIService ProjectSecretAPI service
type ProjectSecretAPIService service

type ApiCreateProjectSecretRequest struct {
	ctx           context.Context
	ApiService    *ProjectSecretAPIService
	projectId     string
	secretRequest *SecretRequest
}

func (r ApiCreateProjectSecretRequest) SecretRequest(secretRequest SecretRequest) ApiCreateProjectSecretRequest {
	r.secretRequest = &secretRequest
	return r
}

func (r ApiCreateProjectSecretRequest) Execute() (*Secret, *http.Response, error) {
	return r.ApiService.CreateProjectSecretExecute(r)
}

/*
CreateProjectSecret Add a secret to the project

- Add a secret to the project.

  - If the secret key already exists, then it will be replaced by the new one.

  - If the secret value points toward an existing secret key, it will be considered as an alias.

    @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
    @param projectId Project ID
    @return ApiCreateProjectSecretRequest
*/
func (a *ProjectSecretAPIService) CreateProjectSecret(ctx context.Context, projectId string) ApiCreateProjectSecretRequest {
	return ApiCreateProjectSecretRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
	}
}

// Execute executes the request
//
//	@return Secret
func (a *ProjectSecretAPIService) CreateProjectSecretExecute(r ApiCreateProjectSecretRequest) (*Secret, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Secret
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectSecretAPIService.CreateProjectSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{projectId}/secret"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

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
	localVarPostBody = r.secretRequest
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

type ApiCreateProjectSecretAliasRequest struct {
	ctx        context.Context
	ApiService *ProjectSecretAPIService
	projectId  string
	secretId   string
	key        *Key
}

func (r ApiCreateProjectSecretAliasRequest) Key(key Key) ApiCreateProjectSecretAliasRequest {
	r.key = &key
	return r
}

func (r ApiCreateProjectSecretAliasRequest) Execute() (*Secret, *http.Response, error) {
	return r.ApiService.CreateProjectSecretAliasExecute(r)
}

/*
CreateProjectSecretAlias Create a secret alias at the project level

- Allows you to add an alias at project level on an existing secret having higher scope, in order to customize its key.
- You only have to specify a key in the request body
- The system will create a new secret at project level with the same value as the one corresponding to the secret id in the path
- The response body will contain the newly created secret
- Information regarding the aliased_secret will be exposed in the "aliased_secret" field of the newly created secret
- You can't create an alias on an alias

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId Project ID
	@param secretId Secret ID
	@return ApiCreateProjectSecretAliasRequest
*/
func (a *ProjectSecretAPIService) CreateProjectSecretAlias(ctx context.Context, projectId string, secretId string) ApiCreateProjectSecretAliasRequest {
	return ApiCreateProjectSecretAliasRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		secretId:   secretId,
	}
}

// Execute executes the request
//
//	@return Secret
func (a *ProjectSecretAPIService) CreateProjectSecretAliasExecute(r ApiCreateProjectSecretAliasRequest) (*Secret, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Secret
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectSecretAPIService.CreateProjectSecretAlias")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{projectId}/secret/{secretId}/alias"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secretId"+"}", url.PathEscape(parameterValueToString(r.secretId, "secretId")), -1)

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

type ApiCreateProjectSecretOverrideRequest struct {
	ctx        context.Context
	ApiService *ProjectSecretAPIService
	projectId  string
	secretId   string
	value      *Value
}

func (r ApiCreateProjectSecretOverrideRequest) Value(value Value) ApiCreateProjectSecretOverrideRequest {
	r.value = &value
	return r
}

func (r ApiCreateProjectSecretOverrideRequest) Execute() (*Secret, *http.Response, error) {
	return r.ApiService.CreateProjectSecretOverrideExecute(r)
}

/*
CreateProjectSecretOverride Create a secret override at the project level

- Allows you to override at project level a secret that has a higher scope.
- You only have to specify a value in the request body
- The system will create a new secret at project level with the same key as the one corresponding to the secret id in the path
- The response body will contain the newly created secret
- Information regarding the overridden_secret will be exposed in the "overridden_secret" field of the newly created secret

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId Project ID
	@param secretId Secret ID
	@return ApiCreateProjectSecretOverrideRequest
*/
func (a *ProjectSecretAPIService) CreateProjectSecretOverride(ctx context.Context, projectId string, secretId string) ApiCreateProjectSecretOverrideRequest {
	return ApiCreateProjectSecretOverrideRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		secretId:   secretId,
	}
}

// Execute executes the request
//
//	@return Secret
func (a *ProjectSecretAPIService) CreateProjectSecretOverrideExecute(r ApiCreateProjectSecretOverrideRequest) (*Secret, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Secret
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectSecretAPIService.CreateProjectSecretOverride")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{projectId}/secret/{secretId}/override"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secretId"+"}", url.PathEscape(parameterValueToString(r.secretId, "secretId")), -1)

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

type ApiDeleteProjectSecretRequest struct {
	ctx        context.Context
	ApiService *ProjectSecretAPIService
	projectId  string
	secretId   string
}

func (r ApiDeleteProjectSecretRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteProjectSecretExecute(r)
}

/*
DeleteProjectSecret Delete a secret from a project

- To delete a secret you must have the project user permission
- You can't delete a BUILT_IN secret
- If you delete a secret having override or alias, the associated override/alias will be deleted as well  operationId: deleteProjectSecret

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId Project ID
	@param secretId Secret ID
	@return ApiDeleteProjectSecretRequest
*/
func (a *ProjectSecretAPIService) DeleteProjectSecret(ctx context.Context, projectId string, secretId string) ApiDeleteProjectSecretRequest {
	return ApiDeleteProjectSecretRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		secretId:   secretId,
	}
}

// Execute executes the request
func (a *ProjectSecretAPIService) DeleteProjectSecretExecute(r ApiDeleteProjectSecretRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectSecretAPIService.DeleteProjectSecret")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{projectId}/secret/{secretId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secretId"+"}", url.PathEscape(parameterValueToString(r.secretId, "secretId")), -1)

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

type ApiEditProjectSecretRequest struct {
	ctx               context.Context
	ApiService        *ProjectSecretAPIService
	projectId         string
	secretId          string
	secretEditRequest *SecretEditRequest
}

func (r ApiEditProjectSecretRequest) SecretEditRequest(secretEditRequest SecretEditRequest) ApiEditProjectSecretRequest {
	r.secretEditRequest = &secretEditRequest
	return r
}

func (r ApiEditProjectSecretRequest) Execute() (*Secret, *http.Response, error) {
	return r.ApiService.EditProjectSecretExecute(r)
}

/*
EditProjectSecret Edit a secret belonging to the project

- You can't edit a BUILT_IN secret
- For an override, you can't edit the key
- For an alias, you can't edit the value
- An override can only have a scope lower to the secret it is overriding (hierarchy is BUILT_IN > PROJECT > ENVIRONMENT > APPLICATION)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId Project ID
	@param secretId Secret ID
	@return ApiEditProjectSecretRequest
*/
func (a *ProjectSecretAPIService) EditProjectSecret(ctx context.Context, projectId string, secretId string) ApiEditProjectSecretRequest {
	return ApiEditProjectSecretRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
		secretId:   secretId,
	}
}

// Execute executes the request
//
//	@return Secret
func (a *ProjectSecretAPIService) EditProjectSecretExecute(r ApiEditProjectSecretRequest) (*Secret, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Secret
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectSecretAPIService.EditProjectSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{projectId}/secret/{secretId}"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secretId"+"}", url.PathEscape(parameterValueToString(r.secretId, "secretId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.secretEditRequest == nil {
		return localVarReturnValue, nil, reportError("secretEditRequest is required and must be specified")
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
	localVarPostBody = r.secretEditRequest
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

type ApiListProjectSecretsRequest struct {
	ctx        context.Context
	ApiService *ProjectSecretAPIService
	projectId  string
}

func (r ApiListProjectSecretsRequest) Execute() (*SecretResponseList, *http.Response, error) {
	return r.ApiService.ListProjectSecretsExecute(r)
}

/*
ListProjectSecrets List project secrets

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId Project ID
	@return ApiListProjectSecretsRequest
*/
func (a *ProjectSecretAPIService) ListProjectSecrets(ctx context.Context, projectId string) ApiListProjectSecretsRequest {
	return ApiListProjectSecretsRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
	}
}

// Execute executes the request
//
//	@return SecretResponseList
func (a *ProjectSecretAPIService) ListProjectSecretsExecute(r ApiListProjectSecretsRequest) (*SecretResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SecretResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProjectSecretAPIService.ListProjectSecrets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{projectId}/secret"
	localVarPath = strings.Replace(localVarPath, "{"+"projectId"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

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
