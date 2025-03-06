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
	"fmt"
)

// ClusterCredentials - struct for ClusterCredentials
type ClusterCredentials struct {
	AwsRoleClusterCredentials   *AwsRoleClusterCredentials
	AwsStaticClusterCredentials *AwsStaticClusterCredentials
	GenericClusterCredentials   *GenericClusterCredentials
	ScalewayClusterCredentials  *ScalewayClusterCredentials
}

// AwsRoleClusterCredentialsAsClusterCredentials is a convenience function that returns AwsRoleClusterCredentials wrapped in ClusterCredentials
func AwsRoleClusterCredentialsAsClusterCredentials(v *AwsRoleClusterCredentials) ClusterCredentials {
	return ClusterCredentials{
		AwsRoleClusterCredentials: v,
	}
}

// AwsStaticClusterCredentialsAsClusterCredentials is a convenience function that returns AwsStaticClusterCredentials wrapped in ClusterCredentials
func AwsStaticClusterCredentialsAsClusterCredentials(v *AwsStaticClusterCredentials) ClusterCredentials {
	return ClusterCredentials{
		AwsStaticClusterCredentials: v,
	}
}

// GenericClusterCredentialsAsClusterCredentials is a convenience function that returns GenericClusterCredentials wrapped in ClusterCredentials
func GenericClusterCredentialsAsClusterCredentials(v *GenericClusterCredentials) ClusterCredentials {
	return ClusterCredentials{
		GenericClusterCredentials: v,
	}
}

// ScalewayClusterCredentialsAsClusterCredentials is a convenience function that returns ScalewayClusterCredentials wrapped in ClusterCredentials
func ScalewayClusterCredentialsAsClusterCredentials(v *ScalewayClusterCredentials) ClusterCredentials {
	return ClusterCredentials{
		ScalewayClusterCredentials: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ClusterCredentials) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'AWS'
	if jsonDict["object_type"] == "AWS" {
		// try to unmarshal JSON data into AwsStaticClusterCredentials
		err = json.Unmarshal(data, &dst.AwsStaticClusterCredentials)
		if err == nil {
			return nil // data stored in dst.AwsStaticClusterCredentials, return on the first match
		} else {
			dst.AwsStaticClusterCredentials = nil
			return fmt.Errorf("failed to unmarshal ClusterCredentials as AwsStaticClusterCredentials: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AWS_ROLE'
	if jsonDict["object_type"] == "AWS_ROLE" {
		// try to unmarshal JSON data into AwsRoleClusterCredentials
		err = json.Unmarshal(data, &dst.AwsRoleClusterCredentials)
		if err == nil {
			return nil // data stored in dst.AwsRoleClusterCredentials, return on the first match
		} else {
			dst.AwsRoleClusterCredentials = nil
			return fmt.Errorf("failed to unmarshal ClusterCredentials as AwsRoleClusterCredentials: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OTHER'
	if jsonDict["object_type"] == "OTHER" {
		// try to unmarshal JSON data into GenericClusterCredentials
		err = json.Unmarshal(data, &dst.GenericClusterCredentials)
		if err == nil {
			return nil // data stored in dst.GenericClusterCredentials, return on the first match
		} else {
			dst.GenericClusterCredentials = nil
			return fmt.Errorf("failed to unmarshal ClusterCredentials as GenericClusterCredentials: %s", err.Error())
		}
	}

	// check if the discriminator value is 'SCW'
	if jsonDict["object_type"] == "SCW" {
		// try to unmarshal JSON data into ScalewayClusterCredentials
		err = json.Unmarshal(data, &dst.ScalewayClusterCredentials)
		if err == nil {
			return nil // data stored in dst.ScalewayClusterCredentials, return on the first match
		} else {
			dst.ScalewayClusterCredentials = nil
			return fmt.Errorf("failed to unmarshal ClusterCredentials as ScalewayClusterCredentials: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AwsRoleClusterCredentials'
	if jsonDict["object_type"] == "AwsRoleClusterCredentials" {
		// try to unmarshal JSON data into AwsRoleClusterCredentials
		err = json.Unmarshal(data, &dst.AwsRoleClusterCredentials)
		if err == nil {
			return nil // data stored in dst.AwsRoleClusterCredentials, return on the first match
		} else {
			dst.AwsRoleClusterCredentials = nil
			return fmt.Errorf("failed to unmarshal ClusterCredentials as AwsRoleClusterCredentials: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AwsStaticClusterCredentials'
	if jsonDict["object_type"] == "AwsStaticClusterCredentials" {
		// try to unmarshal JSON data into AwsStaticClusterCredentials
		err = json.Unmarshal(data, &dst.AwsStaticClusterCredentials)
		if err == nil {
			return nil // data stored in dst.AwsStaticClusterCredentials, return on the first match
		} else {
			dst.AwsStaticClusterCredentials = nil
			return fmt.Errorf("failed to unmarshal ClusterCredentials as AwsStaticClusterCredentials: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GenericClusterCredentials'
	if jsonDict["object_type"] == "GenericClusterCredentials" {
		// try to unmarshal JSON data into GenericClusterCredentials
		err = json.Unmarshal(data, &dst.GenericClusterCredentials)
		if err == nil {
			return nil // data stored in dst.GenericClusterCredentials, return on the first match
		} else {
			dst.GenericClusterCredentials = nil
			return fmt.Errorf("failed to unmarshal ClusterCredentials as GenericClusterCredentials: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ScalewayClusterCredentials'
	if jsonDict["object_type"] == "ScalewayClusterCredentials" {
		// try to unmarshal JSON data into ScalewayClusterCredentials
		err = json.Unmarshal(data, &dst.ScalewayClusterCredentials)
		if err == nil {
			return nil // data stored in dst.ScalewayClusterCredentials, return on the first match
		} else {
			dst.ScalewayClusterCredentials = nil
			return fmt.Errorf("failed to unmarshal ClusterCredentials as ScalewayClusterCredentials: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ClusterCredentials) MarshalJSON() ([]byte, error) {
	if src.AwsRoleClusterCredentials != nil {
		return json.Marshal(&src.AwsRoleClusterCredentials)
	}

	if src.AwsStaticClusterCredentials != nil {
		return json.Marshal(&src.AwsStaticClusterCredentials)
	}

	if src.GenericClusterCredentials != nil {
		return json.Marshal(&src.GenericClusterCredentials)
	}

	if src.ScalewayClusterCredentials != nil {
		return json.Marshal(&src.ScalewayClusterCredentials)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ClusterCredentials) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AwsRoleClusterCredentials != nil {
		return obj.AwsRoleClusterCredentials
	}

	if obj.AwsStaticClusterCredentials != nil {
		return obj.AwsStaticClusterCredentials
	}

	if obj.GenericClusterCredentials != nil {
		return obj.GenericClusterCredentials
	}

	if obj.ScalewayClusterCredentials != nil {
		return obj.ScalewayClusterCredentials
	}

	// all schemas are nil
	return nil
}

type NullableClusterCredentials struct {
	value *ClusterCredentials
	isSet bool
}

func (v NullableClusterCredentials) Get() *ClusterCredentials {
	return v.value
}

func (v *NullableClusterCredentials) Set(val *ClusterCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterCredentials(val *ClusterCredentials) *NullableClusterCredentials {
	return &NullableClusterCredentials{value: val, isSet: true}
}

func (v NullableClusterCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
