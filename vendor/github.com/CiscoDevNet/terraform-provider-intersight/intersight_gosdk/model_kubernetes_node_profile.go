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
	"reflect"
	"strings"
)

// KubernetesNodeProfile A profile specifying configuration settings for a Kubernetes node.
type KubernetesNodeProfile struct {
	PolicyAbstractProfile
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Cloud provider for this node profile. * `noProvider` - Enables the use of no cloud provider. * `external` - Out of tree cloud provider, e.g. CPI for vsphere.
	CloudProvider        *string                                 `json:"CloudProvider,omitempty"`
	NodeGroup            *KubernetesNodeGroupProfileRelationship `json:"NodeGroup,omitempty"`
	Target               *AssetDeviceRegistrationRelationship    `json:"Target,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesNodeProfile KubernetesNodeProfile

// NewKubernetesNodeProfile instantiates a new KubernetesNodeProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesNodeProfile(classId string, objectType string) *KubernetesNodeProfile {
	this := KubernetesNodeProfile{}
	this.ClassId = classId
	this.ObjectType = objectType
	var cloudProvider string = "noProvider"
	this.CloudProvider = &cloudProvider
	return &this
}

// NewKubernetesNodeProfileWithDefaults instantiates a new KubernetesNodeProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesNodeProfileWithDefaults() *KubernetesNodeProfile {
	this := KubernetesNodeProfile{}
	var classId string = "kubernetes.VirtualMachineNodeProfile"
	this.ClassId = classId
	var objectType string = "kubernetes.VirtualMachineNodeProfile"
	this.ObjectType = objectType
	var cloudProvider string = "noProvider"
	this.CloudProvider = &cloudProvider
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesNodeProfile) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeProfile) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesNodeProfile) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesNodeProfile) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesNodeProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesNodeProfile) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *KubernetesNodeProfile) GetCloudProvider() string {
	if o == nil || o.CloudProvider == nil {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeProfile) GetCloudProviderOk() (*string, bool) {
	if o == nil || o.CloudProvider == nil {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *KubernetesNodeProfile) HasCloudProvider() bool {
	if o != nil && o.CloudProvider != nil {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *KubernetesNodeProfile) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetNodeGroup returns the NodeGroup field value if set, zero value otherwise.
func (o *KubernetesNodeProfile) GetNodeGroup() KubernetesNodeGroupProfileRelationship {
	if o == nil || o.NodeGroup == nil {
		var ret KubernetesNodeGroupProfileRelationship
		return ret
	}
	return *o.NodeGroup
}

// GetNodeGroupOk returns a tuple with the NodeGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeProfile) GetNodeGroupOk() (*KubernetesNodeGroupProfileRelationship, bool) {
	if o == nil || o.NodeGroup == nil {
		return nil, false
	}
	return o.NodeGroup, true
}

// HasNodeGroup returns a boolean if a field has been set.
func (o *KubernetesNodeProfile) HasNodeGroup() bool {
	if o != nil && o.NodeGroup != nil {
		return true
	}

	return false
}

// SetNodeGroup gets a reference to the given KubernetesNodeGroupProfileRelationship and assigns it to the NodeGroup field.
func (o *KubernetesNodeProfile) SetNodeGroup(v KubernetesNodeGroupProfileRelationship) {
	o.NodeGroup = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *KubernetesNodeProfile) GetTarget() AssetDeviceRegistrationRelationship {
	if o == nil || o.Target == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesNodeProfile) GetTargetOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *KubernetesNodeProfile) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the Target field.
func (o *KubernetesNodeProfile) SetTarget(v AssetDeviceRegistrationRelationship) {
	o.Target = &v
}

func (o KubernetesNodeProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractProfile, errPolicyAbstractProfile := json.Marshal(o.PolicyAbstractProfile)
	if errPolicyAbstractProfile != nil {
		return []byte{}, errPolicyAbstractProfile
	}
	errPolicyAbstractProfile = json.Unmarshal([]byte(serializedPolicyAbstractProfile), &toSerialize)
	if errPolicyAbstractProfile != nil {
		return []byte{}, errPolicyAbstractProfile
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CloudProvider != nil {
		toSerialize["CloudProvider"] = o.CloudProvider
	}
	if o.NodeGroup != nil {
		toSerialize["NodeGroup"] = o.NodeGroup
	}
	if o.Target != nil {
		toSerialize["Target"] = o.Target
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesNodeProfile) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesNodeProfileWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Cloud provider for this node profile. * `noProvider` - Enables the use of no cloud provider. * `external` - Out of tree cloud provider, e.g. CPI for vsphere.
		CloudProvider *string                                 `json:"CloudProvider,omitempty"`
		NodeGroup     *KubernetesNodeGroupProfileRelationship `json:"NodeGroup,omitempty"`
		Target        *AssetDeviceRegistrationRelationship    `json:"Target,omitempty"`
	}

	varKubernetesNodeProfileWithoutEmbeddedStruct := KubernetesNodeProfileWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesNodeProfileWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesNodeProfile := _KubernetesNodeProfile{}
		varKubernetesNodeProfile.ClassId = varKubernetesNodeProfileWithoutEmbeddedStruct.ClassId
		varKubernetesNodeProfile.ObjectType = varKubernetesNodeProfileWithoutEmbeddedStruct.ObjectType
		varKubernetesNodeProfile.CloudProvider = varKubernetesNodeProfileWithoutEmbeddedStruct.CloudProvider
		varKubernetesNodeProfile.NodeGroup = varKubernetesNodeProfileWithoutEmbeddedStruct.NodeGroup
		varKubernetesNodeProfile.Target = varKubernetesNodeProfileWithoutEmbeddedStruct.Target
		*o = KubernetesNodeProfile(varKubernetesNodeProfile)
	} else {
		return err
	}

	varKubernetesNodeProfile := _KubernetesNodeProfile{}

	err = json.Unmarshal(bytes, &varKubernetesNodeProfile)
	if err == nil {
		o.PolicyAbstractProfile = varKubernetesNodeProfile.PolicyAbstractProfile
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CloudProvider")
		delete(additionalProperties, "NodeGroup")
		delete(additionalProperties, "Target")

		// remove fields from embedded structs
		reflectPolicyAbstractProfile := reflect.ValueOf(o.PolicyAbstractProfile)
		for i := 0; i < reflectPolicyAbstractProfile.Type().NumField(); i++ {
			t := reflectPolicyAbstractProfile.Type().Field(i)

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

type NullableKubernetesNodeProfile struct {
	value *KubernetesNodeProfile
	isSet bool
}

func (v NullableKubernetesNodeProfile) Get() *KubernetesNodeProfile {
	return v.value
}

func (v *NullableKubernetesNodeProfile) Set(val *KubernetesNodeProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesNodeProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesNodeProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesNodeProfile(val *KubernetesNodeProfile) *NullableKubernetesNodeProfile {
	return &NullableKubernetesNodeProfile{value: val, isSet: true}
}

func (v NullableKubernetesNodeProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesNodeProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
