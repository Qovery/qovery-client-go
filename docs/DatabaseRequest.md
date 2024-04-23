# DatabaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **string** | give a description to this database | [optional] 
**Type** | [**DatabaseTypeEnum**](DatabaseTypeEnum.md) |  | 
**Version** | **string** |  | 
**Mode** | [**DatabaseModeEnum**](DatabaseModeEnum.md) |  | 
**Accessibility** | Pointer to [**DatabaseAccessibilityEnum**](DatabaseAccessibilityEnum.md) |  | [optional] [default to DATABASEACCESSIBILITYENUM_PRIVATE]
**Cpu** | Pointer to **int32** | unit is millicores (m). 1000m &#x3D; 1 cpu This field will be ignored for managed DB (instance type will be used instead).  | [optional] [default to 250]
**InstanceType** | Pointer to **string** | Database instance type to be used for this database. The list of values can be retrieved via the endpoint /{CloudProvider}/managedDatabase/instanceType/{region}/{dbType}. This field SHOULD NOT be set for container DB. | [optional] 
**Memory** | Pointer to **int32** | unit is MB. 1024 MB &#x3D; 1GB This field will be ignored for managed DB (instance type will be used instead). Default value is linked to the database type: - MANAGED: &#x60;100&#x60; - CONTAINER   - POSTGRES: &#x60;100&#x60;   - REDIS: &#x60;100&#x60;   - MYSQL: &#x60;512&#x60;   - MONGODB: &#x60;256&#x60;  | [optional] 
**Storage** | Pointer to **int32** | unit is GB | [optional] [default to 10]
**AnnotationsGroups** | Pointer to [**[]ServiceAnnotationRequest**](ServiceAnnotationRequest.md) |  | [optional] 

## Methods

### NewDatabaseRequest

`func NewDatabaseRequest(name string, type_ DatabaseTypeEnum, version string, mode DatabaseModeEnum, ) *DatabaseRequest`

NewDatabaseRequest instantiates a new DatabaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseRequestWithDefaults

`func NewDatabaseRequestWithDefaults() *DatabaseRequest`

NewDatabaseRequestWithDefaults instantiates a new DatabaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatabaseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DatabaseRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatabaseRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatabaseRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatabaseRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *DatabaseRequest) GetType() DatabaseTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseRequest) GetTypeOk() (*DatabaseTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseRequest) SetType(v DatabaseTypeEnum)`

SetType sets Type field to given value.


### GetVersion

`func (o *DatabaseRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DatabaseRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DatabaseRequest) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetMode

`func (o *DatabaseRequest) GetMode() DatabaseModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DatabaseRequest) GetModeOk() (*DatabaseModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DatabaseRequest) SetMode(v DatabaseModeEnum)`

SetMode sets Mode field to given value.


### GetAccessibility

`func (o *DatabaseRequest) GetAccessibility() DatabaseAccessibilityEnum`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *DatabaseRequest) GetAccessibilityOk() (*DatabaseAccessibilityEnum, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *DatabaseRequest) SetAccessibility(v DatabaseAccessibilityEnum)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *DatabaseRequest) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetCpu

`func (o *DatabaseRequest) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *DatabaseRequest) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *DatabaseRequest) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *DatabaseRequest) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetInstanceType

`func (o *DatabaseRequest) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *DatabaseRequest) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *DatabaseRequest) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *DatabaseRequest) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetMemory

`func (o *DatabaseRequest) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *DatabaseRequest) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *DatabaseRequest) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *DatabaseRequest) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetStorage

`func (o *DatabaseRequest) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *DatabaseRequest) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *DatabaseRequest) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *DatabaseRequest) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetAnnotationsGroups

`func (o *DatabaseRequest) GetAnnotationsGroups() []ServiceAnnotationRequest`

GetAnnotationsGroups returns the AnnotationsGroups field if non-nil, zero value otherwise.

### GetAnnotationsGroupsOk

`func (o *DatabaseRequest) GetAnnotationsGroupsOk() (*[]ServiceAnnotationRequest, bool)`

GetAnnotationsGroupsOk returns a tuple with the AnnotationsGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationsGroups

`func (o *DatabaseRequest) SetAnnotationsGroups(v []ServiceAnnotationRequest)`

SetAnnotationsGroups sets AnnotationsGroups field to given value.

### HasAnnotationsGroups

`func (o *DatabaseRequest) HasAnnotationsGroups() bool`

HasAnnotationsGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


