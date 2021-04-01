# SearchLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line** | Pointer to **int32** |  | [optional] [readonly] 
**Segments** | Pointer to [**[]SearchSegment**](SearchSegment.md) |  | [optional] [readonly] 

## Methods

### NewSearchLine

`func NewSearchLine() *SearchLine`

NewSearchLine instantiates a new SearchLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchLineWithDefaults

`func NewSearchLineWithDefaults() *SearchLine`

NewSearchLineWithDefaults instantiates a new SearchLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine

`func (o *SearchLine) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *SearchLine) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *SearchLine) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *SearchLine) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetSegments

`func (o *SearchLine) GetSegments() []SearchSegment`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *SearchLine) GetSegmentsOk() (*[]SearchSegment, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *SearchLine) SetSegments(v []SearchSegment)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *SearchLine) HasSegments() bool`

HasSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


