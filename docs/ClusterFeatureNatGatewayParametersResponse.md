# ClusterFeatureNatGatewayParametersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ClusterFeatureResponseTypeEnum**](ClusterFeatureResponseTypeEnum.md) |  | 
**Value** | [**NullableClusterFeatureNatGatewayParameters**](ClusterFeatureNatGatewayParameters.md) |  | 

## Methods

### NewClusterFeatureNatGatewayParametersResponse

`func NewClusterFeatureNatGatewayParametersResponse(type_ ClusterFeatureResponseTypeEnum, value NullableClusterFeatureNatGatewayParameters, ) *ClusterFeatureNatGatewayParametersResponse`

NewClusterFeatureNatGatewayParametersResponse instantiates a new ClusterFeatureNatGatewayParametersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterFeatureNatGatewayParametersResponseWithDefaults

`func NewClusterFeatureNatGatewayParametersResponseWithDefaults() *ClusterFeatureNatGatewayParametersResponse`

NewClusterFeatureNatGatewayParametersResponseWithDefaults instantiates a new ClusterFeatureNatGatewayParametersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ClusterFeatureNatGatewayParametersResponse) GetType() ClusterFeatureResponseTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterFeatureNatGatewayParametersResponse) GetTypeOk() (*ClusterFeatureResponseTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterFeatureNatGatewayParametersResponse) SetType(v ClusterFeatureResponseTypeEnum)`

SetType sets Type field to given value.


### GetValue

`func (o *ClusterFeatureNatGatewayParametersResponse) GetValue() ClusterFeatureNatGatewayParameters`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ClusterFeatureNatGatewayParametersResponse) GetValueOk() (*ClusterFeatureNatGatewayParameters, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ClusterFeatureNatGatewayParametersResponse) SetValue(v ClusterFeatureNatGatewayParameters)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *ClusterFeatureNatGatewayParametersResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ClusterFeatureNatGatewayParametersResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


