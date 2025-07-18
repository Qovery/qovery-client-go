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

// checks if the ClusterFeatureAwsExistingVpc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterFeatureAwsExistingVpc{}

// ClusterFeatureAwsExistingVpc struct for ClusterFeatureAwsExistingVpc
type ClusterFeatureAwsExistingVpc struct {
	AwsVpcEksId                        string   `json:"aws_vpc_eks_id"`
	EksSubnetsZoneAIds                 []string `json:"eks_subnets_zone_a_ids"`
	EksSubnetsZoneBIds                 []string `json:"eks_subnets_zone_b_ids"`
	EksSubnetsZoneCIds                 []string `json:"eks_subnets_zone_c_ids"`
	DocumentdbSubnetsZoneAIds          []string `json:"documentdb_subnets_zone_a_ids,omitempty"`
	DocumentdbSubnetsZoneBIds          []string `json:"documentdb_subnets_zone_b_ids,omitempty"`
	DocumentdbSubnetsZoneCIds          []string `json:"documentdb_subnets_zone_c_ids,omitempty"`
	ElasticacheSubnetsZoneAIds         []string `json:"elasticache_subnets_zone_a_ids,omitempty"`
	ElasticacheSubnetsZoneBIds         []string `json:"elasticache_subnets_zone_b_ids,omitempty"`
	ElasticacheSubnetsZoneCIds         []string `json:"elasticache_subnets_zone_c_ids,omitempty"`
	RdsSubnetsZoneAIds                 []string `json:"rds_subnets_zone_a_ids,omitempty"`
	RdsSubnetsZoneBIds                 []string `json:"rds_subnets_zone_b_ids,omitempty"`
	RdsSubnetsZoneCIds                 []string `json:"rds_subnets_zone_c_ids,omitempty"`
	EksKarpenterFargateSubnetsZoneAIds []string `json:"eks_karpenter_fargate_subnets_zone_a_ids,omitempty"`
	EksKarpenterFargateSubnetsZoneBIds []string `json:"eks_karpenter_fargate_subnets_zone_b_ids,omitempty"`
	EksKarpenterFargateSubnetsZoneCIds []string `json:"eks_karpenter_fargate_subnets_zone_c_ids,omitempty"`
	AdditionalProperties               map[string]interface{}
}

type _ClusterFeatureAwsExistingVpc ClusterFeatureAwsExistingVpc

// NewClusterFeatureAwsExistingVpc instantiates a new ClusterFeatureAwsExistingVpc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterFeatureAwsExistingVpc(awsVpcEksId string, eksSubnetsZoneAIds []string, eksSubnetsZoneBIds []string, eksSubnetsZoneCIds []string) *ClusterFeatureAwsExistingVpc {
	this := ClusterFeatureAwsExistingVpc{}
	this.AwsVpcEksId = awsVpcEksId
	this.EksSubnetsZoneAIds = eksSubnetsZoneAIds
	this.EksSubnetsZoneBIds = eksSubnetsZoneBIds
	this.EksSubnetsZoneCIds = eksSubnetsZoneCIds
	return &this
}

// NewClusterFeatureAwsExistingVpcWithDefaults instantiates a new ClusterFeatureAwsExistingVpc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterFeatureAwsExistingVpcWithDefaults() *ClusterFeatureAwsExistingVpc {
	this := ClusterFeatureAwsExistingVpc{}
	return &this
}

// GetAwsVpcEksId returns the AwsVpcEksId field value
func (o *ClusterFeatureAwsExistingVpc) GetAwsVpcEksId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AwsVpcEksId
}

// GetAwsVpcEksIdOk returns a tuple with the AwsVpcEksId field value
// and a boolean to check if the value has been set.
func (o *ClusterFeatureAwsExistingVpc) GetAwsVpcEksIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsVpcEksId, true
}

// SetAwsVpcEksId sets field value
func (o *ClusterFeatureAwsExistingVpc) SetAwsVpcEksId(v string) {
	o.AwsVpcEksId = v
}

// GetEksSubnetsZoneAIds returns the EksSubnetsZoneAIds field value
func (o *ClusterFeatureAwsExistingVpc) GetEksSubnetsZoneAIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EksSubnetsZoneAIds
}

// GetEksSubnetsZoneAIdsOk returns a tuple with the EksSubnetsZoneAIds field value
// and a boolean to check if the value has been set.
func (o *ClusterFeatureAwsExistingVpc) GetEksSubnetsZoneAIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EksSubnetsZoneAIds, true
}

