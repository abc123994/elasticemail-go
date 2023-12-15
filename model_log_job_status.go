/*
Elastic Email REST API

This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    The API has a limit of 20 concurrent connections and a hard timeout of 600 seconds per request.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://app.elasticemail.com/marketing/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>

API version: 4.0.0
Contact: support@elasticemail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"encoding/json"
	"fmt"
)

// LogJobStatus the model 'LogJobStatus'
type LogJobStatus string

// List of LogJobStatus
const (
	LOGJOBSTATUS_ALL LogJobStatus = "All"
	LOGJOBSTATUS_READY_TO_SEND LogJobStatus = "ReadyToSend"
	LOGJOBSTATUS_WAITING_TO_RETRY LogJobStatus = "WaitingToRetry"
	LOGJOBSTATUS_SENDING LogJobStatus = "Sending"
	LOGJOBSTATUS_ERROR LogJobStatus = "Error"
	LOGJOBSTATUS_SENT LogJobStatus = "Sent"
	LOGJOBSTATUS_OPENED LogJobStatus = "Opened"
	LOGJOBSTATUS_CLICKED LogJobStatus = "Clicked"
	LOGJOBSTATUS_UNSUBSCRIBED LogJobStatus = "Unsubscribed"
	LOGJOBSTATUS_ABUSE_REPORT LogJobStatus = "AbuseReport"
)

// All allowed values of LogJobStatus enum
var AllowedLogJobStatusEnumValues = []LogJobStatus{
	"All",
	"ReadyToSend",
	"WaitingToRetry",
	"Sending",
	"Error",
	"Sent",
	"Opened",
	"Clicked",
	"Unsubscribed",
	"AbuseReport",
}

func (v *LogJobStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogJobStatus(value)
	for _, existing := range AllowedLogJobStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogJobStatus", value)
}

// NewLogJobStatusFromValue returns a pointer to a valid LogJobStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogJobStatusFromValue(v string) (*LogJobStatus, error) {
	ev := LogJobStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogJobStatus: valid values are %v", v, AllowedLogJobStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogJobStatus) IsValid() bool {
	for _, existing := range AllowedLogJobStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogJobStatus value
func (v LogJobStatus) Ptr() *LogJobStatus {
	return &v
}

type NullableLogJobStatus struct {
	value *LogJobStatus
	isSet bool
}

func (v NullableLogJobStatus) Get() *LogJobStatus {
	return v.value
}

func (v *NullableLogJobStatus) Set(val *LogJobStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableLogJobStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableLogJobStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogJobStatus(val *LogJobStatus) *NullableLogJobStatus {
	return &NullableLogJobStatus{value: val, isSet: true}
}

func (v NullableLogJobStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogJobStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
