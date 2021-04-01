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
	"time"
)

// Comment The base type for all comments. This type should be considered abstract. Each of the \"commentable\" resources defines its own subtypes (e.g. `issue_comment`).
type Comment struct {
	Id *int32 `json:"id,omitempty"`
	CreatedOn *time.Time `json:"created_on,omitempty"`
	UpdatedOn *time.Time `json:"updated_on,omitempty"`
	Content *IssueContent `json:"content,omitempty"`
	User *User `json:"user,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
	Parent *Comment `json:"parent,omitempty"`
	Inline *CommentInline `json:"inline,omitempty"`
	Links *CommentLinks `json:"links,omitempty"`
}

// NewComment instantiates a new Comment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComment() *Comment {
	this := Comment{}
	return &this
}

// NewCommentWithDefaults instantiates a new Comment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentWithDefaults() *Comment {
	this := Comment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Comment) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Comment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Comment) SetId(v int32) {
	o.Id = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *Comment) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *Comment) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *Comment) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *Comment) GetUpdatedOn() time.Time {
	if o == nil || o.UpdatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetUpdatedOnOk() (*time.Time, bool) {
	if o == nil || o.UpdatedOn == nil {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *Comment) HasUpdatedOn() bool {
	if o != nil && o.UpdatedOn != nil {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given time.Time and assigns it to the UpdatedOn field.
func (o *Comment) SetUpdatedOn(v time.Time) {
	o.UpdatedOn = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Comment) GetContent() IssueContent {
	if o == nil || o.Content == nil {
		var ret IssueContent
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetContentOk() (*IssueContent, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Comment) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given IssueContent and assigns it to the Content field.
func (o *Comment) SetContent(v IssueContent) {
	o.Content = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Comment) GetUser() User {
	if o == nil || o.User == nil {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetUserOk() (*User, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Comment) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *Comment) SetUser(v User) {
	o.User = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Comment) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Comment) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *Comment) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *Comment) GetParent() Comment {
	if o == nil || o.Parent == nil {
		var ret Comment
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetParentOk() (*Comment, bool) {
	if o == nil || o.Parent == nil {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *Comment) HasParent() bool {
	if o != nil && o.Parent != nil {
		return true
	}

	return false
}

// SetParent gets a reference to the given Comment and assigns it to the Parent field.
func (o *Comment) SetParent(v Comment) {
	o.Parent = &v
}

// GetInline returns the Inline field value if set, zero value otherwise.
func (o *Comment) GetInline() CommentInline {
	if o == nil || o.Inline == nil {
		var ret CommentInline
		return ret
	}
	return *o.Inline
}

// GetInlineOk returns a tuple with the Inline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetInlineOk() (*CommentInline, bool) {
	if o == nil || o.Inline == nil {
		return nil, false
	}
	return o.Inline, true
}

// HasInline returns a boolean if a field has been set.
func (o *Comment) HasInline() bool {
	if o != nil && o.Inline != nil {
		return true
	}

	return false
}

// SetInline gets a reference to the given CommentInline and assigns it to the Inline field.
func (o *Comment) SetInline(v CommentInline) {
	o.Inline = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Comment) GetLinks() CommentLinks {
	if o == nil || o.Links == nil {
		var ret CommentLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetLinksOk() (*CommentLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Comment) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CommentLinks and assigns it to the Links field.
func (o *Comment) SetLinks(v CommentLinks) {
	o.Links = &v
}

func (o Comment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedOn != nil {
		toSerialize["created_on"] = o.CreatedOn
	}
	if o.UpdatedOn != nil {
		toSerialize["updated_on"] = o.UpdatedOn
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if o.Parent != nil {
		toSerialize["parent"] = o.Parent
	}
	if o.Inline != nil {
		toSerialize["inline"] = o.Inline
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableComment struct {
	value *Comment
	isSet bool
}

func (v NullableComment) Get() *Comment {
	return v.value
}

func (v *NullableComment) Set(val *Comment) {
	v.value = val
	v.isSet = true
}

func (v NullableComment) IsSet() bool {
	return v.isSet
}

func (v *NullableComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComment(val *Comment) *NullableComment {
	return &NullableComment{value: val, isSet: true}
}

func (v NullableComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


