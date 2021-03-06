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

// ServiceAllOf struct for ServiceAllOf
type ServiceAllOf struct {
	Type *ServiceTypeEnum `json:"type,omitempty"`
	// name of the service
	Name *string `json:"name,omitempty"`
	// uuid of the associated service (application, database, job, gateway...)
	Id string `json:"id"`
	// Git commit ID corresponding to the deployed version of the application
	DeployedCommitId *string `json:"deployed_commit_id,omitempty"`
	// uuid of the user that made the last update
	LastUpdatedBy *string `json:"last_updated_by,omitempty"`
	// global overview of resources consumption of the service
	ConsumedResourcesInPercent *float32 `json:"consumed_resources_in_percent,omitempty"`
	// describes the typology of service (container, postgresl, redis...)
	ServiceTypology *string `json:"service_typology,omitempty"`
	// for databases this field exposes the database version
	ServiceVersion *string `json:"service_version,omitempty"`
	ToUpdate       *bool   `json:"to_update,omitempty"`
}

// NewServiceAllOf instantiates a new ServiceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAllOf(id string) *ServiceAllOf {
	this := ServiceAllOf{}
	this.Id = id
	return &this
}

// NewServiceAllOfWithDefaults instantiates a new ServiceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAllOfWithDefaults() *ServiceAllOf {
	this := ServiceAllOf{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceAllOf) GetType() ServiceTypeEnum {
	if o == nil || o.Type == nil {
		var ret ServiceTypeEnum
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAllOf) GetTypeOk() (*ServiceTypeEnum, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given ServiceTypeEnum and assigns it to the Type field.
func (o *ServiceAllOf) SetType(v ServiceTypeEnum) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceAllOf) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value
func (o *ServiceAllOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceAllOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceAllOf) SetId(v string) {
	o.Id = v
}

// GetDeployedCommitId returns the DeployedCommitId field value if set, zero value otherwise.
func (o *ServiceAllOf) GetDeployedCommitId() string {
	if o == nil || o.DeployedCommitId == nil {
		var ret string
		return ret
	}
	return *o.DeployedCommitId
}

// GetDeployedCommitIdOk returns a tuple with the DeployedCommitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAllOf) GetDeployedCommitIdOk() (*string, bool) {
	if o == nil || o.DeployedCommitId == nil {
		return nil, false
	}
	return o.DeployedCommitId, true
}

// HasDeployedCommitId returns a boolean if a field has been set.
func (o *ServiceAllOf) HasDeployedCommitId() bool {
	if o != nil && o.DeployedCommitId != nil {
		return true
	}

	return false
}

// SetDeployedCommitId gets a reference to the given string and assigns it to the DeployedCommitId field.
func (o *ServiceAllOf) SetDeployedCommitId(v string) {
	o.DeployedCommitId = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *ServiceAllOf) GetLastUpdatedBy() string {
	if o == nil || o.LastUpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAllOf) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || o.LastUpdatedBy == nil {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *ServiceAllOf) HasLastUpdatedBy() bool {
	if o != nil && o.LastUpdatedBy != nil {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *ServiceAllOf) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetConsumedResourcesInPercent returns the ConsumedResourcesInPercent field value if set, zero value otherwise.
func (o *ServiceAllOf) GetConsumedResourcesInPercent() float32 {
	if o == nil || o.ConsumedResourcesInPercent == nil {
		var ret float32
		return ret
	}
	return *o.ConsumedResourcesInPercent
}

// GetConsumedResourcesInPercentOk returns a tuple with the ConsumedResourcesInPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAllOf) GetConsumedResourcesInPercentOk() (*float32, bool) {
	if o == nil || o.ConsumedResourcesInPercent == nil {
		return nil, false
	}
	return o.ConsumedResourcesInPercent, true
}

// HasConsumedResourcesInPercent returns a boolean if a field has been set.
func (o *ServiceAllOf) HasConsumedResourcesInPercent() bool {
	if o != nil && o.ConsumedResourcesInPercent != nil {
		return true
	}

	return false
}

