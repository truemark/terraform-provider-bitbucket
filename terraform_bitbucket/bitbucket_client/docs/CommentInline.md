# CommentInline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | Pointer to **int32** | The comment&#39;s anchor line in the new version of the file. If the &#39;from&#39; line is also provided, this value will be removed. | [optional] 
**From** | Pointer to **int32** | The comment&#39;s anchor line in the old version of the file. | [optional] 
**Path** | **string** | The path of the file this comment is anchored to. | 

## Methods

### NewCommentInline

`func NewCommentInline(path string, ) *CommentInline`

NewCommentInline instantiates a new CommentInline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentInlineWithDefaults

`func NewCommentInlineWithDefaults() *CommentInline`

NewCommentInlineWithDefaults instantiates a new CommentInline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *CommentInline) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CommentInline) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CommentInline) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *CommentInline) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetFrom

`func (o *CommentInline) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CommentInline) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CommentInline) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CommentInline) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetPath

`func (o *CommentInline) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CommentInline) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CommentInline) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


