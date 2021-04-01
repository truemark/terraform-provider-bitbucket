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

// Pullrequest struct for Pullrequest
type Pullrequest struct {
	Object
	Links *PullRequestLinks `json:"links,omitempty"`
	// The pull request's unique ID. Note that pull request IDs are only unique within their associated repository.
	Id *int32 `json:"id,omitempty"`
	// Title of the pull request.
	Title *string `json:"title,omitempty"`
	Rendered *RenderedPullRequestMarkup `json:"rendered,omitempty"`
	Summary *IssueContent `json:"summary,omitempty"`
	// The pull request's current status.
	State *string `json:"state,omitempty"`
	Author *Account `json:"author,omitempty"`
	Source *PullrequestEndpoint `json:"source,omitempty"`
	Destination *PullrequestEndpoint `json:"destination,omitempty"`
	MergeCommit *PullRequestCommit `json:"merge_commit,omitempty"`
	// The number of comments for a specific pull request.
	CommentCount *int32 `json:"comment_count,omitempty"`
	// The number of open tasks for a specific pull request.
	TaskCount *int32 `json:"task_count,omitempty"`
	// A boolean flag indicating if merging the pull request closes the source branch.
	CloseSourceBranch *bool `json:"close_source_branch,omitempty"`
	ClosedBy *Account `json:"closed_by,omitempty"`
	// Explains why a pull request was declined. This field is only applicable to pull requests in rejected state.
	Reason *string `json:"reason,omitempty"`
	// The ISO8601 timestamp the request was created.
	CreatedOn *time.Time `json:"created_on,omitempty"`
	// The ISO8601 timestamp the request was last updated.
	UpdatedOn *time.Time `json:"updated_on,omitempty"`
	// The list of users that were added as reviewers on this pull request when it was created. For performance reasons, the API only includes this list on a pull request's `self` URL.
	Reviewers *[]Account `json:"reviewers,omitempty"`
	//         The list of users that are collaborating on this pull request.         Collaborators are user that:          * are added to the pull request as a reviewer (part of the reviewers           list)         * are not explicit reviewers, but have commented on the pull request         * are not explicit reviewers, but have approved the pull request          Each user is wrapped in an object that indicates the user's role and         whether they have approved the pull request. For performance reasons,         the API only returns this list when an API requests a pull request by         id.         
	Participants *[]Participant `json:"participants,omitempty"`
}

// NewPullrequest instantiates a new Pullrequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPullrequest(type_ string) *Pullrequest {
	this := Pullrequest{}
	this.Type = type_
	return &this
}

// NewPullrequestWithDefaults instantiates a new Pullrequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPullrequestWithDefaults() *Pullrequest {
	this := Pullrequest{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Pullrequest) GetLinks() PullRequestLinks {
	if o == nil || o.Links == nil {
		var ret PullRequestLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pullrequest) GetLinksOk() (*PullRequestLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Pullrequest) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PullRequestLinks and assigns it to the Links field.
func (o *Pullrequest) SetLinks(v PullRequestLinks) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Pullrequest) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pullrequest) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Pullrequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Pullrequest) SetId(v int32) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Pullrequest) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pullrequest) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Pullrequest) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Pullrequest) SetTitle(v string) {
	o.Title = &v
}

// GetRendered returns the Rendered field value if set, zero value otherwise.
func (o *Pullrequest) GetRendered() RenderedPullRequestMarkup {
	if o == nil || o.Rendered == nil {
		var ret RenderedPullRequestMarkup
		return ret
	}
	return *o.Rendered
}

// GetRenderedOk returns a tuple with the Rendered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pullrequest) GetRenderedOk() (*RenderedPullRequestMarkup, bool) {
	if o == nil || o.Rendered == nil {
		return nil, false
	}
	return o.Rendered, true
}

// HasRendered returns a boolean if a field has been set.
func (o *Pullrequest) HasRendered() bool {
	if o != nil && o.Rendered != nil {
		return true
	}

	return false
}

