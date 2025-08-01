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

// OrganizationEventAPIService OrganizationEventAPI service
type OrganizationEventAPIService service

type ApiGetOrganizationEventTargetsRequest struct {
	ctx            context.Context
	ApiService     *OrganizationEventAPIService
	organizationId string
	fromTimestamp  *string
	toTimestamp    *string
	eventType      *OrganizationEventType
	targetType     *OrganizationEventTargetType
	triggeredBy    *string
	origin         *OrganizationEventOrigin
	projectId      *string
	environmentId  *string
}

// Display targets available since this timestamp.   A range of date can be specified by using &#x60;from-timestamp&#x60; with &#x60;to-timestamp&#x60; The format is a timestamp with nano precision
func (r ApiGetOrganizationEventTargetsRequest) FromTimestamp(fromTimestamp string) ApiGetOrganizationEventTargetsRequest {
	r.fromTimestamp = &fromTimestamp
	return r
}

// Display targets triggered before this timestamp.   A range of date can be specified by using &#x60;to-timestamp&#x60; with &#x60;from-timestamp&#x60; The format is a timestamp with nano precision
func (r ApiGetOrganizationEventTargetsRequest) ToTimestamp(toTimestamp string) ApiGetOrganizationEventTargetsRequest {
	r.toTimestamp = &toTimestamp
	return r
}

func (r ApiGetOrganizationEventTargetsRequest) EventType(eventType OrganizationEventType) ApiGetOrganizationEventTargetsRequest {
	r.eventType = &eventType
	return r
}

func (r ApiGetOrganizationEventTargetsRequest) TargetType(targetType OrganizationEventTargetType) ApiGetOrganizationEventTargetsRequest {
	r.targetType = &targetType
	return r
}

// Information about the owner of the event (user name / apitoken / automatic action)
func (r ApiGetOrganizationEventTargetsRequest) TriggeredBy(triggeredBy string) ApiGetOrganizationEventTargetsRequest {
	r.triggeredBy = &triggeredBy
	return r
}

func (r ApiGetOrganizationEventTargetsRequest) Origin(origin OrganizationEventOrigin) ApiGetOrganizationEventTargetsRequest {
	r.origin = &origin
	return r
}

// Mandatory when requesting an environment or a service
func (r ApiGetOrganizationEventTargetsRequest) ProjectId(projectId string) ApiGetOrganizationEventTargetsRequest {
	r.projectId = &projectId
	return r
}

// Mandatory when requesting a service
func (r ApiGetOrganizationEventTargetsRequest) EnvironmentId(environmentId string) ApiGetOrganizationEventTargetsRequest {
	r.environmentId = &environmentId
	return r
}

func (r ApiGetOrganizationEventTargetsRequest) Execute() (*OrganizationEventTargetResponseList, *http.Response, error) {
	return r.ApiService.GetOrganizationEventTargetsExecute(r)
}

/*
GetOrganizationEventTargets Get available event targets to filter events

Get available event targets to filter events

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organizationId Organization ID
	@return ApiGetOrganizationEventTargetsRequest
*/
func (a *OrganizationEventAPIService) GetOrganizationEventTargets(ctx context.Context, organizationId string) ApiGetOrganizationEventTargetsRequest {
	return ApiGetOrganizationEventTargetsRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//
//	@return OrganizationEventTargetResponseList
func (a *OrganizationEventAPIService) GetOrganizationEventTargetsExecute(r ApiGetOrganizationEventTargetsRequest) (*OrganizationEventTargetResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrganizationEventTargetResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationEventAPIService.GetOrganizationEventTargets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organization/{organizationId}/targets"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fromTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fromTimestamp", r.fromTimestamp, "")
	}
	if r.toTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "toTimestamp", r.toTimestamp, "")
	}
	if r.eventType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", r.eventType, "")
	}
	if r.targetType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "targetType", r.targetType, "")
	}
	if r.triggeredBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "triggeredBy", r.triggeredBy, "")
	}
	if r.origin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "origin", r.origin, "")
	}
	if r.projectId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "projectId", r.projectId, "")
	}
	if r.environmentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "environmentId", r.environmentId, "")
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

type ApiGetOrganizationEventsRequest struct {
	ctx            context.Context
	ApiService     *OrganizationEventAPIService
	organizationId string
	pageSize       *float32
	fromTimestamp  *string
	toTimestamp    *string
	continueToken  *string
	stepBackToken  *string
	eventType      *OrganizationEventType
	targetType     *OrganizationEventTargetType
	targetId       *string
	subTargetType  *OrganizationEventSubTargetType
	triggeredBy    *string
	origin         *OrganizationEventOrigin
}

