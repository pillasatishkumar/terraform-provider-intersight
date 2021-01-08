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

// VnicFcQosPolicy A Fibre Channel Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vHBA. This system class determines the quality of service for the outgoing traffic. For certain adapters additional controls can also be specified like burst and rate on the outgoing traffic.
type VnicFcQosPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Class of Service to be associated to the traffic on the virtual interface.
	Cos *int64 `json:"Cos,omitempty"`
	// The maximum size of the Fibre Channel frame payload bytes that the virtual interface supports.
	MaxDataFieldSize *int64 `json:"MaxDataFieldSize,omitempty"`
	// The priortity matching the System QoS specified in the fabric profile. * `Best Effort` - QoS Priority for Best-effort traffic. * `FC` - QoS Priority for FC traffic. * `Platinum` - QoS Priority for Platinum traffic. * `Gold` - QoS Priority for Gold traffic. * `Silver` - QoS Priority for Silver traffic. * `Bronze` - QoS Priority for Bronze traffic.
	Priority *string `json:"Priority,omitempty"`
	// The value in Mbps to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off.
	RateLimit            *int64                                `json:"RateLimit,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicFcQosPolicy VnicFcQosPolicy

// NewVnicFcQosPolicy instantiates a new VnicFcQosPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicFcQosPolicy(classId string, objectType string) *VnicFcQosPolicy {
	this := VnicFcQosPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	var cos int64 = 3
	this.Cos = &cos
	var maxDataFieldSize int64 = 2112
	this.MaxDataFieldSize = &maxDataFieldSize
	var priority string = "Best Effort"
	this.Priority = &priority
	var rateLimit int64 = 0
	this.RateLimit = &rateLimit
	return &this
}

// NewVnicFcQosPolicyWithDefaults instantiates a new VnicFcQosPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicFcQosPolicyWithDefaults() *VnicFcQosPolicy {
	this := VnicFcQosPolicy{}
	var classId string = "vnic.FcQosPolicy"
	this.ClassId = classId
	var objectType string = "vnic.FcQosPolicy"
	this.ObjectType = objectType
	var cos int64 = 3
	this.Cos = &cos
	var maxDataFieldSize int64 = 2112
	this.MaxDataFieldSize = &maxDataFieldSize
	var priority string = "Best Effort"
	this.Priority = &priority
	var rateLimit int64 = 0
	this.RateLimit = &rateLimit
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicFcQosPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicFcQosPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicFcQosPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicFcQosPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCos returns the Cos field value if set, zero value otherwise.
func (o *VnicFcQosPolicy) GetCos() int64 {
	if o == nil || o.Cos == nil {
		var ret int64
		return ret
	}
	return *o.Cos
}

// GetCosOk returns a tuple with the Cos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicy) GetCosOk() (*int64, bool) {
	if o == nil || o.Cos == nil {
		return nil, false
	}
	return o.Cos, true
}

// HasCos returns a boolean if a field has been set.
func (o *VnicFcQosPolicy) HasCos() bool {
	if o != nil && o.Cos != nil {
		return true
	}

	return false
}

// SetCos gets a reference to the given int64 and assigns it to the Cos field.
func (o *VnicFcQosPolicy) SetCos(v int64) {
	o.Cos = &v
}

// GetMaxDataFieldSize returns the MaxDataFieldSize field value if set, zero value otherwise.
func (o *VnicFcQosPolicy) GetMaxDataFieldSize() int64 {
	if o == nil || o.MaxDataFieldSize == nil {
		var ret int64
		return ret
	}
	return *o.MaxDataFieldSize
}

// GetMaxDataFieldSizeOk returns a tuple with the MaxDataFieldSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicy) GetMaxDataFieldSizeOk() (*int64, bool) {
	if o == nil || o.MaxDataFieldSize == nil {
		return nil, false
	}
	return o.MaxDataFieldSize, true
}

// HasMaxDataFieldSize returns a boolean if a field has been set.
func (o *VnicFcQosPolicy) HasMaxDataFieldSize() bool {
	if o != nil && o.MaxDataFieldSize != nil {
		return true
	}

	return false
}

// SetMaxDataFieldSize gets a reference to the given int64 and assigns it to the MaxDataFieldSize field.
func (o *VnicFcQosPolicy) SetMaxDataFieldSize(v int64) {
	o.MaxDataFieldSize = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *VnicFcQosPolicy) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicy) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *VnicFcQosPolicy) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *VnicFcQosPolicy) SetPriority(v string) {
	o.Priority = &v
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise.
func (o *VnicFcQosPolicy) GetRateLimit() int64 {
	if o == nil || o.RateLimit == nil {
		var ret int64
		return ret
	}
	return *o.RateLimit
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicy) GetRateLimitOk() (*int64, bool) {
	if o == nil || o.RateLimit == nil {
		return nil, false
	}
	return o.RateLimit, true
}

// HasRateLimit returns a boolean if a field has been set.
func (o *VnicFcQosPolicy) HasRateLimit() bool {
	if o != nil && o.RateLimit != nil {
		return true
	}

	return false
}

// SetRateLimit gets a reference to the given int64 and assigns it to the RateLimit field.
func (o *VnicFcQosPolicy) SetRateLimit(v int64) {
	o.RateLimit = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *VnicFcQosPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcQosPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *VnicFcQosPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VnicFcQosPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o VnicFcQosPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Cos != nil {
		toSerialize["Cos"] = o.Cos
	}
	if o.MaxDataFieldSize != nil {
		toSerialize["MaxDataFieldSize"] = o.MaxDataFieldSize
	}
	if o.Priority != nil {
		toSerialize["Priority"] = o.Priority
	}
	if o.RateLimit != nil {
		toSerialize["RateLimit"] = o.RateLimit
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicFcQosPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type VnicFcQosPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Class of Service to be associated to the traffic on the virtual interface.
		Cos *int64 `json:"Cos,omitempty"`
		// The maximum size of the Fibre Channel frame payload bytes that the virtual interface supports.
		MaxDataFieldSize *int64 `json:"MaxDataFieldSize,omitempty"`
		// The priortity matching the System QoS specified in the fabric profile. * `Best Effort` - QoS Priority for Best-effort traffic. * `FC` - QoS Priority for FC traffic. * `Platinum` - QoS Priority for Platinum traffic. * `Gold` - QoS Priority for Gold traffic. * `Silver` - QoS Priority for Silver traffic. * `Bronze` - QoS Priority for Bronze traffic.
		Priority *string `json:"Priority,omitempty"`
		// The value in Mbps to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off.
		RateLimit    *int64                                `json:"RateLimit,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varVnicFcQosPolicyWithoutEmbeddedStruct := VnicFcQosPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicFcQosPolicyWithoutEmbeddedStruct)
	if err == nil {
		varVnicFcQosPolicy := _VnicFcQosPolicy{}
		varVnicFcQosPolicy.ClassId = varVnicFcQosPolicyWithoutEmbeddedStruct.ClassId
		varVnicFcQosPolicy.ObjectType = varVnicFcQosPolicyWithoutEmbeddedStruct.ObjectType
		varVnicFcQosPolicy.Cos = varVnicFcQosPolicyWithoutEmbeddedStruct.Cos
		varVnicFcQosPolicy.MaxDataFieldSize = varVnicFcQosPolicyWithoutEmbeddedStruct.MaxDataFieldSize
		varVnicFcQosPolicy.Priority = varVnicFcQosPolicyWithoutEmbeddedStruct.Priority
		varVnicFcQosPolicy.RateLimit = varVnicFcQosPolicyWithoutEmbeddedStruct.RateLimit
		varVnicFcQosPolicy.Organization = varVnicFcQosPolicyWithoutEmbeddedStruct.Organization
		*o = VnicFcQosPolicy(varVnicFcQosPolicy)
	} else {
		return err
	}

	varVnicFcQosPolicy := _VnicFcQosPolicy{}

	err = json.Unmarshal(bytes, &varVnicFcQosPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varVnicFcQosPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cos")
		delete(additionalProperties, "MaxDataFieldSize")
		delete(additionalProperties, "Priority")
		delete(additionalProperties, "RateLimit")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableVnicFcQosPolicy struct {
	value *VnicFcQosPolicy
	isSet bool
}

func (v NullableVnicFcQosPolicy) Get() *VnicFcQosPolicy {
	return v.value
}

func (v *NullableVnicFcQosPolicy) Set(val *VnicFcQosPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicFcQosPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicFcQosPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicFcQosPolicy(val *VnicFcQosPolicy) *NullableVnicFcQosPolicy {
	return &NullableVnicFcQosPolicy{value: val, isSet: true}
}

func (v NullableVnicFcQosPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicFcQosPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
