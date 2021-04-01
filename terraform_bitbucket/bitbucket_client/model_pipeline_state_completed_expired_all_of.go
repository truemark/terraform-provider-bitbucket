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

// PipelineStateCompletedExpiredAllOf A Bitbucket Pipelines EXPIRED pipeline result.
type PipelineStateCompletedExpiredAllOf struct {
	// The name of the stopped result (EXPIRED).
	Name *string `json:"name,omitempty"`
}

// NewPipelineStateCompletedExpiredAllOf instantiates a new PipelineStateCompletedExpiredAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStateCompletedExpiredAllOf() *PipelineStateCompletedExpiredAllOf {
	this := PipelineStateCompletedExpiredAllOf{}
	return &this
}

// NewPipelineStateCompletedExpiredAllOfWithDefaults instantiates a new PipelineStateCompletedExpiredAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStateCompletedExpiredAllOfWithDefaults() *PipelineStateCompletedExpiredAllOf {
	this := PipelineStateCompletedExpiredAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PipelineStateCompletedExpiredAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStateCompletedExpiredAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PipelineStateCompletedExpiredAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PipelineStateCompletedExpiredAllOf) SetName(v string) {
	o.Name = &v
}

func (o PipelineStateCompletedExpiredAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineStateCompletedExpiredAllOf struct {
	value *PipelineStateCompletedExpiredAllOf
	isSet bool
}

func (v NullablePipelineStateCompletedExpiredAllOf) Get() *PipelineStateCompletedExpiredAllOf {
	return v.value
}

func (v *NullablePipelineStateCompletedExpiredAllOf) Set(val *PipelineStateCompletedExpiredAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStateCompletedExpiredAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStateCompletedExpiredAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStateCompletedExpiredAllOf(val *PipelineStateCompletedExpiredAllOf) *NullablePipelineStateCompletedExpiredAllOf {
	return &NullablePipelineStateCompletedExpiredAllOf{value: val, isSet: true}
}

func (v NullablePipelineStateCompletedExpiredAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStateCompletedExpiredAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


