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

// PipelineStateCompletedStoppedAllOf A Bitbucket Pipelines STOPPED pipeline result.
type PipelineStateCompletedStoppedAllOf struct {
	// The name of the stopped result (STOPPED).
	Name *string `json:"name,omitempty"`
}

// NewPipelineStateCompletedStoppedAllOf instantiates a new PipelineStateCompletedStoppedAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStateCompletedStoppedAllOf() *PipelineStateCompletedStoppedAllOf {
	this := PipelineStateCompletedStoppedAllOf{}
	return &this
}

// NewPipelineStateCompletedStoppedAllOfWithDefaults instantiates a new PipelineStateCompletedStoppedAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStateCompletedStoppedAllOfWithDefaults() *PipelineStateCompletedStoppedAllOf {
	this := PipelineStateCompletedStoppedAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PipelineStateCompletedStoppedAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStateCompletedStoppedAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PipelineStateCompletedStoppedAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PipelineStateCompletedStoppedAllOf) SetName(v string) {
	o.Name = &v
}

func (o PipelineStateCompletedStoppedAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineStateCompletedStoppedAllOf struct {
	value *PipelineStateCompletedStoppedAllOf
	isSet bool
}

func (v NullablePipelineStateCompletedStoppedAllOf) Get() *PipelineStateCompletedStoppedAllOf {
	return v.value
}

func (v *NullablePipelineStateCompletedStoppedAllOf) Set(val *PipelineStateCompletedStoppedAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStateCompletedStoppedAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStateCompletedStoppedAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStateCompletedStoppedAllOf(val *PipelineStateCompletedStoppedAllOf) *NullablePipelineStateCompletedStoppedAllOf {
	return &NullablePipelineStateCompletedStoppedAllOf{value: val, isSet: true}
}

func (v NullablePipelineStateCompletedStoppedAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStateCompletedStoppedAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


