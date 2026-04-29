# ArgoCdConnectionCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ArgoCdConnectionStatusEnum**](ArgoCdConnectionStatusEnum.md) |  | 
**AppCount** | Pointer to **int32** | Number of ArgoCD applications visible with the provided token. Present only when status is \&quot;connected\&quot;. | [optional] 
**Reason** | Pointer to **string** | Failure reason. Present only when status is \&quot;error\&quot;. | [optional] 

## Methods

### NewArgoCdConnectionCheckResponse

`func NewArgoCdConnectionCheckResponse(status ArgoCdConnectionStatusEnum, ) *ArgoCdConnectionCheckResponse`

NewArgoCdConnectionCheckResponse instantiates a new ArgoCdConnectionCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoCdConnectionCheckResponseWithDefaults

`func NewArgoCdConnectionCheckResponseWithDefaults() *ArgoCdConnectionCheckResponse`

NewArgoCdConnectionCheckResponseWithDefaults instantiates a new ArgoCdConnectionCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ArgoCdConnectionCheckResponse) GetStatus() ArgoCdConnectionStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArgoCdConnectionCheckResponse) GetStatusOk() (*ArgoCdConnectionStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArgoCdConnectionCheckResponse) SetStatus(v ArgoCdConnectionStatusEnum)`

SetStatus sets Status field to given value.


### GetAppCount

`func (o *ArgoCdConnectionCheckResponse) GetAppCount() int32`

GetAppCount returns the AppCount field if non-nil, zero value otherwise.

### GetAppCountOk

`func (o *ArgoCdConnectionCheckResponse) GetAppCountOk() (*int32, bool)`

GetAppCountOk returns a tuple with the AppCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCount

`func (o *ArgoCdConnectionCheckResponse) SetAppCount(v int32)`

SetAppCount sets AppCount field to given value.

### HasAppCount

`func (o *ArgoCdConnectionCheckResponse) HasAppCount() bool`

HasAppCount returns a boolean if a field has been set.

### GetReason

`func (o *ArgoCdConnectionCheckResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ArgoCdConnectionCheckResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ArgoCdConnectionCheckResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ArgoCdConnectionCheckResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


