# KedaTriggerAuthenticationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ConfigYaml** | Pointer to **string** | Optional raw KEDA TriggerAuthentication YAML configuration. | [optional] 

## Methods

### NewKedaTriggerAuthenticationRequest

`func NewKedaTriggerAuthenticationRequest(name string, ) *KedaTriggerAuthenticationRequest`

NewKedaTriggerAuthenticationRequest instantiates a new KedaTriggerAuthenticationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKedaTriggerAuthenticationRequestWithDefaults

`func NewKedaTriggerAuthenticationRequestWithDefaults() *KedaTriggerAuthenticationRequest`

NewKedaTriggerAuthenticationRequestWithDefaults instantiates a new KedaTriggerAuthenticationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KedaTriggerAuthenticationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KedaTriggerAuthenticationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KedaTriggerAuthenticationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetConfigYaml

`func (o *KedaTriggerAuthenticationRequest) GetConfigYaml() string`

GetConfigYaml returns the ConfigYaml field if non-nil, zero value otherwise.

### GetConfigYamlOk

`func (o *KedaTriggerAuthenticationRequest) GetConfigYamlOk() (*string, bool)`

GetConfigYamlOk returns a tuple with the ConfigYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigYaml

`func (o *KedaTriggerAuthenticationRequest) SetConfigYaml(v string)`

SetConfigYaml sets ConfigYaml field to given value.

### HasConfigYaml

`func (o *KedaTriggerAuthenticationRequest) HasConfigYaml() bool`

HasConfigYaml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


