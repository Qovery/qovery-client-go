# KedaScalerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ScalerType** | **string** |  | 
**Enabled** | **bool** |  | 
**Role** | [**KedaScalerRole**](KedaScalerRole.md) |  | 

## Methods

### NewKedaScalerResponse

`func NewKedaScalerResponse(id string, createdAt time.Time, scalerType string, enabled bool, role KedaScalerRole, ) *KedaScalerResponse`

NewKedaScalerResponse instantiates a new KedaScalerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKedaScalerResponseWithDefaults

`func NewKedaScalerResponseWithDefaults() *KedaScalerResponse`

NewKedaScalerResponseWithDefaults instantiates a new KedaScalerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KedaScalerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KedaScalerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KedaScalerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *KedaScalerResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KedaScalerResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KedaScalerResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *KedaScalerResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KedaScalerResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KedaScalerResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KedaScalerResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetScalerType

`func (o *KedaScalerResponse) GetScalerType() string`

GetScalerType returns the ScalerType field if non-nil, zero value otherwise.

### GetScalerTypeOk

`func (o *KedaScalerResponse) GetScalerTypeOk() (*string, bool)`

GetScalerTypeOk returns a tuple with the ScalerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalerType

`func (o *KedaScalerResponse) SetScalerType(v string)`

SetScalerType sets ScalerType field to given value.


### GetEnabled

`func (o *KedaScalerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KedaScalerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KedaScalerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRole

`func (o *KedaScalerResponse) GetRole() KedaScalerRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *KedaScalerResponse) GetRoleOk() (*KedaScalerRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *KedaScalerResponse) SetRole(v KedaScalerRole)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


