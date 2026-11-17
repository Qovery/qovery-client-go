# CustomDomainPortResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the port. | 
**Name** | **string** | Name of the port. | 
**InternalPort** | **int32** | The port the service listens on. | 

## Methods

### NewCustomDomainPortResponse

`func NewCustomDomainPortResponse(id string, name string, internalPort int32, ) *CustomDomainPortResponse`

NewCustomDomainPortResponse instantiates a new CustomDomainPortResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainPortResponseWithDefaults

`func NewCustomDomainPortResponseWithDefaults() *CustomDomainPortResponse`

NewCustomDomainPortResponseWithDefaults instantiates a new CustomDomainPortResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomDomainPortResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomDomainPortResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomDomainPortResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CustomDomainPortResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomDomainPortResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomDomainPortResponse) SetName(v string)`

SetName sets Name field to given value.


### GetInternalPort

`func (o *CustomDomainPortResponse) GetInternalPort() int32`

GetInternalPort returns the InternalPort field if non-nil, zero value otherwise.

### GetInternalPortOk

`func (o *CustomDomainPortResponse) GetInternalPortOk() (*int32, bool)`

GetInternalPortOk returns a tuple with the InternalPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalPort

`func (o *CustomDomainPortResponse) SetInternalPort(v int32)`

SetInternalPort sets InternalPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


