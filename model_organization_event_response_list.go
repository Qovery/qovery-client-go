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

// OrganizationEventResponseList struct for OrganizationEventResponseList
type OrganizationEventResponseList struct {
	Links  *OrganizationEventResponseListLinks `json:"links,omitempty"`
	Events []OrganizationEventResponse         `json:"events,omitempty"`
}

// NewOrganizationEventResponseList instantiates a new OrganizationEventResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationEventResponseList() *OrganizationEventResponseList {
	this := OrganizationEventResponseList{}
	return &this
}

// NewOrganizationEventResponseListWithDefaults instantiates a new OrganizationEventResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationEventResponseListWithDefaults() *OrganizationEventResponseList {
	this := OrganizationEventResponseList{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrganizationEventResponseList) GetLinks() OrganizationEventResponseListLinks {
	if o == nil || o.Links == nil {
		var ret OrganizationEventResponseListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEventResponseList) GetLinksOk() (*OrganizationEventResponseListLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrganizationEventResponseList) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OrganizationEventResponseListLinks and assigns it to the Links field.
func (o *OrganizationEventResponseList) SetLinks(v OrganizationEventResponseListLinks) {
	o.Links = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *OrganizationEventResponseList) GetEvents() []OrganizationEventResponse {
	if o == nil || o.Events == nil {
		var ret []OrganizationEventResponse
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEventResponseList) GetEventsOk() ([]OrganizationEventResponse, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *OrganizationEventResponseList) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []OrganizationEventResponse and assigns it to the Events field.
func (o *OrganizationEventResponseList) SetEvents(v []OrganizationEventResponse) {
	o.Events = v
}

func (o OrganizationEventResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationEventResponseList struct {
	value *OrganizationEventResponseList
	isSet bool
}

func (v NullableOrganizationEventResponseList) Get() *OrganizationEventResponseList {
	return v.value
}

func (v *NullableOrganizationEventResponseList) Set(val *OrganizationEventResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationEventResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationEventResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationEventResponseList(val *OrganizationEventResponseList) *NullableOrganizationEventResponseList {
	return &NullableOrganizationEventResponseList{value: val, isSet: true}
}

func (v NullableOrganizationEventResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationEventResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}