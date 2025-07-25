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
	"time"
)

// checks if the CronJobResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CronJobResponse{}

// CronJobResponse struct for CronJobResponse
type CronJobResponse struct {
	Id          string          `json:"id"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   *time.Time      `json:"updated_at,omitempty"`
	Environment ReferenceObject `json:"environment"`
	// Maximum cpu that can be allocated to the job based on organization cluster configuration. unit is millicores (m). 1000m = 1 cpu
	MaximumCpu int32 `json:"maximum_cpu"`
	// Maximum memory that can be allocated to the job based on organization cluster configuration. unit is MB. 1024 MB = 1GB
	MaximumMemory int32 `json:"maximum_memory"`
	// name is case insensitive
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu
	Cpu int32 `json:"cpu"`
	// unit is MB. 1024 MB = 1GB
	Memory int32 `json:"memory"`
	// Maximum number of restart allowed before the job is considered as failed 0 means that no restart/crash of the job is allowed
	MaxNbRestart *int32 `json:"max_nb_restart,omitempty"`
	// Maximum number of seconds allowed for the job to run before killing it and mark it as failed
	MaxDurationSeconds *int32 `json:"max_duration_seconds,omitempty"`
	// Indicates if the 'environment preview option' is enabled for this container.   If enabled, a preview environment will be automatically cloned when `/preview` endpoint is called.   If not specified, it takes the value of the `auto_preview` property from the associated environment.
	AutoPreview bool `json:"auto_preview"`
	// Port where to run readiness and liveliness probes checks. The port will not be exposed externally
	Port         NullableInt32              `json:"port,omitempty"`
	Source       BaseJobResponseAllOfSource `json:"source"`
	Healthchecks Healthcheck                `json:"healthchecks"`
	// Specify if the job will be automatically updated after receiving a new image tag or a new commit according to the source type.  The new image tag shall be communicated via the \"Auto Deploy job\" endpoint https://api-doc.qovery.com/#tag/Jobs/operation/autoDeployJobEnvironments
	AutoDeploy *bool `json:"auto_deploy,omitempty"`
	// Icon URI representing the job.
	IconUri              string                                 `json:"icon_uri"`
	ServiceType          ServiceTypeEnum                        `json:"service_type"`
	JobType              JobTypeEnum                            `json:"job_type"`
	Schedule             CronJobResponseAllOfSchedule           `json:"schedule"`
	AnnotationsGroups    []OrganizationAnnotationsGroupResponse `json:"annotations_groups,omitempty"`
	LabelsGroups         []OrganizationLabelsGroupResponse      `json:"labels_groups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CronJobResponse CronJobResponse

// NewCronJobResponse instantiates a new CronJobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCronJobResponse(id string, createdAt time.Time, environment ReferenceObject, maximumCpu int32, maximumMemory int32, name string, cpu int32, memory int32, autoPreview bool, source BaseJobResponseAllOfSource, healthchecks Healthcheck, iconUri string, serviceType ServiceTypeEnum, jobType JobTypeEnum, schedule CronJobResponseAllOfSchedule) *CronJobResponse {
	this := CronJobResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Environment = environment
	this.MaximumCpu = maximumCpu
	this.MaximumMemory = maximumMemory
	this.Name = name
	this.Cpu = cpu
	this.Memory = memory
	this.AutoPreview = autoPreview
	this.Source = source
	this.Healthchecks = healthchecks
	this.IconUri = iconUri
	this.ServiceType = serviceType
	this.JobType = jobType
	this.Schedule = schedule
	return &this
}

// NewCronJobResponseWithDefaults instantiates a new CronJobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCronJobResponseWithDefaults() *CronJobResponse {
	this := CronJobResponse{}
	return &this
}

// GetId returns the Id field value
func (o *CronJobResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CronJobResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CronJobResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CronJobResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CronJobResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CronJobResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CronJobResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetEnvironment returns the Environment field value
func (o *CronJobResponse) GetEnvironment() ReferenceObject {
	if o == nil {
		var ret ReferenceObject
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetEnvironmentOk() (*ReferenceObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *CronJobResponse) SetEnvironment(v ReferenceObject) {
	o.Environment = v
}

// GetMaximumCpu returns the MaximumCpu field value
func (o *CronJobResponse) GetMaximumCpu() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaximumCpu
}

// GetMaximumCpuOk returns a tuple with the MaximumCpu field value
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetMaximumCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaximumCpu, true
}

// SetMaximumCpu sets field value
func (o *CronJobResponse) SetMaximumCpu(v int32) {
	o.MaximumCpu = v
}

// GetMaximumMemory returns the MaximumMemory field value
func (o *CronJobResponse) GetMaximumMemory() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaximumMemory
}

