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
	"time"
)

// checks if the AccountInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountInfo{}

// AccountInfo struct for AccountInfo
type AccountInfo struct {
	Id                   *string    `json:"id,omitempty"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	Nickname             *string    `json:"nickname,omitempty"`
	FirstName            *string    `json:"first_name,omitempty"`
	LastName             *string    `json:"last_name,omitempty"`
	ProfilePictureUrl    *string    `json:"profile_picture_url,omitempty"`
	CommunicationEmail   *string    `json:"communication_email,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountInfo AccountInfo

// NewAccountInfo instantiates a new AccountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInfo() *AccountInfo {
	this := AccountInfo{}
	return &this
}

// NewAccountInfoWithDefaults instantiates a new AccountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInfoWithDefaults() *AccountInfo {
	this := AccountInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountInfo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountInfo) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AccountInfo) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AccountInfo) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AccountInfo) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetNickname returns the Nickname field value if set, zero value otherwise.
func (o *AccountInfo) GetNickname() string {
	if o == nil || IsNil(o.Nickname) {
		var ret string
		return ret
	}
	return *o.Nickname
}

// GetNicknameOk returns a tuple with the Nickname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetNicknameOk() (*string, bool) {
	if o == nil || IsNil(o.Nickname) {
		return nil, false
	}
	return o.Nickname, true
}

// HasNickname returns a boolean if a field has been set.
func (o *AccountInfo) HasNickname() bool {
	if o != nil && !IsNil(o.Nickname) {
		return true
	}

	return false
}

// SetNickname gets a reference to the given string and assigns it to the Nickname field.
func (o *AccountInfo) SetNickname(v string) {
	o.Nickname = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *AccountInfo) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *AccountInfo) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *AccountInfo) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *AccountInfo) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *AccountInfo) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *AccountInfo) SetLastName(v string) {
	o.LastName = &v
}

// GetProfilePictureUrl returns the ProfilePictureUrl field value if set, zero value otherwise.
func (o *AccountInfo) GetProfilePictureUrl() string {
	if o == nil || IsNil(o.ProfilePictureUrl) {
		var ret string
		return ret
	}
	return *o.ProfilePictureUrl
}

// GetProfilePictureUrlOk returns a tuple with the ProfilePictureUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetProfilePictureUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ProfilePictureUrl) {
		return nil, false
	}
	return o.ProfilePictureUrl, true
}

// HasProfilePictureUrl returns a boolean if a field has been set.
func (o *AccountInfo) HasProfilePictureUrl() bool {
	if o != nil && !IsNil(o.ProfilePictureUrl) {
		return true
	}

	return false
}

// SetProfilePictureUrl gets a reference to the given string and assigns it to the ProfilePictureUrl field.
func (o *AccountInfo) SetProfilePictureUrl(v string) {
	o.ProfilePictureUrl = &v
}

// GetCommunicationEmail returns the CommunicationEmail field value if set, zero value otherwise.
func (o *AccountInfo) GetCommunicationEmail() string {
	if o == nil || IsNil(o.CommunicationEmail) {
		var ret string
		return ret
	}
	return *o.CommunicationEmail
}

// GetCommunicationEmailOk returns a tuple with the CommunicationEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetCommunicationEmailOk() (*string, bool) {
	if o == nil || IsNil(o.CommunicationEmail) {
		return nil, false
	}
	return o.CommunicationEmail, true
}

// HasCommunicationEmail returns a boolean if a field has been set.
func (o *AccountInfo) HasCommunicationEmail() bool {
	if o != nil && !IsNil(o.CommunicationEmail) {
		return true
	}

	return false
}

// SetCommunicationEmail gets a reference to the given string and assigns it to the CommunicationEmail field.
func (o *AccountInfo) SetCommunicationEmail(v string) {
	o.CommunicationEmail = &v
}

func (o AccountInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Nickname) {
		toSerialize["nickname"] = o.Nickname
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.ProfilePictureUrl) {
		toSerialize["profile_picture_url"] = o.ProfilePictureUrl
	}
	if !IsNil(o.CommunicationEmail) {
		toSerialize["communication_email"] = o.CommunicationEmail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountInfo) UnmarshalJSON(data []byte) (err error) {
	varAccountInfo := _AccountInfo{}

	err = json.Unmarshal(data, &varAccountInfo)

	if err != nil {
		return err
	}

	*o = AccountInfo(varAccountInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "nickname")
		delete(additionalProperties, "first_name")
		delete(additionalProperties, "last_name")
		delete(additionalProperties, "profile_picture_url")
		delete(additionalProperties, "communication_email")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountInfo struct {
	value *AccountInfo
	isSet bool
}

func (v NullableAccountInfo) Get() *AccountInfo {
	return v.value
}

func (v *NullableAccountInfo) Set(val *AccountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInfo(val *AccountInfo) *NullableAccountInfo {
	return &NullableAccountInfo{value: val, isSet: true}
}

func (v NullableAccountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
