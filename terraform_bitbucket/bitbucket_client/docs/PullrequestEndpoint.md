# PullrequestEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to [**Repository**](Repository.md) |  | [optional] 
**Branch** | Pointer to [**PullRequestBranch**](PullRequestBranch.md) |  | [optional] 
**Commit** | Pointer to [**PullRequestCommit**](PullRequestCommit.md) |  | [optional] 

## Methods

### NewPullrequestEndpoint

`func NewPullrequestEndpoint() *PullrequestEndpoint`

NewPullrequestEndpoint instantiates a new PullrequestEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullrequestEndpointWithDefaults

`func NewPullrequestEndpointWithDefaults() *PullrequestEndpoint`

NewPullrequestEndpointWithDefaults instantiates a new PullrequestEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *PullrequestEndpoint) GetRepository() Repository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PullrequestEndpoint) GetRepositoryOk() (*Repository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PullrequestEndpoint) SetRepository(v Repository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PullrequestEndpoint) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetBranch

`func (o *PullrequestEndpoint) GetBranch() PullRequestBranch`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *PullrequestEndpoint) GetBranchOk() (*PullRequestBranch, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *PullrequestEndpoint) SetBranch(v PullRequestBranch)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *PullrequestEndpoint) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCommit

`func (o *PullrequestEndpoint) GetCommit() PullRequestCommit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *PullrequestEndpoint) GetCommitOk() (*PullRequestCommit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *PullrequestEndpoint) SetCommit(v PullRequestCommit)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *PullrequestEndpoint) HasCommit() bool`

HasCommit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


