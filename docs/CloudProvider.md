# CloudProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShortName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**Regions** | Pointer to [**[]ClusterRegion**](ClusterRegion.md) |  | [optional] 

## Methods

### NewCloudProvider

`func NewCloudProvider() *CloudProvider`

NewCloudProvider instantiates a new CloudProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderWithDefaults

`func NewCloudProviderWithDefaults() *CloudProvider`

NewCloudProviderWithDefaults instantiates a new CloudProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortName

`func (o *CloudProvider) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *CloudProvider) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *CloudProvider) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *CloudProvider) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetName

`func (o *CloudProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogoUrl

`func (o *CloudProvider) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *CloudProvider) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *CloudProvider) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *CloudProvider) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetRegions

`func (o *CloudProvider) GetRegions() []ClusterRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *CloudProvider) GetRegionsOk() (*[]ClusterRegion, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *CloudProvider) SetRegions(v []ClusterRegion)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *CloudProvider) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


