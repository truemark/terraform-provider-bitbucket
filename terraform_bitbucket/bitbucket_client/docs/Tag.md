# Tag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | The message associated with the tag, if available. | [optional] 
**Date** | Pointer to **time.Time** | The date that the tag was created, if available | [optional] 
**Tagger** | Pointer to [**Author**](Author.md) |  | [optional] 

## Methods

### NewTag

`func NewTag() *Tag`

NewTag instantiates a new Tag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagWithDefaults

`func NewTagWithDefaults() *Tag`

NewTagWithDefaults instantiates a new Tag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *Tag) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Tag) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Tag) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Tag) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDate

`func (o *Tag) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Tag) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Tag) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *Tag) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetTagger

`func (o *Tag) GetTagger() Author`

GetTagger returns the Tagger field if non-nil, zero value otherwise.

### GetTaggerOk

`func (o *Tag) GetTaggerOk() (*Author, bool)`

GetTaggerOk returns a tuple with the Tagger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagger

`func (o *Tag) SetTagger(v Author)`

SetTagger sets Tagger field to given value.

### HasTagger

`func (o *Tag) HasTagger() bool`

HasTagger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


