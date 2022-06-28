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

// ContainerSecretApiService ContainerSecretApi service
type ContainerSecretApiService service

type ApiCreateContainerSecretRequest struct {
	ctx           context.Context
	ApiService    *ContainerSecretApiService
	containerId   string
	secretRequest *SecretRequest
}

func (r ApiCreateContainerSecretRequest) SecretRequest(secretRequest SecretRequest) ApiCreateContainerSecretRequest {
	r.secretRequest = &secretRequest
	return r
}

func (r ApiCreateContainerSecretRequest) Execute() (*Secret, *http.Response, error) {
	return r.ApiService.CreateContainerSecretExecute(r)
}

/*
CreateContainerSecret Add a secret to the container

- Add a secret to the container.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param containerId Container ID
 @return ApiCreateContainerSecretRequest
*/
func (a *ContainerSecretApiService) CreateContainerSecret(ctx context.Context, containerId string) ApiCreateContainerSecretRequest {
	return ApiCreateContainerSecretRequest{
		ApiService:  a,
		ctx:         ctx,
		containerId: containerId,
	}
}

// Execute executes the request
//  @return Secret
func (a *ContainerSecretApiService) CreateContainerSecretExecute(r ApiCreateContainerSecretRequest) (*Secret, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Secret
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContainerSecretApiService.CreateContainerSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/container/{containerId}/secret"
	localVarPath = strings.Replace(localVarPath, "{"+"containerId"+"}", url.PathEscape(parameterToString(r.containerId, "")), -1)

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

type ApiCreateContainerSecretAliasRequest struct {
	ctx         context.Context
	ApiService  *ContainerSecretApiService
	containerId string
	secretId    string
	key         *Key
}

func (r ApiCreateContainerSecretAliasRequest) Key(key Key) ApiCreateContainerSecretAliasRequest {
	r.key = &key
	return r
}

func (r ApiCreateContainerSecretAliasRequest) Execute() (*Secret, *http.Response, error) {
	return r.ApiService.CreateContainerSecretAliasExecute(r)
}

/*
CreateContainerSecretAlias Create a secret alias at the container level

- Allows you to add an alias at container level on an existing secret having higher scope, in order to customize its key.
- You only have to specify a key in the request body
- The system will create a new secret at container level with the same value as the one corresponding to the secret id in the path
- The response body will contain the newly created secret
- Information regarding the aliased_secret will be exposed in the "aliased_secret" field of the newly created secret
- Only 1 alias level is allowed. You can't create an alias on an alias


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param containerId Container ID
 @param secretId Secret ID
 @return ApiCreateContainerSecretAliasRequest
*/
func (a *ContainerSecretApiService) CreateContainerSecretAlias(ctx context.Context, containerId string, secretId string) ApiCreateContainerSecretAliasRequest {
	return ApiCreateContainerSecretAliasRequest{
		ApiService:  a,
		ctx:         ctx,
		containerId: containerId,
		secretId:    secretId,
	}
}

// Execute executes the request
//  @return Secret
func (a *ContainerSecretApiService) CreateContainerSecretAliasExecute(r ApiCreateContainerSecretAliasRequest) (*Secret, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Secret
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContainerSecretApiService.CreateContainerSecretAlias")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/container/{containerId}/secret/{secretId}/alias"
	localVarPath = strings.Replace(localVarPath, "{"+"containerId"+"}", url.PathEscape(parameterToString(r.containerId, "")), -1)
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

type ApiDeleteContainerSecretRequest struct {
	ctx         context.Context
	ApiService  *ContainerSecretApiService
	containerId string
	secretId    string
}

func (r ApiDeleteContainerSecretRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteContainerSecretExecute(r)
}

/*
DeleteContainerSecret Delete a secret from an container

- To delete a secret you must have the project user permission
- You can't delete a BUILT_IN secret
- If you delete a secret having override or alias, the associated override/alias will be deleted as well


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param containerId Container ID
 @param secretId Secret ID
 @return ApiDeleteContainerSecretRequest
*/
func (a *ContainerSecretApiService) DeleteContainerSecret(ctx context.Context, containerId string, secretId string) ApiDeleteContainerSecretRequest {
	return ApiDeleteContainerSecretRequest{
		ApiService:  a,
		ctx:         ctx,
		containerId: containerId,
		secretId:    secretId,
	}
}

// Execute executes the request
func (a *ContainerSecretApiService) DeleteContainerSecretExecute(r ApiDeleteContainerSecretRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContainerSecretApiService.DeleteContainerSecret")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/container/{containerId}/secret/{secretId}"
	localVarPath = strings.Replace(localVarPath, "{"+"containerId"+"}", url.PathEscape(parameterToString(r.containerId, "")), -1)
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

type ApiEditContainerSecretRequest struct {
	ctx               context.Context
	ApiService        *ContainerSecretApiService
	containerId       string
	secretId          string
	secretEditRequest *SecretEditRequest
}

func (r ApiEditContainerSecretRequest) SecretEditRequest(secretEditRequest SecretEditRequest) ApiEditContainerSecretRequest {
	r.secretEditRequest = &secretEditRequest
	return r
}

func (r ApiEditContainerSecretRequest) Execute() (*Secret, *http.Response, error) {
	return r.ApiService.EditContainerSecretExecute(r)
}

/*
EditContainerSecret Edit a secret belonging to the container

- You can't edit a BUILT_IN secret
- For an override, you can't edit the key
- For an alias, you can't edit the value
- An override can only have a scope lower to the secret it is overriding (hierarchy is BUILT_IN > PROJECT > ENVIRONMENT > CONTAINER)


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param containerId Container ID
 @param secretId Secret ID
 @return ApiEditContainerSecretRequest
*/
func (a *ContainerSecretApiService) EditContainerSecret(ctx context.Context, containerId string, secretId string) ApiEditContainerSecretRequest {
	return ApiEditContainerSecretRequest{
		ApiService:  a,
		ctx:         ctx,
		containerId: containerId,
		secretId:    secretId,
	}
}

// Execute executes the request
//  @return Secret
func (a *ContainerSecretApiService) EditContainerSecretExecute(r ApiEditContainerSecretRequest) (*Secret, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Secret
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContainerSecretApiService.EditContainerSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/container/{containerId}/secret/{secretId}"
	localVarPath = strings.Replace(localVarPath, "{"+"containerId"+"}", url.PathEscape(parameterToString(r.containerId, "")), -1)
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

type ApiListContainerSecretsRequest struct {
	ctx         context.Context
	ApiService  *ContainerSecretApiService
	containerId string
}

func (r ApiListContainerSecretsRequest) Execute() (*SecretResponseList, *http.Response, error) {
	return r.ApiService.ListContainerSecretsExecute(r)
}

/*
ListContainerSecrets List container secrets

Secrets are like environment variables, but they are secured and can't be revealed.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param containerId Container ID
 @return ApiListContainerSecretsRequest
*/
func (a *ContainerSecretApiService) ListContainerSecrets(ctx context.Context, containerId string) ApiListContainerSecretsRequest {
	return ApiListContainerSecretsRequest{
		ApiService:  a,
		ctx:         ctx,
		containerId: containerId,
	}
}

// Execute executes the request
//  @return SecretResponseList
func (a *ContainerSecretApiService) ListContainerSecretsExecute(r ApiListContainerSecretsRequest) (*SecretResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SecretResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContainerSecretApiService.ListContainerSecrets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/container/{containerId}/secret"
	localVarPath = strings.Replace(localVarPath, "{"+"containerId"+"}", url.PathEscape(parameterToString(r.containerId, "")), -1)

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