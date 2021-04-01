# SearchSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** |  | [optional] [readonly] 
**Match** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewSearchSegment

`func NewSearchSegment() *SearchSegment`

NewSearchSegment instantiates a new SearchSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSegmentWithDefaults

`func NewSearchSegmentWithDefaults() *SearchSegment`

NewSearchSegmentWithDefaults instantiates a new SearchSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *SearchSegment) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SearchSegment) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SearchSegment) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *SearchSegment) HasText() bool`

HasText returns a boolean if a field has been set.

### GetMatch

`func (o *SearchSegment) GetMatch() bool`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *SearchSegment) GetMatchOk() (*bool, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *SearchSegment) SetMatch(v bool)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *SearchSegment) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


