# IngressDeploymentStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouterId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**StateEnum**](StateEnum.md) |  | [optional] 

## Methods

### NewIngressDeploymentStatusResponse

`func NewIngressDeploymentStatusResponse() *IngressDeploymentStatusResponse`

NewIngressDeploymentStatusResponse instantiates a new IngressDeploymentStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressDeploymentStatusResponseWithDefaults

`func NewIngressDeploymentStatusResponseWithDefaults() *IngressDeploymentStatusResponse`

NewIngressDeploymentStatusResponseWithDefaults instantiates a new IngressDeploymentStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouterId

`func (o *IngressDeploymentStatusResponse) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *IngressDeploymentStatusResponse) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *IngressDeploymentStatusResponse) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *IngressDeploymentStatusResponse) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetStatus

`func (o *IngressDeploymentStatusResponse) GetStatus() StateEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IngressDeploymentStatusResponse) GetStatusOk() (*StateEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IngressDeploymentStatusResponse) SetStatus(v StateEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IngressDeploymentStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


