# GatewayStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayName** | **string** |  | 
**Conditions** | [**[]GatewayConditionEntryResponse**](GatewayConditionEntryResponse.md) |  | 

## Methods

### NewGatewayStatusResponse

`func NewGatewayStatusResponse(gatewayName string, conditions []GatewayConditionEntryResponse, ) *GatewayStatusResponse`

NewGatewayStatusResponse instantiates a new GatewayStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayStatusResponseWithDefaults

`func NewGatewayStatusResponseWithDefaults() *GatewayStatusResponse`

NewGatewayStatusResponseWithDefaults instantiates a new GatewayStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayName

`func (o *GatewayStatusResponse) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *GatewayStatusResponse) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *GatewayStatusResponse) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.


### GetConditions

`func (o *GatewayStatusResponse) GetConditions() []GatewayConditionEntryResponse`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *GatewayStatusResponse) GetConditionsOk() (*[]GatewayConditionEntryResponse, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *GatewayStatusResponse) SetConditions(v []GatewayConditionEntryResponse)`

SetConditions sets Conditions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


