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

// checks if the OrganizationAnnotationsGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationAnnotationsGroupResponse{}

// OrganizationAnnotationsGroupResponse struct for OrganizationAnnotationsGroupResponse
type OrganizationAnnotationsGroupResponse struct {
	Id          string                                  `json:"id"`
	Name        string                                  `json:"name"`
	Annotations []Annotation                            `json:"annotations"`
	Scopes      []OrganizationAnnotationsGroupScopeEnum `json:"scopes"`
}

// NewOrganizationAnnotationsGroupResponse instantiates a new OrganizationAnnotationsGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationAnnotationsGroupResponse(id string, name string, annotations []Annotation, scopes []OrganizationAnnotationsGroupScopeEnum) *OrganizationAnnotationsGroupResponse {
	this := OrganizationAnnotationsGroupResponse{}
	this.Id = id
	this.Name = name
	this.Annotations = annotations
	this.Scopes = scopes
	return &this
}

// NewOrganizationAnnotationsGroupResponseWithDefaults instantiates a new OrganizationAnnotationsGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationAnnotationsGroupResponseWithDefaults() *OrganizationAnnotationsGroupResponse {
	this := OrganizationAnnotationsGroupResponse{}
	return &this
}

// GetId returns the Id field value
func (o *OrganizationAnnotationsGroupResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationAnnotationsGroupResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationAnnotationsGroupResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *OrganizationAnnotationsGroupResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationAnnotationsGroupResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationAnnotationsGroupResponse) SetName(v string) {
	o.Name = v
}

// GetAnnotations returns the Annotations field value
func (o *OrganizationAnnotationsGroupResponse) GetAnnotations() []Annotation {
	if o == nil {
		var ret []Annotation
		return ret
	}

	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value
// and a boolean to check if the value has been set.
func (o *OrganizationAnnotationsGroupResponse) GetAnnotationsOk() ([]Annotation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Annotations, true
}

// SetAnnotations sets field value
func (o *OrganizationAnnotationsGroupResponse) SetAnnotations(v []Annotation) {
	o.Annotations = v
}

// GetScopes returns the Scopes field value
func (o *OrganizationAnnotationsGroupResponse) GetScopes() []OrganizationAnnotationsGroupScopeEnum {
	if o == nil {
		var ret []OrganizationAnnotationsGroupScopeEnum
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *OrganizationAnnotationsGroupResponse) GetScopesOk() ([]OrganizationAnnotationsGroupScopeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *OrganizationAnnotationsGroupResponse) SetScopes(v []OrganizationAnnotationsGroupScopeEnum) {
	o.Scopes = v
}

func (o OrganizationAnnotationsGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationAnnotationsGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["annotations"] = o.Annotations
	toSerialize["scopes"] = o.Scopes
	return toSerialize, nil
}

type NullableOrganizationAnnotationsGroupResponse struct {
	value *OrganizationAnnotationsGroupResponse
	isSet bool
}

func (v NullableOrganizationAnnotationsGroupResponse) Get() *OrganizationAnnotationsGroupResponse {
	return v.value
}

func (v *NullableOrganizationAnnotationsGroupResponse) Set(val *OrganizationAnnotationsGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationAnnotationsGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationAnnotationsGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationAnnotationsGroupResponse(val *OrganizationAnnotationsGroupResponse) *NullableOrganizationAnnotationsGroupResponse {
	return &NullableOrganizationAnnotationsGroupResponse{value: val, isSet: true}
}

func (v NullableOrganizationAnnotationsGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationAnnotationsGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