// GetMaximumMemoryOk returns a tuple with the MaximumMemory field value
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetMaximumMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaximumMemory, true
}

// SetMaximumMemory sets field value
func (o *CronJobResponse) SetMaximumMemory(v int32) {
	o.MaximumMemory = v
}

// GetName returns the Name field value
func (o *CronJobResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CronJobResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CronJobResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CronJobResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CronJobResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCpu returns the Cpu field value
func (o *CronJobResponse) GetCpu() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value
func (o *CronJobResponse) SetCpu(v int32) {
	o.Cpu = v
}

// GetMemory returns the Memory field value
func (o *CronJobResponse) GetMemory() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memory, true
}

// SetMemory sets field value
func (o *CronJobResponse) SetMemory(v int32) {
	o.Memory = v
}

// GetMaxNbRestart returns the MaxNbRestart field value if set, zero value otherwise.
func (o *CronJobResponse) GetMaxNbRestart() int32 {
	if o == nil || IsNil(o.MaxNbRestart) {
		var ret int32
		return ret
	}
	return *o.MaxNbRestart
}

// GetMaxNbRestartOk returns a tuple with the MaxNbRestart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetMaxNbRestartOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxNbRestart) {
		return nil, false
	}
	return o.MaxNbRestart, true
}

// HasMaxNbRestart returns a boolean if a field has been set.
func (o *CronJobResponse) HasMaxNbRestart() bool {
	if o != nil && !IsNil(o.MaxNbRestart) {
		return true
	}

	return false
}

// SetMaxNbRestart gets a reference to the given int32 and assigns it to the MaxNbRestart field.
func (o *CronJobResponse) SetMaxNbRestart(v int32) {
	o.MaxNbRestart = &v
}

// GetMaxDurationSeconds returns the MaxDurationSeconds field value if set, zero value otherwise.
func (o *CronJobResponse) GetMaxDurationSeconds() int32 {
	if o == nil || IsNil(o.MaxDurationSeconds) {
		var ret int32
		return ret
	}
	return *o.MaxDurationSeconds
}

// GetMaxDurationSecondsOk returns a tuple with the MaxDurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetMaxDurationSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxDurationSeconds) {
		return nil, false
	}
	return o.MaxDurationSeconds, true
}

// HasMaxDurationSeconds returns a boolean if a field has been set.
func (o *CronJobResponse) HasMaxDurationSeconds() bool {
	if o != nil && !IsNil(o.MaxDurationSeconds) {
		return true
	}

	return false
}

// SetMaxDurationSeconds gets a reference to the given int32 and assigns it to the MaxDurationSeconds field.
func (o *CronJobResponse) SetMaxDurationSeconds(v int32) {
	o.MaxDurationSeconds = &v
}

// GetAutoPreview returns the AutoPreview field value
func (o *CronJobResponse) GetAutoPreview() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoPreview
}

// GetAutoPreviewOk returns a tuple with the AutoPreview field value
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetAutoPreviewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoPreview, true
}

// SetAutoPreview sets field value
func (o *CronJobResponse) SetAutoPreview(v bool) {
	o.AutoPreview = v
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CronJobResponse) GetPort() int32 {
	if o == nil || IsNil(o.Port.Get()) {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CronJobResponse) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *CronJobResponse) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *CronJobResponse) SetPort(v int32) {
	o.Port.Set(&v)
}

// SetPortNil sets the value for Port to be an explicit nil
func (o *CronJobResponse) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *CronJobResponse) UnsetPort() {
	o.Port.Unset()
}

