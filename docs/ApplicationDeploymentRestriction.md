# ApplicationDeploymentRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | Match mode will rebuild app only if specified items are updated. Exclude mode will not rebuild app if specified items are updated. | 
**Type** | **string** |  | 
**Value** | **string** |  | 

## Methods

### NewApplicationDeploymentRestriction

`func NewApplicationDeploymentRestriction(mode string, type_ string, value string, ) *ApplicationDeploymentRestriction`

NewApplicationDeploymentRestriction instantiates a new ApplicationDeploymentRestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDeploymentRestrictionWithDefaults

`func NewApplicationDeploymentRestrictionWithDefaults() *ApplicationDeploymentRestriction`

NewApplicationDeploymentRestrictionWithDefaults instantiates a new ApplicationDeploymentRestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *ApplicationDeploymentRestriction) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ApplicationDeploymentRestriction) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ApplicationDeploymentRestriction) SetMode(v string)`

SetMode sets Mode field to given value.


### GetType

`func (o *ApplicationDeploymentRestriction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationDeploymentRestriction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationDeploymentRestriction) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ApplicationDeploymentRestriction) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApplicationDeploymentRestriction) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApplicationDeploymentRestriction) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


