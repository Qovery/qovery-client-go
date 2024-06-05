/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the EnvironmentLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentLog{}

// EnvironmentLog struct for EnvironmentLog
type EnvironmentLog struct {
	Id        string               `json:"id"`
	CreatedAt time.Time            `json:"created_at"`
	Scope     *EnvironmentLogScope `json:"scope,omitempty"`
	State     *StatusKindEnum      `json:"state,omitempty"`
	// Log message
	Message NullableString `json:"message"`
	// Only for errors. Helps Qovery team to investigate.
	ExecutionId *string `json:"execution_id,omitempty"`
	Hint        *string `json:"hint,omitempty"`
}

type _EnvironmentLog EnvironmentLog

// NewEnvironmentLog instantiates a new EnvironmentLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentLog(id string, createdAt time.Time, message NullableString) *EnvironmentLog {
	this := EnvironmentLog{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Message = message
	return &this
}

// NewEnvironmentLogWithDefaults instantiates a new EnvironmentLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentLogWithDefaults() *EnvironmentLog {
	this := EnvironmentLog{}
	return &this
}

// GetId returns the Id field value
func (o *EnvironmentLog) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnvironmentLog) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnvironmentLog) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EnvironmentLog) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EnvironmentLog) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EnvironmentLog) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *EnvironmentLog) GetScope() EnvironmentLogScope {
	if o == nil || IsNil(o.Scope) {
		var ret EnvironmentLogScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLog) GetScopeOk() (*EnvironmentLogScope, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *EnvironmentLog) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given EnvironmentLogScope and assigns it to the Scope field.
func (o *EnvironmentLog) SetScope(v EnvironmentLogScope) {
	o.Scope = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *EnvironmentLog) GetState() StatusKindEnum {
	if o == nil || IsNil(o.State) {
		var ret StatusKindEnum
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLog) GetStateOk() (*StatusKindEnum, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *EnvironmentLog) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given StatusKindEnum and assigns it to the State field.
func (o *EnvironmentLog) SetState(v StatusKindEnum) {
	o.State = &v
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EnvironmentLog) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}

	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnvironmentLog) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// SetMessage sets field value
func (o *EnvironmentLog) SetMessage(v string) {
	o.Message.Set(&v)
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *EnvironmentLog) GetExecutionId() string {
	if o == nil || IsNil(o.ExecutionId) {
		var ret string
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLog) GetExecutionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionId) {
		return nil, false
	}
	return o.ExecutionId, true
}

// HasExecutionId returns a boolean if a field has been set.
func (o *EnvironmentLog) HasExecutionId() bool {
	if o != nil && !IsNil(o.ExecutionId) {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given string and assigns it to the ExecutionId field.
func (o *EnvironmentLog) SetExecutionId(v string) {
	o.ExecutionId = &v
}

// GetHint returns the Hint field value if set, zero value otherwise.
func (o *EnvironmentLog) GetHint() string {
	if o == nil || IsNil(o.Hint) {
		var ret string
		return ret
	}
	return *o.Hint
}

// GetHintOk returns a tuple with the Hint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLog) GetHintOk() (*string, bool) {
	if o == nil || IsNil(o.Hint) {
		return nil, false
	}
	return o.Hint, true
}

// HasHint returns a boolean if a field has been set.
func (o *EnvironmentLog) HasHint() bool {
	if o != nil && !IsNil(o.Hint) {
		return true
	}

	return false
}

// SetHint gets a reference to the given string and assigns it to the Hint field.
func (o *EnvironmentLog) SetHint(v string) {
	o.Hint = &v
}

func (o EnvironmentLog) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	toSerialize["message"] = o.Message.Get()
	if !IsNil(o.ExecutionId) {
		toSerialize["execution_id"] = o.ExecutionId
	}
	if !IsNil(o.Hint) {
		toSerialize["hint"] = o.Hint
	}
	return toSerialize, nil
}

func (o *EnvironmentLog) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"message",
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

	varEnvironmentLog := _EnvironmentLog{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEnvironmentLog)

	if err != nil {
		return err
	}

	*o = EnvironmentLog(varEnvironmentLog)

	return err
}

type NullableEnvironmentLog struct {
	value *EnvironmentLog
	isSet bool
}

func (v NullableEnvironmentLog) Get() *EnvironmentLog {
	return v.value
}

func (v *NullableEnvironmentLog) Set(val *EnvironmentLog) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentLog) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentLog(val *EnvironmentLog) *NullableEnvironmentLog {
	return &NullableEnvironmentLog{value: val, isSet: true}
}

func (v NullableEnvironmentLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
