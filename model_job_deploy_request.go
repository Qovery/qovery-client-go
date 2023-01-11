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

// JobDeployRequest struct for JobDeployRequest
type JobDeployRequest struct {
	// Image tag to deploy.   Cannot be set if `git_commit_id` is defined
	ImageTag *string `json:"image_tag,omitempty"`
	// Commit to deploy Cannot be set if `image_tag` is defined
	GitCommitId *string `json:"git_commit_id,omitempty"`
}

// NewJobDeployRequest instantiates a new JobDeployRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobDeployRequest() *JobDeployRequest {
	this := JobDeployRequest{}
	return &this
}

// NewJobDeployRequestWithDefaults instantiates a new JobDeployRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobDeployRequestWithDefaults() *JobDeployRequest {
	this := JobDeployRequest{}
	return &this
}

// GetImageTag returns the ImageTag field value if set, zero value otherwise.
func (o *JobDeployRequest) GetImageTag() string {
	if o == nil || o.ImageTag == nil {
		var ret string
		return ret
	}
	return *o.ImageTag
}

// GetImageTagOk returns a tuple with the ImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDeployRequest) GetImageTagOk() (*string, bool) {
	if o == nil || o.ImageTag == nil {
		return nil, false
	}
	return o.ImageTag, true
}

// HasImageTag returns a boolean if a field has been set.
func (o *JobDeployRequest) HasImageTag() bool {
	if o != nil && o.ImageTag != nil {
		return true
	}

	return false
}

// SetImageTag gets a reference to the given string and assigns it to the ImageTag field.
func (o *JobDeployRequest) SetImageTag(v string) {
	o.ImageTag = &v
}

// GetGitCommitId returns the GitCommitId field value if set, zero value otherwise.
func (o *JobDeployRequest) GetGitCommitId() string {
	if o == nil || o.GitCommitId == nil {
		var ret string
		return ret
	}
	return *o.GitCommitId
}

// GetGitCommitIdOk returns a tuple with the GitCommitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDeployRequest) GetGitCommitIdOk() (*string, bool) {
	if o == nil || o.GitCommitId == nil {
		return nil, false
	}
	return o.GitCommitId, true
}

// HasGitCommitId returns a boolean if a field has been set.
func (o *JobDeployRequest) HasGitCommitId() bool {
	if o != nil && o.GitCommitId != nil {
		return true
	}

	return false
}

// SetGitCommitId gets a reference to the given string and assigns it to the GitCommitId field.
func (o *JobDeployRequest) SetGitCommitId(v string) {
	o.GitCommitId = &v
}

func (o JobDeployRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImageTag != nil {
		toSerialize["image_tag"] = o.ImageTag
	}
	if o.GitCommitId != nil {
		toSerialize["git_commit_id"] = o.GitCommitId
	}
	return json.Marshal(toSerialize)
}

type NullableJobDeployRequest struct {
	value *JobDeployRequest
	isSet bool
}

func (v NullableJobDeployRequest) Get() *JobDeployRequest {
	return v.value
}

func (v *NullableJobDeployRequest) Set(val *JobDeployRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableJobDeployRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableJobDeployRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobDeployRequest(val *JobDeployRequest) *NullableJobDeployRequest {
	return &NullableJobDeployRequest{value: val, isSet: true}
}

func (v NullableJobDeployRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobDeployRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}