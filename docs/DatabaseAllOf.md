# DatabaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to [**ReferenceObject**](ReferenceObject.md) |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**MaximumCpu** | Pointer to **int32** | Maximum cpu that can be allocated to the database based on organization cluster configuration. unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] [default to 250]
**MaximumMemory** | Pointer to **int32** | Maximum memory that can be allocated to the database based on organization cluster configuration. unit is MB. 1024 MB &#x3D; 1GB | [optional] [default to 256]
**DiskEncrypted** | Pointer to **bool** | indicates if the database disk is encrypted or not | [optional] 

## Methods

### NewDatabaseAllOf

`func NewDatabaseAllOf() *DatabaseAllOf`

NewDatabaseAllOf instantiates a new DatabaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseAllOfWithDefaults

`func NewDatabaseAllOfWithDefaults() *DatabaseAllOf`

NewDatabaseAllOfWithDefaults instantiates a new DatabaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *DatabaseAllOf) GetEnvironment() ReferenceObject`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DatabaseAllOf) GetEnvironmentOk() (*ReferenceObject, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DatabaseAllOf) SetEnvironment(v ReferenceObject)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *DatabaseAllOf) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetHost

`func (o *DatabaseAllOf) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DatabaseAllOf) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DatabaseAllOf) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DatabaseAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *DatabaseAllOf) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DatabaseAllOf) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DatabaseAllOf) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DatabaseAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetMaximumCpu

`func (o *DatabaseAllOf) GetMaximumCpu() int32`

GetMaximumCpu returns the MaximumCpu field if non-nil, zero value otherwise.

### GetMaximumCpuOk

`func (o *DatabaseAllOf) GetMaximumCpuOk() (*int32, bool)`

GetMaximumCpuOk returns a tuple with the MaximumCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumCpu

`func (o *DatabaseAllOf) SetMaximumCpu(v int32)`

SetMaximumCpu sets MaximumCpu field to given value.

### HasMaximumCpu

`func (o *DatabaseAllOf) HasMaximumCpu() bool`

HasMaximumCpu returns a boolean if a field has been set.

### GetMaximumMemory

`func (o *DatabaseAllOf) GetMaximumMemory() int32`

GetMaximumMemory returns the MaximumMemory field if non-nil, zero value otherwise.

### GetMaximumMemoryOk

`func (o *DatabaseAllOf) GetMaximumMemoryOk() (*int32, bool)`

GetMaximumMemoryOk returns a tuple with the MaximumMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMemory

`func (o *DatabaseAllOf) SetMaximumMemory(v int32)`

SetMaximumMemory sets MaximumMemory field to given value.

### HasMaximumMemory

`func (o *DatabaseAllOf) HasMaximumMemory() bool`

HasMaximumMemory returns a boolean if a field has been set.

### GetDiskEncrypted

`func (o *DatabaseAllOf) GetDiskEncrypted() bool`

GetDiskEncrypted returns the DiskEncrypted field if non-nil, zero value otherwise.

### GetDiskEncryptedOk

`func (o *DatabaseAllOf) GetDiskEncryptedOk() (*bool, bool)`

GetDiskEncryptedOk returns a tuple with the DiskEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncrypted

`func (o *DatabaseAllOf) SetDiskEncrypted(v bool)`

SetDiskEncrypted sets DiskEncrypted field to given value.

### HasDiskEncrypted

`func (o *DatabaseAllOf) HasDiskEncrypted() bool`

HasDiskEncrypted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


