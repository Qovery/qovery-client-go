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

// checks if the AksInfrastructureOutputs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AksInfrastructureOutputs{}

// AksInfrastructureOutputs struct for AksInfrastructureOutputs
type AksInfrastructureOutputs struct {
	Kind                 string `json:"kind"`
	ClusterName          string `json:"cluster_name"`
	ClusterOidcIssuer    string `json:"cluster_oidc_issuer"`
	AdditionalProperties map[string]interface{}
}

type _AksInfrastructureOutputs AksInfrastructureOutputs

// NewAksInfrastructureOutputs instantiates a new AksInfrastructureOutputs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAksInfrastructureOutputs(kind string, clusterName string, clusterOidcIssuer string) *AksInfrastructureOutputs {
	this := AksInfrastructureOutputs{}
	this.Kind = kind
	this.ClusterName = clusterName
	this.ClusterOidcIssuer = clusterOidcIssuer
	return &this
}

// NewAksInfrastructureOutputsWithDefaults instantiates a new AksInfrastructureOutputs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAksInfrastructureOutputsWithDefaults() *AksInfrastructureOutputs {
	this := AksInfrastructureOutputs{}
	return &this
}

// GetKind returns the Kind field value
func (o *AksInfrastructureOutputs) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *AksInfrastructureOutputs) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *AksInfrastructureOutputs) SetKind(v string) {
	o.Kind = v
}

// GetClusterName returns the ClusterName field value
func (o *AksInfrastructureOutputs) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *AksInfrastructureOutputs) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value
func (o *AksInfrastructureOutputs) SetClusterName(v string) {
	o.ClusterName = v
}

// GetClusterOidcIssuer returns the ClusterOidcIssuer field value
func (o *AksInfrastructureOutputs) GetClusterOidcIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterOidcIssuer
}

// GetClusterOidcIssuerOk returns a tuple with the ClusterOidcIssuer field value
// and a boolean to check if the value has been set.
func (o *AksInfrastructureOutputs) GetClusterOidcIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterOidcIssuer, true
}

// SetClusterOidcIssuer sets field value
func (o *AksInfrastructureOutputs) SetClusterOidcIssuer(v string) {
	o.ClusterOidcIssuer = v
}

func (o AksInfrastructureOutputs) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AksInfrastructureOutputs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["cluster_name"] = o.ClusterName
	toSerialize["cluster_oidc_issuer"] = o.ClusterOidcIssuer

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AksInfrastructureOutputs) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kind",
		"cluster_name",
		"cluster_oidc_issuer",
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

	varAksInfrastructureOutputs := _AksInfrastructureOutputs{}

	err = json.Unmarshal(data, &varAksInfrastructureOutputs)

	if err != nil {
		return err
	}

	*o = AksInfrastructureOutputs(varAksInfrastructureOutputs)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "kind")
		delete(additionalProperties, "cluster_name")
		delete(additionalProperties, "cluster_oidc_issuer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAksInfrastructureOutputs struct {
	value *AksInfrastructureOutputs
	isSet bool
}

func (v NullableAksInfrastructureOutputs) Get() *AksInfrastructureOutputs {
	return v.value
}

func (v *NullableAksInfrastructureOutputs) Set(val *AksInfrastructureOutputs) {
	v.value = val
	v.isSet = true
}

func (v NullableAksInfrastructureOutputs) IsSet() bool {
	return v.isSet
}

func (v *NullableAksInfrastructureOutputs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAksInfrastructureOutputs(val *AksInfrastructureOutputs) *NullableAksInfrastructureOutputs {
	return &NullableAksInfrastructureOutputs{value: val, isSet: true}
}

func (v NullableAksInfrastructureOutputs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAksInfrastructureOutputs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
