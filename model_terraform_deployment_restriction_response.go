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

// checks if the TerraformDeploymentRestrictionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerraformDeploymentRestrictionResponse{}

// TerraformDeploymentRestrictionResponse struct for TerraformDeploymentRestrictionResponse
type TerraformDeploymentRestrictionResponse struct {
	Id        string                        `json:"id"`
	CreatedAt time.Time                     `json:"created_at"`
	UpdatedAt *time.Time                    `json:"updated_at,omitempty"`
	Mode      DeploymentRestrictionModeEnum `json:"mode"`
	Type      DeploymentRestrictionTypeEnum `json:"type"`
	// ‘For `PATH` restrictions, the value must not start with `/`’
	Value                string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _TerraformDeploymentRestrictionResponse TerraformDeploymentRestrictionResponse

// NewTerraformDeploymentRestrictionResponse instantiates a new TerraformDeploymentRestrictionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerraformDeploymentRestrictionResponse(id string, createdAt time.Time, mode DeploymentRestrictionModeEnum, type_ DeploymentRestrictionTypeEnum, value string) *TerraformDeploymentRestrictionResponse {
	this := TerraformDeploymentRestrictionResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Mode = mode
	this.Type = type_
	this.Value = value
	return &this
}

// NewTerraformDeploymentRestrictionResponseWithDefaults instantiates a new TerraformDeploymentRestrictionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerraformDeploymentRestrictionResponseWithDefaults() *TerraformDeploymentRestrictionResponse {
	this := TerraformDeploymentRestrictionResponse{}
	return &this
}

// GetId returns the Id field value
func (o *TerraformDeploymentRestrictionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TerraformDeploymentRestrictionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TerraformDeploymentRestrictionResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TerraformDeploymentRestrictionResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TerraformDeploymentRestrictionResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TerraformDeploymentRestrictionResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TerraformDeploymentRestrictionResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerraformDeploymentRestrictionResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TerraformDeploymentRestrictionResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *TerraformDeploymentRestrictionResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetMode returns the Mode field value
func (o *TerraformDeploymentRestrictionResponse) GetMode() DeploymentRestrictionModeEnum {
	if o == nil {
		var ret DeploymentRestrictionModeEnum
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *TerraformDeploymentRestrictionResponse) GetModeOk() (*DeploymentRestrictionModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *TerraformDeploymentRestrictionResponse) SetMode(v DeploymentRestrictionModeEnum) {
	o.Mode = v
}

// GetType returns the Type field value
func (o *TerraformDeploymentRestrictionResponse) GetType() DeploymentRestrictionTypeEnum {
	if o == nil {
		var ret DeploymentRestrictionTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TerraformDeploymentRestrictionResponse) GetTypeOk() (*DeploymentRestrictionTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TerraformDeploymentRestrictionResponse) SetType(v DeploymentRestrictionTypeEnum) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *TerraformDeploymentRestrictionResponse) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TerraformDeploymentRestrictionResponse) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TerraformDeploymentRestrictionResponse) SetValue(v string) {
	o.Value = v
}

func (o TerraformDeploymentRestrictionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerraformDeploymentRestrictionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	toSerialize["mode"] = o.Mode
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TerraformDeploymentRestrictionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"mode",
		"type",
		"value",
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

	varTerraformDeploymentRestrictionResponse := _TerraformDeploymentRestrictionResponse{}

	err = json.Unmarshal(data, &varTerraformDeploymentRestrictionResponse)

	if err != nil {
		return err
	}

	*o = TerraformDeploymentRestrictionResponse(varTerraformDeploymentRestrictionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTerraformDeploymentRestrictionResponse struct {
	value *TerraformDeploymentRestrictionResponse
	isSet bool
}

func (v NullableTerraformDeploymentRestrictionResponse) Get() *TerraformDeploymentRestrictionResponse {
	return v.value
}

func (v *NullableTerraformDeploymentRestrictionResponse) Set(val *TerraformDeploymentRestrictionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTerraformDeploymentRestrictionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTerraformDeploymentRestrictionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerraformDeploymentRestrictionResponse(val *TerraformDeploymentRestrictionResponse) *NullableTerraformDeploymentRestrictionResponse {
	return &NullableTerraformDeploymentRestrictionResponse{value: val, isSet: true}
}

func (v NullableTerraformDeploymentRestrictionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerraformDeploymentRestrictionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
