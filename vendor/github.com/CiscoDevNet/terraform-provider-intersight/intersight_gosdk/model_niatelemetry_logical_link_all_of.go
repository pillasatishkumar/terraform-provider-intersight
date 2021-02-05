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
)

// NiatelemetryLogicalLinkAllOf Definition of the list of properties defined in 'niatelemetry.LogicalLink', excluding properties defined in parent classes.
type NiatelemetryLogicalLinkAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Return value of dbId attribute.
	DbId *int64 `json:"DbId,omitempty"`
	// Return value of isPresent attribute.
	IsPresent *bool `json:"IsPresent,omitempty"`
	// Return value of linkAddr1 attribute.
	LinkAddr1 *string `json:"LinkAddr1,omitempty"`
	// Return value of linkAddr2 attribute.
	LinkAddr2 *string `json:"LinkAddr2,omitempty"`
	// Return value of linkState attribute.
	LinkState *string `json:"LinkState,omitempty"`
	// Return value of linkType attribute.
	LinkType *string `json:"LinkType,omitempty"`
	// Return value of uptime attribute.
	Uptime               *string `json:"Uptime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryLogicalLinkAllOf NiatelemetryLogicalLinkAllOf

// NewNiatelemetryLogicalLinkAllOf instantiates a new NiatelemetryLogicalLinkAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryLogicalLinkAllOf(classId string, objectType string) *NiatelemetryLogicalLinkAllOf {
	this := NiatelemetryLogicalLinkAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryLogicalLinkAllOfWithDefaults instantiates a new NiatelemetryLogicalLinkAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryLogicalLinkAllOfWithDefaults() *NiatelemetryLogicalLinkAllOf {
	this := NiatelemetryLogicalLinkAllOf{}
	var classId string = "niatelemetry.LogicalLink"
	this.ClassId = classId
	var objectType string = "niatelemetry.LogicalLink"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryLogicalLinkAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryLogicalLinkAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryLogicalLinkAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryLogicalLinkAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryLogicalLinkAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryLogicalLinkAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDbId returns the DbId field value if set, zero value otherwise.
func (o *NiatelemetryLogicalLinkAllOf) GetDbId() int64 {
	if o == nil || o.DbId == nil {
		var ret int64
		return ret
	}
	return *o.DbId
}

// GetDbIdOk returns a tuple with the DbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLogicalLinkAllOf) GetDbIdOk() (*int64, bool) {
	if o == nil || o.DbId == nil {
		return nil, false
	}
	return o.DbId, true
}

// HasDbId returns a boolean if a field has been set.
func (o *NiatelemetryLogicalLinkAllOf) HasDbId() bool {
	if o != nil && o.DbId != nil {
		return true
	}

	return false
}

// SetDbId gets a reference to the given int64 and assigns it to the DbId field.
func (o *NiatelemetryLogicalLinkAllOf) SetDbId(v int64) {
	o.DbId = &v
}

// GetIsPresent returns the IsPresent field value if set, zero value otherwise.
func (o *NiatelemetryLogicalLinkAllOf) GetIsPresent() bool {
	if o == nil || o.IsPresent == nil {
		var ret bool
		return ret
	}
	return *o.IsPresent
}

// GetIsPresentOk returns a tuple with the IsPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLogicalLinkAllOf) GetIsPresentOk() (*bool, bool) {
	if o == nil || o.IsPresent == nil {
		return nil, false
	}
	return o.IsPresent, true
}

// HasIsPresent returns a boolean if a field has been set.
func (o *NiatelemetryLogicalLinkAllOf) HasIsPresent() bool {
	if o != nil && o.IsPresent != nil {
		return true
	}

	return false
}

// SetIsPresent gets a reference to the given bool and assigns it to the IsPresent field.
func (o *NiatelemetryLogicalLinkAllOf) SetIsPresent(v bool) {
	o.IsPresent = &v
}

// GetLinkAddr1 returns the LinkAddr1 field value if set, zero value otherwise.
func (o *NiatelemetryLogicalLinkAllOf) GetLinkAddr1() string {
	if o == nil || o.LinkAddr1 == nil {
		var ret string
		return ret
	}
	return *o.LinkAddr1
}

// GetLinkAddr1Ok returns a tuple with the LinkAddr1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLogicalLinkAllOf) GetLinkAddr1Ok() (*string, bool) {
	if o == nil || o.LinkAddr1 == nil {
		return nil, false
	}
	return o.LinkAddr1, true
}

// HasLinkAddr1 returns a boolean if a field has been set.
func (o *NiatelemetryLogicalLinkAllOf) HasLinkAddr1() bool {
	if o != nil && o.LinkAddr1 != nil {
		return true
	}

	return false
}

// SetLinkAddr1 gets a reference to the given string and assigns it to the LinkAddr1 field.
func (o *NiatelemetryLogicalLinkAllOf) SetLinkAddr1(v string) {
	o.LinkAddr1 = &v
}

// GetLinkAddr2 returns the LinkAddr2 field value if set, zero value otherwise.
func (o *NiatelemetryLogicalLinkAllOf) GetLinkAddr2() string {
	if o == nil || o.LinkAddr2 == nil {
		var ret string
		return ret
	}
	return *o.LinkAddr2
}

// GetLinkAddr2Ok returns a tuple with the LinkAddr2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLogicalLinkAllOf) GetLinkAddr2Ok() (*string, bool) {
	if o == nil || o.LinkAddr2 == nil {
		return nil, false
	}
	return o.LinkAddr2, true
}

// HasLinkAddr2 returns a boolean if a field has been set.
func (o *NiatelemetryLogicalLinkAllOf) HasLinkAddr2() bool {
	if o != nil && o.LinkAddr2 != nil {
		return true
	}

	return false
}

// SetLinkAddr2 gets a reference to the given string and assigns it to the LinkAddr2 field.
func (o *NiatelemetryLogicalLinkAllOf) SetLinkAddr2(v string) {
	o.LinkAddr2 = &v
}

// GetLinkState returns the LinkState field value if set, zero value otherwise.
func (o *NiatelemetryLogicalLinkAllOf) GetLinkState() string {
	if o == nil || o.LinkState == nil {
		var ret string
		return ret
	}
	return *o.LinkState
}

// GetLinkStateOk returns a tuple with the LinkState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLogicalLinkAllOf) GetLinkStateOk() (*string, bool) {
	if o == nil || o.LinkState == nil {
		return nil, false
	}
	return o.LinkState, true
}

// HasLinkState returns a boolean if a field has been set.
func (o *NiatelemetryLogicalLinkAllOf) HasLinkState() bool {
	if o != nil && o.LinkState != nil {
		return true
	}

	return false
}

// SetLinkState gets a reference to the given string and assigns it to the LinkState field.
func (o *NiatelemetryLogicalLinkAllOf) SetLinkState(v string) {
	o.LinkState = &v
}

// GetLinkType returns the LinkType field value if set, zero value otherwise.
func (o *NiatelemetryLogicalLinkAllOf) GetLinkType() string {
	if o == nil || o.LinkType == nil {
		var ret string
		return ret
	}
	return *o.LinkType
}

// GetLinkTypeOk returns a tuple with the LinkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLogicalLinkAllOf) GetLinkTypeOk() (*string, bool) {
	if o == nil || o.LinkType == nil {
		return nil, false
	}
	return o.LinkType, true
}

// HasLinkType returns a boolean if a field has been set.
func (o *NiatelemetryLogicalLinkAllOf) HasLinkType() bool {
	if o != nil && o.LinkType != nil {
		return true
	}

	return false
}

// SetLinkType gets a reference to the given string and assigns it to the LinkType field.
func (o *NiatelemetryLogicalLinkAllOf) SetLinkType(v string) {
	o.LinkType = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *NiatelemetryLogicalLinkAllOf) GetUptime() string {
	if o == nil || o.Uptime == nil {
		var ret string
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryLogicalLinkAllOf) GetUptimeOk() (*string, bool) {
	if o == nil || o.Uptime == nil {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *NiatelemetryLogicalLinkAllOf) HasUptime() bool {
	if o != nil && o.Uptime != nil {
		return true
	}

	return false
}

// SetUptime gets a reference to the given string and assigns it to the Uptime field.
func (o *NiatelemetryLogicalLinkAllOf) SetUptime(v string) {
	o.Uptime = &v
}

func (o NiatelemetryLogicalLinkAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DbId != nil {
		toSerialize["DbId"] = o.DbId
	}
	if o.IsPresent != nil {
		toSerialize["IsPresent"] = o.IsPresent
	}
	if o.LinkAddr1 != nil {
		toSerialize["LinkAddr1"] = o.LinkAddr1
	}
	if o.LinkAddr2 != nil {
		toSerialize["LinkAddr2"] = o.LinkAddr2
	}
	if o.LinkState != nil {
		toSerialize["LinkState"] = o.LinkState
	}
	if o.LinkType != nil {
		toSerialize["LinkType"] = o.LinkType
	}
	if o.Uptime != nil {
		toSerialize["Uptime"] = o.Uptime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryLogicalLinkAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryLogicalLinkAllOf := _NiatelemetryLogicalLinkAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryLogicalLinkAllOf); err == nil {
		*o = NiatelemetryLogicalLinkAllOf(varNiatelemetryLogicalLinkAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DbId")
		delete(additionalProperties, "IsPresent")
		delete(additionalProperties, "LinkAddr1")
		delete(additionalProperties, "LinkAddr2")
		delete(additionalProperties, "LinkState")
		delete(additionalProperties, "LinkType")
		delete(additionalProperties, "Uptime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryLogicalLinkAllOf struct {
	value *NiatelemetryLogicalLinkAllOf
	isSet bool
}

func (v NullableNiatelemetryLogicalLinkAllOf) Get() *NiatelemetryLogicalLinkAllOf {
	return v.value
}

func (v *NullableNiatelemetryLogicalLinkAllOf) Set(val *NiatelemetryLogicalLinkAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryLogicalLinkAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryLogicalLinkAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryLogicalLinkAllOf(val *NiatelemetryLogicalLinkAllOf) *NullableNiatelemetryLogicalLinkAllOf {
	return &NullableNiatelemetryLogicalLinkAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryLogicalLinkAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryLogicalLinkAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
