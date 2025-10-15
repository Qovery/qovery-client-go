# StatusDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**ServiceActionEnum**](ServiceActionEnum.md) |  | 
**Status** | [**ServiceActionStatusEnum**](ServiceActionStatusEnum.md) |  | 
**SubAction** | [**ServiceSubActionEnum**](ServiceSubActionEnum.md) |  | [default to SERVICESUBACTIONENUM_NONE]

## Methods

### NewStatusDetails

`func NewStatusDetails(action ServiceActionEnum, status ServiceActionStatusEnum, subAction ServiceSubActionEnum, ) *StatusDetails`

NewStatusDetails instantiates a new StatusDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusDetailsWithDefaults

`func NewStatusDetailsWithDefaults() *StatusDetails`

NewStatusDetailsWithDefaults instantiates a new StatusDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *StatusDetails) GetAction() ServiceActionEnum`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *StatusDetails) GetActionOk() (*ServiceActionEnum, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *StatusDetails) SetAction(v ServiceActionEnum)`

SetAction sets Action field to given value.


### GetStatus

`func (o *StatusDetails) GetStatus() ServiceActionStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StatusDetails) GetStatusOk() (*ServiceActionStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StatusDetails) SetStatus(v ServiceActionStatusEnum)`

SetStatus sets Status field to given value.


### GetSubAction

`func (o *StatusDetails) GetSubAction() ServiceSubActionEnum`

GetSubAction returns the SubAction field if non-nil, zero value otherwise.

### GetSubActionOk

`func (o *StatusDetails) GetSubActionOk() (*ServiceSubActionEnum, bool)`

GetSubActionOk returns a tuple with the SubAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAction

`func (o *StatusDetails) SetSubAction(v ServiceSubActionEnum)`

SetSubAction sets SubAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


