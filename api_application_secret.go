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

// ApplicationSecretApiService ApplicationSecretApi service
type ApplicationSecretApiService service

type ApiCreateApplicationSecretRequest struct {
	ctx           context.Context
	ApiService    *ApplicationSecretApiService
	applicationId string
	secretRequest *SecretRequest
}

func (r ApiCreateApplicationSecretRequest) SecretRequest(secretRequest SecretRequest) ApiCreateApplicationSecretRequest {
	r.secretRequest = &secretRequest
	return r
}

func (r ApiCreateApplicationSecretRequest) Execute() (*SecretResponse, *http.Response, error) {
	return r.ApiService.CreateApplicationSecretExecute(r)
}

/*
CreateApplicationSecret Add a secret to the application

- Add a secret to the application.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiCreateApplicationSecretRequest
*/
func (a *ApplicationSecretApiService) CreateApplicationSecret(ctx context.Context, applicationId string) ApiCreateApplicationSecretRequest {
	return ApiCreateApplicationSecretRequest{
		ApiService:    a,
		ctx:           ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return SecretResponse
func (a *ApplicationSecretApiService) CreateApplicationSecretExecute(r ApiCreateApplicationSecretRequest) (*SecretResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SecretResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationSecretApiService.CreateApplicationSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/secret"
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

type ApiCreateApplicationSecretAliasRequest struct {
	ctx           context.Context
	ApiService    *ApplicationSecretApiService
	applicationId string
	secretId      string
	key           *Key
}

func (r ApiCreateApplicationSecretAliasRequest) Key(key Key) ApiCreateApplicationSecretAliasRequest {
	r.key = &key
	return r
}

func (r ApiCreateApplicationSecretAliasRequest) Execute() (*SecretResponse, *http.Response, error) {
	return r.ApiService.CreateApplicationSecretAliasExecute(r)
}

/*
CreateApplicationSecretAlias Create a secret alias at the application level

- Allows you to add an alias at application level on an existing secret having higher scope, in order to customize its key.
- You only have to specify a key in the request body
- The system will create a new secret at application level with the same value as the one corresponding to the secret id in the path
- The response body will contain the newly created secret
- Information regarding the aliased_secret will be exposed in the "aliased_secret" field of the newly created secret
- Only 1 alias level is allowed. You can't create an alias on an alias


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @param secretId Secret ID
 @return ApiCreateApplicationSecretAliasRequest
*/
func (a *ApplicationSecretApiService) CreateApplicationSecretAlias(ctx context.Context, applicationId string, secretId string) ApiCreateApplicationSecretAliasRequest {
	return ApiCreateApplicationSecretAliasRequest{
		ApiService:    a,
		ctx:           ctx,
		applicationId: applicationId,
		secretId:      secretId,
	}
}

// Execute executes the request
//  @return SecretResponse
func (a *ApplicationSecretApiService) CreateApplicationSecretAliasExecute(r ApiCreateApplicationSecretAliasRequest) (*SecretResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SecretResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationSecretApiService.CreateApplicationSecretAlias")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/secret/{secretId}/alias"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterToString(r.applicationId, "")), -1)
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

type ApiCreateApplicationSecretOverrideRequest struct {
	ctx           context.Context
	ApiService    *ApplicationSecretApiService
	applicationId string
	secretId      string
	value         *Value
}

func (r ApiCreateApplicationSecretOverrideRequest) Value(value Value) ApiCreateApplicationSecretOverrideRequest {
	r.value = &value
	return r
}

func (r ApiCreateApplicationSecretOverrideRequest) Execute() (*SecretResponse, *http.Response, error) {
	return r.ApiService.CreateApplicationSecretOverrideExecute(r)
}

/*
CreateApplicationSecretOverride Create a secret override at the application level

- Allows you to override at application level a secret that has a higher scope.
- You only have to specify a value in the request body
- The system will create a new secret at application level with the same key as the one corresponding to the secret id in the path
- The response body will contain the newly created secret
- Information regarding the overridden_secret will be exposed in the "overridden_secret" field of the newly created secret


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @param secretId Secret ID
 @return ApiCreateApplicationSecretOverrideRequest
*/
func (a *ApplicationSecretApiService) CreateApplicationSecretOverride(ctx context.Context, applicationId string, secretId string) ApiCreateApplicationSecretOverrideRequest {
	return ApiCreateApplicationSecretOverrideRequest{
		ApiService:    a,
		ctx:           ctx,
		applicationId: applicationId,
		secretId:      secretId,
	}
}

// Execute executes the request
//  @return SecretResponse
func (a *ApplicationSecretApiService) CreateApplicationSecretOverrideExecute(r ApiCreateApplicationSecretOverrideRequest) (*SecretResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SecretResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationSecretApiService.CreateApplicationSecretOverride")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/secret/{secretId}/override"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterToString(r.applicationId, "")), -1)
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

type ApiDeleteApplicationSecretRequest struct {
	ctx           context.Context
	ApiService    *ApplicationSecretApiService
	applicationId string
	secretId      string
}

func (r ApiDeleteApplicationSecretRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteApplicationSecretExecute(r)
}

/*
DeleteApplicationSecret Delete a secret from an application

- To delete a secret you must have the project user permission
- You can't delete a BUILT_IN secret
- If you delete a secret having override or alias, the associated override/alias will be deleted as well


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @param secretId Secret ID
 @return ApiDeleteApplicationSecretRequest
*/
func (a *ApplicationSecretApiService) DeleteApplicationSecret(ctx context.Context, applicationId string, secretId string) ApiDeleteApplicationSecretRequest {
	return ApiDeleteApplicationSecretRequest{
		ApiService:    a,
		ctx:           ctx,
		applicationId: applicationId,
		secretId:      secretId,
	}
}

// Execute executes the request
func (a *ApplicationSecretApiService) DeleteApplicationSecretExecute(r ApiDeleteApplicationSecretRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationSecretApiService.DeleteApplicationSecret")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/secret/{secretId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterToString(r.applicationId, "")), -1)
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

type ApiEditApplicationSecretRequest struct {
	ctx               context.Context
	ApiService        *ApplicationSecretApiService
	applicationId     string
	secretId          string
	secretEditRequest *SecretEditRequest
}

func (r ApiEditApplicationSecretRequest) SecretEditRequest(secretEditRequest SecretEditRequest) ApiEditApplicationSecretRequest {
	r.secretEditRequest = &secretEditRequest
	return r
}

func (r ApiEditApplicationSecretRequest) Execute() (*SecretResponse, *http.Response, error) {
	return r.ApiService.EditApplicationSecretExecute(r)
}

/*
EditApplicationSecret Edit a secret belonging to the application

- You can't edit a BUILT_IN secret
- For an override, you can't edit the key
- For an alias, you can't edit the value
- An override can only have a scope lower to the secret it is overriding (hierarchy is BUILT_IN > PROJECT > ENVIRONMENT > APPLICATION)


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @param secretId Secret ID
 @return ApiEditApplicationSecretRequest
*/
func (a *ApplicationSecretApiService) EditApplicationSecret(ctx context.Context, applicationId string, secretId string) ApiEditApplicationSecretRequest {
	return ApiEditApplicationSecretRequest{
		ApiService:    a,
		ctx:           ctx,
		applicationId: applicationId,
		secretId:      secretId,
	}
}

// Execute executes the request
//  @return SecretResponse
func (a *ApplicationSecretApiService) EditApplicationSecretExecute(r ApiEditApplicationSecretRequest) (*SecretResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SecretResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationSecretApiService.EditApplicationSecret")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/secret/{secretId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterToString(r.applicationId, "")), -1)
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

type ApiListApplicationSecretsRequest struct {
	ctx           context.Context
	ApiService    *ApplicationSecretApiService
	applicationId string
}

func (r ApiListApplicationSecretsRequest) Execute() (*SecretResponseList, *http.Response, error) {
	return r.ApiService.ListApplicationSecretsExecute(r)
}

/*
ListApplicationSecrets List application secrets

Secrets are like environment variables, but they are secured and can't be revealed.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiListApplicationSecretsRequest
*/
func (a *ApplicationSecretApiService) ListApplicationSecrets(ctx context.Context, applicationId string) ApiListApplicationSecretsRequest {
	return ApiListApplicationSecretsRequest{
		ApiService:    a,
		ctx:           ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return SecretResponseList
func (a *ApplicationSecretApiService) ListApplicationSecretsExecute(r ApiListApplicationSecretsRequest) (*SecretResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SecretResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationSecretApiService.ListApplicationSecrets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/secret"
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
