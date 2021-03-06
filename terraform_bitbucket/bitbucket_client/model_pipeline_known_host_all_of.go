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

// PipelineKnownHostAllOf A Pipelines known host.
type PipelineKnownHostAllOf struct {
	// The UUID identifying the known host.
	Uuid *string `json:"uuid,omitempty"`
	// The hostname of the known host.
	Hostname *string `json:"hostname,omitempty"`
	PublicKey *PipelineSshPublicKey `json:"public_key,omitempty"`
}

// NewPipelineKnownHostAllOf instantiates a new PipelineKnownHostAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineKnownHostAllOf() *PipelineKnownHostAllOf {
	this := PipelineKnownHostAllOf{}
	return &this
}

// NewPipelineKnownHostAllOfWithDefaults instantiates a new PipelineKnownHostAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineKnownHostAllOfWithDefaults() *PipelineKnownHostAllOf {
	this := PipelineKnownHostAllOf{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *PipelineKnownHostAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineKnownHostAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *PipelineKnownHostAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *PipelineKnownHostAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *PipelineKnownHostAllOf) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineKnownHostAllOf) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *PipelineKnownHostAllOf) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *PipelineKnownHostAllOf) SetHostname(v string) {
	o.Hostname = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *PipelineKnownHostAllOf) GetPublicKey() PipelineSshPublicKey {
	if o == nil || o.PublicKey == nil {
		var ret PipelineSshPublicKey
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineKnownHostAllOf) GetPublicKeyOk() (*PipelineSshPublicKey, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *PipelineKnownHostAllOf) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given PipelineSshPublicKey and assigns it to the PublicKey field.
func (o *PipelineKnownHostAllOf) SetPublicKey(v PipelineSshPublicKey) {
	o.PublicKey = &v
}

func (o PipelineKnownHostAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Uuid != nil {
		toSerialize["uuid"] = o.Uuid
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.PublicKey != nil {
		toSerialize["public_key"] = o.PublicKey
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineKnownHostAllOf struct {
	value *PipelineKnownHostAllOf
	isSet bool
}

func (v NullablePipelineKnownHostAllOf) Get() *PipelineKnownHostAllOf {
	return v.value
}

func (v *NullablePipelineKnownHostAllOf) Set(val *PipelineKnownHostAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineKnownHostAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineKnownHostAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineKnownHostAllOf(val *PipelineKnownHostAllOf) *NullablePipelineKnownHostAllOf {
	return &NullablePipelineKnownHostAllOf{value: val, isSet: true}
}

func (v NullablePipelineKnownHostAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineKnownHostAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


