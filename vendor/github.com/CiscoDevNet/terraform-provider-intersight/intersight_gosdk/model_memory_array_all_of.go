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

// MemoryArrayAllOf Definition of the list of properties defined in 'memory.Array', excluding properties defined in parent classes.
type MemoryArrayAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The instance number of the memory array.
	ArrayId *int64 `json:"ArrayId,omitempty"`
	// ID of the CPU that access this memory array.
	CpuId *int64 `json:"CpuId,omitempty"`
	// Current capacity of all the memory units on a server.
	CurrentCapacity *string `json:"CurrentCapacity,omitempty"`
	// The primary hardware error detection or correction method supported by the memory array.
	ErrorCorrection *string `json:"ErrorCorrection,omitempty"`
	// Maximum capacity of all the memory units on a server.
	MaxCapacity *string `json:"MaxCapacity,omitempty"`
	// The maximum number of slots or sockets available for memory devices in the memory array.
	MaxDevices *string `json:"MaxDevices,omitempty"`
	// The power state indicator of the memory array.
	OperPowerState *string `json:"OperPowerState,omitempty"`
	// The presence of atleast one memory device in the array. Valid values are 'equipped' and 'absent'.
	Presence            *string                          `json:"Presence,omitempty"`
	ComputeBlade        *ComputeBladeRelationship        `json:"ComputeBlade,omitempty"`
	ComputeBoard        *ComputeBoardRelationship        `json:"ComputeBoard,omitempty"`
	ComputeRackUnit     *ComputeRackUnitRelationship     `json:"ComputeRackUnit,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	// An array of relationships to memoryPersistentMemoryUnit resources.
	PersistentMemoryUnits []MemoryPersistentMemoryUnitRelationship `json:"PersistentMemoryUnits,omitempty"`
	RegisteredDevice      *AssetDeviceRegistrationRelationship     `json:"RegisteredDevice,omitempty"`
	// An array of relationships to memoryUnit resources.
	Units                []MemoryUnitRelationship `json:"Units,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemoryArrayAllOf MemoryArrayAllOf

// NewMemoryArrayAllOf instantiates a new MemoryArrayAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryArrayAllOf(classId string, objectType string) *MemoryArrayAllOf {
	this := MemoryArrayAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewMemoryArrayAllOfWithDefaults instantiates a new MemoryArrayAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryArrayAllOfWithDefaults() *MemoryArrayAllOf {
	this := MemoryArrayAllOf{}
	var classId string = "memory.Array"
	this.ClassId = classId
	var objectType string = "memory.Array"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MemoryArrayAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MemoryArrayAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MemoryArrayAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MemoryArrayAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MemoryArrayAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MemoryArrayAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetArrayId returns the ArrayId field value if set, zero value otherwise.
func (o *MemoryArrayAllOf) GetArrayId() int64 {
	if o == nil || o.ArrayId == nil {
		var ret int64
		return ret
	}
	return *o.ArrayId
}

// GetArrayIdOk returns a tuple with the ArrayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArrayAllOf) GetArrayIdOk() (*int64, bool) {
	if o == nil || o.ArrayId == nil {
		return nil, false
	}
	return o.ArrayId, true
}

// HasArrayId returns a boolean if a field has been set.
func (o *MemoryArrayAllOf) HasArrayId() bool {
	if o != nil && o.ArrayId != nil {
		return true
	}

	return false
}

// SetArrayId gets a reference to the given int64 and assigns it to the ArrayId field.
func (o *MemoryArrayAllOf) SetArrayId(v int64) {
	o.ArrayId = &v
}

// GetCpuId returns the CpuId field value if set, zero value otherwise.
func (o *MemoryArrayAllOf) GetCpuId() int64 {
	if o == nil || o.CpuId == nil {
		var ret int64
		return ret
	}
	return *o.CpuId
}

// GetCpuIdOk returns a tuple with the CpuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArrayAllOf) GetCpuIdOk() (*int64, bool) {
	if o == nil || o.CpuId == nil {
		return nil, false
	}
	return o.CpuId, true
}

// HasCpuId returns a boolean if a field has been set.
func (o *MemoryArrayAllOf) HasCpuId() bool {
	if o != nil && o.CpuId != nil {
		return true
	}

	return false
}

// SetCpuId gets a reference to the given int64 and assigns it to the CpuId field.
func (o *MemoryArrayAllOf) SetCpuId(v int64) {
	o.CpuId = &v
}

// GetCurrentCapacity returns the CurrentCapacity field value if set, zero value otherwise.
func (o *MemoryArrayAllOf) GetCurrentCapacity() string {
	if o == nil || o.CurrentCapacity == nil {
		var ret string
		return ret
	}
	return *o.CurrentCapacity
}

// GetCurrentCapacityOk returns a tuple with the CurrentCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArrayAllOf) GetCurrentCapacityOk() (*string, bool) {
	if o == nil || o.CurrentCapacity == nil {
		return nil, false
	}
	return o.CurrentCapacity, true
}

