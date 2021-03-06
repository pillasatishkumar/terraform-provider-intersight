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

// NiatelemetryNiaInventoryDcnmAllOf Definition of the list of properties defined in 'niatelemetry.NiaInventoryDcnm', excluding properties defined in parent classes.
type NiatelemetryNiaInventoryDcnmAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Returns the value of the dev Field.
	Dev *bool `json:"Dev,omitempty"`
	// Returns the value of the haEnabled field.
	HaEnabled *bool `json:"HaEnabled,omitempty"`
	// Returns the value of the haReplicationStatus field.
	HaReplicationStatus *string `json:"HaReplicationStatus,omitempty"`
	// Returns the value of the install field.
	Install *string `json:"Install,omitempty"`
	// Returns the value of the isMediaController field.
	IsMediaController *bool `json:"IsMediaController,omitempty"`
	// Returns total number of fabrics in DCNM set-up.
	NumFabrics *int64 `json:"NumFabrics,omitempty"`
	// Returns the number of fabrics in msd.
	NumFabricsInMsd *int64 `json:"NumFabricsInMsd,omitempty"`
	// Returns the number of local users other than admin user.
	NumLocalUsers *int64 `json:"NumLocalUsers,omitempty"`
	// Returns the value of the version field.
	Version              *string                              `json:"Version,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NiatelemetryNiaInventoryDcnmAllOf NiatelemetryNiaInventoryDcnmAllOf

// NewNiatelemetryNiaInventoryDcnmAllOf instantiates a new NiatelemetryNiaInventoryDcnmAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNiaInventoryDcnmAllOf(classId string, objectType string) *NiatelemetryNiaInventoryDcnmAllOf {
	this := NiatelemetryNiaInventoryDcnmAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryNiaInventoryDcnmAllOfWithDefaults instantiates a new NiatelemetryNiaInventoryDcnmAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNiaInventoryDcnmAllOfWithDefaults() *NiatelemetryNiaInventoryDcnmAllOf {
	this := NiatelemetryNiaInventoryDcnmAllOf{}
	var classId string = "niatelemetry.NiaInventoryDcnm"
	this.ClassId = classId
	var objectType string = "niatelemetry.NiaInventoryDcnm"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDev returns the Dev field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetDev() bool {
	if o == nil || o.Dev == nil {
		var ret bool
		return ret
	}
	return *o.Dev
}

// GetDevOk returns a tuple with the Dev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetDevOk() (*bool, bool) {
	if o == nil || o.Dev == nil {
		return nil, false
	}
	return o.Dev, true
}

// HasDev returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasDev() bool {
	if o != nil && o.Dev != nil {
		return true
	}

	return false
}

// SetDev gets a reference to the given bool and assigns it to the Dev field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetDev(v bool) {
	o.Dev = &v
}

// GetHaEnabled returns the HaEnabled field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetHaEnabled() bool {
	if o == nil || o.HaEnabled == nil {
		var ret bool
		return ret
	}
	return *o.HaEnabled
}

// GetHaEnabledOk returns a tuple with the HaEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetHaEnabledOk() (*bool, bool) {
	if o == nil || o.HaEnabled == nil {
		return nil, false
	}
	return o.HaEnabled, true
}

// HasHaEnabled returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasHaEnabled() bool {
	if o != nil && o.HaEnabled != nil {
		return true
	}

	return false
}

// SetHaEnabled gets a reference to the given bool and assigns it to the HaEnabled field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetHaEnabled(v bool) {
	o.HaEnabled = &v
}

// GetHaReplicationStatus returns the HaReplicationStatus field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetHaReplicationStatus() string {
	if o == nil || o.HaReplicationStatus == nil {
		var ret string
		return ret
	}
	return *o.HaReplicationStatus
}

// GetHaReplicationStatusOk returns a tuple with the HaReplicationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetHaReplicationStatusOk() (*string, bool) {
	if o == nil || o.HaReplicationStatus == nil {
		return nil, false
	}
	return o.HaReplicationStatus, true
}

// HasHaReplicationStatus returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasHaReplicationStatus() bool {
	if o != nil && o.HaReplicationStatus != nil {
		return true
	}

	return false
}

// SetHaReplicationStatus gets a reference to the given string and assigns it to the HaReplicationStatus field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetHaReplicationStatus(v string) {
	o.HaReplicationStatus = &v
}

// GetInstall returns the Install field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetInstall() string {
	if o == nil || o.Install == nil {
		var ret string
		return ret
	}
	return *o.Install
}

// GetInstallOk returns a tuple with the Install field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetInstallOk() (*string, bool) {
	if o == nil || o.Install == nil {
		return nil, false
	}
	return o.Install, true
}

// HasInstall returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasInstall() bool {
	if o != nil && o.Install != nil {
		return true
	}

	return false
}

// SetInstall gets a reference to the given string and assigns it to the Install field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetInstall(v string) {
	o.Install = &v
}

// GetIsMediaController returns the IsMediaController field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetIsMediaController() bool {
	if o == nil || o.IsMediaController == nil {
		var ret bool
		return ret
	}
	return *o.IsMediaController
}

// GetIsMediaControllerOk returns a tuple with the IsMediaController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetIsMediaControllerOk() (*bool, bool) {
	if o == nil || o.IsMediaController == nil {
		return nil, false
	}
	return o.IsMediaController, true
}

// HasIsMediaController returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasIsMediaController() bool {
	if o != nil && o.IsMediaController != nil {
		return true
	}

	return false
}

// SetIsMediaController gets a reference to the given bool and assigns it to the IsMediaController field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetIsMediaController(v bool) {
	o.IsMediaController = &v
}

// GetNumFabrics returns the NumFabrics field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumFabrics() int64 {
	if o == nil || o.NumFabrics == nil {
		var ret int64
		return ret
	}
	return *o.NumFabrics
}

// GetNumFabricsOk returns a tuple with the NumFabrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumFabricsOk() (*int64, bool) {
	if o == nil || o.NumFabrics == nil {
		return nil, false
	}
	return o.NumFabrics, true
}

// HasNumFabrics returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNumFabrics() bool {
	if o != nil && o.NumFabrics != nil {
		return true
	}

	return false
}

// SetNumFabrics gets a reference to the given int64 and assigns it to the NumFabrics field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNumFabrics(v int64) {
	o.NumFabrics = &v
}

// GetNumFabricsInMsd returns the NumFabricsInMsd field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumFabricsInMsd() int64 {
	if o == nil || o.NumFabricsInMsd == nil {
		var ret int64
		return ret
	}
	return *o.NumFabricsInMsd
}

// GetNumFabricsInMsdOk returns a tuple with the NumFabricsInMsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumFabricsInMsdOk() (*int64, bool) {
	if o == nil || o.NumFabricsInMsd == nil {
		return nil, false
	}
	return o.NumFabricsInMsd, true
}

// HasNumFabricsInMsd returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNumFabricsInMsd() bool {
	if o != nil && o.NumFabricsInMsd != nil {
		return true
	}

	return false
}

// SetNumFabricsInMsd gets a reference to the given int64 and assigns it to the NumFabricsInMsd field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNumFabricsInMsd(v int64) {
	o.NumFabricsInMsd = &v
}

// GetNumLocalUsers returns the NumLocalUsers field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumLocalUsers() int64 {
	if o == nil || o.NumLocalUsers == nil {
		var ret int64
		return ret
	}
	return *o.NumLocalUsers
}

// GetNumLocalUsersOk returns a tuple with the NumLocalUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetNumLocalUsersOk() (*int64, bool) {
	if o == nil || o.NumLocalUsers == nil {
		return nil, false
	}
	return o.NumLocalUsers, true
}

// HasNumLocalUsers returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasNumLocalUsers() bool {
	if o != nil && o.NumLocalUsers != nil {
		return true
	}

	return false
}

// SetNumLocalUsers gets a reference to the given int64 and assigns it to the NumLocalUsers field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetNumLocalUsers(v int64) {
	o.NumLocalUsers = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryDcnmAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryNiaInventoryDcnmAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryNiaInventoryDcnmAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Dev != nil {
		toSerialize["Dev"] = o.Dev
	}
	if o.HaEnabled != nil {
		toSerialize["HaEnabled"] = o.HaEnabled
	}
	if o.HaReplicationStatus != nil {
		toSerialize["HaReplicationStatus"] = o.HaReplicationStatus
	}
	if o.Install != nil {
		toSerialize["Install"] = o.Install
	}
	if o.IsMediaController != nil {
		toSerialize["IsMediaController"] = o.IsMediaController
	}
	if o.NumFabrics != nil {
		toSerialize["NumFabrics"] = o.NumFabrics
	}
	if o.NumFabricsInMsd != nil {
		toSerialize["NumFabricsInMsd"] = o.NumFabricsInMsd
	}
	if o.NumLocalUsers != nil {
		toSerialize["NumLocalUsers"] = o.NumLocalUsers
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NiatelemetryNiaInventoryDcnmAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varNiatelemetryNiaInventoryDcnmAllOf := _NiatelemetryNiaInventoryDcnmAllOf{}

	if err = json.Unmarshal(bytes, &varNiatelemetryNiaInventoryDcnmAllOf); err == nil {
		*o = NiatelemetryNiaInventoryDcnmAllOf(varNiatelemetryNiaInventoryDcnmAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Dev")
		delete(additionalProperties, "HaEnabled")
		delete(additionalProperties, "HaReplicationStatus")
		delete(additionalProperties, "Install")
		delete(additionalProperties, "IsMediaController")
		delete(additionalProperties, "NumFabrics")
		delete(additionalProperties, "NumFabricsInMsd")
		delete(additionalProperties, "NumLocalUsers")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNiatelemetryNiaInventoryDcnmAllOf struct {
	value *NiatelemetryNiaInventoryDcnmAllOf
	isSet bool
}

func (v NullableNiatelemetryNiaInventoryDcnmAllOf) Get() *NiatelemetryNiaInventoryDcnmAllOf {
	return v.value
}

func (v *NullableNiatelemetryNiaInventoryDcnmAllOf) Set(val *NiatelemetryNiaInventoryDcnmAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNiaInventoryDcnmAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNiaInventoryDcnmAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNiaInventoryDcnmAllOf(val *NiatelemetryNiaInventoryDcnmAllOf) *NullableNiatelemetryNiaInventoryDcnmAllOf {
	return &NullableNiatelemetryNiaInventoryDcnmAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryNiaInventoryDcnmAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNiaInventoryDcnmAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
