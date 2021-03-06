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

// PipelineBuildNumber struct for PipelineBuildNumber
type PipelineBuildNumber struct {
	Object
	// The next number that will be used as build number.
	Next *int32 `json:"next,omitempty"`
}

// NewPipelineBuildNumber instantiates a new PipelineBuildNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineBuildNumber(type_ string) *PipelineBuildNumber {
	this := PipelineBuildNumber{}
	this.Type = type_
	return &this
}

// NewPipelineBuildNumberWithDefaults instantiates a new PipelineBuildNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineBuildNumberWithDefaults() *PipelineBuildNumber {
	this := PipelineBuildNumber{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PipelineBuildNumber) GetNext() int32 {
	if o == nil || o.Next == nil {
		var ret int32
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineBuildNumber) GetNextOk() (*int32, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PipelineBuildNumber) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given int32 and assigns it to the Next field.
func (o *PipelineBuildNumber) SetNext(v int32) {
	o.Next = &v
}

func (o PipelineBuildNumber) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedObject, errObject := json.Marshal(o.Object)
	if errObject != nil {
		return []byte{}, errObject
	}
	errObject = json.Unmarshal([]byte(serializedObject), &toSerialize)
	if errObject != nil {
		return []byte{}, errObject
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineBuildNumber struct {
	value *PipelineBuildNumber
	isSet bool
}

func (v NullablePipelineBuildNumber) Get() *PipelineBuildNumber {
	return v.value
}

func (v *NullablePipelineBuildNumber) Set(val *PipelineBuildNumber) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineBuildNumber) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineBuildNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineBuildNumber(val *PipelineBuildNumber) *NullablePipelineBuildNumber {
	return &NullablePipelineBuildNumber{value: val, isSet: true}
}

func (v NullablePipelineBuildNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineBuildNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