// HasCurrentCapacity returns a boolean if a field has been set.
func (o *MemoryArrayAllOf) HasCurrentCapacity() bool {
	if o != nil && o.CurrentCapacity != nil {
		return true
	}

	return false
}

// SetCurrentCapacity gets a reference to the given string and assigns it to the CurrentCapacity field.
func (o *MemoryArrayAllOf) SetCurrentCapacity(v string) {
	o.CurrentCapacity = &v
}

// GetErrorCorrection returns the ErrorCorrection field value if set, zero value otherwise.
func (o *MemoryArrayAllOf) GetErrorCorrection() string {
	if o == nil || o.ErrorCorrection == nil {
		var ret string
		return ret
	}
	return *o.ErrorCorrection
}

// GetErrorCorrectionOk returns a tuple with the ErrorCorrection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArrayAllOf) GetErrorCorrectionOk() (*string, bool) {
	if o == nil || o.ErrorCorrection == nil {
		return nil, false
	}
	return o.ErrorCorrection, true
}

// HasErrorCorrection returns a boolean if a field has been set.
func (o *MemoryArrayAllOf) HasErrorCorrection() bool {
	if o != nil && o.ErrorCorrection != nil {
		return true
	}

	return false
}

// SetErrorCorrection gets a reference to the given string and assigns it to the ErrorCorrection field.
func (o *MemoryArrayAllOf) SetErrorCorrection(v string) {
	o.ErrorCorrection = &v
}

// GetMaxCapacity returns the MaxCapacity field value if set, zero value otherwise.
func (o *MemoryArrayAllOf) GetMaxCapacity() string {
	if o == nil || o.MaxCapacity == nil {
		var ret string
		return ret
	}
	return *o.MaxCapacity
}

// GetMaxCapacityOk returns a tuple with the MaxCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArrayAllOf) GetMaxCapacityOk() (*string, bool) {
	if o == nil || o.MaxCapacity == nil {
		return nil, false
	}
	return o.MaxCapacity, true
}

// HasMaxCapacity returns a boolean if a field has been set.
func (o *MemoryArrayAllOf) HasMaxCapacity() bool {
	if o != nil && o.MaxCapacity != nil {
		return true
	}

	return false
}

// SetMaxCapacity gets a reference to the given string and assigns it to the MaxCapacity field.
func (o *MemoryArrayAllOf) SetMaxCapacity(v string) {
	o.MaxCapacity = &v
}

// GetMaxDevices returns the MaxDevices field value if set, zero value otherwise.
func (o *MemoryArrayAllOf) GetMaxDevices() string {
	if o == nil || o.MaxDevices == nil {
		var ret string
		return ret
	}
	return *o.MaxDevices
}

// GetMaxDevicesOk returns a tuple with the MaxDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArrayAllOf) GetMaxDevicesOk() (*string, bool) {
	if o == nil || o.MaxDevices == nil {
		return nil, false
	}
	return o.MaxDevices, true
}

// HasMaxDevices returns a boolean if a field has been set.
func (o *MemoryArrayAllOf) HasMaxDevices() bool {
	if o != nil && o.MaxDevices != nil {
		return true
	}

	return false
}

// SetMaxDevices gets a reference to the given string and assigns it to the MaxDevices field.
func (o *MemoryArrayAllOf) SetMaxDevices(v string) {
	o.MaxDevices = &v
}

// GetOperPowerState returns the OperPowerState field value if set, zero value otherwise.
func (o *MemoryArrayAllOf) GetOperPowerState() string {
	if o == nil || o.OperPowerState == nil {
		var ret string
		return ret
	}
	return *o.OperPowerState
}

// GetOperPowerStateOk returns a tuple with the OperPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArrayAllOf) GetOperPowerStateOk() (*string, bool) {
	if o == nil || o.OperPowerState == nil {
		return nil, false
	}
	return o.OperPowerState, true
}

// HasOperPowerState returns a boolean if a field has been set.
func (o *MemoryArrayAllOf) HasOperPowerState() bool {
	if o != nil && o.OperPowerState != nil {
		return true
	}

	return false
}

// SetOperPowerState gets a reference to the given string and assigns it to the OperPowerState field.
func (o *MemoryArrayAllOf) SetOperPowerState(v string) {
	o.OperPowerState = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *MemoryArrayAllOf) GetPresence() string {
	if o == nil || o.Presence == nil {
		var ret string
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArrayAllOf) GetPresenceOk() (*string, bool) {
	if o == nil || o.Presence == nil {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *MemoryArrayAllOf) HasPresence() bool {
	if o != nil && o.Presence != nil {
		return true
	}

	return false
}

// SetPresence gets a reference to the given string and assigns it to the Presence field.
func (o *MemoryArrayAllOf) SetPresence(v string) {
	o.Presence = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *MemoryArrayAllOf) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArrayAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *MemoryArrayAllOf) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *MemoryArrayAllOf) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise.
