# Pullrequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**PullRequestLinks**](PullRequestLinks.md) |  | [optional] 
**Id** | Pointer to **int32** | The pull request&#39;s unique ID. Note that pull request IDs are only unique within their associated repository. | [optional] 
**Title** | Pointer to **string** | Title of the pull request. | [optional] 
**Rendered** | Pointer to [**RenderedPullRequestMarkup**](RenderedPullRequestMarkup.md) |  | [optional] 
**Summary** | Pointer to [**IssueContent**](IssueContent.md) |  | [optional] 
**State** | Pointer to **string** | The pull request&#39;s current status. | [optional] 
**Author** | Pointer to [**Account**](Account.md) |  | [optional] 
**Source** | Pointer to [**PullrequestEndpoint**](PullrequestEndpoint.md) |  | [optional] 
**Destination** | Pointer to [**PullrequestEndpoint**](PullrequestEndpoint.md) |  | [optional] 
**MergeCommit** | Pointer to [**PullRequestCommit**](PullRequestCommit.md) |  | [optional] 
**CommentCount** | Pointer to **int32** | The number of comments for a specific pull request. | [optional] 
**TaskCount** | Pointer to **int32** | The number of open tasks for a specific pull request. | [optional] 
**CloseSourceBranch** | Pointer to **bool** | A boolean flag indicating if merging the pull request closes the source branch. | [optional] 
**ClosedBy** | Pointer to [**Account**](Account.md) |  | [optional] 
**Reason** | Pointer to **string** | Explains why a pull request was declined. This field is only applicable to pull requests in rejected state. | [optional] 
**CreatedOn** | Pointer to **time.Time** | The ISO8601 timestamp the request was created. | [optional] 
**UpdatedOn** | Pointer to **time.Time** | The ISO8601 timestamp the request was last updated. | [optional] 
**Reviewers** | Pointer to [**[]Account**](Account.md) | The list of users that were added as reviewers on this pull request when it was created. For performance reasons, the API only includes this list on a pull request&#39;s &#x60;self&#x60; URL. | [optional] 
**Participants** | Pointer to [**[]Participant**](Participant.md) |         The list of users that are collaborating on this pull request.         Collaborators are user that:          * are added to the pull request as a reviewer (part of the reviewers           list)         * are not explicit reviewers, but have commented on the pull request         * are not explicit reviewers, but have approved the pull request          Each user is wrapped in an object that indicates the user&#39;s role and         whether they have approved the pull request. For performance reasons,         the API only returns this list when an API requests a pull request by         id.          | [optional] 

## Methods

### NewPullrequest

`func NewPullrequest() *Pullrequest`

NewPullrequest instantiates a new Pullrequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullrequestWithDefaults

`func NewPullrequestWithDefaults() *Pullrequest`

NewPullrequestWithDefaults instantiates a new Pullrequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Pullrequest) GetLinks() PullRequestLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Pullrequest) GetLinksOk() (*PullRequestLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Pullrequest) SetLinks(v PullRequestLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Pullrequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *Pullrequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Pullrequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Pullrequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Pullrequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *Pullrequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Pullrequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Pullrequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Pullrequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetRendered

`func (o *Pullrequest) GetRendered() RenderedPullRequestMarkup`

GetRendered returns the Rendered field if non-nil, zero value otherwise.

### GetRenderedOk

`func (o *Pullrequest) GetRenderedOk() (*RenderedPullRequestMarkup, bool)`

GetRenderedOk returns a tuple with the Rendered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRendered

`func (o *Pullrequest) SetRendered(v RenderedPullRequestMarkup)`

SetRendered sets Rendered field to given value.

### HasRendered

`func (o *Pullrequest) HasRendered() bool`

HasRendered returns a boolean if a field has been set.

### GetSummary

`func (o *Pullrequest) GetSummary() IssueContent`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Pullrequest) GetSummaryOk() (*IssueContent, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Pullrequest) SetSummary(v IssueContent)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Pullrequest) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetState

