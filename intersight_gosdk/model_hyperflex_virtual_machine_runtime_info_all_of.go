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

// HyperflexVirtualMachineRuntimeInfoAllOf Definition of the list of properties defined in 'hyperflex.VirtualMachineRuntimeInfo', excluding properties defined in parent classes.
type HyperflexVirtualMachineRuntimeInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// BiosUuid of the Protected Virtual Machine.
	BiosUuid *string `json:"BiosUuid,omitempty"`
	// Connection state of the VM.
	ConnectionState *string `json:"ConnectionState,omitempty"`
	// CPU Usage of Virtual Machine.
	CpuUsage *int64 `json:"CpuUsage,omitempty"`
	// Folder which VM belongs to.
	Folder *string `json:"Folder,omitempty"`
	// Guest operating system family, if known.
	GuestFamily *string `json:"GuestFamily,omitempty"`
	// Guest operating system full name, if known.
	GuestFullName *string `json:"GuestFullName,omitempty"`
	// Guest operating system identifier (short name), if known.
	GuestId *string `json:"GuestId,omitempty"`
	// VMware guest reset, powercycle, or boot action states.
	GuestState *string `json:"GuestState,omitempty"`
	// Hostname of Virtual Machine.
	HostName *string `json:"HostName,omitempty"`
	// InstanceUuid of the Protected Virtual Machine.
	InstanceUuid *string `json:"InstanceUuid,omitempty"`
	// CPU Memory in MB of Virtual Machine.
	MemoryMb *int64 `json:"MemoryMb,omitempty"`
	// Memory usage of Virtual Machine.
	MemoryUsage *int64 `json:"MemoryUsage,omitempty"`
	// Virtual Machine unique MOID.
	Moid *string `json:"Moid,omitempty"`
	// Name of the Virtual Machine.
	Name     *string  `json:"Name,omitempty"`
	Networks []string `json:"Networks,omitempty"`
	// Number of CPUs for the VM.
	NumCpu *int64 `json:"NumCpu,omitempty"`
	// Power state of the Virtual Machine.
	PowerState *string `json:"PowerState,omitempty"`
	// Provisioned Size of Virtual Machine.
	ProvisionedSize *int64 `json:"ProvisionedSize,omitempty"`
	// Resource pool which VM belongs to.
	ResourcePool *string `json:"ResourcePool,omitempty"`
	// Used Size of Virtual Machine.
	UsedSize *int64 `json:"UsedSize,omitempty"`
	// Version of the Virtual Machine.
	Version *string `json:"Version,omitempty"`
	// Vmx Path in VC datastore format.
	VmxPath              *string `json:"VmxPath,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexVirtualMachineRuntimeInfoAllOf HyperflexVirtualMachineRuntimeInfoAllOf

// NewHyperflexVirtualMachineRuntimeInfoAllOf instantiates a new HyperflexVirtualMachineRuntimeInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexVirtualMachineRuntimeInfoAllOf(classId string, objectType string) *HyperflexVirtualMachineRuntimeInfoAllOf {
	this := HyperflexVirtualMachineRuntimeInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewHyperflexVirtualMachineRuntimeInfoAllOfWithDefaults instantiates a new HyperflexVirtualMachineRuntimeInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexVirtualMachineRuntimeInfoAllOfWithDefaults() *HyperflexVirtualMachineRuntimeInfoAllOf {
	this := HyperflexVirtualMachineRuntimeInfoAllOf{}
	var classId string = "hyperflex.VirtualMachineRuntimeInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.VirtualMachineRuntimeInfo"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBiosUuid returns the BiosUuid field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetBiosUuid() string {
	if o == nil || o.BiosUuid == nil {
		var ret string
		return ret
	}
	return *o.BiosUuid
}

// GetBiosUuidOk returns a tuple with the BiosUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetBiosUuidOk() (*string, bool) {
	if o == nil || o.BiosUuid == nil {
		return nil, false
	}
	return o.BiosUuid, true
}

// HasBiosUuid returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasBiosUuid() bool {
	if o != nil && o.BiosUuid != nil {
		return true
	}

	return false
}

// SetBiosUuid gets a reference to the given string and assigns it to the BiosUuid field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetBiosUuid(v string) {
	o.BiosUuid = &v
}

// GetConnectionState returns the ConnectionState field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetConnectionState() string {
	if o == nil || o.ConnectionState == nil {
		var ret string
		return ret
	}
	return *o.ConnectionState
}

// GetConnectionStateOk returns a tuple with the ConnectionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetConnectionStateOk() (*string, bool) {
	if o == nil || o.ConnectionState == nil {
		return nil, false
	}
	return o.ConnectionState, true
}

// HasConnectionState returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasConnectionState() bool {
	if o != nil && o.ConnectionState != nil {
		return true
	}

	return false
}

// SetConnectionState gets a reference to the given string and assigns it to the ConnectionState field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetConnectionState(v string) {
	o.ConnectionState = &v
}

// GetCpuUsage returns the CpuUsage field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetCpuUsage() int64 {
	if o == nil || o.CpuUsage == nil {
		var ret int64
		return ret
	}
	return *o.CpuUsage
}

// GetCpuUsageOk returns a tuple with the CpuUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetCpuUsageOk() (*int64, bool) {
	if o == nil || o.CpuUsage == nil {
		return nil, false
	}
	return o.CpuUsage, true
}

// HasCpuUsage returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasCpuUsage() bool {
	if o != nil && o.CpuUsage != nil {
		return true
	}

	return false
}

// SetCpuUsage gets a reference to the given int64 and assigns it to the CpuUsage field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetCpuUsage(v int64) {
	o.CpuUsage = &v
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetFolder() string {
	if o == nil || o.Folder == nil {
		var ret string
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetFolderOk() (*string, bool) {
	if o == nil || o.Folder == nil {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasFolder() bool {
	if o != nil && o.Folder != nil {
		return true
	}

	return false
}

// SetFolder gets a reference to the given string and assigns it to the Folder field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetFolder(v string) {
	o.Folder = &v
}

// GetGuestFamily returns the GuestFamily field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetGuestFamily() string {
	if o == nil || o.GuestFamily == nil {
		var ret string
		return ret
	}
	return *o.GuestFamily
}

// GetGuestFamilyOk returns a tuple with the GuestFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetGuestFamilyOk() (*string, bool) {
	if o == nil || o.GuestFamily == nil {
		return nil, false
	}
	return o.GuestFamily, true
}

// HasGuestFamily returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasGuestFamily() bool {
	if o != nil && o.GuestFamily != nil {
		return true
	}

	return false
}

// SetGuestFamily gets a reference to the given string and assigns it to the GuestFamily field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetGuestFamily(v string) {
	o.GuestFamily = &v
}

// GetGuestFullName returns the GuestFullName field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetGuestFullName() string {
	if o == nil || o.GuestFullName == nil {
		var ret string
		return ret
	}
	return *o.GuestFullName
}

// GetGuestFullNameOk returns a tuple with the GuestFullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetGuestFullNameOk() (*string, bool) {
	if o == nil || o.GuestFullName == nil {
		return nil, false
	}
	return o.GuestFullName, true
}

// HasGuestFullName returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasGuestFullName() bool {
	if o != nil && o.GuestFullName != nil {
		return true
	}

	return false
}

// SetGuestFullName gets a reference to the given string and assigns it to the GuestFullName field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetGuestFullName(v string) {
	o.GuestFullName = &v
}

// GetGuestId returns the GuestId field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetGuestId() string {
	if o == nil || o.GuestId == nil {
		var ret string
		return ret
	}
	return *o.GuestId
}

// GetGuestIdOk returns a tuple with the GuestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetGuestIdOk() (*string, bool) {
	if o == nil || o.GuestId == nil {
		return nil, false
	}
	return o.GuestId, true
}

// HasGuestId returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasGuestId() bool {
	if o != nil && o.GuestId != nil {
		return true
	}

	return false
}

// SetGuestId gets a reference to the given string and assigns it to the GuestId field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetGuestId(v string) {
	o.GuestId = &v
}

// GetGuestState returns the GuestState field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetGuestState() string {
	if o == nil || o.GuestState == nil {
		var ret string
		return ret
	}
	return *o.GuestState
}

// GetGuestStateOk returns a tuple with the GuestState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetGuestStateOk() (*string, bool) {
	if o == nil || o.GuestState == nil {
		return nil, false
	}
	return o.GuestState, true
}

// HasGuestState returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasGuestState() bool {
	if o != nil && o.GuestState != nil {
		return true
	}

	return false
}

// SetGuestState gets a reference to the given string and assigns it to the GuestState field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetGuestState(v string) {
	o.GuestState = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetHostName(v string) {
	o.HostName = &v
}

// GetInstanceUuid returns the InstanceUuid field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetInstanceUuid() string {
	if o == nil || o.InstanceUuid == nil {
		var ret string
		return ret
	}
	return *o.InstanceUuid
}

// GetInstanceUuidOk returns a tuple with the InstanceUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetInstanceUuidOk() (*string, bool) {
	if o == nil || o.InstanceUuid == nil {
		return nil, false
	}
	return o.InstanceUuid, true
}

// HasInstanceUuid returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasInstanceUuid() bool {
	if o != nil && o.InstanceUuid != nil {
		return true
	}

	return false
}

// SetInstanceUuid gets a reference to the given string and assigns it to the InstanceUuid field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetInstanceUuid(v string) {
	o.InstanceUuid = &v
}

// GetMemoryMb returns the MemoryMb field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetMemoryMb() int64 {
	if o == nil || o.MemoryMb == nil {
		var ret int64
		return ret
	}
	return *o.MemoryMb
}

// GetMemoryMbOk returns a tuple with the MemoryMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetMemoryMbOk() (*int64, bool) {
	if o == nil || o.MemoryMb == nil {
		return nil, false
	}
	return o.MemoryMb, true
}

// HasMemoryMb returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasMemoryMb() bool {
	if o != nil && o.MemoryMb != nil {
		return true
	}

	return false
}

// SetMemoryMb gets a reference to the given int64 and assigns it to the MemoryMb field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetMemoryMb(v int64) {
	o.MemoryMb = &v
}

// GetMemoryUsage returns the MemoryUsage field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetMemoryUsage() int64 {
	if o == nil || o.MemoryUsage == nil {
		var ret int64
		return ret
	}
	return *o.MemoryUsage
}

// GetMemoryUsageOk returns a tuple with the MemoryUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetMemoryUsageOk() (*int64, bool) {
	if o == nil || o.MemoryUsage == nil {
		return nil, false
	}
	return o.MemoryUsage, true
}

// HasMemoryUsage returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasMemoryUsage() bool {
	if o != nil && o.MemoryUsage != nil {
		return true
	}

	return false
}

// SetMemoryUsage gets a reference to the given int64 and assigns it to the MemoryUsage field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetMemoryUsage(v int64) {
	o.MemoryUsage = &v
}

// GetMoid returns the Moid field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetMoid() string {
	if o == nil || o.Moid == nil {
		var ret string
		return ret
	}
	return *o.Moid
}

// GetMoidOk returns a tuple with the Moid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetMoidOk() (*string, bool) {
	if o == nil || o.Moid == nil {
		return nil, false
	}
	return o.Moid, true
}

// HasMoid returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasMoid() bool {
	if o != nil && o.Moid != nil {
		return true
	}

	return false
}

// SetMoid gets a reference to the given string and assigns it to the Moid field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetMoid(v string) {
	o.Moid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetName(v string) {
	o.Name = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetNetworks() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetNetworksOk() (*[]string, bool) {
	if o == nil || o.Networks == nil {
		return nil, false
	}
	return &o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasNetworks() bool {
	if o != nil && o.Networks != nil {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []string and assigns it to the Networks field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetNetworks(v []string) {
	o.Networks = v
}

// GetNumCpu returns the NumCpu field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetNumCpu() int64 {
	if o == nil || o.NumCpu == nil {
		var ret int64
		return ret
	}
	return *o.NumCpu
}

// GetNumCpuOk returns a tuple with the NumCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetNumCpuOk() (*int64, bool) {
	if o == nil || o.NumCpu == nil {
		return nil, false
	}
	return o.NumCpu, true
}

// HasNumCpu returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasNumCpu() bool {
	if o != nil && o.NumCpu != nil {
		return true
	}

	return false
}

// SetNumCpu gets a reference to the given int64 and assigns it to the NumCpu field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetNumCpu(v int64) {
	o.NumCpu = &v
}

// GetPowerState returns the PowerState field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetPowerState() string {
	if o == nil || o.PowerState == nil {
		var ret string
		return ret
	}
	return *o.PowerState
}

// GetPowerStateOk returns a tuple with the PowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetPowerStateOk() (*string, bool) {
	if o == nil || o.PowerState == nil {
		return nil, false
	}
	return o.PowerState, true
}

// HasPowerState returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasPowerState() bool {
	if o != nil && o.PowerState != nil {
		return true
	}

	return false
}

// SetPowerState gets a reference to the given string and assigns it to the PowerState field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetPowerState(v string) {
	o.PowerState = &v
}

// GetProvisionedSize returns the ProvisionedSize field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetProvisionedSize() int64 {
	if o == nil || o.ProvisionedSize == nil {
		var ret int64
		return ret
	}
	return *o.ProvisionedSize
}

// GetProvisionedSizeOk returns a tuple with the ProvisionedSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetProvisionedSizeOk() (*int64, bool) {
	if o == nil || o.ProvisionedSize == nil {
		return nil, false
	}
	return o.ProvisionedSize, true
}

// HasProvisionedSize returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasProvisionedSize() bool {
	if o != nil && o.ProvisionedSize != nil {
		return true
	}

	return false
}

// SetProvisionedSize gets a reference to the given int64 and assigns it to the ProvisionedSize field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetProvisionedSize(v int64) {
	o.ProvisionedSize = &v
}

// GetResourcePool returns the ResourcePool field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetResourcePool() string {
	if o == nil || o.ResourcePool == nil {
		var ret string
		return ret
	}
	return *o.ResourcePool
}

// GetResourcePoolOk returns a tuple with the ResourcePool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetResourcePoolOk() (*string, bool) {
	if o == nil || o.ResourcePool == nil {
		return nil, false
	}
	return o.ResourcePool, true
}

// HasResourcePool returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasResourcePool() bool {
	if o != nil && o.ResourcePool != nil {
		return true
	}

	return false
}

// SetResourcePool gets a reference to the given string and assigns it to the ResourcePool field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetResourcePool(v string) {
	o.ResourcePool = &v
}

// GetUsedSize returns the UsedSize field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetUsedSize() int64 {
	if o == nil || o.UsedSize == nil {
		var ret int64
		return ret
	}
	return *o.UsedSize
}

// GetUsedSizeOk returns a tuple with the UsedSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetUsedSizeOk() (*int64, bool) {
	if o == nil || o.UsedSize == nil {
		return nil, false
	}
	return o.UsedSize, true
}

// HasUsedSize returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasUsedSize() bool {
	if o != nil && o.UsedSize != nil {
		return true
	}

	return false
}

// SetUsedSize gets a reference to the given int64 and assigns it to the UsedSize field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetUsedSize(v int64) {
	o.UsedSize = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetVmxPath returns the VmxPath field value if set, zero value otherwise.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetVmxPath() string {
	if o == nil || o.VmxPath == nil {
		var ret string
		return ret
	}
	return *o.VmxPath
}

// GetVmxPathOk returns a tuple with the VmxPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) GetVmxPathOk() (*string, bool) {
	if o == nil || o.VmxPath == nil {
		return nil, false
	}
	return o.VmxPath, true
}

// HasVmxPath returns a boolean if a field has been set.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) HasVmxPath() bool {
	if o != nil && o.VmxPath != nil {
		return true
	}

	return false
}

// SetVmxPath gets a reference to the given string and assigns it to the VmxPath field.
func (o *HyperflexVirtualMachineRuntimeInfoAllOf) SetVmxPath(v string) {
	o.VmxPath = &v
}

func (o HyperflexVirtualMachineRuntimeInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.BiosUuid != nil {
		toSerialize["BiosUuid"] = o.BiosUuid
	}
	if o.ConnectionState != nil {
		toSerialize["ConnectionState"] = o.ConnectionState
	}
	if o.CpuUsage != nil {
		toSerialize["CpuUsage"] = o.CpuUsage
	}
	if o.Folder != nil {
		toSerialize["Folder"] = o.Folder
	}
	if o.GuestFamily != nil {
		toSerialize["GuestFamily"] = o.GuestFamily
	}
	if o.GuestFullName != nil {
		toSerialize["GuestFullName"] = o.GuestFullName
	}
	if o.GuestId != nil {
		toSerialize["GuestId"] = o.GuestId
	}
	if o.GuestState != nil {
		toSerialize["GuestState"] = o.GuestState
	}
	if o.HostName != nil {
		toSerialize["HostName"] = o.HostName
	}
	if o.InstanceUuid != nil {
		toSerialize["InstanceUuid"] = o.InstanceUuid
	}
	if o.MemoryMb != nil {
		toSerialize["MemoryMb"] = o.MemoryMb
	}
	if o.MemoryUsage != nil {
		toSerialize["MemoryUsage"] = o.MemoryUsage
	}
	if o.Moid != nil {
		toSerialize["Moid"] = o.Moid
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Networks != nil {
		toSerialize["Networks"] = o.Networks
	}
	if o.NumCpu != nil {
		toSerialize["NumCpu"] = o.NumCpu
	}
	if o.PowerState != nil {
		toSerialize["PowerState"] = o.PowerState
	}
	if o.ProvisionedSize != nil {
		toSerialize["ProvisionedSize"] = o.ProvisionedSize
	}
	if o.ResourcePool != nil {
		toSerialize["ResourcePool"] = o.ResourcePool
	}
	if o.UsedSize != nil {
		toSerialize["UsedSize"] = o.UsedSize
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.VmxPath != nil {
		toSerialize["VmxPath"] = o.VmxPath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexVirtualMachineRuntimeInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexVirtualMachineRuntimeInfoAllOf := _HyperflexVirtualMachineRuntimeInfoAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexVirtualMachineRuntimeInfoAllOf); err == nil {
		*o = HyperflexVirtualMachineRuntimeInfoAllOf(varHyperflexVirtualMachineRuntimeInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "BiosUuid")
		delete(additionalProperties, "ConnectionState")
		delete(additionalProperties, "CpuUsage")
		delete(additionalProperties, "Folder")
		delete(additionalProperties, "GuestFamily")
		delete(additionalProperties, "GuestFullName")
		delete(additionalProperties, "GuestId")
		delete(additionalProperties, "GuestState")
		delete(additionalProperties, "HostName")
		delete(additionalProperties, "InstanceUuid")
		delete(additionalProperties, "MemoryMb")
		delete(additionalProperties, "MemoryUsage")
		delete(additionalProperties, "Moid")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Networks")
		delete(additionalProperties, "NumCpu")
		delete(additionalProperties, "PowerState")
		delete(additionalProperties, "ProvisionedSize")
		delete(additionalProperties, "ResourcePool")
		delete(additionalProperties, "UsedSize")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "VmxPath")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexVirtualMachineRuntimeInfoAllOf struct {
	value *HyperflexVirtualMachineRuntimeInfoAllOf
	isSet bool
}

func (v NullableHyperflexVirtualMachineRuntimeInfoAllOf) Get() *HyperflexVirtualMachineRuntimeInfoAllOf {
	return v.value
}

func (v *NullableHyperflexVirtualMachineRuntimeInfoAllOf) Set(val *HyperflexVirtualMachineRuntimeInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexVirtualMachineRuntimeInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexVirtualMachineRuntimeInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexVirtualMachineRuntimeInfoAllOf(val *HyperflexVirtualMachineRuntimeInfoAllOf) *NullableHyperflexVirtualMachineRuntimeInfoAllOf {
	return &NullableHyperflexVirtualMachineRuntimeInfoAllOf{value: val, isSet: true}
}

func (v NullableHyperflexVirtualMachineRuntimeInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexVirtualMachineRuntimeInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
