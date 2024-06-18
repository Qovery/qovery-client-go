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

// checks if the ClusterFeatureResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterFeatureResponse{}

// ClusterFeatureResponse struct for ClusterFeatureResponse
type ClusterFeatureResponse struct {
	Id          *string        `json:"id,omitempty"`
	Title       *string        `json:"title,omitempty"`
	Description NullableString `json:"description,omitempty"`
	// Deprecated
	CostPerMonthInCents NullableInt32 `json:"cost_per_month_in_cents,omitempty"`
	// Deprecated
	CostPerMonth NullableFloat32 `json:"cost_per_month,omitempty"`
	// Deprecated
	CurrencyCode                      NullableString                              `json:"currency_code,omitempty"`
	IsCloudProviderPayingFeature      *bool                                       `json:"is_cloud_provider_paying_feature,omitempty"`
	CloudProviderFeatureDocumentation NullableString                              `json:"cloud_provider_feature_documentation,omitempty"`
	IsQoveryPayingFeature             *bool                                       `json:"is_qovery_paying_feature,omitempty"`
	QoveryFeatureDocumentation        NullableString                              `json:"qovery_feature_documentation,omitempty"`
	ValueType                         *string                                     `json:"value_type,omitempty"`
	ValueObject                       NullableClusterFeatureResponseValueObject   `json:"value_object,omitempty"`
	IsValueUpdatable                  *bool                                       `json:"is_value_updatable,omitempty"`
	AcceptedValues                    []ClusterFeatureResponseAcceptedValuesInner `json:"accepted_values,omitempty"`
	AdditionalProperties              map[string]interface{}
}

type _ClusterFeatureResponse ClusterFeatureResponse

// NewClusterFeatureResponse instantiates a new ClusterFeatureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterFeatureResponse() *ClusterFeatureResponse {
	this := ClusterFeatureResponse{}
	var isValueUpdatable bool = false
	this.IsValueUpdatable = &isValueUpdatable
	return &this
}

// NewClusterFeatureResponseWithDefaults instantiates a new ClusterFeatureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterFeatureResponseWithDefaults() *ClusterFeatureResponse {
	this := ClusterFeatureResponse{}
	var isValueUpdatable bool = false
	this.IsValueUpdatable = &isValueUpdatable
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClusterFeatureResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFeatureResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterFeatureResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClusterFeatureResponse) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ClusterFeatureResponse) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFeatureResponse) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ClusterFeatureResponse) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ClusterFeatureResponse) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeatureResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeatureResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ClusterFeatureResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ClusterFeatureResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ClusterFeatureResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ClusterFeatureResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetCostPerMonthInCents returns the CostPerMonthInCents field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *ClusterFeatureResponse) GetCostPerMonthInCents() int32 {
	if o == nil || IsNil(o.CostPerMonthInCents.Get()) {
		var ret int32
		return ret
	}
	return *o.CostPerMonthInCents.Get()
}

// GetCostPerMonthInCentsOk returns a tuple with the CostPerMonthInCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *ClusterFeatureResponse) GetCostPerMonthInCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CostPerMonthInCents.Get(), o.CostPerMonthInCents.IsSet()
}

// HasCostPerMonthInCents returns a boolean if a field has been set.
func (o *ClusterFeatureResponse) HasCostPerMonthInCents() bool {
	if o != nil && o.CostPerMonthInCents.IsSet() {
		return true
	}

	return false
}

// SetCostPerMonthInCents gets a reference to the given NullableInt32 and assigns it to the CostPerMonthInCents field.
// Deprecated
func (o *ClusterFeatureResponse) SetCostPerMonthInCents(v int32) {
	o.CostPerMonthInCents.Set(&v)
}

// SetCostPerMonthInCentsNil sets the value for CostPerMonthInCents to be an explicit nil
func (o *ClusterFeatureResponse) SetCostPerMonthInCentsNil() {
	o.CostPerMonthInCents.Set(nil)
}

