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

// checks if the DeployAllRequestApplicationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeployAllRequestApplicationsInner{}

// DeployAllRequestApplicationsInner struct for DeployAllRequestApplicationsInner
type DeployAllRequestApplicationsInner struct {
	// id of the application to be deployed.
	ApplicationId string `json:"application_id"`
	// Commit ID to deploy. Can be empty only if the service has been already deployed (in this case the service version won't be changed)
	GitCommitId          *string `json:"git_commit_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeployAllRequestApplicationsInner DeployAllRequestApplicationsInner

// NewDeployAllRequestApplicationsInner instantiates a new DeployAllRequestApplicationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployAllRequestApplicationsInner(applicationId string) *DeployAllRequestApplicationsInner {
	this := DeployAllRequestApplicationsInner{}
	this.ApplicationId = applicationId
	return &this
}

// NewDeployAllRequestApplicationsInnerWithDefaults instantiates a new DeployAllRequestApplicationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeployAllRequestApplicationsInnerWithDefaults() *DeployAllRequestApplicationsInner {
	this := DeployAllRequestApplicationsInner{}
	return &this
}

// GetApplicationId returns the ApplicationId field value
func (o *DeployAllRequestApplicationsInner) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *DeployAllRequestApplicationsInner) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *DeployAllRequestApplicationsInner) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetGitCommitId returns the GitCommitId field value if set, zero value otherwise.
func (o *DeployAllRequestApplicationsInner) GetGitCommitId() string {
	if o == nil || IsNil(o.GitCommitId) {
		var ret string
		return ret
	}
	return *o.GitCommitId
}

// GetGitCommitIdOk returns a tuple with the GitCommitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeployAllRequestApplicationsInner) GetGitCommitIdOk() (*string, bool) {
	if o == nil || IsNil(o.GitCommitId) {
		return nil, false
	}
	return o.GitCommitId, true
}

// HasGitCommitId returns a boolean if a field has been set.
func (o *DeployAllRequestApplicationsInner) HasGitCommitId() bool {
	if o != nil && !IsNil(o.GitCommitId) {
		return true
	}

	return false
}

// SetGitCommitId gets a reference to the given string and assigns it to the GitCommitId field.
func (o *DeployAllRequestApplicationsInner) SetGitCommitId(v string) {
	o.GitCommitId = &v
}

func (o DeployAllRequestApplicationsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeployAllRequestApplicationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["application_id"] = o.ApplicationId
	if !IsNil(o.GitCommitId) {
		toSerialize["git_commit_id"] = o.GitCommitId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeployAllRequestApplicationsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"application_id",
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

	varDeployAllRequestApplicationsInner := _DeployAllRequestApplicationsInner{}

	err = json.Unmarshal(data, &varDeployAllRequestApplicationsInner)

	if err != nil {
		return err
	}

	*o = DeployAllRequestApplicationsInner(varDeployAllRequestApplicationsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "application_id")
		delete(additionalProperties, "git_commit_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeployAllRequestApplicationsInner struct {
	value *DeployAllRequestApplicationsInner
	isSet bool
}

func (v NullableDeployAllRequestApplicationsInner) Get() *DeployAllRequestApplicationsInner {
	return v.value
}

func (v *NullableDeployAllRequestApplicationsInner) Set(val *DeployAllRequestApplicationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDeployAllRequestApplicationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDeployAllRequestApplicationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeployAllRequestApplicationsInner(val *DeployAllRequestApplicationsInner) *NullableDeployAllRequestApplicationsInner {
	return &NullableDeployAllRequestApplicationsInner{value: val, isSet: true}
}

func (v NullableDeployAllRequestApplicationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeployAllRequestApplicationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
