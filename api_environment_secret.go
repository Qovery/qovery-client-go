/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.2
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

// EnvironmentSecretApiService EnvironmentSecretApi service
type EnvironmentSecretApiService service

type ApiCreateEnvironmentSecretRequest struct {
	ctx           context.Context
	ApiService    *EnvironmentSecretApiService
	environmentId string
	secretRequest *SecretRequest
}

func (r ApiCreateEnvironmentSecretRequest) SecretRequest(secretRequest SecretRequest) ApiCreateEnvironmentSecretRequest {
	r.secretRequest = &secretRequest
	return r
}

func (r ApiCreateEnvironmentSecretRequest) Execute() (*SecretResponse, *http.Response, error) {
	return r.ApiService.CreateEnvironmentSecretExecute(r)
}

/*
CreateEnvironmentSecret Add a secret to the environment

- Add a secret to the environment.
  - If the secret key already exists, then it will be replaced by the new one.
  - If the secret value points toward an existing secret key, it will be considered as an alias.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiCreateEnvironmentSecretRequest
*/
func (a *EnvironmentSecretApiService) CreateEnvironmentSecret(ctx context.Context, environmentId string) ApiCreateEnvironmentSecretRequest {
	return ApiCreateEnvironmentSecretRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return SecretResponse
func (a *EnvironmentSecretApiService) CreateEnvironmentSecretExecute(r ApiCreateEnvironmentSecretRequest) (*SecretResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SecretResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentSecretApiService.CreateEnvironmentSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/secret"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterToString(r.environmentId, "")), -1)

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

type ApiCreateEnvironmentSecretAliasRequest struct {
	ctx           context.Context
	ApiService    *EnvironmentSecretApiService
	environmentId string
	secretId      string
	key           *Key
}

func (r ApiCreateEnvironmentSecretAliasRequest) Key(key Key) ApiCreateEnvironmentSecretAliasRequest {
	r.key = &key
	return r
}

func (r ApiCreateEnvironmentSecretAliasRequest) Execute() (*SecretResponse, *http.Response, error) {
	return r.ApiService.CreateEnvironmentSecretAliasExecute(r)
}

/*
CreateEnvironmentSecretAlias Create a secret alias at the environment level

- Allows you to add an alias at environment level on an existing secret having higher scope, in order to customize its key.
- You only have to specify a key in the request body
- The system will create a new secret at environment level with the same value as the one corresponding to the secret id in the path
- The response body will contain the newly created secret
- Information regarding the aliased_secret will be exposed in the "aliased_secret" field of the newly created secret
- Only 1 alias level is allowed. You can't create an alias on an alias


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @param secretId Secret ID
 @return ApiCreateEnvironmentSecretAliasRequest
*/
func (a *EnvironmentSecretApiService) CreateEnvironmentSecretAlias(ctx context.Context, environmentId string, secretId string) ApiCreateEnvironmentSecretAliasRequest {
	return ApiCreateEnvironmentSecretAliasRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
		secretId:      secretId,
	}
}

// Execute executes the request
//  @return SecretResponse
func (a *EnvironmentSecretApiService) CreateEnvironmentSecretAliasExecute(r ApiCreateEnvironmentSecretAliasRequest) (*SecretResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SecretResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentSecretApiService.CreateEnvironmentSecretAlias")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/secret/{secretId}/alias"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secretId"+"}", url.PathEscape(parameterToString(r.secretId, "")), -1)

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

type ApiCreateEnvironmentSecretOverrideRequest struct {
	ctx           context.Context
	ApiService    *EnvironmentSecretApiService
	environmentId string
	secretId      string
	value         *Value
}

func (r ApiCreateEnvironmentSecretOverrideRequest) Value(value Value) ApiCreateEnvironmentSecretOverrideRequest {
	r.value = &value
	return r
}

func (r ApiCreateEnvironmentSecretOverrideRequest) Execute() (*SecretResponse, *http.Response, error) {
	return r.ApiService.CreateEnvironmentSecretOverrideExecute(r)
}

/*
CreateEnvironmentSecretOverride Create a secret override at the environment level

- Allows you to override at environment level a secret that has a higher scope.
- You only have to specify a value in the request body
- The system will create a new secret at environment level with the same key as the one corresponding to the secret id in the path
- The response body will contain the newly created secret
- Information regarding the overridden_secret will be exposed in the "overridden_secret" field of the newly created secret


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @param secretId Secret ID
 @return ApiCreateEnvironmentSecretOverrideRequest
*/
func (a *EnvironmentSecretApiService) CreateEnvironmentSecretOverride(ctx context.Context, environmentId string, secretId string) ApiCreateEnvironmentSecretOverrideRequest {
	return ApiCreateEnvironmentSecretOverrideRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
		secretId:      secretId,
	}
}