// SetConsumedResourcesInPercent gets a reference to the given float32 and assigns it to the ConsumedResourcesInPercent field.
func (o *ServiceAllOf) SetConsumedResourcesInPercent(v float32) {
	o.ConsumedResourcesInPercent = &v
}

// GetServiceTypology returns the ServiceTypology field value if set, zero value otherwise.
func (o *ServiceAllOf) GetServiceTypology() string {
	if o == nil || o.ServiceTypology == nil {
		var ret string
		return ret
	}
	return *o.ServiceTypology
}

// GetServiceTypologyOk returns a tuple with the ServiceTypology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAllOf) GetServiceTypologyOk() (*string, bool) {
	if o == nil || o.ServiceTypology == nil {
		return nil, false
	}
	return o.ServiceTypology, true
}

// HasServiceTypology returns a boolean if a field has been set.
func (o *ServiceAllOf) HasServiceTypology() bool {
	if o != nil && o.ServiceTypology != nil {
		return true
	}

	return false
}

// SetServiceTypology gets a reference to the given string and assigns it to the ServiceTypology field.
func (o *ServiceAllOf) SetServiceTypology(v string) {
	o.ServiceTypology = &v
}

// GetServiceVersion returns the ServiceVersion field value if set, zero value otherwise.
func (o *ServiceAllOf) GetServiceVersion() string {
	if o == nil || o.ServiceVersion == nil {
		var ret string
		return ret
	}
	return *o.ServiceVersion
}

// GetServiceVersionOk returns a tuple with the ServiceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAllOf) GetServiceVersionOk() (*string, bool) {
	if o == nil || o.ServiceVersion == nil {
		return nil, false
	}
	return o.ServiceVersion, true
}

// HasServiceVersion returns a boolean if a field has been set.
func (o *ServiceAllOf) HasServiceVersion() bool {
	if o != nil && o.ServiceVersion != nil {
		return true
	}

	return false
}

// SetServiceVersion gets a reference to the given string and assigns it to the ServiceVersion field.
func (o *ServiceAllOf) SetServiceVersion(v string) {
	o.ServiceVersion = &v
}

// GetToUpdate returns the ToUpdate field value if set, zero value otherwise.
func (o *ServiceAllOf) GetToUpdate() bool {
	if o == nil || o.ToUpdate == nil {
		var ret bool
		return ret
	}
	return *o.ToUpdate
}

// GetToUpdateOk returns a tuple with the ToUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAllOf) GetToUpdateOk() (*bool, bool) {
	if o == nil || o.ToUpdate == nil {
		return nil, false
	}
	return o.ToUpdate, true
}

// HasToUpdate returns a boolean if a field has been set.
func (o *ServiceAllOf) HasToUpdate() bool {
	if o != nil && o.ToUpdate != nil {
		return true
	}

	return false
}

// SetToUpdate gets a reference to the given bool and assigns it to the ToUpdate field.
func (o *ServiceAllOf) SetToUpdate(v bool) {
	o.ToUpdate = &v
}

func (o ServiceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.DeployedCommitId != nil {
		toSerialize["deployed_commit_id"] = o.DeployedCommitId
	}
	if o.LastUpdatedBy != nil {
		toSerialize["last_updated_by"] = o.LastUpdatedBy
	}
	if o.ConsumedResourcesInPercent != nil {
		toSerialize["consumed_resources_in_percent"] = o.ConsumedResourcesInPercent
	}
	if o.ServiceTypology != nil {
		toSerialize["service_typology"] = o.ServiceTypology
	}
	if o.ServiceVersion != nil {
		toSerialize["service_version"] = o.ServiceVersion
	}
	if o.ToUpdate != nil {
		toSerialize["to_update"] = o.ToUpdate
	}
	return json.Marshal(toSerialize)
}

type NullableServiceAllOf struct {
	value *ServiceAllOf
	isSet bool
}

func (v NullableServiceAllOf) Get() *ServiceAllOf {
	return v.value
}

func (v *NullableServiceAllOf) Set(val *ServiceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAllOf(val *ServiceAllOf) *NullableServiceAllOf {
	return &NullableServiceAllOf{value: val, isSet: true}
}

func (v NullableServiceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
