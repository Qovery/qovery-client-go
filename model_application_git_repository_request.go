/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"encoding/json"
)

// ApplicationGitRepositoryRequest struct for ApplicationGitRepositoryRequest
type ApplicationGitRepositoryRequest struct {
	// application git repository URL
	Url string `json:"url"`
	// Name of the branch to use. This is optional If not specified, then the branch used is the `main` or `master` one
	Branch *string `json:"branch,omitempty"`
	// indicates the root path of the application.
	RootPath string `json:"root_path"`
}

// NewApplicationGitRepositoryRequest instantiates a new ApplicationGitRepositoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationGitRepositoryRequest(url string, rootPath string) *ApplicationGitRepositoryRequest {
	this := ApplicationGitRepositoryRequest{}
	this.Url = url
	this.RootPath = rootPath
	return &this
}

// NewApplicationGitRepositoryRequestWithDefaults instantiates a new ApplicationGitRepositoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationGitRepositoryRequestWithDefaults() *ApplicationGitRepositoryRequest {
	this := ApplicationGitRepositoryRequest{}
	var rootPath string = "/"
	this.RootPath = rootPath
	return &this
}

// GetUrl returns the Url field value
func (o *ApplicationGitRepositoryRequest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ApplicationGitRepositoryRequest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ApplicationGitRepositoryRequest) SetUrl(v string) {
	o.Url = v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *ApplicationGitRepositoryRequest) GetBranch() string {
	if o == nil || o.Branch == nil {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationGitRepositoryRequest) GetBranchOk() (*string, bool) {
	if o == nil || o.Branch == nil {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *ApplicationGitRepositoryRequest) HasBranch() bool {
	if o != nil && o.Branch != nil {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *ApplicationGitRepositoryRequest) SetBranch(v string) {
	o.Branch = &v
}

// GetRootPath returns the RootPath field value
func (o *ApplicationGitRepositoryRequest) GetRootPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RootPath
}

// GetRootPathOk returns a tuple with the RootPath field value
// and a boolean to check if the value has been set.
func (o *ApplicationGitRepositoryRequest) GetRootPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootPath, true
}

// SetRootPath sets field value
func (o *ApplicationGitRepositoryRequest) SetRootPath(v string) {
	o.RootPath = v
}

func (o ApplicationGitRepositoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.Branch != nil {
		toSerialize["branch"] = o.Branch
	}
	if true {
		toSerialize["root_path"] = o.RootPath
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationGitRepositoryRequest struct {
	value *ApplicationGitRepositoryRequest
	isSet bool
}

func (v NullableApplicationGitRepositoryRequest) Get() *ApplicationGitRepositoryRequest {
	return v.value
}

func (v *NullableApplicationGitRepositoryRequest) Set(val *ApplicationGitRepositoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationGitRepositoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationGitRepositoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationGitRepositoryRequest(val *ApplicationGitRepositoryRequest) *NullableApplicationGitRepositoryRequest {
	return &NullableApplicationGitRepositoryRequest{value: val, isSet: true}
}

func (v NullableApplicationGitRepositoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationGitRepositoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
