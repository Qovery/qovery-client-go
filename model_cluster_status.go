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
	"time"
)

// checks if the ClusterStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterStatus{}

// ClusterStatus struct for ClusterStatus
type ClusterStatus struct {
	ClusterId               *string           `json:"cluster_id,omitempty"`
	Status                  *ClusterStateEnum `json:"status,omitempty"`
	IsDeployed              *bool             `json:"is_deployed,omitempty"`
	NextK8sAvailableVersion NullableString    `json:"next_k8s_available_version,omitempty"`
	LastExecutionId         *string           `json:"last_execution_id,omitempty"`
	ClusterLock             *ClusterLock      `json:"cluster_lock,omitempty"`
	LastDeploymentDate      *time.Time        `json:"last_deployment_date,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _ClusterStatus ClusterStatus

// NewClusterStatus instantiates a new ClusterStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterStatus() *ClusterStatus {
	this := ClusterStatus{}
	return &this
}

// NewClusterStatusWithDefaults instantiates a new ClusterStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterStatusWithDefaults() *ClusterStatus {
	this := ClusterStatus{}
	return &this
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise.
func (o *ClusterStatus) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}
	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *ClusterStatus) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *ClusterStatus) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClusterStatus) GetStatus() ClusterStateEnum {
	if o == nil || IsNil(o.Status) {
		var ret ClusterStateEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetStatusOk() (*ClusterStateEnum, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClusterStatus) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ClusterStateEnum and assigns it to the Status field.
func (o *ClusterStatus) SetStatus(v ClusterStateEnum) {
	o.Status = &v
}

// GetIsDeployed returns the IsDeployed field value if set, zero value otherwise.
func (o *ClusterStatus) GetIsDeployed() bool {
	if o == nil || IsNil(o.IsDeployed) {
		var ret bool
		return ret
	}
	return *o.IsDeployed
}

// GetIsDeployedOk returns a tuple with the IsDeployed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetIsDeployedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeployed) {
		return nil, false
	}
	return o.IsDeployed, true
}

// HasIsDeployed returns a boolean if a field has been set.
func (o *ClusterStatus) HasIsDeployed() bool {
	if o != nil && !IsNil(o.IsDeployed) {
		return true
	}

	return false
}

// SetIsDeployed gets a reference to the given bool and assigns it to the IsDeployed field.
func (o *ClusterStatus) SetIsDeployed(v bool) {
	o.IsDeployed = &v
}

// GetNextK8sAvailableVersion returns the NextK8sAvailableVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterStatus) GetNextK8sAvailableVersion() string {
	if o == nil || IsNil(o.NextK8sAvailableVersion.Get()) {
		var ret string
		return ret
	}
	return *o.NextK8sAvailableVersion.Get()
}

// GetNextK8sAvailableVersionOk returns a tuple with the NextK8sAvailableVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterStatus) GetNextK8sAvailableVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextK8sAvailableVersion.Get(), o.NextK8sAvailableVersion.IsSet()
}

// HasNextK8sAvailableVersion returns a boolean if a field has been set.
func (o *ClusterStatus) HasNextK8sAvailableVersion() bool {
	if o != nil && o.NextK8sAvailableVersion.IsSet() {
		return true
	}

	return false
}

// SetNextK8sAvailableVersion gets a reference to the given NullableString and assigns it to the NextK8sAvailableVersion field.
func (o *ClusterStatus) SetNextK8sAvailableVersion(v string) {
	o.NextK8sAvailableVersion.Set(&v)
}

// SetNextK8sAvailableVersionNil sets the value for NextK8sAvailableVersion to be an explicit nil
func (o *ClusterStatus) SetNextK8sAvailableVersionNil() {
	o.NextK8sAvailableVersion.Set(nil)
}

// UnsetNextK8sAvailableVersion ensures that no value is present for NextK8sAvailableVersion, not even an explicit nil
func (o *ClusterStatus) UnsetNextK8sAvailableVersion() {
	o.NextK8sAvailableVersion.Unset()
}

// GetLastExecutionId returns the LastExecutionId field value if set, zero value otherwise.
func (o *ClusterStatus) GetLastExecutionId() string {
	if o == nil || IsNil(o.LastExecutionId) {
		var ret string
		return ret
	}
	return *o.LastExecutionId
}

// GetLastExecutionIdOk returns a tuple with the LastExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetLastExecutionIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastExecutionId) {
		return nil, false
	}
	return o.LastExecutionId, true
}

// HasLastExecutionId returns a boolean if a field has been set.
func (o *ClusterStatus) HasLastExecutionId() bool {
	if o != nil && !IsNil(o.LastExecutionId) {
		return true
	}

	return false
}

// SetLastExecutionId gets a reference to the given string and assigns it to the LastExecutionId field.
func (o *ClusterStatus) SetLastExecutionId(v string) {
	o.LastExecutionId = &v
}

// GetClusterLock returns the ClusterLock field value if set, zero value otherwise.
func (o *ClusterStatus) GetClusterLock() ClusterLock {
	if o == nil || IsNil(o.ClusterLock) {
		var ret ClusterLock
		return ret
	}
	return *o.ClusterLock
}

// GetClusterLockOk returns a tuple with the ClusterLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetClusterLockOk() (*ClusterLock, bool) {
	if o == nil || IsNil(o.ClusterLock) {
		return nil, false
	}
	return o.ClusterLock, true
}

// HasClusterLock returns a boolean if a field has been set.
func (o *ClusterStatus) HasClusterLock() bool {
	if o != nil && !IsNil(o.ClusterLock) {
		return true
	}

	return false
}

// SetClusterLock gets a reference to the given ClusterLock and assigns it to the ClusterLock field.
func (o *ClusterStatus) SetClusterLock(v ClusterLock) {
	o.ClusterLock = &v
}

// GetLastDeploymentDate returns the LastDeploymentDate field value if set, zero value otherwise.
func (o *ClusterStatus) GetLastDeploymentDate() time.Time {
	if o == nil || IsNil(o.LastDeploymentDate) {
		var ret time.Time
		return ret
	}
	return *o.LastDeploymentDate
}

// GetLastDeploymentDateOk returns a tuple with the LastDeploymentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterStatus) GetLastDeploymentDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastDeploymentDate) {
		return nil, false
	}
	return o.LastDeploymentDate, true
}

// HasLastDeploymentDate returns a boolean if a field has been set.
func (o *ClusterStatus) HasLastDeploymentDate() bool {
	if o != nil && !IsNil(o.LastDeploymentDate) {
		return true
	}

	return false
}

// SetLastDeploymentDate gets a reference to the given time.Time and assigns it to the LastDeploymentDate field.
func (o *ClusterStatus) SetLastDeploymentDate(v time.Time) {
	o.LastDeploymentDate = &v
}

func (o ClusterStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClusterId) {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.IsDeployed) {
		toSerialize["is_deployed"] = o.IsDeployed
	}
	if o.NextK8sAvailableVersion.IsSet() {
		toSerialize["next_k8s_available_version"] = o.NextK8sAvailableVersion.Get()
	}
	if !IsNil(o.LastExecutionId) {
		toSerialize["last_execution_id"] = o.LastExecutionId
	}
	if !IsNil(o.ClusterLock) {
		toSerialize["cluster_lock"] = o.ClusterLock
	}
	if !IsNil(o.LastDeploymentDate) {
		toSerialize["last_deployment_date"] = o.LastDeploymentDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClusterStatus) UnmarshalJSON(data []byte) (err error) {
	varClusterStatus := _ClusterStatus{}

	err = json.Unmarshal(data, &varClusterStatus)

	if err != nil {
		return err
	}

	*o = ClusterStatus(varClusterStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cluster_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "is_deployed")
		delete(additionalProperties, "next_k8s_available_version")
		delete(additionalProperties, "last_execution_id")
		delete(additionalProperties, "cluster_lock")
		delete(additionalProperties, "last_deployment_date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClusterStatus struct {
	value *ClusterStatus
	isSet bool
}

func (v NullableClusterStatus) Get() *ClusterStatus {
	return v.value
}

func (v *NullableClusterStatus) Set(val *ClusterStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterStatus(val *ClusterStatus) *NullableClusterStatus {
	return &NullableClusterStatus{value: val, isSet: true}
}

func (v NullableClusterStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
