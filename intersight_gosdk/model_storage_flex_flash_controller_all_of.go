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

// StorageFlexFlashControllerAllOf Definition of the list of properties defined in 'storage.FlexFlashController', excluding properties defined in parent classes.
type StorageFlexFlashControllerAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// State of the Flex Flash Storage Controller.
	ControllerState *string `json:"ControllerState,omitempty"`
	// Identifier for the Flex Flash Storage Controller.
	FfControllerId *string                   `json:"FfControllerId,omitempty"`
	ComputeBoard   *ComputeBoardRelationship `json:"ComputeBoard,omitempty"`
	// An array of relationships to storageFlexFlashControllerProps resources.
	FlexFlashControllerProps []StorageFlexFlashControllerPropsRelationship `json:"FlexFlashControllerProps,omitempty"`
	// An array of relationships to storageFlexFlashPhysicalDrive resources.
	FlexFlashPhysicalDrives []StorageFlexFlashPhysicalDriveRelationship `json:"FlexFlashPhysicalDrives,omitempty"`
	// An array of relationships to storageFlexFlashVirtualDrive resources.
	FlexFlashVirtualDrives []StorageFlexFlashVirtualDriveRelationship `json:"FlexFlashVirtualDrives,omitempty"`
	InventoryDeviceInfo    *InventoryDeviceInfoRelationship           `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice       *AssetDeviceRegistrationRelationship       `json:"RegisteredDevice,omitempty"`
	// An array of relationships to firmwareRunningFirmware resources.
	RunningFirmware      []FirmwareRunningFirmwareRelationship `json:"RunningFirmware,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageFlexFlashControllerAllOf StorageFlexFlashControllerAllOf

// NewStorageFlexFlashControllerAllOf instantiates a new StorageFlexFlashControllerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageFlexFlashControllerAllOf(classId string, objectType string) *StorageFlexFlashControllerAllOf {
	this := StorageFlexFlashControllerAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageFlexFlashControllerAllOfWithDefaults instantiates a new StorageFlexFlashControllerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageFlexFlashControllerAllOfWithDefaults() *StorageFlexFlashControllerAllOf {
	this := StorageFlexFlashControllerAllOf{}
	var classId string = "storage.FlexFlashController"
	this.ClassId = classId
	var objectType string = "storage.FlexFlashController"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageFlexFlashControllerAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageFlexFlashControllerAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageFlexFlashControllerAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageFlexFlashControllerAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetControllerState returns the ControllerState field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerAllOf) GetControllerState() string {
	if o == nil || o.ControllerState == nil {
		var ret string
		return ret
	}
	return *o.ControllerState
}

// GetControllerStateOk returns a tuple with the ControllerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerAllOf) GetControllerStateOk() (*string, bool) {
	if o == nil || o.ControllerState == nil {
		return nil, false
	}
	return o.ControllerState, true
}

// HasControllerState returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerAllOf) HasControllerState() bool {
	if o != nil && o.ControllerState != nil {
		return true
	}

	return false
}

// SetControllerState gets a reference to the given string and assigns it to the ControllerState field.
func (o *StorageFlexFlashControllerAllOf) SetControllerState(v string) {
	o.ControllerState = &v
}

// GetFfControllerId returns the FfControllerId field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerAllOf) GetFfControllerId() string {
	if o == nil || o.FfControllerId == nil {
		var ret string
		return ret
	}
	return *o.FfControllerId
}

// GetFfControllerIdOk returns a tuple with the FfControllerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerAllOf) GetFfControllerIdOk() (*string, bool) {
	if o == nil || o.FfControllerId == nil {
		return nil, false
	}
	return o.FfControllerId, true
}

// HasFfControllerId returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerAllOf) HasFfControllerId() bool {
	if o != nil && o.FfControllerId != nil {
		return true
	}

	return false
}

