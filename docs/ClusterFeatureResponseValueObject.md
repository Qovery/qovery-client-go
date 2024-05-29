# ClusterFeatureResponseValueObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ClusterFeatureResponseTypeEnum**](ClusterFeatureResponseTypeEnum.md) |  | 
**Value** | [**ClusterFeatureGcpExistingVpc**](ClusterFeatureGcpExistingVpc.md) |  | 

## Methods

### NewClusterFeatureResponseValueObject

`func NewClusterFeatureResponseValueObject(type_ ClusterFeatureResponseTypeEnum, value ClusterFeatureGcpExistingVpc, ) *ClusterFeatureResponseValueObject`

NewClusterFeatureResponseValueObject instantiates a new ClusterFeatureResponseValueObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterFeatureResponseValueObjectWithDefaults

`func NewClusterFeatureResponseValueObjectWithDefaults() *ClusterFeatureResponseValueObject`

NewClusterFeatureResponseValueObjectWithDefaults instantiates a new ClusterFeatureResponseValueObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ClusterFeatureResponseValueObject) GetType() ClusterFeatureResponseTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterFeatureResponseValueObject) GetTypeOk() (*ClusterFeatureResponseTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterFeatureResponseValueObject) SetType(v ClusterFeatureResponseTypeEnum)`

SetType sets Type field to given value.


### GetValue

`func (o *ClusterFeatureResponseValueObject) GetValue() ClusterFeatureGcpExistingVpc`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ClusterFeatureResponseValueObject) GetValueOk() (*ClusterFeatureGcpExistingVpc, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ClusterFeatureResponseValueObject) SetValue(v ClusterFeatureGcpExistingVpc)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


