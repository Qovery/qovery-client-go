# DatabaseEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | name is case-insensitive | [optional] 
**Description** | Pointer to **string** | give a description to this database | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Accessibility** | Pointer to [**DatabaseAccessibilityEnum**](DatabaseAccessibilityEnum.md) |  | [optional] [default to DATABASEACCESSIBILITYENUM_PRIVATE]
**Cpu** | Pointer to **int32** | unit is millicores (m). 1000m &#x3D; 1 cpu. This field will be ignored for managed DB (instance type will be used instead).  | [optional] [default to 250]
**Memory** | Pointer to **int32** | unit is MB. 1024 MB &#x3D; 1GB This field will be ignored for managed DB (instance type will be used instead). Default value is linked to the database type: - MANAGED: 100 - CONTAINER   - POSTGRES: 100   - REDIS: 100   - MYSQL: 512   - MONGODB: 256  | [optional] 
**Storage** | Pointer to **int32** | unit is GB | [optional] 
**InstanceType** | Pointer to **string** | Database instance type to be used for this database. The list of values can be retrieved via the endpoint /{CloudProvider}/managedDatabase/instanceType/{region}/{dbType}. This field SHOULD NOT be set for container DB. | [optional] 
**AnnotationsGroupIds** | Pointer to **[]string** | list of id of the annotations groups | [optional] 

## Methods

### NewDatabaseEditRequest

`func NewDatabaseEditRequest() *DatabaseEditRequest`

NewDatabaseEditRequest instantiates a new DatabaseEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseEditRequestWithDefaults

`func NewDatabaseEditRequestWithDefaults() *DatabaseEditRequest`

NewDatabaseEditRequestWithDefaults instantiates a new DatabaseEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatabaseEditRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseEditRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseEditRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DatabaseEditRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *DatabaseEditRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatabaseEditRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatabaseEditRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatabaseEditRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *DatabaseEditRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DatabaseEditRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DatabaseEditRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DatabaseEditRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAccessibility

`func (o *DatabaseEditRequest) GetAccessibility() DatabaseAccessibilityEnum`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *DatabaseEditRequest) GetAccessibilityOk() (*DatabaseAccessibilityEnum, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *DatabaseEditRequest) SetAccessibility(v DatabaseAccessibilityEnum)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *DatabaseEditRequest) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetCpu

`func (o *DatabaseEditRequest) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *DatabaseEditRequest) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *DatabaseEditRequest) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *DatabaseEditRequest) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *DatabaseEditRequest) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *DatabaseEditRequest) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *DatabaseEditRequest) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *DatabaseEditRequest) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetStorage

`func (o *DatabaseEditRequest) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *DatabaseEditRequest) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *DatabaseEditRequest) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *DatabaseEditRequest) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetInstanceType

`func (o *DatabaseEditRequest) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *DatabaseEditRequest) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *DatabaseEditRequest) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *DatabaseEditRequest) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetAnnotationsGroupIds

`func (o *DatabaseEditRequest) GetAnnotationsGroupIds() []string`

GetAnnotationsGroupIds returns the AnnotationsGroupIds field if non-nil, zero value otherwise.

### GetAnnotationsGroupIdsOk

`func (o *DatabaseEditRequest) GetAnnotationsGroupIdsOk() (*[]string, bool)`

GetAnnotationsGroupIdsOk returns a tuple with the AnnotationsGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationsGroupIds

`func (o *DatabaseEditRequest) SetAnnotationsGroupIds(v []string)`

SetAnnotationsGroupIds sets AnnotationsGroupIds field to given value.

### HasAnnotationsGroupIds

`func (o *DatabaseEditRequest) HasAnnotationsGroupIds() bool`

HasAnnotationsGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


