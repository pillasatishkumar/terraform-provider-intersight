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
	"reflect"
	"strings"
)

// WorkflowBatchApiExecutor Intersight allows generic API tasks to be created by taking the API request body and a response parser specification in the form of content.Grammar object. Batch API associates the list of API requests to be executed as part of single task execution. Each API request takes the request body and a response parser specification.
type WorkflowBatchApiExecutor struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType  string                          `json:"ObjectType"`
	Batch       []WorkflowApi                   `json:"Batch,omitempty"`
	Constraints NullableWorkflowTaskConstraints `json:"Constraints,omitempty"`
	// A detailed description about the batch APIs.
	Description *string `json:"Description,omitempty"`
	// Name for the batch API task.
	Name *string `json:"Name,omitempty"`
	// All the possible outcomes of this task are captured here. Outcomes property is a collection property of type workflow.Outcome objects. The outcomes can be mapped to the message to be shown. The outcomes are evaluated in the order they are given. At the end of the outcomes list, an catchall success/fail outcome can be added with condition as 'true'. This is an optional property and if not specified the task will be marked as success.
	Outcomes interface{} `json:"Outcomes,omitempty"`
	// Intersight Orchestrator allows the extraction of required values from API responses using the API response grammar. These extracted values can be mapped to task output parameters defined in task definition. The mapping of API output parameters to the task output parameters is provided as JSON in this property.
	Output interface{} `json:"Output,omitempty"`
	// When an execution of a nth API in the Batch fails, Retry from falied API flag indicates if the execution should start from the nth API or the first API during task retry. By default the value is set to false.
	RetryFromFailedApi *bool `json:"RetryFromFailedApi,omitempty"`
	// The skip expression, if provided, allows the batch API executor to skip the task execution when the given expression evaluates to true. The expression is given as such a golang template that has to be evaluated to a final content true/false. The expression is an optional and in case not provided, the API will always be executed.
	SkipOnCondition      *string                                   `json:"SkipOnCondition,omitempty"`
	ErrorResponseHandler *WorkflowErrorResponseHandlerRelationship `json:"ErrorResponseHandler,omitempty"`
	TaskDefinition       *WorkflowTaskDefinitionRelationship       `json:"TaskDefinition,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowBatchApiExecutor WorkflowBatchApiExecutor

// NewWorkflowBatchApiExecutor instantiates a new WorkflowBatchApiExecutor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowBatchApiExecutor(classId string, objectType string) *WorkflowBatchApiExecutor {
	this := WorkflowBatchApiExecutor{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowBatchApiExecutorWithDefaults instantiates a new WorkflowBatchApiExecutor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowBatchApiExecutorWithDefaults() *WorkflowBatchApiExecutor {
	this := WorkflowBatchApiExecutor{}
	var classId string = "workflow.BatchApiExecutor"
	this.ClassId = classId
	var objectType string = "workflow.BatchApiExecutor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowBatchApiExecutor) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutor) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowBatchApiExecutor) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowBatchApiExecutor) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutor) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowBatchApiExecutor) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBatch returns the Batch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBatchApiExecutor) GetBatch() []WorkflowApi {
	if o == nil {
		var ret []WorkflowApi
		return ret
	}
	return o.Batch
}

// GetBatchOk returns a tuple with the Batch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBatchApiExecutor) GetBatchOk() (*[]WorkflowApi, bool) {
	if o == nil || o.Batch == nil {
		return nil, false
	}
	return &o.Batch, true
}

// HasBatch returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutor) HasBatch() bool {
	if o != nil && o.Batch != nil {
		return true
	}

	return false
}

// SetBatch gets a reference to the given []WorkflowApi and assigns it to the Batch field.
func (o *WorkflowBatchApiExecutor) SetBatch(v []WorkflowApi) {
	o.Batch = v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBatchApiExecutor) GetConstraints() WorkflowTaskConstraints {
	if o == nil || o.Constraints.Get() == nil {
		var ret WorkflowTaskConstraints
		return ret
	}
	return *o.Constraints.Get()
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBatchApiExecutor) GetConstraintsOk() (*WorkflowTaskConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints.Get(), o.Constraints.IsSet()
}

// HasConstraints returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutor) HasConstraints() bool {
	if o != nil && o.Constraints.IsSet() {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given NullableWorkflowTaskConstraints and assigns it to the Constraints field.
func (o *WorkflowBatchApiExecutor) SetConstraints(v WorkflowTaskConstraints) {
	o.Constraints.Set(&v)
}

// SetConstraintsNil sets the value for Constraints to be an explicit nil
func (o *WorkflowBatchApiExecutor) SetConstraintsNil() {
	o.Constraints.Set(nil)
}

// UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
func (o *WorkflowBatchApiExecutor) UnsetConstraints() {
	o.Constraints.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowBatchApiExecutor) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutor) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutor) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowBatchApiExecutor) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowBatchApiExecutor) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutor) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutor) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowBatchApiExecutor) SetName(v string) {
	o.Name = &v
}

// GetOutcomes returns the Outcomes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBatchApiExecutor) GetOutcomes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Outcomes
}

// GetOutcomesOk returns a tuple with the Outcomes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBatchApiExecutor) GetOutcomesOk() (*interface{}, bool) {
	if o == nil || o.Outcomes == nil {
		return nil, false
	}
	return &o.Outcomes, true
}

// HasOutcomes returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutor) HasOutcomes() bool {
	if o != nil && o.Outcomes != nil {
		return true
	}

	return false
}

// SetOutcomes gets a reference to the given interface{} and assigns it to the Outcomes field.
func (o *WorkflowBatchApiExecutor) SetOutcomes(v interface{}) {
	o.Outcomes = v
}

// GetOutput returns the Output field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowBatchApiExecutor) GetOutput() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowBatchApiExecutor) GetOutputOk() (*interface{}, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return &o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutor) HasOutput() bool {
	if o != nil && o.Output != nil {
		return true
	}

	return false
}

// SetOutput gets a reference to the given interface{} and assigns it to the Output field.
func (o *WorkflowBatchApiExecutor) SetOutput(v interface{}) {
	o.Output = v
}

// GetRetryFromFailedApi returns the RetryFromFailedApi field value if set, zero value otherwise.
func (o *WorkflowBatchApiExecutor) GetRetryFromFailedApi() bool {
	if o == nil || o.RetryFromFailedApi == nil {
		var ret bool
		return ret
	}
	return *o.RetryFromFailedApi
}

// GetRetryFromFailedApiOk returns a tuple with the RetryFromFailedApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutor) GetRetryFromFailedApiOk() (*bool, bool) {
	if o == nil || o.RetryFromFailedApi == nil {
		return nil, false
	}
	return o.RetryFromFailedApi, true
}

// HasRetryFromFailedApi returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutor) HasRetryFromFailedApi() bool {
	if o != nil && o.RetryFromFailedApi != nil {
		return true
	}

	return false
}

// SetRetryFromFailedApi gets a reference to the given bool and assigns it to the RetryFromFailedApi field.
func (o *WorkflowBatchApiExecutor) SetRetryFromFailedApi(v bool) {
	o.RetryFromFailedApi = &v
}

// GetSkipOnCondition returns the SkipOnCondition field value if set, zero value otherwise.
func (o *WorkflowBatchApiExecutor) GetSkipOnCondition() string {
	if o == nil || o.SkipOnCondition == nil {
		var ret string
		return ret
	}
	return *o.SkipOnCondition
}

// GetSkipOnConditionOk returns a tuple with the SkipOnCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutor) GetSkipOnConditionOk() (*string, bool) {
	if o == nil || o.SkipOnCondition == nil {
		return nil, false
	}
	return o.SkipOnCondition, true
}

// HasSkipOnCondition returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutor) HasSkipOnCondition() bool {
	if o != nil && o.SkipOnCondition != nil {
		return true
	}

	return false
}

// SetSkipOnCondition gets a reference to the given string and assigns it to the SkipOnCondition field.
func (o *WorkflowBatchApiExecutor) SetSkipOnCondition(v string) {
	o.SkipOnCondition = &v
}

// GetErrorResponseHandler returns the ErrorResponseHandler field value if set, zero value otherwise.
func (o *WorkflowBatchApiExecutor) GetErrorResponseHandler() WorkflowErrorResponseHandlerRelationship {
	if o == nil || o.ErrorResponseHandler == nil {
		var ret WorkflowErrorResponseHandlerRelationship
		return ret
	}
	return *o.ErrorResponseHandler
}

// GetErrorResponseHandlerOk returns a tuple with the ErrorResponseHandler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutor) GetErrorResponseHandlerOk() (*WorkflowErrorResponseHandlerRelationship, bool) {
	if o == nil || o.ErrorResponseHandler == nil {
		return nil, false
	}
	return o.ErrorResponseHandler, true
}

// HasErrorResponseHandler returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutor) HasErrorResponseHandler() bool {
	if o != nil && o.ErrorResponseHandler != nil {
		return true
	}

	return false
}

// SetErrorResponseHandler gets a reference to the given WorkflowErrorResponseHandlerRelationship and assigns it to the ErrorResponseHandler field.
func (o *WorkflowBatchApiExecutor) SetErrorResponseHandler(v WorkflowErrorResponseHandlerRelationship) {
	o.ErrorResponseHandler = &v
}

// GetTaskDefinition returns the TaskDefinition field value if set, zero value otherwise.
func (o *WorkflowBatchApiExecutor) GetTaskDefinition() WorkflowTaskDefinitionRelationship {
	if o == nil || o.TaskDefinition == nil {
		var ret WorkflowTaskDefinitionRelationship
		return ret
	}
	return *o.TaskDefinition
}

// GetTaskDefinitionOk returns a tuple with the TaskDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowBatchApiExecutor) GetTaskDefinitionOk() (*WorkflowTaskDefinitionRelationship, bool) {
	if o == nil || o.TaskDefinition == nil {
		return nil, false
	}
	return o.TaskDefinition, true
}

// HasTaskDefinition returns a boolean if a field has been set.
func (o *WorkflowBatchApiExecutor) HasTaskDefinition() bool {
	if o != nil && o.TaskDefinition != nil {
		return true
	}

	return false
}

// SetTaskDefinition gets a reference to the given WorkflowTaskDefinitionRelationship and assigns it to the TaskDefinition field.
func (o *WorkflowBatchApiExecutor) SetTaskDefinition(v WorkflowTaskDefinitionRelationship) {
	o.TaskDefinition = &v
}

func (o WorkflowBatchApiExecutor) MarshalJSON() ([]byte, error) {
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
	if o.Batch != nil {
		toSerialize["Batch"] = o.Batch
	}
	if o.Constraints.IsSet() {
		toSerialize["Constraints"] = o.Constraints.Get()
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Outcomes != nil {
		toSerialize["Outcomes"] = o.Outcomes
	}
	if o.Output != nil {
		toSerialize["Output"] = o.Output
	}
	if o.RetryFromFailedApi != nil {
		toSerialize["RetryFromFailedApi"] = o.RetryFromFailedApi
	}
	if o.SkipOnCondition != nil {
		toSerialize["SkipOnCondition"] = o.SkipOnCondition
	}
	if o.ErrorResponseHandler != nil {
		toSerialize["ErrorResponseHandler"] = o.ErrorResponseHandler
	}
	if o.TaskDefinition != nil {
		toSerialize["TaskDefinition"] = o.TaskDefinition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowBatchApiExecutor) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowBatchApiExecutorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType  string                          `json:"ObjectType"`
		Batch       []WorkflowApi                   `json:"Batch,omitempty"`
		Constraints NullableWorkflowTaskConstraints `json:"Constraints,omitempty"`
		// A detailed description about the batch APIs.
		Description *string `json:"Description,omitempty"`
		// Name for the batch API task.
		Name *string `json:"Name,omitempty"`
		// All the possible outcomes of this task are captured here. Outcomes property is a collection property of type workflow.Outcome objects. The outcomes can be mapped to the message to be shown. The outcomes are evaluated in the order they are given. At the end of the outcomes list, an catchall success/fail outcome can be added with condition as 'true'. This is an optional property and if not specified the task will be marked as success.
		Outcomes interface{} `json:"Outcomes,omitempty"`
		// Intersight Orchestrator allows the extraction of required values from API responses using the API response grammar. These extracted values can be mapped to task output parameters defined in task definition. The mapping of API output parameters to the task output parameters is provided as JSON in this property.
		Output interface{} `json:"Output,omitempty"`
		// When an execution of a nth API in the Batch fails, Retry from falied API flag indicates if the execution should start from the nth API or the first API during task retry. By default the value is set to false.
		RetryFromFailedApi *bool `json:"RetryFromFailedApi,omitempty"`
		// The skip expression, if provided, allows the batch API executor to skip the task execution when the given expression evaluates to true. The expression is given as such a golang template that has to be evaluated to a final content true/false. The expression is an optional and in case not provided, the API will always be executed.
		SkipOnCondition      *string                                   `json:"SkipOnCondition,omitempty"`
		ErrorResponseHandler *WorkflowErrorResponseHandlerRelationship `json:"ErrorResponseHandler,omitempty"`
		TaskDefinition       *WorkflowTaskDefinitionRelationship       `json:"TaskDefinition,omitempty"`
	}

	varWorkflowBatchApiExecutorWithoutEmbeddedStruct := WorkflowBatchApiExecutorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowBatchApiExecutorWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowBatchApiExecutor := _WorkflowBatchApiExecutor{}
		varWorkflowBatchApiExecutor.ClassId = varWorkflowBatchApiExecutorWithoutEmbeddedStruct.ClassId
		varWorkflowBatchApiExecutor.ObjectType = varWorkflowBatchApiExecutorWithoutEmbeddedStruct.ObjectType
		varWorkflowBatchApiExecutor.Batch = varWorkflowBatchApiExecutorWithoutEmbeddedStruct.Batch
		varWorkflowBatchApiExecutor.Constraints = varWorkflowBatchApiExecutorWithoutEmbeddedStruct.Constraints
		varWorkflowBatchApiExecutor.Description = varWorkflowBatchApiExecutorWithoutEmbeddedStruct.Description
		varWorkflowBatchApiExecutor.Name = varWorkflowBatchApiExecutorWithoutEmbeddedStruct.Name
		varWorkflowBatchApiExecutor.Outcomes = varWorkflowBatchApiExecutorWithoutEmbeddedStruct.Outcomes
		varWorkflowBatchApiExecutor.Output = varWorkflowBatchApiExecutorWithoutEmbeddedStruct.Output
		varWorkflowBatchApiExecutor.RetryFromFailedApi = varWorkflowBatchApiExecutorWithoutEmbeddedStruct.RetryFromFailedApi
		varWorkflowBatchApiExecutor.SkipOnCondition = varWorkflowBatchApiExecutorWithoutEmbeddedStruct.SkipOnCondition
		varWorkflowBatchApiExecutor.ErrorResponseHandler = varWorkflowBatchApiExecutorWithoutEmbeddedStruct.ErrorResponseHandler
		varWorkflowBatchApiExecutor.TaskDefinition = varWorkflowBatchApiExecutorWithoutEmbeddedStruct.TaskDefinition
		*o = WorkflowBatchApiExecutor(varWorkflowBatchApiExecutor)
	} else {
		return err
	}

	varWorkflowBatchApiExecutor := _WorkflowBatchApiExecutor{}

	err = json.Unmarshal(bytes, &varWorkflowBatchApiExecutor)
	if err == nil {
		o.MoBaseMo = varWorkflowBatchApiExecutor.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Batch")
		delete(additionalProperties, "Constraints")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Outcomes")
		delete(additionalProperties, "Output")
		delete(additionalProperties, "RetryFromFailedApi")
		delete(additionalProperties, "SkipOnCondition")
		delete(additionalProperties, "ErrorResponseHandler")
		delete(additionalProperties, "TaskDefinition")

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

type NullableWorkflowBatchApiExecutor struct {
	value *WorkflowBatchApiExecutor
	isSet bool
}

func (v NullableWorkflowBatchApiExecutor) Get() *WorkflowBatchApiExecutor {
	return v.value
}

func (v *NullableWorkflowBatchApiExecutor) Set(val *WorkflowBatchApiExecutor) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowBatchApiExecutor) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowBatchApiExecutor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowBatchApiExecutor(val *WorkflowBatchApiExecutor) *NullableWorkflowBatchApiExecutor {
	return &NullableWorkflowBatchApiExecutor{value: val, isSet: true}
}

func (v NullableWorkflowBatchApiExecutor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowBatchApiExecutor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
