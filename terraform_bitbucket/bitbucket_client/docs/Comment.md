# Comment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CreatedOn** | Pointer to **time.Time** |  | [optional] 
**UpdatedOn** | Pointer to **time.Time** |  | [optional] 
**Content** | Pointer to [**IssueContent**](IssueContent.md) |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Parent** | Pointer to [**Comment**](Comment.md) |  | [optional] 
**Inline** | Pointer to [**CommentInline**](CommentInline.md) |  | [optional] 
**Links** | Pointer to [**CommentLinks**](CommentLinks.md) |  | [optional] 

## Methods

### NewComment

`func NewComment() *Comment`

NewComment instantiates a new Comment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentWithDefaults

`func NewCommentWithDefaults() *Comment`

NewCommentWithDefaults instantiates a new Comment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Comment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Comment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Comment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Comment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedOn

`func (o *Comment) GetCreatedOn() time.Time`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *Comment) GetCreatedOnOk() (*time.Time, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *Comment) SetCreatedOn(v time.Time)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *Comment) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *Comment) GetUpdatedOn() time.Time`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *Comment) GetUpdatedOnOk() (*time.Time, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *Comment) SetUpdatedOn(v time.Time)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *Comment) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.

### GetContent

`func (o *Comment) GetContent() IssueContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Comment) GetContentOk() (*IssueContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Comment) SetContent(v IssueContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *Comment) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetUser

`func (o *Comment) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Comment) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Comment) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *Comment) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetDeleted

`func (o *Comment) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Comment) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Comment) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Comment) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetParent

`func (o *Comment) GetParent() Comment`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Comment) GetParentOk() (*Comment, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Comment) SetParent(v Comment)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Comment) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetInline

`func (o *Comment) GetInline() CommentInline`

GetInline returns the Inline field if non-nil, zero value otherwise.

### GetInlineOk

`func (o *Comment) GetInlineOk() (*CommentInline, bool)`

GetInlineOk returns a tuple with the Inline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInline

`func (o *Comment) SetInline(v CommentInline)`

SetInline sets Inline field to given value.

### HasInline

`func (o *Comment) HasInline() bool`

HasInline returns a boolean if a field has been set.

### GetLinks

`func (o *Comment) GetLinks() CommentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Comment) GetLinksOk() (*CommentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Comment) SetLinks(v CommentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Comment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