func (o *MemoryArrayAllOf) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || o.ComputeBoard == nil {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArrayAllOf) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil || o.ComputeBoard == nil {
		return nil, false
	}
	return o.ComputeBoard, true
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *MemoryArrayAllOf) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard != nil {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given ComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *MemoryArrayAllOf) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *MemoryArrayAllOf) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArrayAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *MemoryArrayAllOf) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *MemoryArrayAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *MemoryArrayAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArrayAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *MemoryArrayAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *MemoryArrayAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetPersistentMemoryUnits returns the PersistentMemoryUnits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryArrayAllOf) GetPersistentMemoryUnits() []MemoryPersistentMemoryUnitRelationship {
	if o == nil {
		var ret []MemoryPersistentMemoryUnitRelationship
		return ret
	}
	return o.PersistentMemoryUnits
}

// GetPersistentMemoryUnitsOk returns a tuple with the PersistentMemoryUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryArrayAllOf) GetPersistentMemoryUnitsOk() (*[]MemoryPersistentMemoryUnitRelationship, bool) {
	if o == nil || o.PersistentMemoryUnits == nil {
		return nil, false
	}
	return &o.PersistentMemoryUnits, true
}

// HasPersistentMemoryUnits returns a boolean if a field has been set.
func (o *MemoryArrayAllOf) HasPersistentMemoryUnits() bool {
	if o != nil && o.PersistentMemoryUnits != nil {
		return true
	}

	return false
}

// SetPersistentMemoryUnits gets a reference to the given []MemoryPersistentMemoryUnitRelationship and assigns it to the PersistentMemoryUnits field.
func (o *MemoryArrayAllOf) SetPersistentMemoryUnits(v []MemoryPersistentMemoryUnitRelationship) {
	o.PersistentMemoryUnits = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *MemoryArrayAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryArrayAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *MemoryArrayAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *MemoryArrayAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetUnits returns the Units field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemoryArrayAllOf) GetUnits() []MemoryUnitRelationship {
	if o == nil {
		var ret []MemoryUnitRelationship
		return ret
	}
	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemoryArrayAllOf) GetUnitsOk() (*[]MemoryUnitRelationship, bool) {
	if o == nil || o.Units == nil {
		return nil, false
	}
	return &o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *MemoryArrayAllOf) HasUnits() bool {
	if o != nil && o.Units != nil {
		return true
	}

	return false
}

// SetUnits gets a reference to the given []MemoryUnitRelationship and assigns it to the Units field.
func (o *MemoryArrayAllOf) SetUnits(v []MemoryUnitRelationship) {
	o.Units = v
}

func (o MemoryArrayAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ArrayId != nil {
		toSerialize["ArrayId"] = o.ArrayId
	}
	if o.CpuId != nil {
		toSerialize["CpuId"] = o.CpuId
	}
	if o.CurrentCapacity != nil {
		toSerialize["CurrentCapacity"] = o.CurrentCapacity
	}
	if o.ErrorCorrection != nil {
		toSerialize["ErrorCorrection"] = o.ErrorCorrection
	}
	if o.MaxCapacity != nil {
		toSerialize["MaxCapacity"] = o.MaxCapacity
	}
	if o.MaxDevices != nil {
		toSerialize["MaxDevices"] = o.MaxDevices
	}
	if o.OperPowerState != nil {
		toSerialize["OperPowerState"] = o.OperPowerState
	}
	if o.Presence != nil {
		toSerialize["Presence"] = o.Presence
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeBoard != nil {
		toSerialize["ComputeBoard"] = o.ComputeBoard
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.PersistentMemoryUnits != nil {
		toSerialize["PersistentMemoryUnits"] = o.PersistentMemoryUnits
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.Units != nil {
		toSerialize["Units"] = o.Units
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MemoryArrayAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMemoryArrayAllOf := _MemoryArrayAllOf{}

	if err = json.Unmarshal(bytes, &varMemoryArrayAllOf); err == nil {
		*o = MemoryArrayAllOf(varMemoryArrayAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ArrayId")
		delete(additionalProperties, "CpuId")
		delete(additionalProperties, "CurrentCapacity")
		delete(additionalProperties, "ErrorCorrection")
		delete(additionalProperties, "MaxCapacity")
		delete(additionalProperties, "MaxDevices")
		delete(additionalProperties, "OperPowerState")
		delete(additionalProperties, "Presence")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "ComputeBoard")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PersistentMemoryUnits")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Units")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemoryArrayAllOf struct {
	value *MemoryArrayAllOf
	isSet bool
}

func (v NullableMemoryArrayAllOf) Get() *MemoryArrayAllOf {
	return v.value
}

func (v *NullableMemoryArrayAllOf) Set(val *MemoryArrayAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryArrayAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryArrayAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryArrayAllOf(val *MemoryArrayAllOf) *NullableMemoryArrayAllOf {
	return &NullableMemoryArrayAllOf{value: val, isSet: true}
}

func (v NullableMemoryArrayAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryArrayAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
