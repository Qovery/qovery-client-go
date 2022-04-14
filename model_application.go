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
	"time"
)

// Application struct for Application
type Application struct {
	Id            string                      `json:"id"`
	CreatedAt     time.Time                   `json:"created_at"`
	UpdatedAt     *time.Time                  `json:"updated_at,omitempty"`
	Storage       []ApplicationStorageStorage `json:"storage,omitempty"`
	Ports         []ApplicationPortPorts      `json:"ports,omitempty"`
	Environment   *ReferenceObject            `json:"environment,omitempty"`
	GitRepository *ApplicationGitRepository   `json:"git_repository,omitempty"`
	// Maximum cpu that can be allocated to the application based on organization cluster configuration. unit is millicores (m). 1000m = 1 cpu
	MaximumCpu *int32 `json:"maximum_cpu,omitempty"`
	// Maximum memory that can be allocated to the application based on organization cluster configuration. unit is MB. 1024 MB = 1GB
	MaximumMemory *int32 `json:"maximum_memory,omitempty"`
	// name is case insensitive
	Name *string `json:"name,omitempty"`
	// give a description to this application
	Description NullableString `json:"description,omitempty"`
	BuildMode   *BuildModeEnum `json:"build_mode,omitempty"`
	// The path of the associated Dockerfile. Only if you are using build_mode = DOCKER
	DockerfilePath    NullableString                `json:"dockerfile_path,omitempty"`
	BuildpackLanguage NullableBuildPackLanguageEnum `json:"buildpack_language,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu *int32 `json:"cpu,omitempty"`
	// unit is MB. 1024 MB = 1GB
	Memory *int32 `json:"memory,omitempty"`
	// Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no application running.
	MinRunningInstances *int32 `json:"min_running_instances,omitempty"`
	// Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.
	MaxRunningInstances *int32       `json:"max_running_instances,omitempty"`
	Healthcheck         *Healthcheck `json:"healthcheck,omitempty"`
	// Specify if the environment preview option is activated or not for this application. If activated, a preview environment will be automatically cloned at each pull request.
	AutoPreview *bool `json:"auto_preview,omitempty"`
}

// NewApplication instantiates a new Application object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplication(id string, createdAt time.Time) *Application {
	this := Application{}
	this.Id = id
	this.CreatedAt = createdAt
	var maximumCpu int32 = 250
	this.MaximumCpu = &maximumCpu
	var maximumMemory int32 = 256
	this.MaximumMemory = &maximumMemory
	var buildMode BuildModeEnum = BUILDPACKS
	this.BuildMode = &buildMode
	var cpu int32 = 250
	this.Cpu = &cpu
	var memory int32 = 256
	this.Memory = &memory
	var minRunningInstances int32 = 1
	this.MinRunningInstances = &minRunningInstances
	var maxRunningInstances int32 = 1
	this.MaxRunningInstances = &maxRunningInstances
	var autoPreview bool = true
	this.AutoPreview = &autoPreview
	return &this
}

// NewApplicationWithDefaults instantiates a new Application object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationWithDefaults() *Application {
	this := Application{}
	var maximumCpu int32 = 250
	this.MaximumCpu = &maximumCpu
	var maximumMemory int32 = 256
	this.MaximumMemory = &maximumMemory
	var buildMode BuildModeEnum = BUILDPACKS
	this.BuildMode = &buildMode
	var cpu int32 = 250
	this.Cpu = &cpu
	var memory int32 = 256
	this.Memory = &memory
	var minRunningInstances int32 = 1
	this.MinRunningInstances = &minRunningInstances
	var maxRunningInstances int32 = 1
	this.MaxRunningInstances = &maxRunningInstances
	var autoPreview bool = true
	this.AutoPreview = &autoPreview
	return &this
}

// GetId returns the Id field value
func (o *Application) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Application) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Application) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Application) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Application) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Application) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Application) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Application) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Application) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *Application) GetStorage() []ApplicationStorageStorage {
	if o == nil || o.Storage == nil {
		var ret []ApplicationStorageStorage
		return ret
	}
	return o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetStorageOk() ([]ApplicationStorageStorage, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *Application) HasStorage() bool {
	if o != nil && o.Storage != nil {
		return true
	}

	return false
}

// SetStorage gets a reference to the given []ApplicationStorageStorage and assigns it to the Storage field.
func (o *Application) SetStorage(v []ApplicationStorageStorage) {
	o.Storage = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *Application) GetPorts() []ApplicationPortPorts {
	if o == nil || o.Ports == nil {
		var ret []ApplicationPortPorts
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetPortsOk() ([]ApplicationPortPorts, bool) {
	if o == nil || o.Ports == nil {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *Application) HasPorts() bool {
	if o != nil && o.Ports != nil {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []ApplicationPortPorts and assigns it to the Ports field.
func (o *Application) SetPorts(v []ApplicationPortPorts) {
	o.Ports = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *Application) GetEnvironment() ReferenceObject {
	if o == nil || o.Environment == nil {
		var ret ReferenceObject
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetEnvironmentOk() (*ReferenceObject, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *Application) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ReferenceObject and assigns it to the Environment field.
func (o *Application) SetEnvironment(v ReferenceObject) {
	o.Environment = &v
}

// GetGitRepository returns the GitRepository field value if set, zero value otherwise.
func (o *Application) GetGitRepository() ApplicationGitRepository {
	if o == nil || o.GitRepository == nil {
		var ret ApplicationGitRepository
		return ret
	}
	return *o.GitRepository
}

// GetGitRepositoryOk returns a tuple with the GitRepository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetGitRepositoryOk() (*ApplicationGitRepository, bool) {
	if o == nil || o.GitRepository == nil {
		return nil, false
	}
	return o.GitRepository, true
}

// HasGitRepository returns a boolean if a field has been set.
func (o *Application) HasGitRepository() bool {
	if o != nil && o.GitRepository != nil {
		return true
	}

	return false
}

// SetGitRepository gets a reference to the given ApplicationGitRepository and assigns it to the GitRepository field.
func (o *Application) SetGitRepository(v ApplicationGitRepository) {
	o.GitRepository = &v
}

// GetMaximumCpu returns the MaximumCpu field value if set, zero value otherwise.
func (o *Application) GetMaximumCpu() int32 {
	if o == nil || o.MaximumCpu == nil {
		var ret int32
		return ret
	}
	return *o.MaximumCpu
}

// GetMaximumCpuOk returns a tuple with the MaximumCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetMaximumCpuOk() (*int32, bool) {
	if o == nil || o.MaximumCpu == nil {
		return nil, false
	}
	return o.MaximumCpu, true
}

// HasMaximumCpu returns a boolean if a field has been set.
func (o *Application) HasMaximumCpu() bool {
	if o != nil && o.MaximumCpu != nil {
		return true
	}

	return false
}

// SetMaximumCpu gets a reference to the given int32 and assigns it to the MaximumCpu field.
func (o *Application) SetMaximumCpu(v int32) {
	o.MaximumCpu = &v
}

// GetMaximumMemory returns the MaximumMemory field value if set, zero value otherwise.
func (o *Application) GetMaximumMemory() int32 {
	if o == nil || o.MaximumMemory == nil {
		var ret int32
		return ret
	}
	return *o.MaximumMemory
}

// GetMaximumMemoryOk returns a tuple with the MaximumMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetMaximumMemoryOk() (*int32, bool) {
	if o == nil || o.MaximumMemory == nil {
		return nil, false
	}
	return o.MaximumMemory, true
}

// HasMaximumMemory returns a boolean if a field has been set.
func (o *Application) HasMaximumMemory() bool {
	if o != nil && o.MaximumMemory != nil {
		return true
	}

	return false
}

// SetMaximumMemory gets a reference to the given int32 and assigns it to the MaximumMemory field.
func (o *Application) SetMaximumMemory(v int32) {
	o.MaximumMemory = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Application) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Application) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Application) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Application) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Application) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Application) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Application) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Application) UnsetDescription() {
	o.Description.Unset()
}

// GetBuildMode returns the BuildMode field value if set, zero value otherwise.
func (o *Application) GetBuildMode() BuildModeEnum {
	if o == nil || o.BuildMode == nil {
		var ret BuildModeEnum
		return ret
	}
	return *o.BuildMode
}

// GetBuildModeOk returns a tuple with the BuildMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetBuildModeOk() (*BuildModeEnum, bool) {
	if o == nil || o.BuildMode == nil {
		return nil, false
	}
	return o.BuildMode, true
}

// HasBuildMode returns a boolean if a field has been set.
func (o *Application) HasBuildMode() bool {
	if o != nil && o.BuildMode != nil {
		return true
	}

	return false
}

// SetBuildMode gets a reference to the given BuildModeEnum and assigns it to the BuildMode field.
func (o *Application) SetBuildMode(v BuildModeEnum) {
	o.BuildMode = &v
}

// GetDockerfilePath returns the DockerfilePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Application) GetDockerfilePath() string {
	if o == nil || o.DockerfilePath.Get() == nil {
		var ret string
		return ret
	}
	return *o.DockerfilePath.Get()
}

// GetDockerfilePathOk returns a tuple with the DockerfilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetDockerfilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DockerfilePath.Get(), o.DockerfilePath.IsSet()
}

// HasDockerfilePath returns a boolean if a field has been set.
func (o *Application) HasDockerfilePath() bool {
	if o != nil && o.DockerfilePath.IsSet() {
		return true
	}

	return false
}

// SetDockerfilePath gets a reference to the given NullableString and assigns it to the DockerfilePath field.
func (o *Application) SetDockerfilePath(v string) {
	o.DockerfilePath.Set(&v)
}

// SetDockerfilePathNil sets the value for DockerfilePath to be an explicit nil
func (o *Application) SetDockerfilePathNil() {
	o.DockerfilePath.Set(nil)
}

// UnsetDockerfilePath ensures that no value is present for DockerfilePath, not even an explicit nil
func (o *Application) UnsetDockerfilePath() {
	o.DockerfilePath.Unset()
}

// GetBuildpackLanguage returns the BuildpackLanguage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Application) GetBuildpackLanguage() BuildPackLanguageEnum {
	if o == nil || o.BuildpackLanguage.Get() == nil {
		var ret BuildPackLanguageEnum
		return ret
	}
	return *o.BuildpackLanguage.Get()
}

// GetBuildpackLanguageOk returns a tuple with the BuildpackLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Application) GetBuildpackLanguageOk() (*BuildPackLanguageEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuildpackLanguage.Get(), o.BuildpackLanguage.IsSet()
}

// HasBuildpackLanguage returns a boolean if a field has been set.
func (o *Application) HasBuildpackLanguage() bool {
	if o != nil && o.BuildpackLanguage.IsSet() {
		return true
	}

	return false
}

// SetBuildpackLanguage gets a reference to the given NullableBuildPackLanguageEnum and assigns it to the BuildpackLanguage field.
func (o *Application) SetBuildpackLanguage(v BuildPackLanguageEnum) {
	o.BuildpackLanguage.Set(&v)
}

// SetBuildpackLanguageNil sets the value for BuildpackLanguage to be an explicit nil
func (o *Application) SetBuildpackLanguageNil() {
	o.BuildpackLanguage.Set(nil)
}

// UnsetBuildpackLanguage ensures that no value is present for BuildpackLanguage, not even an explicit nil
func (o *Application) UnsetBuildpackLanguage() {
	o.BuildpackLanguage.Unset()
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *Application) GetCpu() int32 {
	if o == nil || o.Cpu == nil {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetCpuOk() (*int32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *Application) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *Application) SetCpu(v int32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *Application) GetMemory() int32 {
	if o == nil || o.Memory == nil {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetMemoryOk() (*int32, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *Application) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *Application) SetMemory(v int32) {
	o.Memory = &v
}

// GetMinRunningInstances returns the MinRunningInstances field value if set, zero value otherwise.
func (o *Application) GetMinRunningInstances() int32 {
	if o == nil || o.MinRunningInstances == nil {
		var ret int32
		return ret
	}
	return *o.MinRunningInstances
}

// GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetMinRunningInstancesOk() (*int32, bool) {
	if o == nil || o.MinRunningInstances == nil {
		return nil, false
	}
	return o.MinRunningInstances, true
}

// HasMinRunningInstances returns a boolean if a field has been set.
func (o *Application) HasMinRunningInstances() bool {
	if o != nil && o.MinRunningInstances != nil {
		return true
	}

	return false
}

// SetMinRunningInstances gets a reference to the given int32 and assigns it to the MinRunningInstances field.
func (o *Application) SetMinRunningInstances(v int32) {
	o.MinRunningInstances = &v
}

// GetMaxRunningInstances returns the MaxRunningInstances field value if set, zero value otherwise.
func (o *Application) GetMaxRunningInstances() int32 {
	if o == nil || o.MaxRunningInstances == nil {
		var ret int32
		return ret
	}
	return *o.MaxRunningInstances
}

// GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetMaxRunningInstancesOk() (*int32, bool) {
	if o == nil || o.MaxRunningInstances == nil {
		return nil, false
	}
	return o.MaxRunningInstances, true
}

// HasMaxRunningInstances returns a boolean if a field has been set.
func (o *Application) HasMaxRunningInstances() bool {
	if o != nil && o.MaxRunningInstances != nil {
		return true
	}

	return false
}

// SetMaxRunningInstances gets a reference to the given int32 and assigns it to the MaxRunningInstances field.
func (o *Application) SetMaxRunningInstances(v int32) {
	o.MaxRunningInstances = &v
}

// GetHealthcheck returns the Healthcheck field value if set, zero value otherwise.
func (o *Application) GetHealthcheck() Healthcheck {
	if o == nil || o.Healthcheck == nil {
		var ret Healthcheck
		return ret
	}
	return *o.Healthcheck
}

// GetHealthcheckOk returns a tuple with the Healthcheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetHealthcheckOk() (*Healthcheck, bool) {
	if o == nil || o.Healthcheck == nil {
		return nil, false
	}
	return o.Healthcheck, true
}

// HasHealthcheck returns a boolean if a field has been set.
func (o *Application) HasHealthcheck() bool {
	if o != nil && o.Healthcheck != nil {
		return true
	}

	return false
}

// SetHealthcheck gets a reference to the given Healthcheck and assigns it to the Healthcheck field.
func (o *Application) SetHealthcheck(v Healthcheck) {
	o.Healthcheck = &v
}

// GetAutoPreview returns the AutoPreview field value if set, zero value otherwise.
func (o *Application) GetAutoPreview() bool {
	if o == nil || o.AutoPreview == nil {
		var ret bool
		return ret
	}
	return *o.AutoPreview
}

// GetAutoPreviewOk returns a tuple with the AutoPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetAutoPreviewOk() (*bool, bool) {
	if o == nil || o.AutoPreview == nil {
		return nil, false
	}
	return o.AutoPreview, true
}

// HasAutoPreview returns a boolean if a field has been set.
func (o *Application) HasAutoPreview() bool {
	if o != nil && o.AutoPreview != nil {
		return true
	}

	return false
}

// SetAutoPreview gets a reference to the given bool and assigns it to the AutoPreview field.
func (o *Application) SetAutoPreview(v bool) {
	o.AutoPreview = &v
}

func (o Application) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.GitRepository != nil {
		toSerialize["git_repository"] = o.GitRepository
	}
	if o.MaximumCpu != nil {
		toSerialize["maximum_cpu"] = o.MaximumCpu
	}
	if o.MaximumMemory != nil {
		toSerialize["maximum_memory"] = o.MaximumMemory
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.BuildMode != nil {
		toSerialize["build_mode"] = o.BuildMode
	}
	if o.DockerfilePath.IsSet() {
		toSerialize["dockerfile_path"] = o.DockerfilePath.Get()
	}
	if o.BuildpackLanguage.IsSet() {
		toSerialize["buildpack_language"] = o.BuildpackLanguage.Get()
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
	if o.AutoPreview != nil {
		toSerialize["auto_preview"] = o.AutoPreview
	}
	return json.Marshal(toSerialize)
}

type NullableApplication struct {
	value *Application
	isSet bool
}

func (v NullableApplication) Get() *Application {
	return v.value
}

func (v *NullableApplication) Set(val *Application) {
	v.value = val
	v.isSet = true
}

func (v NullableApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplication(val *Application) *NullableApplication {
	return &NullableApplication{value: val, isSet: true}
}

func (v NullableApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}