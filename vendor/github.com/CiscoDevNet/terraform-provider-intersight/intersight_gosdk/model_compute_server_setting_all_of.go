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
)

// ComputeServerSettingAllOf Definition of the list of properties defined in 'compute.ServerSetting', excluding properties defined in parent classes.
type ComputeServerSettingAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// User configured state of the locator LED for the server. * `None` - No operation property for locator led. * `On` - The Locator Led is turned on. * `Off` - The Locator Led is turned off.
	AdminLocatorLedState *string `json:"AdminLocatorLedState,omitempty"`
	// User configured power state of the server. * `Policy` - Power state is set to the default value in the policy. * `PowerOn` - Power state of the server is set to On. * `PowerOff` - Power state is the server set to Off. * `PowerCycle` - Power state the server is reset. * `HardReset` - Power state the server is hard reset. * `Shutdown` - Operating system on the server is shut down. * `Reboot` - Power state of IMC is rebooted.
	AdminPowerState *string `json:"AdminPowerState,omitempty"`
	// The configured state of these settings in the target server. The value is any one of Applied, Applying, Failed. Applied - This state denotes that the settings are applied successfully in the target server. Applying - This state denotes that the settings are being applied in the target server. Failed - This state denotes that the settings could not be applied in the target server. * `Applied` - User configured settings are in applied state. * `Applying` - User settings are being applied on the target server. * `Failed` - User configured settings could not be applied.
	ConfigState *string `json:"ConfigState,omitempty"`
	// The name of the device chosen by user for configuring One-Time Boot device.
	OneTimeBootDevice         *string                                  `json:"OneTimeBootDevice,omitempty"`
	PersistentMemoryOperation NullableComputePersistentMemoryOperation `json:"PersistentMemoryOperation,omitempty"`
	ServerConfig              NullableComputeServerConfig              `json:"ServerConfig,omitempty"`
	LocatorLed                *EquipmentLocatorLedRelationship         `json:"LocatorLed,omitempty"`
	RegisteredDevice          *AssetDeviceRegistrationRelationship     `json:"RegisteredDevice,omitempty"`
	RunningWorkflow           *WorkflowWorkflowInfoRelationship        `json:"RunningWorkflow,omitempty"`
	Server                    *ComputePhysicalRelationship             `json:"Server,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _ComputeServerSettingAllOf ComputeServerSettingAllOf

// NewComputeServerSettingAllOf instantiates a new ComputeServerSettingAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeServerSettingAllOf(classId string, objectType string) *ComputeServerSettingAllOf {
	this := ComputeServerSettingAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var adminLocatorLedState string = "None"
	this.AdminLocatorLedState = &adminLocatorLedState
	var adminPowerState string = "Policy"
	this.AdminPowerState = &adminPowerState
	var configState string = "Applied"
	this.ConfigState = &configState
	return &this
}

// NewComputeServerSettingAllOfWithDefaults instantiates a new ComputeServerSettingAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeServerSettingAllOfWithDefaults() *ComputeServerSettingAllOf {
	this := ComputeServerSettingAllOf{}
	var classId string = "compute.ServerSetting"
	this.ClassId = classId
	var objectType string = "compute.ServerSetting"
	this.ObjectType = objectType
	var adminLocatorLedState string = "None"
	this.AdminLocatorLedState = &adminLocatorLedState
	var adminPowerState string = "Policy"
	this.AdminPowerState = &adminPowerState
	var configState string = "Applied"
	this.ConfigState = &configState
	return &this
}

// GetClassId returns the ClassId field value
func (o *ComputeServerSettingAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ComputeServerSettingAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ComputeServerSettingAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ComputeServerSettingAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ComputeServerSettingAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ComputeServerSettingAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminLocatorLedState returns the AdminLocatorLedState field value if set, zero value otherwise.
func (o *ComputeServerSettingAllOf) GetAdminLocatorLedState() string {
	if o == nil || o.AdminLocatorLedState == nil {
		var ret string
		return ret
	}
	return *o.AdminLocatorLedState
}

// GetAdminLocatorLedStateOk returns a tuple with the AdminLocatorLedState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSettingAllOf) GetAdminLocatorLedStateOk() (*string, bool) {
	if o == nil || o.AdminLocatorLedState == nil {
		return nil, false
	}
	return o.AdminLocatorLedState, true
}

// HasAdminLocatorLedState returns a boolean if a field has been set.
func (o *ComputeServerSettingAllOf) HasAdminLocatorLedState() bool {
	if o != nil && o.AdminLocatorLedState != nil {
		return true
	}

	return false
}

// SetAdminLocatorLedState gets a reference to the given string and assigns it to the AdminLocatorLedState field.
func (o *ComputeServerSettingAllOf) SetAdminLocatorLedState(v string) {
	o.AdminLocatorLedState = &v
}

// GetAdminPowerState returns the AdminPowerState field value if set, zero value otherwise.
func (o *ComputeServerSettingAllOf) GetAdminPowerState() string {
	if o == nil || o.AdminPowerState == nil {
		var ret string
		return ret
	}
	return *o.AdminPowerState
}

// GetAdminPowerStateOk returns a tuple with the AdminPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSettingAllOf) GetAdminPowerStateOk() (*string, bool) {
	if o == nil || o.AdminPowerState == nil {
		return nil, false
	}
	return o.AdminPowerState, true
}

// HasAdminPowerState returns a boolean if a field has been set.
func (o *ComputeServerSettingAllOf) HasAdminPowerState() bool {
	if o != nil && o.AdminPowerState != nil {
		return true
	}

	return false
}

// SetAdminPowerState gets a reference to the given string and assigns it to the AdminPowerState field.
func (o *ComputeServerSettingAllOf) SetAdminPowerState(v string) {
	o.AdminPowerState = &v
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *ComputeServerSettingAllOf) GetConfigState() string {
	if o == nil || o.ConfigState == nil {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSettingAllOf) GetConfigStateOk() (*string, bool) {
	if o == nil || o.ConfigState == nil {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *ComputeServerSettingAllOf) HasConfigState() bool {
	if o != nil && o.ConfigState != nil {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *ComputeServerSettingAllOf) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetOneTimeBootDevice returns the OneTimeBootDevice field value if set, zero value otherwise.
func (o *ComputeServerSettingAllOf) GetOneTimeBootDevice() string {
	if o == nil || o.OneTimeBootDevice == nil {
		var ret string
		return ret
	}
	return *o.OneTimeBootDevice
}

// GetOneTimeBootDeviceOk returns a tuple with the OneTimeBootDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSettingAllOf) GetOneTimeBootDeviceOk() (*string, bool) {
	if o == nil || o.OneTimeBootDevice == nil {
		return nil, false
	}
	return o.OneTimeBootDevice, true
}

// HasOneTimeBootDevice returns a boolean if a field has been set.
func (o *ComputeServerSettingAllOf) HasOneTimeBootDevice() bool {
	if o != nil && o.OneTimeBootDevice != nil {
		return true
	}

	return false
}

// SetOneTimeBootDevice gets a reference to the given string and assigns it to the OneTimeBootDevice field.
func (o *ComputeServerSettingAllOf) SetOneTimeBootDevice(v string) {
	o.OneTimeBootDevice = &v
}

// GetPersistentMemoryOperation returns the PersistentMemoryOperation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeServerSettingAllOf) GetPersistentMemoryOperation() ComputePersistentMemoryOperation {
	if o == nil || o.PersistentMemoryOperation.Get() == nil {
		var ret ComputePersistentMemoryOperation
		return ret
	}
	return *o.PersistentMemoryOperation.Get()
}

// GetPersistentMemoryOperationOk returns a tuple with the PersistentMemoryOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeServerSettingAllOf) GetPersistentMemoryOperationOk() (*ComputePersistentMemoryOperation, bool) {
	if o == nil {
		return nil, false
	}
	return o.PersistentMemoryOperation.Get(), o.PersistentMemoryOperation.IsSet()
}

// HasPersistentMemoryOperation returns a boolean if a field has been set.
func (o *ComputeServerSettingAllOf) HasPersistentMemoryOperation() bool {
	if o != nil && o.PersistentMemoryOperation.IsSet() {
		return true
	}

	return false
}

// SetPersistentMemoryOperation gets a reference to the given NullableComputePersistentMemoryOperation and assigns it to the PersistentMemoryOperation field.
func (o *ComputeServerSettingAllOf) SetPersistentMemoryOperation(v ComputePersistentMemoryOperation) {
	o.PersistentMemoryOperation.Set(&v)
}

// SetPersistentMemoryOperationNil sets the value for PersistentMemoryOperation to be an explicit nil
func (o *ComputeServerSettingAllOf) SetPersistentMemoryOperationNil() {
	o.PersistentMemoryOperation.Set(nil)
}

// UnsetPersistentMemoryOperation ensures that no value is present for PersistentMemoryOperation, not even an explicit nil
func (o *ComputeServerSettingAllOf) UnsetPersistentMemoryOperation() {
	o.PersistentMemoryOperation.Unset()
}

// GetServerConfig returns the ServerConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeServerSettingAllOf) GetServerConfig() ComputeServerConfig {
	if o == nil || o.ServerConfig.Get() == nil {
		var ret ComputeServerConfig
		return ret
	}
	return *o.ServerConfig.Get()
}

// GetServerConfigOk returns a tuple with the ServerConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeServerSettingAllOf) GetServerConfigOk() (*ComputeServerConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerConfig.Get(), o.ServerConfig.IsSet()
}

// HasServerConfig returns a boolean if a field has been set.
func (o *ComputeServerSettingAllOf) HasServerConfig() bool {
	if o != nil && o.ServerConfig.IsSet() {
		return true
	}

	return false
}

// SetServerConfig gets a reference to the given NullableComputeServerConfig and assigns it to the ServerConfig field.
func (o *ComputeServerSettingAllOf) SetServerConfig(v ComputeServerConfig) {
	o.ServerConfig.Set(&v)
}

// SetServerConfigNil sets the value for ServerConfig to be an explicit nil
func (o *ComputeServerSettingAllOf) SetServerConfigNil() {
	o.ServerConfig.Set(nil)
}

// UnsetServerConfig ensures that no value is present for ServerConfig, not even an explicit nil
func (o *ComputeServerSettingAllOf) UnsetServerConfig() {
	o.ServerConfig.Unset()
}

// GetLocatorLed returns the LocatorLed field value if set, zero value otherwise.
func (o *ComputeServerSettingAllOf) GetLocatorLed() EquipmentLocatorLedRelationship {
	if o == nil || o.LocatorLed == nil {
		var ret EquipmentLocatorLedRelationship
		return ret
	}
	return *o.LocatorLed
}

// GetLocatorLedOk returns a tuple with the LocatorLed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSettingAllOf) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool) {
	if o == nil || o.LocatorLed == nil {
		return nil, false
	}
	return o.LocatorLed, true
}

// HasLocatorLed returns a boolean if a field has been set.
func (o *ComputeServerSettingAllOf) HasLocatorLed() bool {
	if o != nil && o.LocatorLed != nil {
		return true
	}

	return false
}

// SetLocatorLed gets a reference to the given EquipmentLocatorLedRelationship and assigns it to the LocatorLed field.
func (o *ComputeServerSettingAllOf) SetLocatorLed(v EquipmentLocatorLedRelationship) {
	o.LocatorLed = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ComputeServerSettingAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSettingAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ComputeServerSettingAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ComputeServerSettingAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetRunningWorkflow returns the RunningWorkflow field value if set, zero value otherwise.
func (o *ComputeServerSettingAllOf) GetRunningWorkflow() WorkflowWorkflowInfoRelationship {
	if o == nil || o.RunningWorkflow == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.RunningWorkflow
}

// GetRunningWorkflowOk returns a tuple with the RunningWorkflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSettingAllOf) GetRunningWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.RunningWorkflow == nil {
		return nil, false
	}
	return o.RunningWorkflow, true
}

// HasRunningWorkflow returns a boolean if a field has been set.
func (o *ComputeServerSettingAllOf) HasRunningWorkflow() bool {
	if o != nil && o.RunningWorkflow != nil {
		return true
	}

	return false
}

// SetRunningWorkflow gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the RunningWorkflow field.
func (o *ComputeServerSettingAllOf) SetRunningWorkflow(v WorkflowWorkflowInfoRelationship) {
	o.RunningWorkflow = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *ComputeServerSettingAllOf) GetServer() ComputePhysicalRelationship {
	if o == nil || o.Server == nil {
		var ret ComputePhysicalRelationship
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeServerSettingAllOf) GetServerOk() (*ComputePhysicalRelationship, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *ComputeServerSettingAllOf) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given ComputePhysicalRelationship and assigns it to the Server field.
func (o *ComputeServerSettingAllOf) SetServer(v ComputePhysicalRelationship) {
	o.Server = &v
}

func (o ComputeServerSettingAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AdminLocatorLedState != nil {
		toSerialize["AdminLocatorLedState"] = o.AdminLocatorLedState
	}
	if o.AdminPowerState != nil {
		toSerialize["AdminPowerState"] = o.AdminPowerState
	}
	if o.ConfigState != nil {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if o.OneTimeBootDevice != nil {
		toSerialize["OneTimeBootDevice"] = o.OneTimeBootDevice
	}
	if o.PersistentMemoryOperation.IsSet() {
		toSerialize["PersistentMemoryOperation"] = o.PersistentMemoryOperation.Get()
	}
	if o.ServerConfig.IsSet() {
		toSerialize["ServerConfig"] = o.ServerConfig.Get()
	}
	if o.LocatorLed != nil {
		toSerialize["LocatorLed"] = o.LocatorLed
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.RunningWorkflow != nil {
		toSerialize["RunningWorkflow"] = o.RunningWorkflow
	}
	if o.Server != nil {
		toSerialize["Server"] = o.Server
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeServerSettingAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varComputeServerSettingAllOf := _ComputeServerSettingAllOf{}

	if err = json.Unmarshal(bytes, &varComputeServerSettingAllOf); err == nil {
		*o = ComputeServerSettingAllOf(varComputeServerSettingAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminLocatorLedState")
		delete(additionalProperties, "AdminPowerState")
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "OneTimeBootDevice")
		delete(additionalProperties, "PersistentMemoryOperation")
		delete(additionalProperties, "ServerConfig")
		delete(additionalProperties, "LocatorLed")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "RunningWorkflow")
		delete(additionalProperties, "Server")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComputeServerSettingAllOf struct {
	value *ComputeServerSettingAllOf
	isSet bool
}

func (v NullableComputeServerSettingAllOf) Get() *ComputeServerSettingAllOf {
	return v.value
}

func (v *NullableComputeServerSettingAllOf) Set(val *ComputeServerSettingAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeServerSettingAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeServerSettingAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeServerSettingAllOf(val *ComputeServerSettingAllOf) *NullableComputeServerSettingAllOf {
	return &NullableComputeServerSettingAllOf{value: val, isSet: true}
}

func (v NullableComputeServerSettingAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeServerSettingAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
