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
	"time"
)

// TamAdvisoryDefinition An Intersight Advisory. An advisory represents an identification of a potential issue and may also include  a recommendation for resolving the said issue. Advisories may be of different kind and severity. for e.g. It could be a security vulnerability or a performance issue or a hardware issue with different recommendations for resolving them.
type TamAdvisoryDefinition struct {
	TamBaseAdvisory
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType      string                         `json:"ObjectType"`
	Actions         []TamAction                    `json:"Actions,omitempty"`
	AdvisoryDetails NullableTamBaseAdvisoryDetails `json:"AdvisoryDetails,omitempty"`
	// Cisco generated identifier for the published security advisory.
	AdvisoryId     *string            `json:"AdvisoryId,omitempty"`
	ApiDataSources []TamApiDataSource `json:"ApiDataSources,omitempty"`
	// Date when the security advisory was first published by Cisco.
	DatePublished *time.Time `json:"DatePublished,omitempty"`
	// Date when the security advisory was last updated by Cisco.
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	// A link to an external URL describing security Advisory in more details.
	ExternalUrl *string `json:"ExternalUrl,omitempty"`
	// Recommended action to resolve the security advisory.
	Recommendation *string `json:"Recommendation,omitempty"`
	// The type (field notice, security advisory etc.) of Intersight advisory. * `securityAdvisory` - Respresents the psirt alert type (https://tools.cisco.com/security/center/publicationListing.x). * `fieldNotice` - Respresents the field notice alert type (https://www.cisco.com/c/en/us/support/web/tsd-products-field-notice-summary.html).
	Type *string `json:"Type,omitempty"`
	// Cisco assigned advisory version after latest revision.
	Version *string `json:"Version,omitempty"`
	// Workarounds available for the advisory.
	Workaround           *string                               `json:"Workaround,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TamAdvisoryDefinition TamAdvisoryDefinition

// NewTamAdvisoryDefinition instantiates a new TamAdvisoryDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTamAdvisoryDefinition(classId string, objectType string) *TamAdvisoryDefinition {
	this := TamAdvisoryDefinition{}
	this.ClassId = classId
	this.ObjectType = objectType
	var type_ string = "securityAdvisory"
	this.Type = &type_
	return &this
}

// NewTamAdvisoryDefinitionWithDefaults instantiates a new TamAdvisoryDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTamAdvisoryDefinitionWithDefaults() *TamAdvisoryDefinition {
	this := TamAdvisoryDefinition{}
	var classId string = "tam.AdvisoryDefinition"
	this.ClassId = classId
	var objectType string = "tam.AdvisoryDefinition"
	this.ObjectType = objectType
	var type_ string = "securityAdvisory"
	this.Type = &type_
	return &this
}

// GetClassId returns the ClassId field value
func (o *TamAdvisoryDefinition) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinition) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *TamAdvisoryDefinition) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *TamAdvisoryDefinition) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinition) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *TamAdvisoryDefinition) SetObjectType(v string) {
	o.ObjectType = v
}

// GetActions returns the Actions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TamAdvisoryDefinition) GetActions() []TamAction {
	if o == nil {
		var ret []TamAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TamAdvisoryDefinition) GetActionsOk() (*[]TamAction, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return &o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *TamAdvisoryDefinition) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []TamAction and assigns it to the Actions field.
func (o *TamAdvisoryDefinition) SetActions(v []TamAction) {
	o.Actions = v
}

// GetAdvisoryDetails returns the AdvisoryDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TamAdvisoryDefinition) GetAdvisoryDetails() TamBaseAdvisoryDetails {
	if o == nil || o.AdvisoryDetails.Get() == nil {
		var ret TamBaseAdvisoryDetails
		return ret
	}
	return *o.AdvisoryDetails.Get()
}

// GetAdvisoryDetailsOk returns a tuple with the AdvisoryDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TamAdvisoryDefinition) GetAdvisoryDetailsOk() (*TamBaseAdvisoryDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdvisoryDetails.Get(), o.AdvisoryDetails.IsSet()
}

// HasAdvisoryDetails returns a boolean if a field has been set.
func (o *TamAdvisoryDefinition) HasAdvisoryDetails() bool {
	if o != nil && o.AdvisoryDetails.IsSet() {
		return true
	}

	return false
}

// SetAdvisoryDetails gets a reference to the given NullableTamBaseAdvisoryDetails and assigns it to the AdvisoryDetails field.
func (o *TamAdvisoryDefinition) SetAdvisoryDetails(v TamBaseAdvisoryDetails) {
	o.AdvisoryDetails.Set(&v)
}

// SetAdvisoryDetailsNil sets the value for AdvisoryDetails to be an explicit nil
func (o *TamAdvisoryDefinition) SetAdvisoryDetailsNil() {
	o.AdvisoryDetails.Set(nil)
}

// UnsetAdvisoryDetails ensures that no value is present for AdvisoryDetails, not even an explicit nil
func (o *TamAdvisoryDefinition) UnsetAdvisoryDetails() {
	o.AdvisoryDetails.Unset()
}

// GetAdvisoryId returns the AdvisoryId field value if set, zero value otherwise.
func (o *TamAdvisoryDefinition) GetAdvisoryId() string {
	if o == nil || o.AdvisoryId == nil {
		var ret string
		return ret
	}
	return *o.AdvisoryId
}

// GetAdvisoryIdOk returns a tuple with the AdvisoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinition) GetAdvisoryIdOk() (*string, bool) {
	if o == nil || o.AdvisoryId == nil {
		return nil, false
	}
	return o.AdvisoryId, true
}

// HasAdvisoryId returns a boolean if a field has been set.
func (o *TamAdvisoryDefinition) HasAdvisoryId() bool {
	if o != nil && o.AdvisoryId != nil {
		return true
	}

	return false
}

// SetAdvisoryId gets a reference to the given string and assigns it to the AdvisoryId field.
func (o *TamAdvisoryDefinition) SetAdvisoryId(v string) {
	o.AdvisoryId = &v
}

// GetApiDataSources returns the ApiDataSources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TamAdvisoryDefinition) GetApiDataSources() []TamApiDataSource {
	if o == nil {
		var ret []TamApiDataSource
		return ret
	}
	return o.ApiDataSources
}

// GetApiDataSourcesOk returns a tuple with the ApiDataSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TamAdvisoryDefinition) GetApiDataSourcesOk() (*[]TamApiDataSource, bool) {
	if o == nil || o.ApiDataSources == nil {
		return nil, false
	}
	return &o.ApiDataSources, true
}

// HasApiDataSources returns a boolean if a field has been set.
func (o *TamAdvisoryDefinition) HasApiDataSources() bool {
	if o != nil && o.ApiDataSources != nil {
		return true
	}

	return false
}

// SetApiDataSources gets a reference to the given []TamApiDataSource and assigns it to the ApiDataSources field.
func (o *TamAdvisoryDefinition) SetApiDataSources(v []TamApiDataSource) {
	o.ApiDataSources = v
}

// GetDatePublished returns the DatePublished field value if set, zero value otherwise.
func (o *TamAdvisoryDefinition) GetDatePublished() time.Time {
	if o == nil || o.DatePublished == nil {
		var ret time.Time
		return ret
	}
	return *o.DatePublished
}

// GetDatePublishedOk returns a tuple with the DatePublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinition) GetDatePublishedOk() (*time.Time, bool) {
	if o == nil || o.DatePublished == nil {
		return nil, false
	}
	return o.DatePublished, true
}

// HasDatePublished returns a boolean if a field has been set.
func (o *TamAdvisoryDefinition) HasDatePublished() bool {
	if o != nil && o.DatePublished != nil {
		return true
	}

	return false
}

// SetDatePublished gets a reference to the given time.Time and assigns it to the DatePublished field.
func (o *TamAdvisoryDefinition) SetDatePublished(v time.Time) {
	o.DatePublished = &v
}

// GetDateUpdated returns the DateUpdated field value if set, zero value otherwise.
func (o *TamAdvisoryDefinition) GetDateUpdated() time.Time {
	if o == nil || o.DateUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateUpdated
}

// GetDateUpdatedOk returns a tuple with the DateUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinition) GetDateUpdatedOk() (*time.Time, bool) {
	if o == nil || o.DateUpdated == nil {
		return nil, false
	}
	return o.DateUpdated, true
}

// HasDateUpdated returns a boolean if a field has been set.
func (o *TamAdvisoryDefinition) HasDateUpdated() bool {
	if o != nil && o.DateUpdated != nil {
		return true
	}

	return false
}

// SetDateUpdated gets a reference to the given time.Time and assigns it to the DateUpdated field.
func (o *TamAdvisoryDefinition) SetDateUpdated(v time.Time) {
	o.DateUpdated = &v
}

// GetExternalUrl returns the ExternalUrl field value if set, zero value otherwise.
func (o *TamAdvisoryDefinition) GetExternalUrl() string {
	if o == nil || o.ExternalUrl == nil {
		var ret string
		return ret
	}
	return *o.ExternalUrl
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinition) GetExternalUrlOk() (*string, bool) {
	if o == nil || o.ExternalUrl == nil {
		return nil, false
	}
	return o.ExternalUrl, true
}

// HasExternalUrl returns a boolean if a field has been set.
func (o *TamAdvisoryDefinition) HasExternalUrl() bool {
	if o != nil && o.ExternalUrl != nil {
		return true
	}

	return false
}

// SetExternalUrl gets a reference to the given string and assigns it to the ExternalUrl field.
func (o *TamAdvisoryDefinition) SetExternalUrl(v string) {
	o.ExternalUrl = &v
}

// GetRecommendation returns the Recommendation field value if set, zero value otherwise.
func (o *TamAdvisoryDefinition) GetRecommendation() string {
	if o == nil || o.Recommendation == nil {
		var ret string
		return ret
	}
	return *o.Recommendation
}

// GetRecommendationOk returns a tuple with the Recommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinition) GetRecommendationOk() (*string, bool) {
	if o == nil || o.Recommendation == nil {
		return nil, false
	}
	return o.Recommendation, true
}

// HasRecommendation returns a boolean if a field has been set.
func (o *TamAdvisoryDefinition) HasRecommendation() bool {
	if o != nil && o.Recommendation != nil {
		return true
	}

	return false
}

// SetRecommendation gets a reference to the given string and assigns it to the Recommendation field.
func (o *TamAdvisoryDefinition) SetRecommendation(v string) {
	o.Recommendation = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TamAdvisoryDefinition) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinition) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TamAdvisoryDefinition) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TamAdvisoryDefinition) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *TamAdvisoryDefinition) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinition) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *TamAdvisoryDefinition) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *TamAdvisoryDefinition) SetVersion(v string) {
	o.Version = &v
}

// GetWorkaround returns the Workaround field value if set, zero value otherwise.
func (o *TamAdvisoryDefinition) GetWorkaround() string {
	if o == nil || o.Workaround == nil {
		var ret string
		return ret
	}
	return *o.Workaround
}

// GetWorkaroundOk returns a tuple with the Workaround field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinition) GetWorkaroundOk() (*string, bool) {
	if o == nil || o.Workaround == nil {
		return nil, false
	}
	return o.Workaround, true
}

// HasWorkaround returns a boolean if a field has been set.
func (o *TamAdvisoryDefinition) HasWorkaround() bool {
	if o != nil && o.Workaround != nil {
		return true
	}

	return false
}

// SetWorkaround gets a reference to the given string and assigns it to the Workaround field.
func (o *TamAdvisoryDefinition) SetWorkaround(v string) {
	o.Workaround = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *TamAdvisoryDefinition) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TamAdvisoryDefinition) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *TamAdvisoryDefinition) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *TamAdvisoryDefinition) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o TamAdvisoryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTamBaseAdvisory, errTamBaseAdvisory := json.Marshal(o.TamBaseAdvisory)
	if errTamBaseAdvisory != nil {
		return []byte{}, errTamBaseAdvisory
	}
	errTamBaseAdvisory = json.Unmarshal([]byte(serializedTamBaseAdvisory), &toSerialize)
	if errTamBaseAdvisory != nil {
		return []byte{}, errTamBaseAdvisory
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Actions != nil {
		toSerialize["Actions"] = o.Actions
	}
	if o.AdvisoryDetails.IsSet() {
		toSerialize["AdvisoryDetails"] = o.AdvisoryDetails.Get()
	}
	if o.AdvisoryId != nil {
		toSerialize["AdvisoryId"] = o.AdvisoryId
	}
	if o.ApiDataSources != nil {
		toSerialize["ApiDataSources"] = o.ApiDataSources
	}
	if o.DatePublished != nil {
		toSerialize["DatePublished"] = o.DatePublished
	}
	if o.DateUpdated != nil {
		toSerialize["DateUpdated"] = o.DateUpdated
	}
	if o.ExternalUrl != nil {
		toSerialize["ExternalUrl"] = o.ExternalUrl
	}
	if o.Recommendation != nil {
		toSerialize["Recommendation"] = o.Recommendation
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Workaround != nil {
		toSerialize["Workaround"] = o.Workaround
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TamAdvisoryDefinition) UnmarshalJSON(bytes []byte) (err error) {
	type TamAdvisoryDefinitionWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType      string                         `json:"ObjectType"`
		Actions         []TamAction                    `json:"Actions,omitempty"`
		AdvisoryDetails NullableTamBaseAdvisoryDetails `json:"AdvisoryDetails,omitempty"`
		// Cisco generated identifier for the published security advisory.
		AdvisoryId     *string            `json:"AdvisoryId,omitempty"`
		ApiDataSources []TamApiDataSource `json:"ApiDataSources,omitempty"`
		// Date when the security advisory was first published by Cisco.
		DatePublished *time.Time `json:"DatePublished,omitempty"`
		// Date when the security advisory was last updated by Cisco.
		DateUpdated *time.Time `json:"DateUpdated,omitempty"`
		// A link to an external URL describing security Advisory in more details.
		ExternalUrl *string `json:"ExternalUrl,omitempty"`
		// Recommended action to resolve the security advisory.
		Recommendation *string `json:"Recommendation,omitempty"`
		// The type (field notice, security advisory etc.) of Intersight advisory. * `securityAdvisory` - Respresents the psirt alert type (https://tools.cisco.com/security/center/publicationListing.x). * `fieldNotice` - Respresents the field notice alert type (https://www.cisco.com/c/en/us/support/web/tsd-products-field-notice-summary.html).
		Type *string `json:"Type,omitempty"`
		// Cisco assigned advisory version after latest revision.
		Version *string `json:"Version,omitempty"`
		// Workarounds available for the advisory.
		Workaround   *string                               `json:"Workaround,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varTamAdvisoryDefinitionWithoutEmbeddedStruct := TamAdvisoryDefinitionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varTamAdvisoryDefinitionWithoutEmbeddedStruct)
	if err == nil {
		varTamAdvisoryDefinition := _TamAdvisoryDefinition{}
		varTamAdvisoryDefinition.ClassId = varTamAdvisoryDefinitionWithoutEmbeddedStruct.ClassId
		varTamAdvisoryDefinition.ObjectType = varTamAdvisoryDefinitionWithoutEmbeddedStruct.ObjectType
		varTamAdvisoryDefinition.Actions = varTamAdvisoryDefinitionWithoutEmbeddedStruct.Actions
		varTamAdvisoryDefinition.AdvisoryDetails = varTamAdvisoryDefinitionWithoutEmbeddedStruct.AdvisoryDetails
		varTamAdvisoryDefinition.AdvisoryId = varTamAdvisoryDefinitionWithoutEmbeddedStruct.AdvisoryId
		varTamAdvisoryDefinition.ApiDataSources = varTamAdvisoryDefinitionWithoutEmbeddedStruct.ApiDataSources
		varTamAdvisoryDefinition.DatePublished = varTamAdvisoryDefinitionWithoutEmbeddedStruct.DatePublished
		varTamAdvisoryDefinition.DateUpdated = varTamAdvisoryDefinitionWithoutEmbeddedStruct.DateUpdated
		varTamAdvisoryDefinition.ExternalUrl = varTamAdvisoryDefinitionWithoutEmbeddedStruct.ExternalUrl
		varTamAdvisoryDefinition.Recommendation = varTamAdvisoryDefinitionWithoutEmbeddedStruct.Recommendation
		varTamAdvisoryDefinition.Type = varTamAdvisoryDefinitionWithoutEmbeddedStruct.Type
		varTamAdvisoryDefinition.Version = varTamAdvisoryDefinitionWithoutEmbeddedStruct.Version
		varTamAdvisoryDefinition.Workaround = varTamAdvisoryDefinitionWithoutEmbeddedStruct.Workaround
		varTamAdvisoryDefinition.Organization = varTamAdvisoryDefinitionWithoutEmbeddedStruct.Organization
		*o = TamAdvisoryDefinition(varTamAdvisoryDefinition)
	} else {
		return err
	}

	varTamAdvisoryDefinition := _TamAdvisoryDefinition{}

	err = json.Unmarshal(bytes, &varTamAdvisoryDefinition)
	if err == nil {
		o.TamBaseAdvisory = varTamAdvisoryDefinition.TamBaseAdvisory
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Actions")
		delete(additionalProperties, "AdvisoryDetails")
		delete(additionalProperties, "AdvisoryId")
		delete(additionalProperties, "ApiDataSources")
		delete(additionalProperties, "DatePublished")
		delete(additionalProperties, "DateUpdated")
		delete(additionalProperties, "ExternalUrl")
		delete(additionalProperties, "Recommendation")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Workaround")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectTamBaseAdvisory := reflect.ValueOf(o.TamBaseAdvisory)
		for i := 0; i < reflectTamBaseAdvisory.Type().NumField(); i++ {
			t := reflectTamBaseAdvisory.Type().Field(i)

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

type NullableTamAdvisoryDefinition struct {
	value *TamAdvisoryDefinition
	isSet bool
}

func (v NullableTamAdvisoryDefinition) Get() *TamAdvisoryDefinition {
	return v.value
}

func (v *NullableTamAdvisoryDefinition) Set(val *TamAdvisoryDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableTamAdvisoryDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableTamAdvisoryDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTamAdvisoryDefinition(val *TamAdvisoryDefinition) *NullableTamAdvisoryDefinition {
	return &NullableTamAdvisoryDefinition{value: val, isSet: true}
}

func (v NullableTamAdvisoryDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTamAdvisoryDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
