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
)

// checks if the ApplicationDeploymentRestrictionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationDeploymentRestrictionRequest{}

// ApplicationDeploymentRestrictionRequest struct for ApplicationDeploymentRestrictionRequest
type ApplicationDeploymentRestrictionRequest struct {
	Mode DeploymentRestrictionModeEnum `json:"mode"`
	Type DeploymentRestrictionTypeEnum `json:"type"`
	// For `PATH` restrictions, the value must not start with `/`
	Value                string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationDeploymentRestrictionRequest ApplicationDeploymentRestrictionRequest

// NewApplicationDeploymentRestrictionRequest instantiates a new ApplicationDeploymentRestrictionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationDeploymentRestrictionRequest(mode DeploymentRestrictionModeEnum, type_ DeploymentRestrictionTypeEnum, value string) *ApplicationDeploymentRestrictionRequest {
	this := ApplicationDeploymentRestrictionRequest{}
	this.Mode = mode
	this.Type = type_
	this.Value = value
	return &this
}

// NewApplicationDeploymentRestrictionRequestWithDefaults instantiates a new ApplicationDeploymentRestrictionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationDeploymentRestrictionRequestWithDefaults() *ApplicationDeploymentRestrictionRequest {
	this := ApplicationDeploymentRestrictionRequest{}
	return &this
}

// GetMode returns the Mode field value
func (o *ApplicationDeploymentRestrictionRequest) GetMode() DeploymentRestrictionModeEnum {
	if o == nil {
		var ret DeploymentRestrictionModeEnum
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *ApplicationDeploymentRestrictionRequest) GetModeOk() (*DeploymentRestrictionModeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *ApplicationDeploymentRestrictionRequest) SetMode(v DeploymentRestrictionModeEnum) {
	o.Mode = v
}

// GetType returns the Type field value
func (o *ApplicationDeploymentRestrictionRequest) GetType() DeploymentRestrictionTypeEnum {
	if o == nil {
		var ret DeploymentRestrictionTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationDeploymentRestrictionRequest) GetTypeOk() (*DeploymentRestrictionTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationDeploymentRestrictionRequest) SetType(v DeploymentRestrictionTypeEnum) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *ApplicationDeploymentRestrictionRequest) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ApplicationDeploymentRestrictionRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ApplicationDeploymentRestrictionRequest) SetValue(v string) {
	o.Value = v
}

func (o ApplicationDeploymentRestrictionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationDeploymentRestrictionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mode"] = o.Mode
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationDeploymentRestrictionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varApplicationDeploymentRestrictionRequest := _ApplicationDeploymentRestrictionRequest{}

	err = json.Unmarshal(data, &varApplicationDeploymentRestrictionRequest)

	if err != nil {
		return err
	}

	*o = ApplicationDeploymentRestrictionRequest(varApplicationDeploymentRestrictionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mode")
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplicationDeploymentRestrictionRequest struct {
	value *ApplicationDeploymentRestrictionRequest
	isSet bool
}

func (v NullableApplicationDeploymentRestrictionRequest) Get() *ApplicationDeploymentRestrictionRequest {
	return v.value
}

func (v *NullableApplicationDeploymentRestrictionRequest) Set(val *ApplicationDeploymentRestrictionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationDeploymentRestrictionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationDeploymentRestrictionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationDeploymentRestrictionRequest(val *ApplicationDeploymentRestrictionRequest) *NullableApplicationDeploymentRestrictionRequest {
	return &NullableApplicationDeploymentRestrictionRequest{value: val, isSet: true}
}

func (v NullableApplicationDeploymentRestrictionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationDeploymentRestrictionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
