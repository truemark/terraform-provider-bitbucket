# PaginatedDiffstats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int32** | Total number of objects in the response. This is an optional element that is not provided in all responses, as it can be expensive to compute. | [optional] 
**Page** | Pointer to **int32** | Page number of the current results. This is an optional element that is not provided in all responses. | [optional] 
**Pagelen** | Pointer to **int32** | Current number of objects on the existing page. The default value is 10 with 100 being the maximum allowed value. Individual APIs may enforce different values. | [optional] 
**Next** | Pointer to **string** | Link to the next page if it exists. The last page of a collection does not have this value. Use this link to navigate the result set and refrain from constructing your own URLs. | [optional] 
**Previous** | Pointer to **string** | Link to previous page if it exists. A collections first page does not have this value. This is an optional element that is not provided in all responses. Some result sets strictly support forward navigation and never provide previous links. Clients must anticipate that backwards navigation is not always available. Use this link to navigate the result set and refrain from constructing your own URLs. | [optional] 
**Values** | Pointer to [**[]Diffstat**](Diffstat.md) |  | [optional] 

## Methods

### NewPaginatedDiffstats

`func NewPaginatedDiffstats() *PaginatedDiffstats`

NewPaginatedDiffstats instantiates a new PaginatedDiffstats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedDiffstatsWithDefaults

`func NewPaginatedDiffstatsWithDefaults() *PaginatedDiffstats`

NewPaginatedDiffstatsWithDefaults instantiates a new PaginatedDiffstats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *PaginatedDiffstats) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PaginatedDiffstats) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PaginatedDiffstats) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PaginatedDiffstats) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetPage

`func (o *PaginatedDiffstats) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PaginatedDiffstats) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PaginatedDiffstats) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PaginatedDiffstats) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPagelen

`func (o *PaginatedDiffstats) GetPagelen() int32`

GetPagelen returns the Pagelen field if non-nil, zero value otherwise.

### GetPagelenOk

`func (o *PaginatedDiffstats) GetPagelenOk() (*int32, bool)`

GetPagelenOk returns a tuple with the Pagelen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagelen

`func (o *PaginatedDiffstats) SetPagelen(v int32)`

SetPagelen sets Pagelen field to given value.

### HasPagelen

`func (o *PaginatedDiffstats) HasPagelen() bool`

HasPagelen returns a boolean if a field has been set.

### GetNext

`func (o *PaginatedDiffstats) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedDiffstats) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedDiffstats) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedDiffstats) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *PaginatedDiffstats) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedDiffstats) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedDiffstats) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedDiffstats) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetValues

`func (o *PaginatedDiffstats) GetValues() []Diffstat`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PaginatedDiffstats) GetValuesOk() (*[]Diffstat, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PaginatedDiffstats) SetValues(v []Diffstat)`

SetValues sets Values field to given value.

### HasValues

`func (o *PaginatedDiffstats) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


