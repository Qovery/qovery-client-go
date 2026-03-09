# ArgoCdCredentialsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ClusterId** | **string** |  | 
**ArgocdUrl** | **string** |  | 
**ArgocdToken** | **string** | Always returned as REDACTED | 
**LastCheckedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewArgoCdCredentialsResponse

`func NewArgoCdCredentialsResponse(id string, clusterId string, argocdUrl string, argocdToken string, createdAt time.Time, updatedAt time.Time, ) *ArgoCdCredentialsResponse`

NewArgoCdCredentialsResponse instantiates a new ArgoCdCredentialsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoCdCredentialsResponseWithDefaults

`func NewArgoCdCredentialsResponseWithDefaults() *ArgoCdCredentialsResponse`

NewArgoCdCredentialsResponseWithDefaults instantiates a new ArgoCdCredentialsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArgoCdCredentialsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArgoCdCredentialsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArgoCdCredentialsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetClusterId

`func (o *ArgoCdCredentialsResponse) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ArgoCdCredentialsResponse) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ArgoCdCredentialsResponse) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetArgocdUrl

`func (o *ArgoCdCredentialsResponse) GetArgocdUrl() string`

GetArgocdUrl returns the ArgocdUrl field if non-nil, zero value otherwise.

### GetArgocdUrlOk

`func (o *ArgoCdCredentialsResponse) GetArgocdUrlOk() (*string, bool)`

GetArgocdUrlOk returns a tuple with the ArgocdUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgocdUrl

`func (o *ArgoCdCredentialsResponse) SetArgocdUrl(v string)`

SetArgocdUrl sets ArgocdUrl field to given value.


### GetArgocdToken

`func (o *ArgoCdCredentialsResponse) GetArgocdToken() string`

GetArgocdToken returns the ArgocdToken field if non-nil, zero value otherwise.

### GetArgocdTokenOk

`func (o *ArgoCdCredentialsResponse) GetArgocdTokenOk() (*string, bool)`

GetArgocdTokenOk returns a tuple with the ArgocdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgocdToken

`func (o *ArgoCdCredentialsResponse) SetArgocdToken(v string)`

SetArgocdToken sets ArgocdToken field to given value.


### GetLastCheckedAt

`func (o *ArgoCdCredentialsResponse) GetLastCheckedAt() time.Time`

GetLastCheckedAt returns the LastCheckedAt field if non-nil, zero value otherwise.

### GetLastCheckedAtOk

`func (o *ArgoCdCredentialsResponse) GetLastCheckedAtOk() (*time.Time, bool)`

GetLastCheckedAtOk returns a tuple with the LastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedAt

`func (o *ArgoCdCredentialsResponse) SetLastCheckedAt(v time.Time)`

SetLastCheckedAt sets LastCheckedAt field to given value.

### HasLastCheckedAt

`func (o *ArgoCdCredentialsResponse) HasLastCheckedAt() bool`

HasLastCheckedAt returns a boolean if a field has been set.

### SetLastCheckedAtNil

`func (o *ArgoCdCredentialsResponse) SetLastCheckedAtNil(b bool)`

 SetLastCheckedAtNil sets the value for LastCheckedAt to be an explicit nil

### UnsetLastCheckedAt
`func (o *ArgoCdCredentialsResponse) UnsetLastCheckedAt()`

UnsetLastCheckedAt ensures that no value is present for LastCheckedAt, not even an explicit nil
### GetCreatedAt

`func (o *ArgoCdCredentialsResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArgoCdCredentialsResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArgoCdCredentialsResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ArgoCdCredentialsResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArgoCdCredentialsResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArgoCdCredentialsResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


