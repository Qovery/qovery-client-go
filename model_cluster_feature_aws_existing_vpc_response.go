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

// checks if the ClusterFeatureAwsExistingVpcResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterFeatureAwsExistingVpcResponse{}

// ClusterFeatureAwsExistingVpcResponse struct for ClusterFeatureAwsExistingVpcResponse
type ClusterFeatureAwsExistingVpcResponse struct {
	Type                 ClusterFeatureResponseTypeEnum `json:"type"`
	Value                ClusterFeatureAwsExistingVpc   `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _ClusterFeatureAwsExistingVpcResponse ClusterFeatureAwsExistingVpcResponse

// NewClusterFeatureAwsExistingVpcResponse instantiates a new ClusterFeatureAwsExistingVpcResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterFeatureAwsExistingVpcResponse(type_ ClusterFeatureResponseTypeEnum, value ClusterFeatureAwsExistingVpc) *ClusterFeatureAwsExistingVpcResponse {
	this := ClusterFeatureAwsExistingVpcResponse{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewClusterFeatureAwsExistingVpcResponseWithDefaults instantiates a new ClusterFeatureAwsExistingVpcResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterFeatureAwsExistingVpcResponseWithDefaults() *ClusterFeatureAwsExistingVpcResponse {
	this := ClusterFeatureAwsExistingVpcResponse{}
	return &this
}

// GetType returns the Type field value
func (o *ClusterFeatureAwsExistingVpcResponse) GetType() ClusterFeatureResponseTypeEnum {
	if o == nil {
		var ret ClusterFeatureResponseTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ClusterFeatureAwsExistingVpcResponse) GetTypeOk() (*ClusterFeatureResponseTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ClusterFeatureAwsExistingVpcResponse) SetType(v ClusterFeatureResponseTypeEnum) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *ClusterFeatureAwsExistingVpcResponse) GetValue() ClusterFeatureAwsExistingVpc {
	if o == nil {
		var ret ClusterFeatureAwsExistingVpc
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ClusterFeatureAwsExistingVpcResponse) GetValueOk() (*ClusterFeatureAwsExistingVpc, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ClusterFeatureAwsExistingVpcResponse) SetValue(v ClusterFeatureAwsExistingVpc) {
	o.Value = v
}

func (o ClusterFeatureAwsExistingVpcResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterFeatureAwsExistingVpcResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClusterFeatureAwsExistingVpcResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varClusterFeatureAwsExistingVpcResponse := _ClusterFeatureAwsExistingVpcResponse{}

	err = json.Unmarshal(data, &varClusterFeatureAwsExistingVpcResponse)

	if err != nil {
		return err
	}

	*o = ClusterFeatureAwsExistingVpcResponse(varClusterFeatureAwsExistingVpcResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClusterFeatureAwsExistingVpcResponse struct {
	value *ClusterFeatureAwsExistingVpcResponse
	isSet bool
}

func (v NullableClusterFeatureAwsExistingVpcResponse) Get() *ClusterFeatureAwsExistingVpcResponse {
	return v.value
}

func (v *NullableClusterFeatureAwsExistingVpcResponse) Set(val *ClusterFeatureAwsExistingVpcResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterFeatureAwsExistingVpcResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterFeatureAwsExistingVpcResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterFeatureAwsExistingVpcResponse(val *ClusterFeatureAwsExistingVpcResponse) *NullableClusterFeatureAwsExistingVpcResponse {
	return &NullableClusterFeatureAwsExistingVpcResponse{value: val, isSet: true}
}

func (v NullableClusterFeatureAwsExistingVpcResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterFeatureAwsExistingVpcResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
