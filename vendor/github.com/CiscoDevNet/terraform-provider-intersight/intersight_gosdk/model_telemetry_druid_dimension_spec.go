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
	"fmt"
)

// TelemetryDruidDimensionSpec - struct for TelemetryDruidDimensionSpec
type TelemetryDruidDimensionSpec struct {
	TelemetryDruidDefaultDimensionSpec    *TelemetryDruidDefaultDimensionSpec
	TelemetryDruidExtractionDimensionSpec *TelemetryDruidExtractionDimensionSpec
}

// TelemetryDruidDefaultDimensionSpecAsTelemetryDruidDimensionSpec is a convenience function that returns TelemetryDruidDefaultDimensionSpec wrapped in TelemetryDruidDimensionSpec
func TelemetryDruidDefaultDimensionSpecAsTelemetryDruidDimensionSpec(v *TelemetryDruidDefaultDimensionSpec) TelemetryDruidDimensionSpec {
	return TelemetryDruidDimensionSpec{TelemetryDruidDefaultDimensionSpec: v}
}

// TelemetryDruidExtractionDimensionSpecAsTelemetryDruidDimensionSpec is a convenience function that returns TelemetryDruidExtractionDimensionSpec wrapped in TelemetryDruidDimensionSpec
func TelemetryDruidExtractionDimensionSpecAsTelemetryDruidDimensionSpec(v *TelemetryDruidExtractionDimensionSpec) TelemetryDruidDimensionSpec {
	return TelemetryDruidDimensionSpec{TelemetryDruidExtractionDimensionSpec: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *TelemetryDruidDimensionSpec) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'default'
	if jsonDict["type"] == "default" {
		// try to unmarshal JSON data into TelemetryDruidDefaultDimensionSpec
		err = json.Unmarshal(data, &dst.TelemetryDruidDefaultDimensionSpec)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidDefaultDimensionSpec, return on the first match
		} else {
			dst.TelemetryDruidDefaultDimensionSpec = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidDimensionSpec as TelemetryDruidDefaultDimensionSpec: %s", err.Error())
		}
	}

	// check if the discriminator value is 'extraction'
	if jsonDict["type"] == "extraction" {
		// try to unmarshal JSON data into TelemetryDruidExtractionDimensionSpec
		err = json.Unmarshal(data, &dst.TelemetryDruidExtractionDimensionSpec)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidExtractionDimensionSpec, return on the first match
		} else {
			dst.TelemetryDruidExtractionDimensionSpec = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidDimensionSpec as TelemetryDruidExtractionDimensionSpec: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidDefaultDimensionSpec'
	if jsonDict["type"] == "telemetry.DruidDefaultDimensionSpec" {
		// try to unmarshal JSON data into TelemetryDruidDefaultDimensionSpec
		err = json.Unmarshal(data, &dst.TelemetryDruidDefaultDimensionSpec)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidDefaultDimensionSpec, return on the first match
		} else {
			dst.TelemetryDruidDefaultDimensionSpec = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidDimensionSpec as TelemetryDruidDefaultDimensionSpec: %s", err.Error())
		}
	}

	// check if the discriminator value is 'telemetry.DruidExtractionDimensionSpec'
	if jsonDict["type"] == "telemetry.DruidExtractionDimensionSpec" {
		// try to unmarshal JSON data into TelemetryDruidExtractionDimensionSpec
		err = json.Unmarshal(data, &dst.TelemetryDruidExtractionDimensionSpec)
		if err == nil {
			return nil // data stored in dst.TelemetryDruidExtractionDimensionSpec, return on the first match
		} else {
			dst.TelemetryDruidExtractionDimensionSpec = nil
			return fmt.Errorf("Failed to unmarshal TelemetryDruidDimensionSpec as TelemetryDruidExtractionDimensionSpec: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TelemetryDruidDimensionSpec) MarshalJSON() ([]byte, error) {
	if src.TelemetryDruidDefaultDimensionSpec != nil {
		return json.Marshal(&src.TelemetryDruidDefaultDimensionSpec)
	}

	if src.TelemetryDruidExtractionDimensionSpec != nil {
		return json.Marshal(&src.TelemetryDruidExtractionDimensionSpec)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TelemetryDruidDimensionSpec) GetActualInstance() interface{} {
	if obj.TelemetryDruidDefaultDimensionSpec != nil {
		return obj.TelemetryDruidDefaultDimensionSpec
	}

	if obj.TelemetryDruidExtractionDimensionSpec != nil {
		return obj.TelemetryDruidExtractionDimensionSpec
	}

	// all schemas are nil
	return nil
}

type NullableTelemetryDruidDimensionSpec struct {
	value *TelemetryDruidDimensionSpec
	isSet bool
}

func (v NullableTelemetryDruidDimensionSpec) Get() *TelemetryDruidDimensionSpec {
	return v.value
}

func (v *NullableTelemetryDruidDimensionSpec) Set(val *TelemetryDruidDimensionSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidDimensionSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidDimensionSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidDimensionSpec(val *TelemetryDruidDimensionSpec) *NullableTelemetryDruidDimensionSpec {
	return &NullableTelemetryDruidDimensionSpec{value: val, isSet: true}
}

func (v NullableTelemetryDruidDimensionSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidDimensionSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
