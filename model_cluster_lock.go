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

// checks if the ClusterLock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterLock{}

// ClusterLock struct for ClusterLock
type ClusterLock struct {
	Reason               string    `json:"reason"`
	TtlInDays            *int32    `json:"ttl_in_days,omitempty"`
	ClusterId            string    `json:"cluster_id"`
	LockedAt             time.Time `json:"locked_at"`
	OwnerName            string    `json:"owner_name"`
	AdditionalProperties map[string]interface{}
}

type _ClusterLock ClusterLock

// NewClusterLock instantiates a new ClusterLock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterLock(reason string, clusterId string, lockedAt time.Time, ownerName string) *ClusterLock {
	this := ClusterLock{}
	this.Reason = reason
	this.ClusterId = clusterId
	this.LockedAt = lockedAt
	this.OwnerName = ownerName
	return &this
}

// NewClusterLockWithDefaults instantiates a new ClusterLock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterLockWithDefaults() *ClusterLock {
	this := ClusterLock{}
	return &this
}

// GetReason returns the Reason field value
func (o *ClusterLock) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *ClusterLock) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *ClusterLock) SetReason(v string) {
	o.Reason = v
}

// GetTtlInDays returns the TtlInDays field value if set, zero value otherwise.
func (o *ClusterLock) GetTtlInDays() int32 {
	if o == nil || IsNil(o.TtlInDays) {
		var ret int32
		return ret
	}
	return *o.TtlInDays
}

// GetTtlInDaysOk returns a tuple with the TtlInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLock) GetTtlInDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.TtlInDays) {
		return nil, false
	}
	return o.TtlInDays, true
}

// HasTtlInDays returns a boolean if a field has been set.
func (o *ClusterLock) HasTtlInDays() bool {
	if o != nil && !IsNil(o.TtlInDays) {
		return true
	}

	return false
}

// SetTtlInDays gets a reference to the given int32 and assigns it to the TtlInDays field.
func (o *ClusterLock) SetTtlInDays(v int32) {
	o.TtlInDays = &v
}

// GetClusterId returns the ClusterId field value
func (o *ClusterLock) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *ClusterLock) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *ClusterLock) SetClusterId(v string) {
	o.ClusterId = v
}

// GetLockedAt returns the LockedAt field value
func (o *ClusterLock) GetLockedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LockedAt
}

// GetLockedAtOk returns a tuple with the LockedAt field value
// and a boolean to check if the value has been set.
func (o *ClusterLock) GetLockedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockedAt, true
}

// SetLockedAt sets field value
func (o *ClusterLock) SetLockedAt(v time.Time) {
	o.LockedAt = v
}

// GetOwnerName returns the OwnerName field value
func (o *ClusterLock) GetOwnerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value
// and a boolean to check if the value has been set.
func (o *ClusterLock) GetOwnerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerName, true
}

// SetOwnerName sets field value
func (o *ClusterLock) SetOwnerName(v string) {
	o.OwnerName = v
}

func (o ClusterLock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterLock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reason"] = o.Reason
	if !IsNil(o.TtlInDays) {
		toSerialize["ttl_in_days"] = o.TtlInDays
	}
	toSerialize["cluster_id"] = o.ClusterId
	toSerialize["locked_at"] = o.LockedAt
	toSerialize["owner_name"] = o.OwnerName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClusterLock) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reason",
		"cluster_id",
		"locked_at",
		"owner_name",
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

	varClusterLock := _ClusterLock{}

	err = json.Unmarshal(data, &varClusterLock)

	if err != nil {
		return err
	}

	*o = ClusterLock(varClusterLock)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reason")
		delete(additionalProperties, "ttl_in_days")
		delete(additionalProperties, "cluster_id")
		delete(additionalProperties, "locked_at")
		delete(additionalProperties, "owner_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClusterLock struct {
	value *ClusterLock
	isSet bool
}

func (v NullableClusterLock) Get() *ClusterLock {
	return v.value
}

func (v *NullableClusterLock) Set(val *ClusterLock) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterLock) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterLock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterLock(val *ClusterLock) *NullableClusterLock {
	return &NullableClusterLock{value: val, isSet: true}
}

func (v NullableClusterLock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterLock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
