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

// OrganizationEventType Type of the organization event
type OrganizationEventType string

// List of OrganizationEventType
const (
	ORGANIZATIONEVENTTYPE_CREATE                   OrganizationEventType = "CREATE"
	ORGANIZATIONEVENTTYPE_UPDATE                   OrganizationEventType = "UPDATE"
	ORGANIZATIONEVENTTYPE_DELETE                   OrganizationEventType = "DELETE"
	ORGANIZATIONEVENTTYPE_ACCEPT                   OrganizationEventType = "ACCEPT"
	ORGANIZATIONEVENTTYPE_EXPORT                   OrganizationEventType = "EXPORT"
	ORGANIZATIONEVENTTYPE_CLONE                    OrganizationEventType = "CLONE"
	ORGANIZATIONEVENTTYPE_DEPLOY_QUEUED            OrganizationEventType = "DEPLOY_QUEUED"
	ORGANIZATIONEVENTTYPE_STOP_QUEUED              OrganizationEventType = "STOP_QUEUED"
	ORGANIZATIONEVENTTYPE_RESTART_QUEUED           OrganizationEventType = "RESTART_QUEUED"
	ORGANIZATIONEVENTTYPE_DELETE_QUEUED            OrganizationEventType = "DELETE_QUEUED"
	ORGANIZATIONEVENTTYPE_UNINSTALL_QUEUED         OrganizationEventType = "UNINSTALL_QUEUED"
	ORGANIZATIONEVENTTYPE_MAINTENANCE              OrganizationEventType = "MAINTENANCE"
	ORGANIZATIONEVENTTYPE_DRY_RUN                  OrganizationEventType = "DRY_RUN"
	ORGANIZATIONEVENTTYPE_TRIGGER_REDEPLOY         OrganizationEventType = "TRIGGER_REDEPLOY"
	ORGANIZATIONEVENTTYPE_TRIGGER_CANCEL           OrganizationEventType = "TRIGGER_CANCEL"
	ORGANIZATIONEVENTTYPE_FORCE_RUN_QUEUED         OrganizationEventType = "FORCE_RUN_QUEUED"
	ORGANIZATIONEVENTTYPE_FORCE_RUN_QUEUED_DEPLOY  OrganizationEventType = "FORCE_RUN_QUEUED_DEPLOY"
	ORGANIZATIONEVENTTYPE_FORCE_RUN_QUEUED_STOP    OrganizationEventType = "FORCE_RUN_QUEUED_STOP"
	ORGANIZATIONEVENTTYPE_FORCE_RUN_QUEUED_DELETE  OrganizationEventType = "FORCE_RUN_QUEUED_DELETE"
	ORGANIZATIONEVENTTYPE_TRIGGER_DEPLOY           OrganizationEventType = "TRIGGER_DEPLOY"
	ORGANIZATIONEVENTTYPE_TRIGGER_STOP             OrganizationEventType = "TRIGGER_STOP"
	ORGANIZATIONEVENTTYPE_TRIGGER_RESTART          OrganizationEventType = "TRIGGER_RESTART"
	ORGANIZATIONEVENTTYPE_TRIGGER_DELETE           OrganizationEventType = "TRIGGER_DELETE"
	ORGANIZATIONEVENTTYPE_TRIGGER_UNINSTALL        OrganizationEventType = "TRIGGER_UNINSTALL"
	ORGANIZATIONEVENTTYPE_TRIGGER_DEPLOY_DRY_RUN   OrganizationEventType = "TRIGGER_DEPLOY_DRY_RUN"
	ORGANIZATIONEVENTTYPE_TRIGGER_FORCE_RUN        OrganizationEventType = "TRIGGER_FORCE_RUN"
	ORGANIZATIONEVENTTYPE_TRIGGER_FORCE_RUN_DEPLOY OrganizationEventType = "TRIGGER_FORCE_RUN_DEPLOY"
	ORGANIZATIONEVENTTYPE_TRIGGER_FORCE_RUN_STOP   OrganizationEventType = "TRIGGER_FORCE_RUN_STOP"
	ORGANIZATIONEVENTTYPE_TRIGGER_FORCE_RUN_DELETE OrganizationEventType = "TRIGGER_FORCE_RUN_DELETE"
	ORGANIZATIONEVENTTYPE_DEPLOYED                 OrganizationEventType = "DEPLOYED"
	ORGANIZATIONEVENTTYPE_STOPPED                  OrganizationEventType = "STOPPED"
	ORGANIZATIONEVENTTYPE_DELETED                  OrganizationEventType = "DELETED"
	ORGANIZATIONEVENTTYPE_UNINSTALLED              OrganizationEventType = "UNINSTALLED"
	ORGANIZATIONEVENTTYPE_RESTARTED                OrganizationEventType = "RESTARTED"
	ORGANIZATIONEVENTTYPE_DEPLOYED_DRY_RUN         OrganizationEventType = "DEPLOYED_DRY_RUN"
	ORGANIZATIONEVENTTYPE_DEPLOY_FAILED            OrganizationEventType = "DEPLOY_FAILED"
	ORGANIZATIONEVENTTYPE_STOP_FAILED              OrganizationEventType = "STOP_FAILED"
	ORGANIZATIONEVENTTYPE_DELETE_FAILED            OrganizationEventType = "DELETE_FAILED"
	ORGANIZATIONEVENTTYPE_UNINSTALL_FAILED         OrganizationEventType = "UNINSTALL_FAILED"
	ORGANIZATIONEVENTTYPE_RESTART_FAILED           OrganizationEventType = "RESTART_FAILED"
	ORGANIZATIONEVENTTYPE_DEPLOYED_DRY_RUN_FAILED  OrganizationEventType = "DEPLOYED_DRY_RUN_FAILED"
	ORGANIZATIONEVENTTYPE_SHELL                    OrganizationEventType = "SHELL"
	ORGANIZATIONEVENTTYPE_PORT_FORWARD             OrganizationEventType = "PORT_FORWARD"
	ORGANIZATIONEVENTTYPE_REMOTE_DEBUG             OrganizationEventType = "REMOTE_DEBUG"
	ORGANIZATIONEVENTTYPE_IMPORT                   OrganizationEventType = "IMPORT"
	ORGANIZATIONEVENTTYPE_LOCK                     OrganizationEventType = "LOCK"
	ORGANIZATIONEVENTTYPE_UNLOCK                   OrganizationEventType = "UNLOCK"
)

