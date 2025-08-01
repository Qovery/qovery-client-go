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
	"fmt"
)

// checks if the DeploymentHistoryEnvironmentPaginatedResponseListV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentHistoryEnvironmentPaginatedResponseListV2{}

// DeploymentHistoryEnvironmentPaginatedResponseListV2 struct for DeploymentHistoryEnvironmentPaginatedResponseListV2
type DeploymentHistoryEnvironmentPaginatedResponseListV2 struct {
	Page                 float32                          `json:"page"`
	PageSize             float32                          `json:"page_size"`
	Results              []DeploymentHistoryEnvironmentV2 `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeploymentHistoryEnvironmentPaginatedResponseListV2 DeploymentHistoryEnvironmentPaginatedResponseListV2

// NewDeploymentHistoryEnvironmentPaginatedResponseListV2 instantiates a new DeploymentHistoryEnvironmentPaginatedResponseListV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistoryEnvironmentPaginatedResponseListV2(page float32, pageSize float32) *DeploymentHistoryEnvironmentPaginatedResponseListV2 {
	this := DeploymentHistoryEnvironmentPaginatedResponseListV2{}
	this.Page = page
	this.PageSize = pageSize
	return &this
}

// NewDeploymentHistoryEnvironmentPaginatedResponseListV2WithDefaults instantiates a new DeploymentHistoryEnvironmentPaginatedResponseListV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryEnvironmentPaginatedResponseListV2WithDefaults() *DeploymentHistoryEnvironmentPaginatedResponseListV2 {
	this := DeploymentHistoryEnvironmentPaginatedResponseListV2{}
	return &this
}

// GetPage returns the Page field value
func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) GetPage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) GetPageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) SetPage(v float32) {
	o.Page = v
}

// GetPageSize returns the PageSize field value
func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) GetPageSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) GetPageSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) SetPageSize(v float32) {
	o.PageSize = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) GetResults() []DeploymentHistoryEnvironmentV2 {
	if o == nil || IsNil(o.Results) {
		var ret []DeploymentHistoryEnvironmentV2
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) GetResultsOk() ([]DeploymentHistoryEnvironmentV2, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DeploymentHistoryEnvironmentV2 and assigns it to the Results field.
func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) SetResults(v []DeploymentHistoryEnvironmentV2) {
	o.Results = v
}

func (o DeploymentHistoryEnvironmentPaginatedResponseListV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentHistoryEnvironmentPaginatedResponseListV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["page"] = o.Page
	toSerialize["page_size"] = o.PageSize
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeploymentHistoryEnvironmentPaginatedResponseListV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"page",
		"page_size",
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

	varDeploymentHistoryEnvironmentPaginatedResponseListV2 := _DeploymentHistoryEnvironmentPaginatedResponseListV2{}

	err = json.Unmarshal(data, &varDeploymentHistoryEnvironmentPaginatedResponseListV2)

	if err != nil {
		return err
	}

	*o = DeploymentHistoryEnvironmentPaginatedResponseListV2(varDeploymentHistoryEnvironmentPaginatedResponseListV2)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "page")
		delete(additionalProperties, "page_size")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeploymentHistoryEnvironmentPaginatedResponseListV2 struct {
	value *DeploymentHistoryEnvironmentPaginatedResponseListV2
	isSet bool
}

func (v NullableDeploymentHistoryEnvironmentPaginatedResponseListV2) Get() *DeploymentHistoryEnvironmentPaginatedResponseListV2 {
	return v.value
}

func (v *NullableDeploymentHistoryEnvironmentPaginatedResponseListV2) Set(val *DeploymentHistoryEnvironmentPaginatedResponseListV2) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistoryEnvironmentPaginatedResponseListV2) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistoryEnvironmentPaginatedResponseListV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistoryEnvironmentPaginatedResponseListV2(val *DeploymentHistoryEnvironmentPaginatedResponseListV2) *NullableDeploymentHistoryEnvironmentPaginatedResponseListV2 {
	return &NullableDeploymentHistoryEnvironmentPaginatedResponseListV2{value: val, isSet: true}
}

func (v NullableDeploymentHistoryEnvironmentPaginatedResponseListV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistoryEnvironmentPaginatedResponseListV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
