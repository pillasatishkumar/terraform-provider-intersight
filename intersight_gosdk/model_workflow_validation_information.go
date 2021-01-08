/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-01-04T05:15:49Z.
 *
 * API version: 1.0.9-3144
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// WorkflowValidationInformation Type used to capture all the validation information for the workflow definition.
type WorkflowValidationInformation struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The current validation state of this workflow. The possible states are Valid, Invalid, NotValidated (default). * `NotValidated` - The state when workflow definition has not been validated. * `Valid` - The state when workflow definition is valid. * `Invalid` - The state when workflow definition is invalid.
	State                *string                   `json:"State,omitempty"`
	ValidationError      []WorkflowValidationError `json:"ValidationError,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowValidationInformation WorkflowValidationInformation

// NewWorkflowValidationInformation instantiates a new WorkflowValidationInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowValidationInformation(classId string, objectType string) *WorkflowValidationInformation {
	this := WorkflowValidationInformation{}
	this.ClassId = classId
	this.ObjectType = objectType
	var state string = "NotValidated"
	this.State = &state
	return &this
}

// NewWorkflowValidationInformationWithDefaults instantiates a new WorkflowValidationInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowValidationInformationWithDefaults() *WorkflowValidationInformation {
	this := WorkflowValidationInformation{}
	var classId string = "workflow.ValidationInformation"
	this.ClassId = classId
	var objectType string = "workflow.ValidationInformation"
	this.ObjectType = objectType
	var state string = "NotValidated"
	this.State = &state
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowValidationInformation) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowValidationInformation) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowValidationInformation) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowValidationInformation) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowValidationInformation) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowValidationInformation) SetObjectType(v string) {
	o.ObjectType = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *WorkflowValidationInformation) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowValidationInformation) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *WorkflowValidationInformation) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *WorkflowValidationInformation) SetState(v string) {
	o.State = &v
}

// GetValidationError returns the ValidationError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowValidationInformation) GetValidationError() []WorkflowValidationError {
	if o == nil {
		var ret []WorkflowValidationError
		return ret
	}
	return o.ValidationError
}

// GetValidationErrorOk returns a tuple with the ValidationError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowValidationInformation) GetValidationErrorOk() (*[]WorkflowValidationError, bool) {
	if o == nil || o.ValidationError == nil {
		return nil, false
	}
	return &o.ValidationError, true
}

// HasValidationError returns a boolean if a field has been set.
func (o *WorkflowValidationInformation) HasValidationError() bool {
	if o != nil && o.ValidationError != nil {
		return true
	}

	return false
}

// SetValidationError gets a reference to the given []WorkflowValidationError and assigns it to the ValidationError field.
func (o *WorkflowValidationInformation) SetValidationError(v []WorkflowValidationError) {
	o.ValidationError = v
}

func (o WorkflowValidationInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.State != nil {
		toSerialize["State"] = o.State
	}
	if o.ValidationError != nil {
		toSerialize["ValidationError"] = o.ValidationError
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowValidationInformation) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowValidationInformationWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The current validation state of this workflow. The possible states are Valid, Invalid, NotValidated (default). * `NotValidated` - The state when workflow definition has not been validated. * `Valid` - The state when workflow definition is valid. * `Invalid` - The state when workflow definition is invalid.
		State           *string                   `json:"State,omitempty"`
		ValidationError []WorkflowValidationError `json:"ValidationError,omitempty"`
	}

	varWorkflowValidationInformationWithoutEmbeddedStruct := WorkflowValidationInformationWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowValidationInformationWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowValidationInformation := _WorkflowValidationInformation{}
		varWorkflowValidationInformation.ClassId = varWorkflowValidationInformationWithoutEmbeddedStruct.ClassId
		varWorkflowValidationInformation.ObjectType = varWorkflowValidationInformationWithoutEmbeddedStruct.ObjectType
		varWorkflowValidationInformation.State = varWorkflowValidationInformationWithoutEmbeddedStruct.State
		varWorkflowValidationInformation.ValidationError = varWorkflowValidationInformationWithoutEmbeddedStruct.ValidationError
		*o = WorkflowValidationInformation(varWorkflowValidationInformation)
	} else {
		return err
	}

	varWorkflowValidationInformation := _WorkflowValidationInformation{}

	err = json.Unmarshal(bytes, &varWorkflowValidationInformation)
	if err == nil {
		o.MoBaseComplexType = varWorkflowValidationInformation.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "State")
		delete(additionalProperties, "ValidationError")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowValidationInformation struct {
	value *WorkflowValidationInformation
	isSet bool
}

func (v NullableWorkflowValidationInformation) Get() *WorkflowValidationInformation {
	return v.value
}

func (v *NullableWorkflowValidationInformation) Set(val *WorkflowValidationInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowValidationInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowValidationInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowValidationInformation(val *WorkflowValidationInformation) *NullableWorkflowValidationInformation {
	return &NullableWorkflowValidationInformation{value: val, isSet: true}
}

func (v NullableWorkflowValidationInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowValidationInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
