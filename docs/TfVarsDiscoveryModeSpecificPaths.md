# TfVarsDiscoveryModeSpecificPaths

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Discriminator for specific paths mode | 
**Paths** | **[]string** | List of specific paths to tfvars files | 

## Methods

### NewTfVarsDiscoveryModeSpecificPaths

`func NewTfVarsDiscoveryModeSpecificPaths(type_ string, paths []string, ) *TfVarsDiscoveryModeSpecificPaths`

NewTfVarsDiscoveryModeSpecificPaths instantiates a new TfVarsDiscoveryModeSpecificPaths object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTfVarsDiscoveryModeSpecificPathsWithDefaults

`func NewTfVarsDiscoveryModeSpecificPathsWithDefaults() *TfVarsDiscoveryModeSpecificPaths`

NewTfVarsDiscoveryModeSpecificPathsWithDefaults instantiates a new TfVarsDiscoveryModeSpecificPaths object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TfVarsDiscoveryModeSpecificPaths) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TfVarsDiscoveryModeSpecificPaths) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TfVarsDiscoveryModeSpecificPaths) SetType(v string)`

SetType sets Type field to given value.


### GetPaths

`func (o *TfVarsDiscoveryModeSpecificPaths) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *TfVarsDiscoveryModeSpecificPaths) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *TfVarsDiscoveryModeSpecificPaths) SetPaths(v []string)`

SetPaths sets Paths field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


