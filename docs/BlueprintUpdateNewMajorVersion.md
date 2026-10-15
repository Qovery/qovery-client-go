# BlueprintUpdateNewMajorVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceVersion** | **string** | The engine major version (e.g. \&quot;17\&quot;). | 
**LatestTag** | **string** | Latest catalog tag for that major version. | 
**Mode** | [**BlueprintMajorSwitchMode**](BlueprintMajorSwitchMode.md) |  | 
**UnsupportedReason** | Pointer to **NullableString** | Why the move is refused. Set only when &#x60;mode&#x60; is &#x60;UNSUPPORTED&#x60;. | [optional] 

## Methods

### NewBlueprintUpdateNewMajorVersion

`func NewBlueprintUpdateNewMajorVersion(serviceVersion string, latestTag string, mode BlueprintMajorSwitchMode, ) *BlueprintUpdateNewMajorVersion`

NewBlueprintUpdateNewMajorVersion instantiates a new BlueprintUpdateNewMajorVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintUpdateNewMajorVersionWithDefaults

`func NewBlueprintUpdateNewMajorVersionWithDefaults() *BlueprintUpdateNewMajorVersion`

NewBlueprintUpdateNewMajorVersionWithDefaults instantiates a new BlueprintUpdateNewMajorVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceVersion

`func (o *BlueprintUpdateNewMajorVersion) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *BlueprintUpdateNewMajorVersion) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *BlueprintUpdateNewMajorVersion) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.


### GetLatestTag

`func (o *BlueprintUpdateNewMajorVersion) GetLatestTag() string`

GetLatestTag returns the LatestTag field if non-nil, zero value otherwise.

### GetLatestTagOk

`func (o *BlueprintUpdateNewMajorVersion) GetLatestTagOk() (*string, bool)`

GetLatestTagOk returns a tuple with the LatestTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTag

`func (o *BlueprintUpdateNewMajorVersion) SetLatestTag(v string)`

SetLatestTag sets LatestTag field to given value.


### GetMode

`func (o *BlueprintUpdateNewMajorVersion) GetMode() BlueprintMajorSwitchMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *BlueprintUpdateNewMajorVersion) GetModeOk() (*BlueprintMajorSwitchMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *BlueprintUpdateNewMajorVersion) SetMode(v BlueprintMajorSwitchMode)`

SetMode sets Mode field to given value.


### GetUnsupportedReason

`func (o *BlueprintUpdateNewMajorVersion) GetUnsupportedReason() string`

GetUnsupportedReason returns the UnsupportedReason field if non-nil, zero value otherwise.

### GetUnsupportedReasonOk

`func (o *BlueprintUpdateNewMajorVersion) GetUnsupportedReasonOk() (*string, bool)`

GetUnsupportedReasonOk returns a tuple with the UnsupportedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsupportedReason

`func (o *BlueprintUpdateNewMajorVersion) SetUnsupportedReason(v string)`

SetUnsupportedReason sets UnsupportedReason field to given value.

### HasUnsupportedReason

`func (o *BlueprintUpdateNewMajorVersion) HasUnsupportedReason() bool`

HasUnsupportedReason returns a boolean if a field has been set.

### SetUnsupportedReasonNil

`func (o *BlueprintUpdateNewMajorVersion) SetUnsupportedReasonNil(b bool)`

 SetUnsupportedReasonNil sets the value for UnsupportedReason to be an explicit nil

### UnsetUnsupportedReason
`func (o *BlueprintUpdateNewMajorVersion) UnsetUnsupportedReason()`

UnsetUnsupportedReason ensures that no value is present for UnsupportedReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