// SetEksSubnetsZoneAIds sets field value
func (o *ClusterFeatureAwsExistingVpc) SetEksSubnetsZoneAIds(v []string) {
	o.EksSubnetsZoneAIds = v
}

// GetEksSubnetsZoneBIds returns the EksSubnetsZoneBIds field value
func (o *ClusterFeatureAwsExistingVpc) GetEksSubnetsZoneBIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EksSubnetsZoneBIds
}

// GetEksSubnetsZoneBIdsOk returns a tuple with the EksSubnetsZoneBIds field value
// and a boolean to check if the value has been set.
func (o *ClusterFeatureAwsExistingVpc) GetEksSubnetsZoneBIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EksSubnetsZoneBIds, true
}

// SetEksSubnetsZoneBIds sets field value
func (o *ClusterFeatureAwsExistingVpc) SetEksSubnetsZoneBIds(v []string) {
	o.EksSubnetsZoneBIds = v
}

// GetEksSubnetsZoneCIds returns the EksSubnetsZoneCIds field value
func (o *ClusterFeatureAwsExistingVpc) GetEksSubnetsZoneCIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EksSubnetsZoneCIds
}

// GetEksSubnetsZoneCIdsOk returns a tuple with the EksSubnetsZoneCIds field value
// and a boolean to check if the value has been set.
func (o *ClusterFeatureAwsExistingVpc) GetEksSubnetsZoneCIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EksSubnetsZoneCIds, true
}

// SetEksSubnetsZoneCIds sets field value
func (o *ClusterFeatureAwsExistingVpc) SetEksSubnetsZoneCIds(v []string) {
	o.EksSubnetsZoneCIds = v
}

// GetDocumentdbSubnetsZoneAIds returns the DocumentdbSubnetsZoneAIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeatureAwsExistingVpc) GetDocumentdbSubnetsZoneAIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DocumentdbSubnetsZoneAIds
}

// GetDocumentdbSubnetsZoneAIdsOk returns a tuple with the DocumentdbSubnetsZoneAIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeatureAwsExistingVpc) GetDocumentdbSubnetsZoneAIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.DocumentdbSubnetsZoneAIds) {
		return nil, false
	}
	return o.DocumentdbSubnetsZoneAIds, true
}

// HasDocumentdbSubnetsZoneAIds returns a boolean if a field has been set.
func (o *ClusterFeatureAwsExistingVpc) HasDocumentdbSubnetsZoneAIds() bool {
	if o != nil && !IsNil(o.DocumentdbSubnetsZoneAIds) {
		return true
	}

	return false
}

// SetDocumentdbSubnetsZoneAIds gets a reference to the given []string and assigns it to the DocumentdbSubnetsZoneAIds field.
func (o *ClusterFeatureAwsExistingVpc) SetDocumentdbSubnetsZoneAIds(v []string) {
	o.DocumentdbSubnetsZoneAIds = v
}

// GetDocumentdbSubnetsZoneBIds returns the DocumentdbSubnetsZoneBIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeatureAwsExistingVpc) GetDocumentdbSubnetsZoneBIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DocumentdbSubnetsZoneBIds
}

// GetDocumentdbSubnetsZoneBIdsOk returns a tuple with the DocumentdbSubnetsZoneBIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeatureAwsExistingVpc) GetDocumentdbSubnetsZoneBIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.DocumentdbSubnetsZoneBIds) {
		return nil, false
	}
	return o.DocumentdbSubnetsZoneBIds, true
}

// HasDocumentdbSubnetsZoneBIds returns a boolean if a field has been set.
func (o *ClusterFeatureAwsExistingVpc) HasDocumentdbSubnetsZoneBIds() bool {
	if o != nil && !IsNil(o.DocumentdbSubnetsZoneBIds) {
		return true
	}

	return false
}

// SetDocumentdbSubnetsZoneBIds gets a reference to the given []string and assigns it to the DocumentdbSubnetsZoneBIds field.
func (o *ClusterFeatureAwsExistingVpc) SetDocumentdbSubnetsZoneBIds(v []string) {
	o.DocumentdbSubnetsZoneBIds = v
}

// GetDocumentdbSubnetsZoneCIds returns the DocumentdbSubnetsZoneCIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeatureAwsExistingVpc) GetDocumentdbSubnetsZoneCIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DocumentdbSubnetsZoneCIds
}

