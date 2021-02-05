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
	"reflect"
	"strings"
)

// StorageVirtualDriveExtension Information of virtual drives as reported by a storage controller. In certain cases like S-series servers, virtual drive information will be reported by the controller separately and this represents such information.
type StorageVirtualDriveExtension struct {
	InventoryBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The ability to boot from the virtual drive.
	Bootable *string `json:"Bootable,omitempty"`
	// The container id of the virtual drive.
	ContainerId *int64 `json:"ContainerId,omitempty"`
	// The state of the virtual drive.
	DriveState *string `json:"DriveState,omitempty"`
	// The name of the Virtual drive.
	Name *string `json:"Name,omitempty"`
	// The operational device id of the virtual drive.
	OperDeviceId *string `json:"OperDeviceId,omitempty"`
	// The UUID assigned to the virtual drive.
	Uuid *string `json:"Uuid,omitempty"`
	// The UUID value of the vendor assigned to the virtual drive.
	VendorUuid *string `json:"VendorUuid,omitempty"`
	// The distinguished name of the virtual drive for which the extended data is provided.
	VirtualDriveDn *string `json:"VirtualDriveDn,omitempty"`
	// The Id of the virtual drive.
	VirtualDriveId       *string                              `json:"VirtualDriveId,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	StorageController    *StorageControllerRelationship       `json:"StorageController,omitempty"`
	VirtualDrive         *StorageVirtualDriveRelationship     `json:"VirtualDrive,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageVirtualDriveExtension StorageVirtualDriveExtension

// NewStorageVirtualDriveExtension instantiates a new StorageVirtualDriveExtension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageVirtualDriveExtension(classId string, objectType string) *StorageVirtualDriveExtension {
	this := StorageVirtualDriveExtension{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageVirtualDriveExtensionWithDefaults instantiates a new StorageVirtualDriveExtension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageVirtualDriveExtensionWithDefaults() *StorageVirtualDriveExtension {
	this := StorageVirtualDriveExtension{}
	var classId string = "storage.VirtualDriveExtension"
	this.ClassId = classId
	var objectType string = "storage.VirtualDriveExtension"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageVirtualDriveExtension) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageVirtualDriveExtension) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageVirtualDriveExtension) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageVirtualDriveExtension) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootable returns the Bootable field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetBootable() string {
	if o == nil || o.Bootable == nil {
		var ret string
		return ret
	}
	return *o.Bootable
}

// GetBootableOk returns a tuple with the Bootable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetBootableOk() (*string, bool) {
	if o == nil || o.Bootable == nil {
		return nil, false
	}
	return o.Bootable, true
}

// HasBootable returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasBootable() bool {
	if o != nil && o.Bootable != nil {
		return true
	}

	return false
}

// SetBootable gets a reference to the given string and assigns it to the Bootable field.
func (o *StorageVirtualDriveExtension) SetBootable(v string) {
	o.Bootable = &v
}

// GetContainerId returns the ContainerId field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetContainerId() int64 {
	if o == nil || o.ContainerId == nil {
		var ret int64
		return ret
	}
	return *o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetContainerIdOk() (*int64, bool) {
	if o == nil || o.ContainerId == nil {
		return nil, false
	}
	return o.ContainerId, true
}

// HasContainerId returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasContainerId() bool {
	if o != nil && o.ContainerId != nil {
		return true
	}

	return false
}

// SetContainerId gets a reference to the given int64 and assigns it to the ContainerId field.
func (o *StorageVirtualDriveExtension) SetContainerId(v int64) {
	o.ContainerId = &v
}

// GetDriveState returns the DriveState field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetDriveState() string {
	if o == nil || o.DriveState == nil {
		var ret string
		return ret
	}
	return *o.DriveState
}

// GetDriveStateOk returns a tuple with the DriveState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetDriveStateOk() (*string, bool) {
	if o == nil || o.DriveState == nil {
		return nil, false
	}
	return o.DriveState, true
}

// HasDriveState returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasDriveState() bool {
	if o != nil && o.DriveState != nil {
		return true
	}

	return false
}

// SetDriveState gets a reference to the given string and assigns it to the DriveState field.
func (o *StorageVirtualDriveExtension) SetDriveState(v string) {
	o.DriveState = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageVirtualDriveExtension) SetName(v string) {
	o.Name = &v
}

// GetOperDeviceId returns the OperDeviceId field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetOperDeviceId() string {
	if o == nil || o.OperDeviceId == nil {
		var ret string
		return ret
	}
	return *o.OperDeviceId
}

// GetOperDeviceIdOk returns a tuple with the OperDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetOperDeviceIdOk() (*string, bool) {
	if o == nil || o.OperDeviceId == nil {
		return nil, false
	}
	return o.OperDeviceId, true
}

// HasOperDeviceId returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasOperDeviceId() bool {
	if o != nil && o.OperDeviceId != nil {
		return true
	}

	return false
}

// SetOperDeviceId gets a reference to the given string and assigns it to the OperDeviceId field.
func (o *StorageVirtualDriveExtension) SetOperDeviceId(v string) {
	o.OperDeviceId = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageVirtualDriveExtension) SetUuid(v string) {
	o.Uuid = &v
}