`func (o *Pullrequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Pullrequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Pullrequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Pullrequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAuthor

`func (o *Pullrequest) GetAuthor() Account`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Pullrequest) GetAuthorOk() (*Account, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Pullrequest) SetAuthor(v Account)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Pullrequest) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetSource

`func (o *Pullrequest) GetSource() PullrequestEndpoint`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Pullrequest) GetSourceOk() (*PullrequestEndpoint, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Pullrequest) SetSource(v PullrequestEndpoint)`

SetSource sets Source field to given value.

### HasSource

`func (o *Pullrequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDestination

`func (o *Pullrequest) GetDestination() PullrequestEndpoint`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Pullrequest) GetDestinationOk() (*PullrequestEndpoint, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Pullrequest) SetDestination(v PullrequestEndpoint)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *Pullrequest) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetMergeCommit

`func (o *Pullrequest) GetMergeCommit() PullRequestCommit`

GetMergeCommit returns the MergeCommit field if non-nil, zero value otherwise.

### GetMergeCommitOk

`func (o *Pullrequest) GetMergeCommitOk() (*PullRequestCommit, bool)`

GetMergeCommitOk returns a tuple with the MergeCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeCommit

`func (o *Pullrequest) SetMergeCommit(v PullRequestCommit)`

SetMergeCommit sets MergeCommit field to given value.

### HasMergeCommit

`func (o *Pullrequest) HasMergeCommit() bool`

HasMergeCommit returns a boolean if a field has been set.

### GetCommentCount

`func (o *Pullrequest) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *Pullrequest) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *Pullrequest) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.

### HasCommentCount

`func (o *Pullrequest) HasCommentCount() bool`

HasCommentCount returns a boolean if a field has been set.

### GetTaskCount

`func (o *Pullrequest) GetTaskCount() int32`

GetTaskCount returns the TaskCount field if non-nil, zero value otherwise.

### GetTaskCountOk

`func (o *Pullrequest) GetTaskCountOk() (*int32, bool)`

GetTaskCountOk returns a tuple with the TaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCount

`func (o *Pullrequest) SetTaskCount(v int32)`

SetTaskCount sets TaskCount field to given value.

### HasTaskCount

`func (o *Pullrequest) HasTaskCount() bool`

HasTaskCount returns a boolean if a field has been set.

### GetCloseSourceBranch

`func (o *Pullrequest) GetCloseSourceBranch() bool`

GetCloseSourceBranch returns the CloseSourceBranch field if non-nil, zero value otherwise.

### GetCloseSourceBranchOk

`func (o *Pullrequest) GetCloseSourceBranchOk() (*bool, bool)`

GetCloseSourceBranchOk returns a tuple with the CloseSourceBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseSourceBranch

`func (o *Pullrequest) SetCloseSourceBranch(v bool)`

SetCloseSourceBranch sets CloseSourceBranch field to given value.

### HasCloseSourceBranch

`func (o *Pullrequest) HasCloseSourceBranch() bool`

HasCloseSourceBranch returns a boolean if a field has been set.

### GetClosedBy

`func (o *Pullrequest) GetClosedBy() Account`

GetClosedBy returns the ClosedBy field if non-nil, zero value otherwise.

### GetClosedByOk

`func (o *Pullrequest) GetClosedByOk() (*Account, bool)`

GetClosedByOk returns a tuple with the ClosedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedBy

`func (o *Pullrequest) SetClosedBy(v Account)`

SetClosedBy sets ClosedBy field to given value.

### HasClosedBy

`func (o *Pullrequest) HasClosedBy() bool`

HasClosedBy returns a boolean if a field has been set.

### GetReason

`func (o *Pullrequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Pullrequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Pullrequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Pullrequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Pullrequest) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Pullrequest) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Pullrequest) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Pullrequest) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *Pullrequest) GetUpdatedOn() time.Time`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *Pullrequest) GetUpdatedOnOk() (*time.Time, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *Pullrequest) SetUpdatedOn(v time.Time)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *Pullrequest) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.

### GetReviewers

`func (o *Pullrequest) GetReviewers() []Account`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *Pullrequest) GetReviewersOk() (*[]Account, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *Pullrequest) SetReviewers(v []Account)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *Pullrequest) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetParticipants

`func (o *Pullrequest) GetParticipants() []Participant`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *Pullrequest) GetParticipantsOk() (*[]Participant, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *Pullrequest) SetParticipants(v []Participant)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *Pullrequest) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