// The number of events to display in the current page
func (r ApiGetOrganizationEventsRequest) PageSize(pageSize float32) ApiGetOrganizationEventsRequest {
	r.pageSize = &pageSize
	return r
}

// Display events triggered since this timestamp.   A range of date can be specified by using &#x60;from-timestamp&#x60; with &#x60;to-timestamp&#x60; The format is a timestamp with nano precision
func (r ApiGetOrganizationEventsRequest) FromTimestamp(fromTimestamp string) ApiGetOrganizationEventsRequest {
	r.fromTimestamp = &fromTimestamp
	return r
}

// Display events triggered before this timestamp.   A range of date can be specified by using &#x60;to-timestamp&#x60; with &#x60;from-timestamp&#x60; The format is a timestamp with nano precision
func (r ApiGetOrganizationEventsRequest) ToTimestamp(toTimestamp string) ApiGetOrganizationEventsRequest {
	r.toTimestamp = &toTimestamp
	return r
}

// Token used to fetch the next page results The format is a timestamp with nano precision
func (r ApiGetOrganizationEventsRequest) ContinueToken(continueToken string) ApiGetOrganizationEventsRequest {
	r.continueToken = &continueToken
	return r
}

// Token used to fetch the previous page results The format is a timestamp with nano precision
func (r ApiGetOrganizationEventsRequest) StepBackToken(stepBackToken string) ApiGetOrganizationEventsRequest {
	r.stepBackToken = &stepBackToken
	return r
}

func (r ApiGetOrganizationEventsRequest) EventType(eventType OrganizationEventType) ApiGetOrganizationEventsRequest {
	r.eventType = &eventType
	return r
}

func (r ApiGetOrganizationEventsRequest) TargetType(targetType OrganizationEventTargetType) ApiGetOrganizationEventsRequest {
	r.targetType = &targetType
	return r
}

// The target resource id to search.   Must be specified with the corresponding &#x60;target_type&#x60;
func (r ApiGetOrganizationEventsRequest) TargetId(targetId string) ApiGetOrganizationEventsRequest {
	r.targetId = &targetId
	return r
}

func (r ApiGetOrganizationEventsRequest) SubTargetType(subTargetType OrganizationEventSubTargetType) ApiGetOrganizationEventsRequest {
	r.subTargetType = &subTargetType
	return r
}

// Information about the owner of the event (user name / apitoken / automatic action)
func (r ApiGetOrganizationEventsRequest) TriggeredBy(triggeredBy string) ApiGetOrganizationEventsRequest {
	r.triggeredBy = &triggeredBy
	return r
}

func (r ApiGetOrganizationEventsRequest) Origin(origin OrganizationEventOrigin) ApiGetOrganizationEventsRequest {
	r.origin = &origin
	return r
}

func (r ApiGetOrganizationEventsRequest) Execute() (*OrganizationEventResponseList, *http.Response, error) {
	return r.ApiService.GetOrganizationEventsExecute(r)
}

/*
GetOrganizationEvents Get all events inside the organization

Get all events inside the organization

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organizationId Organization ID
	@return ApiGetOrganizationEventsRequest
*/
func (a *OrganizationEventAPIService) GetOrganizationEvents(ctx context.Context, organizationId string) ApiGetOrganizationEventsRequest {
	return ApiGetOrganizationEventsRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//
//	@return OrganizationEventResponseList
func (a *OrganizationEventAPIService) GetOrganizationEventsExecute(r ApiGetOrganizationEventsRequest) (*OrganizationEventResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OrganizationEventResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationEventAPIService.GetOrganizationEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organization/{organizationId}/events"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	} else {
		var defaultValue float32 = 10
		r.pageSize = &defaultValue
	}
	if r.fromTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fromTimestamp", r.fromTimestamp, "")
	}
	if r.toTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "toTimestamp", r.toTimestamp, "")
	}
	if r.continueToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "continueToken", r.continueToken, "")
	}
	if r.stepBackToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "stepBackToken", r.stepBackToken, "")
	}
	if r.eventType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "eventType", r.eventType, "")
	}
	if r.targetType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "targetType", r.targetType, "")
	}
	if r.targetId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "targetId", r.targetId, "")
	}
	if r.subTargetType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subTargetType", r.subTargetType, "")
	}
	if r.triggeredBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "triggeredBy", r.triggeredBy, "")
	}
	if r.origin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "origin", r.origin, "")
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