// GetDocumentdbSubnetsZoneCIdsOk returns a tuple with the DocumentdbSubnetsZoneCIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeatureAwsExistingVpc) GetDocumentdbSubnetsZoneCIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.DocumentdbSubnetsZoneCIds) {
		return nil, false
	}
	return o.DocumentdbSubnetsZoneCIds, true
}

// HasDocumentdbSubnetsZoneCIds returns a boolean if a field has been set.
func (o *ClusterFeatureAwsExistingVpc) HasDocumentdbSubnetsZoneCIds() bool {
	if o != nil && !IsNil(o.DocumentdbSubnetsZoneCIds) {
		return true
	}

	return false
}

// SetDocumentdbSubnetsZoneCIds gets a reference to the given []string and assigns it to the DocumentdbSubnetsZoneCIds field.
func (o *ClusterFeatureAwsExistingVpc) SetDocumentdbSubnetsZoneCIds(v []string) {
	o.DocumentdbSubnetsZoneCIds = v
}

// GetElasticacheSubnetsZoneAIds returns the ElasticacheSubnetsZoneAIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeatureAwsExistingVpc) GetElasticacheSubnetsZoneAIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ElasticacheSubnetsZoneAIds
}

// GetElasticacheSubnetsZoneAIdsOk returns a tuple with the ElasticacheSubnetsZoneAIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeatureAwsExistingVpc) GetElasticacheSubnetsZoneAIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ElasticacheSubnetsZoneAIds) {
		return nil, false
	}
	return o.ElasticacheSubnetsZoneAIds, true
}

// HasElasticacheSubnetsZoneAIds returns a boolean if a field has been set.
func (o *ClusterFeatureAwsExistingVpc) HasElasticacheSubnetsZoneAIds() bool {
	if o != nil && !IsNil(o.ElasticacheSubnetsZoneAIds) {
		return true
	}

	return false
}

// SetElasticacheSubnetsZoneAIds gets a reference to the given []string and assigns it to the ElasticacheSubnetsZoneAIds field.
func (o *ClusterFeatureAwsExistingVpc) SetElasticacheSubnetsZoneAIds(v []string) {
	o.ElasticacheSubnetsZoneAIds = v
}

// GetElasticacheSubnetsZoneBIds returns the ElasticacheSubnetsZoneBIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeatureAwsExistingVpc) GetElasticacheSubnetsZoneBIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ElasticacheSubnetsZoneBIds
}

// GetElasticacheSubnetsZoneBIdsOk returns a tuple with the ElasticacheSubnetsZoneBIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeatureAwsExistingVpc) GetElasticacheSubnetsZoneBIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ElasticacheSubnetsZoneBIds) {
		return nil, false
	}
	return o.ElasticacheSubnetsZoneBIds, true
}

// HasElasticacheSubnetsZoneBIds returns a boolean if a field has been set.
func (o *ClusterFeatureAwsExistingVpc) HasElasticacheSubnetsZoneBIds() bool {
	if o != nil && !IsNil(o.ElasticacheSubnetsZoneBIds) {
		return true
	}

	return false
}

// SetElasticacheSubnetsZoneBIds gets a reference to the given []string and assigns it to the ElasticacheSubnetsZoneBIds field.
func (o *ClusterFeatureAwsExistingVpc) SetElasticacheSubnetsZoneBIds(v []string) {
	o.ElasticacheSubnetsZoneBIds = v
}

// GetElasticacheSubnetsZoneCIds returns the ElasticacheSubnetsZoneCIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeatureAwsExistingVpc) GetElasticacheSubnetsZoneCIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ElasticacheSubnetsZoneCIds
}

// GetElasticacheSubnetsZoneCIdsOk returns a tuple with the ElasticacheSubnetsZoneCIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeatureAwsExistingVpc) GetElasticacheSubnetsZoneCIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ElasticacheSubnetsZoneCIds) {
		return nil, false
	}
	return o.ElasticacheSubnetsZoneCIds, true
}

// HasElasticacheSubnetsZoneCIds returns a boolean if a field has been set.
func (o *ClusterFeatureAwsExistingVpc) HasElasticacheSubnetsZoneCIds() bool {
	if o != nil && !IsNil(o.ElasticacheSubnetsZoneCIds) {
		return true
	}

	return false
}

