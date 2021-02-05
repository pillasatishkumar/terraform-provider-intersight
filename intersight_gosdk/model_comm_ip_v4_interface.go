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

// CommIpV4Interface The configuration data of a single IPv4 interface, including IP address, netmask, and default gateway.
type CommIpV4Interface struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The IPv4 address of the default gateway.
	Gateway *string `json:"Gateway,omitempty"`
	// The IPv4 Address, represented in the standard dot-decimal notation, e.g. 192.168.1.3.
	IpAddress *string `json:"IpAddress,omitempty"`
	// The IPv4 Netmask, represented in the standard dot-decimal notation, e.g. 255.255.255.0.
	Netmask              *string `json:"Netmask,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommIpV4Interface CommIpV4Interface

// NewCommIpV4Interface instantiates a new CommIpV4Interface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommIpV4Interface(classId string, objectType string) *CommIpV4Interface {
	this := CommIpV4Interface{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewCommIpV4InterfaceWithDefaults instantiates a new CommIpV4Interface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommIpV4InterfaceWithDefaults() *CommIpV4Interface {
	this := CommIpV4Interface{}
	var classId string = "comm.IpV4Interface"
	this.ClassId = classId
	var objectType string = "comm.IpV4Interface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *CommIpV4Interface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *CommIpV4Interface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *CommIpV4Interface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *CommIpV4Interface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *CommIpV4Interface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *CommIpV4Interface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *CommIpV4Interface) GetGateway() string {
	if o == nil || o.Gateway == nil {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommIpV4Interface) GetGatewayOk() (*string, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *CommIpV4Interface) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *CommIpV4Interface) SetGateway(v string) {
	o.Gateway = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *CommIpV4Interface) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommIpV4Interface) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *CommIpV4Interface) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *CommIpV4Interface) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *CommIpV4Interface) GetNetmask() string {
	if o == nil || o.Netmask == nil {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommIpV4Interface) GetNetmaskOk() (*string, bool) {
	if o == nil || o.Netmask == nil {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *CommIpV4Interface) HasNetmask() bool {
	if o != nil && o.Netmask != nil {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *CommIpV4Interface) SetNetmask(v string) {
	o.Netmask = &v
}

func (o CommIpV4Interface) MarshalJSON() ([]byte, error) {
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
	if o.Gateway != nil {
		toSerialize["Gateway"] = o.Gateway
	}
	if o.IpAddress != nil {
		toSerialize["IpAddress"] = o.IpAddress
	}
	if o.Netmask != nil {
		toSerialize["Netmask"] = o.Netmask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CommIpV4Interface) UnmarshalJSON(bytes []byte) (err error) {
	type CommIpV4InterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The IPv4 address of the default gateway.
		Gateway *string `json:"Gateway,omitempty"`
		// The IPv4 Address, represented in the standard dot-decimal notation, e.g. 192.168.1.3.
		IpAddress *string `json:"IpAddress,omitempty"`
		// The IPv4 Netmask, represented in the standard dot-decimal notation, e.g. 255.255.255.0.
		Netmask *string `json:"Netmask,omitempty"`
	}

	varCommIpV4InterfaceWithoutEmbeddedStruct := CommIpV4InterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCommIpV4InterfaceWithoutEmbeddedStruct)
	if err == nil {
		varCommIpV4Interface := _CommIpV4Interface{}
		varCommIpV4Interface.ClassId = varCommIpV4InterfaceWithoutEmbeddedStruct.ClassId
		varCommIpV4Interface.ObjectType = varCommIpV4InterfaceWithoutEmbeddedStruct.ObjectType
		varCommIpV4Interface.Gateway = varCommIpV4InterfaceWithoutEmbeddedStruct.Gateway
		varCommIpV4Interface.IpAddress = varCommIpV4InterfaceWithoutEmbeddedStruct.IpAddress
		varCommIpV4Interface.Netmask = varCommIpV4InterfaceWithoutEmbeddedStruct.Netmask
		*o = CommIpV4Interface(varCommIpV4Interface)
	} else {
		return err
	}

	varCommIpV4Interface := _CommIpV4Interface{}

	err = json.Unmarshal(bytes, &varCommIpV4Interface)
	if err == nil {
		o.MoBaseComplexType = varCommIpV4Interface.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Gateway")
		delete(additionalProperties, "IpAddress")
		delete(additionalProperties, "Netmask")

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

type NullableCommIpV4Interface struct {
	value *CommIpV4Interface
	isSet bool
}

func (v NullableCommIpV4Interface) Get() *CommIpV4Interface {
	return v.value
}

func (v *NullableCommIpV4Interface) Set(val *CommIpV4Interface) {
	v.value = val
	v.isSet = true
}

func (v NullableCommIpV4Interface) IsSet() bool {
	return v.isSet
}

func (v *NullableCommIpV4Interface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommIpV4Interface(val *CommIpV4Interface) *NullableCommIpV4Interface {
	return &NullableCommIpV4Interface{value: val, isSet: true}
}

func (v NullableCommIpV4Interface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommIpV4Interface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
