# ClusterFeatureNatGatewayParametersNatGatewayType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**Type** | **string** |  | 
**StaticIpsEnabled** | **bool** |  | 
**StaticIpsCount** | **int32** | Number of static IPs to allocate. Must be &gt;&#x3D; 1 when static_ips_enabled is true. | 

## Methods

### NewClusterFeatureNatGatewayParametersNatGatewayType

`func NewClusterFeatureNatGatewayParametersNatGatewayType(provider string, type_ string, staticIpsEnabled bool, staticIpsCount int32, ) *ClusterFeatureNatGatewayParametersNatGatewayType`

NewClusterFeatureNatGatewayParametersNatGatewayType instantiates a new ClusterFeatureNatGatewayParametersNatGatewayType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterFeatureNatGatewayParametersNatGatewayTypeWithDefaults

`func NewClusterFeatureNatGatewayParametersNatGatewayTypeWithDefaults() *ClusterFeatureNatGatewayParametersNatGatewayType`

NewClusterFeatureNatGatewayParametersNatGatewayTypeWithDefaults instantiates a new ClusterFeatureNatGatewayParametersNatGatewayType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *ClusterFeatureNatGatewayParametersNatGatewayType) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ClusterFeatureNatGatewayParametersNatGatewayType) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ClusterFeatureNatGatewayParametersNatGatewayType) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetType

`func (o *ClusterFeatureNatGatewayParametersNatGatewayType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterFeatureNatGatewayParametersNatGatewayType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterFeatureNatGatewayParametersNatGatewayType) SetType(v string)`

SetType sets Type field to given value.


### GetStaticIpsEnabled

`func (o *ClusterFeatureNatGatewayParametersNatGatewayType) GetStaticIpsEnabled() bool`

GetStaticIpsEnabled returns the StaticIpsEnabled field if non-nil, zero value otherwise.

### GetStaticIpsEnabledOk

`func (o *ClusterFeatureNatGatewayParametersNatGatewayType) GetStaticIpsEnabledOk() (*bool, bool)`

GetStaticIpsEnabledOk returns a tuple with the StaticIpsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIpsEnabled

`func (o *ClusterFeatureNatGatewayParametersNatGatewayType) SetStaticIpsEnabled(v bool)`

SetStaticIpsEnabled sets StaticIpsEnabled field to given value.


### GetStaticIpsCount

`func (o *ClusterFeatureNatGatewayParametersNatGatewayType) GetStaticIpsCount() int32`

GetStaticIpsCount returns the StaticIpsCount field if non-nil, zero value otherwise.

### GetStaticIpsCountOk

`func (o *ClusterFeatureNatGatewayParametersNatGatewayType) GetStaticIpsCountOk() (*int32, bool)`

GetStaticIpsCountOk returns a tuple with the StaticIpsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIpsCount

`func (o *ClusterFeatureNatGatewayParametersNatGatewayType) SetStaticIpsCount(v int32)`

SetStaticIpsCount sets StaticIpsCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