// SetRendered gets a reference to the given RenderedPullRequestMarkup and assigns it to the Rendered field.
func (o *Pullrequest) SetRendered(v RenderedPullRequestMarkup) {
	o.Rendered = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *Pullrequest) GetSummary() IssueContent {
	if o == nil || o.Summary == nil {
		var ret IssueContent
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pullrequest) GetSummaryOk() (*IssueContent, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *Pullrequest) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given IssueContent and assigns it to the Summary field.
func (o *Pullrequest) SetSummary(v IssueContent) {
	o.Summary = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Pullrequest) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pullrequest) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Pullrequest) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Pullrequest) SetState(v string) {
	o.State = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *Pullrequest) GetAuthor() Account {
	if o == nil || o.Author == nil {
		var ret Account
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pullrequest) GetAuthorOk() (*Account, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *Pullrequest) HasAuthor() bool {
	if o != nil && o.Author != nil {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given Account and assigns it to the Author field.
func (o *Pullrequest) SetAuthor(v Account) {
	o.Author = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Pullrequest) GetSource() PullrequestEndpoint {
	if o == nil || o.Source == nil {
		var ret PullrequestEndpoint
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pullrequest) GetSourceOk() (*PullrequestEndpoint, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Pullrequest) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given PullrequestEndpoint and assigns it to the Source field.
func (o *Pullrequest) SetSource(v PullrequestEndpoint) {
	o.Source = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *Pullrequest) GetDestination() PullrequestEndpoint {
	if o == nil || o.Destination == nil {
		var ret PullrequestEndpoint
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pullrequest) GetDestinationOk() (*PullrequestEndpoint, bool) {
	if o == nil || o.Destination == nil {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *Pullrequest) HasDestination() bool {
	if o != nil && o.Destination != nil {
		return true
	}

	return false
}

// SetDestination gets a reference to the given PullrequestEndpoint and assigns it to the Destination field.
func (o *Pullrequest) SetDestination(v PullrequestEndpoint) {
	o.Destination = &v
}

// GetMergeCommit returns the MergeCommit field value if set, zero value otherwise.
func (o *Pullrequest) GetMergeCommit() PullRequestCommit {
	if o == nil || o.MergeCommit == nil {
		var ret PullRequestCommit
		return ret
	}
	return *o.MergeCommit
}

// GetMergeCommitOk returns a tuple with the MergeCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pullrequest) GetMergeCommitOk() (*PullRequestCommit, bool) {
	if o == nil || o.MergeCommit == nil {
		return nil, false
	}
	return o.MergeCommit, true
}

// HasMergeCommit returns a boolean if a field has been set.
func (o *Pullrequest) HasMergeCommit() bool {
	if o != nil && o.MergeCommit != nil {
		return true
	}

	return false
}

// SetMergeCommit gets a reference to the given PullRequestCommit and assigns it to the MergeCommit field.
func (o *Pullrequest) SetMergeCommit(v PullRequestCommit) {
	o.MergeCommit = &v
}

// GetCommentCount returns the CommentCount field value if set, zero value otherwise.
func (o *Pullrequest) GetCommentCount() int32 {
	if o == nil || o.CommentCount == nil {
		var ret int32
		return ret
	}
	return *o.CommentCount
}

// GetCommentCountOk returns a tuple with the CommentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pullrequest) GetCommentCountOk() (*int32, bool) {
	if o == nil || o.CommentCount == nil {
		return nil, false
	}
	return o.CommentCount, true
}

// HasCommentCount returns a boolean if a field has been set.
func (o *Pullrequest) HasCommentCount() bool {
	if o != nil && o.CommentCount != nil {
		return true
	}

	return false
}

// SetCommentCount gets a reference to the given int32 and assigns it to the CommentCount field.
func (o *Pullrequest) SetCommentCount(v int32) {
	o.CommentCount = &v
}

// GetTaskCount returns the TaskCount field value if set, zero value otherwise.
func (o *Pullrequest) GetTaskCount() int32 {
	if o == nil || o.TaskCount == nil {
		var ret int32
		return ret
	}
	return *o.TaskCount
}

// GetTaskCountOk returns a tuple with the TaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pullrequest) GetTaskCountOk() (*int32, bool) {
	if o == nil || o.TaskCount == nil {
		return nil, false
	}
	return o.TaskCount, true
}

// HasTaskCount returns a boolean if a field has been set.
func (o *Pullrequest) HasTaskCount() bool {
	if o != nil && o.TaskCount != nil {
		return true
	}

	return false
}

// SetTaskCount gets a reference to the given int32 and assigns it to the TaskCount field.
func (o *Pullrequest) SetTaskCount(v int32) {
	o.TaskCount = &v
}

// GetCloseSourceBranch returns the CloseSourceBranch field value if set, zero value otherwise.
func (o *Pullrequest) GetCloseSourceBranch() bool {
	if o == nil || o.CloseSourceBranch == nil {
		var ret bool
		return ret
	}
	return *o.CloseSourceBranch
}

// GetCloseSourceBranchOk returns a tuple with the CloseSourceBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pullrequest) GetCloseSourceBranchOk() (*bool, bool) {
	if o == nil || o.CloseSourceBranch == nil {
		return nil, false
	}
	return o.CloseSourceBranch, true
}

// HasCloseSourceBranch returns a boolean if a field has been set.
func (o *Pullrequest) HasCloseSourceBranch() bool {
	if o != nil && o.CloseSourceBranch != nil {
		return true
	}

	return false
}

// SetCloseSourceBranch gets a reference to the given bool and assigns it to the CloseSourceBranch field.
func (o *Pullrequest) SetCloseSourceBranch(v bool) {
	o.CloseSourceBranch = &v
}

// GetClosedBy returns the ClosedBy field value if set, zero value otherwise.
func (o *Pullrequest) GetClosedBy() Account {
	if o == nil || o.ClosedBy == nil {
		var ret Account
		return ret
	}
	return *o.ClosedBy
}

// GetClosedByOk returns a tuple with the ClosedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pullrequest) GetClosedByOk() (*Account, bool) {
	if o == nil || o.ClosedBy == nil {
		return nil, false
	}
	return o.ClosedBy, true
}

// HasClosedBy returns a boolean if a field has been set.
func (o *Pullrequest) HasClosedBy() bool {
	if o != nil && o.ClosedBy != nil {
		return true
	}

	return false
}

// SetClosedBy gets a reference to the given Account and assigns it to the ClosedBy field.
func (o *Pullrequest) SetClosedBy(v Account) {
	o.ClosedBy = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *Pullrequest) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pullrequest) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *Pullrequest) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *Pullrequest) SetReason(v string) {
	o.Reason = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *Pullrequest) GetCreatedOn() time.Time {
	if o == nil || o.CreatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pullrequest) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || o.CreatedOn == nil {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *Pullrequest) HasCreatedOn() bool {
	if o != nil && o.CreatedOn != nil {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *Pullrequest) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetUpdatedOn returns the UpdatedOn field value if set, zero value otherwise.
func (o *Pullrequest) GetUpdatedOn() time.Time {
	if o == nil || o.UpdatedOn == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedOn
}

// GetUpdatedOnOk returns a tuple with the UpdatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pullrequest) GetUpdatedOnOk() (*time.Time, bool) {
	if o == nil || o.UpdatedOn == nil {
		return nil, false
	}
	return o.UpdatedOn, true
}

