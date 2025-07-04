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

// checks if the ScalewayCredentialsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScalewayCredentialsRequest{}

// ScalewayCredentialsRequest struct for ScalewayCredentialsRequest
type ScalewayCredentialsRequest struct {
	Name                   string `json:"name"`
	ScalewayAccessKey      string `json:"scaleway_access_key"`
	ScalewaySecretKey      string `json:"scaleway_secret_key"`
	ScalewayProjectId      string `json:"scaleway_project_id"`
	ScalewayOrganizationId string `json:"scaleway_organization_id"`
	AdditionalProperties   map[string]interface{}
}

type _ScalewayCredentialsRequest ScalewayCredentialsRequest

// NewScalewayCredentialsRequest instantiates a new ScalewayCredentialsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScalewayCredentialsRequest(name string, scalewayAccessKey string, scalewaySecretKey string, scalewayProjectId string, scalewayOrganizationId string) *ScalewayCredentialsRequest {
	this := ScalewayCredentialsRequest{}
	this.Name = name
	this.ScalewayAccessKey = scalewayAccessKey
	this.ScalewaySecretKey = scalewaySecretKey
	this.ScalewayProjectId = scalewayProjectId
	this.ScalewayOrganizationId = scalewayOrganizationId
	return &this
}

// NewScalewayCredentialsRequestWithDefaults instantiates a new ScalewayCredentialsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScalewayCredentialsRequestWithDefaults() *ScalewayCredentialsRequest {
	this := ScalewayCredentialsRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ScalewayCredentialsRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScalewayCredentialsRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ScalewayCredentialsRequest) SetName(v string) {
	o.Name = v
}

// GetScalewayAccessKey returns the ScalewayAccessKey field value
func (o *ScalewayCredentialsRequest) GetScalewayAccessKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScalewayAccessKey
}

// GetScalewayAccessKeyOk returns a tuple with the ScalewayAccessKey field value
// and a boolean to check if the value has been set.
func (o *ScalewayCredentialsRequest) GetScalewayAccessKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScalewayAccessKey, true
}

// SetScalewayAccessKey sets field value
func (o *ScalewayCredentialsRequest) SetScalewayAccessKey(v string) {
	o.ScalewayAccessKey = v
}

// GetScalewaySecretKey returns the ScalewaySecretKey field value
func (o *ScalewayCredentialsRequest) GetScalewaySecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScalewaySecretKey
}

// GetScalewaySecretKeyOk returns a tuple with the ScalewaySecretKey field value
// and a boolean to check if the value has been set.
func (o *ScalewayCredentialsRequest) GetScalewaySecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScalewaySecretKey, true
}

// SetScalewaySecretKey sets field value
func (o *ScalewayCredentialsRequest) SetScalewaySecretKey(v string) {
	o.ScalewaySecretKey = v
}

// GetScalewayProjectId returns the ScalewayProjectId field value
func (o *ScalewayCredentialsRequest) GetScalewayProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScalewayProjectId
}

// GetScalewayProjectIdOk returns a tuple with the ScalewayProjectId field value
// and a boolean to check if the value has been set.
func (o *ScalewayCredentialsRequest) GetScalewayProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScalewayProjectId, true
}

// SetScalewayProjectId sets field value
func (o *ScalewayCredentialsRequest) SetScalewayProjectId(v string) {
	o.ScalewayProjectId = v
}

// GetScalewayOrganizationId returns the ScalewayOrganizationId field value
func (o *ScalewayCredentialsRequest) GetScalewayOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScalewayOrganizationId
}

// GetScalewayOrganizationIdOk returns a tuple with the ScalewayOrganizationId field value
// and a boolean to check if the value has been set.
func (o *ScalewayCredentialsRequest) GetScalewayOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScalewayOrganizationId, true
}

// SetScalewayOrganizationId sets field value
func (o *ScalewayCredentialsRequest) SetScalewayOrganizationId(v string) {
	o.ScalewayOrganizationId = v
}

func (o ScalewayCredentialsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScalewayCredentialsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["scaleway_access_key"] = o.ScalewayAccessKey
	toSerialize["scaleway_secret_key"] = o.ScalewaySecretKey
	toSerialize["scaleway_project_id"] = o.ScalewayProjectId
	toSerialize["scaleway_organization_id"] = o.ScalewayOrganizationId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScalewayCredentialsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"scaleway_access_key",
		"scaleway_secret_key",
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

	varScalewayCredentialsRequest := _ScalewayCredentialsRequest{}

	err = json.Unmarshal(data, &varScalewayCredentialsRequest)

	if err != nil {
		return err
	}

	*o = ScalewayCredentialsRequest(varScalewayCredentialsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "scaleway_access_key")
		delete(additionalProperties, "scaleway_secret_key")
		delete(additionalProperties, "scaleway_project_id")
		delete(additionalProperties, "scaleway_organization_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScalewayCredentialsRequest struct {
	value *ScalewayCredentialsRequest
	isSet bool
}

func (v NullableScalewayCredentialsRequest) Get() *ScalewayCredentialsRequest {
	return v.value
}

func (v *NullableScalewayCredentialsRequest) Set(val *ScalewayCredentialsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScalewayCredentialsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScalewayCredentialsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScalewayCredentialsRequest(val *ScalewayCredentialsRequest) *NullableScalewayCredentialsRequest {
	return &NullableScalewayCredentialsRequest{value: val, isSet: true}
}

func (v NullableScalewayCredentialsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScalewayCredentialsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
