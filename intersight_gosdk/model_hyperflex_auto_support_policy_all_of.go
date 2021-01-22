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

// HyperflexAutoSupportPolicyAllOf Definition of the list of properties defined in 'hyperflex.AutoSupportPolicy', excluding properties defined in parent classes.
type HyperflexAutoSupportPolicyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enable or disable Auto-Support.
	AdminState *bool `json:"AdminState,omitempty"`
	// The recipient email address for support tickets.
	ServiceTicketReceipient *string `json:"ServiceTicketReceipient,omitempty"`
	// An array of relationships to hyperflexClusterProfile resources.
	ClusterProfiles      []HyperflexClusterProfileRelationship `json:"ClusterProfiles,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexAutoSupportPolicyAllOf HyperflexAutoSupportPolicyAllOf

// NewHyperflexAutoSupportPolicyAllOf instantiates a new HyperflexAutoSupportPolicyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexAutoSupportPolicyAllOf(classId string, objectType string) *HyperflexAutoSupportPolicyAllOf {
	this := HyperflexAutoSupportPolicyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminState bool = true
	this.AdminState = &adminState
	return &this
}

// NewHyperflexAutoSupportPolicyAllOfWithDefaults instantiates a new HyperflexAutoSupportPolicyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexAutoSupportPolicyAllOfWithDefaults() *HyperflexAutoSupportPolicyAllOf {
	this := HyperflexAutoSupportPolicyAllOf{}
	var classId string = "hyperflex.AutoSupportPolicy"
	this.ClassId = classId
	var objectType string = "hyperflex.AutoSupportPolicy"
	this.ObjectType = objectType
	var adminState bool = true
	this.AdminState = &adminState
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexAutoSupportPolicyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexAutoSupportPolicyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexAutoSupportPolicyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexAutoSupportPolicyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexAutoSupportPolicyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexAutoSupportPolicyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *HyperflexAutoSupportPolicyAllOf) GetAdminState() bool {
	if o == nil || o.AdminState == nil {
		var ret bool
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAutoSupportPolicyAllOf) GetAdminStateOk() (*bool, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *HyperflexAutoSupportPolicyAllOf) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given bool and assigns it to the AdminState field.
func (o *HyperflexAutoSupportPolicyAllOf) SetAdminState(v bool) {
	o.AdminState = &v
}

// GetServiceTicketReceipient returns the ServiceTicketReceipient field value if set, zero value otherwise.
func (o *HyperflexAutoSupportPolicyAllOf) GetServiceTicketReceipient() string {
	if o == nil || o.ServiceTicketReceipient == nil {
		var ret string
		return ret
	}
	return *o.ServiceTicketReceipient
}

// GetServiceTicketReceipientOk returns a tuple with the ServiceTicketReceipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAutoSupportPolicyAllOf) GetServiceTicketReceipientOk() (*string, bool) {
	if o == nil || o.ServiceTicketReceipient == nil {
		return nil, false
	}
	return o.ServiceTicketReceipient, true
}

// HasServiceTicketReceipient returns a boolean if a field has been set.
func (o *HyperflexAutoSupportPolicyAllOf) HasServiceTicketReceipient() bool {
	if o != nil && o.ServiceTicketReceipient != nil {
		return true
	}

	return false
}

// SetServiceTicketReceipient gets a reference to the given string and assigns it to the ServiceTicketReceipient field.
func (o *HyperflexAutoSupportPolicyAllOf) SetServiceTicketReceipient(v string) {
	o.ServiceTicketReceipient = &v
}

// GetClusterProfiles returns the ClusterProfiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexAutoSupportPolicyAllOf) GetClusterProfiles() []HyperflexClusterProfileRelationship {
	if o == nil {
		var ret []HyperflexClusterProfileRelationship
		return ret
	}
	return o.ClusterProfiles
}

// GetClusterProfilesOk returns a tuple with the ClusterProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexAutoSupportPolicyAllOf) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool) {
	if o == nil || o.ClusterProfiles == nil {
		return nil, false
	}
	return &o.ClusterProfiles, true
}

// HasClusterProfiles returns a boolean if a field has been set.
func (o *HyperflexAutoSupportPolicyAllOf) HasClusterProfiles() bool {
	if o != nil && o.ClusterProfiles != nil {
		return true
	}

	return false
}

// SetClusterProfiles gets a reference to the given []HyperflexClusterProfileRelationship and assigns it to the ClusterProfiles field.
func (o *HyperflexAutoSupportPolicyAllOf) SetClusterProfiles(v []HyperflexClusterProfileRelationship) {
	o.ClusterProfiles = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *HyperflexAutoSupportPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexAutoSupportPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *HyperflexAutoSupportPolicyAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *HyperflexAutoSupportPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o HyperflexAutoSupportPolicyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.ServiceTicketReceipient != nil {
		toSerialize["ServiceTicketReceipient"] = o.ServiceTicketReceipient
	}
	if o.ClusterProfiles != nil {
		toSerialize["ClusterProfiles"] = o.ClusterProfiles
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexAutoSupportPolicyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexAutoSupportPolicyAllOf := _HyperflexAutoSupportPolicyAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexAutoSupportPolicyAllOf); err == nil {
		*o = HyperflexAutoSupportPolicyAllOf(varHyperflexAutoSupportPolicyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "ServiceTicketReceipient")
		delete(additionalProperties, "ClusterProfiles")
		delete(additionalProperties, "Organization")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexAutoSupportPolicyAllOf struct {
	value *HyperflexAutoSupportPolicyAllOf
	isSet bool
}

func (v NullableHyperflexAutoSupportPolicyAllOf) Get() *HyperflexAutoSupportPolicyAllOf {
	return v.value
}

func (v *NullableHyperflexAutoSupportPolicyAllOf) Set(val *HyperflexAutoSupportPolicyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexAutoSupportPolicyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexAutoSupportPolicyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexAutoSupportPolicyAllOf(val *HyperflexAutoSupportPolicyAllOf) *NullableHyperflexAutoSupportPolicyAllOf {
	return &NullableHyperflexAutoSupportPolicyAllOf{value: val, isSet: true}
}

func (v NullableHyperflexAutoSupportPolicyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexAutoSupportPolicyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
