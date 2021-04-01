# IssueChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Links** | Pointer to [**IssueChangeLinks**](IssueChangeLinks.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Issue** | Pointer to [**Issue**](Issue.md) |  | [optional] 
**Changes** | Pointer to [**IssueChangeChanges**](IssueChangeChanges.md) |  | [optional] 
**Message** | Pointer to [**IssueContent**](IssueContent.md) |  | [optional] 

## Methods

### NewIssueChange

`func NewIssueChange(type_ string, ) *IssueChange`

NewIssueChange instantiates a new IssueChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueChangeWithDefaults

`func NewIssueChangeWithDefaults() *IssueChange`

NewIssueChangeWithDefaults instantiates a new IssueChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IssueChange) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IssueChange) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IssueChange) SetType(v string)`

SetType sets Type field to given value.


### GetLinks

`func (o *IssueChange) GetLinks() IssueChangeLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IssueChange) GetLinksOk() (*IssueChangeLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IssueChange) SetLinks(v IssueChangeLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IssueChange) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *IssueChange) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueChange) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueChange) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueChange) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedOn

`func (o *IssueChange) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *IssueChange) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *IssueChange) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *IssueChange) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUser

`func (o *IssueChange) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IssueChange) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IssueChange) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *IssueChange) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetIssue

`func (o *IssueChange) GetIssue() Issue`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *IssueChange) GetIssueOk() (*Issue, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *IssueChange) SetIssue(v Issue)`

SetIssue sets Issue field to given value.

### HasIssue

`func (o *IssueChange) HasIssue() bool`

HasIssue returns a boolean if a field has been set.

### GetChanges

`func (o *IssueChange) GetChanges() IssueChangeChanges`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *IssueChange) GetChangesOk() (*IssueChangeChanges, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *IssueChange) SetChanges(v IssueChangeChanges)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *IssueChange) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### GetMessage

`func (o *IssueChange) GetMessage() IssueContent`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *IssueChange) GetMessageOk() (*IssueContent, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *IssueChange) SetMessage(v IssueContent)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *IssueChange) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