// UnsetCostPerMonthInCents ensures that no value is present for CostPerMonthInCents, not even an explicit nil
func (o *ClusterFeatureResponse) UnsetCostPerMonthInCents() {
	o.CostPerMonthInCents.Unset()
}

// GetCostPerMonth returns the CostPerMonth field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *ClusterFeatureResponse) GetCostPerMonth() float32 {
	if o == nil || IsNil(o.CostPerMonth.Get()) {
		var ret float32
		return ret
	}
	return *o.CostPerMonth.Get()
}

// GetCostPerMonthOk returns a tuple with the CostPerMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *ClusterFeatureResponse) GetCostPerMonthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CostPerMonth.Get(), o.CostPerMonth.IsSet()
}

// HasCostPerMonth returns a boolean if a field has been set.
func (o *ClusterFeatureResponse) HasCostPerMonth() bool {
	if o != nil && o.CostPerMonth.IsSet() {
		return true
	}

	return false
}

// SetCostPerMonth gets a reference to the given NullableFloat32 and assigns it to the CostPerMonth field.
// Deprecated
func (o *ClusterFeatureResponse) SetCostPerMonth(v float32) {
	o.CostPerMonth.Set(&v)
}

// SetCostPerMonthNil sets the value for CostPerMonth to be an explicit nil
func (o *ClusterFeatureResponse) SetCostPerMonthNil() {
	o.CostPerMonth.Set(nil)
}

// UnsetCostPerMonth ensures that no value is present for CostPerMonth, not even an explicit nil
func (o *ClusterFeatureResponse) UnsetCostPerMonth() {
	o.CostPerMonth.Unset()
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *ClusterFeatureResponse) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode.Get()) {
		var ret string
		return ret
	}
	return *o.CurrencyCode.Get()
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *ClusterFeatureResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrencyCode.Get(), o.CurrencyCode.IsSet()
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *ClusterFeatureResponse) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given NullableString and assigns it to the CurrencyCode field.
// Deprecated
func (o *ClusterFeatureResponse) SetCurrencyCode(v string) {
	o.CurrencyCode.Set(&v)
}

// SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil
func (o *ClusterFeatureResponse) SetCurrencyCodeNil() {
	o.CurrencyCode.Set(nil)
}

// UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
func (o *ClusterFeatureResponse) UnsetCurrencyCode() {
	o.CurrencyCode.Unset()
}

// GetIsCloudProviderPayingFeature returns the IsCloudProviderPayingFeature field value if set, zero value otherwise.
func (o *ClusterFeatureResponse) GetIsCloudProviderPayingFeature() bool {
	if o == nil || IsNil(o.IsCloudProviderPayingFeature) {
		var ret bool
		return ret
	}
	return *o.IsCloudProviderPayingFeature
}

// GetIsCloudProviderPayingFeatureOk returns a tuple with the IsCloudProviderPayingFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFeatureResponse) GetIsCloudProviderPayingFeatureOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCloudProviderPayingFeature) {
		return nil, false
	}
	return o.IsCloudProviderPayingFeature, true
}

// HasIsCloudProviderPayingFeature returns a boolean if a field has been set.
func (o *ClusterFeatureResponse) HasIsCloudProviderPayingFeature() bool {
	if o != nil && !IsNil(o.IsCloudProviderPayingFeature) {
		return true
	}

	return false
}

// SetIsCloudProviderPayingFeature gets a reference to the given bool and assigns it to the IsCloudProviderPayingFeature field.
func (o *ClusterFeatureResponse) SetIsCloudProviderPayingFeature(v bool) {
	o.IsCloudProviderPayingFeature = &v
}

// GetCloudProviderFeatureDocumentation returns the CloudProviderFeatureDocumentation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeatureResponse) GetCloudProviderFeatureDocumentation() string {
	if o == nil || IsNil(o.CloudProviderFeatureDocumentation.Get()) {
		var ret string
		return ret
	}
	return *o.CloudProviderFeatureDocumentation.Get()
}

