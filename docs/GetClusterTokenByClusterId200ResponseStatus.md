# GetClusterTokenByClusterId200ResponseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**ExpirationTimestamp** | **time.Time** |  | 

## Methods

### NewGetClusterTokenByClusterId200ResponseStatus

`func NewGetClusterTokenByClusterId200ResponseStatus(token string, expirationTimestamp time.Time, ) *GetClusterTokenByClusterId200ResponseStatus`

NewGetClusterTokenByClusterId200ResponseStatus instantiates a new GetClusterTokenByClusterId200ResponseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterTokenByClusterId200ResponseStatusWithDefaults

`func NewGetClusterTokenByClusterId200ResponseStatusWithDefaults() *GetClusterTokenByClusterId200ResponseStatus`

NewGetClusterTokenByClusterId200ResponseStatusWithDefaults instantiates a new GetClusterTokenByClusterId200ResponseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *GetClusterTokenByClusterId200ResponseStatus) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetClusterTokenByClusterId200ResponseStatus) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetClusterTokenByClusterId200ResponseStatus) SetToken(v string)`

SetToken sets Token field to given value.


### GetExpirationTimestamp

`func (o *GetClusterTokenByClusterId200ResponseStatus) GetExpirationTimestamp() time.Time`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *GetClusterTokenByClusterId200ResponseStatus) GetExpirationTimestampOk() (*time.Time, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *GetClusterTokenByClusterId200ResponseStatus) SetExpirationTimestamp(v time.Time)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


