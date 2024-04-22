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
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ApplicationAnnotationsGroupAPIService ApplicationAnnotationsGroupAPI service
type ApplicationAnnotationsGroupAPIService service

type ApiAddAnnotationsGroupToApplicationRequest struct {
	ctx                context.Context
	ApiService         *ApplicationAnnotationsGroupAPIService
	applicationId      string
	annotationsGroupId string
}

func (r ApiAddAnnotationsGroupToApplicationRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddAnnotationsGroupToApplicationExecute(r)
}

/*
AddAnnotationsGroupToApplication Add annotations group to application

Add annotations group to application

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @param annotationsGroupId Organization annotations group ID
 @return ApiAddAnnotationsGroupToApplicationRequest
*/
func (a *ApplicationAnnotationsGroupAPIService) AddAnnotationsGroupToApplication(ctx context.Context, applicationId string, annotationsGroupId string) ApiAddAnnotationsGroupToApplicationRequest {
	return ApiAddAnnotationsGroupToApplicationRequest{
		ApiService:         a,
		ctx:                ctx,
		applicationId:      applicationId,
		annotationsGroupId: annotationsGroupId,
	}
}

// Execute executes the request
func (a *ApplicationAnnotationsGroupAPIService) AddAnnotationsGroupToApplicationExecute(r ApiAddAnnotationsGroupToApplicationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationAnnotationsGroupAPIService.AddAnnotationsGroupToApplication")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/annotationsGroup/{annotationsGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"annotationsGroupId"+"}", url.PathEscape(parameterValueToString(r.annotationsGroupId, "annotationsGroupId")), -1)

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

type ApiDeleteAnnotationsGroupToApplicationRequest struct {
	ctx                context.Context
	ApiService         *ApplicationAnnotationsGroupAPIService
	applicationId      string
	annotationsGroupId string
}

func (r ApiDeleteAnnotationsGroupToApplicationRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAnnotationsGroupToApplicationExecute(r)
}

/*
DeleteAnnotationsGroupToApplication Delete annotations group to application

Delete annotations group to application

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @param annotationsGroupId Organization annotations group ID
 @return ApiDeleteAnnotationsGroupToApplicationRequest
*/
func (a *ApplicationAnnotationsGroupAPIService) DeleteAnnotationsGroupToApplication(ctx context.Context, applicationId string, annotationsGroupId string) ApiDeleteAnnotationsGroupToApplicationRequest {
	return ApiDeleteAnnotationsGroupToApplicationRequest{
		ApiService:         a,
		ctx:                ctx,
		applicationId:      applicationId,
		annotationsGroupId: annotationsGroupId,
	}
}

// Execute executes the request
func (a *ApplicationAnnotationsGroupAPIService) DeleteAnnotationsGroupToApplicationExecute(r ApiDeleteAnnotationsGroupToApplicationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationAnnotationsGroupAPIService.DeleteAnnotationsGroupToApplication")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/annotationsGroup/{annotationsGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"annotationsGroupId"+"}", url.PathEscape(parameterValueToString(r.annotationsGroupId, "annotationsGroupId")), -1)

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

type ApiListApplicationAnnotationsGroupRequest struct {
	ctx           context.Context
	ApiService    *ApplicationAnnotationsGroupAPIService
	applicationId string
}

func (r ApiListApplicationAnnotationsGroupRequest) Execute() ([]OrganizationAnnotationsGroupResponse, *http.Response, error) {
	return r.ApiService.ListApplicationAnnotationsGroupExecute(r)
}

/*
ListApplicationAnnotationsGroup List application annotations group

List application annotations group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID
 @return ApiListApplicationAnnotationsGroupRequest
*/
func (a *ApplicationAnnotationsGroupAPIService) ListApplicationAnnotationsGroup(ctx context.Context, applicationId string) ApiListApplicationAnnotationsGroupRequest {
	return ApiListApplicationAnnotationsGroupRequest{
		ApiService:    a,
		ctx:           ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return []OrganizationAnnotationsGroupResponse
func (a *ApplicationAnnotationsGroupAPIService) ListApplicationAnnotationsGroupExecute(r ApiListApplicationAnnotationsGroupRequest) ([]OrganizationAnnotationsGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []OrganizationAnnotationsGroupResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationAnnotationsGroupAPIService.ListApplicationAnnotationsGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application/{applicationId}/annotationsGroup"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)

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
