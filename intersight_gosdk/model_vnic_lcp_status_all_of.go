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
)

// VnicLcpStatusAllOf Definition of the list of properties defined in 'vnic.LcpStatus', excluding properties defined in parent classes.
type VnicLcpStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The reason for the status - it will be empty if status is ok or validating. If error, it will have the appropriate message indicating the reason for failure.
	Reason *string `json:"Reason,omitempty"`
	// Indicates if the LCP is ready for Deploy or not. * `ok` - No issues with the LCP/SCP/VIF. * `error` - The LCP/SCP/VIF cannot be deployed due to error. * `validating` - Validation in progress for the LCP.
	Status               *string                                  `json:"Status,omitempty"`
	VnicInfo             []VnicVifStatus                          `json:"VnicInfo,omitempty"`
	Profile              *PolicyAbstractConfigProfileRelationship `json:"Profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicLcpStatusAllOf VnicLcpStatusAllOf

// NewVnicLcpStatusAllOf instantiates a new VnicLcpStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicLcpStatusAllOf(classId string, objectType string) *VnicLcpStatusAllOf {
	this := VnicLcpStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var status string = "ok"
	this.Status = &status
	return &this
}

// NewVnicLcpStatusAllOfWithDefaults instantiates a new VnicLcpStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicLcpStatusAllOfWithDefaults() *VnicLcpStatusAllOf {
	this := VnicLcpStatusAllOf{}
	var classId string = "vnic.LcpStatus"
	this.ClassId = classId
	var objectType string = "vnic.LcpStatus"
	this.ObjectType = objectType
	var status string = "ok"
	this.Status = &status
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicLcpStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicLcpStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicLcpStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicLcpStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicLcpStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicLcpStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *VnicLcpStatusAllOf) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLcpStatusAllOf) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *VnicLcpStatusAllOf) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *VnicLcpStatusAllOf) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VnicLcpStatusAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLcpStatusAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VnicLcpStatusAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VnicLcpStatusAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetVnicInfo returns the VnicInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicLcpStatusAllOf) GetVnicInfo() []VnicVifStatus {
	if o == nil {
		var ret []VnicVifStatus
		return ret
	}
	return o.VnicInfo
}

// GetVnicInfoOk returns a tuple with the VnicInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicLcpStatusAllOf) GetVnicInfoOk() (*[]VnicVifStatus, bool) {
	if o == nil || o.VnicInfo == nil {
		return nil, false
	}
	return &o.VnicInfo, true
}

// HasVnicInfo returns a boolean if a field has been set.
func (o *VnicLcpStatusAllOf) HasVnicInfo() bool {
	if o != nil && o.VnicInfo != nil {
		return true
	}

	return false
}

// SetVnicInfo gets a reference to the given []VnicVifStatus and assigns it to the VnicInfo field.
func (o *VnicLcpStatusAllOf) SetVnicInfo(v []VnicVifStatus) {
	o.VnicInfo = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *VnicLcpStatusAllOf) GetProfile() PolicyAbstractConfigProfileRelationship {
	if o == nil || o.Profile == nil {
		var ret PolicyAbstractConfigProfileRelationship
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLcpStatusAllOf) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *VnicLcpStatusAllOf) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given PolicyAbstractConfigProfileRelationship and assigns it to the Profile field.
func (o *VnicLcpStatusAllOf) SetProfile(v PolicyAbstractConfigProfileRelationship) {
	o.Profile = &v
}

func (o VnicLcpStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Reason != nil {
		toSerialize["Reason"] = o.Reason
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.VnicInfo != nil {
		toSerialize["VnicInfo"] = o.VnicInfo
	}
	if o.Profile != nil {
		toSerialize["Profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicLcpStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVnicLcpStatusAllOf := _VnicLcpStatusAllOf{}

	if err = json.Unmarshal(bytes, &varVnicLcpStatusAllOf); err == nil {
		*o = VnicLcpStatusAllOf(varVnicLcpStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Reason")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "VnicInfo")
		delete(additionalProperties, "Profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVnicLcpStatusAllOf struct {
	value *VnicLcpStatusAllOf
	isSet bool
}

func (v NullableVnicLcpStatusAllOf) Get() *VnicLcpStatusAllOf {
	return v.value
}

func (v *NullableVnicLcpStatusAllOf) Set(val *VnicLcpStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicLcpStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicLcpStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicLcpStatusAllOf(val *VnicLcpStatusAllOf) *NullableVnicLcpStatusAllOf {
	return &NullableVnicLcpStatusAllOf{value: val, isSet: true}
}

func (v NullableVnicLcpStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicLcpStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