// GetVendorUuid returns the VendorUuid field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetVendorUuid() string {
	if o == nil || o.VendorUuid == nil {
		var ret string
		return ret
	}
	return *o.VendorUuid
}

// GetVendorUuidOk returns a tuple with the VendorUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetVendorUuidOk() (*string, bool) {
	if o == nil || o.VendorUuid == nil {
		return nil, false
	}
	return o.VendorUuid, true
}

// HasVendorUuid returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasVendorUuid() bool {
	if o != nil && o.VendorUuid != nil {
		return true
	}

	return false
}

// SetVendorUuid gets a reference to the given string and assigns it to the VendorUuid field.
func (o *StorageVirtualDriveExtension) SetVendorUuid(v string) {
	o.VendorUuid = &v
}

// GetVirtualDriveDn returns the VirtualDriveDn field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetVirtualDriveDn() string {
	if o == nil || o.VirtualDriveDn == nil {
		var ret string
		return ret
	}
	return *o.VirtualDriveDn
}

// GetVirtualDriveDnOk returns a tuple with the VirtualDriveDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetVirtualDriveDnOk() (*string, bool) {
	if o == nil || o.VirtualDriveDn == nil {
		return nil, false
	}
	return o.VirtualDriveDn, true
}

// HasVirtualDriveDn returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasVirtualDriveDn() bool {
	if o != nil && o.VirtualDriveDn != nil {
		return true
	}

	return false
}

// SetVirtualDriveDn gets a reference to the given string and assigns it to the VirtualDriveDn field.
func (o *StorageVirtualDriveExtension) SetVirtualDriveDn(v string) {
	o.VirtualDriveDn = &v
}

// GetVirtualDriveId returns the VirtualDriveId field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetVirtualDriveId() string {
	if o == nil || o.VirtualDriveId == nil {
		var ret string
		return ret
	}
	return *o.VirtualDriveId
}

// GetVirtualDriveIdOk returns a tuple with the VirtualDriveId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetVirtualDriveIdOk() (*string, bool) {
	if o == nil || o.VirtualDriveId == nil {
		return nil, false
	}
	return o.VirtualDriveId, true
}

// HasVirtualDriveId returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasVirtualDriveId() bool {
	if o != nil && o.VirtualDriveId != nil {
		return true
	}

	return false
}

// SetVirtualDriveId gets a reference to the given string and assigns it to the VirtualDriveId field.
func (o *StorageVirtualDriveExtension) SetVirtualDriveId(v string) {
	o.VirtualDriveId = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageVirtualDriveExtension) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageVirtualDriveExtension) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageController returns the StorageController field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetStorageController() StorageControllerRelationship {
	if o == nil || o.StorageController == nil {
		var ret StorageControllerRelationship
		return ret
	}
	return *o.StorageController
}

// GetStorageControllerOk returns a tuple with the StorageController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetStorageControllerOk() (*StorageControllerRelationship, bool) {
	if o == nil || o.StorageController == nil {
		return nil, false
	}
	return o.StorageController, true
}

// HasStorageController returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasStorageController() bool {
	if o != nil && o.StorageController != nil {
		return true
	}

	return false
}

// SetStorageController gets a reference to the given StorageControllerRelationship and assigns it to the StorageController field.
func (o *StorageVirtualDriveExtension) SetStorageController(v StorageControllerRelationship) {
	o.StorageController = &v
}

// GetVirtualDrive returns the VirtualDrive field value if set, zero value otherwise.
func (o *StorageVirtualDriveExtension) GetVirtualDrive() StorageVirtualDriveRelationship {
	if o == nil || o.VirtualDrive == nil {
		var ret StorageVirtualDriveRelationship
		return ret
	}
	return *o.VirtualDrive
}

// GetVirtualDriveOk returns a tuple with the VirtualDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveExtension) GetVirtualDriveOk() (*StorageVirtualDriveRelationship, bool) {
	if o == nil || o.VirtualDrive == nil {
		return nil, false
	}
	return o.VirtualDrive, true
}

// HasVirtualDrive returns a boolean if a field has been set.
func (o *StorageVirtualDriveExtension) HasVirtualDrive() bool {
	if o != nil && o.VirtualDrive != nil {
		return true
	}

	return false
}

// SetVirtualDrive gets a reference to the given StorageVirtualDriveRelationship and assigns it to the VirtualDrive field.
func (o *StorageVirtualDriveExtension) SetVirtualDrive(v StorageVirtualDriveRelationship) {
	o.VirtualDrive = &v
}

