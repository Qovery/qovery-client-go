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

// checks if the LifecycleTemplateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LifecycleTemplateResponse{}

// LifecycleTemplateResponse struct for LifecycleTemplateResponse
type LifecycleTemplateResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// location of the template
	SourceUrl     string            `json:"sourceUrl"`
	CloudProvider CloudProviderEnum `json:"cloud_provider"`
	// lis of pre-defined command for each event
	Events               []LifecycleTemplateResponseEventsInner `json:"events"`
	Resources            LifecycleTemplateResponseResources     `json:"resources"`
	Variables            LifecycleTemplateResponseVariables     `json:"variables"`
	AdditionalProperties map[string]interface{}
}

type _LifecycleTemplateResponse LifecycleTemplateResponse

// NewLifecycleTemplateResponse instantiates a new LifecycleTemplateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLifecycleTemplateResponse(id string, name string, description string, sourceUrl string, cloudProvider CloudProviderEnum, events []LifecycleTemplateResponseEventsInner, resources LifecycleTemplateResponseResources, variables LifecycleTemplateResponseVariables) *LifecycleTemplateResponse {
	this := LifecycleTemplateResponse{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.SourceUrl = sourceUrl
	this.CloudProvider = cloudProvider
	this.Events = events
	this.Resources = resources
	this.Variables = variables
	return &this
}

// NewLifecycleTemplateResponseWithDefaults instantiates a new LifecycleTemplateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLifecycleTemplateResponseWithDefaults() *LifecycleTemplateResponse {
	this := LifecycleTemplateResponse{}
	return &this
}

// GetId returns the Id field value
func (o *LifecycleTemplateResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LifecycleTemplateResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LifecycleTemplateResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *LifecycleTemplateResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LifecycleTemplateResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LifecycleTemplateResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *LifecycleTemplateResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *LifecycleTemplateResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *LifecycleTemplateResponse) SetDescription(v string) {
	o.Description = v
}

// GetSourceUrl returns the SourceUrl field value
func (o *LifecycleTemplateResponse) GetSourceUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceUrl
}

// GetSourceUrlOk returns a tuple with the SourceUrl field value
// and a boolean to check if the value has been set.
func (o *LifecycleTemplateResponse) GetSourceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceUrl, true
}

// SetSourceUrl sets field value
func (o *LifecycleTemplateResponse) SetSourceUrl(v string) {
	o.SourceUrl = v
}

// GetCloudProvider returns the CloudProvider field value
func (o *LifecycleTemplateResponse) GetCloudProvider() CloudProviderEnum {
	if o == nil {
		var ret CloudProviderEnum
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *LifecycleTemplateResponse) GetCloudProviderOk() (*CloudProviderEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *LifecycleTemplateResponse) SetCloudProvider(v CloudProviderEnum) {
	o.CloudProvider = v
}

// GetEvents returns the Events field value
func (o *LifecycleTemplateResponse) GetEvents() []LifecycleTemplateResponseEventsInner {
	if o == nil {
		var ret []LifecycleTemplateResponseEventsInner
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *LifecycleTemplateResponse) GetEventsOk() ([]LifecycleTemplateResponseEventsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *LifecycleTemplateResponse) SetEvents(v []LifecycleTemplateResponseEventsInner) {
	o.Events = v
}

// GetResources returns the Resources field value
func (o *LifecycleTemplateResponse) GetResources() LifecycleTemplateResponseResources {
	if o == nil {
		var ret LifecycleTemplateResponseResources
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *LifecycleTemplateResponse) GetResourcesOk() (*LifecycleTemplateResponseResources, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *LifecycleTemplateResponse) SetResources(v LifecycleTemplateResponseResources) {
	o.Resources = v
}

// GetVariables returns the Variables field value
func (o *LifecycleTemplateResponse) GetVariables() LifecycleTemplateResponseVariables {
	if o == nil {
		var ret LifecycleTemplateResponseVariables
		return ret
	}

	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value
// and a boolean to check if the value has been set.
func (o *LifecycleTemplateResponse) GetVariablesOk() (*LifecycleTemplateResponseVariables, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variables, true
}

// SetVariables sets field value
func (o *LifecycleTemplateResponse) SetVariables(v LifecycleTemplateResponseVariables) {
	o.Variables = v
}

func (o LifecycleTemplateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LifecycleTemplateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["sourceUrl"] = o.SourceUrl
	toSerialize["cloud_provider"] = o.CloudProvider
	toSerialize["events"] = o.Events
	toSerialize["resources"] = o.Resources
	toSerialize["variables"] = o.Variables

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LifecycleTemplateResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
		"sourceUrl",
		"cloud_provider",
		"events",
		"resources",
		"variables",
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

	varLifecycleTemplateResponse := _LifecycleTemplateResponse{}

	err = json.Unmarshal(data, &varLifecycleTemplateResponse)

	if err != nil {
		return err
	}

	*o = LifecycleTemplateResponse(varLifecycleTemplateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "sourceUrl")
		delete(additionalProperties, "cloud_provider")
		delete(additionalProperties, "events")
		delete(additionalProperties, "resources")
		delete(additionalProperties, "variables")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLifecycleTemplateResponse struct {
	value *LifecycleTemplateResponse
	isSet bool
}

func (v NullableLifecycleTemplateResponse) Get() *LifecycleTemplateResponse {
	return v.value
}

func (v *NullableLifecycleTemplateResponse) Set(val *LifecycleTemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLifecycleTemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLifecycleTemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLifecycleTemplateResponse(val *LifecycleTemplateResponse) *NullableLifecycleTemplateResponse {
	return &NullableLifecycleTemplateResponse{value: val, isSet: true}
}

func (v NullableLifecycleTemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLifecycleTemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