// All allowed values of OrganizationEventType enum
var AllowedOrganizationEventTypeEnumValues = []OrganizationEventType{
	"CREATE",
	"UPDATE",
	"DELETE",
	"ACCEPT",
	"EXPORT",
	"CLONE",
	"DEPLOY_QUEUED",
	"STOP_QUEUED",
	"RESTART_QUEUED",
	"DELETE_QUEUED",
	"UNINSTALL_QUEUED",
	"MAINTENANCE",
	"DRY_RUN",
	"TRIGGER_REDEPLOY",
	"TRIGGER_CANCEL",
	"FORCE_RUN_QUEUED",
	"FORCE_RUN_QUEUED_DEPLOY",
	"FORCE_RUN_QUEUED_STOP",
	"FORCE_RUN_QUEUED_DELETE",
	"TRIGGER_DEPLOY",
	"TRIGGER_STOP",
	"TRIGGER_RESTART",
	"TRIGGER_DELETE",
	"TRIGGER_UNINSTALL",
	"TRIGGER_DEPLOY_DRY_RUN",
	"TRIGGER_FORCE_RUN",
	"TRIGGER_FORCE_RUN_DEPLOY",
	"TRIGGER_FORCE_RUN_STOP",
	"TRIGGER_FORCE_RUN_DELETE",
	"DEPLOYED",
	"STOPPED",
	"DELETED",
	"UNINSTALLED",
	"RESTARTED",
	"DEPLOYED_DRY_RUN",
	"DEPLOY_FAILED",
	"STOP_FAILED",
	"DELETE_FAILED",
	"UNINSTALL_FAILED",
	"RESTART_FAILED",
	"DEPLOYED_DRY_RUN_FAILED",
	"SHELL",
	"PORT_FORWARD",
	"REMOTE_DEBUG",
	"IMPORT",
	"LOCK",
	"UNLOCK",
}

func (v *OrganizationEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrganizationEventType(value)
	for _, existing := range AllowedOrganizationEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrganizationEventType", value)
}

// NewOrganizationEventTypeFromValue returns a pointer to a valid OrganizationEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrganizationEventTypeFromValue(v string) (*OrganizationEventType, error) {
	ev := OrganizationEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrganizationEventType: valid values are %v", v, AllowedOrganizationEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrganizationEventType) IsValid() bool {
	for _, existing := range AllowedOrganizationEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrganizationEventType value
func (v OrganizationEventType) Ptr() *OrganizationEventType {
	return &v
}

type NullableOrganizationEventType struct {
	value *OrganizationEventType
	isSet bool
}

func (v NullableOrganizationEventType) Get() *OrganizationEventType {
	return v.value
}

func (v *NullableOrganizationEventType) Set(val *OrganizationEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationEventType(val *OrganizationEventType) *NullableOrganizationEventType {
	return &NullableOrganizationEventType{value: val, isSet: true}
}

func (v NullableOrganizationEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
