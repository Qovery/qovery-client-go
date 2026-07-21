# BlueprintUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsUpToDate** | **bool** |  | 
**CurrentTag** | **string** |  | 
**LatestTag** | **string** |  | 
**NewRequiredValues** | [**[]BlueprintUpdateNewRequiredValue**](BlueprintUpdateNewRequiredValue.md) | Variables added in the latest version that are required with no default | 
**NewOptionalValues** | [**[]BlueprintUpdateNewOptionalValue**](BlueprintUpdateNewOptionalValue.md) | Variables added in the latest version that have a default value | 
**NowRequiredValues** | [**[]BlueprintUpdateNewRequiredValue**](BlueprintUpdateNewRequiredValue.md) | Variables that were optional but are now required in the latest version | 
**UpdatedValues** | [**[]BlueprintUpdateUpdatedValue**](BlueprintUpdateUpdatedValue.md) | Variables whose default value changed between the current and latest versions | 
**RemovedValues** | [**[]BlueprintUpdateRemovedValue**](BlueprintUpdateRemovedValue.md) | Variables that no longer exist in the latest version | 
**EngineDiff** | [**BlueprintUpdateEngineDiff**](BlueprintUpdateEngineDiff.md) |  | 
**NewMajorVersions** | [**[]BlueprintUpdateNewMajorVersion**](BlueprintUpdateNewMajorVersion.md) | Major versions of the same blueprint newer than the service&#39;s current one (e.g. service on aws/postgres/16 while 17 and 18 exist). Empty when already on the latest major or when the current version is non-numeric. | 

## Methods

### NewBlueprintUpdateResponse

`func NewBlueprintUpdateResponse(isUpToDate bool, currentTag string, latestTag string, newRequiredValues []BlueprintUpdateNewRequiredValue, newOptionalValues []BlueprintUpdateNewOptionalValue, nowRequiredValues []BlueprintUpdateNewRequiredValue, updatedValues []BlueprintUpdateUpdatedValue, removedValues []BlueprintUpdateRemovedValue, engineDiff BlueprintUpdateEngineDiff, newMajorVersions []BlueprintUpdateNewMajorVersion, ) *BlueprintUpdateResponse`

NewBlueprintUpdateResponse instantiates a new BlueprintUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintUpdateResponseWithDefaults

`func NewBlueprintUpdateResponseWithDefaults() *BlueprintUpdateResponse`

NewBlueprintUpdateResponseWithDefaults instantiates a new BlueprintUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsUpToDate

`func (o *BlueprintUpdateResponse) GetIsUpToDate() bool`

GetIsUpToDate returns the IsUpToDate field if non-nil, zero value otherwise.

### GetIsUpToDateOk

`func (o *BlueprintUpdateResponse) GetIsUpToDateOk() (*bool, bool)`

GetIsUpToDateOk returns a tuple with the IsUpToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpToDate

`func (o *BlueprintUpdateResponse) SetIsUpToDate(v bool)`

SetIsUpToDate sets IsUpToDate field to given value.


### GetCurrentTag

`func (o *BlueprintUpdateResponse) GetCurrentTag() string`

GetCurrentTag returns the CurrentTag field if non-nil, zero value otherwise.

### GetCurrentTagOk

`func (o *BlueprintUpdateResponse) GetCurrentTagOk() (*string, bool)`

GetCurrentTagOk returns a tuple with the CurrentTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTag

`func (o *BlueprintUpdateResponse) SetCurrentTag(v string)`

SetCurrentTag sets CurrentTag field to given value.


### GetLatestTag

`func (o *BlueprintUpdateResponse) GetLatestTag() string`

GetLatestTag returns the LatestTag field if non-nil, zero value otherwise.

### GetLatestTagOk

`func (o *BlueprintUpdateResponse) GetLatestTagOk() (*string, bool)`

GetLatestTagOk returns a tuple with the LatestTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTag

`func (o *BlueprintUpdateResponse) SetLatestTag(v string)`

SetLatestTag sets LatestTag field to given value.


### GetNewRequiredValues

`func (o *BlueprintUpdateResponse) GetNewRequiredValues() []BlueprintUpdateNewRequiredValue`

GetNewRequiredValues returns the NewRequiredValues field if non-nil, zero value otherwise.

### GetNewRequiredValuesOk

`func (o *BlueprintUpdateResponse) GetNewRequiredValuesOk() (*[]BlueprintUpdateNewRequiredValue, bool)`

