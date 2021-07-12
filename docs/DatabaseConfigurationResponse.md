# DatabaseConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseType** | Pointer to **string** |  | [optional] 
**Version** | Pointer to [**[]DatabaseVersionMode**](DatabaseVersionMode.md) |  | [optional] 

## Methods

### NewDatabaseConfigurationResponse

`func NewDatabaseConfigurationResponse() *DatabaseConfigurationResponse`

NewDatabaseConfigurationResponse instantiates a new DatabaseConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseConfigurationResponseWithDefaults

`func NewDatabaseConfigurationResponseWithDefaults() *DatabaseConfigurationResponse`

NewDatabaseConfigurationResponseWithDefaults instantiates a new DatabaseConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseType

`func (o *DatabaseConfigurationResponse) GetDatabaseType() string`

GetDatabaseType returns the DatabaseType field if non-nil, zero value otherwise.

### GetDatabaseTypeOk

`func (o *DatabaseConfigurationResponse) GetDatabaseTypeOk() (*string, bool)`

GetDatabaseTypeOk returns a tuple with the DatabaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseType

`func (o *DatabaseConfigurationResponse) SetDatabaseType(v string)`

SetDatabaseType sets DatabaseType field to given value.

### HasDatabaseType

`func (o *DatabaseConfigurationResponse) HasDatabaseType() bool`

HasDatabaseType returns a boolean if a field has been set.

### GetVersion

`func (o *DatabaseConfigurationResponse) GetVersion() []DatabaseVersionMode`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DatabaseConfigurationResponse) GetVersionOk() (*[]DatabaseVersionMode, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DatabaseConfigurationResponse) SetVersion(v []DatabaseVersionMode)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DatabaseConfigurationResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


