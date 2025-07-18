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

// checks if the DatabaseEditRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseEditRequest{}

// DatabaseEditRequest struct for DatabaseEditRequest
type DatabaseEditRequest struct {
	// name is case-insensitive
	Name *string `json:"name,omitempty"`
	// give a description to this database
	Description   *string                    `json:"description,omitempty"`
	Version       *string                    `json:"version,omitempty"`
	Accessibility *DatabaseAccessibilityEnum `json:"accessibility,omitempty"`
	// unit is millicores (m). 1000m = 1 cpu. This field will be ignored for managed DB (instance type will be used instead).
	Cpu *int32 `json:"cpu,omitempty"`
	// unit is MB. 1024 MB = 1GB This field will be ignored for managed DB (instance type will be used instead). Default value is linked to the database type: - MANAGED: 100 - CONTAINER   - POSTGRES: 100   - REDIS: 100   - MYSQL: 512   - MONGODB: 256
	Memory *int32 `json:"memory,omitempty"`
	// unit is GB
	Storage *int32 `json:"storage,omitempty"`
	// Database instance type to be used for this database. The list of values can be retrieved via the endpoint /{CloudProvider}/managedDatabase/instanceType/{region}/{dbType}. This field SHOULD NOT be set for container DB.
	InstanceType      *string                    `json:"instance_type,omitempty"`
	AnnotationsGroups []ServiceAnnotationRequest `json:"annotations_groups,omitempty"`
	LabelsGroups      []ServiceLabelRequest      `json:"labels_groups,omitempty"`
	// Icon URI representing the database.
	IconUri              *string `json:"icon_uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DatabaseEditRequest DatabaseEditRequest

// NewDatabaseEditRequest instantiates a new DatabaseEditRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseEditRequest() *DatabaseEditRequest {
	this := DatabaseEditRequest{}
	var accessibility DatabaseAccessibilityEnum = DATABASEACCESSIBILITYENUM_PRIVATE
	this.Accessibility = &accessibility
	var cpu int32 = 250
	this.Cpu = &cpu
	return &this
}

// NewDatabaseEditRequestWithDefaults instantiates a new DatabaseEditRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseEditRequestWithDefaults() *DatabaseEditRequest {
	this := DatabaseEditRequest{}
	var accessibility DatabaseAccessibilityEnum = DATABASEACCESSIBILITYENUM_PRIVATE
	this.Accessibility = &accessibility
	var cpu int32 = 250
	this.Cpu = &cpu
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DatabaseEditRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseEditRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DatabaseEditRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DatabaseEditRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DatabaseEditRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseEditRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DatabaseEditRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DatabaseEditRequest) SetDescription(v string) {
	o.Description = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DatabaseEditRequest) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseEditRequest) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DatabaseEditRequest) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DatabaseEditRequest) SetVersion(v string) {
	o.Version = &v
}

// GetAccessibility returns the Accessibility field value if set, zero value otherwise.
func (o *DatabaseEditRequest) GetAccessibility() DatabaseAccessibilityEnum {
	if o == nil || IsNil(o.Accessibility) {
		var ret DatabaseAccessibilityEnum
		return ret
	}
	return *o.Accessibility
}

// GetAccessibilityOk returns a tuple with the Accessibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseEditRequest) GetAccessibilityOk() (*DatabaseAccessibilityEnum, bool) {
	if o == nil || IsNil(o.Accessibility) {
		return nil, false
	}
	return o.Accessibility, true
}

// HasAccessibility returns a boolean if a field has been set.
func (o *DatabaseEditRequest) HasAccessibility() bool {
	if o != nil && !IsNil(o.Accessibility) {
		return true
	}

	return false
}

// SetAccessibility gets a reference to the given DatabaseAccessibilityEnum and assigns it to the Accessibility field.
func (o *DatabaseEditRequest) SetAccessibility(v DatabaseAccessibilityEnum) {
	o.Accessibility = &v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *DatabaseEditRequest) GetCpu() int32 {
	if o == nil || IsNil(o.Cpu) {
		var ret int32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseEditRequest) GetCpuOk() (*int32, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *DatabaseEditRequest) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int32 and assigns it to the Cpu field.
func (o *DatabaseEditRequest) SetCpu(v int32) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *DatabaseEditRequest) GetMemory() int32 {
	if o == nil || IsNil(o.Memory) {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseEditRequest) GetMemoryOk() (*int32, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *DatabaseEditRequest) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *DatabaseEditRequest) SetMemory(v int32) {
	o.Memory = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *DatabaseEditRequest) GetStorage() int32 {
	if o == nil || IsNil(o.Storage) {
		var ret int32
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseEditRequest) GetStorageOk() (*int32, bool) {
	if o == nil || IsNil(o.Storage) {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *DatabaseEditRequest) HasStorage() bool {
	if o != nil && !IsNil(o.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given int32 and assigns it to the Storage field.
func (o *DatabaseEditRequest) SetStorage(v int32) {
	o.Storage = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *DatabaseEditRequest) GetInstanceType() string {
	if o == nil || IsNil(o.InstanceType) {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseEditRequest) GetInstanceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceType) {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *DatabaseEditRequest) HasInstanceType() bool {
	if o != nil && !IsNil(o.InstanceType) {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *DatabaseEditRequest) SetInstanceType(v string) {
	o.InstanceType = &v
}

// GetAnnotationsGroups returns the AnnotationsGroups field value if set, zero value otherwise.
func (o *DatabaseEditRequest) GetAnnotationsGroups() []ServiceAnnotationRequest {
	if o == nil || IsNil(o.AnnotationsGroups) {
		var ret []ServiceAnnotationRequest
		return ret
	}
	return o.AnnotationsGroups
}

// GetAnnotationsGroupsOk returns a tuple with the AnnotationsGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseEditRequest) GetAnnotationsGroupsOk() ([]ServiceAnnotationRequest, bool) {
	if o == nil || IsNil(o.AnnotationsGroups) {
		return nil, false
	}
	return o.AnnotationsGroups, true
}

// HasAnnotationsGroups returns a boolean if a field has been set.
func (o *DatabaseEditRequest) HasAnnotationsGroups() bool {
	if o != nil && !IsNil(o.AnnotationsGroups) {
		return true
	}

	return false
}

// SetAnnotationsGroups gets a reference to the given []ServiceAnnotationRequest and assigns it to the AnnotationsGroups field.
func (o *DatabaseEditRequest) SetAnnotationsGroups(v []ServiceAnnotationRequest) {
	o.AnnotationsGroups = v
}

// GetLabelsGroups returns the LabelsGroups field value if set, zero value otherwise.
func (o *DatabaseEditRequest) GetLabelsGroups() []ServiceLabelRequest {
	if o == nil || IsNil(o.LabelsGroups) {
		var ret []ServiceLabelRequest
		return ret
	}
	return o.LabelsGroups
}

// GetLabelsGroupsOk returns a tuple with the LabelsGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseEditRequest) GetLabelsGroupsOk() ([]ServiceLabelRequest, bool) {
	if o == nil || IsNil(o.LabelsGroups) {
		return nil, false
	}
	return o.LabelsGroups, true
}

// HasLabelsGroups returns a boolean if a field has been set.
func (o *DatabaseEditRequest) HasLabelsGroups() bool {
	if o != nil && !IsNil(o.LabelsGroups) {
		return true
	}

	return false
}

// SetLabelsGroups gets a reference to the given []ServiceLabelRequest and assigns it to the LabelsGroups field.
func (o *DatabaseEditRequest) SetLabelsGroups(v []ServiceLabelRequest) {
	o.LabelsGroups = v
}

// GetIconUri returns the IconUri field value if set, zero value otherwise.
func (o *DatabaseEditRequest) GetIconUri() string {
	if o == nil || IsNil(o.IconUri) {
		var ret string
		return ret
	}
	return *o.IconUri
}

// GetIconUriOk returns a tuple with the IconUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseEditRequest) GetIconUriOk() (*string, bool) {
	if o == nil || IsNil(o.IconUri) {
		return nil, false
	}
	return o.IconUri, true
}

// HasIconUri returns a boolean if a field has been set.
func (o *DatabaseEditRequest) HasIconUri() bool {
	if o != nil && !IsNil(o.IconUri) {
		return true
	}

	return false
}

// SetIconUri gets a reference to the given string and assigns it to the IconUri field.
func (o *DatabaseEditRequest) SetIconUri(v string) {
	o.IconUri = &v
}

func (o DatabaseEditRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseEditRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Accessibility) {
		toSerialize["accessibility"] = o.Accessibility
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !IsNil(o.Storage) {
		toSerialize["storage"] = o.Storage
	}
	if !IsNil(o.InstanceType) {
		toSerialize["instance_type"] = o.InstanceType
	}
	if !IsNil(o.AnnotationsGroups) {
		toSerialize["annotations_groups"] = o.AnnotationsGroups
	}
	if !IsNil(o.LabelsGroups) {
		toSerialize["labels_groups"] = o.LabelsGroups
	}
	if !IsNil(o.IconUri) {
		toSerialize["icon_uri"] = o.IconUri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DatabaseEditRequest) UnmarshalJSON(data []byte) (err error) {
	varDatabaseEditRequest := _DatabaseEditRequest{}

	err = json.Unmarshal(data, &varDatabaseEditRequest)

	if err != nil {
		return err
	}

	*o = DatabaseEditRequest(varDatabaseEditRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "version")
		delete(additionalProperties, "accessibility")
		delete(additionalProperties, "cpu")
		delete(additionalProperties, "memory")
		delete(additionalProperties, "storage")
		delete(additionalProperties, "instance_type")
		delete(additionalProperties, "annotations_groups")
		delete(additionalProperties, "labels_groups")
		delete(additionalProperties, "icon_uri")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDatabaseEditRequest struct {
	value *DatabaseEditRequest
	isSet bool
}

func (v NullableDatabaseEditRequest) Get() *DatabaseEditRequest {
	return v.value
}

func (v *NullableDatabaseEditRequest) Set(val *DatabaseEditRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseEditRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseEditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseEditRequest(val *DatabaseEditRequest) *NullableDatabaseEditRequest {
	return &NullableDatabaseEditRequest{value: val, isSet: true}
}

func (v NullableDatabaseEditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseEditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
