# BlueprintItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Kind** | **string** |  | 
**Description** | **string** |  | 
**Icon** | **string** |  | 
**Categories** | **[]string** |  | 
**Provider** | **string** |  | 
**ServiceFamily** | **string** |  | 
**MajorVersions** | [**[]BlueprintMajorVersion**](BlueprintMajorVersion.md) |  | 

## Methods

### NewBlueprintItem

`func NewBlueprintItem(name string, kind string, description string, icon string, categories []string, provider string, serviceFamily string, majorVersions []BlueprintMajorVersion, ) *BlueprintItem`

NewBlueprintItem instantiates a new BlueprintItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintItemWithDefaults

`func NewBlueprintItemWithDefaults() *BlueprintItem`

NewBlueprintItemWithDefaults instantiates a new BlueprintItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BlueprintItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintItem) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *BlueprintItem) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BlueprintItem) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BlueprintItem) SetKind(v string)`

SetKind sets Kind field to given value.


### GetDescription

`func (o *BlueprintItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BlueprintItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BlueprintItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIcon

`func (o *BlueprintItem) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *BlueprintItem) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *BlueprintItem) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetCategories

`func (o *BlueprintItem) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *BlueprintItem) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *BlueprintItem) SetCategories(v []string)`

SetCategories sets Categories field to given value.


### GetProvider

`func (o *BlueprintItem) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *BlueprintItem) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *BlueprintItem) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetServiceFamily

`func (o *BlueprintItem) GetServiceFamily() string`

GetServiceFamily returns the ServiceFamily field if non-nil, zero value otherwise.

### GetServiceFamilyOk

`func (o *BlueprintItem) GetServiceFamilyOk() (*string, bool)`

GetServiceFamilyOk returns a tuple with the ServiceFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFamily

`func (o *BlueprintItem) SetServiceFamily(v string)`

SetServiceFamily sets ServiceFamily field to given value.


### GetMajorVersions

`func (o *BlueprintItem) GetMajorVersions() []BlueprintMajorVersion`

GetMajorVersions returns the MajorVersions field if non-nil, zero value otherwise.

### GetMajorVersionsOk

`func (o *BlueprintItem) GetMajorVersionsOk() (*[]BlueprintMajorVersion, bool)`

GetMajorVersionsOk returns a tuple with the MajorVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorVersions

`func (o *BlueprintItem) SetMajorVersions(v []BlueprintMajorVersion)`

SetMajorVersions sets MajorVersions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


