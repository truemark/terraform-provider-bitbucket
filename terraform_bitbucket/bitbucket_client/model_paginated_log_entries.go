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

// PaginatedLogEntries A paginated list of issue changes.
type PaginatedLogEntries struct {
	// Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute.
	Size *int32 `json:"size,omitempty"`
	// Page number of the current results. This is an optional element that is not provided in all responses.
	Page *int32 `json:"page,omitempty"`
	// Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values.
	Pagelen *int32 `json:"pagelen,omitempty"`
	// Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs.
	Next *string `json:"next,omitempty"`
	// Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs.
	Previous *string `json:"previous,omitempty"`
	Values *[]IssueChange `json:"values,omitempty"`
}

// NewPaginatedLogEntries instantiates a new PaginatedLogEntries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedLogEntries() *PaginatedLogEntries {
	this := PaginatedLogEntries{}
	return &this
}

// NewPaginatedLogEntriesWithDefaults instantiates a new PaginatedLogEntries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedLogEntriesWithDefaults() *PaginatedLogEntries {
	this := PaginatedLogEntries{}
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *PaginatedLogEntries) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedLogEntries) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *PaginatedLogEntries) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *PaginatedLogEntries) SetSize(v int32) {
	o.Size = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PaginatedLogEntries) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedLogEntries) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PaginatedLogEntries) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *PaginatedLogEntries) SetPage(v int32) {
	o.Page = &v
}

// GetPagelen returns the Pagelen field value if set, zero value otherwise.
func (o *PaginatedLogEntries) GetPagelen() int32 {
	if o == nil || o.Pagelen == nil {
		var ret int32
		return ret
	}
	return *o.Pagelen
}

// GetPagelenOk returns a tuple with the Pagelen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedLogEntries) GetPagelenOk() (*int32, bool) {
	if o == nil || o.Pagelen == nil {
		return nil, false
	}
	return o.Pagelen, true
}

// HasPagelen returns a boolean if a field has been set.
func (o *PaginatedLogEntries) HasPagelen() bool {
	if o != nil && o.Pagelen != nil {
		return true
	}

	return false
}

// SetPagelen gets a reference to the given int32 and assigns it to the Pagelen field.
func (o *PaginatedLogEntries) SetPagelen(v int32) {
	o.Pagelen = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PaginatedLogEntries) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedLogEntries) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedLogEntries) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *PaginatedLogEntries) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *PaginatedLogEntries) GetPrevious() string {
	if o == nil || o.Previous == nil {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedLogEntries) GetPreviousOk() (*string, bool) {
	if o == nil || o.Previous == nil {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedLogEntries) HasPrevious() bool {
	if o != nil && o.Previous != nil {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *PaginatedLogEntries) SetPrevious(v string) {
	o.Previous = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *PaginatedLogEntries) GetValues() []IssueChange {
	if o == nil || o.Values == nil {
		var ret []IssueChange
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedLogEntries) GetValuesOk() (*[]IssueChange, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *PaginatedLogEntries) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []IssueChange and assigns it to the Values field.
func (o *PaginatedLogEntries) SetValues(v []IssueChange) {
	o.Values = &v
}

func (o PaginatedLogEntries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Pagelen != nil {
		toSerialize["pagelen"] = o.Pagelen
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullablePaginatedLogEntries struct {
	value *PaginatedLogEntries
	isSet bool
}

func (v NullablePaginatedLogEntries) Get() *PaginatedLogEntries {
	return v.value
}

func (v *NullablePaginatedLogEntries) Set(val *PaginatedLogEntries) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedLogEntries) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedLogEntries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedLogEntries(val *PaginatedLogEntries) *NullablePaginatedLogEntries {
	return &NullablePaginatedLogEntries{value: val, isSet: true}
}

func (v NullablePaginatedLogEntries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedLogEntries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


