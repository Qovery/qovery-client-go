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

// checks if the ScalewayClusterCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScalewayClusterCredentials{}

// ScalewayClusterCredentials struct for ScalewayClusterCredentials
type ScalewayClusterCredentials struct {
	Id                     string `json:"id"`
	Name                   string `json:"name"`
	ScalewayAccessKey      string `json:"scaleway_access_key"`
	ObjectType             string `json:"object_type"`
	ScalewayProjectId      string `json:"scaleway_project_id"`
	ScalewayOrganizationId string `json:"scaleway_organization_id"`
	AdditionalProperties   map[string]interface{}
}

type _ScalewayClusterCredentials ScalewayClusterCredentials

// NewScalewayClusterCredentials instantiates a new ScalewayClusterCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScalewayClusterCredentials(id string, name string, scalewayAccessKey string, objectType string, scalewayProjectId string, scalewayOrganizationId string) *ScalewayClusterCredentials {
	this := ScalewayClusterCredentials{}
	this.Id = id
	this.Name = name
	this.ScalewayAccessKey = scalewayAccessKey
	this.ObjectType = objectType
	this.ScalewayProjectId = scalewayProjectId
	this.ScalewayOrganizationId = scalewayOrganizationId
	return &this
}

// NewScalewayClusterCredentialsWithDefaults instantiates a new ScalewayClusterCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScalewayClusterCredentialsWithDefaults() *ScalewayClusterCredentials {
	this := ScalewayClusterCredentials{}
	return &this
}

// GetId returns the Id field value
func (o *ScalewayClusterCredentials) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScalewayClusterCredentials) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScalewayClusterCredentials) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ScalewayClusterCredentials) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScalewayClusterCredentials) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ScalewayClusterCredentials) SetName(v string) {
	o.Name = v
}

// GetScalewayAccessKey returns the ScalewayAccessKey field value
func (o *ScalewayClusterCredentials) GetScalewayAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScalewayAccessKey
}

// GetScalewayAccessKeyOk returns a tuple with the ScalewayAccessKey field value
// and a boolean to check if the value has been set.
func (o *ScalewayClusterCredentials) GetScalewayAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScalewayAccessKey, true
}

// SetScalewayAccessKey sets field value
func (o *ScalewayClusterCredentials) SetScalewayAccessKey(v string) {
	o.ScalewayAccessKey = v
}

// GetObjectType returns the ObjectType field value
func (o *ScalewayClusterCredentials) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ScalewayClusterCredentials) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ScalewayClusterCredentials) SetObjectType(v string) {
	o.ObjectType = v
}

// GetScalewayProjectId returns the ScalewayProjectId field value
func (o *ScalewayClusterCredentials) GetScalewayProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScalewayProjectId
}

// GetScalewayProjectIdOk returns a tuple with the ScalewayProjectId field value
// and a boolean to check if the value has been set.
func (o *ScalewayClusterCredentials) GetScalewayProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScalewayProjectId, true
}

// SetScalewayProjectId sets field value
func (o *ScalewayClusterCredentials) SetScalewayProjectId(v string) {
	o.ScalewayProjectId = v
}

// GetScalewayOrganizationId returns the ScalewayOrganizationId field value
func (o *ScalewayClusterCredentials) GetScalewayOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScalewayOrganizationId
}

// GetScalewayOrganizationIdOk returns a tuple with the ScalewayOrganizationId field value
// and a boolean to check if the value has been set.
func (o *ScalewayClusterCredentials) GetScalewayOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScalewayOrganizationId, true
}

// SetScalewayOrganizationId sets field value
func (o *ScalewayClusterCredentials) SetScalewayOrganizationId(v string) {
	o.ScalewayOrganizationId = v
}

func (o ScalewayClusterCredentials) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScalewayClusterCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["scaleway_access_key"] = o.ScalewayAccessKey
	toSerialize["object_type"] = o.ObjectType
	toSerialize["scaleway_project_id"] = o.ScalewayProjectId
	toSerialize["scaleway_organization_id"] = o.ScalewayOrganizationId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScalewayClusterCredentials) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"scaleway_access_key",
		"object_type",
		"scaleway_project_id",
		"scaleway_organization_id",
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

	varScalewayClusterCredentials := _ScalewayClusterCredentials{}

	err = json.Unmarshal(data, &varScalewayClusterCredentials)

	if err != nil {
		return err
	}

	*o = ScalewayClusterCredentials(varScalewayClusterCredentials)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "scaleway_access_key")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "scaleway_project_id")
		delete(additionalProperties, "scaleway_organization_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScalewayClusterCredentials struct {
	value *ScalewayClusterCredentials
	isSet bool
}

func (v NullableScalewayClusterCredentials) Get() *ScalewayClusterCredentials {
	return v.value
}

func (v *NullableScalewayClusterCredentials) Set(val *ScalewayClusterCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableScalewayClusterCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableScalewayClusterCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScalewayClusterCredentials(val *ScalewayClusterCredentials) *NullableScalewayClusterCredentials {
	return &NullableScalewayClusterCredentials{value: val, isSet: true}
}

func (v NullableScalewayClusterCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScalewayClusterCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}