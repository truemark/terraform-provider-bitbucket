# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**GroupLinks**](GroupLinks.md) |  | [optional] 
**Owner** | Pointer to [**Account**](Account.md) |  | [optional] 
**Workspace** | Pointer to [**Workspace**](Workspace.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** | The \&quot;sluggified\&quot; version of the group&#39;s name. This contains only ASCII characters and can therefore be slightly different than the name | [optional] 
**FullSlug** | Pointer to **string** | The concatenation of the workspace&#39;s slug and the group&#39;s slug, separated with a colon (e.g. &#x60;acme:developers&#x60;)  | [optional] 

## Methods

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Group) GetLinks() GroupLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Group) GetLinksOk() (*GroupLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Group) SetLinks(v GroupLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Group) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetOwner

`func (o *Group) GetOwner() Account`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Group) GetOwnerOk() (*Account, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Group) SetOwner(v Account)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Group) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetWorkspace

`func (o *Group) GetWorkspace() Workspace`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *Group) GetWorkspaceOk() (*Workspace, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *Group) SetWorkspace(v Workspace)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *Group) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.

### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Group) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *Group) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Group) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Group) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *Group) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetFullSlug

`func (o *Group) GetFullSlug() string`

GetFullSlug returns the FullSlug field if non-nil, zero value otherwise.

### GetFullSlugOk

`func (o *Group) GetFullSlugOk() (*string, bool)`

GetFullSlugOk returns a tuple with the FullSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullSlug

`func (o *Group) SetFullSlug(v string)`

SetFullSlug sets FullSlug field to given value.

### HasFullSlug

`func (o *Group) HasFullSlug() bool`

HasFullSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


