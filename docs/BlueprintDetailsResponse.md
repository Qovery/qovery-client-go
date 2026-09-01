# BlueprintDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**CatalogUrl** | **string** | URL to the blueprint catalog entry | 
**Tag** | **string** |  | 
**EnvironmentId** | **string** |  | 
**ServiceType** | **string** | Type of the underlying service backing this blueprint | 
**ServiceId** | **NullableString** | The service the dispatch produced. Null while the dispatch is still running, and null if it failed. | 
**LatestDeployment** | [**NullableBlueprintDetailsResponseLatestDeployment**](BlueprintDetailsResponseLatestDeployment.md) |  | 

## Methods

### NewBlueprintDetailsResponse

`func NewBlueprintDetailsResponse(id string, name string, catalogUrl string, tag string, environmentId string, serviceType string, serviceId NullableString, latestDeployment NullableBlueprintDetailsResponseLatestDeployment, ) *BlueprintDetailsResponse`

NewBlueprintDetailsResponse instantiates a new BlueprintDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintDetailsResponseWithDefaults

`func NewBlueprintDetailsResponseWithDefaults() *BlueprintDetailsResponse`

NewBlueprintDetailsResponseWithDefaults instantiates a new BlueprintDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlueprintDetailsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlueprintDetailsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlueprintDetailsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BlueprintDetailsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintDetailsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintDetailsResponse) SetName(v string)`

SetName sets Name field to given value.


### GetCatalogUrl

`func (o *BlueprintDetailsResponse) GetCatalogUrl() string`

GetCatalogUrl returns the CatalogUrl field if non-nil, zero value otherwise.

### GetCatalogUrlOk

`func (o *BlueprintDetailsResponse) GetCatalogUrlOk() (*string, bool)`

GetCatalogUrlOk returns a tuple with the CatalogUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogUrl

`func (o *BlueprintDetailsResponse) SetCatalogUrl(v string)`

SetCatalogUrl sets CatalogUrl field to given value.


### GetTag

`func (o *BlueprintDetailsResponse) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *BlueprintDetailsResponse) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *BlueprintDetailsResponse) SetTag(v string)`

SetTag sets Tag field to given value.


### GetEnvironmentId

`func (o *BlueprintDetailsResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *BlueprintDetailsResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *BlueprintDetailsResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetServiceType

`func (o *BlueprintDetailsResponse) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *BlueprintDetailsResponse) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *BlueprintDetailsResponse) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.


### GetServiceId

`func (o *BlueprintDetailsResponse) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *BlueprintDetailsResponse) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *BlueprintDetailsResponse) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### SetServiceIdNil

`func (o *BlueprintDetailsResponse) SetServiceIdNil(b bool)`

 SetServiceIdNil sets the value for ServiceId to be an explicit nil

### UnsetServiceId
`func (o *BlueprintDetailsResponse) UnsetServiceId()`

UnsetServiceId ensures that no value is present for ServiceId, not even an explicit nil
### GetLatestDeployment

`func (o *BlueprintDetailsResponse) GetLatestDeployment() BlueprintDetailsResponseLatestDeployment`

GetLatestDeployment returns the LatestDeployment field if non-nil, zero value otherwise.

### GetLatestDeploymentOk

`func (o *BlueprintDetailsResponse) GetLatestDeploymentOk() (*BlueprintDetailsResponseLatestDeployment, bool)`

GetLatestDeploymentOk returns a tuple with the LatestDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDeployment

`func (o *BlueprintDetailsResponse) SetLatestDeployment(v BlueprintDetailsResponseLatestDeployment)`

SetLatestDeployment sets LatestDeployment field to given value.


### SetLatestDeploymentNil

`func (o *BlueprintDetailsResponse) SetLatestDeploymentNil(b bool)`

 SetLatestDeploymentNil sets the value for LatestDeployment to be an explicit nil

### UnsetLatestDeployment
`func (o *BlueprintDetailsResponse) UnsetLatestDeployment()`

UnsetLatestDeployment ensures that no value is present for LatestDeployment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


