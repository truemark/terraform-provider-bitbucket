# Participant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**ParticipatedOn** | Pointer to **time.Time** | The ISO8601 timestamp of the participant&#39;s action. For approvers, this is the time of their approval. For commenters and pull request reviewers who are not approvers, this is the time they last commented, or null if they have not commented. | [optional] 

## Methods

### NewParticipant

`func NewParticipant() *Participant`

NewParticipant instantiates a new Participant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantWithDefaults

`func NewParticipantWithDefaults() *Participant`

NewParticipantWithDefaults instantiates a new Participant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *Participant) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Participant) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Participant) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *Participant) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRole

`func (o *Participant) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Participant) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Participant) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Participant) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetApproved

`func (o *Participant) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *Participant) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *Participant) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *Participant) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetState

`func (o *Participant) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Participant) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Participant) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Participant) HasState() bool`

HasState returns a boolean if a field has been set.

### GetParticipatedOn

`func (o *Participant) GetParticipatedOn() time.Time`

GetParticipatedOn returns the ParticipatedOn field if non-nil, zero value otherwise.

### GetParticipatedOnOk

`func (o *Participant) GetParticipatedOnOk() (*time.Time, bool)`

GetParticipatedOnOk returns a tuple with the ParticipatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipatedOn

`func (o *Participant) SetParticipatedOn(v time.Time)`

SetParticipatedOn sets ParticipatedOn field to given value.

### HasParticipatedOn

`func (o *Participant) HasParticipatedOn() bool`

HasParticipatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


