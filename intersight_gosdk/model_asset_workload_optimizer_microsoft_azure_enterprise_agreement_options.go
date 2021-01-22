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

// AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions Captures configuration specific to the Microsoft Azure Enterprise Agreement target for the Workload Optimizer service.
type AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions struct {
	AssetServiceOptions
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Enrollment number for Azure EA.
	EnrollmentNumber     *string `json:"EnrollmentNumber,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions

// NewAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions instantiates a new AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions(classId string, objectType string) *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions {
	this := AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptionsWithDefaults instantiates a new AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptionsWithDefaults() *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions {
	this := AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions{}
	var classId string = "asset.WorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions"
	this.ClassId = classId
	var objectType string = "asset.WorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnrollmentNumber returns the EnrollmentNumber field value if set, zero value otherwise.
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) GetEnrollmentNumber() string {
	if o == nil || o.EnrollmentNumber == nil {
		var ret string
		return ret
	}
	return *o.EnrollmentNumber
}

// GetEnrollmentNumberOk returns a tuple with the EnrollmentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) GetEnrollmentNumberOk() (*string, bool) {
	if o == nil || o.EnrollmentNumber == nil {
		return nil, false
	}
	return o.EnrollmentNumber, true
}

// HasEnrollmentNumber returns a boolean if a field has been set.
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) HasEnrollmentNumber() bool {
	if o != nil && o.EnrollmentNumber != nil {
		return true
	}

	return false
}

// SetEnrollmentNumber gets a reference to the given string and assigns it to the EnrollmentNumber field.
func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) SetEnrollmentNumber(v string) {
	o.EnrollmentNumber = &v
}

func (o AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetServiceOptions, errAssetServiceOptions := json.Marshal(o.AssetServiceOptions)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	errAssetServiceOptions = json.Unmarshal([]byte(serializedAssetServiceOptions), &toSerialize)
	if errAssetServiceOptions != nil {
		return []byte{}, errAssetServiceOptions
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EnrollmentNumber != nil {
		toSerialize["EnrollmentNumber"] = o.EnrollmentNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) UnmarshalJSON(bytes []byte) (err error) {
	type AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptionsWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Enrollment number for Azure EA.
		EnrollmentNumber *string `json:"EnrollmentNumber,omitempty"`
	}

	varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptionsWithoutEmbeddedStruct := AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptionsWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptionsWithoutEmbeddedStruct)
	if err == nil {
		varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions := _AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions{}
		varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions.ClassId = varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptionsWithoutEmbeddedStruct.ClassId
		varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions.ObjectType = varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptionsWithoutEmbeddedStruct.ObjectType
		varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions.EnrollmentNumber = varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptionsWithoutEmbeddedStruct.EnrollmentNumber
		*o = AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions(varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions)
	} else {
		return err
	}

	varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions := _AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions{}

	err = json.Unmarshal(bytes, &varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions)
	if err == nil {
		o.AssetServiceOptions = varAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions.AssetServiceOptions
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnrollmentNumber")

		// remove fields from embedded structs
		reflectAssetServiceOptions := reflect.ValueOf(o.AssetServiceOptions)
		for i := 0; i < reflectAssetServiceOptions.Type().NumField(); i++ {
			t := reflectAssetServiceOptions.Type().Field(i)

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

type NullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions struct {
	value *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions
	isSet bool
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) Get() *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions {
	return v.value
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) Set(val *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions(val *AssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) *NullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions {
	return &NullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions{value: val, isSet: true}
}

func (v NullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWorkloadOptimizerMicrosoftAzureEnterpriseAgreementOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
