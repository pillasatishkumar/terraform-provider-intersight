/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-01-11T18:30:19Z.
 *
 * API version: 1.0.9-3252
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
)

// BootDeviceBaseAllOf Definition of the list of properties defined in 'boot.DeviceBase', excluding properties defined in parent classes.
type BootDeviceBaseAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Specifies if the boot device is enabled or disabled.
	Enabled *bool `json:"Enabled,omitempty"`
	// A name that helps identify a boot device. It can be any string that adheres to the following constraints. It should start and end with an alphanumeric character. It can have underscores and hyphens. It cannot be more than 30 characters.
	Name                 *string `json:"Name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootDeviceBaseAllOf BootDeviceBaseAllOf

// NewBootDeviceBaseAllOf instantiates a new BootDeviceBaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootDeviceBaseAllOf(classId string, objectType string) *BootDeviceBaseAllOf {
	this := BootDeviceBaseAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewBootDeviceBaseAllOfWithDefaults instantiates a new BootDeviceBaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootDeviceBaseAllOfWithDefaults() *BootDeviceBaseAllOf {
	this := BootDeviceBaseAllOf{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootDeviceBaseAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootDeviceBaseAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootDeviceBaseAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BootDeviceBaseAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootDeviceBaseAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootDeviceBaseAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BootDeviceBaseAllOf) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootDeviceBaseAllOf) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BootDeviceBaseAllOf) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BootDeviceBaseAllOf) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BootDeviceBaseAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootDeviceBaseAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BootDeviceBaseAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BootDeviceBaseAllOf) SetName(v string) {
	o.Name = &v
}

func (o BootDeviceBaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BootDeviceBaseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBootDeviceBaseAllOf := _BootDeviceBaseAllOf{}

	if err = json.Unmarshal(bytes, &varBootDeviceBaseAllOf); err == nil {
		*o = BootDeviceBaseAllOf(varBootDeviceBaseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Enabled")
		delete(additionalProperties, "Name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBootDeviceBaseAllOf struct {
	value *BootDeviceBaseAllOf
	isSet bool
}

func (v NullableBootDeviceBaseAllOf) Get() *BootDeviceBaseAllOf {
	return v.value
}

func (v *NullableBootDeviceBaseAllOf) Set(val *BootDeviceBaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBootDeviceBaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBootDeviceBaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootDeviceBaseAllOf(val *BootDeviceBaseAllOf) *NullableBootDeviceBaseAllOf {
	return &NullableBootDeviceBaseAllOf{value: val, isSet: true}
}

func (v NullableBootDeviceBaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootDeviceBaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