// SetElasticacheSubnetsZoneCIds gets a reference to the given []string and assigns it to the ElasticacheSubnetsZoneCIds field.
func (o *ClusterFeatureAwsExistingVpc) SetElasticacheSubnetsZoneCIds(v []string) {
	o.ElasticacheSubnetsZoneCIds = v
}

// GetRdsSubnetsZoneAIds returns the RdsSubnetsZoneAIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeatureAwsExistingVpc) GetRdsSubnetsZoneAIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RdsSubnetsZoneAIds
}

// GetRdsSubnetsZoneAIdsOk returns a tuple with the RdsSubnetsZoneAIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeatureAwsExistingVpc) GetRdsSubnetsZoneAIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RdsSubnetsZoneAIds) {
		return nil, false
	}
	return o.RdsSubnetsZoneAIds, true
}

// HasRdsSubnetsZoneAIds returns a boolean if a field has been set.
func (o *ClusterFeatureAwsExistingVpc) HasRdsSubnetsZoneAIds() bool {
	if o != nil && !IsNil(o.RdsSubnetsZoneAIds) {
		return true
	}

	return false
}

// SetRdsSubnetsZoneAIds gets a reference to the given []string and assigns it to the RdsSubnetsZoneAIds field.
func (o *ClusterFeatureAwsExistingVpc) SetRdsSubnetsZoneAIds(v []string) {
	o.RdsSubnetsZoneAIds = v
}

// GetRdsSubnetsZoneBIds returns the RdsSubnetsZoneBIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeatureAwsExistingVpc) GetRdsSubnetsZoneBIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RdsSubnetsZoneBIds
}

// GetRdsSubnetsZoneBIdsOk returns a tuple with the RdsSubnetsZoneBIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeatureAwsExistingVpc) GetRdsSubnetsZoneBIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RdsSubnetsZoneBIds) {
		return nil, false
	}
	return o.RdsSubnetsZoneBIds, true
}

// HasRdsSubnetsZoneBIds returns a boolean if a field has been set.
func (o *ClusterFeatureAwsExistingVpc) HasRdsSubnetsZoneBIds() bool {
	if o != nil && !IsNil(o.RdsSubnetsZoneBIds) {
		return true
	}

	return false
}

// SetRdsSubnetsZoneBIds gets a reference to the given []string and assigns it to the RdsSubnetsZoneBIds field.
func (o *ClusterFeatureAwsExistingVpc) SetRdsSubnetsZoneBIds(v []string) {
	o.RdsSubnetsZoneBIds = v
}

// GetRdsSubnetsZoneCIds returns the RdsSubnetsZoneCIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeatureAwsExistingVpc) GetRdsSubnetsZoneCIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RdsSubnetsZoneCIds
}

// GetRdsSubnetsZoneCIdsOk returns a tuple with the RdsSubnetsZoneCIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeatureAwsExistingVpc) GetRdsSubnetsZoneCIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RdsSubnetsZoneCIds) {
		return nil, false
	}
	return o.RdsSubnetsZoneCIds, true
}

// HasRdsSubnetsZoneCIds returns a boolean if a field has been set.
func (o *ClusterFeatureAwsExistingVpc) HasRdsSubnetsZoneCIds() bool {
	if o != nil && !IsNil(o.RdsSubnetsZoneCIds) {
		return true
	}

	return false
}

// SetRdsSubnetsZoneCIds gets a reference to the given []string and assigns it to the RdsSubnetsZoneCIds field.
func (o *ClusterFeatureAwsExistingVpc) SetRdsSubnetsZoneCIds(v []string) {
	o.RdsSubnetsZoneCIds = v
}

// GetEksKarpenterFargateSubnetsZoneAIds returns the EksKarpenterFargateSubnetsZoneAIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeatureAwsExistingVpc) GetEksKarpenterFargateSubnetsZoneAIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EksKarpenterFargateSubnetsZoneAIds
}

// GetEksKarpenterFargateSubnetsZoneAIdsOk returns a tuple with the EksKarpenterFargateSubnetsZoneAIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeatureAwsExistingVpc) GetEksKarpenterFargateSubnetsZoneAIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.EksKarpenterFargateSubnetsZoneAIds) {
		return nil, false
	}
	return o.EksKarpenterFargateSubnetsZoneAIds, true
}

