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

// StorageVirtualDriveConfig Virtual Drive type models a single virtual drive that needs to be created through this policy.
type StorageVirtualDriveConfig struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Access policy that host has on this virtual drive. * `Default` - Use platform default access mode. * `ReadWrite` - Enables host to perform read-write on the VD. * `ReadOnly` - Host can only read from the VD. * `Blocked` - Host can neither read nor write to the VD.
	AccessPolicy *string `json:"AccessPolicy,omitempty"`
	// The flag enables the use of this virtual drive as a boot drive.
	BootDrive *bool `json:"BootDrive,omitempty"`
	// Disk group policy that has the disk group in which this virtual drive needs to be created.
	DiskGroupName *string `json:"DiskGroupName,omitempty"`
	// Disk group policy that has the disk group in which this virtual drive needs to be created.
	DiskGroupPolicy *string `json:"DiskGroupPolicy,omitempty"`
	// Drive Cache property expect disk cache policy. * `Default` - Use platform default drive cache mode. * `NoChange` - Drive cache policy is unchanged. * `Enable` - Enables IO caching on the drive. * `Disable` - Disables IO caching on the drive.
	DriveCache *string `json:"DriveCache,omitempty"`
	// The flag enables this virtual drive to use all the available space in the disk group. When this flag is configured, the size property is ignored.
	ExpandToAvailable *bool `json:"ExpandToAvailable,omitempty"`
	// Desired IO mode - direct IO or cached IO. * `Default` - Use platform default IO mode. * `Direct` - Use direct IO for writing directly into the disk. * `Cached` - Use cached IO for writing into cache and then to disk.
	IoPolicy *string `json:"IoPolicy,omitempty"`
	// The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed.
	Name *string `json:"Name,omitempty"`
	// Read ahead mode to be used to read data from this virtual drive. * `Default` - Use platform default read ahead mode. * `ReadAhead` - Use read ahead mode for the policy. * `NoReadAhead` - Do not use read ahead mode for the policy.
	ReadPolicy *string `json:"ReadPolicy,omitempty"`
	// Virtual drive size in MB. Size is mandatory field unless the 'Expand to Available' option is enabled.
	Size *int64 `json:"Size,omitempty"`
	// The strip size is the portion of a stripe that resides on a single drive in the drive group. The stripe consists of the data segments that the RAID controller writes across multiple drives, not including parity drives. * `Default` - Use platform default strip size for a virtual drive. * `32k` - Enables a strip size of 32k for a virtual drive. * `64k` - Enables a strip size of 64k for a virtual drive. * `128k` - Enables a strip size of 128k for a virtual drive. * `256k` - Enables a strip size of 256k for a virtual drive. * `512k` - Enables a strip size of 512k for a virtual drive. * `1024k` - Enables a strip size of 1024k for a virtual drive.
	StripSize *string `json:"StripSize,omitempty"`
	// Unique Id of the Virtual Drive to be used to identify Virtual Drive when name is empty.
	Vdid *string `json:"Vdid,omitempty"`
	// Write mode to be used to write data to this virtual drive. * `Default` - Use platform default write mode. * `WriteThrough` - Data is written through the cache and to the physical drives. Performance is improved, because subsequent reads of that data can be satisfied from the cache. * `WriteBackGoodBbu` - Data is stored in the cache, and is only written to the physical drives when space in the cache is needed. Virtual drives requesting this policy fall back to Write Through caching when the battery backup unit (BBU) cannot guarantee the safety of the cache in the event of a power failure. * `AlwaysWriteBack` - With this policy, write caching remains Write Back even if the battery backup unit is defective or discharged.
	WritePolicy          *string `json:"WritePolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageVirtualDriveConfig StorageVirtualDriveConfig

// NewStorageVirtualDriveConfig instantiates a new StorageVirtualDriveConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageVirtualDriveConfig(classId string, objectType string) *StorageVirtualDriveConfig {
	this := StorageVirtualDriveConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	var accessPolicy string = "Default"
	this.AccessPolicy = &accessPolicy
	var driveCache string = "Default"
	this.DriveCache = &driveCache
	var ioPolicy string = "Default"
	this.IoPolicy = &ioPolicy
	var readPolicy string = "Default"
	this.ReadPolicy = &readPolicy
	var stripSize string = "Default"
	this.StripSize = &stripSize
	var writePolicy string = "Default"
	this.WritePolicy = &writePolicy
	return &this
}

// NewStorageVirtualDriveConfigWithDefaults instantiates a new StorageVirtualDriveConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageVirtualDriveConfigWithDefaults() *StorageVirtualDriveConfig {
	this := StorageVirtualDriveConfig{}
	var classId string = "storage.VirtualDriveConfig"
	this.ClassId = classId
	var objectType string = "storage.VirtualDriveConfig"
	this.ObjectType = objectType
	var accessPolicy string = "Default"
	this.AccessPolicy = &accessPolicy
	var driveCache string = "Default"
	this.DriveCache = &driveCache
	var ioPolicy string = "Default"
	this.IoPolicy = &ioPolicy
	var readPolicy string = "Default"
	this.ReadPolicy = &readPolicy
	var stripSize string = "Default"
	this.StripSize = &stripSize
	var writePolicy string = "Default"
	this.WritePolicy = &writePolicy
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageVirtualDriveConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageVirtualDriveConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageVirtualDriveConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageVirtualDriveConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAccessPolicy returns the AccessPolicy field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfig) GetAccessPolicy() string {
	if o == nil || o.AccessPolicy == nil {
		var ret string
		return ret
	}
	return *o.AccessPolicy
}

// GetAccessPolicyOk returns a tuple with the AccessPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfig) GetAccessPolicyOk() (*string, bool) {
	if o == nil || o.AccessPolicy == nil {
		return nil, false
	}
	return o.AccessPolicy, true
}

// HasAccessPolicy returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfig) HasAccessPolicy() bool {
	if o != nil && o.AccessPolicy != nil {
		return true
	}

	return false
}

// SetAccessPolicy gets a reference to the given string and assigns it to the AccessPolicy field.
func (o *StorageVirtualDriveConfig) SetAccessPolicy(v string) {
	o.AccessPolicy = &v
}

// GetBootDrive returns the BootDrive field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfig) GetBootDrive() bool {
	if o == nil || o.BootDrive == nil {
		var ret bool
		return ret
	}
	return *o.BootDrive
}

// GetBootDriveOk returns a tuple with the BootDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfig) GetBootDriveOk() (*bool, bool) {
	if o == nil || o.BootDrive == nil {
		return nil, false
	}
	return o.BootDrive, true
}

// HasBootDrive returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfig) HasBootDrive() bool {
	if o != nil && o.BootDrive != nil {
		return true
	}

	return false
}

// SetBootDrive gets a reference to the given bool and assigns it to the BootDrive field.
func (o *StorageVirtualDriveConfig) SetBootDrive(v bool) {
	o.BootDrive = &v
}

// GetDiskGroupName returns the DiskGroupName field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfig) GetDiskGroupName() string {
	if o == nil || o.DiskGroupName == nil {
		var ret string
		return ret
	}
	return *o.DiskGroupName
}

// GetDiskGroupNameOk returns a tuple with the DiskGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfig) GetDiskGroupNameOk() (*string, bool) {
	if o == nil || o.DiskGroupName == nil {
		return nil, false
	}
	return o.DiskGroupName, true
}

// HasDiskGroupName returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfig) HasDiskGroupName() bool {
	if o != nil && o.DiskGroupName != nil {
		return true
	}

	return false
}

// SetDiskGroupName gets a reference to the given string and assigns it to the DiskGroupName field.
func (o *StorageVirtualDriveConfig) SetDiskGroupName(v string) {
	o.DiskGroupName = &v
}

// GetDiskGroupPolicy returns the DiskGroupPolicy field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfig) GetDiskGroupPolicy() string {
	if o == nil || o.DiskGroupPolicy == nil {
		var ret string
		return ret
	}
	return *o.DiskGroupPolicy
}

// GetDiskGroupPolicyOk returns a tuple with the DiskGroupPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfig) GetDiskGroupPolicyOk() (*string, bool) {
	if o == nil || o.DiskGroupPolicy == nil {
		return nil, false
	}
	return o.DiskGroupPolicy, true
}

// HasDiskGroupPolicy returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfig) HasDiskGroupPolicy() bool {
	if o != nil && o.DiskGroupPolicy != nil {
		return true
	}

	return false
}

// SetDiskGroupPolicy gets a reference to the given string and assigns it to the DiskGroupPolicy field.
func (o *StorageVirtualDriveConfig) SetDiskGroupPolicy(v string) {
	o.DiskGroupPolicy = &v
}

// GetDriveCache returns the DriveCache field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfig) GetDriveCache() string {
	if o == nil || o.DriveCache == nil {
		var ret string
		return ret
	}
	return *o.DriveCache
}

// GetDriveCacheOk returns a tuple with the DriveCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfig) GetDriveCacheOk() (*string, bool) {
	if o == nil || o.DriveCache == nil {
		return nil, false
	}
	return o.DriveCache, true
}

// HasDriveCache returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfig) HasDriveCache() bool {
	if o != nil && o.DriveCache != nil {
		return true
	}

	return false
}

// SetDriveCache gets a reference to the given string and assigns it to the DriveCache field.
func (o *StorageVirtualDriveConfig) SetDriveCache(v string) {
	o.DriveCache = &v
}

// GetExpandToAvailable returns the ExpandToAvailable field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfig) GetExpandToAvailable() bool {
	if o == nil || o.ExpandToAvailable == nil {
		var ret bool
		return ret
	}
	return *o.ExpandToAvailable
}

// GetExpandToAvailableOk returns a tuple with the ExpandToAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfig) GetExpandToAvailableOk() (*bool, bool) {
	if o == nil || o.ExpandToAvailable == nil {
		return nil, false
	}
	return o.ExpandToAvailable, true
}

// HasExpandToAvailable returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfig) HasExpandToAvailable() bool {
	if o != nil && o.ExpandToAvailable != nil {
		return true
	}

	return false
}

// SetExpandToAvailable gets a reference to the given bool and assigns it to the ExpandToAvailable field.
func (o *StorageVirtualDriveConfig) SetExpandToAvailable(v bool) {
	o.ExpandToAvailable = &v
}

// GetIoPolicy returns the IoPolicy field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfig) GetIoPolicy() string {
	if o == nil || o.IoPolicy == nil {
		var ret string
		return ret
	}
	return *o.IoPolicy
}

// GetIoPolicyOk returns a tuple with the IoPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfig) GetIoPolicyOk() (*string, bool) {
	if o == nil || o.IoPolicy == nil {
		return nil, false
	}
	return o.IoPolicy, true
}

// HasIoPolicy returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfig) HasIoPolicy() bool {
	if o != nil && o.IoPolicy != nil {
		return true
	}

	return false
}

// SetIoPolicy gets a reference to the given string and assigns it to the IoPolicy field.
func (o *StorageVirtualDriveConfig) SetIoPolicy(v string) {
	o.IoPolicy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfig) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfig) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfig) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageVirtualDriveConfig) SetName(v string) {
	o.Name = &v
}

// GetReadPolicy returns the ReadPolicy field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfig) GetReadPolicy() string {
	if o == nil || o.ReadPolicy == nil {
		var ret string
		return ret
	}
	return *o.ReadPolicy
}

// GetReadPolicyOk returns a tuple with the ReadPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfig) GetReadPolicyOk() (*string, bool) {
	if o == nil || o.ReadPolicy == nil {
		return nil, false
	}
	return o.ReadPolicy, true
}

// HasReadPolicy returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfig) HasReadPolicy() bool {
	if o != nil && o.ReadPolicy != nil {
		return true
	}

	return false
}

// SetReadPolicy gets a reference to the given string and assigns it to the ReadPolicy field.
func (o *StorageVirtualDriveConfig) SetReadPolicy(v string) {
	o.ReadPolicy = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfig) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfig) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfig) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *StorageVirtualDriveConfig) SetSize(v int64) {
	o.Size = &v
}

// GetStripSize returns the StripSize field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfig) GetStripSize() string {
	if o == nil || o.StripSize == nil {
		var ret string
		return ret
	}
	return *o.StripSize
}

// GetStripSizeOk returns a tuple with the StripSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfig) GetStripSizeOk() (*string, bool) {
	if o == nil || o.StripSize == nil {
		return nil, false
	}
	return o.StripSize, true
}

// HasStripSize returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfig) HasStripSize() bool {
	if o != nil && o.StripSize != nil {
		return true
	}

	return false
}

// SetStripSize gets a reference to the given string and assigns it to the StripSize field.
func (o *StorageVirtualDriveConfig) SetStripSize(v string) {
	o.StripSize = &v
}

// GetVdid returns the Vdid field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfig) GetVdid() string {
	if o == nil || o.Vdid == nil {
		var ret string
		return ret
	}
	return *o.Vdid
}

// GetVdidOk returns a tuple with the Vdid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfig) GetVdidOk() (*string, bool) {
	if o == nil || o.Vdid == nil {
		return nil, false
	}
	return o.Vdid, true
}

// HasVdid returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfig) HasVdid() bool {
	if o != nil && o.Vdid != nil {
		return true
	}

	return false
}

// SetVdid gets a reference to the given string and assigns it to the Vdid field.
func (o *StorageVirtualDriveConfig) SetVdid(v string) {
	o.Vdid = &v
}

// GetWritePolicy returns the WritePolicy field value if set, zero value otherwise.
func (o *StorageVirtualDriveConfig) GetWritePolicy() string {
	if o == nil || o.WritePolicy == nil {
		var ret string
		return ret
	}
	return *o.WritePolicy
}

// GetWritePolicyOk returns a tuple with the WritePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageVirtualDriveConfig) GetWritePolicyOk() (*string, bool) {
	if o == nil || o.WritePolicy == nil {
		return nil, false
	}
	return o.WritePolicy, true
}

// HasWritePolicy returns a boolean if a field has been set.
func (o *StorageVirtualDriveConfig) HasWritePolicy() bool {
	if o != nil && o.WritePolicy != nil {
		return true
	}

	return false
}

// SetWritePolicy gets a reference to the given string and assigns it to the WritePolicy field.
func (o *StorageVirtualDriveConfig) SetWritePolicy(v string) {
	o.WritePolicy = &v
}

func (o StorageVirtualDriveConfig) MarshalJSON() ([]byte, error) {
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
	if o.AccessPolicy != nil {
		toSerialize["AccessPolicy"] = o.AccessPolicy
	}
	if o.BootDrive != nil {
		toSerialize["BootDrive"] = o.BootDrive
	}
	if o.DiskGroupName != nil {
		toSerialize["DiskGroupName"] = o.DiskGroupName
	}
	if o.DiskGroupPolicy != nil {
		toSerialize["DiskGroupPolicy"] = o.DiskGroupPolicy
	}
	if o.DriveCache != nil {
		toSerialize["DriveCache"] = o.DriveCache
	}
	if o.ExpandToAvailable != nil {
		toSerialize["ExpandToAvailable"] = o.ExpandToAvailable
	}
	if o.IoPolicy != nil {
		toSerialize["IoPolicy"] = o.IoPolicy
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ReadPolicy != nil {
		toSerialize["ReadPolicy"] = o.ReadPolicy
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.StripSize != nil {
		toSerialize["StripSize"] = o.StripSize
	}
	if o.Vdid != nil {
		toSerialize["Vdid"] = o.Vdid
	}
	if o.WritePolicy != nil {
		toSerialize["WritePolicy"] = o.WritePolicy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageVirtualDriveConfig) UnmarshalJSON(bytes []byte) (err error) {
	type StorageVirtualDriveConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Access policy that host has on this virtual drive. * `Default` - Use platform default access mode. * `ReadWrite` - Enables host to perform read-write on the VD. * `ReadOnly` - Host can only read from the VD. * `Blocked` - Host can neither read nor write to the VD.
		AccessPolicy *string `json:"AccessPolicy,omitempty"`
		// The flag enables the use of this virtual drive as a boot drive.
		BootDrive *bool `json:"BootDrive,omitempty"`
		// Disk group policy that has the disk group in which this virtual drive needs to be created.
		DiskGroupName *string `json:"DiskGroupName,omitempty"`
		// Disk group policy that has the disk group in which this virtual drive needs to be created.
		DiskGroupPolicy *string `json:"DiskGroupPolicy,omitempty"`
		// Drive Cache property expect disk cache policy. * `Default` - Use platform default drive cache mode. * `NoChange` - Drive cache policy is unchanged. * `Enable` - Enables IO caching on the drive. * `Disable` - Disables IO caching on the drive.
		DriveCache *string `json:"DriveCache,omitempty"`
		// The flag enables this virtual drive to use all the available space in the disk group. When this flag is configured, the size property is ignored.
		ExpandToAvailable *bool `json:"ExpandToAvailable,omitempty"`
		// Desired IO mode - direct IO or cached IO. * `Default` - Use platform default IO mode. * `Direct` - Use direct IO for writing directly into the disk. * `Cached` - Use cached IO for writing into cache and then to disk.
		IoPolicy *string `json:"IoPolicy,omitempty"`
		// The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed.
		Name *string `json:"Name,omitempty"`
		// Read ahead mode to be used to read data from this virtual drive. * `Default` - Use platform default read ahead mode. * `ReadAhead` - Use read ahead mode for the policy. * `NoReadAhead` - Do not use read ahead mode for the policy.
		ReadPolicy *string `json:"ReadPolicy,omitempty"`
		// Virtual drive size in MB. Size is mandatory field unless the 'Expand to Available' option is enabled.
		Size *int64 `json:"Size,omitempty"`
		// The strip size is the portion of a stripe that resides on a single drive in the drive group. The stripe consists of the data segments that the RAID controller writes across multiple drives, not including parity drives. * `Default` - Use platform default strip size for a virtual drive. * `32k` - Enables a strip size of 32k for a virtual drive. * `64k` - Enables a strip size of 64k for a virtual drive. * `128k` - Enables a strip size of 128k for a virtual drive. * `256k` - Enables a strip size of 256k for a virtual drive. * `512k` - Enables a strip size of 512k for a virtual drive. * `1024k` - Enables a strip size of 1024k for a virtual drive.
		StripSize *string `json:"StripSize,omitempty"`
		// Unique Id of the Virtual Drive to be used to identify Virtual Drive when name is empty.
		Vdid *string `json:"Vdid,omitempty"`
		// Write mode to be used to write data to this virtual drive. * `Default` - Use platform default write mode. * `WriteThrough` - Data is written through the cache and to the physical drives. Performance is improved, because subsequent reads of that data can be satisfied from the cache. * `WriteBackGoodBbu` - Data is stored in the cache, and is only written to the physical drives when space in the cache is needed. Virtual drives requesting this policy fall back to Write Through caching when the battery backup unit (BBU) cannot guarantee the safety of the cache in the event of a power failure. * `AlwaysWriteBack` - With this policy, write caching remains Write Back even if the battery backup unit is defective or discharged.
		WritePolicy *string `json:"WritePolicy,omitempty"`
	}

	varStorageVirtualDriveConfigWithoutEmbeddedStruct := StorageVirtualDriveConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageVirtualDriveConfigWithoutEmbeddedStruct)
	if err == nil {
		varStorageVirtualDriveConfig := _StorageVirtualDriveConfig{}
		varStorageVirtualDriveConfig.ClassId = varStorageVirtualDriveConfigWithoutEmbeddedStruct.ClassId
		varStorageVirtualDriveConfig.ObjectType = varStorageVirtualDriveConfigWithoutEmbeddedStruct.ObjectType
		varStorageVirtualDriveConfig.AccessPolicy = varStorageVirtualDriveConfigWithoutEmbeddedStruct.AccessPolicy
		varStorageVirtualDriveConfig.BootDrive = varStorageVirtualDriveConfigWithoutEmbeddedStruct.BootDrive
		varStorageVirtualDriveConfig.DiskGroupName = varStorageVirtualDriveConfigWithoutEmbeddedStruct.DiskGroupName
		varStorageVirtualDriveConfig.DiskGroupPolicy = varStorageVirtualDriveConfigWithoutEmbeddedStruct.DiskGroupPolicy
		varStorageVirtualDriveConfig.DriveCache = varStorageVirtualDriveConfigWithoutEmbeddedStruct.DriveCache
		varStorageVirtualDriveConfig.ExpandToAvailable = varStorageVirtualDriveConfigWithoutEmbeddedStruct.ExpandToAvailable
		varStorageVirtualDriveConfig.IoPolicy = varStorageVirtualDriveConfigWithoutEmbeddedStruct.IoPolicy
		varStorageVirtualDriveConfig.Name = varStorageVirtualDriveConfigWithoutEmbeddedStruct.Name
		varStorageVirtualDriveConfig.ReadPolicy = varStorageVirtualDriveConfigWithoutEmbeddedStruct.ReadPolicy
		varStorageVirtualDriveConfig.Size = varStorageVirtualDriveConfigWithoutEmbeddedStruct.Size
		varStorageVirtualDriveConfig.StripSize = varStorageVirtualDriveConfigWithoutEmbeddedStruct.StripSize
		varStorageVirtualDriveConfig.Vdid = varStorageVirtualDriveConfigWithoutEmbeddedStruct.Vdid
		varStorageVirtualDriveConfig.WritePolicy = varStorageVirtualDriveConfigWithoutEmbeddedStruct.WritePolicy
		*o = StorageVirtualDriveConfig(varStorageVirtualDriveConfig)
	} else {
		return err
	}

	varStorageVirtualDriveConfig := _StorageVirtualDriveConfig{}

	err = json.Unmarshal(bytes, &varStorageVirtualDriveConfig)
	if err == nil {
		o.MoBaseComplexType = varStorageVirtualDriveConfig.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AccessPolicy")
		delete(additionalProperties, "BootDrive")
		delete(additionalProperties, "DiskGroupName")
		delete(additionalProperties, "DiskGroupPolicy")
		delete(additionalProperties, "DriveCache")
		delete(additionalProperties, "ExpandToAvailable")
		delete(additionalProperties, "IoPolicy")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "ReadPolicy")
		delete(additionalProperties, "Size")
		delete(additionalProperties, "StripSize")
		delete(additionalProperties, "Vdid")
		delete(additionalProperties, "WritePolicy")

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

type NullableStorageVirtualDriveConfig struct {
	value *StorageVirtualDriveConfig
	isSet bool
}

func (v NullableStorageVirtualDriveConfig) Get() *StorageVirtualDriveConfig {
	return v.value
}

func (v *NullableStorageVirtualDriveConfig) Set(val *StorageVirtualDriveConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageVirtualDriveConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageVirtualDriveConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageVirtualDriveConfig(val *StorageVirtualDriveConfig) *NullableStorageVirtualDriveConfig {
	return &NullableStorageVirtualDriveConfig{value: val, isSet: true}
}

func (v NullableStorageVirtualDriveConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageVirtualDriveConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
