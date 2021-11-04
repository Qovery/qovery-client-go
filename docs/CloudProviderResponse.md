# CloudProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShortName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**Regions** | Pointer to [**[]ClusterRegionResponse**](ClusterRegionResponse.md) |  | [optional] 

## Methods

### NewCloudProviderResponse

`func NewCloudProviderResponse() *CloudProviderResponse`

NewCloudProviderResponse instantiates a new CloudProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderResponseWithDefaults

`func NewCloudProviderResponseWithDefaults() *CloudProviderResponse`

NewCloudProviderResponseWithDefaults instantiates a new CloudProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortName

`func (o *CloudProviderResponse) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *CloudProviderResponse) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *CloudProviderResponse) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *CloudProviderResponse) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetName

`func (o *CloudProviderResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudProviderResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudProviderResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudProviderResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogoUrl

`func (o *CloudProviderResponse) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *CloudProviderResponse) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *CloudProviderResponse) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *CloudProviderResponse) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetRegions

`func (o *CloudProviderResponse) GetRegions() []ClusterRegionResponse`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *CloudProviderResponse) GetRegionsOk() (*[]ClusterRegionResponse, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *CloudProviderResponse) SetRegions(v []ClusterRegionResponse)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *CloudProviderResponse) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