// HasEksKarpenterFargateSubnetsZoneAIds returns a boolean if a field has been set.
func (o *ClusterFeatureAwsExistingVpc) HasEksKarpenterFargateSubnetsZoneAIds() bool {
	if o != nil && !IsNil(o.EksKarpenterFargateSubnetsZoneAIds) {
		return true
	}

	return false
}

// SetEksKarpenterFargateSubnetsZoneAIds gets a reference to the given []string and assigns it to the EksKarpenterFargateSubnetsZoneAIds field.
func (o *ClusterFeatureAwsExistingVpc) SetEksKarpenterFargateSubnetsZoneAIds(v []string) {
	o.EksKarpenterFargateSubnetsZoneAIds = v
}

// GetEksKarpenterFargateSubnetsZoneBIds returns the EksKarpenterFargateSubnetsZoneBIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeatureAwsExistingVpc) GetEksKarpenterFargateSubnetsZoneBIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EksKarpenterFargateSubnetsZoneBIds
}

// GetEksKarpenterFargateSubnetsZoneBIdsOk returns a tuple with the EksKarpenterFargateSubnetsZoneBIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeatureAwsExistingVpc) GetEksKarpenterFargateSubnetsZoneBIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.EksKarpenterFargateSubnetsZoneBIds) {
		return nil, false
	}
	return o.EksKarpenterFargateSubnetsZoneBIds, true
}

// HasEksKarpenterFargateSubnetsZoneBIds returns a boolean if a field has been set.
func (o *ClusterFeatureAwsExistingVpc) HasEksKarpenterFargateSubnetsZoneBIds() bool {
	if o != nil && !IsNil(o.EksKarpenterFargateSubnetsZoneBIds) {
		return true
	}

	return false
}

// SetEksKarpenterFargateSubnetsZoneBIds gets a reference to the given []string and assigns it to the EksKarpenterFargateSubnetsZoneBIds field.
func (o *ClusterFeatureAwsExistingVpc) SetEksKarpenterFargateSubnetsZoneBIds(v []string) {
	o.EksKarpenterFargateSubnetsZoneBIds = v
}

// GetEksKarpenterFargateSubnetsZoneCIds returns the EksKarpenterFargateSubnetsZoneCIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClusterFeatureAwsExistingVpc) GetEksKarpenterFargateSubnetsZoneCIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EksKarpenterFargateSubnetsZoneCIds
}

// GetEksKarpenterFargateSubnetsZoneCIdsOk returns a tuple with the EksKarpenterFargateSubnetsZoneCIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterFeatureAwsExistingVpc) GetEksKarpenterFargateSubnetsZoneCIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.EksKarpenterFargateSubnetsZoneCIds) {
		return nil, false
	}
	return o.EksKarpenterFargateSubnetsZoneCIds, true
}

// HasEksKarpenterFargateSubnetsZoneCIds returns a boolean if a field has been set.
func (o *ClusterFeatureAwsExistingVpc) HasEksKarpenterFargateSubnetsZoneCIds() bool {
	if o != nil && !IsNil(o.EksKarpenterFargateSubnetsZoneCIds) {
		return true
	}

	return false
}

// SetEksKarpenterFargateSubnetsZoneCIds gets a reference to the given []string and assigns it to the EksKarpenterFargateSubnetsZoneCIds field.
func (o *ClusterFeatureAwsExistingVpc) SetEksKarpenterFargateSubnetsZoneCIds(v []string) {
	o.EksKarpenterFargateSubnetsZoneCIds = v
}