func (o StorageVirtualDriveExtension) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedInventoryBase, errInventoryBase := json.Marshal(o.InventoryBase)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	errInventoryBase = json.Unmarshal([]byte(serializedInventoryBase), &toSerialize)
	if errInventoryBase != nil {
		return []byte{}, errInventoryBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Bootable != nil {
		toSerialize["Bootable"] = o.Bootable
	}
	if o.ContainerId != nil {
		toSerialize["ContainerId"] = o.ContainerId
	}
	if o.DriveState != nil {
		toSerialize["DriveState"] = o.DriveState
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperDeviceId != nil {
		toSerialize["OperDeviceId"] = o.OperDeviceId
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.VendorUuid != nil {
		toSerialize["VendorUuid"] = o.VendorUuid
	}
	if o.VirtualDriveDn != nil {
		toSerialize["VirtualDriveDn"] = o.VirtualDriveDn
	}
	if o.VirtualDriveId != nil {
		toSerialize["VirtualDriveId"] = o.VirtualDriveId
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageController != nil {
		toSerialize["StorageController"] = o.StorageController
	}
	if o.VirtualDrive != nil {
		toSerialize["VirtualDrive"] = o.VirtualDrive
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageVirtualDriveExtension) UnmarshalJSON(bytes []byte) (err error) {
	type StorageVirtualDriveExtensionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The ability to boot from the virtual drive.
		Bootable *string `json:"Bootable,omitempty"`
		// The container id of the virtual drive.
		ContainerId *int64 `json:"ContainerId,omitempty"`
		// The state of the virtual drive.
		DriveState *string `json:"DriveState,omitempty"`
		// The name of the Virtual drive.
		Name *string `json:"Name,omitempty"`
		// The operational device id of the virtual drive.
		OperDeviceId *string `json:"OperDeviceId,omitempty"`
		// The UUID assigned to the virtual drive.
		Uuid *string `json:"Uuid,omitempty"`
		// The UUID value of the vendor assigned to the virtual drive.
		VendorUuid *string `json:"VendorUuid,omitempty"`
		// The distinguished name of the virtual drive for which the extended data is provided.
		VirtualDriveDn *string `json:"VirtualDriveDn,omitempty"`
		// The Id of the virtual drive.
		VirtualDriveId      *string                              `json:"VirtualDriveId,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		StorageController   *StorageControllerRelationship       `json:"StorageController,omitempty"`
		VirtualDrive        *StorageVirtualDriveRelationship     `json:"VirtualDrive,omitempty"`
	}

	varStorageVirtualDriveExtensionWithoutEmbeddedStruct := StorageVirtualDriveExtensionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageVirtualDriveExtensionWithoutEmbeddedStruct)
	if err == nil {
		varStorageVirtualDriveExtension := _StorageVirtualDriveExtension{}
		varStorageVirtualDriveExtension.ClassId = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.ClassId
		varStorageVirtualDriveExtension.ObjectType = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.ObjectType
		varStorageVirtualDriveExtension.Bootable = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.Bootable
		varStorageVirtualDriveExtension.ContainerId = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.ContainerId
		varStorageVirtualDriveExtension.DriveState = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.DriveState
		varStorageVirtualDriveExtension.Name = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.Name
		varStorageVirtualDriveExtension.OperDeviceId = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.OperDeviceId
		varStorageVirtualDriveExtension.Uuid = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.Uuid
		varStorageVirtualDriveExtension.VendorUuid = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.VendorUuid
		varStorageVirtualDriveExtension.VirtualDriveDn = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.VirtualDriveDn
		varStorageVirtualDriveExtension.VirtualDriveId = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.VirtualDriveId
		varStorageVirtualDriveExtension.InventoryDeviceInfo = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.InventoryDeviceInfo
		varStorageVirtualDriveExtension.RegisteredDevice = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.RegisteredDevice
		varStorageVirtualDriveExtension.StorageController = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.StorageController
		varStorageVirtualDriveExtension.VirtualDrive = varStorageVirtualDriveExtensionWithoutEmbeddedStruct.VirtualDrive
		*o = StorageVirtualDriveExtension(varStorageVirtualDriveExtension)
	} else {
		return err
	}

	varStorageVirtualDriveExtension := _StorageVirtualDriveExtension{}

	err = json.Unmarshal(bytes, &varStorageVirtualDriveExtension)
	if err == nil {
		o.InventoryBase = varStorageVirtualDriveExtension.InventoryBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Bootable")
		delete(additionalProperties, "ContainerId")
		delete(additionalProperties, "DriveState")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperDeviceId")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "VendorUuid")
		delete(additionalProperties, "VirtualDriveDn")
		delete(additionalProperties, "VirtualDriveId")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageController")
		delete(additionalProperties, "VirtualDrive")

		// remove fields from embedded structs
		reflectInventoryBase := reflect.ValueOf(o.InventoryBase)
		for i := 0; i < reflectInventoryBase.Type().NumField(); i++ {
			t := reflectInventoryBase.Type().Field(i)

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

type NullableStorageVirtualDriveExtension struct {
	value *StorageVirtualDriveExtension
	isSet bool
}

func (v NullableStorageVirtualDriveExtension) Get() *StorageVirtualDriveExtension {
	return v.value
}

func (v *NullableStorageVirtualDriveExtension) Set(val *StorageVirtualDriveExtension) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageVirtualDriveExtension) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageVirtualDriveExtension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageVirtualDriveExtension(val *StorageVirtualDriveExtension) *NullableStorageVirtualDriveExtension {
	return &NullableStorageVirtualDriveExtension{value: val, isSet: true}
}

func (v NullableStorageVirtualDriveExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageVirtualDriveExtension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
