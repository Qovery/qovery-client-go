# LifecycleTemplateResponseVariablesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** | Short description to explain the purpose of the variable | 
**Default** | **string** | Default value for the variable | 
**IsSecret** | **string** | If the variable should be injected as a secret | 
**File** | Pointer to [**LifecycleTemplateResponseVariablesInnerFile**](LifecycleTemplateResponseVariablesInnerFile.md) |  | [optional] 

## Methods

### NewLifecycleTemplateResponseVariablesInner

`func NewLifecycleTemplateResponseVariablesInner(name string, description string, default_ string, isSecret string, ) *LifecycleTemplateResponseVariablesInner`

NewLifecycleTemplateResponseVariablesInner instantiates a new LifecycleTemplateResponseVariablesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLifecycleTemplateResponseVariablesInnerWithDefaults

`func NewLifecycleTemplateResponseVariablesInnerWithDefaults() *LifecycleTemplateResponseVariablesInner`

NewLifecycleTemplateResponseVariablesInnerWithDefaults instantiates a new LifecycleTemplateResponseVariablesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LifecycleTemplateResponseVariablesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LifecycleTemplateResponseVariablesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LifecycleTemplateResponseVariablesInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LifecycleTemplateResponseVariablesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LifecycleTemplateResponseVariablesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LifecycleTemplateResponseVariablesInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDefault

`func (o *LifecycleTemplateResponseVariablesInner) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *LifecycleTemplateResponseVariablesInner) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *LifecycleTemplateResponseVariablesInner) SetDefault(v string)`

SetDefault sets Default field to given value.


### GetIsSecret

`func (o *LifecycleTemplateResponseVariablesInner) GetIsSecret() string`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *LifecycleTemplateResponseVariablesInner) GetIsSecretOk() (*string, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *LifecycleTemplateResponseVariablesInner) SetIsSecret(v string)`

SetIsSecret sets IsSecret field to given value.


### GetFile

`func (o *LifecycleTemplateResponseVariablesInner) GetFile() LifecycleTemplateResponseVariablesInnerFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *LifecycleTemplateResponseVariablesInner) GetFileOk() (*LifecycleTemplateResponseVariablesInnerFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *LifecycleTemplateResponseVariablesInner) SetFile(v LifecycleTemplateResponseVariablesInnerFile)`

SetFile sets File field to given value.

### HasFile

`func (o *LifecycleTemplateResponseVariablesInner) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


