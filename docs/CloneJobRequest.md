# CloneJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**EnvironmentId** | **string** |  | 

## Methods

### NewCloneJobRequest

`func NewCloneJobRequest(name string, environmentId string, ) *CloneJobRequest`

NewCloneJobRequest instantiates a new CloneJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneJobRequestWithDefaults

`func NewCloneJobRequestWithDefaults() *CloneJobRequest`

NewCloneJobRequestWithDefaults instantiates a new CloneJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloneJobRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloneJobRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloneJobRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEnvironmentId

`func (o *CloneJobRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *CloneJobRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *CloneJobRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


