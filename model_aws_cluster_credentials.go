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
	"fmt"
)

// checks if the AwsClusterCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsClusterCredentials{}

// AwsClusterCredentials struct for AwsClusterCredentials
type AwsClusterCredentials struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	AccessKeyId          string `json:"access_key_id"`
	ObjectType           string `json:"object_type"`
	AdditionalProperties map[string]interface{}
}

type _AwsClusterCredentials AwsClusterCredentials

// NewAwsClusterCredentials instantiates a new AwsClusterCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsClusterCredentials(id string, name string, accessKeyId string, objectType string) *AwsClusterCredentials {
	this := AwsClusterCredentials{}
	this.Id = id
	this.Name = name
	this.AccessKeyId = accessKeyId
	this.ObjectType = objectType
	return &this
}

// NewAwsClusterCredentialsWithDefaults instantiates a new AwsClusterCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsClusterCredentialsWithDefaults() *AwsClusterCredentials {
	this := AwsClusterCredentials{}
	return &this
}

// GetId returns the Id field value
func (o *AwsClusterCredentials) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AwsClusterCredentials) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AwsClusterCredentials) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AwsClusterCredentials) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AwsClusterCredentials) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AwsClusterCredentials) SetName(v string) {
	o.Name = v
}

// GetAccessKeyId returns the AccessKeyId field value
func (o *AwsClusterCredentials) GetAccessKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value
// and a boolean to check if the value has been set.
func (o *AwsClusterCredentials) GetAccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKeyId, true
}

// SetAccessKeyId sets field value
func (o *AwsClusterCredentials) SetAccessKeyId(v string) {
	o.AccessKeyId = v
}

// GetObjectType returns the ObjectType field value
func (o *AwsClusterCredentials) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AwsClusterCredentials) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AwsClusterCredentials) SetObjectType(v string) {
	o.ObjectType = v
}

func (o AwsClusterCredentials) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsClusterCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["access_key_id"] = o.AccessKeyId
	toSerialize["object_type"] = o.ObjectType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AwsClusterCredentials) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"access_key_id",
		"object_type",
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

	varAwsClusterCredentials := _AwsClusterCredentials{}

	err = json.Unmarshal(data, &varAwsClusterCredentials)

	if err != nil {
		return err
	}

	*o = AwsClusterCredentials(varAwsClusterCredentials)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "access_key_id")
		delete(additionalProperties, "object_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAwsClusterCredentials struct {
	value *AwsClusterCredentials
	isSet bool
}

func (v NullableAwsClusterCredentials) Get() *AwsClusterCredentials {
	return v.value
}

func (v *NullableAwsClusterCredentials) Set(val *AwsClusterCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsClusterCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsClusterCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsClusterCredentials(val *AwsClusterCredentials) *NullableAwsClusterCredentials {
	return &NullableAwsClusterCredentials{value: val, isSet: true}
}

func (v NullableAwsClusterCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsClusterCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
