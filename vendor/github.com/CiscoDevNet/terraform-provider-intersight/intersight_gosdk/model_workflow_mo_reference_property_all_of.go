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

// WorkflowMoReferencePropertyAllOf Definition of the list of properties defined in 'workflow.MoReferenceProperty', excluding properties defined in parent classes.
type WorkflowMoReferencePropertyAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType        string   `json:"ObjectType"`
	DisplayAttributes []string `json:"DisplayAttributes,omitempty"`
	// Field to hold an Intersight API along with an optional filter to narrow down the search options.
	Selector *string `json:"Selector,omitempty"`
	// A property from the Intersight object, value of which can be used as value for referenced input definition.
	ValueAttribute       *string `json:"ValueAttribute,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowMoReferencePropertyAllOf WorkflowMoReferencePropertyAllOf

// NewWorkflowMoReferencePropertyAllOf instantiates a new WorkflowMoReferencePropertyAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowMoReferencePropertyAllOf(classId string, objectType string) *WorkflowMoReferencePropertyAllOf {
	this := WorkflowMoReferencePropertyAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowMoReferencePropertyAllOfWithDefaults instantiates a new WorkflowMoReferencePropertyAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowMoReferencePropertyAllOfWithDefaults() *WorkflowMoReferencePropertyAllOf {
	this := WorkflowMoReferencePropertyAllOf{}
	var classId string = "workflow.MoReferenceProperty"
	this.ClassId = classId
	var objectType string = "workflow.MoReferenceProperty"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowMoReferencePropertyAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowMoReferencePropertyAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowMoReferencePropertyAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowMoReferencePropertyAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowMoReferencePropertyAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowMoReferencePropertyAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplayAttributes returns the DisplayAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowMoReferencePropertyAllOf) GetDisplayAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DisplayAttributes
}

// GetDisplayAttributesOk returns a tuple with the DisplayAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowMoReferencePropertyAllOf) GetDisplayAttributesOk() (*[]string, bool) {
	if o == nil || o.DisplayAttributes == nil {
		return nil, false
	}
	return &o.DisplayAttributes, true
}

// HasDisplayAttributes returns a boolean if a field has been set.
func (o *WorkflowMoReferencePropertyAllOf) HasDisplayAttributes() bool {
	if o != nil && o.DisplayAttributes != nil {
		return true
	}

	return false
}

// SetDisplayAttributes gets a reference to the given []string and assigns it to the DisplayAttributes field.
func (o *WorkflowMoReferencePropertyAllOf) SetDisplayAttributes(v []string) {
	o.DisplayAttributes = v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *WorkflowMoReferencePropertyAllOf) GetSelector() string {
	if o == nil || o.Selector == nil {
		var ret string
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowMoReferencePropertyAllOf) GetSelectorOk() (*string, bool) {
	if o == nil || o.Selector == nil {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *WorkflowMoReferencePropertyAllOf) HasSelector() bool {
	if o != nil && o.Selector != nil {
		return true
	}

	return false
}

// SetSelector gets a reference to the given string and assigns it to the Selector field.
func (o *WorkflowMoReferencePropertyAllOf) SetSelector(v string) {
	o.Selector = &v
}

// GetValueAttribute returns the ValueAttribute field value if set, zero value otherwise.
func (o *WorkflowMoReferencePropertyAllOf) GetValueAttribute() string {
	if o == nil || o.ValueAttribute == nil {
		var ret string
		return ret
	}
	return *o.ValueAttribute
}

// GetValueAttributeOk returns a tuple with the ValueAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowMoReferencePropertyAllOf) GetValueAttributeOk() (*string, bool) {
	if o == nil || o.ValueAttribute == nil {
		return nil, false
	}
	return o.ValueAttribute, true
}

// HasValueAttribute returns a boolean if a field has been set.
func (o *WorkflowMoReferencePropertyAllOf) HasValueAttribute() bool {
	if o != nil && o.ValueAttribute != nil {
		return true
	}

	return false
}

// SetValueAttribute gets a reference to the given string and assigns it to the ValueAttribute field.
func (o *WorkflowMoReferencePropertyAllOf) SetValueAttribute(v string) {
	o.ValueAttribute = &v
}

func (o WorkflowMoReferencePropertyAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DisplayAttributes != nil {
		toSerialize["DisplayAttributes"] = o.DisplayAttributes
	}
	if o.Selector != nil {
		toSerialize["Selector"] = o.Selector
	}
	if o.ValueAttribute != nil {
		toSerialize["ValueAttribute"] = o.ValueAttribute
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowMoReferencePropertyAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowMoReferencePropertyAllOf := _WorkflowMoReferencePropertyAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowMoReferencePropertyAllOf); err == nil {
		*o = WorkflowMoReferencePropertyAllOf(varWorkflowMoReferencePropertyAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DisplayAttributes")
		delete(additionalProperties, "Selector")
		delete(additionalProperties, "ValueAttribute")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowMoReferencePropertyAllOf struct {
	value *WorkflowMoReferencePropertyAllOf
	isSet bool
}

func (v NullableWorkflowMoReferencePropertyAllOf) Get() *WorkflowMoReferencePropertyAllOf {
	return v.value
}

func (v *NullableWorkflowMoReferencePropertyAllOf) Set(val *WorkflowMoReferencePropertyAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowMoReferencePropertyAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowMoReferencePropertyAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowMoReferencePropertyAllOf(val *WorkflowMoReferencePropertyAllOf) *NullableWorkflowMoReferencePropertyAllOf {
	return &NullableWorkflowMoReferencePropertyAllOf{value: val, isSet: true}
}

func (v NullableWorkflowMoReferencePropertyAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowMoReferencePropertyAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