GetNewRequiredValuesOk returns a tuple with the NewRequiredValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRequiredValues

`func (o *BlueprintUpdateResponse) SetNewRequiredValues(v []BlueprintUpdateNewRequiredValue)`

SetNewRequiredValues sets NewRequiredValues field to given value.


### GetNewOptionalValues

`func (o *BlueprintUpdateResponse) GetNewOptionalValues() []BlueprintUpdateNewOptionalValue`

GetNewOptionalValues returns the NewOptionalValues field if non-nil, zero value otherwise.

### GetNewOptionalValuesOk

`func (o *BlueprintUpdateResponse) GetNewOptionalValuesOk() (*[]BlueprintUpdateNewOptionalValue, bool)`

GetNewOptionalValuesOk returns a tuple with the NewOptionalValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOptionalValues

`func (o *BlueprintUpdateResponse) SetNewOptionalValues(v []BlueprintUpdateNewOptionalValue)`

SetNewOptionalValues sets NewOptionalValues field to given value.


### GetNowRequiredValues

`func (o *BlueprintUpdateResponse) GetNowRequiredValues() []BlueprintUpdateNewRequiredValue`

GetNowRequiredValues returns the NowRequiredValues field if non-nil, zero value otherwise.

### GetNowRequiredValuesOk

`func (o *BlueprintUpdateResponse) GetNowRequiredValuesOk() (*[]BlueprintUpdateNewRequiredValue, bool)`

GetNowRequiredValuesOk returns a tuple with the NowRequiredValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowRequiredValues

`func (o *BlueprintUpdateResponse) SetNowRequiredValues(v []BlueprintUpdateNewRequiredValue)`

SetNowRequiredValues sets NowRequiredValues field to given value.


### GetUpdatedValues

`func (o *BlueprintUpdateResponse) GetUpdatedValues() []BlueprintUpdateUpdatedValue`

GetUpdatedValues returns the UpdatedValues field if non-nil, zero value otherwise.

### GetUpdatedValuesOk

`func (o *BlueprintUpdateResponse) GetUpdatedValuesOk() (*[]BlueprintUpdateUpdatedValue, bool)`

GetUpdatedValuesOk returns a tuple with the UpdatedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedValues

`func (o *BlueprintUpdateResponse) SetUpdatedValues(v []BlueprintUpdateUpdatedValue)`

SetUpdatedValues sets UpdatedValues field to given value.


### GetRemovedValues

`func (o *BlueprintUpdateResponse) GetRemovedValues() []BlueprintUpdateRemovedValue`

GetRemovedValues returns the RemovedValues field if non-nil, zero value otherwise.

### GetRemovedValuesOk

`func (o *BlueprintUpdateResponse) GetRemovedValuesOk() (*[]BlueprintUpdateRemovedValue, bool)`

GetRemovedValuesOk returns a tuple with the RemovedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedValues

`func (o *BlueprintUpdateResponse) SetRemovedValues(v []BlueprintUpdateRemovedValue)`

SetRemovedValues sets RemovedValues field to given value.


### GetEngineDiff

`func (o *BlueprintUpdateResponse) GetEngineDiff() BlueprintUpdateEngineDiff`

GetEngineDiff returns the EngineDiff field if non-nil, zero value otherwise.

### GetEngineDiffOk

`func (o *BlueprintUpdateResponse) GetEngineDiffOk() (*BlueprintUpdateEngineDiff, bool)`

GetEngineDiffOk returns a tuple with the EngineDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineDiff

`func (o *BlueprintUpdateResponse) SetEngineDiff(v BlueprintUpdateEngineDiff)`

SetEngineDiff sets EngineDiff field to given value.


### GetNewMajorVersions

`func (o *BlueprintUpdateResponse) GetNewMajorVersions() []BlueprintUpdateNewMajorVersion`

GetNewMajorVersions returns the NewMajorVersions field if non-nil, zero value otherwise.

### GetNewMajorVersionsOk

`func (o *BlueprintUpdateResponse) GetNewMajorVersionsOk() (*[]BlueprintUpdateNewMajorVersion, bool)`

GetNewMajorVersionsOk returns a tuple with the NewMajorVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMajorVersions

`func (o *BlueprintUpdateResponse) SetNewMajorVersions(v []BlueprintUpdateNewMajorVersion)`

SetNewMajorVersions sets NewMajorVersions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


