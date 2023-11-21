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

// checks if the HelmResponseAllOfSourceOneOf1RepositoryRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmResponseAllOfSourceOneOf1RepositoryRepository{}

// HelmResponseAllOfSourceOneOf1RepositoryRepository struct for HelmResponseAllOfSourceOneOf1RepositoryRepository
type HelmResponseAllOfSourceOneOf1RepositoryRepository struct {
	// The id of the helm repository
	Id *string `json:"id,omitempty"`
	// The name of the helm repository
	Name *string `json:"name,omitempty"`
	// The url the helm repository
	Url *string `json:"url,omitempty"`
}

// NewHelmResponseAllOfSourceOneOf1RepositoryRepository instantiates a new HelmResponseAllOfSourceOneOf1RepositoryRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmResponseAllOfSourceOneOf1RepositoryRepository() *HelmResponseAllOfSourceOneOf1RepositoryRepository {
	this := HelmResponseAllOfSourceOneOf1RepositoryRepository{}
	return &this
}

// NewHelmResponseAllOfSourceOneOf1RepositoryRepositoryWithDefaults instantiates a new HelmResponseAllOfSourceOneOf1RepositoryRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmResponseAllOfSourceOneOf1RepositoryRepositoryWithDefaults() *HelmResponseAllOfSourceOneOf1RepositoryRepository {
	this := HelmResponseAllOfSourceOneOf1RepositoryRepository{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HelmResponseAllOfSourceOneOf1RepositoryRepository) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmResponseAllOfSourceOneOf1RepositoryRepository) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HelmResponseAllOfSourceOneOf1RepositoryRepository) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HelmResponseAllOfSourceOneOf1RepositoryRepository) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HelmResponseAllOfSourceOneOf1RepositoryRepository) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmResponseAllOfSourceOneOf1RepositoryRepository) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HelmResponseAllOfSourceOneOf1RepositoryRepository) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HelmResponseAllOfSourceOneOf1RepositoryRepository) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *HelmResponseAllOfSourceOneOf1RepositoryRepository) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmResponseAllOfSourceOneOf1RepositoryRepository) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *HelmResponseAllOfSourceOneOf1RepositoryRepository) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *HelmResponseAllOfSourceOneOf1RepositoryRepository) SetUrl(v string) {
	o.Url = &v
}

func (o HelmResponseAllOfSourceOneOf1RepositoryRepository) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmResponseAllOfSourceOneOf1RepositoryRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableHelmResponseAllOfSourceOneOf1RepositoryRepository struct {
	value *HelmResponseAllOfSourceOneOf1RepositoryRepository
	isSet bool
}

func (v NullableHelmResponseAllOfSourceOneOf1RepositoryRepository) Get() *HelmResponseAllOfSourceOneOf1RepositoryRepository {
	return v.value
}

func (v *NullableHelmResponseAllOfSourceOneOf1RepositoryRepository) Set(val *HelmResponseAllOfSourceOneOf1RepositoryRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmResponseAllOfSourceOneOf1RepositoryRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmResponseAllOfSourceOneOf1RepositoryRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmResponseAllOfSourceOneOf1RepositoryRepository(val *HelmResponseAllOfSourceOneOf1RepositoryRepository) *NullableHelmResponseAllOfSourceOneOf1RepositoryRepository {
	return &NullableHelmResponseAllOfSourceOneOf1RepositoryRepository{value: val, isSet: true}
}

func (v NullableHelmResponseAllOfSourceOneOf1RepositoryRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmResponseAllOfSourceOneOf1RepositoryRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}