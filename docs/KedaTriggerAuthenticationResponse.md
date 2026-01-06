# KedaTriggerAuthenticationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**OrganizationId** | **string** |  | 
**Name** | **string** |  | 
**ConfigYaml** | Pointer to **string** |  Optional raw KEDA TriggerAuthentication YAML configuration. | [optional] 

## Methods

### NewKedaTriggerAuthenticationResponse

`func NewKedaTriggerAuthenticationResponse(id string, createdAt time.Time, organizationId string, name string, ) *KedaTriggerAuthenticationResponse`

NewKedaTriggerAuthenticationResponse instantiates a new KedaTriggerAuthenticationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKedaTriggerAuthenticationResponseWithDefaults

`func NewKedaTriggerAuthenticationResponseWithDefaults() *KedaTriggerAuthenticationResponse`

NewKedaTriggerAuthenticationResponseWithDefaults instantiates a new KedaTriggerAuthenticationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KedaTriggerAuthenticationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KedaTriggerAuthenticationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KedaTriggerAuthenticationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *KedaTriggerAuthenticationResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KedaTriggerAuthenticationResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KedaTriggerAuthenticationResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *KedaTriggerAuthenticationResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KedaTriggerAuthenticationResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KedaTriggerAuthenticationResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KedaTriggerAuthenticationResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *KedaTriggerAuthenticationResponse) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *KedaTriggerAuthenticationResponse) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *KedaTriggerAuthenticationResponse) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetName

`func (o *KedaTriggerAuthenticationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KedaTriggerAuthenticationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KedaTriggerAuthenticationResponse) SetName(v string)`

SetName sets Name field to given value.


### GetConfigYaml

`func (o *KedaTriggerAuthenticationResponse) GetConfigYaml() string`

GetConfigYaml returns the ConfigYaml field if non-nil, zero value otherwise.

### GetConfigYamlOk

`func (o *KedaTriggerAuthenticationResponse) GetConfigYamlOk() (*string, bool)`

GetConfigYamlOk returns a tuple with the ConfigYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigYaml

`func (o *KedaTriggerAuthenticationResponse) SetConfigYaml(v string)`

SetConfigYaml sets ConfigYaml field to given value.

### HasConfigYaml

`func (o *KedaTriggerAuthenticationResponse) HasConfigYaml() bool`

HasConfigYaml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


