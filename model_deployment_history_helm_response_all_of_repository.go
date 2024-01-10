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

// checks if the DeploymentHistoryHelmResponseAllOfRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentHistoryHelmResponseAllOfRepository{}

// DeploymentHistoryHelmResponseAllOfRepository If the chart source if from a repository, the chart name and its version
type DeploymentHistoryHelmResponseAllOfRepository struct {
	ChartName    *string `json:"chart_name,omitempty"`
	ChartVersion *string `json:"chart_version,omitempty"`
}

// NewDeploymentHistoryHelmResponseAllOfRepository instantiates a new DeploymentHistoryHelmResponseAllOfRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentHistoryHelmResponseAllOfRepository() *DeploymentHistoryHelmResponseAllOfRepository {
	this := DeploymentHistoryHelmResponseAllOfRepository{}
	return &this
}

// NewDeploymentHistoryHelmResponseAllOfRepositoryWithDefaults instantiates a new DeploymentHistoryHelmResponseAllOfRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentHistoryHelmResponseAllOfRepositoryWithDefaults() *DeploymentHistoryHelmResponseAllOfRepository {
	this := DeploymentHistoryHelmResponseAllOfRepository{}
	return &this
}

// GetChartName returns the ChartName field value if set, zero value otherwise.
func (o *DeploymentHistoryHelmResponseAllOfRepository) GetChartName() string {
	if o == nil || IsNil(o.ChartName) {
		var ret string
		return ret
	}
	return *o.ChartName
}

// GetChartNameOk returns a tuple with the ChartName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryHelmResponseAllOfRepository) GetChartNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChartName) {
		return nil, false
	}
	return o.ChartName, true
}

// HasChartName returns a boolean if a field has been set.
func (o *DeploymentHistoryHelmResponseAllOfRepository) HasChartName() bool {
	if o != nil && !IsNil(o.ChartName) {
		return true
	}

	return false
}

// SetChartName gets a reference to the given string and assigns it to the ChartName field.
func (o *DeploymentHistoryHelmResponseAllOfRepository) SetChartName(v string) {
	o.ChartName = &v
}

// GetChartVersion returns the ChartVersion field value if set, zero value otherwise.
func (o *DeploymentHistoryHelmResponseAllOfRepository) GetChartVersion() string {
	if o == nil || IsNil(o.ChartVersion) {
		var ret string
		return ret
	}
	return *o.ChartVersion
}

// GetChartVersionOk returns a tuple with the ChartVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentHistoryHelmResponseAllOfRepository) GetChartVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ChartVersion) {
		return nil, false
	}
	return o.ChartVersion, true
}

// HasChartVersion returns a boolean if a field has been set.
func (o *DeploymentHistoryHelmResponseAllOfRepository) HasChartVersion() bool {
	if o != nil && !IsNil(o.ChartVersion) {
		return true
	}

	return false
}

// SetChartVersion gets a reference to the given string and assigns it to the ChartVersion field.
func (o *DeploymentHistoryHelmResponseAllOfRepository) SetChartVersion(v string) {
	o.ChartVersion = &v
}

func (o DeploymentHistoryHelmResponseAllOfRepository) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentHistoryHelmResponseAllOfRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChartName) {
		toSerialize["chart_name"] = o.ChartName
	}
	if !IsNil(o.ChartVersion) {
		toSerialize["chart_version"] = o.ChartVersion
	}
	return toSerialize, nil
}

type NullableDeploymentHistoryHelmResponseAllOfRepository struct {
	value *DeploymentHistoryHelmResponseAllOfRepository
	isSet bool
}

func (v NullableDeploymentHistoryHelmResponseAllOfRepository) Get() *DeploymentHistoryHelmResponseAllOfRepository {
	return v.value
}

func (v *NullableDeploymentHistoryHelmResponseAllOfRepository) Set(val *DeploymentHistoryHelmResponseAllOfRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentHistoryHelmResponseAllOfRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentHistoryHelmResponseAllOfRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentHistoryHelmResponseAllOfRepository(val *DeploymentHistoryHelmResponseAllOfRepository) *NullableDeploymentHistoryHelmResponseAllOfRepository {
	return &NullableDeploymentHistoryHelmResponseAllOfRepository{value: val, isSet: true}
}

func (v NullableDeploymentHistoryHelmResponseAllOfRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentHistoryHelmResponseAllOfRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}