// Execute executes the request
//  @return SecretResponse
func (a *EnvironmentSecretApiService) CreateEnvironmentSecretOverrideExecute(r ApiCreateEnvironmentSecretOverrideRequest) (*SecretResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SecretResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentSecretApiService.CreateEnvironmentSecretOverride")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/secret/{secretId}/override"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secretId"+"}", url.PathEscape(parameterToString(r.secretId, "")), -1)

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

type ApiDeleteEnvironmentSecretRequest struct {
	ctx           context.Context
	ApiService    *EnvironmentSecretApiService
	environmentId string
	secretId      string
}

func (r ApiDeleteEnvironmentSecretRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteEnvironmentSecretExecute(r)
}

/*
DeleteEnvironmentSecret Delete a secret from the environment

- To delete a secret you must have the project user permission
- You can't delete a BUILT_IN secret
- If you delete a secret having override or alias, the associated override/alias will be deleted as well  operationId: deleteEnvironmentSecret


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @param secretId Secret ID
 @return ApiDeleteEnvironmentSecretRequest
*/
func (a *EnvironmentSecretApiService) DeleteEnvironmentSecret(ctx context.Context, environmentId string, secretId string) ApiDeleteEnvironmentSecretRequest {
	return ApiDeleteEnvironmentSecretRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
		secretId:      secretId,
	}
}

// Execute executes the request
func (a *EnvironmentSecretApiService) DeleteEnvironmentSecretExecute(r ApiDeleteEnvironmentSecretRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentSecretApiService.DeleteEnvironmentSecret")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/secret/{secretId}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secretId"+"}", url.PathEscape(parameterToString(r.secretId, "")), -1)

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

type ApiEditEnvironmentSecretRequest struct {
	ctx               context.Context
	ApiService        *EnvironmentSecretApiService
	environmentId     string
	secretId          string
	secretEditRequest *SecretEditRequest
}

func (r ApiEditEnvironmentSecretRequest) SecretEditRequest(secretEditRequest SecretEditRequest) ApiEditEnvironmentSecretRequest {
	r.secretEditRequest = &secretEditRequest
	return r
}

func (r ApiEditEnvironmentSecretRequest) Execute() (*SecretResponse, *http.Response, error) {
	return r.ApiService.EditEnvironmentSecretExecute(r)
}

/*
EditEnvironmentSecret Edit a secret belonging to the environment

- You can't edit a BUILT_IN secret
- For an override, you can't edit the key
- For an alias, you can't edit the value
- An override can only have a scope lower to the secret it is overriding (hierarchy is BUILT_IN > PROJECT > ENVIRONMENT > APPLICATION)


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @param secretId Secret ID
 @return ApiEditEnvironmentSecretRequest
*/
func (a *EnvironmentSecretApiService) EditEnvironmentSecret(ctx context.Context, environmentId string, secretId string) ApiEditEnvironmentSecretRequest {
	return ApiEditEnvironmentSecretRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
		secretId:      secretId,
	}
}

// Execute executes the request
//  @return SecretResponse
func (a *EnvironmentSecretApiService) EditEnvironmentSecretExecute(r ApiEditEnvironmentSecretRequest) (*SecretResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SecretResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentSecretApiService.EditEnvironmentSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/secret/{secretId}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"secretId"+"}", url.PathEscape(parameterToString(r.secretId, "")), -1)

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

type ApiListEnvironmentSecretsRequest struct {
	ctx           context.Context
	ApiService    *EnvironmentSecretApiService
	environmentId string
}

func (r ApiListEnvironmentSecretsRequest) Execute() (*SecretResponseList, *http.Response, error) {
	return r.ApiService.ListEnvironmentSecretsExecute(r)
}

/*
ListEnvironmentSecrets List environment secrets

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiListEnvironmentSecretsRequest
*/
func (a *EnvironmentSecretApiService) ListEnvironmentSecrets(ctx context.Context, environmentId string) ApiListEnvironmentSecretsRequest {
	return ApiListEnvironmentSecretsRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return SecretResponseList
func (a *EnvironmentSecretApiService) ListEnvironmentSecretsExecute(r ApiListEnvironmentSecretsRequest) (*SecretResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SecretResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentSecretApiService.ListEnvironmentSecrets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/secret"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterToString(r.environmentId, "")), -1)

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
