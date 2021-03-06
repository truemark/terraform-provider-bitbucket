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

// IssueChangeLinks struct for IssueChangeLinks
type IssueChangeLinks struct {
	Self *Link `json:"self,omitempty"`
	Issue *Link `json:"issue,omitempty"`
}

// NewIssueChangeLinks instantiates a new IssueChangeLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueChangeLinks() *IssueChangeLinks {
	this := IssueChangeLinks{}
	return &this
}

// NewIssueChangeLinksWithDefaults instantiates a new IssueChangeLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueChangeLinksWithDefaults() *IssueChangeLinks {
	this := IssueChangeLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *IssueChangeLinks) GetSelf() Link {
	if o == nil || o.Self == nil {
		var ret Link
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueChangeLinks) GetSelfOk() (*Link, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *IssueChangeLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Link and assigns it to the Self field.
func (o *IssueChangeLinks) SetSelf(v Link) {
	o.Self = &v
}

// GetIssue returns the Issue field value if set, zero value otherwise.
func (o *IssueChangeLinks) GetIssue() Link {
	if o == nil || o.Issue == nil {
		var ret Link
		return ret
	}
	return *o.Issue
}

// GetIssueOk returns a tuple with the Issue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueChangeLinks) GetIssueOk() (*Link, bool) {
	if o == nil || o.Issue == nil {
		return nil, false
	}
	return o.Issue, true
}

// HasIssue returns a boolean if a field has been set.
func (o *IssueChangeLinks) HasIssue() bool {
	if o != nil && o.Issue != nil {
		return true
	}

	return false
}

// SetIssue gets a reference to the given Link and assigns it to the Issue field.
func (o *IssueChangeLinks) SetIssue(v Link) {
	o.Issue = &v
}

func (o IssueChangeLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Issue != nil {
		toSerialize["issue"] = o.Issue
	}
	return json.Marshal(toSerialize)
}

type NullableIssueChangeLinks struct {
	value *IssueChangeLinks
	isSet bool
}

func (v NullableIssueChangeLinks) Get() *IssueChangeLinks {
	return v.value
}

func (v *NullableIssueChangeLinks) Set(val *IssueChangeLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueChangeLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueChangeLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueChangeLinks(val *IssueChangeLinks) *NullableIssueChangeLinks {
	return &NullableIssueChangeLinks{value: val, isSet: true}
}

func (v NullableIssueChangeLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueChangeLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


