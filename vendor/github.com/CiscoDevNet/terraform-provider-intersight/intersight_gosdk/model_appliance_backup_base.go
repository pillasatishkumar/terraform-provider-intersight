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

// ApplianceBackupBase BackupBase is the parent type of Backup, Restore, and BackupPolicy managed objects. BackupBase holds the common information required for copying the file from the Intersight Appliance to the remote file server and vice versa.
type ApplianceBackupBase struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Backup filename to backup or restore.
	Filename *string `json:"Filename,omitempty"`
	// Communication protocol used by the file server (e.g. scp or sftp). * `scp` - Secure Copy Protocol (SCP) to access the file server. * `sftp` - SSH File Transfer Protocol (SFTP) to access file server.
	Protocol *string `json:"Protocol,omitempty"`
	// Hostname of the remote file server.
	RemoteHost *string `json:"RemoteHost,omitempty"`
	// File server directory to copy the file.
	RemotePath *string `json:"RemotePath,omitempty"`
	// Remote TCP port on the file server (e.g. 22 for scp).
	RemotePort *int64 `json:"RemotePort,omitempty"`
	// Username to authenticate the fileserver.
	Username             *string `json:"Username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceBackupBase ApplianceBackupBase

// NewApplianceBackupBase instantiates a new ApplianceBackupBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceBackupBase(classId string, objectType string) *ApplianceBackupBase {
	this := ApplianceBackupBase{}
	this.ClassId = classId
	this.ObjectType = objectType
	var protocol string = "scp"
	this.Protocol = &protocol
	return &this
}

// NewApplianceBackupBaseWithDefaults instantiates a new ApplianceBackupBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceBackupBaseWithDefaults() *ApplianceBackupBase {
	this := ApplianceBackupBase{}
	var protocol string = "scp"
	this.Protocol = &protocol
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceBackupBase) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceBackupBase) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceBackupBase) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceBackupBase) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceBackupBase) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceBackupBase) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *ApplianceBackupBase) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupBase) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *ApplianceBackupBase) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *ApplianceBackupBase) SetFilename(v string) {
	o.Filename = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ApplianceBackupBase) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupBase) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ApplianceBackupBase) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ApplianceBackupBase) SetProtocol(v string) {
	o.Protocol = &v
}

// GetRemoteHost returns the RemoteHost field value if set, zero value otherwise.
func (o *ApplianceBackupBase) GetRemoteHost() string {
	if o == nil || o.RemoteHost == nil {
		var ret string
		return ret
	}
	return *o.RemoteHost
}

// GetRemoteHostOk returns a tuple with the RemoteHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupBase) GetRemoteHostOk() (*string, bool) {
	if o == nil || o.RemoteHost == nil {
		return nil, false
	}
	return o.RemoteHost, true
}

// HasRemoteHost returns a boolean if a field has been set.
func (o *ApplianceBackupBase) HasRemoteHost() bool {
	if o != nil && o.RemoteHost != nil {
		return true
	}

	return false
}

// SetRemoteHost gets a reference to the given string and assigns it to the RemoteHost field.
func (o *ApplianceBackupBase) SetRemoteHost(v string) {
	o.RemoteHost = &v
}

// GetRemotePath returns the RemotePath field value if set, zero value otherwise.
func (o *ApplianceBackupBase) GetRemotePath() string {
	if o == nil || o.RemotePath == nil {
		var ret string
		return ret
	}
	return *o.RemotePath
}

// GetRemotePathOk returns a tuple with the RemotePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupBase) GetRemotePathOk() (*string, bool) {
	if o == nil || o.RemotePath == nil {
		return nil, false
	}
	return o.RemotePath, true
}

// HasRemotePath returns a boolean if a field has been set.
func (o *ApplianceBackupBase) HasRemotePath() bool {
	if o != nil && o.RemotePath != nil {
		return true
	}

	return false
}

// SetRemotePath gets a reference to the given string and assigns it to the RemotePath field.
func (o *ApplianceBackupBase) SetRemotePath(v string) {
	o.RemotePath = &v
}

// GetRemotePort returns the RemotePort field value if set, zero value otherwise.
func (o *ApplianceBackupBase) GetRemotePort() int64 {
	if o == nil || o.RemotePort == nil {
		var ret int64
		return ret
	}
	return *o.RemotePort
}

// GetRemotePortOk returns a tuple with the RemotePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupBase) GetRemotePortOk() (*int64, bool) {
	if o == nil || o.RemotePort == nil {
		return nil, false
	}
	return o.RemotePort, true
}

// HasRemotePort returns a boolean if a field has been set.
func (o *ApplianceBackupBase) HasRemotePort() bool {
	if o != nil && o.RemotePort != nil {
		return true
	}

	return false
}

// SetRemotePort gets a reference to the given int64 and assigns it to the RemotePort field.
func (o *ApplianceBackupBase) SetRemotePort(v int64) {
	o.RemotePort = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApplianceBackupBase) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceBackupBase) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApplianceBackupBase) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApplianceBackupBase) SetUsername(v string) {
	o.Username = &v
}

func (o ApplianceBackupBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Filename != nil {
		toSerialize["Filename"] = o.Filename
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.RemoteHost != nil {
		toSerialize["RemoteHost"] = o.RemoteHost
	}
	if o.RemotePath != nil {
		toSerialize["RemotePath"] = o.RemotePath
	}
	if o.RemotePort != nil {
		toSerialize["RemotePort"] = o.RemotePort
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceBackupBase) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceBackupBaseWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Backup filename to backup or restore.
		Filename *string `json:"Filename,omitempty"`
		// Communication protocol used by the file server (e.g. scp or sftp). * `scp` - Secure Copy Protocol (SCP) to access the file server. * `sftp` - SSH File Transfer Protocol (SFTP) to access file server.
		Protocol *string `json:"Protocol,omitempty"`
		// Hostname of the remote file server.
		RemoteHost *string `json:"RemoteHost,omitempty"`
		// File server directory to copy the file.
		RemotePath *string `json:"RemotePath,omitempty"`
		// Remote TCP port on the file server (e.g. 22 for scp).
		RemotePort *int64 `json:"RemotePort,omitempty"`
		// Username to authenticate the fileserver.
		Username *string `json:"Username,omitempty"`
	}

	varApplianceBackupBaseWithoutEmbeddedStruct := ApplianceBackupBaseWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceBackupBaseWithoutEmbeddedStruct)
	if err == nil {
		varApplianceBackupBase := _ApplianceBackupBase{}
		varApplianceBackupBase.ClassId = varApplianceBackupBaseWithoutEmbeddedStruct.ClassId
		varApplianceBackupBase.ObjectType = varApplianceBackupBaseWithoutEmbeddedStruct.ObjectType
		varApplianceBackupBase.Filename = varApplianceBackupBaseWithoutEmbeddedStruct.Filename
		varApplianceBackupBase.Protocol = varApplianceBackupBaseWithoutEmbeddedStruct.Protocol
		varApplianceBackupBase.RemoteHost = varApplianceBackupBaseWithoutEmbeddedStruct.RemoteHost
		varApplianceBackupBase.RemotePath = varApplianceBackupBaseWithoutEmbeddedStruct.RemotePath
		varApplianceBackupBase.RemotePort = varApplianceBackupBaseWithoutEmbeddedStruct.RemotePort
		varApplianceBackupBase.Username = varApplianceBackupBaseWithoutEmbeddedStruct.Username
		*o = ApplianceBackupBase(varApplianceBackupBase)
	} else {
		return err
	}

	varApplianceBackupBase := _ApplianceBackupBase{}

	err = json.Unmarshal(bytes, &varApplianceBackupBase)
	if err == nil {
		o.MoBaseMo = varApplianceBackupBase.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Filename")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "RemoteHost")
		delete(additionalProperties, "RemotePath")
		delete(additionalProperties, "RemotePort")
		delete(additionalProperties, "Username")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableApplianceBackupBase struct {
	value *ApplianceBackupBase
	isSet bool
}

func (v NullableApplianceBackupBase) Get() *ApplianceBackupBase {
	return v.value
}

func (v *NullableApplianceBackupBase) Set(val *ApplianceBackupBase) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceBackupBase) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceBackupBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceBackupBase(val *ApplianceBackupBase) *NullableApplianceBackupBase {
	return &NullableApplianceBackupBase{value: val, isSet: true}
}

func (v NullableApplianceBackupBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceBackupBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
