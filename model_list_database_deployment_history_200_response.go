/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ListDatabaseDeploymentHistory200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListDatabaseDeploymentHistory200Response{}

// ListDatabaseDeploymentHistory200Response struct for ListDatabaseDeploymentHistory200Response
type ListDatabaseDeploymentHistory200Response struct {
	Page     float32                     `json:"page"`
	PageSize float32                     `json:"page_size"`
	Results  []DeploymentHistoryDatabase `json:"results,omitempty"`
}

type _ListDatabaseDeploymentHistory200Response ListDatabaseDeploymentHistory200Response

// NewListDatabaseDeploymentHistory200Response instantiates a new ListDatabaseDeploymentHistory200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDatabaseDeploymentHistory200Response(page float32, pageSize float32) *ListDatabaseDeploymentHistory200Response {
	this := ListDatabaseDeploymentHistory200Response{}
	this.Page = page
	this.PageSize = pageSize
	return &this
}

// NewListDatabaseDeploymentHistory200ResponseWithDefaults instantiates a new ListDatabaseDeploymentHistory200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDatabaseDeploymentHistory200ResponseWithDefaults() *ListDatabaseDeploymentHistory200Response {
	this := ListDatabaseDeploymentHistory200Response{}
	return &this
}

// GetPage returns the Page field value
func (o *ListDatabaseDeploymentHistory200Response) GetPage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *ListDatabaseDeploymentHistory200Response) GetPageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *ListDatabaseDeploymentHistory200Response) SetPage(v float32) {
	o.Page = v
}

// GetPageSize returns the PageSize field value
func (o *ListDatabaseDeploymentHistory200Response) GetPageSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *ListDatabaseDeploymentHistory200Response) GetPageSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *ListDatabaseDeploymentHistory200Response) SetPageSize(v float32) {
	o.PageSize = v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListDatabaseDeploymentHistory200Response) GetResults() []DeploymentHistoryDatabase {
	if o == nil || IsNil(o.Results) {
		var ret []DeploymentHistoryDatabase
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDatabaseDeploymentHistory200Response) GetResultsOk() ([]DeploymentHistoryDatabase, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListDatabaseDeploymentHistory200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DeploymentHistoryDatabase and assigns it to the Results field.
func (o *ListDatabaseDeploymentHistory200Response) SetResults(v []DeploymentHistoryDatabase) {
	o.Results = v
}

func (o ListDatabaseDeploymentHistory200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListDatabaseDeploymentHistory200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["page"] = o.Page
	toSerialize["page_size"] = o.PageSize
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

func (o *ListDatabaseDeploymentHistory200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"page",
		"page_size",
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

	varListDatabaseDeploymentHistory200Response := _ListDatabaseDeploymentHistory200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListDatabaseDeploymentHistory200Response)

	if err != nil {
		return err
	}

	*o = ListDatabaseDeploymentHistory200Response(varListDatabaseDeploymentHistory200Response)

	return err
}

type NullableListDatabaseDeploymentHistory200Response struct {
	value *ListDatabaseDeploymentHistory200Response
	isSet bool
}

func (v NullableListDatabaseDeploymentHistory200Response) Get() *ListDatabaseDeploymentHistory200Response {
	return v.value
}

func (v *NullableListDatabaseDeploymentHistory200Response) Set(val *ListDatabaseDeploymentHistory200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListDatabaseDeploymentHistory200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListDatabaseDeploymentHistory200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDatabaseDeploymentHistory200Response(val *ListDatabaseDeploymentHistory200Response) *NullableListDatabaseDeploymentHistory200Response {
	return &NullableListDatabaseDeploymentHistory200Response{value: val, isSet: true}
}

func (v NullableListDatabaseDeploymentHistory200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDatabaseDeploymentHistory200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
