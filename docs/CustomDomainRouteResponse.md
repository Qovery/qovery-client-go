# CustomDomainRouteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the route. | 
**CustomDomainId** | **string** | Id of this domain in the custom domain API of the service (&#x60;/application|container|helm/{id}/customDomain/{customDomainId}&#x60;). | 
**Service** | [**CustomDomainServiceResponse**](CustomDomainServiceResponse.md) |  | 
**Port** | [**NullableCustomDomainPortResponse**](CustomDomainPortResponse.md) |  | 
**Path** | **string** | The path of the requests this route serves. | 
**PathRewrite** | **NullableString** | The path that replaces &#x60;path&#x60; before the request reaches the service. Null when the path is not rewritten. | 

## Methods

### NewCustomDomainRouteResponse

`func NewCustomDomainRouteResponse(id string, customDomainId string, service CustomDomainServiceResponse, port NullableCustomDomainPortResponse, path string, pathRewrite NullableString, ) *CustomDomainRouteResponse`

NewCustomDomainRouteResponse instantiates a new CustomDomainRouteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainRouteResponseWithDefaults

`func NewCustomDomainRouteResponseWithDefaults() *CustomDomainRouteResponse`

NewCustomDomainRouteResponseWithDefaults instantiates a new CustomDomainRouteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomDomainRouteResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomDomainRouteResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomDomainRouteResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCustomDomainId

`func (o *CustomDomainRouteResponse) GetCustomDomainId() string`

GetCustomDomainId returns the CustomDomainId field if non-nil, zero value otherwise.

### GetCustomDomainIdOk

`func (o *CustomDomainRouteResponse) GetCustomDomainIdOk() (*string, bool)`

GetCustomDomainIdOk returns a tuple with the CustomDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomainId

`func (o *CustomDomainRouteResponse) SetCustomDomainId(v string)`

SetCustomDomainId sets CustomDomainId field to given value.


### GetService

`func (o *CustomDomainRouteResponse) GetService() CustomDomainServiceResponse`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CustomDomainRouteResponse) GetServiceOk() (*CustomDomainServiceResponse, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CustomDomainRouteResponse) SetService(v CustomDomainServiceResponse)`

SetService sets Service field to given value.


### GetPort

`func (o *CustomDomainRouteResponse) GetPort() CustomDomainPortResponse`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CustomDomainRouteResponse) GetPortOk() (*CustomDomainPortResponse, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CustomDomainRouteResponse) SetPort(v CustomDomainPortResponse)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *CustomDomainRouteResponse) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *CustomDomainRouteResponse) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetPath

`func (o *CustomDomainRouteResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CustomDomainRouteResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CustomDomainRouteResponse) SetPath(v string)`

SetPath sets Path field to given value.


### GetPathRewrite

`func (o *CustomDomainRouteResponse) GetPathRewrite() string`

GetPathRewrite returns the PathRewrite field if non-nil, zero value otherwise.

### GetPathRewriteOk

`func (o *CustomDomainRouteResponse) GetPathRewriteOk() (*string, bool)`

GetPathRewriteOk returns a tuple with the PathRewrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathRewrite

`func (o *CustomDomainRouteResponse) SetPathRewrite(v string)`

SetPathRewrite sets PathRewrite field to given value.


### SetPathRewriteNil

`func (o *CustomDomainRouteResponse) SetPathRewriteNil(b bool)`

 SetPathRewriteNil sets the value for PathRewrite to be an explicit nil

### UnsetPathRewrite
`func (o *CustomDomainRouteResponse) UnsetPathRewrite()`

UnsetPathRewrite ensures that no value is present for PathRewrite, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


