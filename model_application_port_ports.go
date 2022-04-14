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

// ApplicationPortPorts struct for ApplicationPortPorts
type ApplicationPortPorts struct {
	Id   *string        `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	// The listening port of your application
	InternalPort int32 `json:"internal_port"`
	// The exposed port for your application. This is optional. If not set a default port will be used.
	ExternalPort *int32 `json:"external_port,omitempty"`
	// Expose the port to the world
	PubliclyAccessible bool              `json:"publicly_accessible"`
	Protocol           *PortProtocolEnum `json:"protocol,omitempty"`
}

// NewApplicationPortPorts instantiates a new ApplicationPortPorts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationPortPorts(internalPort int32, publiclyAccessible bool) *ApplicationPortPorts {
	this := ApplicationPortPorts{}
	this.InternalPort = internalPort
	this.PubliclyAccessible = publiclyAccessible
	var protocol PortProtocolEnum = HTTP
	this.Protocol = &protocol
	return &this
}

// NewApplicationPortPortsWithDefaults instantiates a new ApplicationPortPorts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationPortPortsWithDefaults() *ApplicationPortPorts {
	this := ApplicationPortPorts{}
	var protocol PortProtocolEnum = HTTP
	this.Protocol = &protocol
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationPortPorts) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPortPorts) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationPortPorts) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationPortPorts) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationPortPorts) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationPortPorts) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationPortPorts) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ApplicationPortPorts) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *ApplicationPortPorts) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ApplicationPortPorts) UnsetName() {
	o.Name.Unset()
}

// GetInternalPort returns the InternalPort field value
func (o *ApplicationPortPorts) GetInternalPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InternalPort
}

// GetInternalPortOk returns a tuple with the InternalPort field value
// and a boolean to check if the value has been set.
func (o *ApplicationPortPorts) GetInternalPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InternalPort, true
}

// SetInternalPort sets field value
func (o *ApplicationPortPorts) SetInternalPort(v int32) {
	o.InternalPort = v
}

// GetExternalPort returns the ExternalPort field value if set, zero value otherwise.
func (o *ApplicationPortPorts) GetExternalPort() int32 {
	if o == nil || o.ExternalPort == nil {
		var ret int32
		return ret
	}
	return *o.ExternalPort
}

// GetExternalPortOk returns a tuple with the ExternalPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPortPorts) GetExternalPortOk() (*int32, bool) {
	if o == nil || o.ExternalPort == nil {
		return nil, false
	}
	return o.ExternalPort, true
}

// HasExternalPort returns a boolean if a field has been set.
func (o *ApplicationPortPorts) HasExternalPort() bool {
	if o != nil && o.ExternalPort != nil {
		return true
	}

	return false
}

// SetExternalPort gets a reference to the given int32 and assigns it to the ExternalPort field.
func (o *ApplicationPortPorts) SetExternalPort(v int32) {
	o.ExternalPort = &v
}

// GetPubliclyAccessible returns the PubliclyAccessible field value
func (o *ApplicationPortPorts) GetPubliclyAccessible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PubliclyAccessible
}

// GetPubliclyAccessibleOk returns a tuple with the PubliclyAccessible field value
// and a boolean to check if the value has been set.
func (o *ApplicationPortPorts) GetPubliclyAccessibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PubliclyAccessible, true
}

// SetPubliclyAccessible sets field value
func (o *ApplicationPortPorts) SetPubliclyAccessible(v bool) {
	o.PubliclyAccessible = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ApplicationPortPorts) GetProtocol() PortProtocolEnum {
	if o == nil || o.Protocol == nil {
		var ret PortProtocolEnum
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPortPorts) GetProtocolOk() (*PortProtocolEnum, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ApplicationPortPorts) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given PortProtocolEnum and assigns it to the Protocol field.
func (o *ApplicationPortPorts) SetProtocol(v PortProtocolEnum) {
	o.Protocol = &v
}

func (o ApplicationPortPorts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if true {
		toSerialize["internal_port"] = o.InternalPort
	}
	if o.ExternalPort != nil {
		toSerialize["external_port"] = o.ExternalPort
	}
	if true {
		toSerialize["publicly_accessible"] = o.PubliclyAccessible
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationPortPorts struct {
	value *ApplicationPortPorts
	isSet bool
}

func (v NullableApplicationPortPorts) Get() *ApplicationPortPorts {
	return v.value
}

func (v *NullableApplicationPortPorts) Set(val *ApplicationPortPorts) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationPortPorts) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationPortPorts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationPortPorts(val *ApplicationPortPorts) *NullableApplicationPortPorts {
	return &NullableApplicationPortPorts{value: val, isSet: true}
}

func (v NullableApplicationPortPorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationPortPorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}