/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// checks if the ClusterRegion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterRegion{}

// ClusterRegion struct for ClusterRegion
type ClusterRegion struct {
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
	Country     string `json:"country"`
	City        string `json:"city"`
}

// NewClusterRegion instantiates a new ClusterRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterRegion(name string, countryCode string, country string, city string) *ClusterRegion {
	this := ClusterRegion{}
	this.Name = name
	this.CountryCode = countryCode
	this.Country = country
	this.City = city
	return &this
}

// NewClusterRegionWithDefaults instantiates a new ClusterRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterRegionWithDefaults() *ClusterRegion {
	this := ClusterRegion{}
	return &this
}

// GetName returns the Name field value
func (o *ClusterRegion) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterRegion) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterRegion) SetName(v string) {
	o.Name = v
}

// GetCountryCode returns the CountryCode field value
func (o *ClusterRegion) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *ClusterRegion) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *ClusterRegion) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetCountry returns the Country field value
func (o *ClusterRegion) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *ClusterRegion) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *ClusterRegion) SetCountry(v string) {
	o.Country = v
}

// GetCity returns the City field value
func (o *ClusterRegion) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *ClusterRegion) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *ClusterRegion) SetCity(v string) {
	o.City = v
}

func (o ClusterRegion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["country_code"] = o.CountryCode
	toSerialize["country"] = o.Country
	toSerialize["city"] = o.City
	return toSerialize, nil
}

type NullableClusterRegion struct {
	value *ClusterRegion
	isSet bool
}

func (v NullableClusterRegion) Get() *ClusterRegion {
	return v.value
}

func (v *NullableClusterRegion) Set(val *ClusterRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterRegion(val *ClusterRegion) *NullableClusterRegion {
	return &NullableClusterRegion{value: val, isSet: true}
}

func (v NullableClusterRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
