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

// ChassisConfigImportAllOf Definition of the list of properties defined in 'chassis.ConfigImport', excluding properties defined in parent classes.
type ChassisConfigImportAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of the imported profile.
	Description *string `json:"Description,omitempty"`
	// Policy prefix for the policies of the imported chassis profile.
	PolicyPrefix *string  `json:"PolicyPrefix,omitempty"`
	PolicyTypes  []string `json:"PolicyTypes,omitempty"`
	// Profile name for the imported chassis profile.
	ProfileName          *string                               `json:"ProfileName,omitempty"`
	Chassis              *EquipmentChassisRelationship         `json:"Chassis,omitempty"`
	ChassisProfile       *ChassisProfileRelationship           `json:"ChassisProfile,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChassisConfigImportAllOf ChassisConfigImportAllOf

// NewChassisConfigImportAllOf instantiates a new ChassisConfigImportAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChassisConfigImportAllOf(classId string, objectType string) *ChassisConfigImportAllOf {
	this := ChassisConfigImportAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewChassisConfigImportAllOfWithDefaults instantiates a new ChassisConfigImportAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChassisConfigImportAllOfWithDefaults() *ChassisConfigImportAllOf {
	this := ChassisConfigImportAllOf{}
	var classId string = "chassis.ConfigImport"
	this.ClassId = classId
	var objectType string = "chassis.ConfigImport"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ChassisConfigImportAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ChassisConfigImportAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ChassisConfigImportAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ChassisConfigImportAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ChassisConfigImportAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ChassisConfigImportAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ChassisConfigImportAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisConfigImportAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ChassisConfigImportAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ChassisConfigImportAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetPolicyPrefix returns the PolicyPrefix field value if set, zero value otherwise.
func (o *ChassisConfigImportAllOf) GetPolicyPrefix() string {
	if o == nil || o.PolicyPrefix == nil {
		var ret string
		return ret
	}
	return *o.PolicyPrefix
}

// GetPolicyPrefixOk returns a tuple with the PolicyPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisConfigImportAllOf) GetPolicyPrefixOk() (*string, bool) {
	if o == nil || o.PolicyPrefix == nil {
		return nil, false
	}
	return o.PolicyPrefix, true
}

// HasPolicyPrefix returns a boolean if a field has been set.
func (o *ChassisConfigImportAllOf) HasPolicyPrefix() bool {
	if o != nil && o.PolicyPrefix != nil {
		return true
	}

	return false
}

// SetPolicyPrefix gets a reference to the given string and assigns it to the PolicyPrefix field.
func (o *ChassisConfigImportAllOf) SetPolicyPrefix(v string) {
	o.PolicyPrefix = &v
}

// GetPolicyTypes returns the PolicyTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisConfigImportAllOf) GetPolicyTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PolicyTypes
}

// GetPolicyTypesOk returns a tuple with the PolicyTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisConfigImportAllOf) GetPolicyTypesOk() (*[]string, bool) {
	if o == nil || o.PolicyTypes == nil {
		return nil, false
	}
	return &o.PolicyTypes, true
}

// HasPolicyTypes returns a boolean if a field has been set.
func (o *ChassisConfigImportAllOf) HasPolicyTypes() bool {
	if o != nil && o.PolicyTypes != nil {
		return true
	}

	return false
}

// SetPolicyTypes gets a reference to the given []string and assigns it to the PolicyTypes field.
func (o *ChassisConfigImportAllOf) SetPolicyTypes(v []string) {
	o.PolicyTypes = v
}

// GetProfileName returns the ProfileName field value if set, zero value otherwise.
func (o *ChassisConfigImportAllOf) GetProfileName() string {
	if o == nil || o.ProfileName == nil {
		var ret string
		return ret
	}
	return *o.ProfileName
}

// GetProfileNameOk returns a tuple with the ProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisConfigImportAllOf) GetProfileNameOk() (*string, bool) {
	if o == nil || o.ProfileName == nil {
		return nil, false
	}
	return o.ProfileName, true
}

// HasProfileName returns a boolean if a field has been set.
func (o *ChassisConfigImportAllOf) HasProfileName() bool {
	if o != nil && o.ProfileName != nil {
		return true
	}

	return false
}

// SetProfileName gets a reference to the given string and assigns it to the ProfileName field.
func (o *ChassisConfigImportAllOf) SetProfileName(v string) {
	o.ProfileName = &v
}

// GetChassis returns the Chassis field value if set, zero value otherwise.
func (o *ChassisConfigImportAllOf) GetChassis() EquipmentChassisRelationship {
	if o == nil || o.Chassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.Chassis
}

// GetChassisOk returns a tuple with the Chassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisConfigImportAllOf) GetChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.Chassis == nil {
		return nil, false
	}
	return o.Chassis, true
}

// HasChassis returns a boolean if a field has been set.
func (o *ChassisConfigImportAllOf) HasChassis() bool {
	if o != nil && o.Chassis != nil {
		return true
	}

	return false
}

// SetChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the Chassis field.
func (o *ChassisConfigImportAllOf) SetChassis(v EquipmentChassisRelationship) {
	o.Chassis = &v
}

// GetChassisProfile returns the ChassisProfile field value if set, zero value otherwise.
func (o *ChassisConfigImportAllOf) GetChassisProfile() ChassisProfileRelationship {
	if o == nil || o.ChassisProfile == nil {
		var ret ChassisProfileRelationship
		return ret
	}
	return *o.ChassisProfile
}

// GetChassisProfileOk returns a tuple with the ChassisProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisConfigImportAllOf) GetChassisProfileOk() (*ChassisProfileRelationship, bool) {
	if o == nil || o.ChassisProfile == nil {
		return nil, false
	}
	return o.ChassisProfile, true
}

// HasChassisProfile returns a boolean if a field has been set.
func (o *ChassisConfigImportAllOf) HasChassisProfile() bool {
	if o != nil && o.ChassisProfile != nil {
		return true
	}

	return false
}

// SetChassisProfile gets a reference to the given ChassisProfileRelationship and assigns it to the ChassisProfile field.
func (o *ChassisConfigImportAllOf) SetChassisProfile(v ChassisProfileRelationship) {
	o.ChassisProfile = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ChassisConfigImportAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisConfigImportAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ChassisConfigImportAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ChassisConfigImportAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o ChassisConfigImportAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.PolicyPrefix != nil {
		toSerialize["PolicyPrefix"] = o.PolicyPrefix
	}
	if o.PolicyTypes != nil {
		toSerialize["PolicyTypes"] = o.PolicyTypes
	}
	if o.ProfileName != nil {
		toSerialize["ProfileName"] = o.ProfileName
	}
	if o.Chassis != nil {
		toSerialize["Chassis"] = o.Chassis
	}
	if o.ChassisProfile != nil {
		toSerialize["ChassisProfile"] = o.ChassisProfile
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChassisConfigImportAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varChassisConfigImportAllOf := _ChassisConfigImportAllOf{}

	if err = json.Unmarshal(bytes, &varChassisConfigImportAllOf); err == nil {
		*o = ChassisConfigImportAllOf(varChassisConfigImportAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "PolicyPrefix")
		delete(additionalProperties, "PolicyTypes")
		delete(additionalProperties, "ProfileName")
		delete(additionalProperties, "Chassis")
		delete(additionalProperties, "ChassisProfile")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChassisConfigImportAllOf struct {
	value *ChassisConfigImportAllOf
	isSet bool
}

func (v NullableChassisConfigImportAllOf) Get() *ChassisConfigImportAllOf {
	return v.value
}

func (v *NullableChassisConfigImportAllOf) Set(val *ChassisConfigImportAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableChassisConfigImportAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableChassisConfigImportAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChassisConfigImportAllOf(val *ChassisConfigImportAllOf) *NullableChassisConfigImportAllOf {
	return &NullableChassisConfigImportAllOf{value: val, isSet: true}
}

func (v NullableChassisConfigImportAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChassisConfigImportAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