func (o ClusterFeatureAwsExistingVpc) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterFeatureAwsExistingVpc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aws_vpc_eks_id"] = o.AwsVpcEksId
	toSerialize["eks_subnets_zone_a_ids"] = o.EksSubnetsZoneAIds
	toSerialize["eks_subnets_zone_b_ids"] = o.EksSubnetsZoneBIds
	toSerialize["eks_subnets_zone_c_ids"] = o.EksSubnetsZoneCIds
	if o.DocumentdbSubnetsZoneAIds != nil {
		toSerialize["documentdb_subnets_zone_a_ids"] = o.DocumentdbSubnetsZoneAIds
	}
	if o.DocumentdbSubnetsZoneBIds != nil {
		toSerialize["documentdb_subnets_zone_b_ids"] = o.DocumentdbSubnetsZoneBIds
	}
	if o.DocumentdbSubnetsZoneCIds != nil {
		toSerialize["documentdb_subnets_zone_c_ids"] = o.DocumentdbSubnetsZoneCIds
	}
	if o.ElasticacheSubnetsZoneAIds != nil {
		toSerialize["elasticache_subnets_zone_a_ids"] = o.ElasticacheSubnetsZoneAIds
	}
	if o.ElasticacheSubnetsZoneBIds != nil {
		toSerialize["elasticache_subnets_zone_b_ids"] = o.ElasticacheSubnetsZoneBIds
	}
	if o.ElasticacheSubnetsZoneCIds != nil {
		toSerialize["elasticache_subnets_zone_c_ids"] = o.ElasticacheSubnetsZoneCIds
	}
	if o.RdsSubnetsZoneAIds != nil {
		toSerialize["rds_subnets_zone_a_ids"] = o.RdsSubnetsZoneAIds
	}
	if o.RdsSubnetsZoneBIds != nil {
		toSerialize["rds_subnets_zone_b_ids"] = o.RdsSubnetsZoneBIds
	}
	if o.RdsSubnetsZoneCIds != nil {
		toSerialize["rds_subnets_zone_c_ids"] = o.RdsSubnetsZoneCIds
	}
	if o.EksKarpenterFargateSubnetsZoneAIds != nil {
		toSerialize["eks_karpenter_fargate_subnets_zone_a_ids"] = o.EksKarpenterFargateSubnetsZoneAIds
	}
	if o.EksKarpenterFargateSubnetsZoneBIds != nil {
		toSerialize["eks_karpenter_fargate_subnets_zone_b_ids"] = o.EksKarpenterFargateSubnetsZoneBIds
	}
	if o.EksKarpenterFargateSubnetsZoneCIds != nil {
		toSerialize["eks_karpenter_fargate_subnets_zone_c_ids"] = o.EksKarpenterFargateSubnetsZoneCIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ClusterFeatureAwsExistingVpc) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"aws_vpc_eks_id",
		"eks_subnets_zone_a_ids",
		"eks_subnets_zone_b_ids",
		"eks_subnets_zone_c_ids",
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

	varClusterFeatureAwsExistingVpc := _ClusterFeatureAwsExistingVpc{}

	err = json.Unmarshal(data, &varClusterFeatureAwsExistingVpc)

	if err != nil {
		return err
	}

	*o = ClusterFeatureAwsExistingVpc(varClusterFeatureAwsExistingVpc)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "aws_vpc_eks_id")
		delete(additionalProperties, "eks_subnets_zone_a_ids")
		delete(additionalProperties, "eks_subnets_zone_b_ids")
		delete(additionalProperties, "eks_subnets_zone_c_ids")
		delete(additionalProperties, "documentdb_subnets_zone_a_ids")
		delete(additionalProperties, "documentdb_subnets_zone_b_ids")
		delete(additionalProperties, "documentdb_subnets_zone_c_ids")
		delete(additionalProperties, "elasticache_subnets_zone_a_ids")
		delete(additionalProperties, "elasticache_subnets_zone_b_ids")
		delete(additionalProperties, "elasticache_subnets_zone_c_ids")
		delete(additionalProperties, "rds_subnets_zone_a_ids")
		delete(additionalProperties, "rds_subnets_zone_b_ids")
		delete(additionalProperties, "rds_subnets_zone_c_ids")
		delete(additionalProperties, "eks_karpenter_fargate_subnets_zone_a_ids")
		delete(additionalProperties, "eks_karpenter_fargate_subnets_zone_b_ids")
		delete(additionalProperties, "eks_karpenter_fargate_subnets_zone_c_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableClusterFeatureAwsExistingVpc struct {
	value *ClusterFeatureAwsExistingVpc
	isSet bool
}

func (v NullableClusterFeatureAwsExistingVpc) Get() *ClusterFeatureAwsExistingVpc {
	return v.value
}

func (v *NullableClusterFeatureAwsExistingVpc) Set(val *ClusterFeatureAwsExistingVpc) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterFeatureAwsExistingVpc) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterFeatureAwsExistingVpc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterFeatureAwsExistingVpc(val *ClusterFeatureAwsExistingVpc) *NullableClusterFeatureAwsExistingVpc {
	return &NullableClusterFeatureAwsExistingVpc{value: val, isSet: true}
}

func (v NullableClusterFeatureAwsExistingVpc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterFeatureAwsExistingVpc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
