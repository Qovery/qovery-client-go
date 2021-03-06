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

// ContainerEventApiService ContainerEventApi service
type ContainerEventApiService service

type ApiListContainerEventRequest struct {
	ctx         context.Context
	ApiService  *ContainerEventApiService
	containerId string
	startId     *string
}

// Starting point after which to return results
func (r ApiListContainerEventRequest) StartId(startId string) ApiListContainerEventRequest {
	r.startId = &startId
	return r
}

func (r ApiListContainerEventRequest) Execute() (*EventPaginatedResponseList, *http.Response, error) {
	return r.ApiService.ListContainerEventExecute(r)
}

/*
ListContainerEvent NOT YET IMPLEMENTED - List container events

By default it returns the 20 last results. The response is paginated. In order to request the next page, you can use the startId query parameter

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param containerId Container ID
 @return ApiListContainerEventRequest
*/
func (a *ContainerEventApiService) ListContainerEvent(ctx context.Context, containerId string) ApiListContainerEventRequest {
	return ApiListContainerEventRequest{
		ApiService:  a,
		ctx:         ctx,
		containerId: containerId,
	}
}

// Execute executes the request
//  @return EventPaginatedResponseList
func (a *ContainerEventApiService) ListContainerEventExecute(r ApiListContainerEventRequest) (*EventPaginatedResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EventPaginatedResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContainerEventApiService.ListContainerEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/container/{containerId}/event"
	localVarPath = strings.Replace(localVarPath, "{"+"containerId"+"}", url.PathEscape(parameterToString(r.containerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startId != nil {
		localVarQueryParams.Add("startId", parameterToString(*r.startId, ""))
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
