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

// LogicalDatabaseAllOf struct for LogicalDatabaseAllOf
type LogicalDatabaseAllOf struct {
	Database *ReferenceObject `json:"database,omitempty"`
}

// NewLogicalDatabaseAllOf instantiates a new LogicalDatabaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogicalDatabaseAllOf() *LogicalDatabaseAllOf {
	this := LogicalDatabaseAllOf{}
	return &this
}

// NewLogicalDatabaseAllOfWithDefaults instantiates a new LogicalDatabaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogicalDatabaseAllOfWithDefaults() *LogicalDatabaseAllOf {
	this := LogicalDatabaseAllOf{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *LogicalDatabaseAllOf) GetDatabase() ReferenceObject {
	if o == nil || o.Database == nil {
		var ret ReferenceObject
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogicalDatabaseAllOf) GetDatabaseOk() (*ReferenceObject, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *LogicalDatabaseAllOf) HasDatabase() bool {
	if o != nil && o.Database != nil {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given ReferenceObject and assigns it to the Database field.
func (o *LogicalDatabaseAllOf) SetDatabase(v ReferenceObject) {
	o.Database = &v
}

func (o LogicalDatabaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	return json.Marshal(toSerialize)
}

type NullableLogicalDatabaseAllOf struct {
	value *LogicalDatabaseAllOf
	isSet bool
}

func (v NullableLogicalDatabaseAllOf) Get() *LogicalDatabaseAllOf {
	return v.value
}

func (v *NullableLogicalDatabaseAllOf) Set(val *LogicalDatabaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLogicalDatabaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLogicalDatabaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogicalDatabaseAllOf(val *LogicalDatabaseAllOf) *NullableLogicalDatabaseAllOf {
	return &NullableLogicalDatabaseAllOf{value: val, isSet: true}
}

func (v NullableLogicalDatabaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogicalDatabaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}