// GetSource returns the Source field value
func (o *CronJobResponse) GetSource() BaseJobResponseAllOfSource {
	if o == nil {
		var ret BaseJobResponseAllOfSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetSourceOk() (*BaseJobResponseAllOfSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *CronJobResponse) SetSource(v BaseJobResponseAllOfSource) {
	o.Source = v
}

// GetHealthchecks returns the Healthchecks field value
func (o *CronJobResponse) GetHealthchecks() Healthcheck {
	if o == nil {
		var ret Healthcheck
		return ret
	}

	return o.Healthchecks
}

// GetHealthchecksOk returns a tuple with the Healthchecks field value
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetHealthchecksOk() (*Healthcheck, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Healthchecks, true
}

// SetHealthchecks sets field value
func (o *CronJobResponse) SetHealthchecks(v Healthcheck) {
	o.Healthchecks = v
}

// GetAutoDeploy returns the AutoDeploy field value if set, zero value otherwise.
func (o *CronJobResponse) GetAutoDeploy() bool {
	if o == nil || IsNil(o.AutoDeploy) {
		var ret bool
		return ret
	}
	return *o.AutoDeploy
}

// GetAutoDeployOk returns a tuple with the AutoDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetAutoDeployOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoDeploy) {
		return nil, false
	}
	return o.AutoDeploy, true
}

// HasAutoDeploy returns a boolean if a field has been set.
func (o *CronJobResponse) HasAutoDeploy() bool {
	if o != nil && !IsNil(o.AutoDeploy) {
		return true
	}

	return false
}

// SetAutoDeploy gets a reference to the given bool and assigns it to the AutoDeploy field.
func (o *CronJobResponse) SetAutoDeploy(v bool) {
	o.AutoDeploy = &v
}

// GetIconUri returns the IconUri field value
func (o *CronJobResponse) GetIconUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IconUri
}

// GetIconUriOk returns a tuple with the IconUri field value
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetIconUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IconUri, true
}

// SetIconUri sets field value
func (o *CronJobResponse) SetIconUri(v string) {
	o.IconUri = v
}

