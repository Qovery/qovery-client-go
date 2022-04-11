# LogicalDatabase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **string** |  | [optional] 
**Database** | Pointer to [**ReferenceObject**](ReferenceObject.md) |  | [optional] 

## Methods

### NewLogicalDatabase

`func NewLogicalDatabase(id string, createdAt time.Time, name string, ) *LogicalDatabase`

NewLogicalDatabase instantiates a new LogicalDatabase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogicalDatabaseWithDefaults

`func NewLogicalDatabaseWithDefaults() *LogicalDatabase`

NewLogicalDatabaseWithDefaults instantiates a new LogicalDatabase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogicalDatabase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogicalDatabase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogicalDatabase) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *LogicalDatabase) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LogicalDatabase) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LogicalDatabase) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *LogicalDatabase) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LogicalDatabase) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LogicalDatabase) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LogicalDatabase) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *LogicalDatabase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogicalDatabase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogicalDatabase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LogicalDatabase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LogicalDatabase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LogicalDatabase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LogicalDatabase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDatabase

`func (o *LogicalDatabase) GetDatabase() ReferenceObject`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *LogicalDatabase) GetDatabaseOk() (*ReferenceObject, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *LogicalDatabase) SetDatabase(v ReferenceObject)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *LogicalDatabase) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


