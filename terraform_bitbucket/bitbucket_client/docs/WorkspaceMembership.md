# WorkspaceMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**BranchingModelSettingsLinks**](BranchingModelSettingsLinks.md) |  | [optional] 
**User** | Pointer to [**Account**](Account.md) |  | [optional] 
**Workspace** | Pointer to [**Workspace**](Workspace.md) |  | [optional] 

## Methods

### NewWorkspaceMembership

`func NewWorkspaceMembership() *WorkspaceMembership`

NewWorkspaceMembership instantiates a new WorkspaceMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMembershipWithDefaults

`func NewWorkspaceMembershipWithDefaults() *WorkspaceMembership`

NewWorkspaceMembershipWithDefaults instantiates a new WorkspaceMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *WorkspaceMembership) GetLinks() BranchingModelSettingsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WorkspaceMembership) GetLinksOk() (*BranchingModelSettingsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WorkspaceMembership) SetLinks(v BranchingModelSettingsLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WorkspaceMembership) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetUser

`func (o *WorkspaceMembership) GetUser() Account`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *WorkspaceMembership) GetUserOk() (*Account, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *WorkspaceMembership) SetUser(v Account)`

SetUser sets User field to given value.

### HasUser

`func (o *WorkspaceMembership) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetWorkspace

`func (o *WorkspaceMembership) GetWorkspace() Workspace`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *WorkspaceMembership) GetWorkspaceOk() (*Workspace, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *WorkspaceMembership) SetWorkspace(v Workspace)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *WorkspaceMembership) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


