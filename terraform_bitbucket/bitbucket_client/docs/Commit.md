# Commit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to [**Repository**](Repository.md) |  | [optional] 
**Participants** | Pointer to [**[]Participant**](Participant.md) |  | [optional] 

## Methods

### NewCommit

`func NewCommit() *Commit`

NewCommit instantiates a new Commit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitWithDefaults

`func NewCommitWithDefaults() *Commit`

NewCommitWithDefaults instantiates a new Commit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *Commit) GetRepository() Repository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *Commit) GetRepositoryOk() (*Repository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *Commit) SetRepository(v Repository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *Commit) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetParticipants

`func (o *Commit) GetParticipants() []Participant`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *Commit) GetParticipantsOk() (*[]Participant, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *Commit) SetParticipants(v []Participant)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *Commit) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


