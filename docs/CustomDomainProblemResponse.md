# CustomDomainProblemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**Detail** | Pointer to **string** | All the error messages, joined. | [optional] 
**Instance** | Pointer to **string** | The path of the request. | [optional] 
**Errors** | Pointer to [**[]CustomDomainErrorResponse**](CustomDomainErrorResponse.md) |  | [optional] 

## Methods

### NewCustomDomainProblemResponse

`func NewCustomDomainProblemResponse() *CustomDomainProblemResponse`

NewCustomDomainProblemResponse instantiates a new CustomDomainProblemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainProblemResponseWithDefaults

`func NewCustomDomainProblemResponseWithDefaults() *CustomDomainProblemResponse`

NewCustomDomainProblemResponseWithDefaults instantiates a new CustomDomainProblemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomDomainProblemResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomDomainProblemResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomDomainProblemResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomDomainProblemResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *CustomDomainProblemResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomDomainProblemResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomDomainProblemResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CustomDomainProblemResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *CustomDomainProblemResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomDomainProblemResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomDomainProblemResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomDomainProblemResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetail

`func (o *CustomDomainProblemResponse) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *CustomDomainProblemResponse) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *CustomDomainProblemResponse) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *CustomDomainProblemResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetInstance

`func (o *CustomDomainProblemResponse) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *CustomDomainProblemResponse) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *CustomDomainProblemResponse) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *CustomDomainProblemResponse) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetErrors

`func (o *CustomDomainProblemResponse) GetErrors() []CustomDomainErrorResponse`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CustomDomainProblemResponse) GetErrorsOk() (*[]CustomDomainErrorResponse, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CustomDomainProblemResponse) SetErrors(v []CustomDomainErrorResponse)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CustomDomainProblemResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