// SetFfControllerId gets a reference to the given string and assigns it to the FfControllerId field.
func (o *StorageFlexFlashControllerAllOf) SetFfControllerId(v string) {
	o.FfControllerId = &v
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerAllOf) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || o.ComputeBoard == nil {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerAllOf) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil || o.ComputeBoard == nil {
		return nil, false
	}
	return o.ComputeBoard, true
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerAllOf) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard != nil {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given ComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *StorageFlexFlashControllerAllOf) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard = &v
}

// GetFlexFlashControllerProps returns the FlexFlashControllerProps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexFlashControllerAllOf) GetFlexFlashControllerProps() []StorageFlexFlashControllerPropsRelationship {
	if o == nil {
		var ret []StorageFlexFlashControllerPropsRelationship
		return ret
	}
	return o.FlexFlashControllerProps
}

// GetFlexFlashControllerPropsOk returns a tuple with the FlexFlashControllerProps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexFlashControllerAllOf) GetFlexFlashControllerPropsOk() (*[]StorageFlexFlashControllerPropsRelationship, bool) {
	if o == nil || o.FlexFlashControllerProps == nil {
		return nil, false
	}
	return &o.FlexFlashControllerProps, true
}

// HasFlexFlashControllerProps returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerAllOf) HasFlexFlashControllerProps() bool {
	if o != nil && o.FlexFlashControllerProps != nil {
		return true
	}

	return false
}

// SetFlexFlashControllerProps gets a reference to the given []StorageFlexFlashControllerPropsRelationship and assigns it to the FlexFlashControllerProps field.
func (o *StorageFlexFlashControllerAllOf) SetFlexFlashControllerProps(v []StorageFlexFlashControllerPropsRelationship) {
	o.FlexFlashControllerProps = v
}

// GetFlexFlashPhysicalDrives returns the FlexFlashPhysicalDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexFlashControllerAllOf) GetFlexFlashPhysicalDrives() []StorageFlexFlashPhysicalDriveRelationship {
	if o == nil {
		var ret []StorageFlexFlashPhysicalDriveRelationship
		return ret
	}
	return o.FlexFlashPhysicalDrives
}

// GetFlexFlashPhysicalDrivesOk returns a tuple with the FlexFlashPhysicalDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexFlashControllerAllOf) GetFlexFlashPhysicalDrivesOk() (*[]StorageFlexFlashPhysicalDriveRelationship, bool) {
	if o == nil || o.FlexFlashPhysicalDrives == nil {
		return nil, false
	}
	return &o.FlexFlashPhysicalDrives, true
}

// HasFlexFlashPhysicalDrives returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerAllOf) HasFlexFlashPhysicalDrives() bool {
	if o != nil && o.FlexFlashPhysicalDrives != nil {
		return true
	}

	return false
}

// SetFlexFlashPhysicalDrives gets a reference to the given []StorageFlexFlashPhysicalDriveRelationship and assigns it to the FlexFlashPhysicalDrives field.
func (o *StorageFlexFlashControllerAllOf) SetFlexFlashPhysicalDrives(v []StorageFlexFlashPhysicalDriveRelationship) {
	o.FlexFlashPhysicalDrives = v
}

// GetFlexFlashVirtualDrives returns the FlexFlashVirtualDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexFlashControllerAllOf) GetFlexFlashVirtualDrives() []StorageFlexFlashVirtualDriveRelationship {
	if o == nil {
		var ret []StorageFlexFlashVirtualDriveRelationship
		return ret
	}
	return o.FlexFlashVirtualDrives
}

// GetFlexFlashVirtualDrivesOk returns a tuple with the FlexFlashVirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexFlashControllerAllOf) GetFlexFlashVirtualDrivesOk() (*[]StorageFlexFlashVirtualDriveRelationship, bool) {
	if o == nil || o.FlexFlashVirtualDrives == nil {
		return nil, false
	}
	return &o.FlexFlashVirtualDrives, true
}

