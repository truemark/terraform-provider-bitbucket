# Workspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**WorkspaceLinks**](WorkspaceLinks.md) |  | [optional] 
**Uuid** | Pointer to **string** | The workspace&#39;s immutable id. | [optional] 
**Name** | Pointer to **string** | The name of the workspace. | [optional] 
**Slug** | Pointer to **string** | The short label that identifies this workspace. | [optional] 
**IsPrivate** | Pointer to **bool** | Indicates whether the workspace is publicly accessible, or whether it is private to the members and consequently only visible to members. Note that private workspaces cannot contain public repositories. | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**UpdatedOn** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewWorkspace

`func NewWorkspace() *Workspace`

NewWorkspace instantiates a new Workspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceWithDefaults

`func NewWorkspaceWithDefaults() *Workspace`

NewWorkspaceWithDefaults instantiates a new Workspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Workspace) GetLinks() WorkspaceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Workspace) GetLinksOk() (*WorkspaceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Workspace) SetLinks(v WorkspaceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Workspace) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetUuid

`func (o *Workspace) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Workspace) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Workspace) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Workspace) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *Workspace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Workspace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Workspace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Workspace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *Workspace) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Workspace) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Workspace) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Workspace) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetIsPrivate

`func (o *Workspace) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *Workspace) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *Workspace) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *Workspace) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Workspace) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Workspace) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Workspace) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Workspace) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *Workspace) GetUpdatedOn() time.Time`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *Workspace) GetUpdatedOnOk() (*time.Time, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *Workspace) SetUpdatedOn(v time.Time)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *Workspace) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


