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

// checks if the ClusterInstanceTypeResponseListResultsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterInstanceTypeResponseListResultsInner{}

// ClusterInstanceTypeResponseListResultsInner struct for ClusterInstanceTypeResponseListResultsInner
type ClusterInstanceTypeResponseListResultsInner struct {
	Type                 string                     `json:"type"`
	Name                 string                     `json:"name"`
	Cpu                  int32                      `json:"cpu"`
	RamInGb              int32                      `json:"ram_in_gb"`
	BandwidthInGbps      string                     `json:"bandwidth_in_gbps"`
	BandwidthGuarantee   string                     `json:"bandwidth_guarantee"`
	Architecture         *string                    `json:"architecture,omitempty"`
	GpuInfo              *ClusterInstanceGpuInfo    `json:"gpu_info,omitempty"`
	Attributes           *ClusterInstanceAttributes `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ClusterInstanceTypeResponseListResultsInner ClusterInstanceTypeResponseListResultsInner

// NewClusterInstanceTypeResponseListResultsInner instantiates a new ClusterInstanceTypeResponseListResultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterInstanceTypeResponseListResultsInner(type_ string, name string, cpu int32, ramInGb int32, bandwidthInGbps string, bandwidthGuarantee string) *ClusterInstanceTypeResponseListResultsInner {
	this := ClusterInstanceTypeResponseListResultsInner{}
	this.Type = type_
	this.Name = name
	this.Cpu = cpu
	this.RamInGb = ramInGb
	this.BandwidthInGbps = bandwidthInGbps
	this.BandwidthGuarantee = bandwidthGuarantee
	return &this
}

// NewClusterInstanceTypeResponseListResultsInnerWithDefaults instantiates a new ClusterInstanceTypeResponseListResultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterInstanceTypeResponseListResultsInnerWithDefaults() *ClusterInstanceTypeResponseListResultsInner {
	this := ClusterInstanceTypeResponseListResultsInner{}
	return &this
}

// GetType returns the Type field value
func (o *ClusterInstanceTypeResponseListResultsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ClusterInstanceTypeResponseListResultsInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ClusterInstanceTypeResponseListResultsInner) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *ClusterInstanceTypeResponseListResultsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterInstanceTypeResponseListResultsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterInstanceTypeResponseListResultsInner) SetName(v string) {
	o.Name = v
}

// GetCpu returns the Cpu field value
func (o *ClusterInstanceTypeResponseListResultsInner) GetCpu() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *ClusterInstanceTypeResponseListResultsInner) GetCpuOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpu, true
}

// SetCpu sets field value
func (o *ClusterInstanceTypeResponseListResultsInner) SetCpu(v int32) {
	o.Cpu = v
}

// GetRamInGb returns the RamInGb field value
func (o *ClusterInstanceTypeResponseListResultsInner) GetRamInGb() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RamInGb
}

// GetRamInGbOk returns a tuple with the RamInGb field value
// and a boolean to check if the value has been set.
func (o *ClusterInstanceTypeResponseListResultsInner) GetRamInGbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RamInGb, true
}

// SetRamInGb sets field value
func (o *ClusterInstanceTypeResponseListResultsInner) SetRamInGb(v int32) {
	o.RamInGb = v
}

// GetBandwidthInGbps returns the BandwidthInGbps field value
func (o *ClusterInstanceTypeResponseListResultsInner) GetBandwidthInGbps() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BandwidthInGbps
}

// GetBandwidthInGbpsOk returns a tuple with the BandwidthInGbps field value
// and a boolean to check if the value has been set.
func (o *ClusterInstanceTypeResponseListResultsInner) GetBandwidthInGbpsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BandwidthInGbps, true
}

// SetBandwidthInGbps sets field value
func (o *ClusterInstanceTypeResponseListResultsInner) SetBandwidthInGbps(v string) {
	o.BandwidthInGbps = v
}

// GetBandwidthGuarantee returns the BandwidthGuarantee field value
func (o *ClusterInstanceTypeResponseListResultsInner) GetBandwidthGuarantee() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BandwidthGuarantee
}

// GetBandwidthGuaranteeOk returns a tuple with the BandwidthGuarantee field value
// and a boolean to check if the value has been set.
func (o *ClusterInstanceTypeResponseListResultsInner) GetBandwidthGuaranteeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BandwidthGuarantee, true
}

// SetBandwidthGuarantee sets field value
func (o *ClusterInstanceTypeResponseListResultsInner) SetBandwidthGuarantee(v string) {
	o.BandwidthGuarantee = v
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *ClusterInstanceTypeResponseListResultsInner) GetArchitecture() string {
	if o == nil || IsNil(o.Architecture) {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInstanceTypeResponseListResultsInner) GetArchitectureOk() (*string, bool) {
	if o == nil || IsNil(o.Architecture) {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *ClusterInstanceTypeResponseListResultsInner) HasArchitecture() bool {
	if o != nil && !IsNil(o.Architecture) {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *ClusterInstanceTypeResponseListResultsInner) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetGpuInfo returns the GpuInfo field value if set, zero value otherwise.
func (o *ClusterInstanceTypeResponseListResultsInner) GetGpuInfo() ClusterInstanceGpuInfo {
	if o == nil || IsNil(o.GpuInfo) {
		var ret ClusterInstanceGpuInfo
		return ret
	}
	return *o.GpuInfo
}

// GetGpuInfoOk returns a tuple with the GpuInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInstanceTypeResponseListResultsInner) GetGpuInfoOk() (*ClusterInstanceGpuInfo, bool) {
	if o == nil || IsNil(o.GpuInfo) {
		return nil, false
	}
	return o.GpuInfo, true
}

// HasGpuInfo returns a boolean if a field has been set.
func (o *ClusterInstanceTypeResponseListResultsInner) HasGpuInfo() bool {
	if o != nil && !IsNil(o.GpuInfo) {
		return true
	}

	return false
}

// SetGpuInfo gets a reference to the given ClusterInstanceGpuInfo and assigns it to the GpuInfo field.
func (o *ClusterInstanceTypeResponseListResultsInner) SetGpuInfo(v ClusterInstanceGpuInfo) {
	o.GpuInfo = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ClusterInstanceTypeResponseListResultsInner) GetAttributes() ClusterInstanceAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ClusterInstanceAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInstanceTypeResponseListResultsInner) GetAttributesOk() (*ClusterInstanceAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ClusterInstanceTypeResponseListResultsInner) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ClusterInstanceAttributes and assigns it to the Attributes field.
func (o *ClusterInstanceTypeResponseListResultsInner) SetAttributes(v ClusterInstanceAttributes) {
	o.Attributes = &v
}

func (o ClusterInstanceTypeResponseListResultsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterInstanceTypeResponseListResultsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	toSerialize["cpu"] = o.Cpu
	toSerialize["ram_in_gb"] = o.RamInGb
	toSerialize["bandwidth_in_gbps"] = o.BandwidthInGbps
	toSerialize["bandwidth_guarantee"] = o.BandwidthGuarantee
	if !IsNil(o.Architecture) {
		toSerialize["architecture"] = o.Architecture
	}
	if !IsNil(o.GpuInfo) {
		toSerialize["gpu_info"] = o.GpuInfo
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClusterInstanceTypeResponseListResultsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"name",
		"cpu",
		"ram_in_gb",
		"bandwidth_in_gbps",
		"bandwidth_guarantee",
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

	varClusterInstanceTypeResponseListResultsInner := _ClusterInstanceTypeResponseListResultsInner{}

	err = json.Unmarshal(data, &varClusterInstanceTypeResponseListResultsInner)

	if err != nil {
		return err
	}

	*o = ClusterInstanceTypeResponseListResultsInner(varClusterInstanceTypeResponseListResultsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "cpu")
		delete(additionalProperties, "ram_in_gb")
		delete(additionalProperties, "bandwidth_in_gbps")
		delete(additionalProperties, "bandwidth_guarantee")
		delete(additionalProperties, "architecture")
		delete(additionalProperties, "gpu_info")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClusterInstanceTypeResponseListResultsInner struct {
	value *ClusterInstanceTypeResponseListResultsInner
	isSet bool
}

func (v NullableClusterInstanceTypeResponseListResultsInner) Get() *ClusterInstanceTypeResponseListResultsInner {
	return v.value
}

func (v *NullableClusterInstanceTypeResponseListResultsInner) Set(val *ClusterInstanceTypeResponseListResultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterInstanceTypeResponseListResultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterInstanceTypeResponseListResultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterInstanceTypeResponseListResultsInner(val *ClusterInstanceTypeResponseListResultsInner) *NullableClusterInstanceTypeResponseListResultsInner {
	return &NullableClusterInstanceTypeResponseListResultsInner{value: val, isSet: true}
}

func (v NullableClusterInstanceTypeResponseListResultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterInstanceTypeResponseListResultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
