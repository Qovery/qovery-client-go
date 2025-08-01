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

// checks if the OrganizationEditRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationEditRequest{}

// OrganizationEditRequest struct for OrganizationEditRequest
type OrganizationEditRequest struct {
	// name is case insensitive
	Name                 string         `json:"name"`
	Description          *string        `json:"description,omitempty"`
	WebsiteUrl           NullableString `json:"website_url,omitempty"`
	Repository           NullableString `json:"repository,omitempty"`
	LogoUrl              NullableString `json:"logo_url,omitempty"`
	IconUrl              NullableString `json:"icon_url,omitempty"`
	AdminEmails          []string       `json:"admin_emails,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrganizationEditRequest OrganizationEditRequest

// NewOrganizationEditRequest instantiates a new OrganizationEditRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationEditRequest(name string) *OrganizationEditRequest {
	this := OrganizationEditRequest{}
	this.Name = name
	return &this
}

// NewOrganizationEditRequestWithDefaults instantiates a new OrganizationEditRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationEditRequestWithDefaults() *OrganizationEditRequest {
	this := OrganizationEditRequest{}
	return &this
}

// GetName returns the Name field value
func (o *OrganizationEditRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrganizationEditRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrganizationEditRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrganizationEditRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationEditRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrganizationEditRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrganizationEditRequest) SetDescription(v string) {
	o.Description = &v
}

// GetWebsiteUrl returns the WebsiteUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationEditRequest) GetWebsiteUrl() string {
	if o == nil || IsNil(o.WebsiteUrl.Get()) {
		var ret string
		return ret
	}
	return *o.WebsiteUrl.Get()
}

// GetWebsiteUrlOk returns a tuple with the WebsiteUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationEditRequest) GetWebsiteUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebsiteUrl.Get(), o.WebsiteUrl.IsSet()
}

// HasWebsiteUrl returns a boolean if a field has been set.
func (o *OrganizationEditRequest) HasWebsiteUrl() bool {
	if o != nil && o.WebsiteUrl.IsSet() {
		return true
	}

	return false
}

// SetWebsiteUrl gets a reference to the given NullableString and assigns it to the WebsiteUrl field.
func (o *OrganizationEditRequest) SetWebsiteUrl(v string) {
	o.WebsiteUrl.Set(&v)
}

// SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil
func (o *OrganizationEditRequest) SetWebsiteUrlNil() {
	o.WebsiteUrl.Set(nil)
}

// UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
func (o *OrganizationEditRequest) UnsetWebsiteUrl() {
	o.WebsiteUrl.Unset()
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationEditRequest) GetRepository() string {
	if o == nil || IsNil(o.Repository.Get()) {
		var ret string
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationEditRequest) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *OrganizationEditRequest) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableString and assigns it to the Repository field.
func (o *OrganizationEditRequest) SetRepository(v string) {
	o.Repository.Set(&v)
}

// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *OrganizationEditRequest) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *OrganizationEditRequest) UnsetRepository() {
	o.Repository.Unset()
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationEditRequest) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl.Get()) {
		var ret string
		return ret
	}
	return *o.LogoUrl.Get()
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationEditRequest) GetLogoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoUrl.Get(), o.LogoUrl.IsSet()
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *OrganizationEditRequest) HasLogoUrl() bool {
	if o != nil && o.LogoUrl.IsSet() {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given NullableString and assigns it to the LogoUrl field.
func (o *OrganizationEditRequest) SetLogoUrl(v string) {
	o.LogoUrl.Set(&v)
}

// SetLogoUrlNil sets the value for LogoUrl to be an explicit nil
func (o *OrganizationEditRequest) SetLogoUrlNil() {
	o.LogoUrl.Set(nil)
}

// UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
func (o *OrganizationEditRequest) UnsetLogoUrl() {
	o.LogoUrl.Unset()
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationEditRequest) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationEditRequest) GetIconUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *OrganizationEditRequest) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *OrganizationEditRequest) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}

// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *OrganizationEditRequest) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *OrganizationEditRequest) UnsetIconUrl() {
	o.IconUrl.Unset()
}

// GetAdminEmails returns the AdminEmails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationEditRequest) GetAdminEmails() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AdminEmails
}

// GetAdminEmailsOk returns a tuple with the AdminEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationEditRequest) GetAdminEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.AdminEmails) {
		return nil, false
	}
	return o.AdminEmails, true
}

// HasAdminEmails returns a boolean if a field has been set.
func (o *OrganizationEditRequest) HasAdminEmails() bool {
	if o != nil && !IsNil(o.AdminEmails) {
		return true
	}

	return false
}

// SetAdminEmails gets a reference to the given []string and assigns it to the AdminEmails field.
func (o *OrganizationEditRequest) SetAdminEmails(v []string) {
	o.AdminEmails = v
}

func (o OrganizationEditRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationEditRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.WebsiteUrl.IsSet() {
		toSerialize["website_url"] = o.WebsiteUrl.Get()
	}
	if o.Repository.IsSet() {
		toSerialize["repository"] = o.Repository.Get()
	}
	if o.LogoUrl.IsSet() {
		toSerialize["logo_url"] = o.LogoUrl.Get()
	}
	if o.IconUrl.IsSet() {
		toSerialize["icon_url"] = o.IconUrl.Get()
	}
	if o.AdminEmails != nil {
		toSerialize["admin_emails"] = o.AdminEmails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrganizationEditRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varOrganizationEditRequest := _OrganizationEditRequest{}

	err = json.Unmarshal(data, &varOrganizationEditRequest)

	if err != nil {
		return err
	}

	*o = OrganizationEditRequest(varOrganizationEditRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "website_url")
		delete(additionalProperties, "repository")
		delete(additionalProperties, "logo_url")
		delete(additionalProperties, "icon_url")
		delete(additionalProperties, "admin_emails")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrganizationEditRequest struct {
	value *OrganizationEditRequest
	isSet bool
}

func (v NullableOrganizationEditRequest) Get() *OrganizationEditRequest {
	return v.value
}

func (v *NullableOrganizationEditRequest) Set(val *OrganizationEditRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationEditRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationEditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationEditRequest(val *OrganizationEditRequest) *NullableOrganizationEditRequest {
	return &NullableOrganizationEditRequest{value: val, isSet: true}
}

func (v NullableOrganizationEditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationEditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
