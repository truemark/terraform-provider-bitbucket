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

// PipelineScheduleExecutionExecutedAllOf A Pipelines executed schedule execution.
type PipelineScheduleExecutionExecutedAllOf struct {
	Pipeline *Pipeline `json:"pipeline,omitempty"`
}

// NewPipelineScheduleExecutionExecutedAllOf instantiates a new PipelineScheduleExecutionExecutedAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineScheduleExecutionExecutedAllOf() *PipelineScheduleExecutionExecutedAllOf {
	this := PipelineScheduleExecutionExecutedAllOf{}
	return &this
}

// NewPipelineScheduleExecutionExecutedAllOfWithDefaults instantiates a new PipelineScheduleExecutionExecutedAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineScheduleExecutionExecutedAllOfWithDefaults() *PipelineScheduleExecutionExecutedAllOf {
	this := PipelineScheduleExecutionExecutedAllOf{}
	return &this
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *PipelineScheduleExecutionExecutedAllOf) GetPipeline() Pipeline {
	if o == nil || o.Pipeline == nil {
		var ret Pipeline
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineScheduleExecutionExecutedAllOf) GetPipelineOk() (*Pipeline, bool) {
	if o == nil || o.Pipeline == nil {
		return nil, false
	}
	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *PipelineScheduleExecutionExecutedAllOf) HasPipeline() bool {
	if o != nil && o.Pipeline != nil {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given Pipeline and assigns it to the Pipeline field.
func (o *PipelineScheduleExecutionExecutedAllOf) SetPipeline(v Pipeline) {
	o.Pipeline = &v
}

func (o PipelineScheduleExecutionExecutedAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pipeline != nil {
		toSerialize["pipeline"] = o.Pipeline
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineScheduleExecutionExecutedAllOf struct {
	value *PipelineScheduleExecutionExecutedAllOf
	isSet bool
}

func (v NullablePipelineScheduleExecutionExecutedAllOf) Get() *PipelineScheduleExecutionExecutedAllOf {
	return v.value
}

func (v *NullablePipelineScheduleExecutionExecutedAllOf) Set(val *PipelineScheduleExecutionExecutedAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineScheduleExecutionExecutedAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineScheduleExecutionExecutedAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineScheduleExecutionExecutedAllOf(val *PipelineScheduleExecutionExecutedAllOf) *NullablePipelineScheduleExecutionExecutedAllOf {
	return &NullablePipelineScheduleExecutionExecutedAllOf{value: val, isSet: true}
}

func (v NullablePipelineScheduleExecutionExecutedAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineScheduleExecutionExecutedAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


