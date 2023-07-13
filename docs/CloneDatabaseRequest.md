# CloneDatabaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**EnvironmentId** | **string** |  | 

## Methods

### NewCloneDatabaseRequest

`func NewCloneDatabaseRequest(name string, environmentId string, ) *CloneDatabaseRequest`

NewCloneDatabaseRequest instantiates a new CloneDatabaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneDatabaseRequestWithDefaults

`func NewCloneDatabaseRequestWithDefaults() *CloneDatabaseRequest`

NewCloneDatabaseRequestWithDefaults instantiates a new CloneDatabaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloneDatabaseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloneDatabaseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloneDatabaseRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEnvironmentId

`func (o *CloneDatabaseRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *CloneDatabaseRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *CloneDatabaseRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


