# ClusterRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**CountryCode** | **string** |  | 
**Country** | **string** |  | 
**City** | **string** |  | 
**Zones** | Pointer to **[]string** | List of availability zones supported by this region | [optional] 

## Methods

### NewClusterRegion

`func NewClusterRegion(name string, countryCode string, country string, city string, ) *ClusterRegion`

NewClusterRegion instantiates a new ClusterRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRegionWithDefaults

`func NewClusterRegionWithDefaults() *ClusterRegion`

NewClusterRegionWithDefaults instantiates a new ClusterRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterRegion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterRegion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterRegion) SetName(v string)`

SetName sets Name field to given value.


### GetCountryCode

`func (o *ClusterRegion) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ClusterRegion) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ClusterRegion) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetCountry

`func (o *ClusterRegion) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ClusterRegion) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ClusterRegion) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCity

`func (o *ClusterRegion) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ClusterRegion) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ClusterRegion) SetCity(v string)`

SetCity sets City field to given value.


### GetZones

`func (o *ClusterRegion) GetZones() []string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *ClusterRegion) GetZonesOk() (*[]string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *ClusterRegion) SetZones(v []string)`

SetZones sets Zones field to given value.

### HasZones

`func (o *ClusterRegion) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


