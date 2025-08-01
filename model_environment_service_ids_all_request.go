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

// checks if the EnvironmentServiceIdsAllRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentServiceIdsAllRequest{}

// EnvironmentServiceIdsAllRequest struct for EnvironmentServiceIdsAllRequest
type EnvironmentServiceIdsAllRequest struct {
	ApplicationIds       []string `json:"application_ids,omitempty"`
	ContainerIds         []string `json:"container_ids,omitempty"`
	DatabaseIds          []string `json:"database_ids,omitempty"`
	JobIds               []string `json:"job_ids,omitempty"`
	HelmIds              []string `json:"helm_ids,omitempty"`
	TerraformIds         []string `json:"terraform_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnvironmentServiceIdsAllRequest EnvironmentServiceIdsAllRequest

// NewEnvironmentServiceIdsAllRequest instantiates a new EnvironmentServiceIdsAllRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentServiceIdsAllRequest() *EnvironmentServiceIdsAllRequest {
	this := EnvironmentServiceIdsAllRequest{}
	return &this
}

// NewEnvironmentServiceIdsAllRequestWithDefaults instantiates a new EnvironmentServiceIdsAllRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentServiceIdsAllRequestWithDefaults() *EnvironmentServiceIdsAllRequest {
	this := EnvironmentServiceIdsAllRequest{}
	return &this
}

// GetApplicationIds returns the ApplicationIds field value if set, zero value otherwise.
func (o *EnvironmentServiceIdsAllRequest) GetApplicationIds() []string {
	if o == nil || IsNil(o.ApplicationIds) {
		var ret []string
		return ret
	}
	return o.ApplicationIds
}

// GetApplicationIdsOk returns a tuple with the ApplicationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentServiceIdsAllRequest) GetApplicationIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ApplicationIds) {
		return nil, false
	}
	return o.ApplicationIds, true
}

// HasApplicationIds returns a boolean if a field has been set.
func (o *EnvironmentServiceIdsAllRequest) HasApplicationIds() bool {
	if o != nil && !IsNil(o.ApplicationIds) {
		return true
	}

	return false
}

// SetApplicationIds gets a reference to the given []string and assigns it to the ApplicationIds field.
func (o *EnvironmentServiceIdsAllRequest) SetApplicationIds(v []string) {
	o.ApplicationIds = v
}

// GetContainerIds returns the ContainerIds field value if set, zero value otherwise.
func (o *EnvironmentServiceIdsAllRequest) GetContainerIds() []string {
	if o == nil || IsNil(o.ContainerIds) {
		var ret []string
		return ret
	}
	return o.ContainerIds
}

// GetContainerIdsOk returns a tuple with the ContainerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentServiceIdsAllRequest) GetContainerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ContainerIds) {
		return nil, false
	}
	return o.ContainerIds, true
}

// HasContainerIds returns a boolean if a field has been set.
func (o *EnvironmentServiceIdsAllRequest) HasContainerIds() bool {
	if o != nil && !IsNil(o.ContainerIds) {
		return true
	}

	return false
}

// SetContainerIds gets a reference to the given []string and assigns it to the ContainerIds field.
func (o *EnvironmentServiceIdsAllRequest) SetContainerIds(v []string) {
	o.ContainerIds = v
}

// GetDatabaseIds returns the DatabaseIds field value if set, zero value otherwise.
func (o *EnvironmentServiceIdsAllRequest) GetDatabaseIds() []string {
	if o == nil || IsNil(o.DatabaseIds) {
		var ret []string
		return ret
	}
	return o.DatabaseIds
}

// GetDatabaseIdsOk returns a tuple with the DatabaseIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentServiceIdsAllRequest) GetDatabaseIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.DatabaseIds) {
		return nil, false
	}
	return o.DatabaseIds, true
}

// HasDatabaseIds returns a boolean if a field has been set.
func (o *EnvironmentServiceIdsAllRequest) HasDatabaseIds() bool {
	if o != nil && !IsNil(o.DatabaseIds) {
		return true
	}

	return false
}

// SetDatabaseIds gets a reference to the given []string and assigns it to the DatabaseIds field.
func (o *EnvironmentServiceIdsAllRequest) SetDatabaseIds(v []string) {
	o.DatabaseIds = v
}

// GetJobIds returns the JobIds field value if set, zero value otherwise.
func (o *EnvironmentServiceIdsAllRequest) GetJobIds() []string {
	if o == nil || IsNil(o.JobIds) {
		var ret []string
		return ret
	}
	return o.JobIds
}

// GetJobIdsOk returns a tuple with the JobIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentServiceIdsAllRequest) GetJobIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.JobIds) {
		return nil, false
	}
	return o.JobIds, true
}

// HasJobIds returns a boolean if a field has been set.
func (o *EnvironmentServiceIdsAllRequest) HasJobIds() bool {
	if o != nil && !IsNil(o.JobIds) {
		return true
	}

	return false
}

// SetJobIds gets a reference to the given []string and assigns it to the JobIds field.
func (o *EnvironmentServiceIdsAllRequest) SetJobIds(v []string) {
	o.JobIds = v
}

// GetHelmIds returns the HelmIds field value if set, zero value otherwise.
func (o *EnvironmentServiceIdsAllRequest) GetHelmIds() []string {
	if o == nil || IsNil(o.HelmIds) {
		var ret []string
		return ret
	}
	return o.HelmIds
}

// GetHelmIdsOk returns a tuple with the HelmIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentServiceIdsAllRequest) GetHelmIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.HelmIds) {
		return nil, false
	}
	return o.HelmIds, true
}

// HasHelmIds returns a boolean if a field has been set.
func (o *EnvironmentServiceIdsAllRequest) HasHelmIds() bool {
	if o != nil && !IsNil(o.HelmIds) {
		return true
	}

	return false
}

// SetHelmIds gets a reference to the given []string and assigns it to the HelmIds field.
func (o *EnvironmentServiceIdsAllRequest) SetHelmIds(v []string) {
	o.HelmIds = v
}

// GetTerraformIds returns the TerraformIds field value if set, zero value otherwise.
func (o *EnvironmentServiceIdsAllRequest) GetTerraformIds() []string {
	if o == nil || IsNil(o.TerraformIds) {
		var ret []string
		return ret
	}
	return o.TerraformIds
}

// GetTerraformIdsOk returns a tuple with the TerraformIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentServiceIdsAllRequest) GetTerraformIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TerraformIds) {
		return nil, false
	}
	return o.TerraformIds, true
}

// HasTerraformIds returns a boolean if a field has been set.
func (o *EnvironmentServiceIdsAllRequest) HasTerraformIds() bool {
	if o != nil && !IsNil(o.TerraformIds) {
		return true
	}

	return false
}

// SetTerraformIds gets a reference to the given []string and assigns it to the TerraformIds field.
func (o *EnvironmentServiceIdsAllRequest) SetTerraformIds(v []string) {
	o.TerraformIds = v
}

func (o EnvironmentServiceIdsAllRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentServiceIdsAllRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationIds) {
		toSerialize["application_ids"] = o.ApplicationIds
	}
	if !IsNil(o.ContainerIds) {
		toSerialize["container_ids"] = o.ContainerIds
	}
	if !IsNil(o.DatabaseIds) {
		toSerialize["database_ids"] = o.DatabaseIds
	}
	if !IsNil(o.JobIds) {
		toSerialize["job_ids"] = o.JobIds
	}
	if !IsNil(o.HelmIds) {
		toSerialize["helm_ids"] = o.HelmIds
	}
	if !IsNil(o.TerraformIds) {
		toSerialize["terraform_ids"] = o.TerraformIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnvironmentServiceIdsAllRequest) UnmarshalJSON(data []byte) (err error) {
	varEnvironmentServiceIdsAllRequest := _EnvironmentServiceIdsAllRequest{}

	err = json.Unmarshal(data, &varEnvironmentServiceIdsAllRequest)

	if err != nil {
		return err
	}

	*o = EnvironmentServiceIdsAllRequest(varEnvironmentServiceIdsAllRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "application_ids")
		delete(additionalProperties, "container_ids")
		delete(additionalProperties, "database_ids")
		delete(additionalProperties, "job_ids")
		delete(additionalProperties, "helm_ids")
		delete(additionalProperties, "terraform_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnvironmentServiceIdsAllRequest struct {
	value *EnvironmentServiceIdsAllRequest
	isSet bool
}

func (v NullableEnvironmentServiceIdsAllRequest) Get() *EnvironmentServiceIdsAllRequest {
	return v.value
}

func (v *NullableEnvironmentServiceIdsAllRequest) Set(val *EnvironmentServiceIdsAllRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentServiceIdsAllRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentServiceIdsAllRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentServiceIdsAllRequest(val *EnvironmentServiceIdsAllRequest) *NullableEnvironmentServiceIdsAllRequest {
	return &NullableEnvironmentServiceIdsAllRequest{value: val, isSet: true}
}

func (v NullableEnvironmentServiceIdsAllRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentServiceIdsAllRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
