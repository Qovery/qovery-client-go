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

// checks if the HelmPortRequestPortsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HelmPortRequestPortsInner{}

// HelmPortRequestPortsInner struct for HelmPortRequestPortsInner
type HelmPortRequestPortsInner struct {
	Name *string `json:"name,omitempty"`
	// The listening port of your service.
	InternalPort int32 `json:"internal_port"`
	// The exposed port for your service. This is optional. If not set a default port will be used.
	ExternalPort *int32                `json:"external_port,omitempty"`
	Namespace    *string               `json:"namespace,omitempty"`
	Protocol     *HelmPortProtocolEnum `json:"protocol,omitempty"`
	// is the default port to use for domain
	IsDefault            *bool                `json:"is_default,omitempty"`
	ServiceSelectors     []KubernetesSelector `json:"service_selectors,omitempty"`
	ServiceName          *string              `json:"service_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HelmPortRequestPortsInner HelmPortRequestPortsInner

// NewHelmPortRequestPortsInner instantiates a new HelmPortRequestPortsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHelmPortRequestPortsInner(internalPort int32) *HelmPortRequestPortsInner {
	this := HelmPortRequestPortsInner{}
	this.InternalPort = internalPort
	var protocol HelmPortProtocolEnum = HELMPORTPROTOCOLENUM_HTTP
	this.Protocol = &protocol
	return &this
}

// NewHelmPortRequestPortsInnerWithDefaults instantiates a new HelmPortRequestPortsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmPortRequestPortsInnerWithDefaults() *HelmPortRequestPortsInner {
	this := HelmPortRequestPortsInner{}
	var protocol HelmPortProtocolEnum = HELMPORTPROTOCOLENUM_HTTP
	this.Protocol = &protocol
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HelmPortRequestPortsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmPortRequestPortsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HelmPortRequestPortsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HelmPortRequestPortsInner) SetName(v string) {
	o.Name = &v
}

// GetInternalPort returns the InternalPort field value
func (o *HelmPortRequestPortsInner) GetInternalPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InternalPort
}

// GetInternalPortOk returns a tuple with the InternalPort field value
// and a boolean to check if the value has been set.
func (o *HelmPortRequestPortsInner) GetInternalPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InternalPort, true
}

// SetInternalPort sets field value
func (o *HelmPortRequestPortsInner) SetInternalPort(v int32) {
	o.InternalPort = v
}

// GetExternalPort returns the ExternalPort field value if set, zero value otherwise.
func (o *HelmPortRequestPortsInner) GetExternalPort() int32 {
	if o == nil || IsNil(o.ExternalPort) {
		var ret int32
		return ret
	}
	return *o.ExternalPort
}

// GetExternalPortOk returns a tuple with the ExternalPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmPortRequestPortsInner) GetExternalPortOk() (*int32, bool) {
	if o == nil || IsNil(o.ExternalPort) {
		return nil, false
	}
	return o.ExternalPort, true
}

// HasExternalPort returns a boolean if a field has been set.
func (o *HelmPortRequestPortsInner) HasExternalPort() bool {
	if o != nil && !IsNil(o.ExternalPort) {
		return true
	}

	return false
}

// SetExternalPort gets a reference to the given int32 and assigns it to the ExternalPort field.
func (o *HelmPortRequestPortsInner) SetExternalPort(v int32) {
	o.ExternalPort = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *HelmPortRequestPortsInner) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmPortRequestPortsInner) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *HelmPortRequestPortsInner) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *HelmPortRequestPortsInner) SetNamespace(v string) {
	o.Namespace = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *HelmPortRequestPortsInner) GetProtocol() HelmPortProtocolEnum {
	if o == nil || IsNil(o.Protocol) {
		var ret HelmPortProtocolEnum
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmPortRequestPortsInner) GetProtocolOk() (*HelmPortProtocolEnum, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *HelmPortRequestPortsInner) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given HelmPortProtocolEnum and assigns it to the Protocol field.
func (o *HelmPortRequestPortsInner) SetProtocol(v HelmPortProtocolEnum) {
	o.Protocol = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *HelmPortRequestPortsInner) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmPortRequestPortsInner) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *HelmPortRequestPortsInner) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *HelmPortRequestPortsInner) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetServiceSelectors returns the ServiceSelectors field value if set, zero value otherwise.
func (o *HelmPortRequestPortsInner) GetServiceSelectors() []KubernetesSelector {
	if o == nil || IsNil(o.ServiceSelectors) {
		var ret []KubernetesSelector
		return ret
	}
	return o.ServiceSelectors
}

// GetServiceSelectorsOk returns a tuple with the ServiceSelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmPortRequestPortsInner) GetServiceSelectorsOk() ([]KubernetesSelector, bool) {
	if o == nil || IsNil(o.ServiceSelectors) {
		return nil, false
	}
	return o.ServiceSelectors, true
}

// HasServiceSelectors returns a boolean if a field has been set.
func (o *HelmPortRequestPortsInner) HasServiceSelectors() bool {
	if o != nil && !IsNil(o.ServiceSelectors) {
		return true
	}

	return false
}

// SetServiceSelectors gets a reference to the given []KubernetesSelector and assigns it to the ServiceSelectors field.
func (o *HelmPortRequestPortsInner) SetServiceSelectors(v []KubernetesSelector) {
	o.ServiceSelectors = v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *HelmPortRequestPortsInner) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmPortRequestPortsInner) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *HelmPortRequestPortsInner) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *HelmPortRequestPortsInner) SetServiceName(v string) {
	o.ServiceName = &v
}

func (o HelmPortRequestPortsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmPortRequestPortsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["internal_port"] = o.InternalPort
	if !IsNil(o.ExternalPort) {
		toSerialize["external_port"] = o.ExternalPort
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.IsDefault) {
		toSerialize["is_default"] = o.IsDefault
	}
	if !IsNil(o.ServiceSelectors) {
		toSerialize["service_selectors"] = o.ServiceSelectors
	}
	if !IsNil(o.ServiceName) {
		toSerialize["service_name"] = o.ServiceName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HelmPortRequestPortsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"internal_port",
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

	varHelmPortRequestPortsInner := _HelmPortRequestPortsInner{}

	err = json.Unmarshal(data, &varHelmPortRequestPortsInner)

	if err != nil {
		return err
	}

	*o = HelmPortRequestPortsInner(varHelmPortRequestPortsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "internal_port")
		delete(additionalProperties, "external_port")
		delete(additionalProperties, "namespace")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "is_default")
		delete(additionalProperties, "service_selectors")
		delete(additionalProperties, "service_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHelmPortRequestPortsInner struct {
	value *HelmPortRequestPortsInner
	isSet bool
}

func (v NullableHelmPortRequestPortsInner) Get() *HelmPortRequestPortsInner {
	return v.value
}

func (v *NullableHelmPortRequestPortsInner) Set(val *HelmPortRequestPortsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableHelmPortRequestPortsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmPortRequestPortsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmPortRequestPortsInner(val *HelmPortRequestPortsInner) *NullableHelmPortRequestPortsInner {
	return &NullableHelmPortRequestPortsInner{value: val, isSet: true}
}

func (v NullableHelmPortRequestPortsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmPortRequestPortsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
