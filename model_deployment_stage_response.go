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
	"time"
)

// checks if the DeploymentStageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeploymentStageResponse{}

// DeploymentStageResponse struct for DeploymentStageResponse
type DeploymentStageResponse struct {
	Id          string          `json:"id"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   *time.Time      `json:"updated_at,omitempty"`
	Environment ReferenceObject `json:"environment"`
	// name is case insensitive
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	// Position of the deployment stage within the environment
	DeploymentOrder      *int32                           `json:"deployment_order,omitempty"`
	Services             []DeploymentStageServiceResponse `json:"services,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeploymentStageResponse DeploymentStageResponse

// NewDeploymentStageResponse instantiates a new DeploymentStageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentStageResponse(id string, createdAt time.Time, environment ReferenceObject) *DeploymentStageResponse {
	this := DeploymentStageResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Environment = environment
	return &this
}

// NewDeploymentStageResponseWithDefaults instantiates a new DeploymentStageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentStageResponseWithDefaults() *DeploymentStageResponse {
	this := DeploymentStageResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DeploymentStageResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeploymentStageResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeploymentStageResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DeploymentStageResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DeploymentStageResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DeploymentStageResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DeploymentStageResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStageResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DeploymentStageResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DeploymentStageResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetEnvironment returns the Environment field value
func (o *DeploymentStageResponse) GetEnvironment() ReferenceObject {
	if o == nil {
		var ret ReferenceObject
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *DeploymentStageResponse) GetEnvironmentOk() (*ReferenceObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *DeploymentStageResponse) SetEnvironment(v ReferenceObject) {
	o.Environment = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentStageResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStageResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentStageResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentStageResponse) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeploymentStageResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStageResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeploymentStageResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeploymentStageResponse) SetDescription(v string) {
	o.Description = &v
}

// GetDeploymentOrder returns the DeploymentOrder field value if set, zero value otherwise.
func (o *DeploymentStageResponse) GetDeploymentOrder() int32 {
	if o == nil || IsNil(o.DeploymentOrder) {
		var ret int32
		return ret
	}
	return *o.DeploymentOrder
}

// GetDeploymentOrderOk returns a tuple with the DeploymentOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStageResponse) GetDeploymentOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.DeploymentOrder) {
		return nil, false
	}
	return o.DeploymentOrder, true
}

// HasDeploymentOrder returns a boolean if a field has been set.
func (o *DeploymentStageResponse) HasDeploymentOrder() bool {
	if o != nil && !IsNil(o.DeploymentOrder) {
		return true
	}

	return false
}

// SetDeploymentOrder gets a reference to the given int32 and assigns it to the DeploymentOrder field.
func (o *DeploymentStageResponse) SetDeploymentOrder(v int32) {
	o.DeploymentOrder = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *DeploymentStageResponse) GetServices() []DeploymentStageServiceResponse {
	if o == nil || IsNil(o.Services) {
		var ret []DeploymentStageServiceResponse
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStageResponse) GetServicesOk() ([]DeploymentStageServiceResponse, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *DeploymentStageResponse) HasServices() bool {
	if o != nil && !IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given []DeploymentStageServiceResponse and assigns it to the Services field.
func (o *DeploymentStageResponse) SetServices(v []DeploymentStageServiceResponse) {
	o.Services = v
}

func (o DeploymentStageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploymentStageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	toSerialize["environment"] = o.Environment
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DeploymentOrder) {
		toSerialize["deployment_order"] = o.DeploymentOrder
	}
	if !IsNil(o.Services) {
		toSerialize["services"] = o.Services
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeploymentStageResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"environment",
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

	varDeploymentStageResponse := _DeploymentStageResponse{}

	err = json.Unmarshal(data, &varDeploymentStageResponse)

	if err != nil {
		return err
	}

	*o = DeploymentStageResponse(varDeploymentStageResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "environment")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "deployment_order")
		delete(additionalProperties, "services")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeploymentStageResponse struct {
	value *DeploymentStageResponse
	isSet bool
}

func (v NullableDeploymentStageResponse) Get() *DeploymentStageResponse {
	return v.value
}

func (v *NullableDeploymentStageResponse) Set(val *DeploymentStageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentStageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentStageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentStageResponse(val *DeploymentStageResponse) *NullableDeploymentStageResponse {
	return &NullableDeploymentStageResponse{value: val, isSet: true}
}

func (v NullableDeploymentStageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentStageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
