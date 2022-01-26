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

// ApplicationEditRequest struct for ApplicationEditRequest
type ApplicationEditRequest struct {
	// name is case insensitive
	Name *string `json:"name,omitempty"`
	// give a description to this application
	Description   *string                          `json:"description,omitempty"`
	GitRepository *ApplicationGitRepositoryRequest `json:"git_repository,omitempty"`
	// `DOCKER` requires `dockerfile_path` `BUILDPACKS` does not require any `dockerfile_path`
	BuildMode *string `json:"build_mode,omitempty"`
	// The path of the associated Dockerfile
	DockerfilePath *string `json:"dockerfile_path,omitempty"`
	// Development language of the application
	BuildpackLanguage *string `json:"buildpack_language,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu *float32 `json:"cpu,omitempty"`
	// unit is MB. 1024 MB = 1GB
	Memory *float32 `json:"memory,omitempty"`
	// Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no application running.
	MinRunningInstances *int32 `json:"min_running_instances,omitempty"`
	// Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.
	MaxRunningInstances *int32       `json:"max_running_instances,omitempty"`
	Healthcheck         *Healthcheck `json:"healthcheck,omitempty"`
	// Specify if the environment preview option is activated or not for this application. If activated, a preview environment will be automatically cloned at each pull request.
	EnvPreview *bool `json:"env_preview,omitempty"`
	// Specify if the sticky session option (also called persistant session) is activated or not for this application. If activated, user will be redirected by the load balancer to the same instance each time he access to the application.
	StickySession *bool                                `json:"sticky-session,omitempty"`
	Storage       *[]ApplicationStorageResponseStorage `json:"storage,omitempty"`
	Ports         *[]ApplicationPortResponsePorts      `json:"ports,omitempty"`
}

// NewApplicationEditRequest instantiates a new ApplicationEditRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationEditRequest() *ApplicationEditRequest {
	this := ApplicationEditRequest{}
	return &this
}

// NewApplicationEditRequestWithDefaults instantiates a new ApplicationEditRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationEditRequestWithDefaults() *ApplicationEditRequest {
	this := ApplicationEditRequest{}
	var cpu float32 = 250
	this.Cpu = &cpu
	var memory float32 = 256
	this.Memory = &memory
	var minRunningInstances int32 = 1
	this.MinRunningInstances = &minRunningInstances
	var maxRunningInstances int32 = 1
	this.MaxRunningInstances = &maxRunningInstances
	var envPreview bool = true
	this.EnvPreview = &envPreview
	var stickySession bool = false
	this.StickySession = &stickySession
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationEditRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationEditRequest) SetDescription(v string) {
	o.Description = &v
}

// GetGitRepository returns the GitRepository field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetGitRepository() ApplicationGitRepositoryRequest {
	if o == nil || o.GitRepository == nil {
		var ret ApplicationGitRepositoryRequest
		return ret
	}
	return *o.GitRepository
}

// GetGitRepositoryOk returns a tuple with the GitRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetGitRepositoryOk() (*ApplicationGitRepositoryRequest, bool) {
	if o == nil || o.GitRepository == nil {
		return nil, false
	}
	return o.GitRepository, true
}

// HasGitRepository returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasGitRepository() bool {
	if o != nil && o.GitRepository != nil {
		return true
	}

	return false
}

// SetGitRepository gets a reference to the given ApplicationGitRepositoryRequest and assigns it to the GitRepository field.
func (o *ApplicationEditRequest) SetGitRepository(v ApplicationGitRepositoryRequest) {
	o.GitRepository = &v
}

// GetBuildMode returns the BuildMode field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetBuildMode() string {
	if o == nil || o.BuildMode == nil {
		var ret string
		return ret
	}
	return *o.BuildMode
}

// GetBuildModeOk returns a tuple with the BuildMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetBuildModeOk() (*string, bool) {
	if o == nil || o.BuildMode == nil {
		return nil, false
	}
	return o.BuildMode, true
}

// HasBuildMode returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasBuildMode() bool {
	if o != nil && o.BuildMode != nil {
		return true
	}

	return false
}

// SetBuildMode gets a reference to the given string and assigns it to the BuildMode field.
func (o *ApplicationEditRequest) SetBuildMode(v string) {
	o.BuildMode = &v
}

// GetDockerfilePath returns the DockerfilePath field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetDockerfilePath() string {
	if o == nil || o.DockerfilePath == nil {
		var ret string
		return ret
	}
	return *o.DockerfilePath
}

// GetDockerfilePathOk returns a tuple with the DockerfilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetDockerfilePathOk() (*string, bool) {
	if o == nil || o.DockerfilePath == nil {
		return nil, false
	}
	return o.DockerfilePath, true
}

// HasDockerfilePath returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasDockerfilePath() bool {
	if o != nil && o.DockerfilePath != nil {
		return true
	}

	return false
}

// SetDockerfilePath gets a reference to the given string and assigns it to the DockerfilePath field.
func (o *ApplicationEditRequest) SetDockerfilePath(v string) {
	o.DockerfilePath = &v
}

// GetBuildpackLanguage returns the BuildpackLanguage field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetBuildpackLanguage() string {
	if o == nil || o.BuildpackLanguage == nil {
		var ret string
		return ret
	}
	return *o.BuildpackLanguage
}

// GetBuildpackLanguageOk returns a tuple with the BuildpackLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetBuildpackLanguageOk() (*string, bool) {
	if o == nil || o.BuildpackLanguage == nil {
		return nil, false
	}
	return o.BuildpackLanguage, true
}

// HasBuildpackLanguage returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasBuildpackLanguage() bool {
	if o != nil && o.BuildpackLanguage != nil {
		return true
	}

	return false
}

// SetBuildpackLanguage gets a reference to the given string and assigns it to the BuildpackLanguage field.
func (o *ApplicationEditRequest) SetBuildpackLanguage(v string) {
	o.BuildpackLanguage = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetCpu() float32 {
	if o == nil || o.Cpu == nil {
		var ret float32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetCpuOk() (*float32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given float32 and assigns it to the Cpu field.
func (o *ApplicationEditRequest) SetCpu(v float32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetMemory() float32 {
	if o == nil || o.Memory == nil {
		var ret float32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetMemoryOk() (*float32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given float32 and assigns it to the Memory field.
func (o *ApplicationEditRequest) SetMemory(v float32) {
	o.Memory = &v
}

// GetMinRunningInstances returns the MinRunningInstances field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetMinRunningInstances() int32 {
	if o == nil || o.MinRunningInstances == nil {
		var ret int32
		return ret
	}
	return *o.MinRunningInstances
}

// GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetMinRunningInstancesOk() (*int32, bool) {
	if o == nil || o.MinRunningInstances == nil {
		return nil, false
	}
	return o.MinRunningInstances, true
}

// HasMinRunningInstances returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasMinRunningInstances() bool {
	if o != nil && o.MinRunningInstances != nil {
		return true
	}

	return false
}

// SetMinRunningInstances gets a reference to the given int32 and assigns it to the MinRunningInstances field.
func (o *ApplicationEditRequest) SetMinRunningInstances(v int32) {
	o.MinRunningInstances = &v
}

// GetMaxRunningInstances returns the MaxRunningInstances field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetMaxRunningInstances() int32 {
	if o == nil || o.MaxRunningInstances == nil {
		var ret int32
		return ret
	}
	return *o.MaxRunningInstances
}

// GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetMaxRunningInstancesOk() (*int32, bool) {
	if o == nil || o.MaxRunningInstances == nil {
		return nil, false
	}
	return o.MaxRunningInstances, true
}

// HasMaxRunningInstances returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasMaxRunningInstances() bool {
	if o != nil && o.MaxRunningInstances != nil {
		return true
	}

	return false
}

// SetMaxRunningInstances gets a reference to the given int32 and assigns it to the MaxRunningInstances field.
func (o *ApplicationEditRequest) SetMaxRunningInstances(v int32) {
	o.MaxRunningInstances = &v
}

// GetHealthcheck returns the Healthcheck field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetHealthcheck() Healthcheck {
	if o == nil || o.Healthcheck == nil {
		var ret Healthcheck
		return ret
	}
	return *o.Healthcheck
}

// GetHealthcheckOk returns a tuple with the Healthcheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetHealthcheckOk() (*Healthcheck, bool) {
	if o == nil || o.Healthcheck == nil {
		return nil, false
	}
	return o.Healthcheck, true
}

// HasHealthcheck returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasHealthcheck() bool {
	if o != nil && o.Healthcheck != nil {
		return true
	}

	return false
}

// SetHealthcheck gets a reference to the given Healthcheck and assigns it to the Healthcheck field.
func (o *ApplicationEditRequest) SetHealthcheck(v Healthcheck) {
	o.Healthcheck = &v
}

// GetEnvPreview returns the EnvPreview field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetEnvPreview() bool {
	if o == nil || o.EnvPreview == nil {
		var ret bool
		return ret
	}
	return *o.EnvPreview
}

// GetEnvPreviewOk returns a tuple with the EnvPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetEnvPreviewOk() (*bool, bool) {
	if o == nil || o.EnvPreview == nil {
		return nil, false
	}
	return o.EnvPreview, true
}

// HasEnvPreview returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasEnvPreview() bool {
	if o != nil && o.EnvPreview != nil {
		return true
	}

	return false
}

// SetEnvPreview gets a reference to the given bool and assigns it to the EnvPreview field.
func (o *ApplicationEditRequest) SetEnvPreview(v bool) {
	o.EnvPreview = &v
}

// GetStickySession returns the StickySession field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetStickySession() bool {
	if o == nil || o.StickySession == nil {
		var ret bool
		return ret
	}
	return *o.StickySession
}

// GetStickySessionOk returns a tuple with the StickySession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetStickySessionOk() (*bool, bool) {
	if o == nil || o.StickySession == nil {
		return nil, false
	}
	return o.StickySession, true
}

// HasStickySession returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasStickySession() bool {
	if o != nil && o.StickySession != nil {
		return true
	}

	return false
}

// SetStickySession gets a reference to the given bool and assigns it to the StickySession field.
func (o *ApplicationEditRequest) SetStickySession(v bool) {
	o.StickySession = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetStorage() []ApplicationStorageResponseStorage {
	if o == nil || o.Storage == nil {
		var ret []ApplicationStorageResponseStorage
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetStorageOk() (*[]ApplicationStorageResponseStorage, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasStorage() bool {
	if o != nil && o.Storage != nil {
		return true
	}

	return false
}

// SetStorage gets a reference to the given []ApplicationStorageResponseStorage and assigns it to the Storage field.
func (o *ApplicationEditRequest) SetStorage(v []ApplicationStorageResponseStorage) {
	o.Storage = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *ApplicationEditRequest) GetPorts() []ApplicationPortResponsePorts {
	if o == nil || o.Ports == nil {
		var ret []ApplicationPortResponsePorts
		return ret
	}
	return *o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationEditRequest) GetPortsOk() (*[]ApplicationPortResponsePorts, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *ApplicationEditRequest) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []ApplicationPortResponsePorts and assigns it to the Ports field.
func (o *ApplicationEditRequest) SetPorts(v []ApplicationPortResponsePorts) {
	o.Ports = &v
}

func (o ApplicationEditRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.GitRepository != nil {
		toSerialize["git_repository"] = o.GitRepository
	}
	if o.BuildMode != nil {
		toSerialize["build_mode"] = o.BuildMode
	}
	if o.DockerfilePath != nil {
		toSerialize["dockerfile_path"] = o.DockerfilePath
	}
	if o.BuildpackLanguage != nil {
		toSerialize["buildpack_language"] = o.BuildpackLanguage
	}
	if o.Cpu != nil {
		toSerialize["cpu"] = o.Cpu
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.MinRunningInstances != nil {
		toSerialize["min_running_instances"] = o.MinRunningInstances
	}
	if o.MaxRunningInstances != nil {
		toSerialize["max_running_instances"] = o.MaxRunningInstances
	}
	if o.Healthcheck != nil {
		toSerialize["healthcheck"] = o.Healthcheck
	}
	if o.EnvPreview != nil {
		toSerialize["env_preview"] = o.EnvPreview
	}
	if o.StickySession != nil {
		toSerialize["sticky-session"] = o.StickySession
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationEditRequest struct {
	value *ApplicationEditRequest
	isSet bool
}

func (v NullableApplicationEditRequest) Get() *ApplicationEditRequest {
	return v.value
}

func (v *NullableApplicationEditRequest) Set(val *ApplicationEditRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationEditRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationEditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationEditRequest(val *ApplicationEditRequest) *NullableApplicationEditRequest {
	return &NullableApplicationEditRequest{value: val, isSet: true}
}

func (v NullableApplicationEditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationEditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