// GetCloudProviderFeatureDocumentationOk returns a tuple with the CloudProviderFeatureDocumentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeatureResponse) GetCloudProviderFeatureDocumentationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CloudProviderFeatureDocumentation.Get(), o.CloudProviderFeatureDocumentation.IsSet()
}

// HasCloudProviderFeatureDocumentation returns a boolean if a field has been set.
func (o *ClusterFeatureResponse) HasCloudProviderFeatureDocumentation() bool {
	if o != nil && o.CloudProviderFeatureDocumentation.IsSet() {
		return true
	}

	return false
}

// SetCloudProviderFeatureDocumentation gets a reference to the given NullableString and assigns it to the CloudProviderFeatureDocumentation field.
func (o *ClusterFeatureResponse) SetCloudProviderFeatureDocumentation(v string) {
	o.CloudProviderFeatureDocumentation.Set(&v)
}

// SetCloudProviderFeatureDocumentationNil sets the value for CloudProviderFeatureDocumentation to be an explicit nil
func (o *ClusterFeatureResponse) SetCloudProviderFeatureDocumentationNil() {
	o.CloudProviderFeatureDocumentation.Set(nil)
}

// UnsetCloudProviderFeatureDocumentation ensures that no value is present for CloudProviderFeatureDocumentation, not even an explicit nil
func (o *ClusterFeatureResponse) UnsetCloudProviderFeatureDocumentation() {
	o.CloudProviderFeatureDocumentation.Unset()
}

// GetIsQoveryPayingFeature returns the IsQoveryPayingFeature field value if set, zero value otherwise.
func (o *ClusterFeatureResponse) GetIsQoveryPayingFeature() bool {
	if o == nil || IsNil(o.IsQoveryPayingFeature) {
		var ret bool
		return ret
	}
	return *o.IsQoveryPayingFeature
}

// GetIsQoveryPayingFeatureOk returns a tuple with the IsQoveryPayingFeature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFeatureResponse) GetIsQoveryPayingFeatureOk() (*bool, bool) {
	if o == nil || IsNil(o.IsQoveryPayingFeature) {
		return nil, false
	}
	return o.IsQoveryPayingFeature, true
}

// HasIsQoveryPayingFeature returns a boolean if a field has been set.
func (o *ClusterFeatureResponse) HasIsQoveryPayingFeature() bool {
	if o != nil && !IsNil(o.IsQoveryPayingFeature) {
		return true
	}

	return false
}

// SetIsQoveryPayingFeature gets a reference to the given bool and assigns it to the IsQoveryPayingFeature field.
func (o *ClusterFeatureResponse) SetIsQoveryPayingFeature(v bool) {
	o.IsQoveryPayingFeature = &v
}

// GetQoveryFeatureDocumentation returns the QoveryFeatureDocumentation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeatureResponse) GetQoveryFeatureDocumentation() string {
	if o == nil || IsNil(o.QoveryFeatureDocumentation.Get()) {
		var ret string
		return ret
	}
	return *o.QoveryFeatureDocumentation.Get()
}

// GetQoveryFeatureDocumentationOk returns a tuple with the QoveryFeatureDocumentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeatureResponse) GetQoveryFeatureDocumentationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QoveryFeatureDocumentation.Get(), o.QoveryFeatureDocumentation.IsSet()
}

// HasQoveryFeatureDocumentation returns a boolean if a field has been set.
func (o *ClusterFeatureResponse) HasQoveryFeatureDocumentation() bool {
	if o != nil && o.QoveryFeatureDocumentation.IsSet() {
		return true
	}

	return false
}

// SetQoveryFeatureDocumentation gets a reference to the given NullableString and assigns it to the QoveryFeatureDocumentation field.
func (o *ClusterFeatureResponse) SetQoveryFeatureDocumentation(v string) {
	o.QoveryFeatureDocumentation.Set(&v)
}

