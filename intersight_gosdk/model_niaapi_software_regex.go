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

// NiaapiSoftwareRegex The regular expression to parse the software version strings.
type NiaapiSoftwareRegex struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Regular Expression pattern used to reconginze the version string.
	Regex *string `json:"Regex,omitempty"`
	// Software release. A set of Software releases seperated by comma which can be recongized by according Regex pattern.
	SoftwareVersion      *string `json:"SoftwareVersion,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiaapiSoftwareRegex NiaapiSoftwareRegex

// NewNiaapiSoftwareRegex instantiates a new NiaapiSoftwareRegex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiaapiSoftwareRegex(classId string, objectType string) *NiaapiSoftwareRegex {
	this := NiaapiSoftwareRegex{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiaapiSoftwareRegexWithDefaults instantiates a new NiaapiSoftwareRegex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiaapiSoftwareRegexWithDefaults() *NiaapiSoftwareRegex {
	this := NiaapiSoftwareRegex{}
	var classId string = "niaapi.SoftwareRegex"
	this.ClassId = classId
	var objectType string = "niaapi.SoftwareRegex"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiaapiSoftwareRegex) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareRegex) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiaapiSoftwareRegex) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiaapiSoftwareRegex) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareRegex) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiaapiSoftwareRegex) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRegex returns the Regex field value if set, zero value otherwise.
func (o *NiaapiSoftwareRegex) GetRegex() string {
	if o == nil || o.Regex == nil {
		var ret string
		return ret
	}
	return *o.Regex
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareRegex) GetRegexOk() (*string, bool) {
	if o == nil || o.Regex == nil {
		return nil, false
	}
	return o.Regex, true
}

// HasRegex returns a boolean if a field has been set.
func (o *NiaapiSoftwareRegex) HasRegex() bool {
	if o != nil && o.Regex != nil {
		return true
	}

	return false
}

// SetRegex gets a reference to the given string and assigns it to the Regex field.
func (o *NiaapiSoftwareRegex) SetRegex(v string) {
	o.Regex = &v
}

// GetSoftwareVersion returns the SoftwareVersion field value if set, zero value otherwise.
func (o *NiaapiSoftwareRegex) GetSoftwareVersion() string {
	if o == nil || o.SoftwareVersion == nil {
		var ret string
		return ret
	}
	return *o.SoftwareVersion
}

// GetSoftwareVersionOk returns a tuple with the SoftwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiaapiSoftwareRegex) GetSoftwareVersionOk() (*string, bool) {
	if o == nil || o.SoftwareVersion == nil {
		return nil, false
	}
	return o.SoftwareVersion, true
}

// HasSoftwareVersion returns a boolean if a field has been set.
func (o *NiaapiSoftwareRegex) HasSoftwareVersion() bool {
	if o != nil && o.SoftwareVersion != nil {
		return true
	}

	return false
}

// SetSoftwareVersion gets a reference to the given string and assigns it to the SoftwareVersion field.
func (o *NiaapiSoftwareRegex) SetSoftwareVersion(v string) {
	o.SoftwareVersion = &v
}

func (o NiaapiSoftwareRegex) MarshalJSON() ([]byte, error) {
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
	if o.Regex != nil {
		toSerialize["Regex"] = o.Regex
	}
	if o.SoftwareVersion != nil {
		toSerialize["SoftwareVersion"] = o.SoftwareVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiaapiSoftwareRegex) UnmarshalJSON(bytes []byte) (err error) {
	type NiaapiSoftwareRegexWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Regular Expression pattern used to reconginze the version string.
		Regex *string `json:"Regex,omitempty"`
		// Software release. A set of Software releases seperated by comma which can be recongized by according Regex pattern.
		SoftwareVersion *string `json:"SoftwareVersion,omitempty"`
	}

	varNiaapiSoftwareRegexWithoutEmbeddedStruct := NiaapiSoftwareRegexWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varNiaapiSoftwareRegexWithoutEmbeddedStruct)
	if err == nil {
		varNiaapiSoftwareRegex := _NiaapiSoftwareRegex{}
		varNiaapiSoftwareRegex.ClassId = varNiaapiSoftwareRegexWithoutEmbeddedStruct.ClassId
		varNiaapiSoftwareRegex.ObjectType = varNiaapiSoftwareRegexWithoutEmbeddedStruct.ObjectType
		varNiaapiSoftwareRegex.Regex = varNiaapiSoftwareRegexWithoutEmbeddedStruct.Regex
		varNiaapiSoftwareRegex.SoftwareVersion = varNiaapiSoftwareRegexWithoutEmbeddedStruct.SoftwareVersion
		*o = NiaapiSoftwareRegex(varNiaapiSoftwareRegex)
	} else {
		return err
	}

	varNiaapiSoftwareRegex := _NiaapiSoftwareRegex{}

	err = json.Unmarshal(bytes, &varNiaapiSoftwareRegex)
	if err == nil {
		o.MoBaseComplexType = varNiaapiSoftwareRegex.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Regex")
		delete(additionalProperties, "SoftwareVersion")

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

type NullableNiaapiSoftwareRegex struct {
	value *NiaapiSoftwareRegex
	isSet bool
}

func (v NullableNiaapiSoftwareRegex) Get() *NiaapiSoftwareRegex {
	return v.value
}

func (v *NullableNiaapiSoftwareRegex) Set(val *NiaapiSoftwareRegex) {
	v.value = val
	v.isSet = true
}

func (v NullableNiaapiSoftwareRegex) IsSet() bool {
	return v.isSet
}

func (v *NullableNiaapiSoftwareRegex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiaapiSoftwareRegex(val *NiaapiSoftwareRegex) *NullableNiaapiSoftwareRegex {
	return &NullableNiaapiSoftwareRegex{value: val, isSet: true}
}

func (v NullableNiaapiSoftwareRegex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiaapiSoftwareRegex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
