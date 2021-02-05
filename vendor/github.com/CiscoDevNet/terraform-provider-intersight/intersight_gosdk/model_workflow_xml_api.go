/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-02-05T15:05:56Z.
 *
 * API version: 1.0.9-3562
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// WorkflowXmlApi This models a single XML API request that can be sent to any Cisco UCS devices that support Cisco UCS XML API interface.
type WorkflowXmlApi struct {
	WorkflowApi
	AdditionalProperties map[string]interface{}
}

type _WorkflowXmlApi WorkflowXmlApi

// NewWorkflowXmlApi instantiates a new WorkflowXmlApi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowXmlApi() *WorkflowXmlApi {
	this := WorkflowXmlApi{}
	return &this
}

// NewWorkflowXmlApiWithDefaults instantiates a new WorkflowXmlApi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowXmlApiWithDefaults() *WorkflowXmlApi {
	this := WorkflowXmlApi{}
	return &this
}

func (o WorkflowXmlApi) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowApi, errWorkflowApi := json.Marshal(o.WorkflowApi)
	if errWorkflowApi != nil {
		return []byte{}, errWorkflowApi
	}
	errWorkflowApi = json.Unmarshal([]byte(serializedWorkflowApi), &toSerialize)
	if errWorkflowApi != nil {
		return []byte{}, errWorkflowApi
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowXmlApi) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowXmlApiWithoutEmbeddedStruct struct {
	}

	varWorkflowXmlApiWithoutEmbeddedStruct := WorkflowXmlApiWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowXmlApiWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowXmlApi := _WorkflowXmlApi{}
		*o = WorkflowXmlApi(varWorkflowXmlApi)
	} else {
		return err
	}

	varWorkflowXmlApi := _WorkflowXmlApi{}

	err = json.Unmarshal(bytes, &varWorkflowXmlApi)
	if err == nil {
		o.WorkflowApi = varWorkflowXmlApi.WorkflowApi
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectWorkflowApi := reflect.ValueOf(o.WorkflowApi)
		for i := 0; i < reflectWorkflowApi.Type().NumField(); i++ {
			t := reflectWorkflowApi.Type().Field(i)

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

type NullableWorkflowXmlApi struct {
	value *WorkflowXmlApi
	isSet bool
}

func (v NullableWorkflowXmlApi) Get() *WorkflowXmlApi {
	return v.value
}

func (v *NullableWorkflowXmlApi) Set(val *WorkflowXmlApi) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowXmlApi) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowXmlApi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowXmlApi(val *WorkflowXmlApi) *NullableWorkflowXmlApi {
	return &NullableWorkflowXmlApi{value: val, isSet: true}
}

func (v NullableWorkflowXmlApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowXmlApi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
