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

// StorageHitachiParityGroupAllOf Definition of the list of properties defined in 'storage.HitachiParityGroup', excluding properties defined in parent classes.
type StorageHitachiParityGroupAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Speed (rpm) of the disk belonging to the parity group.
	DiskSpeed *string `json:"DiskSpeed,omitempty"`
	// Type of the disk belonging to the parity group.
	DiskType *string `json:"DiskType,omitempty"`
	// Value of the accelerated compression of the parity group. true, Accelerated compression for the parity group is enabled. false, Accelerated compression for the parity group is disabled.
	IsAcceleratedCompressionEnabled *bool `json:"IsAcceleratedCompressionEnabled,omitempty"`
	// Value of the copy back mode setting of the parity group. true, Copy back mode is enabled. false, Copy back mode is disabled.
	IsCopyBackModeEnabled *bool `json:"IsCopyBackModeEnabled,omitempty"`
	// Value of the encryption setting of the parity group. true, Encryption is enabled. false, Encryption is disabled.
	IsEncryptionEnabled  *bool                                `json:"IsEncryptionEnabled,omitempty"`
	Array                *StorageHitachiArrayRelationship     `json:"Array,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageHitachiParityGroupAllOf StorageHitachiParityGroupAllOf

// NewStorageHitachiParityGroupAllOf instantiates a new StorageHitachiParityGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageHitachiParityGroupAllOf(classId string, objectType string) *StorageHitachiParityGroupAllOf {
	this := StorageHitachiParityGroupAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageHitachiParityGroupAllOfWithDefaults instantiates a new StorageHitachiParityGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageHitachiParityGroupAllOfWithDefaults() *StorageHitachiParityGroupAllOf {
	this := StorageHitachiParityGroupAllOf{}
	var classId string = "storage.HitachiParityGroup"
	this.ClassId = classId
	var objectType string = "storage.HitachiParityGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageHitachiParityGroupAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroupAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageHitachiParityGroupAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageHitachiParityGroupAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroupAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageHitachiParityGroupAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDiskSpeed returns the DiskSpeed field value if set, zero value otherwise.
func (o *StorageHitachiParityGroupAllOf) GetDiskSpeed() string {
	if o == nil || o.DiskSpeed == nil {
		var ret string
		return ret
	}
	return *o.DiskSpeed
}

// GetDiskSpeedOk returns a tuple with the DiskSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroupAllOf) GetDiskSpeedOk() (*string, bool) {
	if o == nil || o.DiskSpeed == nil {
		return nil, false
	}
	return o.DiskSpeed, true
}

// HasDiskSpeed returns a boolean if a field has been set.
func (o *StorageHitachiParityGroupAllOf) HasDiskSpeed() bool {
	if o != nil && o.DiskSpeed != nil {
		return true
	}

	return false
}

// SetDiskSpeed gets a reference to the given string and assigns it to the DiskSpeed field.
func (o *StorageHitachiParityGroupAllOf) SetDiskSpeed(v string) {
	o.DiskSpeed = &v
}

// GetDiskType returns the DiskType field value if set, zero value otherwise.
func (o *StorageHitachiParityGroupAllOf) GetDiskType() string {
	if o == nil || o.DiskType == nil {
		var ret string
		return ret
	}
	return *o.DiskType
}

// GetDiskTypeOk returns a tuple with the DiskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroupAllOf) GetDiskTypeOk() (*string, bool) {
	if o == nil || o.DiskType == nil {
		return nil, false
	}
	return o.DiskType, true
}

// HasDiskType returns a boolean if a field has been set.
func (o *StorageHitachiParityGroupAllOf) HasDiskType() bool {
	if o != nil && o.DiskType != nil {
		return true
	}

	return false
}

// SetDiskType gets a reference to the given string and assigns it to the DiskType field.
func (o *StorageHitachiParityGroupAllOf) SetDiskType(v string) {
	o.DiskType = &v
}

// GetIsAcceleratedCompressionEnabled returns the IsAcceleratedCompressionEnabled field value if set, zero value otherwise.
func (o *StorageHitachiParityGroupAllOf) GetIsAcceleratedCompressionEnabled() bool {
	if o == nil || o.IsAcceleratedCompressionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsAcceleratedCompressionEnabled
}

// GetIsAcceleratedCompressionEnabledOk returns a tuple with the IsAcceleratedCompressionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroupAllOf) GetIsAcceleratedCompressionEnabledOk() (*bool, bool) {
	if o == nil || o.IsAcceleratedCompressionEnabled == nil {
		return nil, false
	}
	return o.IsAcceleratedCompressionEnabled, true
}

// HasIsAcceleratedCompressionEnabled returns a boolean if a field has been set.
func (o *StorageHitachiParityGroupAllOf) HasIsAcceleratedCompressionEnabled() bool {
	if o != nil && o.IsAcceleratedCompressionEnabled != nil {
		return true
	}

	return false
}

// SetIsAcceleratedCompressionEnabled gets a reference to the given bool and assigns it to the IsAcceleratedCompressionEnabled field.
func (o *StorageHitachiParityGroupAllOf) SetIsAcceleratedCompressionEnabled(v bool) {
	o.IsAcceleratedCompressionEnabled = &v
}

// GetIsCopyBackModeEnabled returns the IsCopyBackModeEnabled field value if set, zero value otherwise.
func (o *StorageHitachiParityGroupAllOf) GetIsCopyBackModeEnabled() bool {
	if o == nil || o.IsCopyBackModeEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsCopyBackModeEnabled
}

// GetIsCopyBackModeEnabledOk returns a tuple with the IsCopyBackModeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroupAllOf) GetIsCopyBackModeEnabledOk() (*bool, bool) {
	if o == nil || o.IsCopyBackModeEnabled == nil {
		return nil, false
	}
	return o.IsCopyBackModeEnabled, true
}

// HasIsCopyBackModeEnabled returns a boolean if a field has been set.
func (o *StorageHitachiParityGroupAllOf) HasIsCopyBackModeEnabled() bool {
	if o != nil && o.IsCopyBackModeEnabled != nil {
		return true
	}

	return false
}

// SetIsCopyBackModeEnabled gets a reference to the given bool and assigns it to the IsCopyBackModeEnabled field.
func (o *StorageHitachiParityGroupAllOf) SetIsCopyBackModeEnabled(v bool) {
	o.IsCopyBackModeEnabled = &v
}

// GetIsEncryptionEnabled returns the IsEncryptionEnabled field value if set, zero value otherwise.
func (o *StorageHitachiParityGroupAllOf) GetIsEncryptionEnabled() bool {
	if o == nil || o.IsEncryptionEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEncryptionEnabled
}

// GetIsEncryptionEnabledOk returns a tuple with the IsEncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroupAllOf) GetIsEncryptionEnabledOk() (*bool, bool) {
	if o == nil || o.IsEncryptionEnabled == nil {
		return nil, false
	}
	return o.IsEncryptionEnabled, true
}

// HasIsEncryptionEnabled returns a boolean if a field has been set.
func (o *StorageHitachiParityGroupAllOf) HasIsEncryptionEnabled() bool {
	if o != nil && o.IsEncryptionEnabled != nil {
		return true
	}

	return false
}

// SetIsEncryptionEnabled gets a reference to the given bool and assigns it to the IsEncryptionEnabled field.
func (o *StorageHitachiParityGroupAllOf) SetIsEncryptionEnabled(v bool) {
	o.IsEncryptionEnabled = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageHitachiParityGroupAllOf) GetArray() StorageHitachiArrayRelationship {
	if o == nil || o.Array == nil {
		var ret StorageHitachiArrayRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroupAllOf) GetArrayOk() (*StorageHitachiArrayRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageHitachiParityGroupAllOf) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageHitachiArrayRelationship and assigns it to the Array field.
func (o *StorageHitachiParityGroupAllOf) SetArray(v StorageHitachiArrayRelationship) {
	o.Array = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageHitachiParityGroupAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageHitachiParityGroupAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageHitachiParityGroupAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageHitachiParityGroupAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o StorageHitachiParityGroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DiskSpeed != nil {
		toSerialize["DiskSpeed"] = o.DiskSpeed
	}
	if o.DiskType != nil {
		toSerialize["DiskType"] = o.DiskType
	}
	if o.IsAcceleratedCompressionEnabled != nil {
		toSerialize["IsAcceleratedCompressionEnabled"] = o.IsAcceleratedCompressionEnabled
	}
	if o.IsCopyBackModeEnabled != nil {
		toSerialize["IsCopyBackModeEnabled"] = o.IsCopyBackModeEnabled
	}
	if o.IsEncryptionEnabled != nil {
		toSerialize["IsEncryptionEnabled"] = o.IsEncryptionEnabled
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageHitachiParityGroupAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageHitachiParityGroupAllOf := _StorageHitachiParityGroupAllOf{}

	if err = json.Unmarshal(bytes, &varStorageHitachiParityGroupAllOf); err == nil {
		*o = StorageHitachiParityGroupAllOf(varStorageHitachiParityGroupAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DiskSpeed")
		delete(additionalProperties, "DiskType")
		delete(additionalProperties, "IsAcceleratedCompressionEnabled")
		delete(additionalProperties, "IsCopyBackModeEnabled")
		delete(additionalProperties, "IsEncryptionEnabled")
		delete(additionalProperties, "Array")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageHitachiParityGroupAllOf struct {
	value *StorageHitachiParityGroupAllOf
	isSet bool
}

func (v NullableStorageHitachiParityGroupAllOf) Get() *StorageHitachiParityGroupAllOf {
	return v.value
}

func (v *NullableStorageHitachiParityGroupAllOf) Set(val *StorageHitachiParityGroupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageHitachiParityGroupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageHitachiParityGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageHitachiParityGroupAllOf(val *StorageHitachiParityGroupAllOf) *NullableStorageHitachiParityGroupAllOf {
	return &NullableStorageHitachiParityGroupAllOf{value: val, isSet: true}
}

func (v NullableStorageHitachiParityGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageHitachiParityGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
