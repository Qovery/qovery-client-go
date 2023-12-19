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

// checks if the HelmDeployRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmDeployRequest{}

// HelmDeployRequest struct for HelmDeployRequest
type HelmDeployRequest struct {
	// version of the chart to deploy. Cannot be set if `git_commit_id` is defined
	ChartVersion *string `json:"chart_version,omitempty"`
	// Commit to deploy for chart source. Cannot be set if `version` is defined
	GitCommitId *string `json:"git_commit_id,omitempty"`
	// Commit to deploy for values override
	ValuesOverrideGitCommitId *string `json:"values_override_git_commit_id,omitempty"`
}

// NewHelmDeployRequest instantiates a new HelmDeployRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmDeployRequest() *HelmDeployRequest {
	this := HelmDeployRequest{}
	return &this
}

// NewHelmDeployRequestWithDefaults instantiates a new HelmDeployRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmDeployRequestWithDefaults() *HelmDeployRequest {
	this := HelmDeployRequest{}
	return &this
}

// GetChartVersion returns the ChartVersion field value if set, zero value otherwise.
func (o *HelmDeployRequest) GetChartVersion() string {
	if o == nil || IsNil(o.ChartVersion) {
		var ret string
		return ret
	}
	return *o.ChartVersion
}

// GetChartVersionOk returns a tuple with the ChartVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmDeployRequest) GetChartVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ChartVersion) {
		return nil, false
	}
	return o.ChartVersion, true
}

// HasChartVersion returns a boolean if a field has been set.
func (o *HelmDeployRequest) HasChartVersion() bool {
	if o != nil && !IsNil(o.ChartVersion) {
		return true
	}

	return false
}

// SetChartVersion gets a reference to the given string and assigns it to the ChartVersion field.
func (o *HelmDeployRequest) SetChartVersion(v string) {
	o.ChartVersion = &v
}

// GetGitCommitId returns the GitCommitId field value if set, zero value otherwise.
func (o *HelmDeployRequest) GetGitCommitId() string {
	if o == nil || IsNil(o.GitCommitId) {
		var ret string
		return ret
	}
	return *o.GitCommitId
}

// GetGitCommitIdOk returns a tuple with the GitCommitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmDeployRequest) GetGitCommitIdOk() (*string, bool) {
	if o == nil || IsNil(o.GitCommitId) {
		return nil, false
	}
	return o.GitCommitId, true
}

// HasGitCommitId returns a boolean if a field has been set.
func (o *HelmDeployRequest) HasGitCommitId() bool {
	if o != nil && !IsNil(o.GitCommitId) {
		return true
	}

	return false
}

// SetGitCommitId gets a reference to the given string and assigns it to the GitCommitId field.
func (o *HelmDeployRequest) SetGitCommitId(v string) {
	o.GitCommitId = &v
}

// GetValuesOverrideGitCommitId returns the ValuesOverrideGitCommitId field value if set, zero value otherwise.
func (o *HelmDeployRequest) GetValuesOverrideGitCommitId() string {
	if o == nil || IsNil(o.ValuesOverrideGitCommitId) {
		var ret string
		return ret
	}
	return *o.ValuesOverrideGitCommitId
}

// GetValuesOverrideGitCommitIdOk returns a tuple with the ValuesOverrideGitCommitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmDeployRequest) GetValuesOverrideGitCommitIdOk() (*string, bool) {
	if o == nil || IsNil(o.ValuesOverrideGitCommitId) {
		return nil, false
	}
	return o.ValuesOverrideGitCommitId, true
}

// HasValuesOverrideGitCommitId returns a boolean if a field has been set.
func (o *HelmDeployRequest) HasValuesOverrideGitCommitId() bool {
	if o != nil && !IsNil(o.ValuesOverrideGitCommitId) {
		return true
	}

	return false
}

// SetValuesOverrideGitCommitId gets a reference to the given string and assigns it to the ValuesOverrideGitCommitId field.
func (o *HelmDeployRequest) SetValuesOverrideGitCommitId(v string) {
	o.ValuesOverrideGitCommitId = &v
}

func (o HelmDeployRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmDeployRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChartVersion) {
		toSerialize["chart_version"] = o.ChartVersion
	}
	if !IsNil(o.GitCommitId) {
		toSerialize["git_commit_id"] = o.GitCommitId
	}
	if !IsNil(o.ValuesOverrideGitCommitId) {
		toSerialize["values_override_git_commit_id"] = o.ValuesOverrideGitCommitId
	}
	return toSerialize, nil
}

type NullableHelmDeployRequest struct {
	value *HelmDeployRequest
	isSet bool
}

func (v NullableHelmDeployRequest) Get() *HelmDeployRequest {
	return v.value
}

func (v *NullableHelmDeployRequest) Set(val *HelmDeployRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmDeployRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmDeployRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmDeployRequest(val *HelmDeployRequest) *NullableHelmDeployRequest {
	return &NullableHelmDeployRequest{value: val, isSet: true}
}

func (v NullableHelmDeployRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmDeployRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}