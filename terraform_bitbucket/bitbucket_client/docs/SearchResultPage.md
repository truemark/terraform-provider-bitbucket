# SearchResultPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int64** |  | [optional] [readonly] 
**Page** | Pointer to **int32** |  | [optional] [readonly] 
**Pagelen** | Pointer to **int32** |  | [optional] [readonly] 
**QuerySubstituted** | Pointer to **bool** |  | [optional] [readonly] 
**Next** | Pointer to **string** |  | [optional] [readonly] 
**Previous** | Pointer to **string** |  | [optional] [readonly] 
**Values** | Pointer to [**[]SearchCodeSearchResult**](SearchCodeSearchResult.md) |  | [optional] [readonly] 

## Methods

### NewSearchResultPage

`func NewSearchResultPage() *SearchResultPage`

NewSearchResultPage instantiates a new SearchResultPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultPageWithDefaults

`func NewSearchResultPageWithDefaults() *SearchResultPage`

NewSearchResultPageWithDefaults instantiates a new SearchResultPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *SearchResultPage) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SearchResultPage) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SearchResultPage) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *SearchResultPage) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetPage

`func (o *SearchResultPage) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SearchResultPage) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SearchResultPage) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *SearchResultPage) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPagelen

`func (o *SearchResultPage) GetPagelen() int32`

GetPagelen returns the Pagelen field if non-nil, zero value otherwise.

### GetPagelenOk

`func (o *SearchResultPage) GetPagelenOk() (*int32, bool)`

GetPagelenOk returns a tuple with the Pagelen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagelen

`func (o *SearchResultPage) SetPagelen(v int32)`

SetPagelen sets Pagelen field to given value.

### HasPagelen

`func (o *SearchResultPage) HasPagelen() bool`

HasPagelen returns a boolean if a field has been set.

### GetQuerySubstituted

`func (o *SearchResultPage) GetQuerySubstituted() bool`

GetQuerySubstituted returns the QuerySubstituted field if non-nil, zero value otherwise.

### GetQuerySubstitutedOk

`func (o *SearchResultPage) GetQuerySubstitutedOk() (*bool, bool)`

GetQuerySubstitutedOk returns a tuple with the QuerySubstituted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerySubstituted

`func (o *SearchResultPage) SetQuerySubstituted(v bool)`

SetQuerySubstituted sets QuerySubstituted field to given value.

### HasQuerySubstituted

`func (o *SearchResultPage) HasQuerySubstituted() bool`

HasQuerySubstituted returns a boolean if a field has been set.

### GetNext

`func (o *SearchResultPage) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *SearchResultPage) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *SearchResultPage) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *SearchResultPage) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *SearchResultPage) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *SearchResultPage) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *SearchResultPage) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *SearchResultPage) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetValues

`func (o *SearchResultPage) GetValues() []SearchCodeSearchResult`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SearchResultPage) GetValuesOk() (*[]SearchCodeSearchResult, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SearchResultPage) SetValues(v []SearchCodeSearchResult)`

SetValues sets Values field to given value.

### HasValues

`func (o *SearchResultPage) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


