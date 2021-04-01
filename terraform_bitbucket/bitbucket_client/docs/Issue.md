# Issue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**IssueLinks**](IssueLinks.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Repository** | Pointer to [**Repository**](Repository.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Reporter** | Pointer to [**User**](User.md) |  | [optional] 
**Assignee** | Pointer to [**User**](User.md) |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**UpdatedOn** | Pointer to **time.Time** |  | [optional] 
**EditedOn** | Pointer to **time.Time** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Milestone** | Pointer to [**Milestone**](Milestone.md) |  | [optional] 
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Component** | Pointer to [**Component**](Component.md) |  | [optional] 
**Votes** | Pointer to **int32** |  | [optional] 
**Content** | Pointer to [**IssueContent**](IssueContent.md) |  | [optional] 

## Methods

### NewIssue

`func NewIssue() *Issue`

NewIssue instantiates a new Issue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueWithDefaults

`func NewIssueWithDefaults() *Issue`

NewIssueWithDefaults instantiates a new Issue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Issue) GetLinks() IssueLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Issue) GetLinksOk() (*IssueLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Issue) SetLinks(v IssueLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Issue) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *Issue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Issue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Issue) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Issue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRepository

`func (o *Issue) GetRepository() Repository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *Issue) GetRepositoryOk() (*Repository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *Issue) SetRepository(v Repository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *Issue) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetTitle

`func (o *Issue) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Issue) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Issue) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Issue) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetReporter

`func (o *Issue) GetReporter() User`

GetReporter returns the Reporter field if non-nil, zero value otherwise.

### GetReporterOk

`func (o *Issue) GetReporterOk() (*User, bool)`

GetReporterOk returns a tuple with the Reporter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporter

`func (o *Issue) SetReporter(v User)`

SetReporter sets Reporter field to given value.

### HasReporter

`func (o *Issue) HasReporter() bool`

HasReporter returns a boolean if a field has been set.

### GetAssignee

`func (o *Issue) GetAssignee() User`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *Issue) GetAssigneeOk() (*User, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *Issue) SetAssignee(v User)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *Issue) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Issue) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Issue) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Issue) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Issue) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *Issue) GetUpdatedOn() time.Time`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *Issue) GetUpdatedOnOk() (*time.Time, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *Issue) SetUpdatedOn(v time.Time)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *Issue) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.

### GetEditedOn

`func (o *Issue) GetEditedOn() time.Time`

GetEditedOn returns the EditedOn field if non-nil, zero value otherwise.

### GetEditedOnOk

`func (o *Issue) GetEditedOnOk() (*time.Time, bool)`

GetEditedOnOk returns a tuple with the EditedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedOn

`func (o *Issue) SetEditedOn(v time.Time)`

SetEditedOn sets EditedOn field to given value.

### HasEditedOn

`func (o *Issue) HasEditedOn() bool`

HasEditedOn returns a boolean if a field has been set.

### GetState

`func (o *Issue) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Issue) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Issue) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Issue) HasState() bool`

HasState returns a boolean if a field has been set.

### GetKind

`func (o *Issue) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Issue) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Issue) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Issue) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPriority

`func (o *Issue) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Issue) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Issue) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Issue) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetMilestone

`func (o *Issue) GetMilestone() Milestone`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *Issue) GetMilestoneOk() (*Milestone, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *Issue) SetMilestone(v Milestone)`

SetMilestone sets Milestone field to given value.

### HasMilestone

`func (o *Issue) HasMilestone() bool`

HasMilestone returns a boolean if a field has been set.

### GetVersion

`func (o *Issue) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Issue) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Issue) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Issue) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetComponent

`func (o *Issue) GetComponent() Component`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *Issue) GetComponentOk() (*Component, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *Issue) SetComponent(v Component)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *Issue) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetVotes

`func (o *Issue) GetVotes() int32`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *Issue) GetVotesOk() (*int32, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *Issue) SetVotes(v int32)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *Issue) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### GetContent

`func (o *Issue) GetContent() IssueContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Issue) GetContentOk() (*IssueContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Issue) SetContent(v IssueContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *Issue) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


