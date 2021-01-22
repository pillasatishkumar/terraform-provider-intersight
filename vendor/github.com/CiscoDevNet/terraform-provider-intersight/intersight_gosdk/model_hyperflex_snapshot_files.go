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

// HyperflexSnapshotFiles Information about files tracked with the snapshot.
type HyperflexSnapshotFiles struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                          `json:"ObjectType"`
	NameTrackedFiles     []HyperflexFilePath             `json:"NameTrackedFiles,omitempty"`
	UuidTrackedDisksMap  []HyperflexMapUuidToTrackedDisk `json:"UuidTrackedDisksMap,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexSnapshotFiles HyperflexSnapshotFiles

// NewHyperflexSnapshotFiles instantiates a new HyperflexSnapshotFiles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexSnapshotFiles(classId string, objectType string) *HyperflexSnapshotFiles {
	this := HyperflexSnapshotFiles{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexSnapshotFilesWithDefaults instantiates a new HyperflexSnapshotFiles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexSnapshotFilesWithDefaults() *HyperflexSnapshotFiles {
	this := HyperflexSnapshotFiles{}
	var classId string = "hyperflex.SnapshotFiles"
	this.ClassId = classId
	var objectType string = "hyperflex.SnapshotFiles"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexSnapshotFiles) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotFiles) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexSnapshotFiles) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexSnapshotFiles) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexSnapshotFiles) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexSnapshotFiles) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNameTrackedFiles returns the NameTrackedFiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSnapshotFiles) GetNameTrackedFiles() []HyperflexFilePath {
	if o == nil {
		var ret []HyperflexFilePath
		return ret
	}
	return o.NameTrackedFiles
}

// GetNameTrackedFilesOk returns a tuple with the NameTrackedFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSnapshotFiles) GetNameTrackedFilesOk() (*[]HyperflexFilePath, bool) {
	if o == nil || o.NameTrackedFiles == nil {
		return nil, false
	}
	return &o.NameTrackedFiles, true
}

// HasNameTrackedFiles returns a boolean if a field has been set.
func (o *HyperflexSnapshotFiles) HasNameTrackedFiles() bool {
	if o != nil && o.NameTrackedFiles != nil {
		return true
	}

	return false
}

// SetNameTrackedFiles gets a reference to the given []HyperflexFilePath and assigns it to the NameTrackedFiles field.
func (o *HyperflexSnapshotFiles) SetNameTrackedFiles(v []HyperflexFilePath) {
	o.NameTrackedFiles = v
}

// GetUuidTrackedDisksMap returns the UuidTrackedDisksMap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexSnapshotFiles) GetUuidTrackedDisksMap() []HyperflexMapUuidToTrackedDisk {
	if o == nil {
		var ret []HyperflexMapUuidToTrackedDisk
		return ret
	}
	return o.UuidTrackedDisksMap
}

// GetUuidTrackedDisksMapOk returns a tuple with the UuidTrackedDisksMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexSnapshotFiles) GetUuidTrackedDisksMapOk() (*[]HyperflexMapUuidToTrackedDisk, bool) {
	if o == nil || o.UuidTrackedDisksMap == nil {
		return nil, false
	}
	return &o.UuidTrackedDisksMap, true
}

// HasUuidTrackedDisksMap returns a boolean if a field has been set.
func (o *HyperflexSnapshotFiles) HasUuidTrackedDisksMap() bool {
	if o != nil && o.UuidTrackedDisksMap != nil {
		return true
	}

	return false
}

// SetUuidTrackedDisksMap gets a reference to the given []HyperflexMapUuidToTrackedDisk and assigns it to the UuidTrackedDisksMap field.
func (o *HyperflexSnapshotFiles) SetUuidTrackedDisksMap(v []HyperflexMapUuidToTrackedDisk) {
	o.UuidTrackedDisksMap = v
}

func (o HyperflexSnapshotFiles) MarshalJSON() ([]byte, error) {
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
	if o.NameTrackedFiles != nil {
		toSerialize["NameTrackedFiles"] = o.NameTrackedFiles
	}
	if o.UuidTrackedDisksMap != nil {
		toSerialize["UuidTrackedDisksMap"] = o.UuidTrackedDisksMap
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexSnapshotFiles) UnmarshalJSON(bytes []byte) (err error) {
	type HyperflexSnapshotFilesWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType          string                          `json:"ObjectType"`
		NameTrackedFiles    []HyperflexFilePath             `json:"NameTrackedFiles,omitempty"`
		UuidTrackedDisksMap []HyperflexMapUuidToTrackedDisk `json:"UuidTrackedDisksMap,omitempty"`
	}

	varHyperflexSnapshotFilesWithoutEmbeddedStruct := HyperflexSnapshotFilesWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varHyperflexSnapshotFilesWithoutEmbeddedStruct)
	if err == nil {
		varHyperflexSnapshotFiles := _HyperflexSnapshotFiles{}
		varHyperflexSnapshotFiles.ClassId = varHyperflexSnapshotFilesWithoutEmbeddedStruct.ClassId
		varHyperflexSnapshotFiles.ObjectType = varHyperflexSnapshotFilesWithoutEmbeddedStruct.ObjectType
		varHyperflexSnapshotFiles.NameTrackedFiles = varHyperflexSnapshotFilesWithoutEmbeddedStruct.NameTrackedFiles
		varHyperflexSnapshotFiles.UuidTrackedDisksMap = varHyperflexSnapshotFilesWithoutEmbeddedStruct.UuidTrackedDisksMap
		*o = HyperflexSnapshotFiles(varHyperflexSnapshotFiles)
	} else {
		return err
	}

	varHyperflexSnapshotFiles := _HyperflexSnapshotFiles{}

	err = json.Unmarshal(bytes, &varHyperflexSnapshotFiles)
	if err == nil {
		o.MoBaseComplexType = varHyperflexSnapshotFiles.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "NameTrackedFiles")
		delete(additionalProperties, "UuidTrackedDisksMap")

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

type NullableHyperflexSnapshotFiles struct {
	value *HyperflexSnapshotFiles
	isSet bool
}

func (v NullableHyperflexSnapshotFiles) Get() *HyperflexSnapshotFiles {
	return v.value
}

func (v *NullableHyperflexSnapshotFiles) Set(val *HyperflexSnapshotFiles) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexSnapshotFiles) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexSnapshotFiles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexSnapshotFiles(val *HyperflexSnapshotFiles) *NullableHyperflexSnapshotFiles {
	return &NullableHyperflexSnapshotFiles{value: val, isSet: true}
}

func (v NullableHyperflexSnapshotFiles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexSnapshotFiles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
