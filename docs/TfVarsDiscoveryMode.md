# TfVarsDiscoveryMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator for auto-discovery mode | 
**Paths** | **[]string** | List of specific paths to tfvars files | 

## Methods

### NewTfVarsDiscoveryMode

`func NewTfVarsDiscoveryMode(type_ string, paths []string, ) *TfVarsDiscoveryMode`

NewTfVarsDiscoveryMode instantiates a new TfVarsDiscoveryMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTfVarsDiscoveryModeWithDefaults

`func NewTfVarsDiscoveryModeWithDefaults() *TfVarsDiscoveryMode`

NewTfVarsDiscoveryModeWithDefaults instantiates a new TfVarsDiscoveryMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TfVarsDiscoveryMode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TfVarsDiscoveryMode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TfVarsDiscoveryMode) SetType(v string)`

SetType sets Type field to given value.


### GetPaths

`func (o *TfVarsDiscoveryMode) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *TfVarsDiscoveryMode) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *TfVarsDiscoveryMode) SetPaths(v []string)`

SetPaths sets Paths field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