// HasFlexFlashVirtualDrives returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerAllOf) HasFlexFlashVirtualDrives() bool {
	if o != nil && o.FlexFlashVirtualDrives != nil {
		return true
	}

	return false
}

// SetFlexFlashVirtualDrives gets a reference to the given []StorageFlexFlashVirtualDriveRelationship and assigns it to the FlexFlashVirtualDrives field.
func (o *StorageFlexFlashControllerAllOf) SetFlexFlashVirtualDrives(v []StorageFlexFlashVirtualDriveRelationship) {
	o.FlexFlashVirtualDrives = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageFlexFlashControllerAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageFlexFlashControllerAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashControllerAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageFlexFlashControllerAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetRunningFirmware returns the RunningFirmware field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexFlashControllerAllOf) GetRunningFirmware() []FirmwareRunningFirmwareRelationship {
	if o == nil {
		var ret []FirmwareRunningFirmwareRelationship
		return ret
	}
	return o.RunningFirmware
}

// GetRunningFirmwareOk returns a tuple with the RunningFirmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexFlashControllerAllOf) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool) {
	if o == nil || o.RunningFirmware == nil {
		return nil, false
	}
	return &o.RunningFirmware, true
}

// HasRunningFirmware returns a boolean if a field has been set.
func (o *StorageFlexFlashControllerAllOf) HasRunningFirmware() bool {
	if o != nil && o.RunningFirmware != nil {
		return true
	}

	return false
}

// SetRunningFirmware gets a reference to the given []FirmwareRunningFirmwareRelationship and assigns it to the RunningFirmware field.
func (o *StorageFlexFlashControllerAllOf) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship) {
	o.RunningFirmware = v
}

func (o StorageFlexFlashControllerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ControllerState != nil {
		toSerialize["ControllerState"] = o.ControllerState
	}
	if o.FfControllerId != nil {
		toSerialize["FfControllerId"] = o.FfControllerId
	}
	if o.ComputeBoard != nil {
		toSerialize["ComputeBoard"] = o.ComputeBoard
	}
	if o.FlexFlashControllerProps != nil {
		toSerialize["FlexFlashControllerProps"] = o.FlexFlashControllerProps
	}
	if o.FlexFlashPhysicalDrives != nil {
		toSerialize["FlexFlashPhysicalDrives"] = o.FlexFlashPhysicalDrives
	}
	if o.FlexFlashVirtualDrives != nil {
		toSerialize["FlexFlashVirtualDrives"] = o.FlexFlashVirtualDrives
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.RunningFirmware != nil {
		toSerialize["RunningFirmware"] = o.RunningFirmware
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageFlexFlashControllerAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varStorageFlexFlashControllerAllOf := _StorageFlexFlashControllerAllOf{}

	if err = json.Unmarshal(bytes, &varStorageFlexFlashControllerAllOf); err == nil {
		*o = StorageFlexFlashControllerAllOf(varStorageFlexFlashControllerAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ControllerState")
		delete(additionalProperties, "FfControllerId")
		delete(additionalProperties, "ComputeBoard")
		delete(additionalProperties, "FlexFlashControllerProps")
		delete(additionalProperties, "FlexFlashPhysicalDrives")
		delete(additionalProperties, "FlexFlashVirtualDrives")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "RunningFirmware")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStorageFlexFlashControllerAllOf struct {
	value *StorageFlexFlashControllerAllOf
	isSet bool
}

func (v NullableStorageFlexFlashControllerAllOf) Get() *StorageFlexFlashControllerAllOf {
	return v.value
}

func (v *NullableStorageFlexFlashControllerAllOf) Set(val *StorageFlexFlashControllerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageFlexFlashControllerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageFlexFlashControllerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageFlexFlashControllerAllOf(val *StorageFlexFlashControllerAllOf) *NullableStorageFlexFlashControllerAllOf {
	return &NullableStorageFlexFlashControllerAllOf{value: val, isSet: true}
}

func (v NullableStorageFlexFlashControllerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageFlexFlashControllerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
