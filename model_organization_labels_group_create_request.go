/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the OrganizationLabelsGroupCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationLabelsGroupCreateRequest{}

// OrganizationLabelsGroupCreateRequest struct for OrganizationLabelsGroupCreateRequest
type OrganizationLabelsGroupCreateRequest struct {
	// name of the labels group
	Name   string  `json:"name"`
	Labels []Label `json:"labels"`
}

type _OrganizationLabelsGroupCreateRequest OrganizationLabelsGroupCreateRequest

// NewOrganizationLabelsGroupCreateRequest instantiates a new OrganizationLabelsGroupCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationLabelsGroupCreateRequest(name string, labels []Label) *OrganizationLabelsGroupCreateRequest {
	this := OrganizationLabelsGroupCreateRequest{}
	this.Name = name
	this.Labels = labels
	return &this
}

// NewOrganizationLabelsGroupCreateRequestWithDefaults instantiates a new OrganizationLabelsGroupCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationLabelsGroupCreateRequestWithDefaults() *OrganizationLabelsGroupCreateRequest {
	this := OrganizationLabelsGroupCreateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *OrganizationLabelsGroupCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationLabelsGroupCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationLabelsGroupCreateRequest) SetName(v string) {
	o.Name = v
}

// GetLabels returns the Labels field value
func (o *OrganizationLabelsGroupCreateRequest) GetLabels() []Label {
	if o == nil {
		var ret []Label
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *OrganizationLabelsGroupCreateRequest) GetLabelsOk() ([]Label, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *OrganizationLabelsGroupCreateRequest) SetLabels(v []Label) {
	o.Labels = v
}

func (o OrganizationLabelsGroupCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationLabelsGroupCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["labels"] = o.Labels
	return toSerialize, nil
}

func (o *OrganizationLabelsGroupCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"labels",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrganizationLabelsGroupCreateRequest := _OrganizationLabelsGroupCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrganizationLabelsGroupCreateRequest)

	if err != nil {
		return err
	}

	*o = OrganizationLabelsGroupCreateRequest(varOrganizationLabelsGroupCreateRequest)

	return err
}

type NullableOrganizationLabelsGroupCreateRequest struct {
	value *OrganizationLabelsGroupCreateRequest
	isSet bool
}

func (v NullableOrganizationLabelsGroupCreateRequest) Get() *OrganizationLabelsGroupCreateRequest {
	return v.value
}

func (v *NullableOrganizationLabelsGroupCreateRequest) Set(val *OrganizationLabelsGroupCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationLabelsGroupCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationLabelsGroupCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationLabelsGroupCreateRequest(val *OrganizationLabelsGroupCreateRequest) *NullableOrganizationLabelsGroupCreateRequest {
	return &NullableOrganizationLabelsGroupCreateRequest{value: val, isSet: true}
}

func (v NullableOrganizationLabelsGroupCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationLabelsGroupCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