// HasUpdatedOn returns a boolean if a field has been set.
func (o *Pullrequest) HasUpdatedOn() bool {
	if o != nil && o.UpdatedOn != nil {
		return true
	}

	return false
}

// SetUpdatedOn gets a reference to the given time.Time and assigns it to the UpdatedOn field.
func (o *Pullrequest) SetUpdatedOn(v time.Time) {
	o.UpdatedOn = &v
}

// GetReviewers returns the Reviewers field value if set, zero value otherwise.
func (o *Pullrequest) GetReviewers() []Account {
	if o == nil || o.Reviewers == nil {
		var ret []Account
		return ret
	}
	return *o.Reviewers
}

// GetReviewersOk returns a tuple with the Reviewers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pullrequest) GetReviewersOk() (*[]Account, bool) {
	if o == nil || o.Reviewers == nil {
		return nil, false
	}
	return o.Reviewers, true
}

// HasReviewers returns a boolean if a field has been set.
func (o *Pullrequest) HasReviewers() bool {
	if o != nil && o.Reviewers != nil {
		return true
	}

	return false
}

// SetReviewers gets a reference to the given []Account and assigns it to the Reviewers field.
func (o *Pullrequest) SetReviewers(v []Account) {
	o.Reviewers = &v
}

// GetParticipants returns the Participants field value if set, zero value otherwise.
func (o *Pullrequest) GetParticipants() []Participant {
	if o == nil || o.Participants == nil {
		var ret []Participant
		return ret
	}
	return *o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pullrequest) GetParticipantsOk() (*[]Participant, bool) {
	if o == nil || o.Participants == nil {
		return nil, false
	}
	return o.Participants, true
}

// HasParticipants returns a boolean if a field has been set.
func (o *Pullrequest) HasParticipants() bool {
	if o != nil && o.Participants != nil {
		return true
	}

	return false
}

// SetParticipants gets a reference to the given []Participant and assigns it to the Participants field.
func (o *Pullrequest) SetParticipants(v []Participant) {
	o.Participants = &v
}

func (o Pullrequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedObject, errObject := json.Marshal(o.Object)
	if errObject != nil {
		return []byte{}, errObject
	}
	errObject = json.Unmarshal([]byte(serializedObject), &toSerialize)
	if errObject != nil {
		return []byte{}, errObject
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Rendered != nil {
		toSerialize["rendered"] = o.Rendered
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Destination != nil {
		toSerialize["destination"] = o.Destination
	}
	if o.MergeCommit != nil {
		toSerialize["merge_commit"] = o.MergeCommit
	}
	if o.CommentCount != nil {
		toSerialize["comment_count"] = o.CommentCount
	}
	if o.TaskCount != nil {
		toSerialize["task_count"] = o.TaskCount
	}
	if o.CloseSourceBranch != nil {
		toSerialize["close_source_branch"] = o.CloseSourceBranch
	}
	if o.ClosedBy != nil {
		toSerialize["closed_by"] = o.ClosedBy
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.CreatedOn != nil {
		toSerialize["created_on"] = o.CreatedOn
	}
	if o.UpdatedOn != nil {
		toSerialize["updated_on"] = o.UpdatedOn
	}
	if o.Reviewers != nil {
		toSerialize["reviewers"] = o.Reviewers
	}
	if o.Participants != nil {
		toSerialize["participants"] = o.Participants
	}
	return json.Marshal(toSerialize)
}

type NullablePullrequest struct {
	value *Pullrequest
	isSet bool
}

func (v NullablePullrequest) Get() *Pullrequest {
	return v.value
}

func (v *NullablePullrequest) Set(val *Pullrequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePullrequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePullrequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePullrequest(val *Pullrequest) *NullablePullrequest {
	return &NullablePullrequest{value: val, isSet: true}
}

func (v NullablePullrequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePullrequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


