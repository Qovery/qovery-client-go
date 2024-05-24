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
	"time"
)

// checks if the OrganizationAnnotationsGroupEnrichedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationAnnotationsGroupEnrichedResponse{}

// OrganizationAnnotationsGroupEnrichedResponse struct for OrganizationAnnotationsGroupEnrichedResponse
type OrganizationAnnotationsGroupEnrichedResponse struct {
	Id                   string                                  `json:"id"`
	CreatedAt            time.Time                               `json:"created_at"`
	UpdatedAt            *time.Time                              `json:"updated_at,omitempty"`
	Name                 string                                  `json:"name"`
	Annotations          []Annotation                            `json:"annotations"`
	Scopes               []OrganizationAnnotationsGroupScopeEnum `json:"scopes"`
	AssociatedItemsCount *int32                                  `json:"associated_items_count,omitempty"`
}

// NewOrganizationAnnotationsGroupEnrichedResponse instantiates a new OrganizationAnnotationsGroupEnrichedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationAnnotationsGroupEnrichedResponse(id string, createdAt time.Time, name string, annotations []Annotation, scopes []OrganizationAnnotationsGroupScopeEnum) *OrganizationAnnotationsGroupEnrichedResponse {
	this := OrganizationAnnotationsGroupEnrichedResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Name = name
	this.Annotations = annotations
	this.Scopes = scopes
	return &this
}

// NewOrganizationAnnotationsGroupEnrichedResponseWithDefaults instantiates a new OrganizationAnnotationsGroupEnrichedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationAnnotationsGroupEnrichedResponseWithDefaults() *OrganizationAnnotationsGroupEnrichedResponse {
	this := OrganizationAnnotationsGroupEnrichedResponse{}
	return &this
}

// GetId returns the Id field value
func (o *OrganizationAnnotationsGroupEnrichedResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrganizationAnnotationsGroupEnrichedResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrganizationAnnotationsGroupEnrichedResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *OrganizationAnnotationsGroupEnrichedResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OrganizationAnnotationsGroupEnrichedResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *OrganizationAnnotationsGroupEnrichedResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OrganizationAnnotationsGroupEnrichedResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAnnotationsGroupEnrichedResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OrganizationAnnotationsGroupEnrichedResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *OrganizationAnnotationsGroupEnrichedResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value
func (o *OrganizationAnnotationsGroupEnrichedResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationAnnotationsGroupEnrichedResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationAnnotationsGroupEnrichedResponse) SetName(v string) {
	o.Name = v
}

// GetAnnotations returns the Annotations field value
func (o *OrganizationAnnotationsGroupEnrichedResponse) GetAnnotations() []Annotation {
	if o == nil {
		var ret []Annotation
		return ret
	}

	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value
// and a boolean to check if the value has been set.
func (o *OrganizationAnnotationsGroupEnrichedResponse) GetAnnotationsOk() ([]Annotation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Annotations, true
}

// SetAnnotations sets field value
func (o *OrganizationAnnotationsGroupEnrichedResponse) SetAnnotations(v []Annotation) {
	o.Annotations = v
}

// GetScopes returns the Scopes field value
func (o *OrganizationAnnotationsGroupEnrichedResponse) GetScopes() []OrganizationAnnotationsGroupScopeEnum {
	if o == nil {
		var ret []OrganizationAnnotationsGroupScopeEnum
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *OrganizationAnnotationsGroupEnrichedResponse) GetScopesOk() ([]OrganizationAnnotationsGroupScopeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *OrganizationAnnotationsGroupEnrichedResponse) SetScopes(v []OrganizationAnnotationsGroupScopeEnum) {
	o.Scopes = v
}

// GetAssociatedItemsCount returns the AssociatedItemsCount field value if set, zero value otherwise.
func (o *OrganizationAnnotationsGroupEnrichedResponse) GetAssociatedItemsCount() int32 {
	if o == nil || IsNil(o.AssociatedItemsCount) {
		var ret int32
		return ret
	}
	return *o.AssociatedItemsCount
}

// GetAssociatedItemsCountOk returns a tuple with the AssociatedItemsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAnnotationsGroupEnrichedResponse) GetAssociatedItemsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AssociatedItemsCount) {
		return nil, false
	}
	return o.AssociatedItemsCount, true
}

// HasAssociatedItemsCount returns a boolean if a field has been set.
func (o *OrganizationAnnotationsGroupEnrichedResponse) HasAssociatedItemsCount() bool {
	if o != nil && !IsNil(o.AssociatedItemsCount) {
		return true
	}

	return false
}

// SetAssociatedItemsCount gets a reference to the given int32 and assigns it to the AssociatedItemsCount field.
func (o *OrganizationAnnotationsGroupEnrichedResponse) SetAssociatedItemsCount(v int32) {
	o.AssociatedItemsCount = &v
}

func (o OrganizationAnnotationsGroupEnrichedResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationAnnotationsGroupEnrichedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	toSerialize["name"] = o.Name
	toSerialize["annotations"] = o.Annotations
	toSerialize["scopes"] = o.Scopes
	if !IsNil(o.AssociatedItemsCount) {
		toSerialize["associated_items_count"] = o.AssociatedItemsCount
	}
	return toSerialize, nil
}

type NullableOrganizationAnnotationsGroupEnrichedResponse struct {
	value *OrganizationAnnotationsGroupEnrichedResponse
	isSet bool
}

func (v NullableOrganizationAnnotationsGroupEnrichedResponse) Get() *OrganizationAnnotationsGroupEnrichedResponse {
	return v.value
}

func (v *NullableOrganizationAnnotationsGroupEnrichedResponse) Set(val *OrganizationAnnotationsGroupEnrichedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationAnnotationsGroupEnrichedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationAnnotationsGroupEnrichedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationAnnotationsGroupEnrichedResponse(val *OrganizationAnnotationsGroupEnrichedResponse) *NullableOrganizationAnnotationsGroupEnrichedResponse {
	return &NullableOrganizationAnnotationsGroupEnrichedResponse{value: val, isSet: true}
}

func (v NullableOrganizationAnnotationsGroupEnrichedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationAnnotationsGroupEnrichedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}