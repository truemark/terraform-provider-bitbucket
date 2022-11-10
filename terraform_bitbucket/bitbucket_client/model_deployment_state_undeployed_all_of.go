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

// DeploymentStateUndeployedAllOf A Bitbucket Deployment UNDEPLOYED deployment state.
type DeploymentStateUndeployedAllOf struct {
	// The name of deployment state (UNDEPLOYED).
	Name *string `json:"name,omitempty"`
	// Link to trigger the deployment.
	TriggerUrl *string `json:"trigger_url,omitempty"`
}

// NewDeploymentStateUndeployedAllOf instantiates a new DeploymentStateUndeployedAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentStateUndeployedAllOf() *DeploymentStateUndeployedAllOf {
	this := DeploymentStateUndeployedAllOf{}
	return &this
}

// NewDeploymentStateUndeployedAllOfWithDefaults instantiates a new DeploymentStateUndeployedAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentStateUndeployedAllOfWithDefaults() *DeploymentStateUndeployedAllOf {
	this := DeploymentStateUndeployedAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeploymentStateUndeployedAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateUndeployedAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeploymentStateUndeployedAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeploymentStateUndeployedAllOf) SetName(v string) {
	o.Name = &v
}

// GetTriggerUrl returns the TriggerUrl field value if set, zero value otherwise.
func (o *DeploymentStateUndeployedAllOf) GetTriggerUrl() string {
	if o == nil || o.TriggerUrl == nil {
		var ret string
		return ret
	}
	return *o.TriggerUrl
}

// GetTriggerUrlOk returns a tuple with the TriggerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentStateUndeployedAllOf) GetTriggerUrlOk() (*string, bool) {
	if o == nil || o.TriggerUrl == nil {
		return nil, false
	}
	return o.TriggerUrl, true
}

// HasTriggerUrl returns a boolean if a field has been set.
func (o *DeploymentStateUndeployedAllOf) HasTriggerUrl() bool {
	if o != nil && o.TriggerUrl != nil {
		return true
	}

	return false
}

// SetTriggerUrl gets a reference to the given string and assigns it to the TriggerUrl field.
func (o *DeploymentStateUndeployedAllOf) SetTriggerUrl(v string) {
	o.TriggerUrl = &v
}

func (o DeploymentStateUndeployedAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TriggerUrl != nil {
		toSerialize["trigger_url"] = o.TriggerUrl
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentStateUndeployedAllOf struct {
	value *DeploymentStateUndeployedAllOf
	isSet bool
}

func (v NullableDeploymentStateUndeployedAllOf) Get() *DeploymentStateUndeployedAllOf {
	return v.value
}

func (v *NullableDeploymentStateUndeployedAllOf) Set(val *DeploymentStateUndeployedAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentStateUndeployedAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentStateUndeployedAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentStateUndeployedAllOf(val *DeploymentStateUndeployedAllOf) *NullableDeploymentStateUndeployedAllOf {
	return &NullableDeploymentStateUndeployedAllOf{value: val, isSet: true}
}

func (v NullableDeploymentStateUndeployedAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentStateUndeployedAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

