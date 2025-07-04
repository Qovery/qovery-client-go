/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.4
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// checks if the CloudProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudProvider{}

// CloudProvider struct for CloudProvider
type CloudProvider struct {
	ShortName            *string         `json:"short_name,omitempty"`
	Name                 *string         `json:"name,omitempty"`
	LogoUrl              *string         `json:"logo_url,omitempty"`
	Regions              []ClusterRegion `json:"regions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudProvider CloudProvider

// NewCloudProvider instantiates a new CloudProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProvider() *CloudProvider {
	this := CloudProvider{}
	return &this
}

// NewCloudProviderWithDefaults instantiates a new CloudProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderWithDefaults() *CloudProvider {
	this := CloudProvider{}
	return &this
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *CloudProvider) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProvider) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *CloudProvider) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *CloudProvider) SetShortName(v string) {
	o.ShortName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudProvider) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProvider) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudProvider) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudProvider) SetName(v string) {
	o.Name = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *CloudProvider) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProvider) GetLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *CloudProvider) HasLogoUrl() bool {
	if o != nil && !IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *CloudProvider) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *CloudProvider) GetRegions() []ClusterRegion {
	if o == nil || IsNil(o.Regions) {
		var ret []ClusterRegion
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProvider) GetRegionsOk() ([]ClusterRegion, bool) {
	if o == nil || IsNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *CloudProvider) HasRegions() bool {
	if o != nil && !IsNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []ClusterRegion and assigns it to the Regions field.
func (o *CloudProvider) SetRegions(v []ClusterRegion) {
	o.Regions = v
}

func (o CloudProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShortName) {
		toSerialize["short_name"] = o.ShortName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.LogoUrl) {
		toSerialize["logo_url"] = o.LogoUrl
	}
	if !IsNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudProvider) UnmarshalJSON(data []byte) (err error) {
	varCloudProvider := _CloudProvider{}

	err = json.Unmarshal(data, &varCloudProvider)

	if err != nil {
		return err
	}

	*o = CloudProvider(varCloudProvider)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "short_name")
		delete(additionalProperties, "name")
		delete(additionalProperties, "logo_url")
		delete(additionalProperties, "regions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudProvider struct {
	value *CloudProvider
	isSet bool
}

func (v NullableCloudProvider) Get() *CloudProvider {
	return v.value
}

func (v *NullableCloudProvider) Set(val *CloudProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProvider(val *CloudProvider) *NullableCloudProvider {
	return &NullableCloudProvider{value: val, isSet: true}
}

func (v NullableCloudProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
