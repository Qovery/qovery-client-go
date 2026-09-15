# PlatformConfigurationSectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceComponentKey** | **string** | Component key to use for resolve requests and stored configuration. | 
**FieldKeys** | **[]string** | Unique top-level source fields displayed here instead of in the source component editor. | 

## Methods

### NewPlatformConfigurationSectionResponse

`func NewPlatformConfigurationSectionResponse(sourceComponentKey string, fieldKeys []string, ) *PlatformConfigurationSectionResponse`

NewPlatformConfigurationSectionResponse instantiates a new PlatformConfigurationSectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformConfigurationSectionResponseWithDefaults

`func NewPlatformConfigurationSectionResponseWithDefaults() *PlatformConfigurationSectionResponse`

NewPlatformConfigurationSectionResponseWithDefaults instantiates a new PlatformConfigurationSectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceComponentKey

`func (o *PlatformConfigurationSectionResponse) GetSourceComponentKey() string`

GetSourceComponentKey returns the SourceComponentKey field if non-nil, zero value otherwise.

### GetSourceComponentKeyOk

`func (o *PlatformConfigurationSectionResponse) GetSourceComponentKeyOk() (*string, bool)`

GetSourceComponentKeyOk returns a tuple with the SourceComponentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceComponentKey

`func (o *PlatformConfigurationSectionResponse) SetSourceComponentKey(v string)`

SetSourceComponentKey sets SourceComponentKey field to given value.


### GetFieldKeys

`func (o *PlatformConfigurationSectionResponse) GetFieldKeys() []string`

GetFieldKeys returns the FieldKeys field if non-nil, zero value otherwise.

### GetFieldKeysOk

`func (o *PlatformConfigurationSectionResponse) GetFieldKeysOk() (*[]string, bool)`

GetFieldKeysOk returns a tuple with the FieldKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldKeys

`func (o *PlatformConfigurationSectionResponse) SetFieldKeys(v []string)`

SetFieldKeys sets FieldKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


