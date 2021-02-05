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

// AdapterAdapterConfig Global adapter level settings.
type AdapterAdapterConfig struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                             `json:"ObjectType"`
	DceInterfaceSettings []AdapterDceInterfaceSettings      `json:"DceInterfaceSettings,omitempty"`
	EthSettings          NullableAdapterEthSettings         `json:"EthSettings,omitempty"`
	FcSettings           NullableAdapterFcSettings          `json:"FcSettings,omitempty"`
	PortChannelSettings  NullableAdapterPortChannelSettings `json:"PortChannelSettings,omitempty"`
	// PCIe slot where the VIC adapter is installed. Supported values are (1-15) and MLOM.
	SlotId               *string `json:"SlotId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdapterAdapterConfig AdapterAdapterConfig

// NewAdapterAdapterConfig instantiates a new AdapterAdapterConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdapterAdapterConfig(classId string, objectType string) *AdapterAdapterConfig {
	this := AdapterAdapterConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAdapterAdapterConfigWithDefaults instantiates a new AdapterAdapterConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdapterAdapterConfigWithDefaults() *AdapterAdapterConfig {
	this := AdapterAdapterConfig{}
	var classId string = "adapter.AdapterConfig"
	this.ClassId = classId
	var objectType string = "adapter.AdapterConfig"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AdapterAdapterConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AdapterAdapterConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AdapterAdapterConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AdapterAdapterConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AdapterAdapterConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AdapterAdapterConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDceInterfaceSettings returns the DceInterfaceSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdapterAdapterConfig) GetDceInterfaceSettings() []AdapterDceInterfaceSettings {
	if o == nil {
		var ret []AdapterDceInterfaceSettings
		return ret
	}
	return o.DceInterfaceSettings
}

// GetDceInterfaceSettingsOk returns a tuple with the DceInterfaceSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdapterAdapterConfig) GetDceInterfaceSettingsOk() (*[]AdapterDceInterfaceSettings, bool) {
	if o == nil || o.DceInterfaceSettings == nil {
		return nil, false
	}
	return &o.DceInterfaceSettings, true
}

// HasDceInterfaceSettings returns a boolean if a field has been set.
func (o *AdapterAdapterConfig) HasDceInterfaceSettings() bool {
	if o != nil && o.DceInterfaceSettings != nil {
		return true
	}

	return false
}

// SetDceInterfaceSettings gets a reference to the given []AdapterDceInterfaceSettings and assigns it to the DceInterfaceSettings field.
func (o *AdapterAdapterConfig) SetDceInterfaceSettings(v []AdapterDceInterfaceSettings) {
	o.DceInterfaceSettings = v
}

// GetEthSettings returns the EthSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdapterAdapterConfig) GetEthSettings() AdapterEthSettings {
	if o == nil || o.EthSettings.Get() == nil {
		var ret AdapterEthSettings
		return ret
	}
	return *o.EthSettings.Get()
}

// GetEthSettingsOk returns a tuple with the EthSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdapterAdapterConfig) GetEthSettingsOk() (*AdapterEthSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.EthSettings.Get(), o.EthSettings.IsSet()
}

// HasEthSettings returns a boolean if a field has been set.
func (o *AdapterAdapterConfig) HasEthSettings() bool {
	if o != nil && o.EthSettings.IsSet() {
		return true
	}

	return false
}

// SetEthSettings gets a reference to the given NullableAdapterEthSettings and assigns it to the EthSettings field.
func (o *AdapterAdapterConfig) SetEthSettings(v AdapterEthSettings) {
	o.EthSettings.Set(&v)
}

// SetEthSettingsNil sets the value for EthSettings to be an explicit nil
func (o *AdapterAdapterConfig) SetEthSettingsNil() {
	o.EthSettings.Set(nil)
}

// UnsetEthSettings ensures that no value is present for EthSettings, not even an explicit nil
func (o *AdapterAdapterConfig) UnsetEthSettings() {
	o.EthSettings.Unset()
}

// GetFcSettings returns the FcSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdapterAdapterConfig) GetFcSettings() AdapterFcSettings {
	if o == nil || o.FcSettings.Get() == nil {
		var ret AdapterFcSettings
		return ret
	}
	return *o.FcSettings.Get()
}

// GetFcSettingsOk returns a tuple with the FcSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdapterAdapterConfig) GetFcSettingsOk() (*AdapterFcSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.FcSettings.Get(), o.FcSettings.IsSet()
}

// HasFcSettings returns a boolean if a field has been set.
func (o *AdapterAdapterConfig) HasFcSettings() bool {
	if o != nil && o.FcSettings.IsSet() {
		return true
	}

	return false
}

// SetFcSettings gets a reference to the given NullableAdapterFcSettings and assigns it to the FcSettings field.
func (o *AdapterAdapterConfig) SetFcSettings(v AdapterFcSettings) {
	o.FcSettings.Set(&v)
}

// SetFcSettingsNil sets the value for FcSettings to be an explicit nil
func (o *AdapterAdapterConfig) SetFcSettingsNil() {
	o.FcSettings.Set(nil)
}

// UnsetFcSettings ensures that no value is present for FcSettings, not even an explicit nil
func (o *AdapterAdapterConfig) UnsetFcSettings() {
	o.FcSettings.Unset()
}

// GetPortChannelSettings returns the PortChannelSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdapterAdapterConfig) GetPortChannelSettings() AdapterPortChannelSettings {
	if o == nil || o.PortChannelSettings.Get() == nil {
		var ret AdapterPortChannelSettings
		return ret
	}
	return *o.PortChannelSettings.Get()
}

// GetPortChannelSettingsOk returns a tuple with the PortChannelSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdapterAdapterConfig) GetPortChannelSettingsOk() (*AdapterPortChannelSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.PortChannelSettings.Get(), o.PortChannelSettings.IsSet()
}

// HasPortChannelSettings returns a boolean if a field has been set.
func (o *AdapterAdapterConfig) HasPortChannelSettings() bool {
	if o != nil && o.PortChannelSettings.IsSet() {
		return true
	}

	return false
}

// SetPortChannelSettings gets a reference to the given NullableAdapterPortChannelSettings and assigns it to the PortChannelSettings field.
func (o *AdapterAdapterConfig) SetPortChannelSettings(v AdapterPortChannelSettings) {
	o.PortChannelSettings.Set(&v)
}

// SetPortChannelSettingsNil sets the value for PortChannelSettings to be an explicit nil
func (o *AdapterAdapterConfig) SetPortChannelSettingsNil() {
	o.PortChannelSettings.Set(nil)
}

// UnsetPortChannelSettings ensures that no value is present for PortChannelSettings, not even an explicit nil
func (o *AdapterAdapterConfig) UnsetPortChannelSettings() {
	o.PortChannelSettings.Unset()
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *AdapterAdapterConfig) GetSlotId() string {
	if o == nil || o.SlotId == nil {
		var ret string
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterAdapterConfig) GetSlotIdOk() (*string, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *AdapterAdapterConfig) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given string and assigns it to the SlotId field.
func (o *AdapterAdapterConfig) SetSlotId(v string) {
	o.SlotId = &v
}

func (o AdapterAdapterConfig) MarshalJSON() ([]byte, error) {
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
	if o.DceInterfaceSettings != nil {
		toSerialize["DceInterfaceSettings"] = o.DceInterfaceSettings
	}
	if o.EthSettings.IsSet() {
		toSerialize["EthSettings"] = o.EthSettings.Get()
	}
	if o.FcSettings.IsSet() {
		toSerialize["FcSettings"] = o.FcSettings.Get()
	}
	if o.PortChannelSettings.IsSet() {
		toSerialize["PortChannelSettings"] = o.PortChannelSettings.Get()
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AdapterAdapterConfig) UnmarshalJSON(bytes []byte) (err error) {
	type AdapterAdapterConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType           string                             `json:"ObjectType"`
		DceInterfaceSettings []AdapterDceInterfaceSettings      `json:"DceInterfaceSettings,omitempty"`
		EthSettings          NullableAdapterEthSettings         `json:"EthSettings,omitempty"`
		FcSettings           NullableAdapterFcSettings          `json:"FcSettings,omitempty"`
		PortChannelSettings  NullableAdapterPortChannelSettings `json:"PortChannelSettings,omitempty"`
		// PCIe slot where the VIC adapter is installed. Supported values are (1-15) and MLOM.
		SlotId *string `json:"SlotId,omitempty"`
	}

	varAdapterAdapterConfigWithoutEmbeddedStruct := AdapterAdapterConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAdapterAdapterConfigWithoutEmbeddedStruct)
	if err == nil {
		varAdapterAdapterConfig := _AdapterAdapterConfig{}
		varAdapterAdapterConfig.ClassId = varAdapterAdapterConfigWithoutEmbeddedStruct.ClassId
		varAdapterAdapterConfig.ObjectType = varAdapterAdapterConfigWithoutEmbeddedStruct.ObjectType
		varAdapterAdapterConfig.DceInterfaceSettings = varAdapterAdapterConfigWithoutEmbeddedStruct.DceInterfaceSettings
		varAdapterAdapterConfig.EthSettings = varAdapterAdapterConfigWithoutEmbeddedStruct.EthSettings
		varAdapterAdapterConfig.FcSettings = varAdapterAdapterConfigWithoutEmbeddedStruct.FcSettings
		varAdapterAdapterConfig.PortChannelSettings = varAdapterAdapterConfigWithoutEmbeddedStruct.PortChannelSettings
		varAdapterAdapterConfig.SlotId = varAdapterAdapterConfigWithoutEmbeddedStruct.SlotId
		*o = AdapterAdapterConfig(varAdapterAdapterConfig)
	} else {
		return err
	}

	varAdapterAdapterConfig := _AdapterAdapterConfig{}

	err = json.Unmarshal(bytes, &varAdapterAdapterConfig)
	if err == nil {
		o.MoBaseComplexType = varAdapterAdapterConfig.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DceInterfaceSettings")
		delete(additionalProperties, "EthSettings")
		delete(additionalProperties, "FcSettings")
		delete(additionalProperties, "PortChannelSettings")
		delete(additionalProperties, "SlotId")

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

type NullableAdapterAdapterConfig struct {
	value *AdapterAdapterConfig
	isSet bool
}

func (v NullableAdapterAdapterConfig) Get() *AdapterAdapterConfig {
	return v.value
}

func (v *NullableAdapterAdapterConfig) Set(val *AdapterAdapterConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAdapterAdapterConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAdapterAdapterConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdapterAdapterConfig(val *AdapterAdapterConfig) *NullableAdapterAdapterConfig {
	return &NullableAdapterAdapterConfig{value: val, isSet: true}
}

func (v NullableAdapterAdapterConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdapterAdapterConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