// SetQoveryFeatureDocumentationNil sets the value for QoveryFeatureDocumentation to be an explicit nil
func (o *ClusterFeatureResponse) SetQoveryFeatureDocumentationNil() {
	o.QoveryFeatureDocumentation.Set(nil)
}

// UnsetQoveryFeatureDocumentation ensures that no value is present for QoveryFeatureDocumentation, not even an explicit nil
func (o *ClusterFeatureResponse) UnsetQoveryFeatureDocumentation() {
	o.QoveryFeatureDocumentation.Unset()
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *ClusterFeatureResponse) GetValueType() string {
	if o == nil || IsNil(o.ValueType) {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFeatureResponse) GetValueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *ClusterFeatureResponse) HasValueType() bool {
	if o != nil && !IsNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *ClusterFeatureResponse) SetValueType(v string) {
	o.ValueType = &v
}

// GetValueObject returns the ValueObject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeatureResponse) GetValueObject() ClusterFeatureResponseValueObject {
	if o == nil || IsNil(o.ValueObject.Get()) {
		var ret ClusterFeatureResponseValueObject
		return ret
	}
	return *o.ValueObject.Get()
}

// GetValueObjectOk returns a tuple with the ValueObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeatureResponse) GetValueObjectOk() (*ClusterFeatureResponseValueObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValueObject.Get(), o.ValueObject.IsSet()
}

// HasValueObject returns a boolean if a field has been set.
func (o *ClusterFeatureResponse) HasValueObject() bool {
	if o != nil && o.ValueObject.IsSet() {
		return true
	}

	return false
}

// SetValueObject gets a reference to the given NullableClusterFeatureResponseValueObject and assigns it to the ValueObject field.
func (o *ClusterFeatureResponse) SetValueObject(v ClusterFeatureResponseValueObject) {
	o.ValueObject.Set(&v)
}

// SetValueObjectNil sets the value for ValueObject to be an explicit nil
func (o *ClusterFeatureResponse) SetValueObjectNil() {
	o.ValueObject.Set(nil)
}

// UnsetValueObject ensures that no value is present for ValueObject, not even an explicit nil
func (o *ClusterFeatureResponse) UnsetValueObject() {
	o.ValueObject.Unset()
}

// GetIsValueUpdatable returns the IsValueUpdatable field value if set, zero value otherwise.
func (o *ClusterFeatureResponse) GetIsValueUpdatable() bool {
	if o == nil || IsNil(o.IsValueUpdatable) {
		var ret bool
		return ret
	}
	return *o.IsValueUpdatable
}

// GetIsValueUpdatableOk returns a tuple with the IsValueUpdatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFeatureResponse) GetIsValueUpdatableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsValueUpdatable) {
		return nil, false
	}
	return o.IsValueUpdatable, true
}

// HasIsValueUpdatable returns a boolean if a field has been set.
func (o *ClusterFeatureResponse) HasIsValueUpdatable() bool {
	if o != nil && !IsNil(o.IsValueUpdatable) {
		return true
	}

	return false
}

// SetIsValueUpdatable gets a reference to the given bool and assigns it to the IsValueUpdatable field.
func (o *ClusterFeatureResponse) SetIsValueUpdatable(v bool) {
	o.IsValueUpdatable = &v
}

// GetAcceptedValues returns the AcceptedValues field value if set, zero value otherwise.
func (o *ClusterFeatureResponse) GetAcceptedValues() []ClusterFeatureResponseAcceptedValuesInner {
	if o == nil || IsNil(o.AcceptedValues) {
		var ret []ClusterFeatureResponseAcceptedValuesInner
		return ret
	}
	return o.AcceptedValues
}

