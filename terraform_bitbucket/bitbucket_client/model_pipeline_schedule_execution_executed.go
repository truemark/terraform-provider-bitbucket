/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bitbucket_client

import (
	"encoding/json"
)

// PipelineScheduleExecutionExecuted struct for PipelineScheduleExecutionExecuted
type PipelineScheduleExecutionExecuted struct {
	PipelineScheduleExecution
	Pipeline *Pipeline `json:"pipeline,omitempty"`
}

// NewPipelineScheduleExecutionExecuted instantiates a new PipelineScheduleExecutionExecuted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineScheduleExecutionExecuted(type_ string) *PipelineScheduleExecutionExecuted {
	this := PipelineScheduleExecutionExecuted{}
	this.Type = type_
	return &this
}

// NewPipelineScheduleExecutionExecutedWithDefaults instantiates a new PipelineScheduleExecutionExecuted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineScheduleExecutionExecutedWithDefaults() *PipelineScheduleExecutionExecuted {
	this := PipelineScheduleExecutionExecuted{}
	return &this
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *PipelineScheduleExecutionExecuted) GetPipeline() Pipeline {
	if o == nil || o.Pipeline == nil {
		var ret Pipeline
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineScheduleExecutionExecuted) GetPipelineOk() (*Pipeline, bool) {
	if o == nil || o.Pipeline == nil {
		return nil, false
	}
	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *PipelineScheduleExecutionExecuted) HasPipeline() bool {
	if o != nil && o.Pipeline != nil {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given Pipeline and assigns it to the Pipeline field.
func (o *PipelineScheduleExecutionExecuted) SetPipeline(v Pipeline) {
	o.Pipeline = &v
}

func (o PipelineScheduleExecutionExecuted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPipelineScheduleExecution, errPipelineScheduleExecution := json.Marshal(o.PipelineScheduleExecution)
	if errPipelineScheduleExecution != nil {
		return []byte{}, errPipelineScheduleExecution
	}
	errPipelineScheduleExecution = json.Unmarshal([]byte(serializedPipelineScheduleExecution), &toSerialize)
	if errPipelineScheduleExecution != nil {
		return []byte{}, errPipelineScheduleExecution
	}
	if o.Pipeline != nil {
		toSerialize["pipeline"] = o.Pipeline
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineScheduleExecutionExecuted struct {
	value *PipelineScheduleExecutionExecuted
	isSet bool
}

func (v NullablePipelineScheduleExecutionExecuted) Get() *PipelineScheduleExecutionExecuted {
	return v.value
}

func (v *NullablePipelineScheduleExecutionExecuted) Set(val *PipelineScheduleExecutionExecuted) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineScheduleExecutionExecuted) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineScheduleExecutionExecuted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineScheduleExecutionExecuted(val *PipelineScheduleExecutionExecuted) *NullablePipelineScheduleExecutionExecuted {
	return &NullablePipelineScheduleExecutionExecuted{value: val, isSet: true}
}

func (v NullablePipelineScheduleExecutionExecuted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineScheduleExecutionExecuted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


