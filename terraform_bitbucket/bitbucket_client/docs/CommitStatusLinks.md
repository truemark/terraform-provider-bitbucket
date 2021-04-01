# CommitStatusLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to [**Link**](Link.md) |  | [optional] 
**Commit** | Pointer to [**Link**](Link.md) |  | [optional] 

## Methods

### NewCommitStatusLinks

`func NewCommitStatusLinks() *CommitStatusLinks`

NewCommitStatusLinks instantiates a new CommitStatusLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitStatusLinksWithDefaults

`func NewCommitStatusLinksWithDefaults() *CommitStatusLinks`

NewCommitStatusLinksWithDefaults instantiates a new CommitStatusLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *CommitStatusLinks) GetSelf() Link`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *CommitStatusLinks) GetSelfOk() (*Link, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *CommitStatusLinks) SetSelf(v Link)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *CommitStatusLinks) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetCommit

`func (o *CommitStatusLinks) GetCommit() Link`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *CommitStatusLinks) GetCommitOk() (*Link, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *CommitStatusLinks) SetCommit(v Link)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *CommitStatusLinks) HasCommit() bool`

HasCommit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


