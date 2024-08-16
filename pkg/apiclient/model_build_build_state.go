/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
	"fmt"
)

// BuildBuildState the model 'BuildBuildState'
type BuildBuildState string

// List of build.BuildState
const (
	BuildStatePending   BuildBuildState = "pending"
	BuildStateRunning   BuildBuildState = "running"
	BuildStateError     BuildBuildState = "error"
	BuildStateSuccess   BuildBuildState = "success"
	BuildStatePublished BuildBuildState = "published"
)

// All allowed values of BuildBuildState enum
var AllowedBuildBuildStateEnumValues = []BuildBuildState{
	"pending",
	"running",
	"error",
	"success",
	"published",
}

func (v *BuildBuildState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BuildBuildState(value)
	for _, existing := range AllowedBuildBuildStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BuildBuildState", value)
}

// NewBuildBuildStateFromValue returns a pointer to a valid BuildBuildState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBuildBuildStateFromValue(v string) (*BuildBuildState, error) {
	ev := BuildBuildState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BuildBuildState: valid values are %v", v, AllowedBuildBuildStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BuildBuildState) IsValid() bool {
	for _, existing := range AllowedBuildBuildStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to build.BuildState value
func (v BuildBuildState) Ptr() *BuildBuildState {
	return &v
}

type NullableBuildBuildState struct {
	value *BuildBuildState
	isSet bool
}

func (v NullableBuildBuildState) Get() *BuildBuildState {
	return v.value
}

func (v *NullableBuildBuildState) Set(val *BuildBuildState) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildBuildState) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildBuildState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildBuildState(val *BuildBuildState) *NullableBuildBuildState {
	return &NullableBuildBuildState{value: val, isSet: true}
}

func (v NullableBuildBuildState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildBuildState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