// GetAcceptedValuesOk returns a tuple with the AcceptedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterFeatureResponse) GetAcceptedValuesOk() ([]ClusterFeatureResponseAcceptedValuesInner, bool) {
	if o == nil || IsNil(o.AcceptedValues) {
		return nil, false
	}
	return o.AcceptedValues, true
}

// HasAcceptedValues returns a boolean if a field has been set.
func (o *ClusterFeatureResponse) HasAcceptedValues() bool {
	if o != nil && !IsNil(o.AcceptedValues) {
		return true
	}

	return false
}

// SetAcceptedValues gets a reference to the given []ClusterFeatureResponseAcceptedValuesInner and assigns it to the AcceptedValues field.
func (o *ClusterFeatureResponse) SetAcceptedValues(v []ClusterFeatureResponseAcceptedValuesInner) {
	o.AcceptedValues = v
}

func (o ClusterFeatureResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterFeatureResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.CostPerMonthInCents.IsSet() {
		toSerialize["cost_per_month_in_cents"] = o.CostPerMonthInCents.Get()
	}
	if o.CostPerMonth.IsSet() {
		toSerialize["cost_per_month"] = o.CostPerMonth.Get()
	}
	if o.CurrencyCode.IsSet() {
		toSerialize["currency_code"] = o.CurrencyCode.Get()
	}
	if !IsNil(o.IsCloudProviderPayingFeature) {
		toSerialize["is_cloud_provider_paying_feature"] = o.IsCloudProviderPayingFeature
	}
	if o.CloudProviderFeatureDocumentation.IsSet() {
		toSerialize["cloud_provider_feature_documentation"] = o.CloudProviderFeatureDocumentation.Get()
	}
	if !IsNil(o.IsQoveryPayingFeature) {
		toSerialize["is_qovery_paying_feature"] = o.IsQoveryPayingFeature
	}
	if o.QoveryFeatureDocumentation.IsSet() {
		toSerialize["qovery_feature_documentation"] = o.QoveryFeatureDocumentation.Get()
	}
	if !IsNil(o.ValueType) {
		toSerialize["value_type"] = o.ValueType
	}
	if o.ValueObject.IsSet() {
		toSerialize["value_object"] = o.ValueObject.Get()
	}
	if !IsNil(o.IsValueUpdatable) {
		toSerialize["is_value_updatable"] = o.IsValueUpdatable
	}
	if !IsNil(o.AcceptedValues) {
		toSerialize["accepted_values"] = o.AcceptedValues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClusterFeatureResponse) UnmarshalJSON(data []byte) (err error) {
	varClusterFeatureResponse := _ClusterFeatureResponse{}

	err = json.Unmarshal(data, &varClusterFeatureResponse)

	if err != nil {
		return err
	}

	*o = ClusterFeatureResponse(varClusterFeatureResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "title")
		delete(additionalProperties, "description")
		delete(additionalProperties, "cost_per_month_in_cents")
		delete(additionalProperties, "cost_per_month")
		delete(additionalProperties, "currency_code")
		delete(additionalProperties, "is_cloud_provider_paying_feature")
		delete(additionalProperties, "cloud_provider_feature_documentation")
		delete(additionalProperties, "is_qovery_paying_feature")
		delete(additionalProperties, "qovery_feature_documentation")
		delete(additionalProperties, "value_type")
		delete(additionalProperties, "value_object")
		delete(additionalProperties, "is_value_updatable")
		delete(additionalProperties, "accepted_values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClusterFeatureResponse struct {
	value *ClusterFeatureResponse
	isSet bool
}

func (v NullableClusterFeatureResponse) Get() *ClusterFeatureResponse {
	return v.value
}

func (v *NullableClusterFeatureResponse) Set(val *ClusterFeatureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterFeatureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterFeatureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterFeatureResponse(val *ClusterFeatureResponse) *NullableClusterFeatureResponse {
	return &NullableClusterFeatureResponse{value: val, isSet: true}
}

func (v NullableClusterFeatureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterFeatureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}