// GetServiceType returns the ServiceType field value
func (o *CronJobResponse) GetServiceType() ServiceTypeEnum {
	if o == nil {
		var ret ServiceTypeEnum
		return ret
	}

	return o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetServiceTypeOk() (*ServiceTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceType, true
}

// SetServiceType sets field value
func (o *CronJobResponse) SetServiceType(v ServiceTypeEnum) {
	o.ServiceType = v
}

// GetJobType returns the JobType field value
func (o *CronJobResponse) GetJobType() JobTypeEnum {
	if o == nil {
		var ret JobTypeEnum
		return ret
	}

	return o.JobType
}

// GetJobTypeOk returns a tuple with the JobType field value
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetJobTypeOk() (*JobTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobType, true
}

// SetJobType sets field value
func (o *CronJobResponse) SetJobType(v JobTypeEnum) {
	o.JobType = v
}

// GetSchedule returns the Schedule field value
func (o *CronJobResponse) GetSchedule() CronJobResponseAllOfSchedule {
	if o == nil {
		var ret CronJobResponseAllOfSchedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetScheduleOk() (*CronJobResponseAllOfSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *CronJobResponse) SetSchedule(v CronJobResponseAllOfSchedule) {
	o.Schedule = v
}

// GetAnnotationsGroups returns the AnnotationsGroups field value if set, zero value otherwise.
func (o *CronJobResponse) GetAnnotationsGroups() []OrganizationAnnotationsGroupResponse {
	if o == nil || IsNil(o.AnnotationsGroups) {
		var ret []OrganizationAnnotationsGroupResponse
		return ret
	}
	return o.AnnotationsGroups
}

// GetAnnotationsGroupsOk returns a tuple with the AnnotationsGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetAnnotationsGroupsOk() ([]OrganizationAnnotationsGroupResponse, bool) {
	if o == nil || IsNil(o.AnnotationsGroups) {
		return nil, false
	}
	return o.AnnotationsGroups, true
}

// HasAnnotationsGroups returns a boolean if a field has been set.
func (o *CronJobResponse) HasAnnotationsGroups() bool {
	if o != nil && !IsNil(o.AnnotationsGroups) {
		return true
	}

	return false
}

// SetAnnotationsGroups gets a reference to the given []OrganizationAnnotationsGroupResponse and assigns it to the AnnotationsGroups field.
func (o *CronJobResponse) SetAnnotationsGroups(v []OrganizationAnnotationsGroupResponse) {
	o.AnnotationsGroups = v
}

// GetLabelsGroups returns the LabelsGroups field value if set, zero value otherwise.
func (o *CronJobResponse) GetLabelsGroups() []OrganizationLabelsGroupResponse {
	if o == nil || IsNil(o.LabelsGroups) {
		var ret []OrganizationLabelsGroupResponse
		return ret
	}
	return o.LabelsGroups
}

// GetLabelsGroupsOk returns a tuple with the LabelsGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronJobResponse) GetLabelsGroupsOk() ([]OrganizationLabelsGroupResponse, bool) {
	if o == nil || IsNil(o.LabelsGroups) {
		return nil, false
	}
	return o.LabelsGroups, true
}

// HasLabelsGroups returns a boolean if a field has been set.
func (o *CronJobResponse) HasLabelsGroups() bool {
	if o != nil && !IsNil(o.LabelsGroups) {
		return true
	}

	return false
}

// SetLabelsGroups gets a reference to the given []OrganizationLabelsGroupResponse and assigns it to the LabelsGroups field.
func (o *CronJobResponse) SetLabelsGroups(v []OrganizationLabelsGroupResponse) {
	o.LabelsGroups = v
}

func (o CronJobResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CronJobResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	toSerialize["environment"] = o.Environment
	toSerialize["maximum_cpu"] = o.MaximumCpu
	toSerialize["maximum_memory"] = o.MaximumMemory
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["cpu"] = o.Cpu
	toSerialize["memory"] = o.Memory
	if !IsNil(o.MaxNbRestart) {
		toSerialize["max_nb_restart"] = o.MaxNbRestart
	}
	if !IsNil(o.MaxDurationSeconds) {
		toSerialize["max_duration_seconds"] = o.MaxDurationSeconds
	}
	toSerialize["auto_preview"] = o.AutoPreview
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	toSerialize["source"] = o.Source
	toSerialize["healthchecks"] = o.Healthchecks
	if !IsNil(o.AutoDeploy) {
		toSerialize["auto_deploy"] = o.AutoDeploy
	}
	toSerialize["icon_uri"] = o.IconUri
	toSerialize["service_type"] = o.ServiceType
	toSerialize["job_type"] = o.JobType
	toSerialize["schedule"] = o.Schedule
	if !IsNil(o.AnnotationsGroups) {
		toSerialize["annotations_groups"] = o.AnnotationsGroups
	}
	if !IsNil(o.LabelsGroups) {
		toSerialize["labels_groups"] = o.LabelsGroups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CronJobResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"environment",
		"maximum_cpu",
		"maximum_memory",
		"name",
		"cpu",
		"memory",
		"auto_preview",
		"source",
		"healthchecks",
		"icon_uri",
		"service_type",
		"job_type",
		"schedule",
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

	varCronJobResponse := _CronJobResponse{}

	err = json.Unmarshal(data, &varCronJobResponse)

	if err != nil {
		return err
	}

	*o = CronJobResponse(varCronJobResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "environment")
		delete(additionalProperties, "maximum_cpu")
		delete(additionalProperties, "maximum_memory")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "cpu")
		delete(additionalProperties, "memory")
		delete(additionalProperties, "max_nb_restart")
		delete(additionalProperties, "max_duration_seconds")
		delete(additionalProperties, "auto_preview")
		delete(additionalProperties, "port")
		delete(additionalProperties, "source")
		delete(additionalProperties, "healthchecks")
		delete(additionalProperties, "auto_deploy")
		delete(additionalProperties, "icon_uri")
		delete(additionalProperties, "service_type")
		delete(additionalProperties, "job_type")
		delete(additionalProperties, "schedule")
		delete(additionalProperties, "annotations_groups")
		delete(additionalProperties, "labels_groups")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCronJobResponse struct {
	value *CronJobResponse
	isSet bool
}

func (v NullableCronJobResponse) Get() *CronJobResponse {
	return v.value
}

func (v *NullableCronJobResponse) Set(val *CronJobResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCronJobResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCronJobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCronJobResponse(val *CronJobResponse) *NullableCronJobResponse {
	return &NullableCronJobResponse{value: val, isSet: true}
}

func (v NullableCronJobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCronJobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
